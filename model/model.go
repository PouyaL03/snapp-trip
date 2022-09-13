package model

type tmp interface {
	Info() 
}

type (
	amountType string
)

const (
	FIXED      amountType = "FIXED"
	PERCENTAGE amountType = "PERCENTAGE"
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

type Rule struct {
	Routes      []Route    `json:"routes"`
	Airlines    []string   `json:"airlines"`
	Agencies    []string   `json:"agencies"`
	Suppliers   []string   `json:"suppliers"`
	AmountType  amountType `json:"amountType"`
	AmountValue int        `json:"amountValue"`
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
