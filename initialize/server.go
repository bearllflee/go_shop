package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"viper/global"
	"viper/router"
)

func MustRunWindowServer() {
	engine := gin.Default()

	userGroup := router.UserGroup{}
	userGroup.InitUserRouters(engine)

	address := fmt.Sprintf(":%d", global.CONFIG.App.Port)
	fmt.Println("启动服务器，监听端口：", address)
	if err := engine.Run(address); err != nil {
		panic(err)
	}
}
