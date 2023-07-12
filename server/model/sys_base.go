package model

type SysBase struct {
	ID        int    `gorm:"column:id;primaryKey;comment:编号" json:"id,omitempty"`
	Name      string `gorm:"column:name;comment:系统名称" json:"name,omitempty"`
	Author    string `gorm:"column:author;comment:版权人" json:"author,omitempty"`
	Url       string `gorm:"column:url;comment:版权人地址" json:"url,omitempty"`
	Installed bool   `gorm:"column:installed;comment:是否安装" json:"installed,omitempty"`
	Common
}
