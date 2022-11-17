package server

import (
	_ "SoftwareEngine/docs"
	"SoftwareEngine/internal/pkg/constant"
	"SoftwareEngine/internal/pkg/middleware"
	"SoftwareEngine/internal/server/controller/v1/problem"
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
	}

	v1 := g.Group("/v1")

	// 题目模块
	{
		problemH := problem.NewProblemController(storeIns)
		problem := v1.Group("/problem", authMiddleware("2"))
		{
			problem.GET("tag", problemH.GetTags)
			problem.POST("all", problemH.GetAll)
			problem.GET(":problemName", problemH.GetProblem)
		}

		//// 题目进行修改，添加删除
		authProblemv1 := v1.Group("/problem", authMiddleware("1"))
		{
			authProblemv1.POST("create", problemH.Create)
			//authProblemv1.PUT("")
			//authProblemv1.DELETE("name")
		}
	}
	// 评论区模块
	{

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
