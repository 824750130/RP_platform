package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Uid      int    `orm:"column(uid);pk"`
	OpenId   string `orm:"column(open_id);unique"`
	Avatar   string `orm:"unique"`
	City     string
	Province string
	Country  string

	NickName    string       `orm:"column(nick_name);size(32)"`
	CreatedTime time.Time    `orm:"column(create_time);auto_now_add;type(datetime)"`
	Inviter     *User        `orm:"column(inviter);null;rel(fk);on_delete(set_null)"`
	User        []*User      `orm:"reverse(many)"`
	AccountInfo *AccountInfo `orm:"rel(one);on_delete(cascade)"`
}

type AccountInfo struct {
	Aid      int     `orm:"column(aid);pk"`
	Balance  float64 `orm:"column(balance);default(0)"`
	Integral int     `orm:"default(0)"`
	User     *User   `orm:"reverse(one)"`
}

//type Notice  struct {
//	id        int         `orm:"pk"`
//	Title     string
//	Content   string
//}

//type Activity struct {
//	id        int         `orm:"pk"`
//}

func init() {
	orm.RegisterModel(new(User), new(AccountInfo))
}

func ReadUser(uid int) User {
	o := orm.NewOrm()
	user := User{Uid: uid}
	o.Read(&user)
	return user
}
