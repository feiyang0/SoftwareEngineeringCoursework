package server

import (
	_ "SoftwareEngine/docs"
	"SoftwareEngine/internal/pkg/middleware"
	"SoftwareEngine/internal/server/controller/v1/user"
	"SoftwareEngine/internal/server/store/mysql"
	"SoftwareEngine/pkg/core"
	"SoftwareEngine/pkg/errno"

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
	// api
	{
		userController := user.NewUserController(storeIns)
		g.GET("/test", userController.Test)
		g.POST("/login", userController.Login)
		g.POST("/register", userController.Register)
		g.GET("/resetPasswd", userController.GetCaptcha)
		g.POST("/resetPasswd", userController.CheckCaptcha)
		g.PUT("/resetPasswd", userController.SetNewPasswd)
	}

	//v1 := g.Group("/v1")
	//{
	//
	//}
}

func loadRouter(g *gin.Engine, mw ...gin.HandlerFunc) {
	installMiddleware(g, mw...)
	installController(g)
}
