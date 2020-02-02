package menu

type service struct {
	repo Repository
}

// NewService creates new auth service
func NewSales(sales Repository) (Service, error) {
	s := service{sales}
	return &s, nil
}

type Service interface {
	RegisterSaleService(UserID string, sale *RegisterSalesModel) (interface{}, error)
	ConfirmSaleService(AdminID string, SaleCode string, UserID string) (interface{}, error)
	ShowlistConfirmSales(Limit int64) (interface{}, error)
}

func (s *service) RegisterSaleService(UserID string, sale *RegisterSalesModel) (interface{}, error) {
	resp, err := s.repo.RegisterSaleRepo(UserID, sale)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) ConfirmSaleService(AdminID string, SaleCode string, UserID string) (interface{}, error) {
	resp, err := s.repo.SaleConfirmRepo(AdminID, SaleCode, UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) ShowlistConfirmSales(Limit int64) (interface{}, error) {
	resp, err := s.repo.ShowlistConfirmRepo(Limit)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
