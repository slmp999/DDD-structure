package report

type service struct {
	repo Repository
}

// NewService creates new auth service
func NewReport(report Repository) (Service, error) {
	s := service{report}
	return &s, nil
}

type Service interface {
	ListReportSale(startdate string, enddate string) (interface{}, error)
}

func (s *service) ListReportSale(startdate string, enddate string) (interface{}, error) {

	resp, err := s.repo.ListReportSaleRepo(startdate, enddate)
	if err != nil {
		return nil, err
	}
	amount, count, err := s.repo.SumReportSale(startdate, enddate)
	if err != nil {
		return nil, err
	}
	sale := map[string]interface{}{
		"date":       startdate + " --> " + enddate,
		"sum_amount": amount + 0,
		"length":     count,
		"list":       resp,
	}
	return sale, nil
}
