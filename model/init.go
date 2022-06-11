package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Database(connstring string) {
	dblogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(connstring), &gorm.Config{
		Logger: dblogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	fmt.Println("数据库连接成功")
	if err != nil {
		fmt.Println("connect db error", err)
	}
	mysqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	mysqlDB.SetMaxIdleConns(20)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	mysqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	mysqlDB.SetConnMaxLifetime(time.Second * 30)

	DB = db
	migration()
}

func migration() {
	// 自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}, &UserPassword{}, &VideoTable{}, &Favorite{}, &Comment{}, &Follow{})
}
