package model

type UserAuth struct {
	UID string `gorm:"column:uid;primaryKey;comment:用户ID" json:"uid,omitempty"`
	QQ  string `gorm:"column:qq;uniqueIndex;comment:QQ" json:"qq,omitempty"`
	Wx  string `gorm:"column:wx;uniqueIndex;comment:微信" json:"wx,omitempty"`
	Common
}
