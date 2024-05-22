package model

type Student struct {
	Id    int64  `db:"id"`
	Name  string `db:"name"`   // 姓名
	StuNo string `db:"stu_no"` // 学号
}

func (Student) TableName() string {
	return "student"
}
