package service

import "tools/service/converter"

type Service struct{
	Converter converter.Converter
}
func NewService()Service{
	return Service{Converter: converter.NewConverter()}
	
}