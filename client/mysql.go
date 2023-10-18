package client

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web_server_2.0/config"
	data "web_server_2.0/data/mysql"
)

func NewMysql(env *config.Env) *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		env.Mysql.User,
		env.Mysql.Password,
		//env.Mysql.Host,
		env.Mysql.Host,
		env.Mysql.Port,
		env.Mysql.Database,
		env.Mysql.Charset)

	//dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
	//	"root",
	//	"12345678",
	//	"db",
	//	"13306",
	//	"app",
	//	"utf8mb4",
	//)
	fmt.Println("dns ->", dns)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("db ->", db, err)
	useAutoMigrate(db)
	return db
}

func useAutoMigrate(db *gorm.DB) {
	//db.AutoMigrate(&model.User{})
	data.NewMysqlTable(db)
}
