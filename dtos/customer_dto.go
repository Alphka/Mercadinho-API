package dtos

import "strings"

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
