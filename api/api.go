package api

import (
	"fmt"
	"tools/api/handlers"
	"tools/db"
)

func Api(UserRepo db.UserRepo){
	h:=handlers.NewHandlers(UserRepo)
	cmd:=0
	for{	fmt.Println(`
	1.Registration
	2.Login
	3.Logout
	4.Get users
	5.Convert(bin to dec )
	6.Create Todo
	7.Get todos
	-1.Exit

	`)
	fmt.Scanln(&cmd)
	switch cmd{
	case 1: h.Registration()
	case 2: h.Login()
	case 4: h.GetUsers()
	case 5: h.BinToDec()
	case 6:h.CreateTodo()
	case 7:h.GetTodosByUserId()
	case -1:fmt.Println("Thanks, for using") 
		return
	}}
	
}