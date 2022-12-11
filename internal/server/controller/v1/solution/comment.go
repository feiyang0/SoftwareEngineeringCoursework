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
