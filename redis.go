package main

import (
	"fmt"
	"log"
	"time"

	redigo "github.com/garyburd/redigo/redis"
)

var (
	redisAddress string = "127.0.0.1:6379"
	redisConn    *redigo.Pool
)

func initRedis() {
	redisConn = &redigo.Pool{
		MaxIdle:     100,
		MaxActive:   100,
		IdleTimeout: 300 * time.Second,
		Wait:        true,
		Dial: func() (redigo.Conn, error) {
			return redigo.Dial("tcp", redisAddress)
		},
	}

	// do PING, if failed then Fatal
	if err := pingRedis(); err != nil {
		log.Fatal(err)
	}
}

func pingRedis() error {
	c := redisConn.Get()
	defer c.Close()
	if _, err := c.Do("PING"); err != nil {
		return err
	}
	return nil
}

// sample command to do a command to redis
func sampleDoCommand(command string) {
	c := redisConn.Get()
	defer c.Close()

	reply, err := c.Do(command)
	fmt.Println(reply, err)
}

func SaddCommand(key, member string) {
	c := redisConn.Get()
	defer c.Close()

	reply, err := c.Do("SADD", key, member)
	fmt.Println(reply, err)
}

func SRandMember(key string) ([]string, error) {
	c := redisConn.Get()
	defer c.Close()

	// Just dummy error
	tNow := time.Now().UnixNano()
	if tNow%2 == 0 {
		return []string{}, redigo.ErrPoolExhausted
	}

	reply, err := redigo.Strings(c.Do("SRANDMEMBER", key, 1))
	fmt.Println(reply, err)

	return reply, err
}

func DelRedisKey(key string) {
	c := redisConn.Get()
	defer c.Close()

	reply, err := c.Do("DEL", key)
	fmt.Println(reply, err)

}
