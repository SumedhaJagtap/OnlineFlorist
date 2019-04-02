package usercrudhandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	logger "log"
	"net/http"

	"gohttpexamples/sample4/dbrepo/userrepo"
	customerrors "gohttpexamples/sample4/delivery/restapplication/packages/errors"
	"gohttpexamples/sample4/delivery/restapplication/packages/httphandlers"
	mthdroutr "gohttpexamples/sample4/delivery/restapplication/packages/mthdrouter"
	"gohttpexamples/sample4/delivery/restapplication/packages/resputl"

	"github.com/gorilla/mux"
)

type UserCrudHandler struct {
	httphandlers.BaseHandler
	usersvc userrepo.Repository
}

func NewUserCrudHandler(usersvc userrepo.Repository) *UserCrudHandler {
	return &UserCrudHandler{usersvc: usersvc}
}

func (p *UserCrudHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}

//Get http method to get data
func (p *UserCrudHandler) Get(r *http.Request) resputl.SrvcRes {
	pathParam := mux.Vars(r)

	usID := pathParam["id"]
	if usID == "" {
		logger.Printf("%s", pathParam)
		//return resputl.Response200OK(generateSampleResponseObj())
		resp, err := p.usersvc.GetAll()

		if err != nil {
			return resputl.ReponseCustomError(err)
		}

		responseObj := transformobjListToResponse(resp)

		return resputl.Response200OK(responseObj)
	} else {
		obj, err := p.usersvc.GetByID(usID)

		if err != nil {
			return resputl.ProcessError(customerrors.NotFoundError("User Object Not found"), "")
		}

		userObj := UserGetRespDTO{
			ID:        obj.ID,
			FirstName: obj.Firstname,
			LastName:  obj.Lastname,
			CreatedOn: obj.CreatedOn,
			Age:       obj.Age,
		}

		return resputl.Response200OK(userObj)

	}

}

//Post method creates new temporary schedule
func (p *UserCrudHandler) Post(r *http.Request) resputl.SrvcRes {
	//logger.Print(r.URL)
	body, err := ioutil.ReadAll(r.Body)
	logger.Print(body, err)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	e, err := ValidateUserCreateUpdateRequest(string(body))
	if e == false {
		return resputl.ProcessError(err, body)
		//return resputl.SimpleBadRequest("Invalid Input Data")

	}
	logger.Printf("Received POST request to Create schedule %s ", string(body))
	var requestdata *UserCreateReqDTO
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	f := userrepo.Factory{}
	userObj := f.NewUser(requestdata.FirstName, requestdata.LastName, requestdata.Age)
	id, err := p.usersvc.Create(userObj)
	//fmt.Println("POST:", id)
	if err != nil {
		logger.Printf("POST : Error while creating in DB: %s", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK(&UserCreateRespDTO{ID: id})
}

//Put method modifies temporary schedule contents
func (p *UserCrudHandler) Put(r *http.Request) resputl.SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	logger.Print(body, err)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	e, err := ValidateUserCreateUpdateRequest(string(body))
	if e == false {
		return resputl.ProcessError(err, body)
		//return resputl.SimpleBadRequest("Invalid Input Data")
	}
	logger.Printf("Received PUT request to Create schedule %s ", string(body))
	var requestdata *UserUpdateReqDTO
	err = json.Unmarshal(body, &requestdata)
	fmt.Println(requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}
	fmt.Println("in put : ", requestdata)

	f := userrepo.Factory{}
	fmt.Println("ID : ", requestdata)
	userObj := f.UpdateUser(requestdata.ID, requestdata.FirstName, requestdata.LastName, requestdata.Age, requestdata.CreatedOn)
	fmt.Println("user obj:", userObj)
	err = p.usersvc.Update(userObj)
	if err != nil {
		logger.Printf("PUT : Error while creating in DB: %s", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK("Implememted")
}

//Delete method removes temporary schedule from db
func (p *UserCrudHandler) Delete(r *http.Request) resputl.SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	logger.Print(body, err)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	e, err := ValidateUserCreateUpdateRequest(string(body))
	if e == false {
		return resputl.ProcessError(err, body)
		//return resputl.SimpleBadRequest("Invalid Input Data")
	}
	logger.Printf("Received PUT request to Create schedule %s ", string(body))
	var requestdata *UserDeteleReqDTO
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}
	fmt.Println("in put : ", requestdata)

	f := userrepo.Factory{}
	fmt.Println("ID : ", requestdata)
	userObj := f.UpdateUser(requestdata.ID, requestdata.FirstName, requestdata.LastName, requestdata.Age, requestdata.CreatedOn)
	fmt.Println("user obj:", userObj)
	err = p.usersvc.Delete(userObj)
	if err != nil {
		logger.Printf("PUT : Error while creating in DB: %s", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}

	return resputl.Response200OK("Implememted")
}
