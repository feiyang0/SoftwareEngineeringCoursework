package v1

import "github.com/go-playground/validator/v10"

// Comment  ans下的讨论，一个答案多个讨论
type Comment struct {
	BaseModel
	SchoolId uint64 `json:"schoolId" gorm:"column:schoolId"`
	Content  string `json:"content" gorm:"column:content"`
	Status   bool   `json:"status" gorm:"column:status"`
	ReplyId  uint64 `json:"reply" gorm:"column:reply"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *Comment) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

type SolutionComment struct {
	SolutionId uint64 `gorm:"primary_key"`
	CommentId  uint64 `gorm:"primary_key"`
}

func (sc *SolutionComment) TableName() string {
	return "solution_comment"
}
