package dbrepository

import domain "OnlineFlorist/backend/microservices/florist_shop/domain"

//Reader read from db
type Reader interface {
	Get(shopid domain.ShopID) (*domain.FloristShop, error)
	GetAll() ([]*domain.FloristShop, error)
	//Regex Substring Match on the name field
	FindByName(name string) ([]*domain.FloristShop, error)
}

//Writer  write to db
type Writer interface {
	//Create Or update
	Create(b *domain.FloristShop) (domain.ShopID, error)
	Store(b *domain.FloristShop) (domain.ShopID, error)
	Delete(id domain.ShopID) error
}

//Filter Find objects by additional filters
type Filter interface {
	FindByItemCategory(ItemCategory string) ([]*domain.FloristShop, error)
	FindByPostCode(postCode string) ([]*domain.FloristShop, error)
	//Search --> across all string fields regex match with case insensitive
	//substring match accross all string fields
	Search(query string) ([]*domain.FloristShop, error)
}

//Repository db interface
type Repository interface {
	Reader
	Writer
	Filter
}
