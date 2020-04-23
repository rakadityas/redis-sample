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
