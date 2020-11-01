package handlers

import (
	"encoding/json"
	"golang/calendar/entities"
	Authentication "golang/calendar/infrastructure"
	"golang/calendar/infrastructure/database"
	sqlcmd "golang/calendar/interfaces/database"
	"golang/calendar/services"
	"golang/calendar/utils"
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
	type Request struct {
		EventID    string `json:"EventID"`
		Date       string `json:"Date"`
		InputEvent string `json:"InputEvent"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	if DecodeErr := decoder.Decode(&request); DecodeErr != nil {
		return 500, nil, DecodeErr
	}
	if ServicesErr := e.Service.CreateEvent(Authentication.FirebaseUID, utils.StrToInt(request.EventID), request.Date, request.InputEvent); ServicesErr != nil {
		return 500, nil, ServicesErr
	}
	return 200, nil, nil
}
func (e *EventHandler) GetEventsByUID(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	type Response struct {
		Events      entities.Events `json:"Events"`
		NextEventID int             `json:"NextEventID"`
	}
	Events, NextEventID, err := e.Service.GetEventsByUID(Authentication.FirebaseUID)
	if err != nil {
		return 500, nil, err
	}
	_Response := Response{Events: Events, NextEventID: NextEventID}

	return 200, _Response, nil
}
func (e *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	type Request struct {
		EventID int `json:"EventID,string"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		log.Println(err)
		return 500, nil, err
	}
	e.Service.DeleteEvent(Authentication.FirebaseUID, request.EventID)
	return 200, nil, nil
}
func (e *EventHandler) GetNextEventID(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	NextEventID, err := e.Service.GetNextEventID(Authentication.FirebaseUID)
	if err != nil {
		log.Println(err)
		return 500, nil, err
	}
	type Response struct {
		NextEventID int `json:"NextEventID"`
	}
	_Response := Response{NextEventID: NextEventID}
	return 200, _Response, nil
}
func (e *EventHandler) EditEvent(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	type Request struct {
		EventID int `json:"EventID,string"`
		// Date       string `json:"Date"`
		InputEvent      string `json:"InputEvent"`
		BackgroundColor string `json:"BackgroundColor"`
		BorderColor     string `json:"BorderColor"`
		TextColor       string `json:"TextColor"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	if err := decoder.Decode(&request); err != nil {
		log.Println(err)
		return 500, nil, err
	}
	if err := e.Service.EditEvent(
		Authentication.FirebaseUID,
		request.EventID,
		request.InputEvent,
		request.BackgroundColor,
		request.BorderColor,
		request.TextColor); err != nil {
		return 500, nil, err
	}
	return 200, nil, nil
}
