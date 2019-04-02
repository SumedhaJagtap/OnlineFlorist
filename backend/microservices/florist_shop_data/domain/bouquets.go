package domain
  
//customer is a domain object
type  bouquet struct {
          BouquetID         ID      `json:"_bouquetid" bson:"_bouquetid"`
          Name         string  `json:"name" bson:"name"`
          Qty     int  `json:"qty" bson:"qty"`
          Rating     int  `json:"rating" bson:"rating"`

          Photo         string  `json:"photo" bson:"photo"`
}        
