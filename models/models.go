package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"goPlayer/models/drivers"
	"os"
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
	Id          int       `orm:"pk;auto"`
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

func (d *Driver) CheckData() error {
	if d.DriverType != "local" {
		return errors.New("only local can use");
	}
	stat, err := os.Stat(d.Path);
	if err != nil {
		return err;
	}
	if !stat.IsDir() {
		return errors.New("only directory can use")
	}
	return nil
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Driver), new(Member))
}
