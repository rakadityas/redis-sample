package main

import (
	"log"

	"github.com/DataDog/datadog-go/statsd"
)

var (
	ddStats *statsd.Client
)

func initDD() {

	var err error

	ddStats, err = statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
}
