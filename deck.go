package main

import (
	"fmt"
)

//create card deck!
//its a slice of strings
type deck []string

func (d deck) print(){
	for i, card := range d{
		fmt.Println(i, card)
	}
}