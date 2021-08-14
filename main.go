package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanagami1369/CodingShare/ent/migrate"
	"github.com/nanagami1369/CodingShare/ent/user"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/module"
)

func main() {
	sm := module.NewStupModule()
	// db
	config, err := sm.ReadConfigFromEnv()
	if err != nil {
		log.Panicf("read config err:&#v", err)
	}
	client, err := sm.ConnentDB(config)
	if err != nil {
		log.Panicf("open db err:&#v", err)
	}
	context := context.Background()
	uam := sm.GetUserAccountModule(client, context)
	defer client.Close()

	// migrate
	err = client.Schema.Create(
		context,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Panicf("open db err:&#v", err)
	}

	// set sample data
	client.User.Create().
		SetUserID("1821141").
		SetPassword("sample").
		SetAccountType(user.AccountTypeStudent).
		SetStudentNumber(1821141).
		Save(context)
	client.User.Create().
		SetUserID("tanaka").
		SetPassword("sample").
		SetAccountType(user.AccountTypeTeacher).
		Save(context)
	client.User.Create().
		SetUserID("suzuki").
		SetPassword("sample").
		SetAccountType(user.AccountTypeGeneral).
		Save(context)

	// router
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
			"message": "Login Ok Hello " + user.UserID + "!",
		})
	})
	router.Run(":8081")
}
