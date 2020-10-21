package database

type Users_table struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
type Events_table struct {
	ID              int    `db:"id"`
	UID             string `db:"uid"`
	Date            string `db:"date"`
	EventID         int    `db:"event_id"`
	Event           string `db:"event"`
	BackgroundColor string `db:"background_color"`
	BorderColor     string `db:"border_color"`
	TextColor       string `db:"text_color"`
	CreatedAt       string `db:"created_at"`
	UpdatedAt       string `db:"updated_at"`
}
type Next_event_id_table struct {
	ID          int    `db:"id"`
	UID         string `db:"uid"`
	NextEventID int    `db:"next_event_id"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

type Todos_table struct {
	ID        int    `db:"id"`
	UID       string `db:"uid"`
	TodoID    int    `db:"todo_id"`
	Todo      string `db:"todo"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
