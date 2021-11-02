package setup

import (
	"context"
	"errors"
	"net/http"

	"os"

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

func (sm *SetupManager) GetVideoModule(client *ent.Client, context context.Context) module.VideoModule {
	r := repository.NewVideoRepository(context, client)
	return module.NewVideoModule(r)
}

func (sm *SetupManager) GetRouter(middleware *middleware.Middleware) (router *gin.Engine, err error) {
	router = gin.Default()
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
		Addr:                 config.DBHost,
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
	dbUser, err := sm.readEnv("MYSQL_USER")
	if err != nil {
		return nil, err
	}
	dbPassword, err := sm.readEnv("MYSQL_PASSWORD")
	if err != nil {
		return nil, err
	}
	dbHost, err := sm.readEnv("MYSQL_HOST")
	if err != nil {
		return nil, err
	}
	dbName, err := sm.readEnv("MYSQL_DATABASE")
	if err != nil {
		return nil, err
	}
	config := &model.Config{
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBHost:     dbHost,
		DBName:     dbName,
	}
	return config, nil
}
