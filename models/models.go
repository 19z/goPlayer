package models

import (
	"github.com/astaxie/beego/orm"
	"goPlayer/models/drivers"
	"time"
)

type Member struct {
	Id          int       `orm:"pk"`
	Username    string    `orm:"unique;size(100)"`
	Password    string    `orm:"size(100)"`
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedTime time.Time `orm:"auto_now;type(datetime)"`
}

type Driver struct {
	Id          int       `orm:"pk"`
	Name        string    `orm:"unique"`
	Path        string    `orm:"size(255)"`
	Username    string    `orm:"null;size(1000)"`
	Password    string    `orm:"null;size(1000)"`
	Endpoint    string    `orm:"null;size(100)"`
	Bucket      string    `orm:"null;size(100)"`
	DriverType  string    `orm:"index"`
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedTime time.Time `orm:"auto_now;type(datetime)"`
}

func (d *Driver) GetDriver() drivers.BaseDriver {
	if d.DriverType == "local" {
		return drivers.LocalDriver{Id: d.Id, Path: d.Path}
	}
	return nil
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Driver), new(Member))
}
