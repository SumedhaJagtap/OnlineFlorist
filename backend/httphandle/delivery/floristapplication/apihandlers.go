package main

import (
	floristshopcrudhandler "OnlineFlorist/backend/httphandle/delivery/floristapplication/floristshop_crudhandler"
	floristshopitemscrudhandler "OnlineFlorist/backend/httphandle/delivery/floristapplication/floristshop_items_crudhandler"
	handlerlib "OnlineFlorist/backend/httphandle/delivery/floristapplication/packages/httphandlers"
	usercrudhandler "OnlineFlorist/backend/httphandle/delivery/floristapplication/usercrudhandler"
	customer_dbrepo "OnlineFlorist/backend/microservices/customer/dbrepository"
	shop_items_dbrepo "OnlineFlorist/backend/microservices/florist_shop_items/dbrepository"
	"net/http"
	"time"

	logger "log"

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

func customer_handler() *usercrudhandler.UserCrudHandler {
	dbname := "customer"
	repoaccess := customer_dbrepo.NewMongoRepository(mongoSession, dbname)
	dbsvc := customer_dbrepo.NewService(repoaccess)
	hndlr := usercrudhandler.NewUserCrudHandler(dbsvc)
	return hndlr
}

func main() {
	shop_item_hndlr := shop_item_handler()
	shop_item_pingHandler := &handlerlib.PingHandler{}
	logger.Println("Setting up shop item  resources.")
	logger.Println("Starting shop item service")
	h1 := mux.NewRouter()
	h1.Handle("/ping/", shop_item_pingHandler)
	h1.Handle("/onlineflorist/shopitems/", shop_item_hndlr)
	logger.Println("shop item Resource Setup Done.")

	shop_hndlr := shop_handler()
	shop_pingHandler := &handlerlib.PingHandler{}
	logger.Println("Setting up shop resources.")
	logger.Println("Starting shop service")
	h2 := mux.NewRouter()
	h2.Handle("/ping/", shop_pingHandler)
	h2.Handle("/onlineflorist/shop/", shop_hndlr)
	logger.Println("shop Resource Setup Done.")

	user_hndlr := customer_handler()
	user_pingHandler := &handlerlib.PingHandler{}
	logger.Println("Setting up Cust resources.")
	logger.Println("Starting Cust service")
	h3 := mux.NewRouter()
	h3.Handle("/user/ping/", user_pingHandler)
	h3.Handle("/onlineflorist/customer/", user_hndlr)
	logger.Println("CUst Resource Setup Done.")

	go func() {
		logger.Println(http.ListenAndServe(":8080", h1))
	}()
	go func() {
		logger.Println(http.ListenAndServe(":8081", h2))
	}()
	go func() {
		logger.Println(http.ListenAndServe(":8082", h3))
	}()
	time.Sleep(time.Second * 1000)
	logger.Println("DONE")
}
