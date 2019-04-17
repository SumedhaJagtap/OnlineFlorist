package floristshopcrudhandler

import "OnlineFlorist/backend/microservices/florist_shop/domain"

type ShopGetRespDTO struct {
	ShopID       domain.ShopID `json:"_shopid" bson:"_shopid"`
	Name         string        `json:"name" bson:"name"`
	Address      string        `json:"address" bson:"address"`
	AddressLine2 string        `json:"addressLine2" bson:"addressLine2"`
	Email        string        `json:"email" bson:"email"`
	Postcode     string        `json:"postcode" bson:"postcode"`
	Phone        string        `json:"phone" bson:"phone"`
	Rating       float32       `json:"rating" bson:"rating"`
}

type ShopGetListRespDTO struct {
	Shops []ShopGetRespDTO `json:"shops"`
	Count int              `json:"count"`
}

type ShopCreateRespDTO struct {
	ShopID domain.ShopID `json:"_shopid"`
}

type ShopUpdateRespDTO struct {
	ShopID domain.ShopID `json:"_shopid"`
}
