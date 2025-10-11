package models

type Customer struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	Address      string  `json:"address"`
	Complement   *string `json:"complement"`
	Number       string  `json:"number"`
	Neighborhood string  `json:"neighborhood"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	ZIPCode      string  `json:"zipCode"`
	PaymentDay   int     `json:"paymentDay"`
}
