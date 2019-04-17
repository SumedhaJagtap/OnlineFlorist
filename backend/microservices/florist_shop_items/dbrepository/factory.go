package dbrepository

import (
	shop "OnlineFlorist/backend/microservices/florist_shop/domain"
	"OnlineFlorist/backend/microservices/florist_shop_items/domain"
)

type Factory struct {
}

func (f *Factory) UpdateShopItemsDTO(ItemID domain.ItemID, Name string, Qty int, Rating float32, Category string, Photo string, ShopID shop.ShopID) *domain.FloristShopItems {
	return &domain.FloristShopItems{
		ItemID:   ItemID,
		Name:     Name,
		Qty:      Qty,
		Rating:   Rating,
		Category: Category,
		Photo:    Photo,
		ShopID:   ShopID,
	}

}

func (f *Factory) NewShopItemsDTO(ItemID domain.ItemID, Name string, Qty int, Rating float32, Category string, Photo string, ShopID shop.ShopID) *domain.FloristShopItems {
	return &domain.FloristShopItems{
		ItemID:   ItemID,
		Name:     Name,
		Qty:      Qty,
		Rating:   Rating,
		Category: Category,
		Photo:    Photo,
		ShopID:   ShopID,
	}

}
func (f *Factory) DeleteShopItemsDTO(ItemID domain.ItemID, Name string, Qty int, Rating float32, Category string, Photo string, ShopID shop.ShopID) *domain.FloristShopItems {
	return &domain.FloristShopItems{
		ItemID:   ItemID,
		Name:     Name,
		Qty:      Qty,
		Rating:   Rating,
		Category: Category,
		Photo:    Photo,
		ShopID:   ShopID,
	}

}
