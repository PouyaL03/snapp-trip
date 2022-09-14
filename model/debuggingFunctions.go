package model

import (
	"fmt"
)

func (trip *Trip) Info() {
	fmt.Println("origin is: ", trip.Origin)
	fmt.Println("destination is: ", trip.Destination)
	fmt.Println("agency is: ", trip.Agency)
	fmt.Println("supplier is: ", trip.Supplier)
	fmt.Println("basePrice is: ", trip.Baseprice)
}

func (rule *Rule) Info() {
}