package model

import "golang.org/x/crypto/bcrypt"

type UserPassword struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// 加密
func (user *UserPassword) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// 用户登录时密码检查
func (user *UserPassword) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
