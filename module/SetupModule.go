package module

import (
	"context"
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/repository"
)

type SetupModule struct {
}

func NewStupModule() *SetupModule {
	return &SetupModule{}
}

func (sm *SetupModule) GetUserAccountModule(client *ent.Client, context context.Context) (module UserAccountModule) {
	r := repository.NewUserAccountRepository(context, client)
	return NewUserAccountModule(r)
}

func (sm *SetupModule) GetRouter() (router *gin.Engine, err error) {
	router = gin.Default()
	return router, nil
}

func (sm *SetupModule) ConnentDB(config *model.Config) (Client *ent.Client, err error) {
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

func (sm *SetupModule) ReadConfigFromEnv() (config *model.Config, err error) {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, errors.New("環境変数「DB_USER」が定義されていません")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return nil, errors.New("環境変数「DB_PASSWORD」が定義されていません")
	}
	dbIp := os.Getenv("DB_IP")
	if dbIp == "" {
		return nil, errors.New("環境変数「DB_IP」が定義されていません")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return nil, errors.New("環境変数「DB_PORT」が定義されていません")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, errors.New("環境変数「DB_NAME」が定義されていません")
	}
	config = &model.Config{
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBIp:       dbIp,
		DBPort:     dbPort,
		DBName:     dbName,
	}
	return config, nil
}
