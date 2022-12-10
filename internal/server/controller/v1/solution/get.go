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

func (s *SolutionController) AddComment(c *gin.Context) {
	log.L(c).Info("call addComment by id function.")

	var r v1.Comment
	var err error

	if err = c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	idKey := c.GetString(constant.XUserIdKey)
	schoolId, _ := strconv.ParseInt(idKey, 10, 64)
	r.SchoolId = uint64(schoolId)

	sid, _ := strconv.ParseInt(c.Param("solutionId"), 10, 64)

	if err = s.solutionS.AddComment(uint64(sid), &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, errno.OK, nil)
}

func (s *SolutionController) DeleteComment(c *gin.Context) {
	log.L(c).Info("call deleteComment by id function.")
	sid, _ := strconv.ParseInt(c.Param("solutionId"), 10, 64)
	cid, _ := strconv.ParseInt(c.Param("commentId"), 10, 64)
	if err := s.solutionS.DelComment(uint64(sid), uint64(cid)); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, errno.OK, nil)
}
