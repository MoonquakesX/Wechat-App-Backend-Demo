package mApi

import (
	"WangYiYunDemo/music/mDAO"
	"WangYiYunDemo/music/mDAO/mdDef"
	"WangYiYunDemo/music/mService"
	"WangYiYunDemo/music/mService/msDef"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Hello(context *gin.Context) {
	log.Println(">>>>> Hello gin Success! <<<<<")
	context.String(http.StatusOK, "Music Ok!")

}

// ListSongs 为了swag识别接口， 通常需要加一些接口描述， 是格式化的， 需要阅读文档
// todo: complete swag description
func ListSongs(context *gin.Context) {
	log.Println("come into ListSongs function")
	req := new(mdDef.ListSongsReq)
	err := context.ShouldBind(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	resp, err := mService.ListSongs(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	//todo : 完善http response结构体封装， 在data的基础上加上success， code， msg等
	context.JSON(http.StatusOK, resp)
}

func ListComments(context *gin.Context) {
	//log.Println("come into ListComments function")
	req := new(mdDef.ReturnCommentReq)
	err := context.ShouldBind(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
	}

	resp, err := mService.ListComment(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
	}
	context.JSON(http.StatusOK, resp)
}

func UploadComment(context *gin.Context) {
	log.Println("come into UploadComment function")
	req := new(mdDef.UploadCommentReq)
	err := context.ShouldBind(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
	}
	err = mService.WriteComment(req)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, nil)
	}
	context.JSON(http.StatusOK, nil)
}

func PlaySong(context *gin.Context) {
	songId := context.Param("id")
	id, err := strconv.Atoi(songId)
	if err != nil {
		log.Println("string conversion error : " + err.Error())
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	err = mService.PlaySong(uint(id))
	if err != nil {
		log.Println("PlaySong service interface error : " + err.Error())
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	context.JSON(http.StatusOK, nil)
}

func GetPlayTimes(context *gin.Context) {
	songId := context.Param("id")
	var resp int
	id, err := strconv.Atoi(songId)
	if err != nil {
		log.Println("string conversion error : " + err.Error())
		context.JSON(http.StatusBadRequest, resp)
		return
	}
	resp, err = mService.GetPlayedTimes(uint(id))
	if err != nil {
		log.Println("GetPlayTimes service interface error : " + err.Error())
		context.JSON(http.StatusBadRequest, resp)
		return
	}
	context.JSON(http.StatusOK, resp)
}

func GetPlayList(context *gin.Context) {
	resp, err := mService.GetPlayList()
	if err != nil {
		log.Println("GetPlayList service interface error : " + err.Error())
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	context.JSON(http.StatusOK, resp)
}

func Login(context *gin.Context) {
	var u msDef.UserReq
	// 应该使用 ShouldBindJSON, 以便使用自定义的 handler.SendResponse
	if err := context.ShouldBindJSON(&u); err != nil {
		log.Println("login service interface error : " + err.Error())
		return
	}

	user, err := mDAO.GetUserByName(u.Username)
	if err != nil {
		log.Println("get name interface error : " + err.Error())
		return
	}

	if err := user.Compare(u.Password); err != nil {
		log.Println("password error : " + err.Error())
		return
	}
	t, err := mService.Sign(user.ID, user.Name)
	if err != nil {
		log.Println("sign error : " + err.Error())
		return
	}

	context.JSON(http.StatusOK, msDef.Token{Token: t})
}
