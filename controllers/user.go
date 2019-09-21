package controllers

import (
	"RP_platform/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	MainController
}

func (this *RegisterController) Post() {

	res := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &res)
	code := res["code"].(string)
	appid := beego.AppConfig.String("appid")
	appsecret := beego.AppConfig.String("appsecret")
	openidUrl := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appid,
		appsecret,
		code,
	)
	req := httplib.Get(openidUrl)
	str, _ := req.String()
	openidInfo := make(map[string]string)
	json.Unmarshal([]byte(str), &openidInfo)
	openid := openidInfo["openid"]
	//openId

	o := orm.NewOrm()
	user := models.User{}

	user.OpenId = openid
	user.Avatar = res["avatar"].(string)
	user.City = res["city"].(string)
	user.Province = res["province"].(string)
	user.Country = res["country"].(string)
	user.NickName = res["nickName"].(string)

	_, err := o.Insert(&user)
	if err != nil {
		this.Error("录入失败")
	} else {
		this.Error("录入成功")
	}
}
