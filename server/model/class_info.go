package model

type ClassInfo struct {
	UID             string `gorm:"column:uid;primaryKey;comment:用户ID" json:"uid,omitempty"`
	CourseCode      string `gorm:"column:course_code;comment:用户ID" json:"course_code,omitempty"`
	CourseName      string `gorm:"column:course_name;comment:用户ID" json:"course_name,omitempty"`
	TotalClassroom  string `gorm:"column:total_classroom;comment:用户ID" json:"total_classroom,omitempty"`
	TotalEvaluation string `gorm:"column:total_evaluation;comment:用户ID" json:"total_evaluation,omitempty"`
	StartTime       string `gorm:"column:start_time;comment:用户ID" json:"start_time,omitempty"`
	EndTime         string `gorm:"column:end_time;comment:用户ID" json:"end_time,omitempty"`
	Common
}
