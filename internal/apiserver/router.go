package apiserver

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	router := gin.New()
	// user
	v1 := router.Group("/v1")
	{
		v1.GET("/users", ListUser)
		//v1.GET("/users/:id")
		//v1.POST("/users")
		//v1.DELETE("/users")
	}

	return router
}
