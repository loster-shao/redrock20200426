package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"redrock20200417lv2/mysql"
)

func Top(c *gin.Context)  {
	var movie []Movie
	db := mysql.DbConn()
	db.Where("ok = ?",1).Find(&movie)
	c.JSON(200,gin.H{"status":http.StatusOK,"message": movie})
}

func (movies Movie)Send(movie []Movie)  {
	//movie
}
