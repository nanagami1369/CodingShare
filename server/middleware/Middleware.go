package middleware

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nanagami1369/CodingShare/module"
)

type Middleware struct {
	sem module.SessionModule
}

func NewMiddleware(sem module.SessionModule) *Middleware {
	return &Middleware{sem: sem}
}

func (m *Middleware) LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		isLogin, err := m.sem.IsLogin(session)
		if err != nil {
			// セッション周りのエラーはログに吐き出す
			log.Println("session err :", err)
		}
		if isLogin {
			c.Next()
		} else {
			c.Status(http.StatusForbidden)
			c.Abort()
		}
	}
}