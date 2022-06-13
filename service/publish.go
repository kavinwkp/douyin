package service

import (
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"

	"douyin/config"
	"douyin/model"
	"douyin/serializer"
	"douyin/utils"

	"gorm.io/gorm"
)

type PublishService struct {
	Token      string
	Title      string
	FileHeader *multipart.FileHeader
}

type PublishListService struct {
	Token  string `json:"token"`
	UserID int64  `json:"user_id"`
}

func (service *PublishService) Publish() serializer.PublishResponse {
	claim, _ := utils.ParseToken(service.Token)

	var user model.User
	user.Id = claim.UserId
	if err := model.DB.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.PublishResponse{
				Response: serializer.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			}
		}
	}
	filename := filepath.Base(service.FileHeader.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := utils.UploadVideo(service.FileHeader, saveFile); err != nil {
		return serializer.PublishResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		}
	}
	rand.Seed(time.Now().Unix())
	// 视频保存成功就在Video表中插入视记录
	var video = model.VideoTable{
		UserID:        user.Id,
		Title:         service.Title,
		PlayUrl:       finalName,
		CoverUrl:      "cover_" + strconv.Itoa(rand.Intn(6)) + ".png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
	if err := model.DB.Create(&video).Error; err != nil {
		return serializer.PublishResponse{
			Response: serializer.Response{StatusCode: 1, StatusMsg: "Database save video failed"},
		}
	}
	return serializer.PublishResponse{
		Response: serializer.Response{
			StatusCode: 0,
			StatusMsg:  finalName + " uploaded successfully",
		},
	}
}

func (service *PublishListService) PublishList() serializer.VideoListResponse {
	// print(service.Token)
	// print(service.UserID)
	var videosTable []model.VideoTable
	result := model.DB.Where("user_id=?", service.UserID).Find(&videosTable)

	if result.Error != nil || result.RowsAffected == 0 {
		return serializer.VideoListResponse{
			Response: serializer.Response{
				StatusCode: 0,
			},
			VideoList: []model.Video{},
		}
	}

	var videos []model.Video

	var user model.User
	model.DB.First(&user, videosTable[0].UserID)

	for _, v := range videosTable {
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
