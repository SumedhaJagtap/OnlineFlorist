package main

import (
	floristshopcrudhandler "OnlineFlorist/backend/httphandle/delivery/floristapplication/floristshop_crudhandler"
	floristshopitemscrudhandler "OnlineFlorist/backend/httphandle/delivery/floristapplication/floristshop_items_crudhandler"
	handlerlib "OnlineFlorist/backend/httphandle/delivery/floristapplication/packages/httphandlers"
	shop_items_dbrepo "OnlineFlorist/backend/microservices/florist_shop_items/dbrepository"
	logger "log"
	"net/http"

	shop_dbrepo "OnlineFlorist/backend/microservices/florist_shop/dbrepository"
	mongoutils "mongorestaurantsample/utils"
	"os"

	"github.com/gorilla/mux"
)

var mongoSession, _ = mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

func shop_handler() *floristshopcrudhandler.DbCrudHandler {
	dbname := "floristshop"
	repoaccess := shop_dbrepo.NewMongoRepository(mongoSession, dbname)
	dbsvc := shop_dbrepo.NewDBService(repoaccess)
	hndlr := floristshopcrudhandler.NewDbCrudHandler(dbsvc)
	return hndlr
}

func shop_item_handler() *floristshopitemscrudhandler.DbCrudHandler {
	dbname := "floristshopitems"
	repoaccess := shop_items_dbrepo.NewMongoRepository(mongoSession, dbname)
	dbsvc := shop_items_dbrepo.NewDBService(repoaccess)
	hndlr := floristshopitemscrudhandler.NewDbCrudHandler(dbsvc)
	return hndlr
}

func main() {
	shop_item_hndlr := shop_item_handler()
	shop_item_pingHandler := &handlerlib.PingHandler{}
	logger.Println("Setting up resources.")
	logger.Println("Starting service")
	h1 := mux.NewRouter()
	h1.Handle("/ping/", shop_item_pingHandler)
	h1.Handle("/onlineflorist/shopitems/", shop_item_hndlr)
	// logger.Println("Resource Setup Done.")
	// logger.Println(http.ListenAndServe(":8080", h1))

	shop_hndlr := shop_handler()
	shop_pingHandler := &handlerlib.PingHandler{}
	// logger.Println("Setting up resources.")
	// logger.Println("Starting service")
	h2 := mux.NewRouter()
	h2.Handle("/ping/", shop_pingHandler)
	h2.Handle("/onlineflorist/shop/", shop_hndlr)
	logger.Println("Resource Setup Done.")
	logger.Println(http.ListenAndServe(":8080", h2))

}
