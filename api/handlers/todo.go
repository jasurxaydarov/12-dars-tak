package handlers

import (
	"fmt"
	"time"
	"tools/db"
	"tools/modles"
)

var todoRepo=db.NewTodoRepo()
func(h Handlers)CreateTodo(){

	var newTodo modles.Todo
	if UserToken!=nil{
		if UserToken.Subcriber.Todo!=true{
			cmd:=0
			fmt.Println("you are not subcribed ")
			fmt.Println("1.subcribe")
			fmt.Scanln(&cmd)
			if cmd==1{
				UserToken.Subcriber.Todo=true
				fmt.Println("you are subcribed")
				if UserToken.Subcriber.Todo{
					fmt.Println("Enter task:")
					fmt.Scanln(&newTodo.Task)
					newTodo.CreatedAt=time.Now()
		
					newTodo.UserId=UserToken.UserId
					todoRepo.CreateTodo(newTodo)
					
				}
			}else{
				fmt.Println("wrong oprt")
			}


		} 
		
		

	}else{
		fmt.Println("You are not logedin")
		return 
	}
}


func (h Handlers)GetTodosByUserId(){

	if UserToken!=nil {
		if UserToken.Subcriber.Todo!=true{
			cmd:=0
			fmt.Println("you are not subcribed ")
			fmt.Println(`
			1.Subcribe 
		
			`)
			fmt.Scanln(&cmd)
			if cmd==1{
				UserToken.Subcriber.Todo=true
				fmt.Println("you are subcribed")
			}

		
		}else if UserToken.Subcriber.Todo{

			
			todos:=todoRepo.GetTodosByUserId(UserToken.UserId)
			for _,v:=range todos{
				fmt.Println(v)
			}
		
		}
		
	}else{
	fmt.Println("You are not login")
	return 

	}
}