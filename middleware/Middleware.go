package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	session sessions.Session
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionId := session.Get("session_id")
		if session_id, ok := sessionId.(string); ok && session_id == "Logind" {
			c.Next()
		} else {
			c.Status(http.StatusForbidden)
			c.Abort()
		}
	}
}
