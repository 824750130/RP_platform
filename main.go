package main

import (
	_ "RP_platform/models"
	_ "RP_platform/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	username := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("mysqlhost")
	port := beego.AppConfig.String("mysqlport")
	db := beego.AppConfig.String("mysqldb")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, db)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink)
	orm.RunSyncdb("default", false, false)
	orm.DefaultTimeLoc = time.UTC
}

func main() {
	beego.Run()
}
