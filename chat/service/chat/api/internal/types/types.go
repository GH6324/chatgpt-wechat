// Code generated by goctl. DO NOT EDIT.
package types

type ChatReply struct {
	Message string `json:"message"`
}

type ChatReq struct {
	Channel string `json:"channel,optional" validate:"required,oneof=wecom openai" label:"渠道"`
	MSG     string `json:"msg,optional"`
	UserID  string `json:"user_id,optional" validate:"required,max=500" label:"用户标识"`
	AgentID int64  `json:"agent_id,optional" validate:"required" label:"应用标识"`
}

type CustomerChatReply struct {
	Message string `json:"message"`
}

type CustomerChatReq struct {
	MsgID      string `json:"msg_id"`
	Msg        string `json:"msg"`
	CustomerID string `json:"customer_id"`
	OpenKfID   string `json:"open_kf_id"`
}

type UserDetailReply struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserDetailReq struct {
}

type UserLoginReply struct {
	Token string `json:"token"`
}

type UserLoginReq struct {
	Email    string `json:"email,optional" validate:"required,email" label:"邮箱"`
	Password string `json:"password,optional" validate:"required" label:"密码"`
}

type UserRegisterReply struct {
	Message string `json:"message"`
}

type UserRegisterReq struct {
	Email    string `json:"email,optional" validate:"required,email" label:"邮箱"`
	Name     string `json:"name,optional" validate:"required,max=50" label:"用户名"`
	Password string `json:"password,optional" validate:"required" label:"密码"`
}
