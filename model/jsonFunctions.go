package model

import (
	"encoding/json"
	"log"
)

func GetTripsFromJson(data []byte) []Trip{
	var trips []Trip

	err := json.Unmarshal(data, &trips)
	if err != nil {
		log.Fatal(err)
	}

	return trips
}

func GetRulesFromJson(data []byte) []Rule{
	var rules []Rule

	err := json.Unmarshal(data, &rules)
	if err != nil {
		log.Fatal(err)
	}

	return rules
}