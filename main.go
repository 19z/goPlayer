package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"goPlayer/helper"
	_ "goPlayer/routers"
	"os"
)

func init() {
	_ = orm.RegisterDriver("sqlite3", orm.DRSqlite)
	_ = orm.RegisterDataBase("default", "sqlite3", "data.db")
	_ = beego.AddFuncMap("FileSizeFormat", helper.FileSizeFormat)
	_ = beego.AddFuncMap("TimeFormat", helper.TimeFormat)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "orm" {
		// go run main.go orm
		//go run main.go orm syncdb  --force
		orm.RunCommand()
	} else {
		beego.Run()
	}
}
