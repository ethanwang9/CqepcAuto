package model

type UserBase struct {
	UID       string `gorm:"column:uid;primaryKey;comment:用户ID" json:"uid,omitempty"`
	Password  string `gorm:"column:password;comment:面板密码" json:"password,omitempty"`
	CqepcId   string `gorm:"column:cqepc_id;uniqueIndex;comment:cqepc学号" json:"cqepc_id,omitempty"`
	CqepcPass string `gorm:"column:cqepc_pass;comment:cqepc密码" json:"cqepc_pass,omitempty"`
	Auth      int    `gorm:"column:auth;comment:权限" json:"auth,omitempty"`
	Avatar    string `gorm:"column:avatar;comment:头像" json:"avatar,omitempty"`
	IsStop    bool   `gorm:"column:is_stop;comment:是否停用" json:"is_stop,omitempty"`
	Common
}
