package router

import (
	"ginlayout/internal/config"
	"ginlayout/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	_ "ginlayout/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var ProviderSet = wire.NewSet(NewUserRouter, NewRouter)

type Router struct {
	cfg        *config.Config
	userRouter *UserRouter
}

func NewRouter(cfg *config.Config, userRouter *UserRouter) *Router {
	return &Router{
		cfg:        cfg,
		userRouter: userRouter,
	}
}

func (r *Router) Setup() *gin.Engine {
	if r.cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())
	engine.Use(middleware.Cors())

	if r.cfg.Server.SwaggerEnabled {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	api := engine.Group("/api/v1")
	{
		r.userRouter.Setup(api)
	}

	return engine
}
