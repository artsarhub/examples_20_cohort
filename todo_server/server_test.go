package todo_server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"yp-examples/todo_server/internal/handler"
	"yp-examples/todo_server/internal/model"
	"yp-examples/todo_server/mocks"
)

func TestGetTodos(t *testing.T) {
	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	repository := mocks.NewMockRepository(ctrl)
	repository.EXPECT().GetTodos().Times(1).Return([]model.Todo{})
	todosHandler := handler.NewTodosHandler(repository)

	responseRecorder := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(todosHandler.GetTodos)

	handlerFunc.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handlerFunc returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "[]"
	if strings.TrimSpace(responseRecorder.Body.String()) != expected {
		t.Errorf("handlerFunc returned unexpected body: got %v want %v", responseRecorder.Body.String(), expected)
	}
}

func TestAddTodo(t *testing.T) {
	todo := model.Todo{
		Task: "Test task",
		Done: false,
	}
	payload, _ := json.Marshal(todo)
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	repository := mocks.NewMockRepository(ctrl)
	repository.EXPECT().AddTodo(todo).Times(1)
	todosHandler := handler.NewTodosHandler(repository)

	responseRecorder := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(todosHandler.AddTodo)

	handlerFunc.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handlerFunc returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateTodo(t *testing.T) {
	todo := model.Todo{
		ID:   uuid.NewString(),
		Task: "Test task",
		Done: true,
	}

	ctrl := gomock.NewController(t)
	repository := mocks.NewMockRepository(ctrl)
	repository.EXPECT().UpdateTodo(todo).Times(1).Return(todo, nil)
	todosHandler := handler.NewTodosHandler(repository)

	payload, _ := json.Marshal(todo)
	updateReq, _ := http.NewRequest("PUT", "/todos", bytes.NewBuffer(payload))
	responseRecorder := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(todosHandler.UpdateTodo)
	handlerFunc.ServeHTTP(responseRecorder, updateReq)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handlerFunc returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var updatedTodo model.Todo
	err := json.NewDecoder(responseRecorder.Body).Decode(&updatedTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, todo.Done, updatedTodo.Done, "handlerFunc returned unexpected body: got %+v want %+v", todo.Done, updatedTodo.Done)
}

func TestDeleteTodo(t *testing.T) {
	todoID := uuid.NewString()

	ctrl := gomock.NewController(t)
	repository := mocks.NewMockRepository(ctrl)
	repository.EXPECT().DeleteTodo(todoID).Times(1).Return(nil)
	todosHandler := handler.NewTodosHandler(repository)

	deleteReq, _ := http.NewRequest("DELETE", fmt.Sprintf("/todos?id=%s", todoID), nil)
	responseRecorder := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(todosHandler.DeleteTodo)
	handlerFunc.ServeHTTP(responseRecorder, deleteReq)

	if status := responseRecorder.Code; status != http.StatusNoContent {
		t.Errorf("handlerFunc returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}
