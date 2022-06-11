package controller

import (
	"net/http"
	"strconv"

	"douyin/service"

	"douyin/serializer"

	"github.com/gin-gonic/gin"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	action_type := c.Query("action_type")
	content := c.Query("comment_text")
	var commentServer = service.CommentService{
		Token:   token,
		VideoID: video_id,
		Content: content,
	}
	switch action_type {
	case "1":
		c.JSON(200, commentServer.Comment())
	case "2":
		comment_id, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		commentServer.CommentID = comment_id
		c.JSON(200, commentServer.CommentCancle())
	default:
		c.JSON(http.StatusOK, serializer.Response{StatusCode: 1, StatusMsg: "Action_Type error"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	var commentServer = service.CommentService{
		Token:   token,
		VideoID: video_id,
	}
	c.JSON(200, commentServer.CommentList())
	// c.JSON(http.StatusOK, serializer.CommentListResponse{
	// 	Response:    serializer.Response{StatusCode: 0},
	// 	CommentList: DemoComments,
	// })
}
