package userrepo

import (
	"gohttpexamples/sample4/domain"
	"gohttpexamples/sample4/utils"
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

func (s *Service) GetAll() ([]*domain.User, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(ID string) (*domain.User, error) {
	return s.repo.GetByID(ID)
}

func (s *Service) Create(u *domain.User) (string, error) {
	u.ID = utils.NewUUID()
	//u.CreatedOn = utils.GetUTCTimeNow()
	return s.repo.Create(u)

}

func (s *Service) Update(inp *domain.User) error {

	//inp.UpdatedOn = utils.GetUTCTImeNow()
	return s.repo.Update(inp)
}

func (s *Service) Delete(inp *domain.User) error {

	//inp.UpdatedOn = utils.GetUTCTImeNow()
	return s.repo.Delete(inp)
}
