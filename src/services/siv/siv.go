package siv

import (
	"fmt"
	"math/rand"
)

const

type carInfo struct {
	Immat string
	Couleur int
}

var (
	numVehicles int
	mapVehicules = make(map[string]*carInfo)
)

func getRandomPlaque() string{

}

func randomLetter() string{
	rank := rand.Intn(26)
	letter := int('A') + rank
	return string(letter)
}

func init(){
	fmt.Println("siv init")

	for i := 0; i < 10; i++{
		//ajout vehicule
	}
}

func NewCarInfo() *carInfo{
	ci := &carInfo{}
	ci.Immat = randomLetter()
	numVehicles++
	return &carInfo{}
}

func NumVehicles() int{
	return numVehicles
}

func GetCarInfo(string) CarInfo{
	return CarInfo{}
}