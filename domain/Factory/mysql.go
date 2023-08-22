package Factory

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// GetDb 初始化数据库连接
func GetDb() *gorm.DB {
	dsn := "root:root@galaxy1024@tcp(localhost:3306)/bingcool?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	dbOnce.Do(func() {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			// 处理数据库连接错误
			panic(err)
		}

		// 设置连接池大小
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		} else {
			sqlDB.SetMaxIdleConns(3)
			sqlDB.SetMaxOpenConns(5)
		}
	})
	return db
}
