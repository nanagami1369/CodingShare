package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

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
	client, err := sm.ConnentDB()
	if err != nil {
		log.Panicln("open db err:", err)
	}
	context := context.Background()
	uam := sm.GetUserAccountModule(client, context)
	sem := sm.GetSessionModule(client, context)
	vim := sm.GetVideoModule(client, context)
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
	router, _ := sm.GetRouter(middleware)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.GET("/loadvideo/:id", func(c *gin.Context) {
		stringId := c.Param("id")
		id, err := strconv.Atoi(stringId)
		if err != nil {
			// 数値以外は404
			c.Status(http.StatusNotFound)
			return
		}

		video, err := vim.Load(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println("video load err :", err)
			return
		}
		if video == nil {
			// 存在しなければ404
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, video)
	})
	router.GET("/searchvideo", func(c *gin.Context) {
		query, isSet := c.GetQuery("q")
		if !isSet {
			c.Status(http.StatusBadRequest)
			return
		}
		videos, err := vim.Search(query)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, videos)
	})
	router.POST("/login", func(c *gin.Context) {
		loginRequest := &model.LoginRequest{}
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			// JSONが読み込めなかったらエラーを返す
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		AccountTypeStudent := user.AccountTypeStudent
		user, err := uam.Confirm(loginRequest)
		if err != nil {
			// 認証に失敗したら失敗した事だけ伝える
			c.JSON(http.StatusUnauthorized, "Authentication failure")
			loginLog.Printf("login err request: %v ,err: %v", loginRequest, err)
			return
		}
		// 学生は/loginからログインさせない
		if user.AccountType == AccountTypeStudent {
			loginLog.Printf("login err request: %v ,err:  student login stop from /login", loginRequest)
			c.String(http.StatusBadRequest, "400 Bad Request")
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
	router.POST("/auth", func(c *gin.Context) {
		store := sessions.Default(c)
		_, user, err := sem.Get(store)
		if err != nil {
			log.Println("get user err:", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		if user == nil {
			c.JSON(http.StatusOK, gin.H{
				"userId": "",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"userId": user.UserID,
		})
	})
	router.GET("uservideo/:userid", func(c *gin.Context) {
		userid := c.Param("userid")
		videos, err := vim.GetUserVideos(userid)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, videos)

	})
	// Kbook
	kbook := router.Group("/kbook")
	kbook.Use(middleware.CheckConnectFromKbookMiddleware())
	kbook.GET("/home", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome Kbook!")
	})
	kbook.POST("/home", func(c *gin.Context) {
		// フォームの取得
		c.Request.ParseForm()
		uid := c.Request.PostForm.Get("uid")
		// 学籍番号の取得
		userId := GetStudentNumber(uid)
		studentNumber, err := strconv.Atoi(userId)
		if err != nil {
			// 学籍番号で無かった(数字に変換できない値だった)場合
			c.String(http.StatusBadRequest, "400 Bad Request")
			loginLog.Printf("login err request: %v ,err: %v", uid, err)
			return
		}
		// アカウントの存在確認
		isExists, err := uam.Exists(userId)
		if err != nil {
			c.String(http.StatusInternalServerError, "500 InternalServerError")
			loginLog.Printf("login err request: %v ,err: %v", userId, err)
			return
		}
		if !isExists {
			// アカウントがなければ作ってからログイン
			signInRequest := &model.SignInRequest{
				Id:            userId,
				RowPassword:   "gabagaba password",
				AccountType:   user.AccountTypeStudent,
				StudentNumber: &studentNumber,
			}
			_, err := uam.SignIn(signInRequest)
			if err != nil {
				c.String(http.StatusInternalServerError, "500 InternalServerError")
				loginLog.Printf("singin err request: %v ,err: %v", userId, err)
				return
			}
		}
		// 認証
		loginRequest := &model.LoginRequest{
			Id:          userId,
			RowPassword: "gabagaba password",
		}
		user, err := uam.Confirm(loginRequest)
		if err != nil {
			// 認証に失敗したら失敗した事だけ伝える
			c.String(http.StatusForbidden, "403 Forbidden")
			loginLog.Printf("login err request: %v ,err: %v", userId, err)
			return
		}
		// 認証成功
		store := sessions.Default(c)
		dateOfExpiry := time.Now().AddDate(0, 0, 3)
		sem.Login(store, user, dateOfExpiry)
		c.HTML(http.StatusOK, "welcome_kbook.html", gin.H{
			"name": user.UserID,
		})
		loginLog.Println("login success request:", user.UserID)
	})
	private := router.Group("/private")
	private.Use(middleware.LoginCheckMiddleware())
	private.POST("/logout", func(c *gin.Context) {
		store := sessions.Default(c)
		user, err := sem.Logout(store)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println("logout err :", err)
			return
		}
		c.String(http.StatusOK, "ログアウトOK")
		loginLog.Println("logout success request:", user.UserID)
	})
	private.POST("/savevideo", func(c *gin.Context) {
		saveVideoRequest := &model.SaveVideoRequest{}
		if err := c.ShouldBindJSON(saveVideoRequest); err != nil {
			// JSONが読み込めなかったらエラーを返す
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		store := sessions.Default(c)
		_, user, err := sem.Get(store)
		if err != nil {
			log.Println("session get err:", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		_, err = vim.Save(saveVideoRequest, user)
		if err != nil {
			log.Println("save video err:", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusOK)
	})
	private.GET("/myvideo", func(c *gin.Context) {
		store := sessions.Default(c)
		_, user, err := sem.Get(store)
		if err != nil {
			log.Println("get user err:", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		videos, err := vim.GetUserVideos(user.UserID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
		}
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, videos)
	})
	private.DELETE("/removevideo/:videoId", func(c *gin.Context) {
		videoId := c.Param("videoId")
		log.Println(videoId)
		id, err := strconv.Atoi(videoId)
		if err != nil {
			// 数値以外は404
			c.Status(http.StatusBadRequest)
			return
		}
		store := sessions.Default(c)
		_, user, err := sem.Get(store)
		if err != nil {
			log.Println("get user err:", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		err = vim.Remove(id, user)
		if err != nil {
			log.Println("video removed err:", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusOK)
	})
	router.Run(":8080")
}

func GetStudentNumber(userId string) string {
	if utf8.RuneCountInString(userId) == 0 {
		return ""
	}
	return strings.TrimLeft(userId, "s")
}
