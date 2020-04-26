package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"redrock20200417lv2/mysql"
	"strconv"
)

func Find(c *gin.Context)  {
	var person  []Person
	var student []Students
	var class   []Class
	xh := c.PostForm("id")
	id, err := strconv.Atoi(xh)
	if err != nil{
		fmt.Println("error!!:",err)
		return
	}
	if len(xh) != 10 || id > 2019215203 || id < 2019210001{
		fmt.Println("查无此人")
		c.JSON(404, gin.H{"status": 404, "message": "查无此人"})
		return
	}
	db := mysql.DbConn()
	db.Where("xh=?", xh).Find(&person)
	for i := 0; i < len(person); i++ {
		class = append(class, Class{
			Name:    person[i].Name,
			Class:   person[i].Class,
			Bx:      person[i].Bx,
			Status:  person[i].Status,
			Time:    person[i].Time,
			Where:   person[i].Where,
			Teacher: person[i].Teacher,
		})
	}
	name := person[0].Stu
	student = append(student, Students{
		Stu:   name,
		Xh:    id,
		Class: class,
	})
	c.JSON(200, gin.H{"status": 200, "课表": student})
	fmt.Println(person)
}

