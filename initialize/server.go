package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"viper/global"
)

func MustRunWindowServer() {
	engine := gin.Default()

	address := fmt.Sprintf(":%d", global.CONFIG.App.Port)
	fmt.Println("启动服务器，监听端口：", address)
	if err := engine.Run(address); err != nil {
		panic(err)
	}
}
