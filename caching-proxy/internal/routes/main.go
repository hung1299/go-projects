package routes

import (
	"fmt"
	"hung1299/go-projects/caching-proxy/global"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)
  
func InitRoute() {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request)
		c.JSON(http.StatusOK, gin.H{
		"message": "i see u",
		})
	})

	c := global.Config

	if err := r.Run(fmt.Sprintf(":%v", c.Port)); err != nil {
		fmt.Println("Can not start server at ", c.Port)
		os.Exit(1)
	}
}