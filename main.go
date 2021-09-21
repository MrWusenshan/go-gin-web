package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "go-gin-web/common"
	"go-gin-web/routers"
)

func main() {
	gin.SetMode(viper.GetString("server.mode"))
	app := gin.Default()
	routers.User(app)
	port := viper.GetString("server.port")
	panic(app.Run(port))
}
