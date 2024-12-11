package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"viper/global"
	"viper/model"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
	)
}

func MustInitDB() {
	mysqlConf := global.CONFIG.MySQL
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database, mysqlConf.Config)
	fmt.Println(dataSource)
	db, err := gorm.Open(mysql.Open(dataSource))
	if err != nil {
		panic(err)
	}
	global.DB = db
}
