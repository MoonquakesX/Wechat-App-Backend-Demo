package mDAO

import (
	"WangYiYunDemo/music/mDAO/mdDef"
	"container/list"
	"gorm.io/gorm/clause"
	"log"
)

func BatchListSongs(req *mdDef.ListSongsReq) (map[uint]*mdDef.SongBasic, error) {
	var songs []*mdDef.SongBasic
	cond := MysqlDB
	log.Println("ids is", len(req.IDs))
	if len(req.IDs) > 0 {
		cond = cond.Where("id in ?", req.IDs)
	} else {
		cond = cond.Where("id < 0") //means find none
	}
	if len(req.Names) > 0 {
		cond = cond.Or("name in ? ", req.Names)
	}
	if len(req.Singers) > 0 {
		cond = cond.Or("singer in ? ", req.Singers)
	}
	if len(req.Tags) > 0 {
		cond = cond.Or("tags in ? ", req.Tags)
	}
	if err := cond.Find(&songs).Error; err != nil {
		log.Println("mysql inner Error : " + err.Error())
		return nil, err
	}
	resp := make(map[uint]*mdDef.SongBasic)
	for _, song := range songs {
		resp[song.ID] = song
	}
	return resp, nil
}

func ListSongs(req *mdDef.ListSongsReq) ([]*mdDef.SongBasic, error) {
	roughResp, err := BatchListSongs(req)
	if err != nil {
		log.Println("mysql inner Error : " + err.Error())
		return nil, err

	}
	resp := make([]*mdDef.SongBasic, 0, len(roughResp))
	for _, song := range roughResp {
		resp = append(resp, song)
	}
	return resp, nil
}

func PlaySong(id uint) error {
	tx := MysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	cond := make(map[string]interface{})
	cond["id"] = id
	var songDetail mdDef.SongBasic

	r := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id", id).Find(&songDetail)
	if r.Error != nil {
		log.Println("DB error : " + r.Error.Error())
		return r.Error
	}

	r = tx.Where(cond).Find(&songDetail)
	if r.Error != nil {
		log.Println("DB error : " + r.Error.Error())
		return r.Error
	}

	r = tx.Model(&songDetail).Update("played_times", songDetail.PlayedTimes+1)
	if r.Error != nil {
		log.Println("DB error : " + r.Error.Error())
		return r.Error
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("mysql transaction error : " + err.Error())
		return err
	}

	if err := InsertIntoPlayBuffer(id); err != nil {
		log.Println("insert into play buffer error :" + err.Error())
		return err
	}
	return nil
}

func GetPlayedTimes(id uint) (int, error) {
	var songDetail mdDef.SongBasic
	cond := make(map[string]interface{})
	cond["id"] = id
	if err := MysqlDB.Where(cond).Find(&songDetail).Error; err != nil {
		log.Println("Mysql Select Error : " + err.Error())
		return 0, err
	}
	return songDetail.PlayedTimes, nil
}

var PlayBuffer list.List

func InsertIntoPlayBuffer(id uint) error {
	for i := PlayBuffer.Front(); i != nil; i = i.Next() {
		if i.Value == id {
			PlayBuffer.Remove(i)
			break
		}
	}

	PlayBuffer.PushBack(id)
	return nil
}

func GetPlayList() (*mdDef.PlayList, error) {
	ids := make([]uint, 0, PlayBuffer.Len())
	for i := PlayBuffer.Front(); i != nil; i = i.Next() {
		ids = append(ids, i.Value.(uint))
	}
	songMap, err := BatchListSongs(&mdDef.ListSongsReq{IDs: ids})
	if err != nil {
		log.Println("mysql inner Error : " + err.Error())
		return nil, err

	}
	resp := make([]*mdDef.SongBasic, 0, len(songMap))
	for _, id := range ids {
		resp = append(resp, songMap[id])
	}
	return &mdDef.PlayList{
		Length: len(ids),
		Songs:  resp,
	}, nil
}
