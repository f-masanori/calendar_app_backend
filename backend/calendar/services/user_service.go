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
	// Find(int) (entities.User, error)
	// Save(*entities.User) (entities.User, error)
	// Update(int, *entities.User) (entities.User, error)
}
type UserService struct {
	UserRepository UserRepository
}

func (s *UserService) StoreNewUser(UID string, Email string) (entities.User, error) {
	fmt.Println("StoreNewUser")
	user, err := s.UserRepository.CreateUser(UID, Email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created New user ID=" + strconv.Itoa(user.ID) + " name=" + user.Name)
	}
	return user, err
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
