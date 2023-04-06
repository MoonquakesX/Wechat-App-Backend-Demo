package mService

import (
	"WangYiYunDemo/music/mDAO"
	"WangYiYunDemo/music/mDAO/mdDef"
	"WangYiYunDemo/music/mService/msDef"
	"log"
)

func ListComment(req *mdDef.ReturnCommentReq) (*msDef.ListCommentResp, error) {
	//log.Println("not implemented")
	daoResp, err := mDAO.ListComment(req)
	if err != nil {
		log.Println("Dao Error : " + err.Error())
		return nil, err
	}
	resp := new(msDef.ListCommentResp)
	resp.CommentMusicData = daoResp
	resp.Length = len(daoResp)
	return resp, nil
}

func WriteComment(req *mdDef.UploadCommentReq) error {
	//log.Println("not implemented")
	err := mDAO.WriteComment(req)
	if err != nil {
		log.Println("Dao Error : " + err.Error())
		return err
	}
	return nil
}
