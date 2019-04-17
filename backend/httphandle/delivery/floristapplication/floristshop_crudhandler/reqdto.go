package floristshopcrudhandler

import "OnlineFlorist/backend/microservices/florist_shop/domain"

type ShopCreateReqDTO struct {
	ShopID       domain.ShopID `json:"_shopid" bson:"_shopid"`
	Name         string        `json:"name" bson:"name"`
	Address      string        `json:"address" bson:"address"`
	AddressLine2 string        `json:"addressLine2" bson:"addressLine2"`
	Email        string        `json:"email" bson:"email"`
	Postcode     string        `json:"postcode" bson:"postcode"`
	Phone        string        `json:"phone" bson:"phone"`
	Rating       float32       `json:"rating" bson:"rating"`
}

type ShopUpdateReqDTO struct {
	ShopID       domain.ShopID `json:"_shopid" bson:"_shopid"`
	Name         string        `json:"name" bson:"name"`
	Address      string        `json:"address" bson:"address"`
	AddressLine2 string        `json:"addressLine2" bson:"addressLine2"`
	Email        string        `json:"email" bson:"email"`
	Postcode     string        `json:"postcode" bson:"postcode"`
	Phone        string        `json:"phone" bson:"phone"`
	Rating       float32       `json:"rating" bson:"rating"`
}

type ShopDeteleReqDTO struct {
	ShopID       domain.ShopID `json:"_shopid" bson:"_shopid"`
	Name         string        `json:"name" bson:"name"`
	Address      string        `json:"address" bson:"address"`
	AddressLine2 string        `json:"addressLine2" bson:"addressLine2"`
	Email        string        `json:"email" bson:"email"`
	Postcode     string        `json:"postcode" bson:"postcode"`
	Phone        string        `json:"phone" bson:"phone"`
	Rating       float32       `json:"rating" bson:"rating"`
}
