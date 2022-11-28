package problem

import (
	"SoftwareEngine/internal/pkg/constant"
	"SoftwareEngine/internal/pkg/log"
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (p *ProblemController) Create(c *gin.Context) {
	log.L(c).Info("problem create function called.")

	var r v1.Problem
	var err error

	if err = c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	//r.SchoolId = uint64(c.GetInt(constant.XUserIdKey))
	idKey := c.GetString(constant.XUserIdKey)
	schoolId, _ := strconv.ParseInt(idKey, 10, 64)
	r.SchoolId = uint64(schoolId)

	if err = p.problemS.Create(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, errno.OK, nil)
}
