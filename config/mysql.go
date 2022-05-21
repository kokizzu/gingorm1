package config

import "gorm.io/gorm"
import "gorm.io/driver/mysql"
import "github.com/kokizzu/gotro/L"

func ConnectDB() *gorm.DB {
	const dsn = "root:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	L.PanicIf(err, `gorm.Open`, err)
	return DB
}
