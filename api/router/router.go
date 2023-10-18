package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"web_server_2.0/api/html"
	"web_server_2.0/bootstrap"
	"web_server_2.0/docs"
	"web_server_2.0/middleware"
)

func Setup(app *bootstrap.Application) {
	//health check
	app.Web.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	docs.SwaggerInfo.BasePath = "/v1/api"
	rhtml := app.Web.Group("/html")
	html.NewArticleTmp(app, rhtml)

	v1 := app.Web.Group("/v1/api")
	v1.Use(cors.Default())

	publicRouter := v1.Group("")
	NewLoginRoute(app, publicRouter)
	NewUserArticle(app, publicRouter)

	protectedRouter := v1.Group("")
	protectedRouter.Use(middleware.AuthMiddleware(app))
	NewUserRoute(app, protectedRouter)

	//swagger
	app.Web.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
