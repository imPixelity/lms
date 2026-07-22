package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	svc *service
}

func NewHandler(svc *service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		log.Fatalf("TODO %v", err)
	}

	user, err := h.svc.get(r.Context(), id)
	if err != nil {
		log.Fatalf("TODO %v", err)
	}

	resp := UserResponse{
		ID:       user.id,
		Username: user.username,
		Email:    user.email,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("TODO %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
