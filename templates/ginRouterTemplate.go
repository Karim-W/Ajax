package templates

var GinRouterTemplate = `
package router

import (
	"context"
	"net/http"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/rs/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
// @schemes   http

//Add your controllers here
func SetupRoutes() *gin.Engine {
	router := gin.Default()
	// uncomment if you have the index page
	// router.LoadHTMLFiles("./index.html")
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", nil)
	// })
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	v1 := router.Group("api/v1")
	router.Run()
	return router
}

func registerHooks(lifecycle fx.Lifecycle, ginRouter *gin.Engine, logger *zap.SugaredLogger) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Initializing server")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Terminating server")
			logger.Sync()
			return nil
		},
	})
}

var Module = fx.Options(fx.Provide(SetupRoutes), fx.Invoke(registerHooks))

`
