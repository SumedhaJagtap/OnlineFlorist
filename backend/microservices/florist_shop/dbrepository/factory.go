package dbrepository

import "OnlineFlorist/backend/microservices/florist_shop/domain"

type Factory struct {
}

func (f *Factory) UpdateShopDTO(ShopID domain.ShopID, Name string, Address string, AddressLine2 string, Email string, Postcode string, Phone string, Rating float32) *domain.FloristShop {
	return &domain.FloristShop{
		ShopID:       ShopID,
		Name:         Name,
		Address:      Address,
		AddressLine2: AddressLine2,
		Email:        Email,
		Postcode:     Postcode,
		Rating:       Rating,
		Phone:        Phone,
	}

}

func (f *Factory) NewShopDTO(ShopID domain.ShopID, Name string, Address string, AddressLine2 string, Email string, Postcode string, Phone string, Rating float32) *domain.FloristShop {
	return &domain.FloristShop{
		ShopID:       ShopID,
		Name:         Name,
		Address:      Address,
		AddressLine2: AddressLine2,
		Email:        Email,
		Postcode:     Postcode,
		Rating:       Rating,
		Phone:        Phone,
	}

}

func (f *Factory) DeleteShopDTO(ShopID domain.ShopID, Name string, Address string, AddressLine2 string, Email string, Postcode string, Phone string, Rating float32) *domain.FloristShop {
	return &domain.FloristShop{
		ShopID:       ShopID,
		Name:         Name,
		Address:      Address,
		AddressLine2: AddressLine2,
		Email:        Email,
		Postcode:     Postcode,
		Rating:       Rating,
		Phone:        Phone,
	}

}
