package services

import (
	"fmt"
	"golang/calendar/entities"
	"golang/calendar/infrastructure/database"
	sqlcmd "golang/calendar/interfaces/database"
	"strconv"
)

type UserRepository interface {
	FindAll() (entities.Users, error)
	CreateUser(string, string) (entities.User, error)
	DeleteUser(int) (int, error)
	CreateNextEventID(string) (int, error)
}
type UserService struct {
	UserRepository UserRepository
}

func (s *UserService) StoreNewUser(UID string, Email string) (entities.User, error) {
	fmt.Println("StoreNewUser")
	fmt.Println(UID, Email)
	user, err := s.UserRepository.CreateUser(UID, Email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created New user ID=" + strconv.Itoa(user.ID) + " name=" + user.Name)
	}
	_, err2 := s.UserRepository.CreateNextEventID(UID)
	if err2 != nil {
		fmt.Println(err2)
	}
	return user, nil
}

/****/

/* for test */
func NewUserService(sqlHandler *database.SqlHandler) *UserService {
	return &UserService{
		UserRepository: &sqlcmd.UserRepository{
			SqlHandler: sqlHandler,
		},
	}
}

/* ******** */

// Index
func (s *UserService) GetAll() (entities.Users, error) {
	users, err := s.UserRepository.FindAll()
	return users, err
}

func (s *UserService) DeleteUser(ID int) int {
	fmt.Println("DeleteUser")
	returnId, err := s.UserRepository.DeleteUser(ID)
	fmt.Println(err)
	return returnId
}
