package product

// Service is the device service
type Service interface {
	FindByID(productID int64) (interface{}, error)
	FindCategoryService() ([]Category, error)
	FindAllProductService() ([]Item, error)
	FindCategoryByCodeService(Code string) ([]Item, error)
	FindProductByIDService(ID int64) ([]ItemDetail, error)
	FavoriteProductService() ([]Item, error)
	FindPackageService() ([]Item, error)
	FindYoutubeList() ([]YoutubeList, error)
	FindCampaign() ([]Campaign, error)
	FavoritePromotion() ([]Item, error)
	ItemList(ID int64) ([]ItemList, error)
	UnitList() ([]ItemUnit, error)
	ItemHistory(ID int64, UserID string) ([]ItemHistory, error)

	AddCategoryService(UserID string, req CategoryAdd) (interface{}, error)
	AddItemService(UserID string, req ItemAdd) (interface{}, error)
	AddPackage(UserID string, req ItemAdd) (interface{}, error)
	AddPromotion(UserID string, req ItemAdd) (interface{}, error)
	DeletePicture(ID int64, UserID string) (interface{}, error)
	DeleteProSub(ID int64, UserID string) (interface{}, error)
	StoreItem()
}

// NewService creates new device service
func NewService(product Repository) (Service, error) {
	s := service{product}
	return &s, nil
}

type service struct {
	repo Repository
}

func (s *service) FindByID(producID int64) (interface{}, error) {
	s.repo.FindByID(12)
	return nil, nil
}

func (s *service) StoreItem() {
	return
}

func (s *service) FindCategoryService() ([]Category, error) {
	resp, err := s.repo.FindCategoryRepo()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) FindAllProductService() ([]Item, error) {
	resp, err := s.repo.FindAllProductRepo()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) FindCategoryByCodeService(Code string) ([]Item, error) {
	resp, err := s.repo.FindCategoryByCodeRepo(Code)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) FindProductByIDService(ID int64) ([]ItemDetail, error) {
	resp, err := s.repo.FindProductByIDRepo(ID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) AddCategoryService(UserID string, req CategoryAdd) (interface{}, error) {
	resp, err := s.repo.AddCategoryRepo(UserID, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) FavoriteProductService() ([]Item, error) {
	resp, err := s.repo.FavoriteProductRepo()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) AddItemService(UserID string, req ItemAdd) (interface{}, error) {
	resp, err := s.repo.AddItemRepo(UserID, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) AddPackage(UserID string, req ItemAdd) (interface{}, error) {
	resp, err := s.repo.AddPackage(UserID, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) AddPromotion(UserID string, req ItemAdd) (interface{}, error) {
	resp, err := s.repo.AddPromotion(UserID, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) FindPackageService() ([]Item, error) {
	resp, err := s.repo.FindPackageRepo()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) FindYoutubeList() ([]YoutubeList, error) {
	resp, err := s.repo.FindYoutubeList()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) FindCampaign() ([]Campaign, error) {
	resp, err := s.repo.FindCampaign()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) FavoritePromotion() ([]Item, error) {
	resp, err := s.repo.FavoritePromotion()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) ItemList(ID int64) ([]ItemList, error) {
	resp, err := s.repo.ItemList(ID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) UnitList() ([]ItemUnit, error) {
	resp, err := s.repo.UnitList()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) ItemHistory(ID int64, UserID string) ([]ItemHistory, error) {
	resp, err := s.repo.ItemHistory(ID, UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) DeletePicture(ID int64, UserID string) (interface{}, error) {
	resp, err := s.repo.DeletePicture(ID, UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) DeleteProSub(ID int64, UserID string) (interface{}, error) {
	resp, err := s.repo.DeleteProSub(ID, UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
