package usercrudhandler

import "OnlineFlorist/backend/microservices/customer/domain"

type UserCreateReqDTO struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// Age       int    `json:"age"`
}

type UserUpdateReqDTO struct {
	CustID    domain.ID `json:"_custid"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	// Age       int    `json:"age"`
	CreatedOn uint64 `json:"createdOn"`
}

type UserDeteleReqDTO struct {
	CustID    domain.ID `json:"_custid"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	// Age       int    `json:"age"`
	CreatedOn uint64 `json:"createdOn"`
}
