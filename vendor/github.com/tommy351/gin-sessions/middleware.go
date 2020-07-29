package sessions

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
)

const (
	sessionName  = "sessionName"
	sessionStore = "sessionStore"
)

// Middleware helps you set up a session in context
func Middleware(name string, store Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(sessionName, name)
		c.Set(sessionStore, store)
		defer context.Clear(c.Request)
		c.Next()
	}
}

// Get returns a session in the context
func Get(c *gin.Context) Session {
	name := c.MustGet(sessionName).(string)
	store := c.MustGet(sessionStore).(Store)

	return &session{
		name:    name,
		request: c.Request,
		writer:  c.Writer,
		store:   store,
	}
}
