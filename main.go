package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/module"
)

func main() {
	sm := module.NewStupModule()
	config, err := sm.ReadConfigFromEnv()
	if err != nil {
		log.Panicf("read config err:&#v", err)
	}
	db, err := sm.OpenDB(config)
	if err != nil {
		log.Panicf("open db err:&#v", err)
	}
	uam := sm.GetUserAccountModule(db)
	router, _ := sm.GetRouter()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.POST("/login", func(c *gin.Context) {
		loginRequest := &model.LoginRequest{}
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			// JSONが読み込めなかったらエラーを返す
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		user, err := uam.Login(loginRequest.Id, loginRequest.Password)
		if err != nil {
			// ログインに失敗したら失敗した事だけ伝える
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication failure",
			})
			return
		}
		// ログイン成功
		c.JSON(200, gin.H{
			"message": "Login Ok Hello " + user.UserId + "!",
		})
	})
	router.Run(":8081")
}
