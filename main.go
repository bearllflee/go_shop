package main

import (
	"viper/initialize"
)

func main() {
	initialize.MustConfig()
	initialize.MustRunWindowServer()
	// fmt.Println(global.CONFIG)
}
