package main

import (
	shop_dbrepo "OnlineFlorist/backend/microservices/florist_shop/dbrepository"
	mongoutils "OnlineFlorist/backend/microservices/florist_shop/utils"
	"os"
)

func main() {
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	shop_db := "floristshop"
	shop_repoaccess := shop_dbrepo.NewMongoRepository(mongoSession, shop_db)

}
