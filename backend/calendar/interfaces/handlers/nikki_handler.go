package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// "go_docker/calender/entities"
	"golang/calendar/infrastructure/database"
	sqlcmd "golang/calendar/interfaces/database"
	"golang/calendar/services"

	"github.com/gorilla/mux"
)

type NikkiHandler struct {
	Service *services.NikkiService
}

func NewNikkiHandler(sqlHandler *database.SqlHandler) *NikkiHandler {
	return &NikkiHandler{
		Service: &services.NikkiService{
			NikkiRepository: &sqlcmd.NikkiRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (h *NikkiHandler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("nikkihandler index")

	/* handler call service  */
	nikkis, err := h.Service.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	/* ************ */

	/* presenter */
	// users構造体 → json変換
	json_nikkis, err := json.Marshal(nikkis)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json_nikkis)
	/* ********* */
}

func (h *NikkiHandler) GetNikki(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Println("userID : " + vars["userID"])
	fmt.Println("date : " + vars["date"])
	userID, _ := strconv.Atoi(vars["userID"])
	date, _ := strconv.Atoi(vars["date"])
	/* handler call service  */
	nikki, err := h.Service.GetNikki(userID, date)
	/* ************ */
	/* presenter */
	// users構造体 → json変換
	json_nikki, err := json.Marshal(nikki)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json_nikki)
	/* ********* */
}
func (h *NikkiHandler) RegisterPhoto(w http.ResponseWriter, r *http.Request) {
	/* handler マッピング*/
	type Request struct {
		NikkiId int    `json:"NikkiId"`
		UserId  int    `json:"UserId"`
		Date    int    `json:"Date"`
		PhotoID int    `json:"PhotoID"`
		Photo   string `json:"Photo"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	log.Println(request)
	/* ******* */
	/* handler call service  */
	h.Service.RegisterPhoto(request.NikkiId, request.UserId, request.Date, request.PhotoID, request.Photo)
	/* ************ */
}

func (h *NikkiHandler) CreateNikki(w http.ResponseWriter, r *http.Request) {
	/* handler マッピング*/
	type Request struct {
		UserId         int    `json:"UserId"`
		Date           int    `json:"Date"`
		Content        string `json:"Content"`
		Title          string `json:"Title"`
		NumberOfPhotos int    `json:"NumberOfPhotos"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	log.Println(request)
	/* ******* */

	/* handler service呼び出し */
	nikki, err := h.Service.CreateNikki(request.UserId, request.Date, request.Title, request.Content, request.NumberOfPhotos)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("succused call Service.StoreNewUser")
	}
	/* ******* */

	/* Presenter */
	json_nikki, err := json.Marshal(nikki)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_nikki)
	/* ******* */
}
func (h *NikkiHandler) EditNikki(w http.ResponseWriter, r *http.Request) {
	/* handler マッピング*/
	type Request struct {
		UserId  int    `json:"UserId"`
		Date    int    `json:"Date"`
		Content string `json:"Content"`
		Title   string `json:"Title"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	log.Println(request)
	/* ******* */
	h.Service.EditNikki(request.UserId, request.Date, request.Title, request.Content)
}
func (h *NikkiHandler) DeleteNikki(w http.ResponseWriter, r *http.Request) {
	/* handler マッピング*/
	type Request struct {
		UserId int `json:"UserId"`
		Date   int `json:"Date"`
	}
	decoder := json.NewDecoder(r.Body)
	request := new(Request)
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	log.Println(request)
	/* ******* */
	/* service 呼び出し */
	confirmDelete := h.Service.DeleteNikki(request.UserId, request.Date)
	fmt.Println(confirmDelete)
	/* ******* */
	/* Presenter */
	json_confirmDelete, err := json.Marshal(confirmDelete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_confirmDelete)
	/* ******* */
}
func (h *NikkiHandler) AddEvent(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.Body)
	fmt.Println(r.Body)
	fmt.Println(r.Method)
	wweew := r.Header.Get("token")
	fmt.Println(wweew)
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "token")

}

// func (h *NikkiHandler) GetAllPhotos(w http.ResponseWriter, r *http.Request) {
// 	/* handler call service  */
// 	h.Service.GetAllPhotos()
// 	/* ************ */
// }
