package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"myblogtest/models"
)

/**
 * 相册控制器
 */
type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		logs.Error(err.Error())
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
	//this.TplName = "album.html"
}
