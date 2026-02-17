package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type Admin struct {
	Id        int       `orm:"pk;auto"`
	Username  string    `orm:"size(30);unique"`
	Password  string    `orm:"size(100)"`
	Nickname  string    `orm:"size(50)"`
	Status    int       `orm:"default(1)"`
	LoginTime time.Time `orm:"null;type(datetime)"`
}

func (a *Admin) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hash)
	return nil
}

func (a *Admin) CheckPassword(pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pwd)) == nil
}

func init() {
	dbPath := "./data/blog.db"
	os.MkdirAll("./data", 0755)
	orm.RegisterDataBase("default", "sqlite3", dbPath)
	orm.RegisterModel(new(Admin), new(Article), new(Game), new(Tool))
	orm.BootStrap()
	orm.RunSyncdb("default", false, true)

	o := orm.NewOrm()
	var admin Admin
	err := o.QueryTable("admin").Filter("username", "admin").One(&admin)
	if err != nil {
		defaultAdmin := Admin{
			Username: "admin",
			Password: "admin123",
			Nickname: "管理员",
		}
		defaultAdmin.EncryptPassword()
		_, err := o.Insert(&defaultAdmin)
		if err != nil {
			logs.Error("创建默认管理员失败:", err)
		}
	}
}
