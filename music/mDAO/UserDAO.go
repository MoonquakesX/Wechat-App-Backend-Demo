package mDAO

import (
	"WangYiYunDemo/music/mDAO/mdDef"
	"log"
)

// GetUserByName 基于名字获取用户
func GetUserByName(username string) (*mdDef.UserBasic, error) {
	user := &mdDef.UserBasic{}
	result := MysqlDB.Where("name = ?", username).First(user)
	if result.Error != nil {
		log.Println("DB error : " + result.Error.Error())
	}
	return user, nil
}
