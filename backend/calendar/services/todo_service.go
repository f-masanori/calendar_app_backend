package services

import (
	"fmt"
	"golang/calendar/entities"
)

type TodoRepository interface {
	CreateTodo(string, int, string)
	DeleteTodo(string, int)
	GetTodosByUID(string) (entities.Todos, int, error)
}
type TodoService struct {
	TodoRepository TodoRepository
}

func (s *TodoService) AddTodo(uid string, todoID int, todo string) {
	/* AllInOne作成時にNextAllInOneIDを更新する必要あり
	AllInOne作成時には必ず必要な動作なのでe.AllInOneRepository.CreateAllInOneに
	入れ込む(トランザクション処理も可能になるため) */
	s.TodoRepository.CreateTodo(uid, todoID, todo)
}
func (s *TodoService) GetTodosByUID(uid string) (entities.Todos, int) {
	/* AllInOne作成時にNextAllInOneIDを更新する必要あり
	AllInOne作成時には必ず必要な動作なのでe.AllInOneRepository.CreateAllInOneに
	入れ込む(トランザクション処理も可能になるため) */
	todos, nextTodoID, err := s.TodoRepository.GetTodosByUID(uid)
	if err != nil {
		fmt.Println(err)
	}
	return todos, nextTodoID
}
func (s *TodoService) DeleteTodo(uid string, todoID int) {
	/* AllInOne作成時にNextAllInOneIDを更新する必要あり
	AllInOne作成時には必ず必要な動作なのでe.AllInOneRepository.CreateAllInOneに
	入れ込む(トランザクション処理も可能になるため) */
	s.TodoRepository.DeleteTodo(uid, todoID)

}
