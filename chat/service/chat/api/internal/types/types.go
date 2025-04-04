// Code generated by goctl. DO NOT EDIT.
package types

type BotChatContent struct {
	MimeType string `json:"mime_type" label:"消息类型"`
	Data     string `json:"data" label:"消息内容"`
}

type BotChatHistoryClearReply struct {
}

type BotChatHistoryClearReq struct {
	BotID int64 `json:"bot_id" validate:"required,min=1" label:"机器人ID"`
}

type BotChatHistoryReply struct {
	List []*BotChatWholeReply `json:"list" label:"消息列表"`
}

type BotChatHistoryReq struct {
	BotID int64 `json:"bot_id" validate:"required,min=1" label:"机器人ID"`
}

type BotChatReply struct {
	MessageID string `json:"message_id" label:"消息标识"`
}

type BotChatReq struct {
	MSG   string `json:"msg" validate:"required" label:"消息"`
	BotID int64  `json:"bot_id" validate:"required,min=1" label:"机器人ID"`
}

type BotChatWholeReply struct {
	Role    string         `json:"role" label:"角色"`
	Content BotChatContent `json:"content" label:"消息内容"`
}

type BotCreateReply struct {
	ID string `json:"id" label:"机器人ID"`
}

type BotCreateReq struct {
	Name   string `json:"name" validate:"required,max=50" label:"机器人名称"`
	Avatar string `json:"avatar" validate:"max=255" label:"机器人头像"`
	Desc   string `json:"desc,optional" validate:"max=255" label:"机器人描述"`
}

type BotCustomListDetail struct {
	OpenKfid        string `json:"open_kf_id" label:"客服ID"`
	Name            string `json:"name" label:"客服名称"`
	Avatar          string `json:"avatar" label:"客服头像"`
	ManagePrivilege bool   `json:"manage_privilege" label:"是否有管理权限"`
}

type BotCustomListReply struct {
	Total    int                    `json:"total" label:"总数"`
	Page     int                    `json:"page" label:"页码"`
	PageSize int                    `json:"page_size" label:"每页数量"`
	List     []*BotCustomListDetail `json:"list" label:"机器人列表"`
}

type BotCustomListReq struct {
	Page     int `json:"page" validate:"required,min=1" label:"页码"`
	PageSize int `json:"page_size" validate:"required,min=1" label:"每页数量"`
}

type BotCustomUpdateReply struct {
}

type BotCustomUpdateReq struct {
	BotID    int64  `json:"bot_id" validate:"required,min=1" label:"机器人ID"`
	OpenKfid string `json:"open_kfid" validate:"required" label:"客服ID"`
}

type BotDeleteReply struct {
}

type BotDeleteReq struct {
	ID int64 `json:"id" validate:"required,min=1" label:"机器人ID"`
}

type BotDetailReply struct {
	ID     int64  `json:"id" label:"机器人ID"`
	Name   string `json:"name" label:"机器人名称"`
	Avatar string `json:"avatar" label:"机器人头像"`
	Desc   string `json:"desc" label:"机器人描述"`
	Prompt string `json:"prompt" label:"机器人基础提示词配置"`
}

type BotDetailReq struct {
	ID int64 `json:"id" validate:"required,min=1" label:"机器人ID"`
}

type BotExploreListReply struct {
	Total    int               `json:"total" label:"总数"`
	Page     int               `json:"page" label:"页码"`
	PageSize int               `json:"page_size" label:"每页数量"`
	List     []*ExploreListBot `json:"list" label:"机器人列表"`
}

type BotExploreListReq struct {
	Page     int `json:"page" validate:"required,min=1" label:"页码"`
	PageSize int `json:"page_size" validate:"required,min=1" label:"每页数量"`
}

type BotKnowledgeUpdateReply struct {
}

