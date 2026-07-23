package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatalf("TODO %v", err)
	}

	if err := req.Validate(); err != nil {
		log.Fatalf("TODO %v", err)
	}

	user := req.ToModel()
	if err := h.svc.Create(r.Context(), user); err != nil {
		log.Fatalf("TODO %v", err)
	}

	resp := NewUserResponse(user)
	b, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("TODO %v", err)
	}
	w.Write(b)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := strconv.ParseInt(r.PathValue("userId"), 10, 64)
	if err != nil {
		log.Fatalf("TODO %v", err)
	}

	user, err := h.svc.FindByID(r.Context(), id)
	if err != nil {
		log.Fatalf("TODO %v", err)
	}

	resp := NewUserResponse(user)
	b, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("TODO %v", err)
	}
	w.Write(b)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var cursor int64
	if raw := r.URL.Query().Get("cursor"); raw != "" {
		parsed, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			log.Fatalf("TODO %v", err)
		}
		cursor = parsed
	}

	limit := 20
	if raw := r.URL.Query().Get("limit"); raw != "" {
		parsed, err := strconv.Atoi(raw)
		if err != nil {
			log.Fatalf("TODO %v", err)
		}
		limit = parsed
	}

	users, hasMore, err := h.svc.List(r.Context(), cursor, limit)
	if err != nil {
		log.Fatalf("TODO %v", err)
	}

	var nextCursor int64
	if hasMore {
		nextCursor = users[len(users)-1].ID
	}

	resp := NewUsersResponse(users, hasMore, nextCursor)
	b, err := json.Marshal(&resp)
	if err != nil {
		log.Fatalf("TODO %v", err)
	}
	w.Write(b)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := strconv.ParseInt(r.PathValue("userId"), 10, 64)
	if err != nil {
		log.Fatalf("TODO %v", err)
	}

	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatalf("TODO %v", err)
	}

	if err := req.Validate(); err != nil {
		log.Fatalf("TODO %v", err)
	}

	user := req.ToModel()
	user.ID = id

	if err := h.svc.Update(r.Context(), user); err != nil {
		log.Fatalf("TODO %v", err)
	}

	resp := NewUserResponse(user)
	b, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("TODO %v", err)
	}
	w.Write(b)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := strconv.ParseInt(r.PathValue("userId"), 10, 64)
	if err != nil {
		log.Fatalf("TODO %v", err)
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		log.Fatalf("TODO %v", err)
	}

	w.WriteHeader(http.StatusNoContent)
}
