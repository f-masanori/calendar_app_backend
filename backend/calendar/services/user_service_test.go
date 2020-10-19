package services

import (
	"golang/calendar/conf"
	"testing"

	// "go_docker/mynikki/entities"
	"golang/calendar/infrastructure/database"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

//このGetaAllは開発者向け
func TestGetAllSuccess(t *testing.T) {

	conf.Test()
	DBhandler := database.TestNewSqlHandler()
	NewUserService := NewUserService(DBhandler)

	users, err := NewUserService.GetAll()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	expected := "entities.Users"
	autual := reflect.TypeOf(users).String()

	if autual != expected {
		t.Fatalf("failed test %#v", "user_seivece :GetAll - 返り値型エラー")
	}

}

// func TestStoreNewUserSuccess(t *testing.T) {
// 	conf.Init()
// 	DBhandler := database.NewSqlHandler()
// 	rows, err := DBhandler.DB.Query("SELECT * from users;")
// 	fmt.Println(rows)
// 	if err !=nil{
// 		t.Fatalf("failed test %#v", err)
// 	}
// }
