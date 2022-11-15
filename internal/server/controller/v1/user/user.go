package user

import (
	"SoftwareEngine/internal/server/store"
)

type UserController struct {
	userS store.UserStore
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{
		userS: store.Users(),
	}
}
