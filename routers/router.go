package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mygetzu/karyakita-backend-go/controllers"
)

func createRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			root := new(controllers.RootController)
			v1.GET("/", root.GetRoot)
		}
	}

	return router
}
