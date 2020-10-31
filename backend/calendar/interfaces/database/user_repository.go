package database

import (
	"database/sql"
	"fmt"
	"golang/calendar/entities"
	"golang/calendar/infrastructure/database"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *UserRepository) CreateNextEventID(UID string) (int, error) {
	/* NextEventID Create Process */
	statement2 := "INSERT INTO next_event_ids(uid,next_event_id) VALUES(?,?)"
	stmtInsert2, err := repo.SqlHandler.DB.Prepare(statement2)
	if err != nil {
		fmt.Println("NextEventID Create Process error")
		return 0, err
	}
	defer stmtInsert2.Close()
	res2, err := stmtInsert2.Exec(UID, 1)
	if err != nil {
		fmt.Println("error2")
		fmt.Println(res2)
		return 0, err
	}
	fmt.Println("success create next event ID record")
	return 1, nil
}
func (repo *UserRepository) CreateUser(UID string, Email string) (entities.User, error) {
	/* User Create process*/
	var user entities.User
	statement := "INSERT INTO users(uid,email) VALUES(?,?)"
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		fmt.Println("User Create process error")
		return user, err
	}
	defer stmtInsert.Close()
	result, err := stmtInsert.Exec(UID, Email)
	if err != nil {
		fmt.Println("error2")
		return user, err
	}
	lastInsertID, err := result.LastInsertId()
	user.ID = int(lastInsertID)
	user.Name = "name"
	user.UID = "uid"
	/* */

	/* NextTodoID Create Process */
	statement3 := "INSERT INTO next_todo_ids(uid,next_todo_id) VALUES(?,?)"
	stmtInsert3, err := repo.SqlHandler.DB.Prepare(statement3)
	if err != nil {
		fmt.Println("NextTodoID Create Process error")
		return user, err
	}
	defer stmtInsert3.Close()
	res3, err := stmtInsert3.Exec(UID, 1)
	if err != nil {
		fmt.Println("error3")
		fmt.Println(res3)
		return user, err
	}

	return user, nil
}

func (repo *UserRepository) FindAll() (entities.Users, error) {
	var users entities.Users

	fmt.Println("show users")
	rows, err := repo.SqlHandler.DB.Query("SELECT * from users;")
	if err != nil {
		log.Print("error executing database query: ", err)
	}
	defer rows.Close() // make sure rows is closed when the handler exits
	defer fmt.Println("どこで終了かの確認")
	// type users_table struct {
	// 	IS         int    `db:"id"`
	// 	Name       string `db:"name"`
	// 	CreatedAt string `db:"CreatedAt"`
	// 	UpdatedAt string `db:"updated_at"`
	// }
	var users_table_colum Users_table
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&users_table_colum.ID, &users_table_colum.Name, &users_table_colum.CreatedAt, &users_table_colum.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		user.ID = users_table_colum.ID
		user.Name = users_table_colum.Name
		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepository) DeleteUser(id int) (int, error) {
	//トランザクションを使用する
	tx, err := repo.SqlHandler.DB.Begin() // トランザクションを開始
	if err != nil {
		return 0, err
	}
	// Transactionのための関数
	trans := func(tx *sql.Tx) (int64, error) {
		stmt1, _ := tx.Prepare("DELETE FROM users WHERE id = ?")
		result, err := stmt1.Exec(id)
		if err != nil {
			return -1, err
		}
		rowsAffect_int64, err := result.RowsAffected()

		stmt2, _ := tx.Prepare("DELETE FROM user WHERE user_id = ?")
		_, err = stmt2.Exec(id)
		if err != nil {
			return -1, err
		}
		return rowsAffect_int64, nil
	}

	rowsAffect_int64, err := trans(tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	rowsAffect := int(rowsAffect_int64)
	if rowsAffect == 0 {
		fmt.Println("UserID= " + strconv.Itoa(id) + " は存在しません")
	} else if rowsAffect == 1 {
		fmt.Println("UserID = " + strconv.Itoa(id) + " 削除")
	} else {
		fmt.Println("DBエラー")
		fmt.Println("_rowsAffect" + strconv.Itoa(rowsAffect))
	}
	// Commit
	tx.Commit()
	return id, nil
}

func NewUserRepository(sqlHandler *database.SqlHandler) *UserRepository {
	return &UserRepository{
		SqlHandler: sqlHandler,
	}
}
