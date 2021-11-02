package controllers

import (
	"fmt"
	"log"
	"myblogtest/models"
)

/**
 * 删除文章控制器
 */
type DeleteArticleController struct {
	BaseController
}

//点击删除后重定向到首页
func (this *DeleteArticleController) Get() {
	artID, _ := this.GetInt("id")
	fmt.Println("删除 id:", artID)
	art:=models.QueryArticleWithId(artID)
	author:=art.Author
	user:=this.Loginuser
	if author==user{
		_, err := models.DeleteArticle(artID)
		if err != nil {
			log.Println(err)
		}
	}else {
		fmt.Println("文章的创建者和删除者不是同一个人")
		this.Data["json"] = map[string]interface{}{"code": 500, "message": "文章的创建者和删除者不是同一个人"}
		this.ServeJSON()
	}
	this.Redirect("/", 302)
}
