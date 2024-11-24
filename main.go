package main

import "fmt"

type MyPersonalStruct struct {
	name    string
	surname string
	age     int16
	score   int16
}

type FunctionalInterface interface {
	calculateAgeScore() int16
	calculateOtherThings() int64
}

// Implement the interface method for MyPersonalStruct
func (ps *MyPersonalStruct) calculateAgeScore() int16 {
	return 12
}

// Factory function to create a new MyPersonalStruct
func NewPersonalStruct() *MyPersonalStruct {
	return &MyPersonalStruct{
		name:    "Dachi",
		surname: "Imedadze",
		age:     23,
		score:   89,
	}
}

func main() {
	fara := NewPersonalStruct()
	fara.calculateAgeScore()
	structPersonale := NewPersonalStruct()
	fmt.Println(structPersonale.calculateAgeScore())
	fmt.Println("krebsona dabrundashvili")
}
