package dbrepository

import (
	"mongorestaurantsample/domain"
)

type DBService struct {
	repo Repository
}

func NewDBService(repo Repository) *DBService {
	return &DBService{repo: repo}
}

func (s *DBService) Create(b *domain.Restaurant) (domain.ID, error) {
	return s.repo.Create(b)
}

func (s *DBService) Get(id domain.ID) (*domain.Restaurant, error) {
	return s.repo.Get(id)

}

func (s *DBService) GetAll() ([]*domain.Restaurant, error) {
	return s.repo.GetAll()
}

func (s *DBService) FindByName(name string) ([]*domain.Restaurant, error) {
	return s.repo.FindByName(name)
}

func (s *DBService) Store(b *domain.Restaurant) (domain.ID, error) {
	return s.repo.Store(b)
}

func (s *DBService) Delete(id domain.ID) error {
	return s.repo.Delete(id)
}

func (s *DBService) FindByTypeOfFood(foodType string) ([]*domain.Restaurant, error) {
	return s.repo.FindByTypeOfFood(foodType)
}

func (s *DBService) FindByTypeOfPostCode(postCode string) ([]*domain.Restaurant, error) {
	return s.repo.FindByTypeOfPostCode(postCode)
}

func (s *DBService) Search(query string) ([]*domain.Restaurant, error) {
	return s.repo.Search(query)
}
