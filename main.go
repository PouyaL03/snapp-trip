package main

import (
	"fmt"
	model "snapp-trip/model"
)

func main() {
	fmt.Println("just using fmt so there wouldn't be any problem")
	trip := model.Trip{
		Agency: "agency",
	}

	fmt.Printf("agency is: %s", trip.Agency)
}