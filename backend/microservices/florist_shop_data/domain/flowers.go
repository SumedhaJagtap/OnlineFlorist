package domain
import  shop "OnlineFlorist/backend/microservices/florist_shop/domain"
  
//customer is a domain object
type  flower struct {
          FlowerID         ID      `json:"_flowerid" bson:"_flowerid"`
          Name         string  `json:"name" bson:"name"`
          Qty     int  `json:"qty" bson:"qty"`
          Rating     int  `json:"rating" bson:"rating"`
          Photo         string  `json:"photo" bson:"photo"`
          ShopID         shop.ID      `json:"_shopid" bson:"_shopid"`

}        
