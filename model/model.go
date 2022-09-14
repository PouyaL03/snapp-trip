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
	Origin      string `json:"origin" gorm:"origin"`
	Destination string `json:"destination" gorm:"origin"`
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

type Airline struct {
	Code string `json:"code" gorm:"code"`
	EnName string `json:"enName" gorm:"enName"`
	FaName string `json:"faName" gorm:"faName"`
}

type Agency struct {
	Name string `json:"name" gorm:"name"`
}

type City struct {
	Code string `json:"code" gorm:"code"`
	FaName string `json:"faName" gorm:"faName"`
}

type Supplier struct {
	EnName string `json:"enName" gorm:"enName"`
	FaName string `json:"faName" gorm:"faName"`
}