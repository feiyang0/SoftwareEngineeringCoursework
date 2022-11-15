package user

import (
	"SoftwareEngine/internal/pkg/log"
	"SoftwareEngine/pkg/core"

	"github.com/gin-gonic/gin"
)

//Test
//@Summary 测试各个模块间信息能否流通
//@Description
//@Tags 测试
//Param id path integer true "test name"
//@Success 200 {string} json
//@Failure 200 {string} json
//@Router /users/test [get]
func (u *UserController) Test(c *gin.Context) {
	log.L(c).Info("test function called.")

	core.WriteResponse(c, nil, "test succeed")
}
