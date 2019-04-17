package dbrepository

import "mongorestaurantsample/domain"

type Factory struct {
}

func (f *Factory) UpdateRestoDTO(DBID domain.ID, Name string, Address string, AddressLine2 string, URL string, Outcode string, Postcode string, Rating float32, TypeOfFood string) *domain.Restaurant {
	return &domain.Restaurant{
		DBID:         DBID,
		Name:         Name,
		Address:      Address,
		AddressLine2: AddressLine2,
		URL:          URL,
		Outcode:      Outcode,
		Postcode:     Postcode,
		Rating:       Rating,
		TypeOfFood:   TypeOfFood,
	}

}

func (f *Factory) NewRestoDTO(DBID domain.ID, Name string, Address string, AddressLine2 string, URL string, Outcode string, Postcode string, Rating float32, TypeOfFood string) *domain.Restaurant {
	return &domain.Restaurant{
		DBID:         DBID,
		Name:         Name,
		Address:      Address,
		AddressLine2: AddressLine2,
		URL:          URL,
		Outcode:      Outcode,
		Postcode:     Postcode,
		Rating:       Rating,
		TypeOfFood:   TypeOfFood,
	}

}

func (f *Factory) DeleteRestoDTO(DBID domain.ID, Name string, Address string, AddressLine2 string, URL string, Outcode string, Postcode string, Rating float32, TypeOfFood string) *domain.Restaurant {
	return &domain.Restaurant{
		DBID:         DBID,
		Name:         Name,
		Address:      Address,
		AddressLine2: AddressLine2,
		URL:          URL,
		Outcode:      Outcode,
		Postcode:     Postcode,
		Rating:       Rating,
		TypeOfFood:   TypeOfFood,
	}

}
