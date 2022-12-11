package solution

import (
	"SoftwareEngine/internal/pkg/log"
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/core"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *SolutionController) List(c *gin.Context) {
	log.L(c).Info("call solution list by id function.")

	var r v1.SolutionListOption
	var err error

	if err = c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	pid, _ := strconv.ParseInt(c.Param("problemId"), 10, 64)
	r.Pid = uint64(pid)
	solutions, sNumber, err := s.solutionS.GetSolutionList(&r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	data := map[string]interface{}{
		"solutionsNumber": sNumber,
		"solutions":       solutions,
	}
	core.WriteResponse(c, nil, data)
}
