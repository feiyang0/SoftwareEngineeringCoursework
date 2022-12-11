package solution

import (
	"SoftwareEngine/internal/pkg/log"
	"SoftwareEngine/pkg/core"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *SolutionController) Get(c *gin.Context) {
	log.L(c).Info("call solution get by id function.")

	var err error
	sid, _ := strconv.ParseInt(c.Param("solutionId"), 10, 64)
	solu, err := s.solutionS.GetSolution(uint64(sid))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, solu)
}
