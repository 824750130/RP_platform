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
type LoginController struct {
	MainController
}

func (this *RegisterController) Post() {
	//register api
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
	//openId
	req := httplib.Get(openidUrl)
	str, _ := req.String()
	openidInfo := make(map[string]string)
	json.Unmarshal([]byte(str), &openidInfo)
	openid := openidInfo["openid"]

	o := orm.NewOrm()

	newUser := models.User{}
	newAccount := models.AccountInfo{}

	newUser.OpenId = openid
	newUser.Avatar = res["avatar"].(string)
	newUser.City = res["city"].(string)
	newUser.Province = res["province"].(string)
	newUser.Country = res["country"].(string)
	newUser.NickName = res["nickName"].(string)
	newAccount.User = &newUser
	o.Insert(&newUser)
	o.Insert(&newAccount)

	this.Success("录入成功", map[string]string{"openid": openid})
}

//func (this *LoginController) Post() {
//	//登录
//	o := orm.NewOrm()
//	var user models.User
//	qs := o.QueryTable("user")
//	userobj := qs.Filter("open_id", openid).One(&user)
//	fmt.Println(userobj)
//}
