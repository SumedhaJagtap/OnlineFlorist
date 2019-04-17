package domain

import shop "OnlineFlorist/backend/microservices/florist_shop/domain"

//customer is a domain object
type bouquet struct {
	BouquetID ItemID      `json:"_bouquetid" bson:"_bouquetid"`
	Name      string      `json:"name" bson:"name"`
	Qty       int         `json:"qty" bson:"qty"`
	Rating    int         `json:"rating" bson:"rating"`
	Category  string      `json:"category" bson:"category"`
	Photo     string      `json:"photo" bson:"photo"`
	ShopID    shop.ShopID `json:"_shopid" bson:"_shopid"`
}
