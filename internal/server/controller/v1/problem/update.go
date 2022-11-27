package problem

import (
	"SoftwareEngine/internal/pkg/log"
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (p *ProblemController) Update(c *gin.Context) {
	log.L(c).Info("problem update function called.")

	var r v1.Problem
	var err error

	if err = c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	id, _ := strconv.ParseInt(c.Param("problemId"), 10, 64)
	r.ID = uint64(id)

	if err = p.problemS.Update(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, errno.OK, nil)
}
