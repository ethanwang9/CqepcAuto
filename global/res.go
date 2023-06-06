package global

// 请求返回

// ResLogin 账号密码登录返回
type ResLogin struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token      string `json:"token"`
		Id         string `json:"id"`
		Username   string `json:"username"`
		Nickname   string `json:"nickname"`
		HasStudent int    `json:"has_student"`
		HasCadre   int    `json:"has_cadre"`
		Openid     string `json:"openid"`
		ClassName  string `json:"class_name"`
		ClassCode  string `json:"class_code"`
	} `json:"data"`
}

// ResClassAll 本学期课表返回
type ResClassAll struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		CourseCode string `json:"course_code"`
		CourseName string `json:"course_name"`
	} `json:"data"`
}

// ResClassAllDetails 课程评课数据
type ResClassAllDetails struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TotalClassroom  int     `json:"totalClassroom"`
		StartTime       string  `json:"startTime"`
		EndTime         string  `json:"endTime"`
		Join            float64 `json:"join"`
		TotalEvaluation int     `json:"totalEvaluation"`
		JoinStr         string  `json:"joinStr"`
	} `json:"data"`
}

// ResTodayClass 今日课表返回
type ResTodayClass struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Id             string      `json:"id"`
		CourseName     string      `json:"course_name"`
		CourseCode     string      `json:"course_code"`
		ClassCode      string      `json:"class_code"`
		ClassClassromm string      `json:"class_classromm"`
		ClassCategory  string      `json:"class_category"`
		ClassMajor     string      `json:"class_major"`
		ClassName      string      `json:"class_name"`
		ClassNum       string      `json:"class_num"`
		NumOfPeople    int         `json:"num_of_people"`
		TeacherName    string      `json:"teacher_name"`
		TeacherCode    string      `json:"teacher_code"`
		College        string      `json:"college"`
		OpenCourseDep  string      `json:"open_course_dep"`
		CampusName     string      `json:"campus_name"`
		ChooseCourseNo string      `json:"choose_course_no"`
		TeachTime      string      `json:"teach_time"`
		Week           int         `json:"week"`
		WeekBegin      string      `json:"week_begin"`
		WeekEnd        string      `json:"week_end"`
		WeekDay        int         `json:"week_day"`
		Node           string      `json:"node"`
		NodeBegin      string      `json:"node_begin"`
		NodeEnd        string      `json:"node_end"`
		Version        int         `json:"version"`
		StartEndWeeks  interface{} `json:"start_end_weeks"`
		SadWeeks       interface{} `json:"sad_weeks"`
		WeekDayStr     interface{} `json:"week_day_str"`
		Eqkey          interface{} `json:"eqkey"`
		RowNum         interface{} `json:"rowNum"`
	} `json:"data"`
}

// ResTodayClassData 今日课表数据返回消息
type ResTodayClassData struct {
	Id             string      `json:"id"`
	CourseName     string      `json:"course_name"`
	CourseCode     string      `json:"course_code"`
	ClassCode      string      `json:"class_code"`
	ClassClassromm string      `json:"class_classromm"`
	ClassCategory  string      `json:"class_category"`
	ClassMajor     string      `json:"class_major"`
	ClassName      string      `json:"class_name"`
	ClassNum       string      `json:"class_num"`
	NumOfPeople    int         `json:"num_of_people"`
	TeacherName    string      `json:"teacher_name"`
	TeacherCode    string      `json:"teacher_code"`
	College        string      `json:"college"`
	OpenCourseDep  string      `json:"open_course_dep"`
	CampusName     string      `json:"campus_name"`
	ChooseCourseNo string      `json:"choose_course_no"`
	TeachTime      string      `json:"teach_time"`
	Week           int         `json:"week"`
	WeekBegin      string      `json:"week_begin"`
	WeekEnd        string      `json:"week_end"`
	WeekDay        int         `json:"week_day"`
	Node           string      `json:"node"`
	NodeBegin      string      `json:"node_begin"`
	NodeEnd        string      `json:"node_end"`
	Version        int         `json:"version"`
	StartEndWeeks  interface{} `json:"start_end_weeks"`
	SadWeeks       interface{} `json:"sad_weeks"`
	WeekDayStr     interface{} `json:"week_day_str"`
	Eqkey          interface{} `json:"eqkey"`
	RowNum         interface{} `json:"rowNum"`
}

// ResPKGet 获取评课数据
type ResPKGet struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		UnitTitle string `json:"unit_title"`
		Id        int    `json:"id"`
	} `json:"data"`
}

// ResPKSend 获取评课返回结果
type ResPKSend struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// PkTj 系统评课统计
type PkTj struct {
	CourseCode string `json:"course_code"`
	CourseName string `json:"course_name"`
	Data       struct {
		TotalClassroom  int     `json:"totalClassroom"`
		StartTime       string  `json:"startTime"`
		EndTime         string  `json:"endTime"`
		Join            float64 `json:"join"`
		TotalEvaluation int     `json:"totalEvaluation"`
		JoinStr         string  `json:"joinStr"`
	} `json:"data"`
}

// ResAppUpdate 更新返回数据
type ResAppUpdate struct {
	Version string `json:"version"`
	Doc     string `json:"doc"`
	Key     string `json:"key"`
}
