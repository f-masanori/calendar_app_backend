package database

import (
	"fmt"
	"golang/calendar/entities"
	"golang/calendar/infrastructure/database"
	"log"
)

type TodoRepository struct {
	SqlHandler *database.SqlHandler
}

func (repo *TodoRepository) CreateTodo(UID string, todoID int, todo string) {
	/* CreateTodo */
	fmt.Println("CreateTodo")
	statement := "INSERT INTO todos(uid,todo_id,todo) VALUES(?,?,?)"
	stmtInsert, err := repo.SqlHandler.DB.Prepare(statement)
	if err != nil {
		log.Println("Prepare(statement) error")
	}
	defer stmtInsert.Close()
	// log.Println(UID, date, event)
	result, err := stmtInsert.Exec(UID, todoID, todo)
	fmt.Println(result)
	if err != nil {
		log.Println("stmtInsert.Exec error")
	}
	/*******/

	/* NextTodoID */
	result2, err2 := repo.SqlHandler.DB.Exec("UPDATE next_todo_ids SET next_todo_id = next_todo_id+1 WHERE uid = ?", UID)
	if err2 != nil {
		log.Println("NextEevntID????? repo.SqlHandler.DB.Exec error")
		log.Fatal(err2)
	}
	fmt.Println(result2)
	/*******/
}

func (repo *TodoRepository) GetTodosByUID(UID string) (entities.Todos, int, error) {
	var todos entities.Todos
	fmt.Println("GetTodoByUID")
	rows, err := repo.SqlHandler.DB.Query("SELECT * from todos WHERE uid = ?;", UID)
	if err != nil {
		log.Print("error executing database query: ", err)
	}
	defer rows.Close() // make sure rows is closed when the handler exits
	var todos_table_colum Todos_table
	for rows.Next() {
		var todo entities.Todo
		err := rows.Scan(&todos_table_colum.ID, &todos_table_colum.UID, &todos_table_colum.TodoID, &todos_table_colum.Todo, &todos_table_colum.CreatedAt, &todos_table_colum.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		todo.ID = todos_table_colum.TodoID
		todo.Todo = todos_table_colum.Todo
		todos = append(todos, todo)
	}

	/* NextEventID Read ?? */
	var _NextTodoID int
	if err := repo.SqlHandler.DB.QueryRow("SELECT next_todo_id FROM next_todo_ids WHERE uid = ?", UID).Scan(&_NextTodoID); err != nil {
		fmt.Println("NextEventID Read error")
		log.Fatal(err)
	}
	return todos, _NextTodoID, nil
}

func (repo *TodoRepository) DeleteTodo(UID string, todoID int) {
	stmtDelete, err := repo.SqlHandler.DB.Prepare("DELETE FROM todos WHERE uid = ? and todo_id = ?")
	if err != nil {
		log.Panicln("(repo *TodoRepository) DeleteTodo error")
		panic(err.Error())
	}
	defer stmtDelete.Close()

	result, err := stmtDelete.Exec(UID, todoID)
	if err != nil {
		panic(err.Error())
	}
	_rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(_rowsAffect)
	rowsAffect := int(_rowsAffect)
	if rowsAffect == 0 {
		fmt.Println("???????")
	} else if rowsAffect == 1 {
		fmt.Println("complete delete")
	} else {
		fmt.Println("DB table error") //??????2??????????
	}
}
