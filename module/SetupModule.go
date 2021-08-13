package module

import (
	"errors"
	"os"

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

func (sm *SetupModule) GetUserAccountModule(db *gorm.DB) (module UserAccountModule) {
	r := repository.NewUserAccountRepository(db)
	return NewUserAccountModule(r)
}

func (sm *SetupModule) GetRouter() (router *gin.Engine, err error) {
	router = gin.Default()
	return router, nil
}

func (sm *SetupModule) OpenDB(config *model.Config) (db *gorm.DB, err error) {
	connectURL := config.DBUser + ":" + config.DBPassword + "@tcp(" + config.DBIp + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(connectURL), &gorm.Config{})
}

func (sm *SetupModule) CloseDB(db *gorm.DB) {
	originalDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if err = originalDB.Close(); err != nil {
		panic(err)
	}
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
