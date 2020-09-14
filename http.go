package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	key = "key"
)

func initHandler() {
	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/create_prize_pool", handleCreatePrizePool)
	http.HandleFunc("/get_prize", handleGetPrize)
	// TODO: Adding more APIs

	fmt.Println("serving in 8181 ")
	http.ListenAndServe(":8181", nil)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	err := pingRedis()
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte("pong"))

	return
}

func handleCreatePrizePool(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var prizes Prizes
	err = json.Unmarshal(jsonData, &prizes)
	if err != nil {
		log.Fatal(err)
		return
	}

	DelRedisKey(key)

	for _, p := range prizes.Prize {
		for i := 0; i < p.Percentage; i++ {
			SaddCommand(key, fmt.Sprintf("%03d%s", i, p.Name))
		}
	}

	return
}

func handleGetPrize(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	userIDStr := queryValues.Get("user_id")

	userID, _ := strconv.Atoi(userIDStr)

	if userID <= 0 {
		w.Write([]byte("your userID is not valid"))
		return
	}

	arrayValue := SRandMember(key)

	fmt.Println(arrayValue[0][3:])
	w.Write([]byte(arrayValue[0][3:]))
	return
}
