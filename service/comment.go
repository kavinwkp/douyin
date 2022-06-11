package service

import (
	"time"

	"douyin/model"
	"douyin/serializer"
	"douyin/utils"
)

type CommentService struct {
	Token     string `form:"token" json:"token"`
	VideoID   int64  `form:"video_id" json:"video_id"`
	Content   string `form:"user_name" json:"user_name"`
	CommentID int64  `form:"comment_id" json:"comment_id"`
}

func (service *CommentService) Comment() serializer.CommentActionResponse {
	claim, _ := utils.ParseToken(service.Token)
	user_id := claim.UserId

	video_id := service.VideoID
	content := service.Content

	var user model.User
	model.DB.First(&user, user_id)

	var comment = model.Comment{
		UserID:      user_id,
		VideoID:     video_id,
		Content:     content,
		CreatedDate: time.Now().Format("01-02"),
	}

	if err := model.DB.Create(&comment).Error; err != nil {
		return serializer.CommentActionResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "DataBase save comment failed",
			},
		}
	}

	return serializer.CommentActionResponse{
		Response: serializer.Response{StatusCode: 0},
		CommentResponse: model.CommentResponse{
			Id:         1,
			User:       user,
			Content:    content,
			CreateDate: comment.CreatedDate,
		},
	}
}

func (service *CommentService) CommentCancle() serializer.CommentCancleResponse {
	var comment model.Comment
	comment.Id = service.CommentID
	model.DB.Delete(&comment)
	return serializer.CommentCancleResponse{
		Response: serializer.Response{
			StatusCode: 0,
			StatusMsg:  "Cancle comment successfully",
		},
	}
}

func (service *CommentService) CommentList() serializer.CommentListResponse {
	var comments []model.Comment
	video_id := service.VideoID
	model.DB.Where("video_id=?", video_id).Order("id DESC").Find(&comments)

	var commentsResponse []model.CommentResponse
	var commentResponse model.CommentResponse
	for _, v := range comments {
		var user model.User
		user.Id = v.UserID
		model.DB.First(&user)
		commentResponse.Id = v.Id
		commentResponse.Content = v.Content
		commentResponse.CreateDate = v.CreatedDate
		commentResponse.User = user
		commentsResponse = append(commentsResponse, commentResponse)
	}
	return serializer.CommentListResponse{
		Response:    serializer.Response{StatusCode: 0},
		CommentList: commentsResponse,
	}
}
