package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"example.com/gopracs2-z5-borisovda/internal/student"
)

type Store interface {
	GetByID(ctx context.Context, id int64) (student.Student, error)
	GetByEmail(ctx context.Context, email string) (student.Student, error)
}

type Handler struct {
	store Store
}

func NewHandler(store Store) *Handler {
	return &Handler{store: store}
}

func RegisterRoutes(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/students", handler.GetStudentByID)
	mux.HandleFunc("/students/by-email", handler.GetStudentByEmail)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"scheme": "https",
	})
}

func (h *Handler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rawID := r.URL.Query().Get("id")
	if rawID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, ok := parsePositiveID(rawID)
	if !ok {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	st, err := h.store.GetByID(r.Context(), id)
	if err != nil {
		handleStoreError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, st)
}

func (h *Handler) GetStudentByEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rawEmail := r.URL.Query().Get("email")
	if rawEmail == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	email, ok := normalizeEmail(rawEmail)
	if !ok {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}

	st, err := h.store.GetByEmail(r.Context(), email)
	if err != nil {
		handleStoreError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, st)
}

func handleStoreError(w http.ResponseWriter, err error) {
	if errors.Is(err, student.ErrStudentNotFound) {
		http.Error(w, "student not found", http.StatusNotFound)
		return
	}

	http.Error(w, "internal server error", http.StatusInternalServerError)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}