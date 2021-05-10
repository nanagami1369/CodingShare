package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/module"
)

func main() {
	sm := module.NewStupModule()
	db, err := sm.OpenDB()
	if err != nil {
		panic(err)
	}
	router := sm.GetRouter()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.Run(":8081")
	user := &model.User{}
	db.First(user)
	fmt.Println(user)
}
