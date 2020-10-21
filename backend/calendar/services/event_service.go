package services

import (
	"fmt"
	"golang/calendar/entities"
)

type EventRepository interface {
	CreateEvent(string, int, string, string)
	GetEventsByUID(string) (entities.Events, int, error)
	DeleteEvent(string, int)
	GetNextEventID(string) int
	EditEvent(string, int, string, string, string, string)
}
type EventService struct {
	EventRepository EventRepository
}

func (e *EventService) CreateEvent(uid string, eventID int, date string, event string) {
	/* Event作成時にNextEventIDを更新する必要あり
	Event作成時には必ず必要な動作なのでe.EventRepository.CreateEventに
	入れ込む(トランザクション処理も可能になるため) */
	e.EventRepository.CreateEvent(uid, eventID, date, event)

}
func (e *EventService) GetEventsByUID(uid string) (entities.Events, int) {
	events, nextEventID, err := e.EventRepository.GetEventsByUID(uid)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("events, nextEventID : ")
	fmt.Println(events, nextEventID)
	return events, nextEventID

}
func (e *EventService) DeleteEvent(UID string, eventID int) {
	e.EventRepository.DeleteEvent(UID, eventID)
}
func (e *EventService) GetNextEventID(UID string) int {
	NextEventID := e.EventRepository.GetNextEventID(UID)
	return NextEventID
}
func (e *EventService) EditEvent(
	UID string,
	EventID int,
	InputEvent string,
	BackgroundColor string,
	BorderColor string,
	TextColor string) {
	e.EventRepository.EditEvent(UID, EventID, InputEvent, BackgroundColor, BorderColor, TextColor)
}
