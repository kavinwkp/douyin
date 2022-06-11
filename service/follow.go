package service

import (
	"errors"
	"fmt"

	"douyin/model"
	"douyin/serializer"
	"douyin/utils"

	"gorm.io/gorm"
)

type FollowActionService struct {
	Token    string
	ToUserID int64
}

type FollowListService struct {
	Token  string
	UserID int64
}

type FollowerListService struct {
	Token    string
	ToUserID int64
}

func (service *FollowActionService) FollowAction() serializer.FollowActionResponse {
	claim, _ := utils.ParseToken(service.Token)
	user_id := claim.UserId
	to_user_id := service.ToUserID

	var follow = model.Follow{
		UserID:   user_id,
		ToUserID: to_user_id,
	}

	if err := model.DB.Create(&follow).Error; err != nil {
		return serializer.FollowActionResponse{
			Response: serializer.Response{StatusCode: 1, StatusMsg: "Database save relation failed"},
		}
	}

	// 关注+1，被关注者粉丝+1，标记为已关注
	model.DB.Model(&model.User{}).Where("id = ?", user_id).Update("follow_count", gorm.Expr("follow_count+1"))
	model.DB.Model(&model.User{}).Where("id = ?", to_user_id).Update("follower_count", gorm.Expr("follower_count+1"))
	model.DB.Model(&model.User{}).Where("id = ?", to_user_id).Update("is_follow", true)

	return serializer.FollowActionResponse{
		Response: serializer.Response{
			StatusCode: 0,
			StatusMsg:  "Follow successfully",
		},
	}
}

func (service *FollowActionService) FollowCancleAction() serializer.FollowActionResponse {
	claim, _ := utils.ParseToken(service.Token)
	user_id := claim.UserId
	to_user_id := service.ToUserID

	var follow model.Follow

	// 删除关注记录
	if err := model.DB.Where("user_id=? AND to_user_id=?", user_id, to_user_id).Delete(&follow).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.FollowActionResponse{
				Response: serializer.Response{StatusCode: 1, StatusMsg: "Not follow"},
			}
		}
	}

	// 关注-1，被关注者粉丝-1，标记为未关注
	model.DB.Model(&model.User{}).Where("id = ?", user_id).Update("follow_count", gorm.Expr("follow_count-1"))
	model.DB.Model(&model.User{}).Where("id = ?", to_user_id).Update("follower_count", gorm.Expr("follower_count-1"))
	model.DB.Model(&model.User{}).Where("id = ?", to_user_id).Update("is_follow", false)

	return serializer.FollowActionResponse{
		Response: serializer.Response{
			StatusCode: 0,
			StatusMsg:  "Cancle follow successfully",
		},
	}
}

func (service *FollowListService) FollowList() serializer.FollowListResponse {

	var follow []model.Follow

	model.DB.Where("user_id=?", service.UserID).Find(&follow)

	var to_user_ids []int64
	for _, v := range follow {
		to_user_ids = append(to_user_ids, v.ToUserID)
	}

	var users []model.User
	model.DB.Find(&users, to_user_ids)

	return serializer.FollowListResponse{
		Response: serializer.Response{
			StatusCode: 0,
		},
		UserList: users,
	}
}

func (service *FollowerListService) FollowerList() serializer.FollowerListResponse {
	var follow []model.Follow

	res := model.DB.Where("to_user_id=?", service.ToUserID).Find(&follow)
	if res.RowsAffected == 0 {
		return serializer.FollowerListResponse{
			Response: serializer.Response{
				StatusCode: 0,
			},
			UserList: []model.User{},
		}
	}

	var user_ids []int64
	for _, v := range follow {
		user_ids = append(user_ids, v.UserID)
	}

	fmt.Println(user_ids)

	var users []model.User
	model.DB.Find(&users, user_ids)

	return serializer.FollowerListResponse{
		Response: serializer.Response{
			StatusCode: 0,
		},
		UserList: users,
	}
}
