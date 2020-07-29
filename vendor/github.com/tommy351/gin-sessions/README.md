# gin-sessions

[![Build Status](https://travis-ci.org/tommy351/gin-sessions.svg?branch=master)](https://travis-ci.org/tommy351/gin-sessions)

Session middleware for [Gin].

## Installation

``` bash
$ go get github.com/tommy351/gin-sessions
```

## Usage

``` go
import (
    "github.com/gin-gonic/gin"
    "github.com/tommy351/gin-sessions"
)

func main(){
    g := gin.New()
    store := sessions.NewCookieStore([]byte("secret123"))
    g.Use(sessions.Middleware("my_session", store))
    
    g.GET("/set", func(c *gin.Context){
        session := sessions.Get(c)
        session.Set("hello", "world")
        session.Save()
    })
    
    g.GET("/get", func(c *gin.Context){
        session := sessions.Get(c)
        session.Get("hello")
    })
}
```

[Gin]: http://gin-gonic.github.io/gin/
