package report

type Repository interface {
	SumReportSale(startdate string, enddate string) (float64, float64, error)
	ListReportSaleRepo(startdate string, enddate string) (interface{}, error)
}
