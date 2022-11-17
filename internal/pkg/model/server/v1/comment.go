package v1

import "github.com/go-playground/validator/v10"

// Comment  ans下的讨论，一个答案多个讨论
type Comment struct {
	BaseModel
	AnsId   uint64 `json:"ansId" gorm:"column:ansId"`
	UserId  uint64 `json:"userId" gorm:"column:userId"`
	Content string `json:"content" gorm:"column:content"`
	Status  bool   `json:"status" gorm:"column:status"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *Comment) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
