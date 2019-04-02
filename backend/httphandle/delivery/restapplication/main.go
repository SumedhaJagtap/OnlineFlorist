package main

import (
	handlerlib "gohttpexamples/sample4/delivery/restapplication/packages/httphandlers"
	restocrudhandler "gohttpexamples/sample4/delivery/restapplication/restocrudhandler"

	logger "log"
	dbrepo "mongorestaurantsample/dbrepository"
	mongoutils "mongorestaurantsample/utils"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	   https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/
	http.DefaultClient.Timeout = time.Minute * 10
}

func main() {
	// repoaccess := userrepo.NewUserInMemRepository()
	// usersvc := userrepo.NewService(repoaccess)
	// hndlr := usercrudhandler.NewUserCrudHandler(usersvc)

	//dbname := "restaurant"

	//repoaccess := userrepo.NewUserInMemRepository()

	//mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	//	dbname := "restaurant"
	//repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	//usersvc := userrepo.NewService(repoaccess)
	//repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	//dbsvc := userrepo.NewDbService(repoaccess)
	//hndlr := usercrudhandler.NewDbCrudHandler(dbsvc)
	//hndlr := userrepo.NewUserCrudHandler(usersvc)

	dbname := "restaurant"
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	dbsvc := dbrepo.NewDBService(repoaccess)
	hndlr := restocrudhandler.NewDbCrudHandler(dbsvc)

	pingHandler := &handlerlib.PingHandler{}
	logger.Println("Setting up resources.")
	logger.Println("Starting service")
	h := mux.NewRouter()
	// h.Handle("/ping/", pingHandler)
	//h.Handle("/user/{id}", hndlr)
	// h.Handle("/user/", hndlr)
	// h.Handle("/user/update/", hndlr)
	//h.Handle("/get/{id}", hndlr)
	//h.Handle("/post/{}", hndlr)
	h.Handle("/ping/", pingHandler)
	h.Handle("/restaurantservice/restaurant/", hndlr)
	//h.Handle("")
	//h.Handle("restaurantservice/restaurant/", hndlr)
	logger.Println("Resource Setup Done.")
	// h.HandleFunc("/restaurantservice/restaurant/{method}/find/{data}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	method := vars["method"]
	// 	data := vars["data"]
	// 	fmt.Fprintf(w, "Method:%s and data=%s", method, data)
	// })
	// //http.ListenAndServe(":8080", h)

	logger.Println(http.ListenAndServe(":8080", h))
}
