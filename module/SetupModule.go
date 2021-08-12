package module

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nanagami1369/CodingShare/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SetupModule struct {
}

func NewStupModule() *SetupModule {
	return &SetupModule{}
}

func (sm *SetupModule) GetUserAccountModule() (module UserAccountModule, err error) {
	db, err := sm.openDB()
	if err != nil {
		return nil, err
	}
	r := repository.NewUserAccountRepository(db)
	return NewUserAccountModule(r), nil
}

func (sm *SetupModule) GetRouter() (router *gin.Engine, err error) {
	router = gin.Default()
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
