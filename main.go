package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println(action)
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})

	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "kaindy")
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})

	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.String(http.StatusOK, fmt.Sprintf("types: %s, username: %s, password: %s", types, username, password))
	})

	r.POST("/login", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if login.Username != "root" || login.Password != "password" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "304",
			})
			return
		}

		fmt.Println(login)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
		})
	})

	_ = r.Run(":8080")
}
