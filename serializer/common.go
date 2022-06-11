package serializer

import "douyin/model"

// 基础序列化器
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

type PublishResponse struct {
	Response
}

type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}

type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

type FavoriteActionResponse struct {
	Response
}

type CommentListResponse struct {
	Response
	CommentList []model.CommentResponse `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	CommentResponse model.CommentResponse `json:"comment,omitempty"`
}

type CommentCancleResponse struct {
	Response
}

type FollowActionResponse struct {
	Response
}

type FollowListResponse struct {
	Response
	UserList []model.User `json:"user_list"`
}

type FollowerListResponse struct {
	Response
	UserList []model.User `json:"user_list"`
}
