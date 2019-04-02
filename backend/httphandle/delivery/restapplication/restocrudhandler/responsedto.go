package restocrudhandler

import "mongorestaurantsample/domain"

type RestoGetRespDTO struct {
	DBID       domain.ID `json:"_id"`
	Name       string    `json:"name"`
	Postcode   string    `json:"postcode" `
	Rating     float32   `json:"rating" `
	TypeOfFood string    `json:"type_of_food" `
	Address    string    `json:"address"`
}

type RestoGetListRespDTO struct {
	Resto []RestoGetRespDTO `json:"resto"`
	Count int               `json:"count"`
}

type RestoCreateRespDTO struct {
	DBID domain.ID `json:"_id"`
}

type RestoUpdateRespDTO struct {
	DBID domain.ID `json:"_id"`
}
