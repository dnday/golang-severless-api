package handler

import (
	"net/http"
)

func UsrerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World from /api/user!"))
}
