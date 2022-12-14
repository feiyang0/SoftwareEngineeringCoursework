package v1

// Solution 题目答案
type Solution struct {
	BaseModel
	SchoolId uint64    `json:"schoolId" gorm:"column:schoolId"`
	Content  string    `json:"content" gorm:"column:content"`
	Title    string    `json:"title" gorm:"column:title"`
	Comments []Comment `json:"comments" gorm:"many2many:solution_comment"`
}

func (a *Solution) TableName() string {
	return "solution"
}

type ProblemSolution struct {
	ProblemId  uint64 `gorm:"primary_key"`
	SolutionId uint64 `gorm:"primary_key"`
}

func (a *ProblemSolution) TableName() string {
	return "problem_solution"
}

type UserSolution struct {
}
