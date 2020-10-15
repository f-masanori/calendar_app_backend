package entities

type Event struct {
	ID              int
	UID             string
	EventID         int
	Date            string
	Event           string
	BackgroundColor string
	BorderColor     string
	TextColor       string
}

type Events []Event
