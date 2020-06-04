package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"goPlayer/models/drivers"
	"os"
	"strings"
	"time"
)

// Member 用户
type Member struct {
	Id          int       `orm:"pk"`                          // 用户ID
	Username    string    `orm:"unique;size(100)"`            // 用户名
	Password    string    `orm:"size(100)"`                   // 密码哈希
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
	UpdatedTime time.Time `orm:"auto_now;type(datetime)"`     // 更新时间
}

// Driver 设备
type Driver struct {
	Id          int       `orm:"pk;auto"`                     // 设备id
	Name        string    `orm:"unique"`                      // 设备名
	Path        string    `orm:"size(255)"`                   // 路径
	Username    string    `orm:"null;size(1000)"`             // 访问需要的用户名或 app id
	Password    string    `orm:"null;size(1000)"`             //访问需要的密码或者 app secret
	Endpoint    string    `orm:"null;size(100)"`              // 访问的endpoint，aws，阿里云等云存储会需要
	Bucket      string    `orm:"null;size(100)"`              // bucket，aws，阿里云等云存储会需要
	DriverType  string    `orm:"index"`                       // 设备类型，local 为本地，目前只支持 local
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
	UpdatedTime time.Time `orm:"auto_now;type(datetime)"`     //更新时间
}

func (d *Driver) GetDriver() drivers.BaseDriver {
	if d.DriverType == "local" {
		return drivers.LocalDriver{Id: d.Id, Path: d.Path}
	}
	return nil
}

func (d *Driver) CheckData() error {
	if d.DriverType != "local" {
		return errors.New("only local can use now")
	}
	d.Path = strings.ReplaceAll(d.Path, "\\", "/")
	stat, err := os.Stat(d.Path)
	if err != nil {
		return err
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
