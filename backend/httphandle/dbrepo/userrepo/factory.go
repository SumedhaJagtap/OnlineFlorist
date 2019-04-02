package userrepo

import (
	"gohttpexamples/sample4/domain"
	"gohttpexamples/sample4/utils"
)

type Factory struct {
}

func (f *Factory) NewUser(firstname string, lastname string, age int) *domain.User {
	return &domain.User{
		Firstname: firstname,
		Lastname:  lastname,
		Age:       age,
		CreatedOn: utils.GetUTCTimeNow(),
	}

}

func (f *Factory) UpdateUser(id string, firstname string, lastname string, age int, CreatedOn uint64) *domain.User {
	return &domain.User{
		ID:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Age:       age,
		CreatedOn: CreatedOn,
	}

}
