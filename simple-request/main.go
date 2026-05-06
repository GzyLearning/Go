package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Ggo",
			"code": 200,
		})
	})
	r.GET("/hello/:name/*houmain", func(c *gin.Context) {
		name := c.Param("name")
		houmain := c.Param("houmain")
		c.JSON(http.StatusOK, gin.H{
			"message": "你好 " + name + "houmian" + houmain,
		})
	})
	r.GET("/welcome", func(c *gin.Context) {
		who := c.DefaultQuery("who", "郭正宇")
		do := c.Query("do")
		c.String(http.StatusOK, "Hello %s %s", who, do)
	})
	r.POST("/post", func(c *gin.Context) {
		form := c.DefaultPostForm("name", "郭正宇")
		value := c.PostForm("do")
		c.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"who":    form,
			"value":  value,
		})

	})
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, "C:\\Users\\15843\\Desktop\\"+"123.txt")
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	v1 := r.Group("/v1")
	{
		v1.POST("/upload", func(c *gin.Context) {
			file, _ := c.FormFile("file")
			log.Println(file.Filename)
			c.SaveUploadedFile(file, "C:\\Users\\15843\\Desktop\\"+"123.txt")
			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		})
	}
	r.Run(":8080")
}
