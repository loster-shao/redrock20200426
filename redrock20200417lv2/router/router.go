package router

import (
	"github.com/gin-gonic/gin"
	"redrock20200417lv2/controller"
)

func SetupRouter(app *gin.Engine) {
	app.POST("/douban",controller.Douban)
	app.POST("/jwzx",controller.Jwzx)
	app.POST("/find", controller.Find)
	app.POST("/top",  controller.Top)
}
