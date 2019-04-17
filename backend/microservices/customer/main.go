package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	//"io/ioutil"
	dbrepo "OnlineFlorist/backend/microservices/customer/dbrepository"
	domain "OnlineFlorist/backend/microservices/customer/domain"
	mongoutils "OnlineFlorist/backend/microservices/customer/utils"
)

func main() {
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "customer"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	// fmt.Println(repoaccess)
	// //Run sample commands
	jsonFile, err := os.Open("/home/sumedha/go/src/OnlineFlorist/backend/microservices/customer/data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Json file opened successfully")
	defer jsonFile.Close()

	scanner := bufio.NewScanner(jsonFile)

	var cust domain.Customer
	//data :='{"name": "sumedha", "address": "SPPU PUNE"}'
	for scanner.Scan() {
		str := scanner.Text()
		json.Unmarshal([]byte(str), &cust)
		cust.CustID = domain.NewID()
		dbid, err1 := repoaccess.Store(&cust)
		fmt.Println(dbid, err1)
	}
	// args := os.Args[1:]
	// if len(args) == 2 {
	// 	query := strings.Replace(args[1], "--", "", 1)
	// 	data := strings.Split(query, "=")
	// 	key, value := data[0], data[1]
	// 	//key = strings.Replace(key, "--", "", 1)
	// 	if args[0] == "find" {
	// 		fmt.Println("\n\nResult By individual find function of ", key, "\n\n")
	// 		if key == "type_of_food" {
	// 			dbresult, dberr := repoaccess.FindByTypeOfFood(value)
	// 			fmt.Println(dbresult, dberr)
	// 		} else if key == "postcode" {
	// 			dbresult, dberr := repoaccess.FindByTypeOfPostCode(value)
	// 			fmt.Println(dbresult, dberr)
	// 		} else if key == "name" {
	// 			dbresult, dberr := repoaccess.FindByName(value)
	// 			fmt.Println(dbresult, dberr)

	// 		}
	// 	} else if args[0] == "count" {
	// 		if key == "type_of_food" {
	// 			dbresult, dberr := repoaccess.CountByTypeOfFood(value)
	// 			fmt.Println(dbresult, dberr)
	// 		} else if key == "postcode" {
	// 			dbresult, dberr := repoaccess.CountByTypeOfPostCode(value)
	// 			fmt.Println(dbresult, dberr)
	// 		}
	// 	} else if args[0] == "search" {
	// 		dbresult, dberr := repoaccess.Search(query)
	// 		fmt.Println(dbresult, dberr)

	// 	}
	// } else if args[0] == "get" {
	// 	fmt.Println("Hey")
	// 	dbresult, dberr := repoaccess.GetAll()
	// 	fmt.Println(dbresult, dberr)
	// }
	// //res.Name = "ABC"
	// //id, err := repoaccess.Create(&res)
	// //fmt.Println("create", id, err)
}
