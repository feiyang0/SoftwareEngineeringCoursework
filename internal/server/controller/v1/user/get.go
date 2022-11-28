package user

import (
	"SoftwareEngine/internal/pkg/constant"
	"SoftwareEngine/internal/pkg/log"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

type userInfo struct {
	SchoolId uint64 `json:"schoolId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Gender   string `json:"gender"`
}

func (u *UserController) Get(c *gin.Context) {
	log.L(c).Info("call user get function.")

	schoolId, _ := strconv.ParseInt(c.GetString(constant.XUserIdKey), 10, 64)
	uid := uint64(schoolId)

	user, err := u.userS.GetById(uid)
	if err != nil {
		core.WriteResponse(c, errno.ErrUserNotFound, nil)
	}
	info := userInfo{
		SchoolId: uid,
		Username: user.Username,
		Email:    user.Email,
		Role:     *user.Role,
		Gender:   user.Gender,
	}
	core.WriteResponse(c, nil, info)
}
