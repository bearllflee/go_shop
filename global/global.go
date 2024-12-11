package global

import (
	"gorm.io/gorm"
	"viper/config"
)

var (
	CONFIG config.Server
	DB     *gorm.DB
)
