package serverredis

import (
	"time"

	"ssf4g/common/tlog"

	"github.com/garyburd/redigo/redis"
)

type ServerRedis struct {
	_pool *redis.Pool
}

// Func - 初始化ServerDao
func (dao *ServerRedis) InitServerRedis(redisurl string, maxidleconn, timeout int, redisauth string) *tlog.ErrData {
	dao._pool = &redis.Pool{
		MaxIdle: maxidleconn,

		IdleTimeout: time.Duration(timeout) * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisurl)

			if err != nil {
				return nil, err
			}

			if redisauth == "" {
				return c, nil
			}

			if _, err := c.Do("AUTH", redisauth); err != nil {

				c.Close()
				return nil, err
			}

			return c, nil
		},
		/*
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		*/
	}

	return nil
}