type BotKnowledgeUpdateReq struct {
	BotID       int64 `json:"bot_id" validate:"required,min=1" label:"机器人ID"`
	KnowledgeID int64 `json:"knowledge_id" validate:"required,min=1" label:"知识库ID"`
}

type BotListDetail struct {
	ID     int64  `json:"id" label:"机器人ID"`
	Name   string `json:"name" label:"机器人名称"`
	Avatar string `json:"avatar" label:"机器人头像"`
	Desc   string `json:"desc" label:"机器人描述"`
}

type BotListReply struct {
	Total    int              `json:"total" label:"总数"`
	Page     int              `json:"page" label:"页码"`
	PageSize int              `json:"page_size" label:"每页数量"`
	List     []*BotListDetail `json:"list" label:"机器人列表"`
}

type BotListReq struct {
	Page     int `json:"page" validate:"required,min=1" label:"页码"`
	PageSize int `json:"page_size" validate:"required,min=1" label:"每页数量"`
}

type BotModelDetailReply struct {
	ModelType   string  `json:"model_type"  label:"模型类型"`
	ModelName   string  `json:"model_name" " label:"模型名称"`
	Temperature float64 `json:"temperature" label:"温度"`
}

type BotModelDetailReq struct {
	BotID int64 `json:"bot_id" validate:"required,min=1" label:"机器人ID"`
}

type BotModelUpdateReply struct {
}

type BotModelUpdateReq struct {
	BotID       int64   `json:"bot_id" validate:"required,min=1" label:"机器人ID"`
	ModelType   string  `json:"model_type" validate:"required,oneof=openai gemini" label:"模型类型"`
	ModelName   string  `json:"model_name" validate:"required" label:"模型名称"`
	Temperature float64 `json:"temperature" validate:"required" label:"温度"`
}

type BotOptimizePromptReply struct {
	MessageID string `json:"message_id" label:"消息标识"`
}

type BotOptimizePromptReq struct {
	BotID        int64  `json:"id" validate:"required,min=1" label:"机器人ID"`
	OriginPrompt string `json:"prompt" validate:"required" label:"机器人基础提示词配置"`
}

type BotPromptUpdateReply struct {
}

type BotPromptUpdateReq struct {
	ID     int64  `json:"id" validate:"required,min=1" label:"机器人ID"`
	Prompt string `json:"prompt" validate:"required" label:"机器人基础提示词配置"`
}

type BotReplicateReply struct {
	ID int64 `json:"id" label:"机器人ID"`
}

type BotReplicateReq struct {
	ID         int64 `json:"id" validate:"required,min=1" label:"机器人ID"`
	OriginType int   `json:"origin_type,optional,default=1" label:"复制类型 1:复制机器人 2: 复制prompt配置"`
}

type BotUpdateReply struct {
}

type BotUpdateReq struct {
	ID     int64  `json:"id" validate:"required,min=1" label:"机器人ID"`
	Name   string `json:"name" validate:"required,max=50" label:"机器人名称"`
	Avatar string `json:"avatar" validate:"max=255" label:"机器人头像"`
	Desc   string `json:"desc,optional" validate:"max=255" label:"机器人描述"`
}

type ChatReply struct {
	Message string `json:"message" label:"消息"`
}

type ChatReq struct {
	Channel string `json:"channel,optional" validate:"required,oneof=wecom openai gemini dify deepseek" label:"渠道"`
	MSG     string `json:"msg,optional" label:"消息"`
	UserID  string `json:"user_id,optional" validate:"required,max=500" label:"用户标识"`
	AgentID int64  `json:"agent_id,optional" validate:"required" label:"应用标识"`
}

type CustomerChatReply struct {
	Message string `json:"message" label:"响应消息"`
}

type CustomerChatReq struct {
	MsgID      string `json:"msg_id" label:"消息标识"`
	Msg        string `json:"msg" label:"消息"`
	CustomerID string `json:"customer_id" label:"客户标识"`
	OpenKfID   string `json:"open_kf_id" label:"客服标识"`
}

