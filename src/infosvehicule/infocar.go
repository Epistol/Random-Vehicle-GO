package main

import (
	"log"
	"services/siv"
	"fmt"
)

func main(){

	siv.PrintVehicle()
	return

	c10 := siv.NewCarInfo()
	siv.NewCarInfo()

	fmt.Println("C10 : ", c10)


	log.Println(siv.NumVehicles())
}