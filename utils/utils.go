package utils

import (
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var JWTsecret = []byte("ABAB")

type Claims struct {
	UserId   int64  `json:"id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// 签发token
func GenerateToken(id int64, username string) (string, error) {
	notTime := time.Now()
	expireTime := notTime.Add(24 * time.Hour) // 24小时后过期
	claims := Claims{
		UserId:   id,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "douyin",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}

// 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func UploadVideo(fileHeader *multipart.FileHeader, path string) error {
	ctx := &gin.Context{}
	err := ctx.SaveUploadedFile(fileHeader, path)
	if err != nil {
		return err
	}
	return nil
}
