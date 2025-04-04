package redis

import (
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

const WelcomeCacheKey = "chat:wecome:%d:%s"
const CursorCacheKey = "chat:cursor:%s"
const EmbeddingsCacheKey = "chat:embeddings:%s"
const UserSessionAgentDefaultKey = "session_agent:default:%s:%s"
const UserSessionListKey = "user:session:list:%s"
const SessionKey = "session:%s"

// DifyConversationKey Dify会话ID存储
const DifyConversationKey = "dify:conversation:%d:%s"

// DifyCustomerConversationKey Dify客户会话ID存储
const DifyCustomerConversationKey = "dify:conversation:%s:%s"

// ImageTemporaryKey 图片临时存储
const ImageTemporaryKey = "chat:image:temporary:%d-%s"

func Init(Host, Pass string) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Host,
		Password: Pass,
		DB:       1,
	})
}

func Close() {
	err := Rdb.Close()
	if err != nil {
		return
	}
}
