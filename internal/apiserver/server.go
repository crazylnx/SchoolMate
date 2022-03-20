package apiserver

import "github.com/gin-gonic/gin"

func Start() {

	logger, _ := ZapConfig().Build()

	logger.Info("start server")
	route := Router()

	err := route.Run(":8181")

	if err != nil {
		logger.Error("Start api server error")
	}
}

func ListUser(c *gin.Context) {
	c.String(200, "pong")
}
