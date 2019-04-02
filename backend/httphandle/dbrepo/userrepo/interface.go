package userrepo

import "gohttpexamples/sample4/domain"

type Reader interface {
	GetAll() ([]*domain.User, error)
	GetByID(ID string) (*domain.User, error)
}

type Writer interface {
	Create(*domain.User) (string, error)
	Update(*domain.User) error
	Delete(*domain.User) error
	//Archive(*domain.User) error
}

type Repository interface {
	Reader
	Writer
}
