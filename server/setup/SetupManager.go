package setup

import (
	"context"
	"errors"
	"net/http"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/middleware"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/module"
	"github.com/nanagami1369/CodingShare/repository"
)

type SetupManager struct {
}

func NewSetupManager() *SetupManager {
	return &SetupManager{}
}

func (sm *SetupManager) GetUserAccountModule(client *ent.Client, context context.Context) module.UserAccountModule {
	r := repository.NewUserAccountRepository(context, client)
	return module.NewUserAccountModule(r)
}

func (sm *SetupManager) GetSessionModule(client *ent.Client, context context.Context) module.SessionModule {
	r := repository.NewSessionRepository(context, client)
	return module.NewSessionModule(r)
}

func (sm *SetupManager) GetRouter(config *model.Config, middleware *middleware.Middleware) (router *gin.Engine, err error) {
	router = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		config.WebUrl,
	}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{
		"Access-Control-Allow-Credentials",
	}
	router.Use(cors.New(corsConfig))
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Path:     "/",
		// 60s * 60 = 1h
		// 1h * 24 = 1day
		// 24h * 3 = 3day
		MaxAge: 60 * 60 * 24 * 3,
	})
	router.Use(sessions.Sessions("codingshare", store))
	return router, nil
}

func (sm *SetupManager) ConnentDB(config *model.Config) (Client *ent.Client, err error) {
	dbConfig := &mysql.Config{
		User:                 config.DBUser,
		Passwd:               config.DBPassword,
		Net:                  "tcp",
		Addr:                 config.DBIp + ":" + config.DBPort,
		DBName:               config.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	return ent.Open("mysql", dbConfig.FormatDSN())
}

func (sm *SetupManager) readEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("環境変数「" + key + "」が定義されていません")
	}
	return value, nil

}

func (sm *SetupManager) ReadConfigFromEnv() (*model.Config, error) {
	dbUser, err := sm.readEnv("CODING_SHARE_DB_USER")
	if err != nil {
		return nil, err
	}
	dbPassword, err := sm.readEnv("CODING_SHARE_DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	dbIp, err := sm.readEnv("CODING_SHARE_DB_IP")
	if err != nil {
		return nil, err
	}
	dbPort, err := sm.readEnv("CODING_SHARE_DB_PORT")
	if err != nil {
		return nil, err
	}
	dbName, err := sm.readEnv("CODING_SHARE_DB_NAME")
	if err != nil {
		return nil, err
	}
	apiUrl, err := sm.readEnv("CODING_SHARE_API_URL")
	if err != nil {
		return nil, err
	}
	webUrl, err := sm.readEnv("CODING_SHARE_WEB_URL")
	if err != nil {
		return nil, err
	}
	config := &model.Config{
		DBUser:              dbUser,
		DBPassword:          dbPassword,
		DBIp:                dbIp,
		DBPort:              dbPort,
		DBName:              dbName,
		ApiUrl:              apiUrl,
		WebUrl:              webUrl,
	}
	return config, nil
}
