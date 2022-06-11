package service

import (
	"errors"

	"douyin/model"
	"douyin/serializer"
	"douyin/utils"

	"gorm.io/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

type UserInfoService struct {
	Token  string `form:"token" json:"token"`
	UserID int64  `form:"user_id" json:"user_id"`
}

func (service *UserService) Register() serializer.Response {

	var user model.UserPassword
	result := model.DB.Where("name=?", service.UserName).First(&user)
	if result.RowsAffected == 1 {
		return serializer.Response{
			StatusCode: 1,
			StatusMsg:  "User already exist",
		}
	}
	// atomic.AddInt64(&userIdSequence, 1)
	user.Name = service.UserName
	user.SetPassword(service.Password)
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			StatusCode: 1,
			StatusMsg:  "DataBase save user failed",
		}
	}
	model.DB.Create(&model.User{
		Name:          user.Name,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	})
	return serializer.Response{
		StatusCode: 0,
		StatusMsg:  "register successfully",
	}
}

func (service *UserService) Login() serializer.UserLoginResponse {
	var user model.UserPassword

	if err := model.DB.Where("name=?", service.UserName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.UserLoginResponse{
				Response: serializer.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			}
		}
	}
	// 用户名没错就验证密码
	if user.CheckPassword(service.Password) == false {
		return serializer.UserLoginResponse{
			Response: serializer.Response{StatusCode: 1, StatusMsg: "Password error"},
		}
	}
	// 密码正确，签发token
	token, _ := utils.GenerateToken(user.Id, user.Name)
	return serializer.UserLoginResponse{
		Response: serializer.Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    token,
	}

}

func (service *UserInfoService) Info() serializer.UserResponse {
	var user model.User
	if err := model.DB.Where("id=?", service.UserID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.UserResponse{
				Response: serializer.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			}
		}
	}
	return serializer.UserResponse{
		Response: serializer.Response{StatusCode: 0},
		User:     user,
	}
}
