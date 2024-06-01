package main

import (
	"tools/api"
	"tools/db"
)

func main(){

	userRepo:=db.NewUserRepo()
api.Api(userRepo)
	
}

