package initialize

import (
	"fmt"

	"github.com/bearllflee/go_shop/global"
	"github.com/bearllflee/go_shop/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.OperationRecord{},
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

func MustInitRedis() {
	redisConf := global.CONFIG.Redis
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})
}
