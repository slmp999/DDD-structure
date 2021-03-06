package sales

type Repository interface {
	RegisterSaleRepo(UserID string, sale *RegisterSalesModel) (interface{}, error)
	SaleConfirmRepo(AdminID string, SaleCode string, USerID string) (interface{}, error)
	ListSaleALL(Limit int64) (interface{}, error)
	ListSaleConfirmed(Limit int64) (interface{}, error)
	ListSaleNoConfirm(Limit int64) (interface{}, error)
	GetProfileSales(SaleCode string) (interface{}, error)
	GetProfileSalesbyUser(UserID string) (interface{}, error)
	UpdateSaleRepo(UserID string, sale *UpdateSaleModel) (interface{}, error)
	GetSalesTeamRepo(UserID string) (interface{}, error)
	RemoveSalesRepo(AdminID string, SaleCode string, USerID string) (interface{}, error)
	CallTokenLine(ID int64) (token string, err error)
	GetProfileSalesbyUserV2(UserID string) (interface{}, error)
	UpdateSalebyAdminRepo(AdminID string, sale *UpdateSaleModelAdmin) (interface{}, error)
	ListCommsionRepo(userID string) (interface{}, error)
	GenarateCommisionRepo(userID string) (interface{}, error)
	GetCommisionDocNoRepo(userID string, docno string) (*ModelCommision, error)
	ListCommsionBackendRepo(status int64, search string, limit int64) (interface{}, error)
	CountListCommmisonBackEnd(status int64, search string) (int64, error)
	ApproveCommisionRepo(adminID string, docno string, status int64, slipapprove string) (interface{}, error)
	GetCommisionDocNo(docno string) (*ModelCommision, error)
	CancelCommisionRepo(userID string, docno string) (interface{}, error)
	CancelCommisionFontRepo(userID string, docno string) (interface{}, error)
}
