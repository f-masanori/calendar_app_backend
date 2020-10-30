package database

import (
	"fmt"
	"golang/calendar/entities"
	"golang/calendar/infrastructure/database"
	"log"
)

type EventRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *EventRepository) CreateEvent(UID string, eventID int, date string, event string) {
	/* Event?Create?? */
	fmt.Println("Event?Create process")
	statement := "INSERT INTO events(uid,event_id,date,event,background_color,border_color,text_color) VALUES(?,?,?,?,?,?,?)"
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		log.Println("Prepare(statement) error")
	}
	defer stmtInsert.Close()
	log.Println(UID, date, event)
	result, err := stmtInsert.Exec(UID, eventID, date, event, "skyblue", "skyblue", "black")
	fmt.Println(result)
	if err != nil {
		log.Println("stmtInsert.Exec error")
	}
	/*******/

	/* NextEevntID????? */
	fmt.Println("NextEevntID?Update process")
	result2, err2 := repo.SqlHandler.DB.Exec("UPDATE next_event_ids SET next_event_id = next_event_id+1 WHERE uid = ?", UID)
	if err2 != nil {
		log.Println("NextEevntID????? repo.SqlHandler.DB.Exec error")
		log.Print(err2)
	}
	fmt.Println(result2)
	/*******/
}

func (repo *EventRepository) GetEventsByUID(UID string) (entities.Events, int, error) {
	/* Event Read ?? */
	var events entities.Events
	// fmt.Println(UID)
	rows, err := repo.SqlHandler.DB.Query("SELECT * from events WHERE uid = ?;", UID)
	if err != nil {
		log.Println("error executing database query: ", err)
	}
	defer func() {
		rows.Close()
		fmt.Println("End func")
		if err := recover(); err != nil {
			fmt.Println("Runtime Error:", err)
		}
	}()
	var events_table_colum Events_table
	for rows.Next() {
		var event entities.Event
		err := rows.Scan(
			&events_table_colum.ID,
			&events_table_colum.UID,
			&events_table_colum.EventID,
			&events_table_colum.Date,
			&events_table_colum.Event,
			&events_table_colum.BackgroundColor,
			&events_table_colum.BorderColor,
			&events_table_colum.TextColor,
			&events_table_colum.CreatedAt,
			&events_table_colum.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, 0, err
		}
		event.ID = events_table_colum.ID
		event.UID = events_table_colum.UID
		event.EventID = events_table_colum.EventID
		event.Date = events_table_colum.Date
		event.Event = events_table_colum.Event
		event.BackgroundColor = events_table_colum.BackgroundColor
		event.BorderColor = events_table_colum.BorderColor
		event.TextColor = events_table_colum.TextColor
		events = append(events, event)
	}
	/**************/

	/* NextEventID Read ?? */
	var _NextEventID int
	if err := repo.SqlHandler.DB.QueryRow("SELECT next_event_id FROM next_event_ids WHERE uid = ?", UID).Scan(&_NextEventID); err != nil {
		log.Println("NextEventID Read ??")
		log.Println(err)
		return nil, 0, err
	}
	fmt.Println(_NextEventID)

	return events, _NextEventID, nil
}

func (repo *EventRepository) DeleteEvent(UID string, eventID int) {
	stmtDelete, err := repo.SqlHandler.DB.Prepare("DELETE FROM events WHERE uid = ? and event_id = ?")
	if err != nil {
		log.Panicln("(repo *EventRepository) DeleteEvent error")
		panic(err.Error())
	}
	defer stmtDelete.Close()

	result, err := stmtDelete.Exec(UID, eventID)
	if err != nil {
		panic(err.Error())
	}
	_rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(_rowsAffect)
	rowsAffect := int(_rowsAffect)
	if rowsAffect == 0 {
		fmt.Println("???????")
	} else if rowsAffect == 1 {
		fmt.Println("complete delete")
	} else {
		fmt.Println("DB table error") //??????2??????????
	}
}
func (repo *EventRepository) GetNextEventID(UID string) int {
	var _NextEventID int
	if err := repo.SqlHandler.DB.QueryRow("SELECT next_event_id FROM next_event_ids WHERE uid = ?", UID).Scan(&_NextEventID); err != nil {
		log.Println("NextEventID Read ??")
		log.Print(err)
		return 0
	}
	fmt.Println(_NextEventID)
	return _NextEventID
}
func (repo *EventRepository) EditEvent(
	UID string,
	EventID int,
	InputEvent string,
	BackgroundColor string,
	BorderColor string,
	TextColor string) {
	fmt.Println("EventEdit process")
	statement := "UPDATE events set event = ?,background_color = ?,border_color = ?,text_color = ? where uid = ? and event_id = ? "
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		log.Print(err)
	}
	defer stmtInsert.Close()

	stmtInsert.Exec(InputEvent, BackgroundColor, BorderColor, TextColor, UID, EventID)
}
