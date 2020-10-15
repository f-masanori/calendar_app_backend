package handlers

import (
	"encoding/json"
	"fmt"
	"golang/calendar/infrastructure/database"
	sqlcmd "golang/calendar/interfaces/database"
	"golang/calendar/services"
	"log"
	"net/http"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(sqlHandler *database.SqlHandler) *UserHandler {
	return &UserHandler{
		Service: &services.UserService{
			UserRepository: &sqlcmd.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}
func (h *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	log.Println("(h *UserHandler) NewUser")
	/* handler マッピング*/
	type Request struct {
		UID   string `json:"UID"`
		Email string `json:"Email"`
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
	user, err := h.Service.StoreNewUser(request.UID, request.Email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("succused call Service.StoreNewUser")
	}
	/* ******* */

	/* Presenter */
	json_user, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_user)
	/* ******* */
}

/******************以下 calendar appでは未使用************/
/***************************************************/

// func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
// 	/* handler call service  */
// 	users, error := h.Service.GetAll() //GetAllの返り値はエンティティのusersでいい？
// 	if error != nil {
// 		fmt.Println(error)
// 		return
// 	}
// 	/* ************ */

// 	/* presenter */
// 	// users構造体 → json変換
// 	json_users, err := json.Marshal(users)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(json_users)
// 	/* ********* */
// }

// func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	/* handler マッピング*/
// 	type Request struct {
// 		Id int `json:"Id"`
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	request := new(Request)
// 	err := decoder.Decode(&request)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(request)
// 	/* ******* */
// 	/* handler service呼び出し */
// 	returnId := h.Service.DeleteUser(request.Id)
// 	/* ******* */
// 	/* Presenter */
// 	json_returnId, err := json.Marshal(returnId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(json_returnId)
// 	/* ******* */
// }
func (h *UserHandler) Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler  Test")
	// 抽象的にGetALL
	// users, error := h.Service.GetAll()

	// fmt.Println(users)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }

	// // json変換
	// res, err := json.Marshal(users)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(res)
	// reqres.W.Write([]byte("uuu"))
}
