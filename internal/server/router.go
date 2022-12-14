package server

import (
	_ "SoftwareEngine/docs"
	"SoftwareEngine/internal/pkg/constant"
	"SoftwareEngine/internal/pkg/middleware"
	"SoftwareEngine/internal/server/controller/v1/problem"
	"SoftwareEngine/internal/server/controller/v1/solution"
	"SoftwareEngine/internal/server/controller/v1/student"
	"SoftwareEngine/internal/server/controller/v1/user"
	"SoftwareEngine/internal/server/store/mysql"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"
	"SoftwareEngine/pkg/token"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func installMiddleware(g *gin.Engine, mw ...gin.HandlerFunc) {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
}

func authMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		identityKey, err := token.ParseRequest(c)

		// 权限等级: 学生:2 < 老师:1 < admin:0
		if err != nil && identityKey[0:1] < role {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set(constant.XUserIdKey, identityKey[1:])
		c.Next()
	}
}

func installController(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	// swagger api docs
	g.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	storeIns, _ := mysql.GetMySQLFactoryOr()
	// 注册登录api
	{
		userController := user.NewUserController(storeIns)

		g.POST("/login", userController.Login)       // 登录
		g.POST("/register", userController.Register) // 注册
		// 密码找回：使用验证码重置密码
		g.GET("/resetPasswd", userController.GetCaptcha)
		g.POST("/resetPasswd", userController.CheckCaptcha)
		g.PUT("/resetPasswd", userController.SetNewPasswd)

		u := g.Group("user", authMiddleware("2"))
		u.GET("info", userController.Get)
		u.GET(":userId", userController.GetById)
	}
	v1 := g.Group("/v1")

	// 题目模块
	{
		problemH := problem.NewProblemController(storeIns)
		stuH := student.NewStudentController(storeIns)
		p := v1.Group("/problem", authMiddleware("2"))
		{

			p.GET("tag", problemH.GetTags)
			p.POST("all", problemH.GetAll)
			p.POST("allWithTag", problemH.GetAllWithTags)
			p.GET(":problemId", problemH.GetProblem)

			p.GET("commit", stuH.Commit)
			p.DELETE("commit", stuH.CancelCommit)

			p.GET("collection", stuH.Collect)
			p.DELETE("collection", stuH.CancelCollect)

		}

		// 题目修改,添加,删除
		teacherG := v1.Group("/problem", authMiddleware("1"))
		{
			teacherG.POST("create", problemH.Create)
			teacherG.PUT(":problemId", problemH.Update)
			teacherG.DELETE(":problemId", problemH.Delete)
		}

		// 题解评论
		solutionG := p.Group(":problemId/solution", authMiddleware("2"))
		{
			solutionH := solution.NewSolutionController(storeIns)
			solutionG.POST("", solutionH.Create)
			solutionG.GET(":solutionId", solutionH.Get)
			solutionG.DELETE(":solutionId", solutionH.Delete)
			solutionG.PUT(":solutionId", solutionH.Update)
			solutionG.POST("all", solutionH.List)

			commentG := solutionG.Group(":solutionId/comment", authMiddleware("2"))
			commentG.POST("", solutionH.AddComment)
			commentG.DELETE(":commentId", solutionH.DeleteComment)
		}
	}

	// 班级模块
	{
		//stuClassH := v1.Group("/class")
		//{
		//
		//	stuClassH.GET(":classId")
		//	stuClassH.POST(":classId") // 加入班级
		//	stuClassH.DELETE(":classId")
		//}
		//
		//teacherClassH := v1.Group("class")
		//{
		//	teacherClassH.GET("/all")
		//}
	}
	// 通知
	// 管理员账号管理
}

func loadRouter(g *gin.Engine, mw ...gin.HandlerFunc) {
	installMiddleware(g, mw...)
	installController(g)
}
