package solution

import (
	"SoftwareEngine/internal/pkg/constant"
	"SoftwareEngine/internal/pkg/log"
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *SolutionController) Create(c *gin.Context) {
	log.L(c).Info("call solution create by id function.")

	var r v1.Solution
	var err error

	if err = c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	idKey := c.GetString(constant.XUserIdKey)
	schoolId, _ := strconv.ParseInt(idKey, 10, 64)
	pid, _ := strconv.ParseInt(c.Param("problemId"), 10, 64)
	r.SchoolId = uint64(schoolId)

	if err = s.solutionS.Create(uint64(pid), &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, errno.OK, nil)

}
