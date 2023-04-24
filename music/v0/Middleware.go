package v0

import (
	"WangYiYunDemo/music/mService"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func AuthJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.GetHeader("Authorization")
		headerList := strings.Split(header, " ")
		if len(headerList) != 2 {
			err := errors.New("无法解析 Authorization 字段")
			log.Println("head error" + err.Error())
			context.Abort()
			return
		}
		t := headerList[0]
		content := headerList[1]

		if t != "Bearer" {
			err := errors.New("认证类型错误, 当前只支持 Bearer")
			log.Println("head error" + err.Error())
			context.Abort()
			return
		}
		if _, err := mService.Verify([]byte(content)); err != nil {
			log.Println("verify error " + err.Error())
			context.Abort()
			return
		}

		context.Next()
	}
}
