package dtos

import "strings"

type CreateCustomerDTO struct {
	Name         string  `json:"name" binding:"required,min=3,max=255"`
	Email        string  `json:"email" binding:"required,email,max=255"`
	Phone        string  `json:"phone" binding:"required,phone"`
	Address      string  `json:"address" binding:"required,min=10,max=255"`
	Complement   *string `json:"complement" binding:"max=100"`
	Number       string  `json:"number" binding:"required,max=10"`
	Neighborhood string  `json:"neighborhood" binding:"required,min=3,max=100"`
	City         string  `json:"city" binding:"required,min=2,max=100"`
	State        string  `json:"state" binding:"required,len=2"`
	ZIPCode      string  `json:"zipCode" binding:"required,len=8"`
	PaymentDay   int     `json:"paymentDay" binding:"required,gte=1,lte=31" example:"10"`
}

func (dto *CreateCustomerDTO) Sanitize() {
	dto.Name = strings.TrimSpace(dto.Name)
	dto.Email = strings.TrimSpace(dto.Email)
	dto.Phone = strings.TrimSpace(dto.Phone)
	dto.Address = strings.TrimSpace(dto.Address)
	dto.Number = strings.TrimSpace(dto.Number)
	dto.Neighborhood = strings.TrimSpace(dto.Neighborhood)
	dto.City = strings.TrimSpace(dto.City)
	dto.State = strings.TrimSpace(dto.State)
	dto.ZIPCode = strings.TrimSpace(dto.ZIPCode)

	if dto.Complement != nil {
		trimmedComplement := strings.TrimSpace(*dto.Complement)
		dto.Complement = &trimmedComplement
	}
}

type EditCustomerDTO struct {
	Name         string  `json:"name" binding:"required,min=3,max=255"`
	Email        string  `json:"email" binding:"required,email,max=255"`
	Phone        string  `json:"phone" binding:"required,phone"`
	Address      string  `json:"address" binding:"required,min=10,max=255"`
	Complement   *string `json:"complement" binding:"max=100"`
	Number       string  `json:"number" binding:"required,max=10"`
	Neighborhood string  `json:"neighborhood" binding:"required,min=3,max=100"`
	City         string  `json:"city" binding:"required,min=2,max=100"`
	State        string  `json:"state" binding:"required,len=2"`
	ZIPCode      string  `json:"zipCode" binding:"required,len=8"`
	PaymentDay   int     `json:"paymentDay" binding:"required,gte=1,lte=31" example:"10"`
}

func (dto *EditCustomerDTO) Sanitize() {
	dto.Name = strings.TrimSpace(dto.Name)
	dto.Email = strings.TrimSpace(dto.Email)
	dto.Phone = strings.TrimSpace(dto.Phone)
	dto.Address = strings.TrimSpace(dto.Address)
	dto.Number = strings.TrimSpace(dto.Number)
	dto.Neighborhood = strings.TrimSpace(dto.Neighborhood)
	dto.City = strings.TrimSpace(dto.City)
	dto.State = strings.TrimSpace(dto.State)
	dto.ZIPCode = strings.TrimSpace(dto.ZIPCode)

	if dto.Complement != nil {
		trimmedComplement := strings.TrimSpace(*dto.Complement)
		dto.Complement = &trimmedComplement
	}
}
