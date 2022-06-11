package controller

import (
	"strconv"

	"douyin/service"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	// 传入token video_id action_type
	token := c.Query("token")
	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64) // string to int64
	action_type := c.Query("action_type")
	var favoriteServer = service.FavoriteActionService{
		Token:   token,
		VideoId: video_id,
	}
	switch action_type {
	case "1":
		c.JSON(200, favoriteServer.FavoriteAction())
	case "2":
		c.JSON(200, favoriteServer.FavoriteCancleAction())
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64) // string to int64
	var favoriteListServer = service.FavoriteListService{
		Token:  token,
		UserId: user_id,
	}
	c.JSON(200, favoriteListServer.FavoriteList())
}
