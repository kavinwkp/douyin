package controller

import (
	"net/http"
	"strconv"

	"douyin/serializer"
	"douyin/service"

	"github.com/gin-gonic/gin"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	to_user_id, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	action_type := c.Query("action_type")

	var followServer = service.FollowActionService{
		Token:    token,
		ToUserID: to_user_id,
	}

	switch action_type {
	case "1":
		c.JSON(200, followServer.FollowAction())
	case "2":
		c.JSON(200, followServer.FollowCancleAction())
	default:
		c.JSON(http.StatusOK, serializer.Response{StatusCode: 1, StatusMsg: "Action type error"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	var followlistServer = service.FollowListService{
		Token:  token,
		UserID: user_id,
	}
	c.JSON(200, followlistServer.FollowList())
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	var followerlistServer = service.FollowerListService{
		Token:    token,
		ToUserID: user_id,
	}
	c.JSON(200, followerlistServer.FollowerList())
	// c.JSON(http.StatusOK, serializer.FollowListResponse{
	// 	Response: serializer.Response{
	// 		StatusCode: 0,
	// 	},
	// 	UserList: []model.User{DemoUser},
	// })
}
