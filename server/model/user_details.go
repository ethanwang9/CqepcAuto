package model

type UserDetails struct {
	CqepcId   string `gorm:"column:cqepc_id;primaryKey;comment:cqepc学号" json:"cqepc_id,omitempty"`
	Name      string `gorm:"column:name;comment:姓名" json:"name,omitempty"`
	Token     string `gorm:"column:token;comment:登录token" json:"token,omitempty"`
	OpenID    string `gorm:"column:openid;comment:微信小程序Openid" json:"openid,omitempty"`
	Class     string `gorm:"column:class;comment:二级学院" json:"class,omitempty"`
	ClassName string `gorm:"column:class_name;comment:专业名称" json:"class_name,omitempty"`
	ClassCode string `gorm:"column:class_code;comment:班级号" json:"class_code,omitempty"`
	Common
}
