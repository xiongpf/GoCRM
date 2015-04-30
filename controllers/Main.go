package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

//func (c *MainController) URLMapping() {
//    c.Mapping("Get", c.Get())
//    c.Mapping("Hello", c.Test())
//}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.html"
}

func (c *MainController) Test() {
	c.TplNames = "contact/helloworld.html"
}