type ExploreListBot struct {
	ID  int64  `json:"id" label:"提示词ID"`
	Key string `json:"key" label:"key"`
}

type Knowledge struct {
	ID         int64  `json:"id" label:"知识库ID"`
	Name       string `json:"name" label:"知识库名称"`
	Avatar     string `json:"avatar" label:"知识库头像"`
	Desc       string `json:"desc" label:"知识库描述"`
	CreateTime string `json:"create_time" label:"创建时间"`
	UpdateTime string `json:"update_time" label:"更新时间"`
}

type KnowledgeCreateReply struct {
	ID int64 `json:"id" label:"知识库ID"`
}

type KnowledgeCreateReq struct {
	Name   string `json:"name" validate:"required,max=50" label:"知识库名称"`
	Avatar string `json:"avatar" validate:"max=255" label:"知识库头像"`
	Desc   string `json:"desc,optional" validate:"max=255" label:"知识库描述"`
}

type KnowledgeDeleteReply struct {
}

type KnowledgeDeleteReq struct {
	ID int64 `json:"id" validate:"required" label:"知识库ID"`
}

type KnowledgeListReply struct {
	List  []Knowledge `json:"list" label:"知识库列表"`
	Total int         `json:"total" label:"总数"`
}

type KnowledgeListReq struct {
	Page     int `json:"page" validate:"required,min=1" label:"页码"`
	PageSize int `json:"page_size" validate:"required,min=1,max=100" label:"每页数量"`
}

type KnowledgeSegment struct {
	ID              int64  `json:"id" label:"知识点ID"`
	KnowledgeID     int64  `json:"knowledge_id" label:"知识库ID"`
	KnowledgeUnitID int64  `json:"knowledge_unit_id" label:"知识库单元ID"`
	Value           string `json:"value" label:"知识点内容"`
	CreateTime      string `json:"create_time" label:"创建时间"`
	UpdateTime      string `json:"update_time" label:"更新时间"`
}

type KnowledgeSegmentsCreateReply struct {
	ID int64 `json:"id" label:"知识点ID"`
}

type KnowledgeSegmentsCreateReq struct {
	KnowledgeID     int64  `json:"knowledge_id" validate:"required" label:"知识库ID"`
	KnowledgeUnitID int64  `json:"knowledge_unit_id" validate:"required" label:"知识库单元ID"`
	Value           string `json:"value" validate:"required,max=2000" label:"知识点内容"`
}

type KnowledgeSegmentsDeleteReply struct {
}

type KnowledgeSegmentsDeleteReq struct {
	ID int64 `json:"id" validate:"required" label:"知识点ID"`
}

type KnowledgeSegmentsListReply struct {
	List  []KnowledgeSegment `json:"list" label:"知识点列表"`
	Total int                `json:"total" label:"总数"`
}

type KnowledgeSegmentsListReq struct {
	KnowledgeID     int64 `json:"knowledge_id" validate:"required" label:"知识库ID"`
	KnowledgeUnitID int64 `json:"knowledge_unit_id" validate:"required" label:"知识库单元ID"`
	Page            int   `json:"page" validate:"required,min=1" label:"页码"`
	PageSize        int   `json:"page_size" validate:"required,min=1,max=100" label:"每页数量"`
}

type KnowledgeSegmentsUpdateReply struct {
}

type KnowledgeSegmentsUpdateReq struct {
	ID    int64  `json:"id" validate:"required" label:"知识点ID"`
	Value string `json:"value" validate:"required,max=2000" label:"知识点内容"`
}

type KnowledgeUnit struct {
	ID          int64  `json:"id" label:"知识库单元ID"`
	KnowledgeID int64  `json:"knowledge_id" label:"知识库ID"`
	Name        string `json:"name" label:"知识库单元名称"`
	Type        string `json:"type" label:"知识库单元类型"`
	Source      string `json:"source" label:"知识库单元来源"`
	Enable      bool   `json:"enable" label:"知识库单元开关"`
	CreateTime  string `json:"create_time" label:"创建时间"`
	UpdateTime  string `json:"update_time" label:"更新时间"`
}

