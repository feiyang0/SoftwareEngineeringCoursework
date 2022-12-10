package student

import (
	"SoftwareEngine/internal/pkg/constant"
	"SoftwareEngine/internal/pkg/log"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *StudentController) Commit(c *gin.Context) {
	log.L(c).Info("problem commit function called.")

	uid, _ := strconv.ParseInt(c.GetString(constant.XUserIdKey), 10, 64)
	pid, err := strconv.ParseInt(c.Query("problemId"), 10, 64)
	if err != nil {
		core.WriteResponse(c, errno.ErrProblemIdError, nil)
		return
	}

	if err = s.studentS.Commit(uint64(uid), uint64(pid)); err != nil {
		core.WriteResponse(c, nil, err)
		return
	}
	core.WriteResponse(c, errno.OK, nil)
}

func (s *StudentController) CancelCommit(c *gin.Context) {
	log.L(c).Info("problem cancel commit function called.")

	uid, _ := strconv.ParseInt(c.GetString(constant.XUserIdKey), 10, 64)
	pid, err := strconv.ParseInt(c.Query("problemId"), 10, 64)
	if err != nil {
		core.WriteResponse(c, errno.ErrProblemIdError, nil)
		return
	}

	if err = s.studentS.CancelCommit(uint64(uid), uint64(pid)); err != nil {
		core.WriteResponse(c, nil, err)
		return
	}
	core.WriteResponse(c, errno.OK, nil)
}
