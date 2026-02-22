package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"todo-app/internal/model"
	"todo-app/internal/response"
	"todo-app/internal/service"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{service: s}
}

func (h *TodoHandler) RegisterRoutes(r *httprouter.Router) {
	r.POST("/todos", h.Create)
	r.GET("/todos/:id", h.Get)
	r.GET("/todos", h.List)
	r.PUT("/todos/:id", h.Update)
	r.DELETE("/todos/:id", h.Delete)
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req struct {
		Task    string    `json:"task"`
		DueDate time.Time `json:"due_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.Error("INVALID_JSON", "invalid request body"))
		return
	}

	todo, err := h.service.Create(r.Context(), req.Task, req.DueDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.Error("INVALID_INPUT", err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.Success(todo))
}

func (h *TodoHandler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	todo, err := h.service.Get(r.Context(), ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response.Error("NOT_FOUND", "todo not found"))
		return
	}
	json.NewEncoder(w).Encode(response.Success(todo))
}

func (h *TodoHandler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	include := r.URL.Query().Get("include_completed") == "true"

	todos, err := h.service.List(r.Context(), include)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.Error("INTERNAL_ERROR", "failed to fetch todos"))
		return
	}
	json.NewEncoder(w).Encode(response.Success(todos))
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var t model.Todo

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.Error("INVALID_INPUT", err.Error()))
		return
	}

	t.ID = ps.ByName("id")

	if err := h.service.Update(r.Context(), t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.Error("UPDATE_FAILED", err.Error()))
		return
	}
	json.NewEncoder(w).Encode(response.Success(t))
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := h.service.Delete(r.Context(), ps.ByName("id")); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response.Error("NOT_FOUND", err.Error()))
		return
	}
	json.NewEncoder(w).Encode(response.Success(map[string]interface{}{"message": "todo deleted"}))
}
