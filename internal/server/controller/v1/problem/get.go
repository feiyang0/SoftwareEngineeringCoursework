package problem

import (
	"SoftwareEngine/internal/pkg/log"
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
)

func (p *ProblemController) GetTags(c *gin.Context) {
	log.L(c).Info("problem getTags function called.")

	tags, err := p.problemS.GetTags()
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, tags)
}

func (p *ProblemController) GetAll(c *gin.Context) {
	log.L(c).Info("problem get all problem function called.")
	var r v1.ListOption

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	problems, err := p.problemS.GetAll(&r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, problems)
}
func (p *ProblemController) GetProblem(c *gin.Context) {
	log.L(c).Info("problem get problem function called.")

	problem, err := p.problemS.GetProblem(c.Param("problemName"))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, problem)
}
