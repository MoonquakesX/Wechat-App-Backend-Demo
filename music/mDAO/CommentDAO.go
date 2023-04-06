package mDAO

import (
	"WangYiYunDemo/music/mDAO/mdDef"
	"log"
	"strconv"
)

func ListComment(req *mdDef.ReturnCommentReq) ([]*mdDef.CommentBasic, error) {
	cond := make(map[string]interface{})
	cond["id"] = req.Id
	cond["Name"] = req.Name

	if len(req.Tags) > 0 {
		cond["Tags"] = req.Tags
	}
	if req.Limit == "" {
		cond["Limit"] = 20
	} else {
		limit, _ := strconv.Atoi(req.Limit)
		cond["Limit"] = limit
	}
	if req.BeforeTime == "" {
		cond["BeforeTime"] = "0"
	} else {
		cond["BeforeTime"] = req.BeforeTime
	}
	if req.Offset == "" {
		cond["Offset"] = 0
	} else {
		offset, _ := strconv.Atoi(req.Offset)
		cond["Offset"] = offset
	}
	resp := make([]*mdDef.CommentBasic, 0)
	//result := MysqlDB.Table("music_comments").Where("id = ", cond["id"]).Offset(cond["Offset"].(int)).Limit(cond["Limit"].(int)).Scan(&resp)
	result := MysqlDB.Where("name = ?", cond["Name"]).Find(&resp)
	if nil != result.Error {
		log.Println("mysql inner Error : " + result.Error.Error())
		return nil, result.Error
	}
	return resp, nil
}

func WriteComment(req *mdDef.UploadCommentReq) error {
	newItem := mdDef.CommentBasic{Name: req.Name, Content: req.Content, Tags: req.Tags}
	//result := MysqlDB.Table("music_comments").Create(&newItem)
	result := MysqlDB.Create(&newItem)
	if result.Error != nil {
		log.Println("mysql inner Error : " + result.Error.Error())
		return result.Error
	}
	return nil

}
