package solution

import (
	"SoftwareEngine/internal/pkg/log"
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *SolutionController) Update(c *gin.Context) {
	log.L(c).Info("call solution update by id function.")
	var r v1.Solution
	var err error

	if err = c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	sid, _ := strconv.ParseInt(c.Param("solutionId"), 10, 64)
	r.ID = uint64(sid)

	if err = s.solutionS.Update(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, errno.OK, nil)
}
