package router

import (
	"fmt"

	Authentication "golang/calendar/infrastructure"
	"golang/calendar/infrastructure/database"
	"golang/calendar/infrastructure/middleware"

	"golang/calendar/interfaces/handlers"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run(datasource string, serviceAccountKeyPath string, port int) {
	router := mux.NewRouter()
	DBhandler := database.NewSqlHandler(datasource)
	router.Use(middleware.CORS)
	userHandler := handlers.NewUserHandler(DBhandler)
	eventHandler := handlers.NewEventHandler(DBhandler)
	allInOneHandler := handlers.NewAllInOneHandler(DBhandler)
	todoHandler := handlers.NewTodoHandler(DBhandler)

	router.HandleFunc("/addEvent", Authentication.AuthMiddleware(eventHandler.AddEvent))
	router.HandleFunc("/getEventsByUID", Authentication.AuthMiddleware(eventHandler.GetEventsByUID))
	router.HandleFunc("/registerUser", userHandler.NewUser)
	router.HandleFunc("/deleteEvent", Authentication.AuthMiddleware(eventHandler.DeleteEvent))
	router.HandleFunc("/editEvent", Authentication.AuthMiddleware(eventHandler.EditEvent))
	router.HandleFunc("/getNextEventID", Authentication.AuthMiddleware(eventHandler.GetNextEventID))

	router.HandleFunc("/addScript", Authentication.AuthMiddleware(allInOneHandler.AddScript))

	router.HandleFunc("/addTodo", Authentication.AuthMiddleware(todoHandler.AddTodo))
	router.HandleFunc("/deleteTodo", Authentication.AuthMiddleware(todoHandler.DeleteTodo))

	router.HandleFunc("/getTodosByUID", Authentication.AuthMiddleware(todoHandler.GetTodosByUID))

	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
