package todo_server

import (
	"log"
	"net/http"

	handler2 "examples_20_cohort/todo_server/internal/handler"
	"examples_20_cohort/todo_server/internal/repository"
	"examples_20_cohort/todo_server/pkg"
)

func RunServer() {
	repo := repository.NewInMemoryRepo()
	handler := handler2.NewTodosHandler(repo)

	http.Handle("/todos", pkg.CorsMiddleware(pkg.LoggingMiddleware(http.HandlerFunc(handler.GetTodos))))
	http.Handle("/add", pkg.CorsMiddleware(pkg.LoggingMiddleware(http.HandlerFunc(handler.AddTodo))))
	http.Handle("/update", pkg.CorsMiddleware(pkg.LoggingMiddleware(http.HandlerFunc(handler.UpdateTodo))))
	http.Handle("/delete", pkg.CorsMiddleware(pkg.LoggingMiddleware(http.HandlerFunc(handler.DeleteTodo))))

	log.Println("Starting the server...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("can not start the server: %+v", err)
	}
}
