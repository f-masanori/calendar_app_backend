package services

type AllInOneRepository interface {
	CreateAllInOne(string, int, string, string)
	// GetEventsByUID(string) (entities.Events, int, error)
	// DeleteEvent(string, int)
	// GetNextEventID(string) int
	// EditEvent(string, int, string, string, string, string)
}
type AllInOneService struct {
	AllInOneRepository AllInOneRepository
}

func (e *AllInOneService) CreateAllInOne(uid string, AllInOneID int, date string, AllInOne string) {
	/* AllInOne作成時にNextAllInOneIDを更新する必要あり
	AllInOne作成時には必ず必要な動作なのでe.AllInOneRepository.CreateAllInOneに
	入れ込む(トランザクション処理も可能になるため) */
	e.AllInOneRepository.CreateAllInOne(uid, AllInOneID, date, AllInOne)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(nikki)
	// return nikki, err
}

// func (e *EventService) GetEventsByUID(uid string) (entities.Events, int) {
// 	events, nextEventID, err := e.EventRepository.GetEventsByUID(uid)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Print("events, nextEventID : ")
// 	fmt.Println(events, nextEventID)
// 	return events, nextEventID
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// fmt.Println(nikki)
// 	// return nikki, err
// }
// func (e *EventService) DeleteEvent(UID string, eventID int) {
// 	e.EventRepository.DeleteEvent(UID, eventID)
// }
// func (e *EventService) GetNextEventID(UID string) int {
// 	NextEventID := e.EventRepository.GetNextEventID(UID)
// 	return NextEventID
// }
// func (e *EventService) EditEvent(
// 	UID string,
// 	EventID int,
// 	InputEvent string,
// 	BackgroundColor string,
// 	BorderColor string,
// 	TextColor string) {
// 	e.EventRepository.EditEvent(UID, EventID, InputEvent, BackgroundColor, BorderColor, TextColor)
// }
