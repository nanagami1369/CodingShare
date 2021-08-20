package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
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
	sampleStudenNumber := 1821141
	uam.SignIn(&model.SignInRequest{
		Id:            "1821141",
		RowPassword:   "sampleData",
		AccountType:   user.AccountTypeStudent,
		StudentNumber: &sampleStudenNumber,
	})
	uam.SignIn(&model.SignInRequest{
		Id:          "tanaka",
		RowPassword: "sampleData",
		AccountType: user.AccountTypeTeacher,
	})
	uam.SignIn(&model.SignInRequest{
		Id:          "suzuki",
		RowPassword: "sampleData",
		AccountType: user.AccountTypeGeneral,
	})

	// logger
	loginLog := log.New(os.Stdout, "[LOGIN]", log.LstdFlags|log.LUTC)

	// router
	router, _ := sm.GetRouter(config)
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
		user, err := uam.Login(loginRequest)
		if err != nil {
			// ログインに失敗したら失敗した事だけ伝える
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication failure",
			})
			loginLog.Printf("login err request: %v ,err: %v", loginRequest, err)
			return
		}
		// ログイン成功
		session := sessions.Default(c)
		session.Set("session_id", "Logind")
		session.Save()
		c.JSON(200, gin.H{
			"message": "Login Ok Hello " + user.UserID + "!",
		})
		loginLog.Printf("login success request: %v, user: %v", loginRequest, user.UserID)
	})
	router.RunTLS(":8081",
		config.CertificateFilePath,
		config.KeyFilePath)
}
