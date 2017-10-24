package main

import (
	"log"
	"services/siv"
)

func main(){

	siv.NewCarInfo()
	siv.NewCarInfo()

	log.Println(siv.NumVehicles())
}