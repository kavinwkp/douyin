package model

type Comment struct {
	Id          int64  `json:"id,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	VideoID     int64  `json:"video_id,omitempty"`
	Content     string `json:"content,omitempty"`
	CreatedDate string `json:"create_data,omitempty"`
}

type CommentResponse struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
