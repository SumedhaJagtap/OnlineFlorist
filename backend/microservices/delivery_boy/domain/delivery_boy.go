package domain
  
//customer is a domain object
type deliveryBoy struct {
          DelBoyID         ID      `json:"_delboyid" bson:"_delboyid"`
          Name         string  `json:"name" bson:"name"`
          Address      string  `json:"address" bson:"address"`
          AddressLine2 string  `json:"addressLine2" bson:"addressLine2"`
          email          string  `json:"email" bson:"email"`
          Postcode     string  `json:"postcode" bson:"postcode"`
          Phone   string  `json:"phone" bson:"phone"`
          Photo   string  `json:"photo" bson:"photo"`
          Bike   string  `json:"bike" bson:"bike"`
}        
