package handlers

import (
	"tools/db"
	"tools/service"
)
type Handlers struct{
	UserRepo db.UserRepo
	Service service.Service

}
func NewHandlers(UserRepo db.UserRepo)Handlers{
	return Handlers{UserRepo: UserRepo}
}

