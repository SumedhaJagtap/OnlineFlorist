package domain

//customer is a domain object
type FloristShop struct {
	ShopID       ShopID  `json:"_shopid" bson:"_shopid"`
	Name         string  `json:"name" bson:"name"`
	Address      string  `json:"address" bson:"address"`
	AddressLine2 string  `json:"addressLine2" bson:"addressLine2"`
	Email        string  `json:"email" bson:"email"`
	Postcode     string  `json:"postcode" bson:"postcode"`
	Phone        string  `json:"phone" bson:"phone"`
	Rating       float32 `json:"rating" bson:"rating"`
}
