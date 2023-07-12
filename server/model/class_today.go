package model

type ClassToday struct {
	UUID           string `gorm:"column:uuid;primaryKey;comment:数据库记录号" json:"uuid,omitempty"`
	UID            string `gorm:"column:uid;comment:用户ID" json:"uid,omitempty"`
	Cid            string `gorm:"column:cid;comment:系统课程ID" json:"cid,omitempty"`
	CourseCode     string `gorm:"column:course_code;comment:课程ID" json:"course_code,omitempty"`
	CourseName     string `gorm:"column:course_name;comment:课程名称" json:"course_name,omitempty"`
	ClassClassroom string `gorm:"column:class_classroom;comment:课程教室" json:"class_classroom,omitempty"`
	ClassCategory  string `gorm:"column:class_category;comment:课程类别" json:"class_category,omitempty"`
	TeacherName    string `gorm:"column:teacher_name;comment:老师名称" json:"teacher_name,omitempty"`
	Node           string `gorm:"column:node;comment:课程上课时间节数" json:"node,omitempty"`
	NodeBegin      string `gorm:"column:node_begin;comment:上课时间" json:"node_begin,omitempty"`
	NodeEnd        string `gorm:"column:node_end;comment:下课时间" json:"node_end,omitempty"`
	Common
}
