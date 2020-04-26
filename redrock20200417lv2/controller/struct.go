package controller

type Spider struct {
	url    string
	header map[string]string
}

type Spiders struct {
	url    string
	//header map[string]string
}

//豆瓣电影
type Movie struct {
	Id    uint     `gorm:"primary_key"`
	Name  string   //名字
	P     string   //评语
	Num   string   //评论数量
	Score float64  //评分
	Jpg   string   //图片地址
	D     string   //导演
	Ok    int
}

//type Person struct {
//	Stu     string
//	Xh      int
//	Name    string  //课程名
//	Class   string  //编号
//	Bx      string  //必修/选修/重修
//	Status  string  //课程状态
//	Time    string  //时间
//	Where   string  //地点
//	Teacher string  //老师
//}

//学生课程
type Student struct {
	Stu     string
	Xh      int      `gorm:"primary_key"`
	L1      string
	L2      string
	L3      string
	L4      string
	L5      string
	L6      string
	L7      string
	L8      string
	L9      string
	L10     string
	L11     string
	L12     string
	L13     string
	L14     string
	L15     string
	L16     string
}

type Class struct {
	Id      string `gorm:"primary_key"`
	Name    string                       //课程名
	Class   string                       //编号
	Bx      string                       //必修/选修/重修
	Status  string                       //课程状态
	Time    string                       //时间
	Where   string                       //地点
	Teacher string                       //老师
}
