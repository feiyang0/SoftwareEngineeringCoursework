package student

import (
	"SoftwareEngine/internal/pkg/constant"
	"SoftwareEngine/internal/pkg/log"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *StudentController) Collect(c *gin.Context) {
	log.L(c).Info("problem collect function called.")

	uid, _ := strconv.ParseInt(c.GetString(constant.XUserIdKey), 10, 64)
	pid, err := strconv.ParseInt(c.Query("problemId"), 10, 64)
	if err != nil {
		core.WriteResponse(c, errno.ErrProblemIdError, nil)
	}
	if err = s.studentS.Collect(uint64(uid), uint64(pid)); err != nil {
		core.WriteResponse(c, nil, err)
	}
	core.WriteResponse(c, errno.OK, nil)
}

func (s *StudentController) CancelCollect(c *gin.Context) {
	log.L(c).Info("problem cancel collect function called.")

	uid, _ := strconv.ParseInt(c.GetString(constant.XUserIdKey), 10, 64)
	pid, err := strconv.ParseInt(c.Query("problemId"), 10, 64)
	if err != nil {
		core.WriteResponse(c, errno.ErrProblemIdError, nil)
	}
	if err = s.studentS.CancelCollect(uint64(uid), uint64(pid)); err != nil {
		core.WriteResponse(c, nil, err)
	}
	core.WriteResponse(c, errno.OK, nil)
}
