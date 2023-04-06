package mDAO

import (
	"WangYiYunDemo/music/mDAO/mdDef"
	"fmt"
	"testing"
)

func TestCreateComments(t *testing.T) {
	DBInit()
	Comments := []mdDef.CommentBasic{
		{
			Name:    "一万次悲伤",
			Content: "逃跑计划-世界",
			Tags:    "rap",
		},
		{
			Name:    "一万次悲伤",
			Content: "我一直在最后的地方等你",
			Tags:    "rap",
		},
		{
			Name:    "一万次悲伤",
			Content: "每一颗眼泪是一万道光",
			Tags:    "rap",
		},
	}
	err := MysqlDB.Create(&Comments).Error
	if err != nil {
		fmt.Println("error : " + err.Error())
		panic(t)
	}

}
