package module

import (
	"context"
	"errors"

	"os"

	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
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

func (sm *SetupModule) readEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("環境変数「" + key + "」が定義されていません")
	}
	return value, nil

}

func (sm *SetupModule) ReadConfigFromEnv() (*model.Config, error) {
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
	certificateFilePath, err := sm.readEnv("CODING_SHARE_CERTIFICATE_FILE_PATH")
	if err != nil {
		return nil, err
	}
	keyFilePath, err := sm.readEnv("CODING_SHARE_KEY_FILE_PATH")
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
		CertificateFilePath: certificateFilePath,
		KeyFilePath:         keyFilePath,
	}
	return config, nil
}
