package domain

//customer is a domain object
type Customer struct {
	CustID    ID     `json:"_custid" bson:"_custid"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CreatedOn uint64 `json:"createdOn"`
}
