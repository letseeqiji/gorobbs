package queue

import (
	"encoding/json"
	"gorobbs/package/setting"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

type RegEmailStruct struct {
	Host  string `json:"host"`
	Email string `json:"email"`
}

func init() {
	Setup()
}

// Setup Initialize the Redis instance
func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.RedisMaxidle,
		MaxActive:   setting.RedisSetting.RedisMaxActive,
		IdleTimeout: setting.RedisSetting.RedisIdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.RedisHost, redis.DialDatabase(13))
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.RedisPassword != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.RedisPassword); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

// Set a key/value
func Set(list string, data interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("LPUSH", list, value)
	if err != nil {
		return err
	}

	return nil
}

// Get get a key
func Get(list string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("RPOP", list))
	if err != nil {
		return nil, err
	}

	return reply, nil
}
