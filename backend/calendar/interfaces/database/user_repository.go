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
	statement := "INSERT INTO next_event_ids(uid,next_event_id) VALUES(?,?)"
	stmtInsert, PrepareErr := repo.SqlHandler.DB.Prepare(statement)
	if PrepareErr != nil {
		fmt.Println(PrepareErr)
		return 0, PrepareErr
	}
	defer stmtInsert.Close()
	_, ExecErr := stmtInsert.Exec(UID, 1)
	if ExecErr != nil {
		fmt.Println(ExecErr)
		return 0, ExecErr
	}
	return 1, nil
}
func (repo *UserRepository) CreateUser(UID string, Email string) (entities.User, error) {
	var user entities.User
	statement := "INSERT INTO users(uid,email) VALUES(?,?)"
	stmtInsert, PrepareErr := repo.SqlHandler.DB.Prepare(statement)
	if PrepareErr != nil {
		log.Println(PrepareErr)
		return user, PrepareErr
	}
	defer stmtInsert.Close()
	result, ExecErr := stmtInsert.Exec(UID, Email)
	if ExecErr != nil {
		log.Println(ExecErr)
		return user, ExecErr
	}
	lastInsertID, LastInsertIDErr := result.LastInsertId()
	if LastInsertIDErr != nil {
		log.Println(LastInsertIDErr)
		return user, LastInsertIDErr
	}
	user.ID = int(lastInsertID)
	user.Name = "name"
	user.UID = "uid"

	/* NextTodoID Create Process */
	statement2 := "INSERT INTO next_todo_ids(uid,next_todo_id) VALUES(?,?)"
	stmtInsert2, NextTodoIDPrepareErr := repo.SqlHandler.DB.Prepare(statement2)
	if NextTodoIDPrepareErr != nil {
		log.Println(NextTodoIDPrepareErr)
		return user, NextTodoIDPrepareErr
	}
	defer stmtInsert2.Close()
	_, ExecErr2 := stmtInsert2.Exec(UID, 1)
	if ExecErr2 != nil {
		log.Println(ExecErr2)
		return user, ExecErr2
	}

	return user, nil
}

func (repo *UserRepository) FindAll() (entities.Users, error) {
	var users entities.Users
	rows, QueryErr := repo.SqlHandler.DB.Query("SELECT * from users;")
	if QueryErr != nil {
		log.Println(QueryErr)
	}
	defer rows.Close()
	var users_table_colum Users_table
	for rows.Next() {
		var user entities.User

		if err := rows.Scan(&users_table_colum.ID, &users_table_colum.UID, &users_table_colum.Name, &users_table_colum.Email, &users_table_colum.CreatedAt, &users_table_colum.UpdatedAt); err != nil {
			log.Println(err)
			return nil, err
		}
		user.ID = users_table_colum.ID
		user.UID = users_table_colum.UID
		user.Name = users_table_colum.Name

		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepository) DeleteUser(id int) (int, error) {
	// ここでTransaction処理をするのはおかしい。servicesで行う

	tx, err := repo.SqlHandler.DB.Begin()
	if err != nil {
		return 0, err
	}
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
	tx.Commit()
	return id, nil
}

func NewUserRepository(sqlHandler *database.SqlHandler) *UserRepository {
	return &UserRepository{
		SqlHandler: sqlHandler,
	}
}