type KnowledgeUnitCreateReply struct {
	ID int64 `json:"id" label:"知识库单元ID"`
}

type KnowledgeUnitCreateReq struct {
	KnowledgeID int64  `json:"knowledge_id" validate:"required" label:"知识库ID"`
	Name        string `json:"name" validate:"required,max=191" label:"知识库单元名称"`
}

type KnowledgeUnitDeleteReply struct {
}

type KnowledgeUnitDeleteReq struct {
	ID int64 `json:"id" validate:"required" label:"知识库单元ID"`
}

type KnowledgeUnitDetailReply struct {
	ID          int64  `json:"id" label:"知识库单元ID"`
	KnowledgeID int64  `json:"knowledge_id" label:"知识库ID"`
	Name        string `json:"name" label:"知识库单元名称"`
	Type        string `json:"type" label:"知识库单元类型"`
	Source      string `json:"source" label:"知识库单元来源"`
	Enable      bool   `json:"enable" label:"知识库单元开关"`
	CreateTime  string `json:"create_time" label:"创建时间"`
	UpdateTime  string `json:"update_time" label:"更新时间"`
}

type KnowledgeUnitDetailReq struct {
	ID int64 `json:"id" label:"知识库单元ID"`
}

type KnowledgeUnitListReply struct {
	List  []KnowledgeUnit `json:"list" label:"知识库单元列表"`
	Total int             `json:"total" label:"总数"`
}

type KnowledgeUnitListReq struct {
	KnowledgeID int64 `json:"knowledge_id" validate:"required" label:"知识库ID"`
	Page        int   `json:"page" validate:"required,min=1" label:"页码"`
	PageSize    int   `json:"page_size" validate:"required,min=1,max=100" label:"每页数量"`
}

type KnowledgeUnitSwitchReply struct {
}

type KnowledgeUnitSwitchReq struct {
	ID int64 `json:"id" validate:"required" label:"知识库单元ID"`
}

type KnowledgeUnitUpdateReply struct {
}

type KnowledgeUnitUpdateReq struct {
	ID   int64  `json:"id" validate:"required" label:"知识库单元ID"`
	Name string `json:"name" validate:"required,max=191" label:"知识库单元名称"`
}

type KnowledgeUpdateReply struct {
}

type KnowledgeUpdateReq struct {
	ID     int64  `json:"id" validate:"required" label:"知识库ID"`
	Name   string `json:"name" validate:"required,max=50" label:"知识库名称"`
	Avatar string `json:"avatar" validate:"max=255" label:"知识库头像"`
	Desc   string `json:"desc,optional" validate:"max=255" label:"知识库描述"`
}

type UserDetailReply struct {
	ID      int64  `json:"id" label:"用户id"`
	Email   string `json:"email" label:"邮箱"`
	Name    string `json:"name" label:"用户名"`
	IsAdmin bool   `json:"is_admin" label:"是否是管理员"`
}

type UserDetailReq struct {
}

type UserLoginReply struct {
	Token string `json:"token" label:"认证token"`
}

type UserLoginReq struct {
	Email    string `json:"email,optional" validate:"required,email" label:"邮箱"`
	Password string `json:"password,optional" validate:"required" label:"密码"`
}

type UserRegisterReply struct {
	Message string `json:"message"`
}

type UserRegisterReq struct {
	Avatar   string `json:"avatar,optional"  label:"头像"`
	Email    string `json:"email,optional" validate:"required,email" label:"邮箱"`
	Name     string `json:"name,optional" validate:"required,max=50" label:"用户名"`
	Password string `json:"password,optional" validate:"required" label:"密码"`
}
