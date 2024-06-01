package handlers

import (
	"fmt"
	
	"tools/modles"
	"tools/pkg/utils"
	"tools/storage/cache"

	"github.com/google/uuid"
)

func ( h *Handlers)Registration(){
	var newUser modles.User

	uid:=uuid.New()
	newUser.UserId=uid.String()
	fmt.Print("Enter Firstname: ")
	fmt.Scanln(&newUser.FirstName)
	
	fmt.Print("Enter lastname: ")
	fmt.Scanln(&newUser.Lastname)
	
	fmt.Print("Enter age: ")
	fmt.Scanln(&newUser.Age)
	
	fmt.Print(" Enter your password: ")
	fmt.Scanln(&newUser.Password)

	fmt.Print("Enter phone number: ")
	fmt.Scanln(&newUser.PhoneNumber)
	var otpNum,phoneNumber int
	otp:=utils.GenerateOtp()
	cache.NewCache().Set(otp,newUser)
	fmt.Println("This is your otp code: ",otp)

	fmt.Print(`	
	WE sent otp code write in here 
	code: `,)
	fmt.Scanln(&otpNum)
	fmt.Print(" Again Enter phone number: ")
	fmt.Scanln(&phoneNumber)
	
	user,err:=cache.NewCache().Get(otpNum,phoneNumber)

	if err!=nil{
		fmt.Println("Error in caches' func: ",err)
		return
	}

	h.UserRepo.CreateUser(*user)
	fmt.Println("congratulations u are registred!!")

	h.UserRepo.CreateUser(newUser)
}


func (h *Handlers)Login(){

	password:=""
	phoneNumber:=0
	fmt.Print("Enter your phone number: ")
	fmt.Scanln(&phoneNumber)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)
	user,err:=h.UserRepo.Login(phoneNumber,password)

	if err!=nil{
		fmt.Println("err on login: ",err)
		return
	}

	UserToken=&Token{}
	UserToken.UserId=user.UserId
	UserToken.FirstName=user.FirstName
	UserToken.Subcriber=user.Subcriber
	fmt.Println("you are login")
	

}

func (h *Handlers)GetUsers(){

		users:=h.UserRepo.GetUser()
	for _,v:=range users{

		fmt.Println("User Id: ",v.UserId)
		fmt.Println("FirstName: ",v.FirstName)
		fmt.Println("LastName: ",v.Lastname)
		fmt.Println("Age: ",v.Age)
		fmt.Println("PhoneNember: ",v.PhoneNumber)
		fmt.Println("")
	}
	return 
}
