package cache

import (
	"errors"
	"tools/database"
	"tools/modles"
)

var Caches map[int]modles.User=map[int]modles.User{}
type Cache struct{}

func NewCache()Cache{

	return Cache{}
}
func (c Cache)Set(otp int,user modles.User)error{

	for _,v:=range database.Users{

		if v.PhoneNumber==user.PhoneNumber{
			return errors.New("This phone number is aleardy used")
		}
	}
	if _,ok:=Caches[otp];ok{
		
		return errors.New("This phone number is aleardy used")
	}


	Caches[otp]=user
	return nil
}
func (c Cache)Get(otp,phoneNumber int)(*modles.User,error){
	
	if _,ok:=Caches[otp];!ok{
		return nil,errors.New("this otp code is wrong")
		
	}
	user:=Caches[otp]
	if user.PhoneNumber!=phoneNumber{
		return nil,errors.New("this is wrong otp")
	}

	delete(Caches,otp)
	return &user,nil
}