package solution

import (
	"SoftwareEngine/internal/pkg/log"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *SolutionController) Delete(c *gin.Context) {
	log.L(c).Info("call solution delete by id function.")
	pid, _ := strconv.ParseInt(c.Param("problemId"), 10, 64)
	sid, _ := strconv.ParseInt(c.Param("solutionId"), 10, 64)
	if err := s.solutionS.Delete(uint64(pid), uint64(sid)); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, errno.OK, nil)
}
