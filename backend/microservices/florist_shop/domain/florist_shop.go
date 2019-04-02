package domain
  
//customer is a domain object
type floristshop struct {
          ShopID         ID      `json:"_shopid" bson:"_shopid"`
          Name         string  `json:"name" bson:"name"`
          Address      string  `json:"address" bson:"address"`
          AddressLine2 string  `json:"addressLine2" bson:"addressLine2"`
          email          string  `json:"email" bson:"email"`
          Postcode     string  `json:"postcode" bson:"postcode"`
          Phone   string  `json:"phone" bson:"phone"`
}        
