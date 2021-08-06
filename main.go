package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/module"
)

func main() {
	sm := module.NewStupModule()
	uam, err := sm.GetUserAccountModule()
	if err != nil {
		panic(err)
	}
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
		user, err := uam.Login(loginRequest.Email, loginRequest.Password)
		if err != nil {
			// ログインに失敗したら失敗した事だけ伝える
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication failure",
			})
			return
		}
		// ログイン成功
		c.JSON(200, gin.H{
			"message": "Login Ok Hello " + user.Name + "!",
		})
	})
	router.Run(":8081")
}
