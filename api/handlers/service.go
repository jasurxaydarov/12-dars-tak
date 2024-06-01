package handlers

import "fmt"

func (h *Handlers)BinToDec(){
	var binNum int

	if UserToken!=nil{

		if UserToken.Subcriber.Converter{

			fmt.Println("enter bin number: ")
			fmt.Scanln(&binNum)
			dec:=h.Service.Converter.BinToDec(binNum)
			fmt.Println("result: ",dec)

		}else {
			fmt.Println("you are not subcribed")
			cmd:=0
			fmt.Println(`
			1.subcribe
			2.unsubcribe`)
			fmt.Scanln(&cmd)
			switch cmd{
			case 1: UserToken.Subcriber.Converter=true
					fmt.Println("you are subcribed")
			case 2:UserToken.Subcriber.Converter=false
				fmt.Println("you are unsubcribed")
			}
		}

	}else {
		fmt.Println("you are not registred!! ")
	}


}