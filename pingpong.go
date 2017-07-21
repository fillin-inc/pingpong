package pingpong

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

// PONG is a string for response
const PONG = "Pong"

type jsonRes struct {
	Msg       string    `json:"msg"`
	CreatedAt time.Time `json:"created_at"`
}

// Handler is a HandlerFunc that returns a response in text/plain format
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(PONG)))
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, PONG)
}

// HandlerJSON is a HandlerFunc that returns a response in application/json format
func HandlerJSON(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(jsonRes{
		Msg:       PONG,
		CreatedAt: time.Now(),
	})
	resStr := string(res)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resStr)))
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, resStr)
}
