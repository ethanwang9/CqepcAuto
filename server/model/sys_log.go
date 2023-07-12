package model

type SysLog struct {
	ID      string `gorm:"column:id;primaryKey;comment:编号" json:"id,omitempty"`
	UID     string `gorm:"column:uid;comment:用户ID" json:"uid,omitempty"`
	Action  string `gorm:"column:action;comment:操作" json:"action,omitempty"`
	Details string `gorm:"column:details;comment:详细内容" json:"details,omitempty"`
	Common
}
