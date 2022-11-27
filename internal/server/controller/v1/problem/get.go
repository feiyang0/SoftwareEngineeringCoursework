package problem

import (
	"SoftwareEngine/internal/pkg/constant"
	"SoftwareEngine/internal/pkg/log"
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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
	var r v1.ProblemListOption

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

func (p *ProblemController) GetAllWithTags(c *gin.Context) {
	log.L(c).Info("problem get all problem function called.")
	var r v1.ProblemListOption

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	//uid := uint64(c.GetInt(constant.XUserIdKey))
	schoolId, err := strconv.ParseInt(c.GetString(constant.XUserIdKey), 10, 64)
	uid := uint64(schoolId)

	fmt.Println("problem get id:", uid)

	problems, err := p.problemS.GetAllWithTag(uid, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, problems)
}
