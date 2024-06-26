package openai

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	copenai "github.com/sashabaranov/go-openai"
	"golang.org/x/net/proxy"
)

const SystemRole = "system"
const UserRole = "user"
const ModelRole = "assistant"

const TextModel = "text-davinci-003"
const ChatModel = "gpt-3.5-turbo"
const ChatModel4 = "gpt-4"

var TotalToken = 3900
var MaxToken = 2000
var Temperature = 0.8

const MimetypeTextPlain = "text/plain"

type ChatModelMessage struct {
	MessageId string      `json:"message_id"`
	Role      string      `json:"role"`
	Content   ChatContent `json:"content"`
}

func NewChatContent(data ...string) ChatContent {
	if len(data) == 0 {
		return ChatContent{}
	}
	if len(data) == 1 {
		return ChatContent{
			MIMEType: MimetypeTextPlain,
			Data:     data[0],
		}
	}
	return ChatContent{
		Data:     data[0],
		MIMEType: data[1],
	}
}

type ChatContent struct {
	MIMEType string `json:"mime_type"`
	Data     string `json:"data"`
}

type ChatClient struct {
	APIKey        string  `json:"api_key"`
	Origin        string  `json:"origin"`
	Engine        string  `json:"engine"`
	HttpProxy     string  `json:"http_proxy"`
	Socks5Proxy   string  `json:"socks5_proxy"`
	ProxyUserName string  `json:"proxy_user_name"`
	ProxyPassword string  `json:"proxy_password"`
	Model         string  `json:"model"`
	BaseHost      string  `json:"base_host"`
	MaxToken      int     `json:"max_token"`
	TotalToken    int     `json:"total_token"`
	Temperature   float32 `json:"temperature"`
}

func NewChatClient(apiKey string) *ChatClient {
	return &ChatClient{
		APIKey:      apiKey,
		MaxToken:    MaxToken,
		TotalToken:  TotalToken,
		Temperature: float32(Temperature),
	}
}

// WithOrigin 设置origin
func (c *ChatClient) WithOrigin(origin string) *ChatClient {
	c.Origin = origin
	return c
}

// WithEngine 设置engine
func (c *ChatClient) WithEngine(engine string) *ChatClient {
	c.Engine = engine
	return c
}

// WithModel 设置模型
func (c *ChatClient) WithModel(model string) *ChatClient {
	c.Model = model
	return c
}

// WithBaseHost 设置baseHost
func (c *ChatClient) WithBaseHost(baseHost string) *ChatClient {
	c.BaseHost = baseHost
	return c
}

// WithMaxToken 设置最大token数
func (c *ChatClient) WithMaxToken(maxToken int) *ChatClient {
	c.MaxToken = maxToken
	return c
}

// WithTemperature 设置响应灵活程度
func (c *ChatClient) WithTemperature(temperature float32) *ChatClient {
	c.Temperature = temperature
	return c
}

// WithTotalToken 设置总token数
func (c *ChatClient) WithTotalToken(totalToken int) *ChatClient {
	c.TotalToken = totalToken
	return c
}

// WithHttpProxy 设置http代理
func (c *ChatClient) WithHttpProxy(proxyUrl string) *ChatClient {
	c.HttpProxy = proxyUrl
	return c
}

// WithSocks5Proxy 设置socks5代理
func (c *ChatClient) WithSocks5Proxy(proxyUrl string) *ChatClient {
	c.Socks5Proxy = proxyUrl
	return c
}

// WithProxyUserName 设置代理用户名
func (c *ChatClient) WithProxyUserName(userName string) *ChatClient {
	c.ProxyUserName = userName
	return c
}

// WithProxyPassword 设置代理密码
func (c *ChatClient) WithProxyPassword(password string) *ChatClient {
	c.ProxyPassword = password
	return c
}

func (c *ChatClient) buildConfig() copenai.ClientConfig {
	config := copenai.DefaultConfig(c.APIKey)
	if c.Origin == "azure" || c.Origin == "azure_ad" {
		config = copenai.DefaultAzureConfig(c.APIKey, c.BaseHost)

		// 默认只使用 一个部署的模型
		config.AzureModelMapperFunc = func(model string) string {
			//azureModelMapping := map[string]string{
			//	defaultModelType: c.Engine,
			//}
			//return azureModelMapping[model]
			return c.Engine
		}
	}
	if c.HttpProxy != "" {
		proxyUrl, _ := url.Parse(c.HttpProxy)
		if c.ProxyUserName != "" && c.ProxyPassword != "" {
			proxyUrl.User = url.UserPassword(c.ProxyUserName, c.ProxyPassword)
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		config.HTTPClient = &http.Client{
			Transport: transport,
			Timeout:   300 * time.Second,
		}
	} else if c.Socks5Proxy != "" {
		socks5Transport := &http.Transport{}
		auth := proxy.Auth{}
		if c.ProxyUserName != "" && c.ProxyPassword != "" {
			auth.Password = c.ProxyPassword
			auth.User = c.ProxyUserName
		}
		dialer, _ := proxy.SOCKS5("tcp", c.Socks5Proxy, &auth, proxy.Direct)
		socks5Transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		}
		config.HTTPClient = &http.Client{
			Transport: socks5Transport,
			Timeout:   300 * time.Second,
		}
	}

	if c.BaseHost != "" && c.Origin == "open_ai" {
		// trim last slash
		config.BaseURL = strings.TrimRight(c.BaseHost, "/") + "/v1"
	}
	return config
}
