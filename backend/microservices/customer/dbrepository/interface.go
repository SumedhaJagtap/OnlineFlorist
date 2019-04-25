package dbrepository

import "OnlineFlorist/backend/microservices/customer/domain"

type Reader interface {
	GetAll() ([]*domain.Customer, error)
	GetByID(ID domain.ID) (*domain.Customer, error)
}

type Writer interface {
	Create(*domain.Customer) (domain.ID, error)
	Update(*domain.Customer) error
	Delete(domain.ID) error
	//Archive(*domain.User) error
}

type Repository interface {
	Reader
	Writer
}
