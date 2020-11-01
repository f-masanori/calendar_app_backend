package database

import (
	"errors"
	"fmt"
	"golang/calendar/entities"
	"golang/calendar/infrastructure/database"
	"log"
)

type EventRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *EventRepository) CreateEvent(UID string, eventID int, date string, event string) error {
	statement := "INSERT INTO events(uid,event_id,date,event,background_color,border_color,text_color) VALUES(?,?,?,?,?,?,?)"
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmtInsert.Close()
	_, err2 := stmtInsert.Exec(UID, eventID, date, event, "skyblue", "skyblue", "black")
	if err2 != nil {
		log.Println(err2)
		return err2
	}

	fmt.Println("NextEevntID?Update process")
	_, err3 := repo.SqlHandler.DB.Exec("UPDATE next_event_ids SET next_event_id = next_event_id+1 WHERE uid = ?", UID)
	if err3 != nil {
		log.Print(err3)
		return err3
	}
	return nil

}

func (repo *EventRepository) GetEventsByUID(UID string) (entities.Events, int, error) {
	var events entities.Events
	rows, err := repo.SqlHandler.DB.Query("SELECT * from events WHERE uid = ?;", UID)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	defer func() {
		rows.Close()
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
	var _NextEventID int
	if err := repo.SqlHandler.DB.QueryRow("SELECT next_event_id FROM next_event_ids WHERE uid = ?", UID).Scan(&_NextEventID); err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return events, _NextEventID, nil
}

func (repo *EventRepository) DeleteEvent(UID string, eventID int) error {
	stmtDelete, PrepareErr := repo.SqlHandler.DB.Prepare("DELETE FROM events WHERE uid = ? and event_id = ?")
	if PrepareErr != nil {
		log.Panicln(PrepareErr)
		return PrepareErr
	}
	defer stmtDelete.Close()

	result, ExecErr := stmtDelete.Exec(UID, eventID)
	if ExecErr != nil {
		log.Panicln(ExecErr)
		return ExecErr
	}
	_rowsAffect, RowAffwctedeErr := result.RowsAffected()
	if RowAffwctedeErr != nil {
		log.Panicln(RowAffwctedeErr)
		return RowAffwctedeErr
	}
	rowsAffect := int(_rowsAffect)
	if rowsAffect == 0 {
		log.Println("not deleted")
	} else if rowsAffect == 1 {
		log.Println("complete delete")
	} else {
		fmt.Println("DB table error")
		return errors.New("anomaly detection. DB error")
	}
	return nil
}
func (repo *EventRepository) GetNextEventID(UID string) (int, error) {
	var _NextEventID int
	if QueryRowErr := repo.SqlHandler.DB.QueryRow("SELECT next_event_id FROM next_event_ids WHERE uid = ?", UID).Scan(&_NextEventID); QueryRowErr != nil {
		log.Print(QueryRowErr)
		return 0, QueryRowErr
	}
	return _NextEventID, nil
}
func (repo *EventRepository) EditEvent(
	UID string,
	EventID int,
	InputEvent string,
	BackgroundColor string,
	BorderColor string,
	TextColor string) error {
	statement := "UPDATE events set event = ?,background_color = ?,border_color = ?,text_color = ? where uid = ? and event_id = ? "
	stmtInsert, PrepareErr := repo.SqlHandler.DB.Prepare(statement)
	if PrepareErr != nil {
		log.Print(PrepareErr)
		return PrepareErr
	}
	defer stmtInsert.Close()

	_, ExecErr := stmtInsert.Exec(InputEvent, BackgroundColor, BorderColor, TextColor, UID, EventID)
	if ExecErr != nil {
		log.Print(ExecErr)
		return ExecErr
	}
	return nil
}
