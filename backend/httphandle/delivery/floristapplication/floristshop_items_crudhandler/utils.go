package floristshopitemscrudhandler

import (
	"OnlineFlorist/backend/microservices/florist_shop_items/domain"
	"fmt"
)

func transformobjListToResponse(resp []*domain.FloristShopItems) ShopItemsGetListRespDTO {
	responseObj := ShopItemsGetListRespDTO{}
	fmt.Println(resp)
	for _, obj := range resp {
		shopItemsObj := ShopItemsGetRespDTO{
			ItemID:   obj.ItemID,
			ShopID:   obj.ShopID,
			Name:     obj.Name,
			Qty:      obj.Qty,
			Photo:    obj.Photo,
			Rating:   obj.Rating,
			Category: obj.Category,
		}
		responseObj.Items = append(responseObj.Items, shopItemsObj)
		fmt.Println(shopItemsObj)
	}

	responseObj.Count = len(responseObj.Items)

	return responseObj
}
