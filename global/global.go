package global

import (
	"github.com/bearllflee/go_shop/config"
	"gorm.io/gorm"
)

var (
	CONFIG config.Server
	DB     *gorm.DB
)
