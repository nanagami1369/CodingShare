package module

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SetupModule struct {
}

func NewStupModule() *SetupModule {
	return &SetupModule{}
}

func (sm *SetupModule) getUserAccountModule() (module UserAccountModule, err error) {
	db, err := sm.openDB()
	if err != nil {
		return nil, err
	}
	r := repository.NewUserAccountRepository(db)
	return NewUserAccountRepository(r), nil
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["userId"],
		"name":   claims["name"],
		"text":   "Hello World.",
	})
}

func (sm *SetupModule) GetRouter() (router *gin.Engine, err error) {

	uam, err := sm.getUserAccountModule()
	if err != nil {
		return nil, err
	}
	securityKey := os.Getenv("SECURITY_KEY")
	if securityKey == "" {
		return nil, errors.New("環境変数「SECURITY_KEY」が定義されていません")
	}
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "api",
		Key:        []byte(securityKey),
		Timeout:    time.Minute * 15,
		MaxRefresh: time.Hour,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			loginRequest := &model.LoginRequest{}
			if err := c.ShouldBindJSON(loginRequest); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			user, err := uam.Login(loginRequest.Email, loginRequest.Password)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return user, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*model.User); ok {
				return jwt.MapClaims{"userId": user.UserId, "name": user.Name, "type": user.Type, "studentNumber": user.StudentNumber}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{UserId: claims["userId"].(float64), Name: claims["name"].(string), Type: claims["type"].(string), StudentNumber: claims["studentNumber"].(float64), Password: ""}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if user, ok := data.(*model.User); ok &&
				(user.Type == "student" || user.Type == "teacher" || user.Type == "general") {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		IdentityKey:    "userId",
		TokenLookup:    "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:  "Bearer",
		TimeFunc:       time.Now,
		SendCookie:     true,
		SecureCookie:   true,
		CookieHTTPOnly: true,
		CookieDomain:   "localhost",
		CookieName:     "token",
		CookieSameSite: http.SameSiteDefaultMode,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	router = gin.Default()
	router.POST("/login", authMiddleware.LoginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	api := router.Group("/api")
	// Refresh time can be longer than token timeout
	api.POST("/refresh_token", authMiddleware.RefreshHandler)
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.POST("/hello", helloHandler)
	}
	return router, nil
}

func (sm *SetupModule) openDB() (db *gorm.DB, err error) {
	url, err := sm.createConnectURLFromEnv()
	if err != nil {
		return nil, err
	}
	return gorm.Open(mysql.Open(url), &gorm.Config{})
}

func (sm *SetupModule) createConnectURLFromEnv() (connectURL string, err error) {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return "", errors.New("環境変数「DB_USER」が定義されていません")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return "", errors.New("環境変数「DB_PASSWORD」が定義されていません")
	}
	dbIp := os.Getenv("DB_IP")
	if dbIp == "" {
		return "", errors.New("環境変数「DB_IP」が定義されていません")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return "", errors.New("環境変数「DB_PORT」が定義されていません")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return "", errors.New("環境変数「DB_NAME」が定義されていません")
	}
	connectURL = dbUser + ":" + dbPassword + "@tcp(" + dbIp + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return connectURL, nil
}
