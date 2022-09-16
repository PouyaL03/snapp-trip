package controller

import ()

func TripValidate(origin string, destination string, airline string, agency string, supplier string) bool {
	return isCityInDatabase(origin) && isCityInDatabase(destination) && isAgencyInDatabse(agency) && isSupplierInDatabase(supplier) && isAirlineInDatabase(airline)
}

func RuleValidate() bool {
	return true
}