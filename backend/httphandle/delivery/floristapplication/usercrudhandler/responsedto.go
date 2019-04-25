package usercrudhandler

import "OnlineFlorist/backend/microservices/customer/domain"

type UserGetRespDTO struct {
	CustID    domain.ID `json:"_custid"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	// Age       int    `json:"age"`
	CreatedOn uint64 `json:"createdOn"`
}

type UserGetListRespDTO struct {
	Users []UserGetRespDTO `json:"users"`
	Count int              `json:"count"`
}

type UserCreateRespDTO struct {
	CustID domain.ID `json:"_custid"`
}

type UserUpdateRespDTO struct {
	CustID domain.ID `json:"_custid"`
}
