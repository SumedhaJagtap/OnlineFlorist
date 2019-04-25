package restocrudhandler

import (
	"fmt"
	"mongorestaurantsample/domain"
)

func transformobjListToResponse(resp []*domain.Restaurant) RestoGetListRespDTO {
	responseObj := RestoGetListRespDTO{}
	fmt.Println(resp)
	for _, obj := range resp {
		restoObj := RestoGetRespDTO{
			DBID:       obj.DBID,
			Name:       obj.Name,
			Address:    obj.Address,
			Postcode:   obj.Postcode,
			Rating:     obj.Rating,
			TypeOfFood: obj.TypeOfFood,
		}
		responseObj.Resto = append(responseObj.Resto, restoObj)
		fmt.Println(restoObj)
	}

	responseObj.Count = len(responseObj.Resto)

	return responseObj
}
