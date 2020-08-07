package model

import (
	"github.com/jinzhu/gorm"
	"noodlenote/config"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func SetUp() {
	d, err := gorm.Open("sqlite3", "./gorm.db")

	if err != nil {
		panic(err)
	}

	//不加s建表
	d.SingularTable(true)
	d.LogMode(config.Cfg.AppConfig.DBLog)

	if config.Cfg.AppConfig.MakeMigration {
		migration(d) //迁移  首次创建数据库需要迁移创建表
	}

	db = d
}

func migration(DB *gorm.DB) {
	// 自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4")
	DB.AutoMigrate(&NoteBook{}).
		AutoMigrate(&Note{}).
		AutoMigrate(&NoteBookLink{})
}
