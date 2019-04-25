package dbrepository

import (
	"OnlineFlorist/backend/microservices/customer/domain"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// type DbService struct {
// 	dbrepo dbrepo.Repository
// }

// func NewDbService(dbrepo dbrepo.Repository) *Service {
// 	return &Service{repo: dbrepo}
// }

// type Reader interface {
// 	GetAll() ([]*domain.User, error)
// 	GetByID() (*domain.User, error)
// }

// type Writer interface {
// 	Create(*domain.User) (string, error)
// 	Update(*domain.User) error
// 	Archive(*domain.User) error
// }

func (s *Service) GetAll() ([]*domain.Customer, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(ID domain.ID) (*domain.Customer, error) {
	return s.repo.GetByID(ID)
}

func (s *Service) Create(u *domain.Customer) (domain.ID, error) {
	// u.CustID = utils.NewUUID()
	// u.CreatedOn = utils.GetUTCTimeNow()
	return s.repo.Create(u)

}

func (s *Service) Update(inp *domain.Customer) error {

	//inp.UpdatedOn = utils.GetUTCTImeNow()
	return s.repo.Update(inp)
}

func (s *Service) Delete(id domain.ID) error {

	//inp.UpdatedOn = utils.GetUTCTImeNow()
	return s.repo.Delete(id)
}
