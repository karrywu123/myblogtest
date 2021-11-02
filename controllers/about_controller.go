package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {

	c.Data["wechat"] = "微信：123456789"
	c.Data["qq"] = "QQ：123456789"
	c.Data["tel"] = "Tel：123456789"

	c.TplName = "aboutme.html"
}
