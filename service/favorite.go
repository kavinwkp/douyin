package service

import (
	"errors"

	"douyin/config"
	"douyin/model"
	"douyin/serializer"
	"douyin/utils"

	"gorm.io/gorm"
)

type FavoriteActionService struct {
	Token   string
	VideoId int64
}

type FavoriteListService struct {
	Token  string
	UserId int64
}

func (service *FavoriteActionService) FavoriteAction() serializer.FavoriteActionResponse {

	claim, _ := utils.ParseToken(service.Token)
	video_id := service.VideoId
	user_id := claim.UserId

	var favorite model.Favorite

	result := model.DB.Where("user_id=? AND video_id=?", user_id, video_id).First(&favorite)
	if result.RowsAffected == 1 {
		return serializer.FavoriteActionResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "Already favorite",
			},
		}
	}
	// 没点过赞就新增一条记录
	favorite.UserID = user_id
	favorite.VideoId = video_id

	err := model.DB.Create(&favorite).Error
	if err != nil {
		return serializer.FavoriteActionResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "DataBase save favorite failed",
			},
		}
	}
	// 标记视频已点赞和更新点赞数
	model.DB.Model(&model.VideoTable{}).Where("id = ?", video_id).Update("is_favorite", true)
	model.DB.Model(&model.VideoTable{}).Where("id = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count+1"))

	return serializer.FavoriteActionResponse{
		Response: serializer.Response{
			StatusCode: 0,
			StatusMsg:  "Favorite successfully",
		},
	}
}

func (service *FavoriteActionService) FavoriteCancleAction() serializer.FavoriteActionResponse {
	claim, _ := utils.ParseToken(service.Token)
	video_id := service.VideoId
	user_id := claim.UserId

	var favorite model.Favorite

	if err := model.DB.Where("user_id=? AND video_id=?", user_id, video_id).First(&favorite).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.FavoriteActionResponse{
				Response: serializer.Response{StatusCode: 1, StatusMsg: "Not favorite"},
			}
		}
	}
	model.DB.Delete(&favorite)
	model.DB.Model(&model.VideoTable{}).Where("id = ?", video_id).Update("is_favorite", false)
	model.DB.Model(&model.VideoTable{}).Where("id = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count-1"))

	return serializer.FavoriteActionResponse{
		Response: serializer.Response{
			StatusCode: 0,
			StatusMsg:  "Cancle favorite successfully",
		},
	}
}

func (service *FavoriteListService) FavoriteList() serializer.VideoListResponse {
	var favorites []model.Favorite
	res := model.DB.Where("user_id=?", service.UserId).Find(&favorites)
	if res.RowsAffected == 0 {
		return serializer.VideoListResponse{
			Response: serializer.Response{
				StatusCode: 0,
			},
			VideoList: []model.Video{},
		}
	}
	var video_ids []int64
	for _, v := range favorites {
		video_ids = append(video_ids, v.VideoId)
	}

	var videosTable []model.VideoTable
	model.DB.Find(&videosTable, video_ids)

	var videos []model.Video
	for _, v := range videosTable {
		var user model.User
		model.DB.First(&user, v.UserID)
		var video = model.Video{
			Id:            v.Id,
			Title:         v.Title,
			Author:        user,
			PlayUrl:       config.BaseURL + v.PlayUrl,
			CoverUrl:      config.BaseURL + v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
		}
		videos = append(videos, video)
	}

	return serializer.VideoListResponse{
		Response: serializer.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	}
}
