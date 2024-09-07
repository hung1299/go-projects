package routes

import (
	"fmt"
	"hung1299/go-projects/caching-proxy/global"
	"hung1299/go-projects/caching-proxy/internal/routes/proxy"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	proxy.InitProxyRouter(r)

	if err := r.Run(fmt.Sprintf(":%v", global.Config.Port)); err != nil {
		fmt.Println("Can not start server at ", global.Config.Port)
		os.Exit(1)
	}
}
