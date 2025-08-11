package bootstrap

import (
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/dao/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func NewDb() *gorm.DB {
	initDb()
	return db
}

func initDb() {
	var err error

	sql := config.Get().Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		sql.Username,
		sql.Password,
		sql.Host,
		sql.Port,
		sql.Dbname,
		sql.Charset,
	)

	fmt.Println(dsn)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}))

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// TODO : 连接池配置 , 暂时写死
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = migrate()
	if err != nil {
		panic(err)
	}
}

func migrate() error {
	return db.AutoMigrate(
		&model.User{},
		&model.News{},
		&model.NewsCategories{},
	)
}
