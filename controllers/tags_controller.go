package controllers

import (
	"fmt"
	"myblogtest/models"
)

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")

	fmt.Println(models.HandleTagsListData(tags))

	this.Data["Tags"] = models.HandleTagsListData(tags)

	this.TplName = "tags.html"
}


