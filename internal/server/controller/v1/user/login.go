package user

import (
	"SoftwareEngine/internal/pkg/log"
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/auth"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"SoftwareEngine/pkg/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// LoginResponse defines the response fields for `/login`.
type LoginResponse struct {
	Token string `json:"token"`
}

// LoginRequest defines the request fields for `/login` .
type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// Login :return a jwt token.
// @Summary Login generates the authentication token
// @Produce  json
// @Param username body string true "id or email"
// @Param password body string true "Password"
// @Success 200 {string} json
// "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Failure 200 {string} json "{"code":errno,"message":"err_msg"}"
// @Router /login [post].
func (u *UserController) Login(c *gin.Context) {
	log.L(c).Info("user create function called.")
	var r LoginRequest
	var err error
	if err = c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	var user *v1.User
	if find := strings.Contains(r.Account, "@"); find {
		user, err = u.userS.GetByEmail(r.Account)
	} else {
		schoolId, _ := strconv.ParseInt(r.Account, 10, 64)
		user, err = u.userS.GetById(uint64(schoolId))
	}

	if err != nil {
		core.WriteResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err = auth.Compare(user.Password, r.Password); err != nil {
		core.WriteResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// use Role-userId Sign the json web token.
	t, err := token.Sign(fmt.Sprintf("%d%d", *user.Role, user.ID))
	//fmt.Println("sign: ", *user.Role, user.ID)
	if err != nil {
		core.WriteResponse(c, errno.ErrToken, nil)
		return
	}

	core.WriteResponse(c, nil, LoginResponse{Token: t})
}
