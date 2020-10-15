package router

import (
	"fmt"

	Auth "golang/calendar/infrastructure"
	"golang/calendar/infrastructure/database"
	"golang/calendar/infrastructure/httputil"
	"golang/calendar/infrastructure/middleware"
	"golang/calendar/interfaces/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rs/cors"
)

type AppHandler struct {
	h func(http.ResponseWriter, *http.Request) (int, interface{}, error)
}

func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	status, res, err := a.h(w, r)
	if err != nil {
		httputil.RespondErrorJson(w, status, err)
		return
	}
	httputil.RespondJSON(w, status, res)
	return
}
func Run(datasource string, serviceAccountKeyPath string, port int) {
	router := mux.NewRouter()
	DBhandler := database.NewSqlHandler(datasource)
	router.Use(middleware.CORS)
	auth := Auth.NewFirebaseAuth(serviceAccountKeyPath)
	authMiddleware := Auth.NewFirebaseAuth(serviceAccountKeyPath)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})
	authChain := alice.New(
		corsMiddleware.Handler,
		authMiddleware.FBAuth,
	)
	userHandler := handlers.NewUserHandler(DBhandler)
	eventHandler := handlers.NewEventHandler(DBhandler)
	allInOneHandler := handlers.NewAllInOneHandler(DBhandler)
	todoHandler := handlers.NewTodoHandler(DBhandler)
	router.Methods(http.MethodGet).Path("/me").Handler(authChain.Then(AppHandler{eventHandler.DeleteEvent}))
	// router.HandleFunc("/addEvent", auth.FBAuth(eventHandler.AddEvent))
	// router.HandleFunc("/getEventsByUID", auth.FBAuth(eventHandler.GetEventsByUID))
	// router.HandleFunc("/registerUser", userHandler.NewUser)
	// router.HandleFunc("/deleteEvent", auth.FBAuth(eventHandler.DeleteEvent))
	// router.HandleFunc("/editEvent", auth.FBAuth(eventHandler.EditEvent))
	// router.HandleFunc("/getNextEventID", auth.FBAuth(eventHandler.GetNextEventID))

	// router.HandleFunc("/addScript", auth.FBAuth(allInOneHandler.AddScript))

	// router.HandleFunc("/addTodo", auth.FBAuth(todoHandler.AddTodo))
	// router.HandleFunc("/deleteTodo", auth.FBAuth(todoHandler.DeleteTodo))

	// router.HandleFunc("/getTodosByUID", auth.FBAuth(todoHandler.GetTodosByUID))

	fmt.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
