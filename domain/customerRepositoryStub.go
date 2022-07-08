package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Alish", "New Delphi", "110011", "2000-01-01", "1"},
		{"1002", "Rob", "New Delphi", "110022", "2000-01-01", "2"},
	}

	return CustomerRepositoryStub{customers: customers}
}
