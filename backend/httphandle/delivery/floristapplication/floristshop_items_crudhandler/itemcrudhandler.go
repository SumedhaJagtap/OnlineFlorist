package floristshopitemscrudhandler

import (
	//"gohttpexamples/sample4/dbrepo/userrepo"
	customerrors "OnlineFlorist/backend/httphandle/delivery/floristapplication/packages/errors"
	mthdroutr "OnlineFlorist/backend/httphandle/delivery/floristapplication/packages/mthdrouter"
	"OnlineFlorist/backend/httphandle/delivery/floristapplication/packages/resputl"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gorilla/mux"

	"OnlineFlorist/backend/httphandle/delivery/floristapplication/packages/httphandlers"
	dbrepo "OnlineFlorist/backend/microservices/florist_shop_items/dbrepository"
	"OnlineFlorist/backend/microservices/florist_shop_items/domain"
	logger "log"
	"net/http"
)

type DbCrudHandler struct {
	httphandlers.BaseHandler
	dbsvc dbrepo.Repository
}

func NewDbCrudHandler(dbsvc dbrepo.Repository) *DbCrudHandler {
	return &DbCrudHandler{dbsvc: dbsvc}
}
func (p *DbCrudHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}

func (p *DbCrudHandler) Get(r *http.Request) resputl.SrvcRes {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("k:", k, "v:", v[0])
	}
	pathParam := mux.Vars(r)
	for k, v := range pathParam {
		fmt.Println("k:", k, "v params:", v)
	}
	//fmt.Println("name:", pathParam["name"])
	//fmt.Println("body:", r.Body)
	usID := pathParam["id"]
	if usID == "" {
		//logger.Printf("%s", pathParam)
		//fmt.Println("Path: ", r.URL.Path, " params :", r.Form["name"][0])
		//return resputl.Response200OK(generateSampleResponseObj())
		if len(r.Form) > 0 {

			// if len(r.Form["name"]) > 0 {
			// 	resp, err := p.dbsvc.FindByName(r.Form["name"][0])
			// 	responseObj := transformobjListToResponse(resp)

			// 	if err != nil {
			// 		return resputl.ReponseCustomError(err)
			// 	} else {
			// 		return resputl.Response200OK(responseObj)
			// 	}

			// }

			// if r.Form["type_of_food"][0] != "" {
			// 	resp, err := p.dbsvc.FindByTypeOfFood(r.Form["type_of_food"][0])
			// 	responseObj := transformobjListToResponse(resp)

			// 	if err != nil {
			// 		return resputl.ReponseCustomError(err)
			// 	} else {
			// 		return resputl.Response200OK(responseObj)
			// 	}
			// }
			if r.Form["searchTerm"][0] != "" {
				resp, err := p.dbsvc.Search("name=" + r.Form["searchTerm"][0])
				responseObj := transformobjListToResponse(resp)

				if err != nil {
					return resputl.ReponseCustomError(err)
				} else {
					return resputl.Response200OK(responseObj)
				}
			}
		}
		resp, err := p.dbsvc.GetAll()
		responseObj := transformobjListToResponse(resp)
		if err != nil {
			return resputl.ReponseCustomError(err)
		} else {
			return resputl.Response200OK(responseObj)
		}
	} else {
		obj, err := p.dbsvc.Get(domain.ItemID(usID))

		if err != nil {
			return resputl.ProcessError(customerrors.NotFoundError("User Object Not found"), "")
		}

		itemObj := ShopItemsGetRespDTO{
			ItemID:   obj.ItemID,
			Name:     obj.Name,
			Qty:      obj.Qty,
			Category: obj.Category,
			Photo:    obj.Photo,
			Rating:   obj.Rating,
			ShopID:   obj.ShopID,
		}

		return resputl.Response200OK(itemObj)

	}

}

func (p *DbCrudHandler) Put(r *http.Request) resputl.SrvcRes {
	fmt.Println("In Put")
	body, err := ioutil.ReadAll(r.Body)
	logger.Print(body, err)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	e, err := ValidateShopItemsCreateUpdateRequest(string(body))
	if e == false {
		fmt.Println("Error***********")
		return resputl.ProcessError(err, body)
		//return resputl.SimpleBadRequest("Invalid Input Data")
	}
	logger.Printf("Received PUT request to Create schedule %s ", string(body))
	var requestdata *ShopItemsUpdateReqDTO
	err = json.Unmarshal(body, &requestdata)
	fmt.Println(requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}
	fmt.Println("in put : ", requestdata)

	f := dbrepo.Factory{}
	fmt.Println("ID : ", requestdata)
	userObj := f.UpdateShopItemsDTO(requestdata.ItemID, requestdata.Name, requestdata.Qty, requestdata.Rating, requestdata.Category, requestdata.Photo, requestdata.ShopID)
	fmt.Println("user obj:", userObj)
	_, err = p.dbsvc.Store(userObj)
	if err != nil {
		logger.Printf("PUT : Error while creating in DB: %s", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK("Implememted")
}

func (p *DbCrudHandler) Post(r *http.Request) resputl.SrvcRes {
	fmt.Println("In POST")
	body, err := ioutil.ReadAll(r.Body)
	logger.Print(body, err)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	e, err := ValidateShopItemsCreateUpdateRequest(string(body))
	if e == false {
		fmt.Println("Errroooooooooooor")
		return resputl.ProcessError(err, body)
		//return resputl.SimpleBadRequest("Invalid Input Data")
	}
	logger.Printf("Received POST request to Create schedule %s ", string(body))
	var requestdata *ShopItemsCreateReqDTO
	err = json.Unmarshal(body, &requestdata)
	fmt.Println(requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}
	fmt.Println("in post : ", requestdata)

	f := dbrepo.Factory{}
	fmt.Println("ID : ", requestdata)
	userObj := f.NewShopItemsDTO(requestdata.ItemID, requestdata.Name, requestdata.Qty, requestdata.Rating, requestdata.Category, requestdata.Photo, requestdata.ShopID)
	fmt.Println("user obj:", userObj)
	_, err = p.dbsvc.Create(userObj)
	if err != nil {
		logger.Printf("PUT : Error while creating in DB: %s", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK("Implememted")
}

func (p *DbCrudHandler) Delete(r *http.Request) resputl.SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	logger.Print(body, err)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	e, err := ValidateShopItemsCreateUpdateRequest(string(body))
	if e == false {
		return resputl.ProcessError(err, body)
		//return resputl.SimpleBadRequest("Invalid Input Data")
	}
	logger.Printf("Received PUT request to Create schedule %s ", string(body))
	var requestdata *ShopItemsDeteleReqDTO
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}
	fmt.Println("in put : ", requestdata)

	f := dbrepo.Factory{}
	fmt.Println("ID : ", requestdata)
	userObj := f.DeleteShopItemsDTO(requestdata.ItemID, requestdata.Name, requestdata.Qty, requestdata.Rating, requestdata.Category, requestdata.Photo, requestdata.ShopID)
	fmt.Println("user obj:", userObj)
	err = p.dbsvc.Delete(userObj.ItemID)
	if err != nil {
		logger.Printf("PUT : Error while creating in DB: %s", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK("Implememted")
}
