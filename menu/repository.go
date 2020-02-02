package menu

type Repository interface {
	RegisterSaleRepo(UserID string, sale *RegisterSalesModel) (interface{}, error)
	SaleConfirmRepo(AdminID string, SaleCode string, USerID string) (interface{}, error)
	ShowlistConfirmRepo(Limit int64) (interface{}, error)
}
