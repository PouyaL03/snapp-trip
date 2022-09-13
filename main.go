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

	fmt.Printf("agency is: %s\n", trip.Agency)
	
	data := model.ReadFromFile("./resources/request-response/change_price_request.json")
	trips := model.GetTripsFromJson(data)
	
	for index, trip := range trips {
		fmt.Println("trip number ", index + 1)
		trip.Info()
	}
}