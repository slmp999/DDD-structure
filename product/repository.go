package product

type Repository interface {
	StoreItem(tokenID string) error
	FindByID(product int64) (interface{}, error)
	FindAllProductRepo() ([]Item, error)
	FindProductByIDRepo(ID int64) ([]ItemDetail, error)
	FindCategoryRepo() ([]Category, error)
	FindCategoryByCodeRepo(Code string) ([]Item, error)
	FavoriteProductRepo() ([]Item, error)
	FindPackageRepo() ([]Item, error)
	FindYoutubeList() ([]YoutubeList, error)
	FindCampaign() ([]Campaign, error)
	FavoritePromotion() ([]Item, error)
	ItemList(ID int64) ([]ItemList, error)
	UnitList() ([]ItemUnit, error)
	ItemHistory(ID int64, UserID string) ([]ItemHistory, error)

	AddCategoryRepo(UserID string, req CategoryAdd) (interface{}, error)
	AddItemRepo(UserID string, req ItemAdd) (interface{}, error)
	AddPackage(UserID string, req ItemAdd) (interface{}, error)
	AddPromotion(UserID string, req ItemAdd) (interface{}, error)
	DeletePicture(ID int64, UserID string) (interface{}, error)
	DeleteProSub(ID int64, UserID string) (interface{}, error)
}
