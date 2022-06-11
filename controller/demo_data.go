package controller

import "douyin/model"

var DemoVideos = []model.Video{
	{
		Id:     1,
		Author: DemoUser,
		// PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		// CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		PlayUrl:       "http://10.37.62.58:8080/static/bear.mp4",
		CoverUrl:      "http://10.37.62.58:8080/static/bear.jpg", // 不能写127.0.0.1
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []model.CommentResponse{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = model.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
