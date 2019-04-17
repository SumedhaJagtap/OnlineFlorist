package usercrudhandler

type UserCreateReqDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

type UserUpdateReqDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	CreatedOn uint64 `json:"createdOn"`
}

type UserDeteleReqDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	CreatedOn uint64 `json:"createdOn"`
}
