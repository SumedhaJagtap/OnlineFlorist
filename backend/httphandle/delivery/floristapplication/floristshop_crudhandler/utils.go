package floristshopcrudhandler

import (
	"OnlineFlorist/backend/microservices/florist_shop/domain"
	"fmt"
)

func transformobjListToResponse(resp []*domain.FloristShop) ShopGetListRespDTO {
	responseObj := ShopGetListRespDTO{}
	fmt.Println(resp)
	for _, obj := range resp {
		shopObj := ShopGetRespDTO{
			ShopID:   obj.ShopID,
			Name:     obj.Name,
			Address:  obj.Address,
			Postcode: obj.Postcode,
			Email:    obj.Email,
			Phone:    obj.Phone,
			Rating:   obj.Rating,
		}
		responseObj.Shops = append(responseObj.Shops, shopObj)
		fmt.Println(shopObj)
	}

	responseObj.Count = len(responseObj.Shops)

	return responseObj
}
