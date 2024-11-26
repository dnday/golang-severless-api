package handler

import "net/http"

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from /api/admin"))
}
