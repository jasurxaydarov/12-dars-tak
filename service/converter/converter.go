package converter

type Converter struct{}

func NewConverter()Converter{
	return Converter{}
}
func (c Converter)BinToDec(bin int)(dec int){

	dec=0
	base:=1
		for bin>0{
			rem:=bin%10
			dec +=rem
			base*=2
			bin/=10
		}
	return
}