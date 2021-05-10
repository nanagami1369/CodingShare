package module

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SetupModule struct {
}

func NewStupModule() *SetupModule {
	return &SetupModule{}
}

func (cm *SetupModule) GetRouter() *gin.Engine {
	return gin.Default()
}

func (cm *SetupModule) OpenDB() (db *gorm.DB, err error) {
	url, err := cm.createConnectURLFromEnv()
	if err != nil {
		return nil, err
	}
	return gorm.Open(mysql.Open(url), &gorm.Config{})
}

func (cm *SetupModule) createConnectURLFromEnv() (connectURL string, err error) {
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
		return "", errors.New("環境変数「DB_USER」が定義されていません")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return "", errors.New("環境変数「DB_USER」が定義されていません")
	}
	connectURL = dbUser + ":" + dbPassword + "@tcp(" + dbIp + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return connectURL, nil
}
