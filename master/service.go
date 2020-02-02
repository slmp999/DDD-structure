package master

type service struct {
	repo Repository
}

// NewService creates new auth service
func NewService(master Repository) (Service, error) {
	s := service{master}
	return &s, nil
}

//  Define Service
type Service interface {
	BankList() (interface{}, error)
}

func (s *service) BankList() (interface{}, error) {
	resp, err := s.repo.GetBankList()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
