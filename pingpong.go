package pingpong

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

const PONG = "Pong"

type jsonRes struct {
	Msg       string    `json:"msg"`
	CreatedAt time.Time `json:"created_at"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(PONG)))
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, PONG)
}

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
