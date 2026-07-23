package app

import (
	"lms/internal/user"
	"net/http"
)

func NewRouter(h *user.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("GET /api/users", nil)
	mux.HandleFunc("GET /api/users/{userId}", h.Get)
	mux.HandleFunc("POST /api/users", h.Create)
	mux.HandleFunc("PUT /api/users/{userId}", nil)
	mux.HandleFunc("DELETE /api/users/{userId}", nil)

	return mux
}
