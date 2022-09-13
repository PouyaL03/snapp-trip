package model

const (
	FIXED      string = "FIXED"
	PERCENTAGE string = "PERCENTAGE"
)

type Route struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type Trip struct {
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Airline      string `json:"airline"`
	Agency       string `json:"agency"`
	Supplier     string `json:"supplier"`
	Baseprice    int    `json:"basePrice"`
	Markup       int    `json:"markup"`
	PayablePrice int    `json:"payablePrice"`
}

type TripResponse struct {
	RuleID       int    `json:"ruleId"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Airline      string `json:"airline"`
	Agency       string `json:"agency"`
	Supplier     string `json:"supplier"`
	BasePrice    int    `json:"basePrice"`
	Markup       int    `json:"markup"`
	PayablePrice int    `json:"payablePrice"`
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
