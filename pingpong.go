package pingpong

import (
	"io"
	"net/http"
	"strconv"
)

const PONG = "Pong"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(PONG)))
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, PONG)
}
