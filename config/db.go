package config

import (
	"ginDemo/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("初始化数据库失败->", err)
		return
	}
	sqlDB, err2 := db.DB()
	if err2 != nil {
		log.Fatal("初始化数据库失败->", err2)
		return
	}
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)       //设置连接池最大链接的空闲数量
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)       //设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(AppConfig.Database.ConnMaxLifetime) //每个连接的最大连接时长

	global.Db = db
}
