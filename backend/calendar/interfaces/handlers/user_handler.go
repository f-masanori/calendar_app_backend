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
func (h *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
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

	user, err := h.Service.StoreNewUser(request.UID, request.Email)
	if err != nil {
		fmt.Println(err)
		return 500, user, nil
	}

	return 200, user, nil
}
