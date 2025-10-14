package dtos

type CreateCustomerDTO struct {
	Name         string  `json:"name" binding:"required"`
	Email        string  `json:"email" binding:"required,email"`
	Phone        string  `json:"phone" binding:"required,len=13"`
	Address      string  `json:"address" binding:"required"`
	Complement   *string `json:"complement"`
	Number       string  `json:"number" binding:"required"`
	Neighborhood string  `json:"neighborhood" binding:"required"`
	City         string  `json:"city" binding:"required"`
	State        string  `json:"state" binding:"required,len=2"`
	ZIPCode      string  `json:"zipCode" binding:"required,len=8"`
	PaymentDay   int     `json:"paymentDay" binding:"required,gte=1,lte=31" example:"10"`
}

type EditCustomerDTO struct {
	Name         string  `json:"name" binding:"required"`
	Email        string  `json:"email" binding:"required,email"`
	Phone        string  `json:"phone" binding:"required,len=13"`
	Address      string  `json:"address" binding:"required"`
	Complement   *string `json:"complement"`
	Number       string  `json:"number" binding:"required"`
	Neighborhood string  `json:"neighborhood" binding:"required"`
	City         string  `json:"city" binding:"required"`
	State        string  `json:"state" binding:"required,len=2"`
	ZIPCode      string  `json:"zipCode" binding:"required,len=8"`
	PaymentDay   int     `json:"paymentDay" binding:"required,gte=1,lte=31" example:"10"`
}
