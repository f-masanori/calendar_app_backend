package handlers

import (
	"encoding/json"
	"fmt"
	"golang/calendar/entities"
	Authentication "golang/calendar/infrastructure"
	"golang/calendar/infrastructure/database"
	sqlcmd "golang/calendar/interfaces/database"
	"golang/calendar/services"
	"log"
	"net/http"
)

type TodoHandler struct {
	Service *services.TodoService
}

func NewTodoHandler(sqlHandler *database.SqlHandler) *TodoHandler {
	return &TodoHandler{
		Service: &services.TodoService{
			TodoRepository: &sqlcmd.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (h *TodoHandler) AddTodo(w http.ResponseWriter, r *http.Request) {
	log.Println(" (e *TodoHandler) AddTodo")
	type Request struct {
		// EventID    int    `json:"EventID"`
		// Date       string `json:"Date"`
		// InputEvent string `json:"InputEvent"`
		TodoID int    `json:"TodoID,string"`
		Todo   string `json:"Todo"`
	}
	decoder := json.NewDecoder(r.Body)
	// fmt.Println(decoder)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}

	// newTodoID, _ := strconv.Atoi(request.TodoID)
	// fmt.Println("newTodoID")

	fmt.Println(request.TodoID)
	h.Service.AddTodo(Authentication.FirebaseUID, request.TodoID, request.Todo)
}

func (h *TodoHandler) GetTodosByUID(w http.ResponseWriter, r *http.Request) {
	log.Println(" GetTodosByUID")

	/* Presenter */
	type Response struct {
		Todos      entities.Todos `json:"Todos"`
		NextTodoID int            `json:"NextTodoID"`
	}
	Todos, NextTodoID := h.Service.GetTodosByUID(Authentication.FirebaseUID)
	fmt.Println(Todos)
	_Response := Response{Todos: Todos, NextTodoID: NextTodoID}
	jsonTodos, err := json.Marshal(_Response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonTodos)
}
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println(" (e *TodoHandler) DeleteTodo")
	type Request struct {
		TodoID int    `json:"TodoID,string"`
		Todo   string `json:"Todo"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}

	fmt.Println(request.TodoID)
	h.Service.DeleteTodo(Authentication.FirebaseUID, request.TodoID)
}
