package model

import (
	"fmt"
	"gorobbs/package/setting"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

const PAGE_SIZE int = 20

type Model struct {
	gorm.Model
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = setting.DatabaseSetting.Type
	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Forum{})
	db.AutoMigrate(&Thread{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Mythread{})
	db.AutoMigrate(&Attach{})
	db.AutoMigrate(&PostUpdateLog{})
	db.AutoMigrate(&MyFavourite{})
}

func CloseDB() {
	defer db.Close()
}

func GetDb() *gorm.DB {
	return db
}

// 自增
func Increment(table string, id int, colum string) error {
	//pcolum := colum + " + 1"
	return db.Exec("update ? set files_num = files_num + 1 where id = 1", "bbs_post").Error
}
