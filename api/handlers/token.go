package handlers

import "tools/modles"

type Token struct{

	FirstName 	string
	UserId 		string
	Subcriber modles.Subcriber

}
var UserToken *Token