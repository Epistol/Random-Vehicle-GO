package siv

import (
	"fmt"
	"math/rand"
	"time"
)

const{
	COULEUR_INDEFINIE =
}

type carInfo struct {
	Immat string
	Couleur int
}

var (
	numVehicles int
	mapVehicules = make(map[string]*carInfo)
)

func getRandomPlaque() string{
	s1 := randomLetter()+randomLetter()
	s2 := randomLetter()+randomLetter()
	nb := rand.Intn(1000)
	fmt.Sprintf("", s1, nb, s2)
}

func randomLetter() string{
	rank := rand.Intn(26)
	letter := int('A') + rank
	return string(letter)
}

func init(){
	fmt.Println("siv init")

	rand.Seed(time.Now().UnixNano)

	for i := 0; i < 10; i++{
		plaque:= getRandomPlaque()
		ci := &carInfo{
			Immat:plaque,
		}
		mapVehicules{plaque} = ci
	}
}

func PrintVehicle(){
	for _ ,v := range mapVehicules{
		fmt.Println(v)
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

func GetCarInfo() carInfo{
	return carInfo{}
}