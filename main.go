package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/nanagami1369/CodingShare/ent/migrate"
	"github.com/nanagami1369/CodingShare/ent/user"
	"github.com/nanagami1369/CodingShare/middleware"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/setup"
)

func main() {
	sm := setup.NewSetupManager()
	// db
	config, err := sm.ReadConfigFromEnv()
	if err != nil {
		log.Panicln("read config err:", err)
	}
	client, err := sm.ConnentDB(config)
	if err != nil {
		log.Panicln("open db err:", err)
	}
	context := context.Background()
	uam := sm.GetUserAccountModule(client, context)
	sem := sm.GetSessionModule(client, context)
	defer client.Close()
	middleware := middleware.NewMiddleware(sem)

	// migrate
	err = client.Schema.Create(
		context,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Panicln("open db err:", err)
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
	router, _ := sm.GetRouter(config, middleware)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.POST("/login", func(c *gin.Context) {
		loginRequest := &model.LoginRequest{}
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			// JSONが読み込めなかったらエラーを返す
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		user, err := uam.Confirm(loginRequest)
		if err != nil {
			// 認証に失敗したら失敗した事だけ伝える
			c.JSON(http.StatusUnauthorized, "Authentication failure")
			loginLog.Printf("login err request: %v ,err: %v", loginRequest, err)
			return
		}
		// 認証成功
		store := sessions.Default(c)
		dateOfExpiry := time.Now().AddDate(0, 0, 3)
		sem.Login(store, user, dateOfExpiry)
		c.JSON(200, gin.H{
			"message": "Login Ok Hello " + user.UserID + "!",
		})
		loginLog.Println("login success request:", user.UserID)
	})

	api := router.Group("/api")
	api.Use(middleware.LoginCheckMiddleware())
	api.POST("/islogin", func(c *gin.Context) {
		c.String(http.StatusOK, "ログイン済み")
	})
	api.POST("/auth", func(c *gin.Context) {
		store := sessions.Default(c)
		_, user, err := sem.Get(store)
		if err != nil {
			log.Println("get user err:", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"userId": user.UserID,
		})
	})
	api.POST("/logout", func(c *gin.Context) {
		store := sessions.Default(c)
		user, err := sem.Logout(store)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println("logout err :", err)
		}
		c.String(http.StatusOK, "ログアウトOK")
		loginLog.Println("logout success request:", user.UserID)
	})

	router.RunTLS(":8081",
		config.CertificateFilePath,
		config.KeyFilePath)
}
