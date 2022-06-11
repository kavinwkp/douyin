package controller

import (
	"fmt"
	"strconv"

	"douyin/model"
	"douyin/service"

	"github.com/gin-gonic/gin"
)

var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

func UserRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var userRegister = service.UserService{
		UserName: username,
		Password: password,
	}
	c.JSON(200, userRegister.Register())
	return
}

func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var userLogin = service.UserService{
		UserName: username,
		Password: password,
	}
	fmt.Println(userLogin)
	c.JSON(200, userLogin.Login())
	return
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	var userInfo = service.UserInfoService{
		Token:  token,
		UserID: user_id,
	}
	c.JSON(200, userInfo.Info())
	return
}
