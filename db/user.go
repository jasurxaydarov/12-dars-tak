package db

import (
	"errors"
	"tools/database"
	"tools/modles"
)

type UserRepo struct{

}

func NewUserRepo()UserRepo{
	return UserRepo{}
}

func (u *UserRepo)CreateUser( user modles.User){

	database.Users = append(database.Users, user)
	
}

func (u *UserRepo)GetUser()[]modles.User{

	return database.Users
}

func (u *UserRepo)Login(phoneNumber int,password string)(*modles.User,error){

	for _,v:=range database.Users{

		if phoneNumber==v.PhoneNumber&&password==v.Password{
			return &v,nil
		}
	}
	return nil,errors.New("No account: please create a account ")

}