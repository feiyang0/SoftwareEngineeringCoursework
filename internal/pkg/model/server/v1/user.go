package v1

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID uint64 `json:"schoolId,omitempty"        gorm:"primary_key;AUTO_INCREMENT;column:id"`
	// Required: true
	Username string `json:"username" gorm:"column:username;not null;size:32" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password,omitempty" gorm:"column:password;not null;size:128" binding:"required" validate:"min=5,max=128"`
	Email    string `json:"email" gorm:"column:email;size:100" binding:"required" validate:"required,email,min=1,max=100"`
	Role     int    `json:"role" gorm:"column:role;size:10;" binding:"required" validate:"required"`
	Gender   string `json:"gender" gorm:"column:gender;size:10" binding:"" validate:""`

	CreatedAt time.Time      `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deletedAt;index:idx_deletedAt"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
