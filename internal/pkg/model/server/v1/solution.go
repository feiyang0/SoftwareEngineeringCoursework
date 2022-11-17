package v1

// Solution 题目答案
type Solution struct {
	BaseModel
	SchoolId uint64 `json:"schoolId" gorm:"column:schoolId"`
	Content  string `json:"content" gorm:"column:content"`
}

func (a *Solution) TableName() string {
	return "solution"
}

type ProblemSolution struct {
	ProblemId  uint64 `gorm:"column:problemId"`
	SolutionId uint64 `gorm:"column:solutionId"`
}

func (a *ProblemSolution) TableName() string {
	return "problemSolution"
}
