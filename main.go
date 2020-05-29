package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	_ "goPlayer/routers"
	"os"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
}

func main() {
	if len(os.Args) > 1 {
		orm.RunCommand()
	} else {
		beego.Run()
	}
}
