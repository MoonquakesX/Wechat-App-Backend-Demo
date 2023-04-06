package mDAO

import (
	"WangYiYunDemo/music/mDAO/mdDef"
	"fmt"
	"testing"
)

func TestCreateSongs(t *testing.T) {
	DBInit()
	songs := []mdDef.SongBasic{
		{
			Name:   "淘汰",
			Singer: "陈奕迅",
			Tags:   "pop",
		},
		{
			Name:   "Hey Kong",
			Singer: "刘聪",
			Tags:   "rap",
		},
		{
			Name:   "凄美地",
			Singer: "郭顶",
			Tags:   "pop",
		},
		{
			Name:   "明明就",
			Singer: "周杰伦",
			Tags:   "pop",
		},
	}
	err := MysqlDB.Create(&songs).Error
	if err != nil {
		fmt.Println("error : " + err.Error())
		panic(t)
	}
}
