package floristshopitemscrudhandler

import (
	shop "OnlineFlorist/backend/microservices/florist_shop/domain"
	"OnlineFlorist/backend/microservices/florist_shop_items/domain"
)

type ShopItemsCreateReqDTO struct {
	ItemID   domain.ItemID `json:"_itemid" bson:"_itemid"`
	Name     string        `json:"name" bson:"name"`
	Qty      int           `json:"qty" bson:"qty"`
	Rating   float32       `json:"rating" bson:"rating"`
	Category string        `json:"category" bson:"category"`
	Photo    string        `json:"photo" bson:"photo"`
	ShopID   shop.ShopID   `json:"_shopid" bson:"_shopid"`
}

type ShopItemsUpdateReqDTO struct {
	ItemID   domain.ItemID `json:"_itemid" bson:"_itemid"`
	Name     string        `json:"name" bson:"name"`
	Qty      int           `json:"qty" bson:"qty"`
	Rating   float32       `json:"rating" bson:"rating"`
	Category string        `json:"category" bson:"category"`
	Photo    string        `json:"photo" bson:"photo"`
	ShopID   shop.ShopID   `json:"_shopid" bson:"_shopid"`
}

type ShopItemsDeteleReqDTO struct {
	ItemID   domain.ItemID `json:"_itemid" bson:"_itemid"`
	Name     string        `json:"name" bson:"name"`
	Qty      int           `json:"qty" bson:"qty"`
	Rating   float32       `json:"rating" bson:"rating"`
	Category string        `json:"category" bson:"category"`
	Photo    string        `json:"photo" bson:"photo"`
	ShopID   shop.ShopID   `json:"_shopid" bson:"_shopid"`
}
