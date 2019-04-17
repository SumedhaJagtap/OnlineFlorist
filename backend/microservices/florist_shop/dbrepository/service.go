package dbrepository

import (
	domain "OnlineFlorist/backend/microservices/florist_shop/domain"
)

type DBService struct {
	repo Repository
}

func NewDBService(repo Repository) *DBService {
	return &DBService{repo: repo}
}

func (s *DBService) Create(b *domain.FloristShop) (domain.ShopID, error) {
	return s.repo.Create(b)
}

func (s *DBService) Get(id domain.ShopID) (*domain.FloristShop, error) {
	return s.repo.Get(id)

}

func (s *DBService) GetAll() ([]*domain.FloristShop, error) {
	return s.repo.GetAll()
}

func (s *DBService) FindByName(name string) ([]*domain.FloristShop, error) {
	return s.repo.FindByName(name)
}

func (s *DBService) Store(b *domain.FloristShop) (domain.ShopID, error) {
	return s.repo.Store(b)
}

func (s *DBService) Delete(id domain.ShopID) error {
	return s.repo.Delete(id)
}

func (s *DBService) FindByItemCategory(ItemCategory string) ([]*domain.FloristShop, error) {
	return s.repo.FindByItemCategory(ItemCategory)
}

func (s *DBService) FindByPostCode(postCode string) ([]*domain.FloristShop, error) {
	return s.repo.FindByPostCode(postCode)
}

func (s *DBService) Search(query string) ([]*domain.FloristShop, error) {
	return s.repo.Search(query)
}
