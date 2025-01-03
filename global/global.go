package global

import (
	"github.com/bearllflee/go_shop/config"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var (
	CONFIG config.Server
	DB     *gorm.DB
	Cron   *cron.Cron
	Redis  *redis.Client
)
