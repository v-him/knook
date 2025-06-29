package api

import (
	"net/http"
	"time"
)

const lichessApi = "https://lichess.org/api"

func useToken(token string, req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+token)
}

func waitMinute() {
	time.Sleep(time.Minute)
}
