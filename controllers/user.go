package controllers

import (
	"RP_platform/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type RegisterController struct {
	MainController
}
type LoginController struct {
	MainController
}

type UserController struct {
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

	//信息录入数据库
	o := orm.NewOrm()
	newUser := models.User{}
	newAccount := models.AccountInfo{}

	newUser.OpenId = openid
	newUser.Avatar = res["avatar"].(string)
	newUser.City = res["city"].(string)
	newUser.Province = res["province"].(string)
	newUser.Country = res["country"].(string)
	newUser.NickName = res["nickName"].(string)
	newUser.AccountInfo = &newAccount
	o.Insert(&newUser)
	o.Insert(&newAccount)

	this.Success("录入成功", map[string]int{"uid": newUser.Uid})
}

func (this *LoginController) Post() {
	//登录
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
	var user models.User
	qs := orm.NewOrm().QueryTable("user")
	qs.Filter("open_id", openid).One(&user)

	this.Success("success", map[string]int{"uid": user.Uid})
}

//func (this *UserController) Get() {
////	uid := this.GetString("uid")
////	id,err := strconv.Atoi(uid)
////	if err != nil {
////		this.Error("uid错误")
////	}
////	user := models.ReadUser(id)
////
////	this.Data["uid"] = uid
////	this.Data["username"] = user.NickName
////	this.Data["integral"] = user.AccountInfo.Integral
////	this.Data["balance"] = user.AccountInfo.Balance
////	this.Data["avatarUrl"] = user.Avatar
////	this.TplName = "userpage.html"
////}
func (this *UserController) Get() {
	uid := this.GetString("uid")
	id, err := strconv.Atoi(uid)
	if err != nil {
		this.Error("uid错误")
	}
	user := models.ReadUser(id)
	this.Success("success", map[string]interface{}{
		"uid":      uid,
		"avatar":   user.Avatar,
		"username": user.NickName,
		"balance":  user.AccountInfo.Balance,
		"integral": user.AccountInfo.Integral,
	})
}
