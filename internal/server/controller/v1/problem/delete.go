package problem

import (
	"SoftwareEngine/internal/pkg/log"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (p *ProblemController) Delete(c *gin.Context) {
	log.L(c).Info("problem get problem function called.")

	id, _ := strconv.ParseInt(c.Param("problemId"), 10, 64)

	if err := p.problemS.Delete(uint64(id)); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, errno.OK, nil)
}
