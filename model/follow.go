package model

type Follow struct {
	Id       int64 `json:"id,omitempty"`
	UserID   int64 `json:"user_id,omitempty"`
	ToUserID int64 `json:"to_user_id,omitempty"`
}
