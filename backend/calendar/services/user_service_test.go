package services

import (
	"golang/calendar/infrastructure/database"
	"golang/conf"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

//このGetaAllは開発者向け
func TestGetAllSuccess(t *testing.T) {

	conf.Test()
	DBhandler := database.TestNewSqlHandler()
	UserService := NewUserService(DBhandler)

	users, err := UserService.GetAll()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	expected := "entities.Users"
	autual := reflect.TypeOf(users).String()

	if autual != expected {
		t.Fatalf("failed test %#v", "user_seivece :GetAll - 返り値型エラー")
	}

}

func TestStoreNewUserSuccess(t *testing.T) {

	conf.Test()
	DBhandler := database.TestNewSqlHandler()
	UserService := NewUserService(DBhandler)

	newUser, StoreNewUserErr := UserService.StoreNewUser("testUID", "test@com")
	if StoreNewUserErr != nil {
		t.Fatalf("failed test %#v", StoreNewUserErr)
	}

	expectedUID := "testUID"
	autualUID := newUser.UID
	// userID := newUser.ID
	// returnedID, DeleteUserErr := UserService.DeleteUser(userID)
	// if DeleteUserErr != nil {
	// 	t.Fatalf("failed test %#v", DeleteUserErr)
	// }

	if autualUID != expectedUID {
		t.Fatalf("failed test %#v", "user_seivece :StoreNewUser - UID返り値型エラー")
	}
	// if returnedID != userID {
	// 	t.Fatalf("failed test %#v", "user_seivece :DeleteUser - UID返り値型エラー")
	// }
}
