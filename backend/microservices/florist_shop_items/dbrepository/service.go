package dbrepository

import (
	shop "OnlineFlorist/backend/microservices/florist_shop/domain"
	domain "OnlineFlorist/backend/microservices/florist_shop_items/domain"
)

type DBService struct {
	repo Repository
}

func NewDBService(repo Repository) *DBService {
	return &DBService{repo: repo}
}
func (s *DBService) SortByRating() ([]*domain.FloristShopItems, error) {
	return s.repo.SortByRating()
}
func (s *DBService) Create(b *domain.FloristShopItems) (domain.ItemID, error) {
	return s.repo.Create(b)
}

func (s *DBService) Get(id domain.ItemID) (*domain.FloristShopItems, error) {
	return s.repo.Get(id)

}

func (s *DBService) GetAll() ([]*domain.FloristShopItems, error) {
	return s.repo.GetAll()
}

func (s *DBService) FindByCategory(category string) ([]*domain.FloristShopItems, error) {
	return s.repo.FindByCategory(category)
}

func (s *DBService) Store(b *domain.FloristShopItems) (domain.ItemID, error) {
	return s.repo.Store(b)
}

func (s *DBService) Delete(id domain.ItemID) error {
	return s.repo.Delete(id)
}

func (s *DBService) FindByTypeOfFood(foodType string) ([]*domain.FloristShopItems, error) {
	return s.repo.FindByTypeOfFood(foodType)
}

func (s *DBService) FindByTypeOfPostCode(postCode string) ([]*domain.FloristShopItems, error) {
	return s.repo.FindByTypeOfPostCode(postCode)
}

func (s *DBService) Search(query string) ([]*domain.FloristShopItems, error) {
	return s.repo.Search(query)
}

func (s *DBService) FindShop(id domain.ItemID) *shop.FloristShop {
	return s.repo.FindShop(id)
}
