package usercrudhandler

import "OnlineFlorist/backend/microservices/customer/domain"

func transformobjListToResponse(resp []*domain.Customer) UserGetListRespDTO {
	responseObj := UserGetListRespDTO{}
	for _, obj := range resp {
		userObj := UserGetRespDTO{
			CustID:    obj.CustID,
			Email:     obj.Email,
			Phone:     obj.Phone,
			FirstName: obj.FirstName,
			LastName:  obj.LastName,
			// Age:       obj.Age,
			CreatedOn: obj.CreatedOn,
		}
		responseObj.Users = append(responseObj.Users, userObj)
	}
	responseObj.Count = len(responseObj.Users)

	return responseObj
}
