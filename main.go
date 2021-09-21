package main

import (
	"github.com/gin-gonic/gin"
	_ "go-gin-web/common"
	"go-gin-web/routers"
)

func main() {
	app := gin.Default()
	routers.User(app)
	panic(app.Run(":8000"))
}
