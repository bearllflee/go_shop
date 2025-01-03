package main

import (
	"github.com/bearllflee/go_shop/global"
	"github.com/bearllflee/go_shop/initialize"
)

func main() {
	initialize.MustConfig()
	initialize.MustInitDB()
	initialize.AutoMigrate(global.DB)
	initialize.MustInitRedis()
	initialize.MustRunWindowServer()
	// fmt.Println(global.CONFIG)
}
