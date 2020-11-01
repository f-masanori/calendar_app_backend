package services

import (
	"golang/calendar/entities"
	"log"
)

type EventRepository interface {
	CreateEvent(string, int, string, string) error
	GetEventsByUID(string) (entities.Events, int, error)
	DeleteEvent(string, int) error
	GetNextEventID(string) (int, error)
	EditEvent(string, int, string, string, string, string) error
}

/*
EventServiceのEventRepositoryはEventRepositoryインターフェイスを満たすメソッド群を持っておけばなんでも良い
→DBの種類には依存していない
*/
type EventService struct {
	EventRepository EventRepository
}

func (e *EventService) CreateEvent(uid string, eventID int, date string, event string) error {
	/* Event作成時にNextEventIDを更新する必要あり
	Event作成時には必ず必要な動作なのでe.EventRepository.CreateEventに
	入れ込む(トランザクション処理も可能になるため) */
	err := e.EventRepository.CreateEvent(uid, eventID, date, event)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (e *EventService) GetEventsByUID(uid string) (entities.Events, int, error) {
	events, nextEventID, err := e.EventRepository.GetEventsByUID(uid)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	return events, nextEventID, nil

}
func (e *EventService) DeleteEvent(UID string, eventID int) error {
	err := e.EventRepository.DeleteEvent(UID, eventID)
	if err != nil {
		return err
	}
	return nil
}
func (e *EventService) GetNextEventID(UID string) (int, error) {
	NextEventID, err := e.EventRepository.GetNextEventID(UID)
	if err != nil {
		return -1, err
	}
	return NextEventID, nil
}
func (e *EventService) EditEvent(
	UID string,
	EventID int,
	InputEvent string,
	BackgroundColor string,
	BorderColor string,
	TextColor string) error {
	err := e.EventRepository.EditEvent(UID, EventID, InputEvent, BackgroundColor, BorderColor, TextColor)
	if err != nil {
		return err
	}
	return nil
}
