package dbrepository

import "OnlineFlorist/backend/microservices/customer/domain"

type Factory struct {
}

func (f *Factory) UpdateCustomerDTO(CustID domain.ID, Email string, Phone string, FirstName string, LastName string, CreatedOn uint64) *domain.Customer {
	return &domain.Customer{

		CustID:    CustID,
		Email:     Email,
		Phone:     Phone,
		FirstName: FirstName,

		LastName:  LastName,
		CreatedOn: CreatedOn,
	}

}

func (f *Factory) NewCustomerDTO(Email string, Phone string, FirstName string, LastName string) *domain.Customer {
	return &domain.Customer{
		Email:     Email,
		Phone:     Phone,
		FirstName: FirstName,

		LastName: LastName,
	}

}

func (f *Factory) DeleteCustomerDTO(CustID domain.ID, Email string, Phone string, FirstName string, LastName string, CreatedOn uint64) *domain.Customer {
	return &domain.Customer{

		CustID:    CustID,
		Email:     Email,
		Phone:     Phone,
		FirstName: FirstName,

		LastName:  LastName,
		CreatedOn: CreatedOn,
	}

}
