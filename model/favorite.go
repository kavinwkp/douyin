package model

type Favorite struct {
	Id      int64
	UserID  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}
