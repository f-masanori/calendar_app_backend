package services

import (
	"fmt"
	"golang/calendar/entities"
	"golang/calendar/infrastructure/database"
	sqlcmd "golang/calendar/interfaces/database"
)

type UserRepository interface {
	FindAll() (entities.Users, error)
	CreateUser(string, string, string) (entities.User, error)
	DeleteUser(int) (int, error)
	CreateNextEventID(string) (int, error)
}
type UserService struct {
	UserRepository UserRepository
}

func (s *UserService) StoreNewUser(UID string, Email string) (entities.User, error) {
	//　トランザクション処理を実装する
	user, err := s.UserRepository.CreateUser(UID, Email, "user")
	if err != nil {
		fmt.Println(err)
		u := entities.User{}
		// ここはもっとスマートな方法がないのか？
		return u, err
	}
	_, RepErr := s.UserRepository.CreateNextEventID(UID)
	if RepErr != nil {
		fmt.Println(RepErr)
		u := entities.User{}
		return u, RepErr

	}
	return user, nil
}
func (s *UserService) GetAll() (entities.Users, error) {
	users, RepErr := s.UserRepository.FindAll()
	if RepErr != nil {
		fmt.Println(RepErr)
		u := entities.Users{}
		return u, RepErr
	}
	return users, nil
}

func (s *UserService) DeleteUser(ID int) (int, error) {
	returnID, RepErr := s.UserRepository.DeleteUser(ID)
	if RepErr != nil {
		fmt.Println(RepErr)
		return -1, RepErr
	}
	return returnID, nil
}

/* for test */
func NewUserService(sqlHandler *database.SqlHandler) *UserService {
	return &UserService{
		UserRepository: &sqlcmd.UserRepository{
			SqlHandler: sqlHandler,
		},
	}
}
