package v1

type Student struct {
	SchoolId uint64
}

type StudentProblem struct {
	BaseModel
	SchoolId  uint64 `json:"schoolId" gorm:"column:schoolId"`
	ProblemId uint64 `json:"problemId" gorm:"column:problemId"`
	Pass      bool   `json:"pass" gorm:"column:pass"`
	Count     int    `json:"count" gorm:"column:count"`
}

func (s *StudentProblem) TableName() string {
	return "studentProblem"
}

type StudentClass struct {
	SchoolId uint64 `json:"schoolId" gorm:"column:schoolId"`
	ClassId  uint64 `json:"classId" gorm:"column:classId"`
}

func (s *StudentClass) TableName() string {
	return "studentClass"
}
