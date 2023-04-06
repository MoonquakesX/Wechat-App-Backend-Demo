package mService

import (
	"WangYiYunDemo/music/mDAO"
	"WangYiYunDemo/music/mDAO/mdDef"
	"WangYiYunDemo/music/mService/msDef"
	"log"
)

func ListSongs(req *mdDef.ListSongsReq) (*msDef.ListSongsResp, error) {
	//log.Println("not implemented")
	daoResp, err := mDAO.ListSongs(req)
	if err != nil {
		log.Println("Dao Error : " + err.Error())
		return nil, err
	}
	resp := new(msDef.ListSongsResp)
	resp.SongsData = daoResp
	resp.Length = len(daoResp)
	return resp, nil
}

func PlaySong(id uint) error {
	err := mDAO.PlaySong(id)
	if err != nil {
		log.Println("PlaySong DAO error :" + err.Error())
		return err
	}
	return nil
}

func GetPlayedTimes(id uint) (int, error) {
	resp, err := mDAO.GetPlayedTimes(id)
	if err != nil {
		log.Println("GetPlayedTimes DAO error :" + err.Error())
		return 0, err
	}
	return resp, nil
}

func GetPlayList() (*mdDef.PlayList, error) {
	resp, err := mDAO.GetPlayList()
	if err != nil {
		log.Println("GetPlayList DAO error : ", err.Error())
		return nil, err
	}
	return resp, nil
}
