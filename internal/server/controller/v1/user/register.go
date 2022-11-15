package user

import (
	"SoftwareEngine/internal/pkg/log"
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/auth"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
)

// Register :return true or false
// @Summary 注册
// @Description 输入账号密码，并选择角色
// @Produce  json
// @Param regRequest body RegRequest true "{"email":"email@qq.com","password":"paswd","username":"name","schoolId":123,"role":1}"
// @Success 200 {string} string "{"code":0,"message":"OK"}"
// @Failure 200 {string} string "{"code":errno,"message":"err_msg"}"
// @Router /users/register [post]
func (u *UserController) Register(c *gin.Context) {

	log.L(c).Info("user create function called.")
	var r v1.User
	var err error

	if err = c.ShouldBind(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, err)
		return
	}
	if err = r.Validate(); err != nil {
		core.WriteResponse(c, errno.ErrValidation, nil)
		return
	}
	// 检查用户是否存在
	if temp, _ := u.userS.GetById(r.ID); temp != nil {
		core.WriteResponse(c, errno.ErrUserAlreadyExist, nil)
		return
	}
	// 加密密码并存放信息
	r.Password, err = auth.Encrypt(r.Password)
	if err != nil {
		core.WriteResponse(c, errno.ErrEncrypt, nil)
		return
	}
	if err = u.userS.Create(&r); err != nil {
		core.WriteResponse(c, err, nil)
	}
	core.WriteResponse(c, nil, r)
}

//https://app.getpostman.com/join-team?invite_code=1dcb88bc2ca88cc8fc814e164c9bc629

// @Param email body string true "email"
// @Param password body string true "password"
// @Param username body string true "username"
// @Param schoolId body string true "schoolId"
// @Param role body int true "role"
