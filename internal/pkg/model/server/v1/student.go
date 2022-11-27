package v1

import "time"

type StudentProblem struct {
	UserId    uint64 `gorm:"primary_key"`
	ProblemId uint64 `gorm:"primary_key"`
	Pass      bool   `json:"pass" gorm:"column:pass"`
	Favour    bool   `json:"favour" gorm:"column:favour"`

	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
}

func (s *StudentProblem) TableName() string {
	return "student_problem"
}

type StudentClass struct {
	SchoolId uint64 `json:"schoolId" gorm:"column:schoolId"`
	ClassId  uint64 `json:"classId" gorm:"column:classId"`
}

func (s *StudentClass) TableName() string {
	return "studentClass"
}
