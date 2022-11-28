package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
)

type User struct {
	Id            int64
	Email         string    `orm:"size(64);unique" form:"Email" valid:"Required;Email"`
	Password      string    `orm:"size(32)" form:"Password" valid:"Required;MinSize(6)"`
	Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
	Lastlogintime time.Time `orm:"type(datetime);null" form:"-"`
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now;type(datetime)"`
}

func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		_ = v.SetError("Repassword", "Does not matched password, repassword")
	}
}

func (u *User) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(u, field, fields...)
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}

func Users() orm.QuerySeter {
	var table User
	return orm.NewOrm().QueryTable(table).OrderBy("-Id")
}

func init() {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("dbprefix"),
		new(User))

}
