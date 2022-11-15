package user

import (
	"SoftwareEngine/internal/pkg/email"
	"SoftwareEngine/internal/pkg/log"
	"SoftwareEngine/pkg/auth"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
	"sync"
)

type ResetReq struct {
	Email     string `json:"email" validate:"require,email,min=1,max=100"`
	Captcha   string `json:"captcha" validate:"omitempty,captcha,eq=6"`
	NewPasswd string `json:"newPasswd" validate:"omitempty,min=5,max=128"`
}

var CaptchaStore map[string]string
var once sync.Once

func (u *UserController) GetCaptcha(c *gin.Context) {
	once.Do(func() {
		CaptchaStore = make(map[string]string)
	})
	log.L(c).Info("user getCaptcha function called.")
	var r ResetReq
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	if t, _ := u.userS.GetByEmail(r.Email); t == nil {
		core.WriteResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	Captcha, err := email.SendCaptcha(r.Email)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	CaptchaStore[r.Email] = Captcha
	core.WriteResponse(c, errno.OK, nil)
}

func (u *UserController) CheckCaptcha(c *gin.Context) {
	var r ResetReq
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	if CaptchaStore[r.Email] != r.Captcha {
		core.WriteResponse(c, errno.ErrCaptchaIncorrect, nil)
		return
	}
	core.WriteResponse(c, errno.OK, nil)
}

func (u *UserController) SetNewPasswd(c *gin.Context) {
	var r ResetReq
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	user, err := u.userS.GetByEmail(r.Email)
	if err != nil {
		core.WriteResponse(c, err, nil)
	}
	user.Password, err = auth.Encrypt(r.NewPasswd)
	if err != nil {
		core.WriteResponse(c, errno.ErrEncrypt, nil)
		return
	}
	if err = u.userS.Update(user); err != nil {
		core.WriteResponse(c, err, nil)
	}
	core.WriteResponse(c, errno.OK, nil)
}
