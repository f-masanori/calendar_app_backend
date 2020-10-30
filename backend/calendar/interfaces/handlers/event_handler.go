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

type EventHandler struct {
	Service *services.EventService
}

func NewEventHandler(sqlHandler *database.SqlHandler) *EventHandler {
	return &EventHandler{
		Service: &services.EventService{
			EventRepository: &sqlcmd.EventRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (e *EventHandler) AddEvent(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	log.Println(" (e *EventHandler) AddEvent")
	type Request struct {
		EventID    int    `json:"EventID"`
		Date       string `json:"Date"`
		InputEvent string `json:"InputEvent"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
		return 400, nil, err

	}
	log.Println(request)
	fmt.Println(r.Body)
	// fmt.Println(r.Method)
	// uid := Authentication.FirebaseUID
	e.Service.CreateEvent(Authentication.FirebaseUID, request.EventID, request.Date, request.InputEvent)
	return 200, nil, nil
}
func (e *EventHandler) GetEventsByUID(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	type Response struct {
		Events      entities.Events `json:"Events"`
		NextEventID int             `json:"NextEventID"`
	}
	Events, NextEventID := e.Service.GetEventsByUID(Authentication.FirebaseUID)
	fmt.Println(Events)
	_Response := Response{Events: Events, NextEventID: NextEventID}

	return 200, _Response, nil
}
func (e *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	log.Println(" (e *EventHandler) DeleteEvent")
	type Request struct {
		EventID int `json:"EventID,string"`
	}
	decoder := json.NewDecoder(r.Body)
	// fmt.Println(decoder)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	log.Println(request)
	// fmt.Println(r.Body)
	// fmt.Println(r.Method)
	// uid := Authentication.FirebaseUID
	e.Service.DeleteEvent(Authentication.FirebaseUID, request.EventID)
	return 200, nil, nil
}
func (e *EventHandler) GetNextEventID(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	log.Println(" (e *EventHandler) GetNextEventID")
	NextEventID := e.Service.GetNextEventID(Authentication.FirebaseUID)
	type Response struct {
		NextEventID int `json:"NextEventID"`
	}
	fmt.Println(NextEventID)
	_Response := Response{NextEventID: NextEventID}
	return 200, _Response, nil
}
func (e *EventHandler) EditEvent(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		EventID int `json:"EventID,string"`
		// Date       string `json:"Date"`
		InputEvent      string `json:"InputEvent"`
		BackgroundColor string `json:"BackgroundColor"`
		BorderColor     string `json:"BorderColor"`
		TextColor       string `json:"TextColor"`
	}
	decoder := json.NewDecoder(r.Body)
	// fmt.Println(decoder)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	log.Println(request)

	e.Service.EditEvent(
		Authentication.FirebaseUID,
		request.EventID,
		request.InputEvent,
		request.BackgroundColor,
		request.BorderColor,
		request.TextColor)

}
