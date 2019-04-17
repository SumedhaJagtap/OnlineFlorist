package dbrepository

import domain "OnlineFlorist/backend/microservices/florist_shop_items/domain"
import shop "OnlineFlorist/backend/microservices/florist_shop/domain"

//Reader read from db
type Reader interface {
	Get(id domain.ItemID) (*domain.FloristShopItems, error)
	GetAll() ([]*domain.FloristShopItems, error)
	//Regex Substring Match on the name field
	FindByCategory(category string) ([]*domain.FloristShopItems, error)
	SortByRating() ([]*domain.FloristShopItems, error)
	FindShop(id domain.ItemID) *shop.FloristShop
}

//Writer  write to db
type Writer interface {
	//Create Or update
	Create(b *domain.FloristShopItems) (domain.ItemID, error)
	Store(b *domain.FloristShopItems) (domain.ItemID, error)
	Delete(id domain.ItemID) error
}

//Filter Find objects by additional filters
type Filter interface {
	FindByTypeOfFood(foodType string) ([]*domain.FloristShopItems, error)
	FindByTypeOfPostCode(postCode string) ([]*domain.FloristShopItems, error)
	//Search --> across all string fields regex match with case insensitive
	//substring match accross all string fields
	Search(query string) ([]*domain.FloristShopItems, error)
}

//Repository db interface
type Repository interface {
	Reader
	Writer
	Filter
}
