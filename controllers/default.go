package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Success(Msg string, data interface{}) {
	code := 0
	this.Data["json"] = map[string]interface{}{"code": code, "msg": Msg, "data": data}
	this.ServeJSON()

}
func (this *MainController) Error(Msg string) {
	code := 1000
	this.Data["json"] = map[string]interface{}{"code": code, "msg": Msg}
	this.ServeJSON()
}
