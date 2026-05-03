package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/gopracs2-z5-borisovda/internal/student"
)

type fakeStore struct {
	byID       map[int64]student.Student
	byEmail    map[string]student.Student
	idCalls    int
	emailCalls int
}

func (f *fakeStore) GetByID(ctx context.Context, id int64) (student.Student, error) {
	f.idCalls++

	st, ok := f.byID[id]
	if !ok {
		return student.Student{}, student.ErrStudentNotFound
	}

	return st, nil
}

func (f *fakeStore) GetByEmail(ctx context.Context, email string) (student.Student, error) {
	f.emailCalls++

	st, ok := f.byEmail[email]
	if !ok {
		return student.Student{}, student.ErrStudentNotFound
	}

	return st, nil
}

func TestHealthOK(t *testing.T) {
	store := &fakeStore{}
	handler := NewHandler(store)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	handler.Health(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestGetStudentByIDOK(t *testing.T) {
	expected := student.Student{
		ID:         1,
		FullName:   "Борисов Денис Александрович",
		StudyGroup: "ЭФМО-01-25",
		Email:      "borisov@example.com",
	}

	store := &fakeStore{
		byID: map[int64]student.Student{
			1: expected,
		},
	}

	handler := NewHandler(store)

	req := httptest.NewRequest(http.MethodGet, "/students?id=1", nil)
	rr := httptest.NewRecorder()

	handler.GetStudentByID(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var actual student.Student
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("expected %+v, got %+v", expected, actual)
	}
}

func TestGetStudentByIDRejectsSQLInjection(t *testing.T) {
	store := &fakeStore{}
	handler := NewHandler(store)

	req := httptest.NewRequest(http.MethodGet, "/students?id=1%20OR%201%3D1", nil)
	rr := httptest.NewRecorder()

	handler.GetStudentByID(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}

	if store.idCalls != 0 {
		t.Fatalf("expected 0 store calls, got %d", store.idCalls)
	}
}

func TestGetStudentByEmailOK(t *testing.T) {
	expected := student.Student{
		ID:         1,
		FullName:   "Иванов Иван Иванович",
		StudyGroup: "ИВБО-01-25",
		Email:      "ivanov@example.com",
	}

	store := &fakeStore{
		byEmail: map[string]student.Student{
			"ivanov@example.com": expected,
		},
	}

	handler := NewHandler(store)

	req := httptest.NewRequest(http.MethodGet, "/students/by-email?email=ivanov@example.com", nil)
	rr := httptest.NewRecorder()

	handler.GetStudentByEmail(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var actual student.Student
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("expected %+v, got %+v", expected, actual)
	}
}

func TestGetStudentByEmailRejectsInvalidEmail(t *testing.T) {
	store := &fakeStore{}
	handler := NewHandler(store)

	req := httptest.NewRequest(http.MethodGet, "/students/by-email?email=ivanov@example.com%27%20OR%20%271%27%3D%271", nil)
	rr := httptest.NewRecorder()

	handler.GetStudentByEmail(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}

	if store.emailCalls != 0 {
		t.Fatalf("expected 0 store calls, got %d", store.emailCalls)
	}
}