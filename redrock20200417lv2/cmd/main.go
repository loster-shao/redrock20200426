package main

import (
	"github.com/gin-gonic/gin"
	"redrock20200417lv2/router"
)

func main(){

	app := gin.Default()

	router.SetupRouter(app)

	app.Run(":8080")
}
