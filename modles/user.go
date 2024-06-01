package modles

type Subcriber struct{
	Converter 	bool
	Calculator 	bool
	Todo 		bool
}
type User struct{
	UserId string
	FirstName 	string
	Lastname  	string
	Age 		int
	PhoneNumber int
	Password 	string
	Subcriber	Subcriber
	
}
