package database

import "tools/modles"


var Users []modles.User=[]modles.User{

	{
	UserId:"9e345e01-ae31-468e-86c4-20b469e00e40" ,
	FirstName: "Jasur",
	Lastname: "xaydarov",
	Age: 19,
	PhoneNumber: 99,
	Password: "999",
	Subcriber:modles.Subcriber{ 
		Converter:false,
		Calculator :false,
		Todo :false,
	},

},{
	UserId:"c7aea64a-b66c-4536-baba-6fd82d8c6f81" ,
	FirstName: "doston ",
	Lastname: "ubaydullayev",
	Age: 19,
	PhoneNumber: 88,
	Password: "888",
	Subcriber:modles.Subcriber{ 
		Converter:false,
		Calculator :false,
		Todo :false,
	},
	

},
{
	UserId:"11d5be34-7ce9-4561-9d3b-fecbbe14195d" ,
	FirstName: "nodir",
	Lastname: "xaydarov",
	Age: 19,
	PhoneNumber: 33,
	Password: "333",
	Subcriber:modles.Subcriber{ 
		Converter:false,
		Calculator :false,
		Todo :false,
	},

},

}


