package routers

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"rest-api/controllers/todo"
	"rest-api/controllers/user"
	_ "rest-api/docs"
	"rest-api/routers/middleware"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// add middleware
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Rất tiếc! đã xảy ra lỗi")
	})

	v1 := g.Group("/api/v1")
	{
		// swagger api docs
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		pprof.Register(g, "")

		// auth group
		au := v1.Group("/auth")
		{
			au.POST("/login", user.Login)
		}

		// user group
		u := v1.Group("/user")
		{
			u.POST("", user.Create)
		}

		// todo group
		td := v1.Group("/todo")
		td.Use(middleware.AuthMiddleware())
		{
			td.POST("", todo.Create)
			td.POST("/list", todo.List)
			td.PUT("/:id", todo.Update)
			td.DELETE("/:id", todo.Delete)
		}
	}

	return g
}
