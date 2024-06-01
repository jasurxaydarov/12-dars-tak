package db

import (
	"tools/database"
	"tools/modles"
)

type TodoRepo struct{

}
func NewTodoRepo()TodoRepo{
	return TodoRepo{}
}

func (t TodoRepo)CreateTodo( todo modles.Todo)error{
		database.Todo=append(database.Todo, todo)

	return nil
}

func(t TodoRepo)GetTodosByUserId(userId string)([]modles.Todo){

	var todo[]modles.Todo
	for _,v:=range database.Todo{

		if v.UserId==userId{
			todo=append(todo, v)
		}
	}

	return todo
	
}
func (t TodoRepo)GetAllTodos()([]modles.Todo){

	var todo []modles.Todo
	for _,v:=range database.Todo{
		todo=append(todo, v)
	}
	return todo
}