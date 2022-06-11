package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	var feedServer = service.FeedService{}
	c.JSON(200, feedServer.Feed())
	return
}
