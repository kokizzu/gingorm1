package config

import "gorm.io/gorm"
import "gorm.io/driver/mysql"
import "github.com/kokizzu/gotro/L"

// mysql -u root -p -h 127.0.0.01 -P 3306
// CREATE DATABASE gingorm1;

var MysqlDsn = "root:password@tcp(127.0.0.1:3306)/gingorm1?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDB() *gorm.DB {
	DB, err := gorm.Open(mysql.Open(MysqlDsn), &gorm.Config{})
	L.PanicIf(err, `gorm.Open`, err)
	return DB
}
