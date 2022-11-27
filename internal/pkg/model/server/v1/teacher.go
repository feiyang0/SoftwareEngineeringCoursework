package v1

type Teacher struct {
}

type Class struct {
	BaseModel
	// 老师id和科目id
	TeacherId uint64 `json:"schoolId" gorm:"column:schoolId"`
	CourseId  uint64 `json:"courseId" gorm:"column:courseId"`
}

func (c *Class) TableName() string {
	return "class"
}
