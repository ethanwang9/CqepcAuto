package model

type UserMessage struct {
	UID           string `gorm:"column:uid;primaryKey;comment:用户ID" json:"uid,omitempty"`
	Type          string `gorm:"column:type;comment:消息类型" json:"type,omitempty"`
	PushplusToken string `gorm:"column:pushplus_token;comment:推送加Token" json:"pushplus_token,omitempty"`
	WorkWxWebhook string `gorm:"column:work_wx_webhook;comment:企业微信群机器人" json:"work_wx_webhook,omitempty"`
	Common
}
