package mDAO

import (
	"WangYiYunDemo/music/mDAO/mdDef"
	"fmt"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	DBInit()

	users := []mdDef.UserBasic{
		{
			Name:     "alliance",
			PassWord: "423319",
		},
	}

	// 验证结构
	for idx, _ := range users {
		if err := users[idx].Validate(); err != nil {
			log.Println("CreateUser service interface error : " + err.Error())
			return
		}
	}

	// 加密密码
	for idx, _ := range users {
		if err := users[idx].Encrypt(); err != nil {
			log.Println("CreateUser service interface error : " + err.Error())
			return
		}
	}
	err := MysqlDB.Create(&users).Error
	if err != nil {
		fmt.Println("error : " + err.Error())
		panic(t)
	}
}
