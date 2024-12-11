package main

import (
	"viper/global"
	"viper/initialize"
)

func main() {
	initialize.MustConfig()
	initialize.MustInitDB()
	initialize.AutoMigrate(global.DB)
	initialize.MustRunWindowServer()
	// fmt.Println(global.CONFIG)
}
