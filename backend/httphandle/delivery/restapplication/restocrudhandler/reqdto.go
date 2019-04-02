package restocrudhandler

import "mongorestaurantsample/domain"

type RestoCreateReqDTO struct {
	DBID         domain.ID `json:"_id" `
	Name         string    `json:"name" `
	Address      string    `json:"address" `
	AddressLine2 string    `json:"addressLine2" `
	URL          string    `json:"url" `
	Outcode      string    `json:"outcode" `
	Postcode     string    `json:"postcode" `
	Rating       float32   `json:"rating" `
	TypeOfFood   string    `json:"type_of_food" `
}

type RestoUpdateReqDTO struct {
	DBID         domain.ID `json:"_id" `
	Name         string    `json:"name" `
	Address      string    `json:"address" `
	AddressLine2 string    `json:"addressLine2" `
	URL          string    `json:"url" `
	Outcode      string    `json:"outcode" `
	Postcode     string    `json:"postcode" `
	Rating       float32   `json:"rating" `
	TypeOfFood   string    `json:"type_of_food" `
}

type RestoDeteleReqDTO struct {
	DBID         domain.ID `json:"_id" `
	Name         string    `json:"name" `
	Address      string    `json:"address" `
	AddressLine2 string    `json:"addressLine2" `
	URL          string    `json:"url" `
	Outcode      string    `json:"outcode" `
	Postcode     string    `json:"postcode" `
	Rating       float32   `json:"rating" `
	TypeOfFood   string    `json:"type_of_food" `
}
