package product

import (
	"context"
	"fmt"

	"gitlab.com/satit13/perfect_api/auth"
)

//"github.com/labstack/gommon/log"

func makeFindAllProductEndpoint(s Service) interface{} {
	// type request struct {
	// 	Limit int `json:"limit"`
	// }
	type response struct {
		Response string `json:"response"`
		Data     []Item `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		resp, err := s.FindAllProductService()
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeFindCategoryEndpoint(s Service) interface{} {
	// type request struct {
	// 	ID int64 `json:"id"`
	// }
	type response struct {
		Response string     `json:"response"`
		Data     []Category `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		resp, err := s.FindCategoryService()
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeFindCategoryByCodeEndpoint(s Service) interface{} {
	type request struct {
		Code string `json:"code"`
	}
	type response struct {
		Response string `json:"response"`
		Data     []Item `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		resp, err := s.FindCategoryByCodeService(req.Code)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeFindProductByIDEndpoint(s Service) interface{} {
	type request struct {
		ID int64 `json:"id"`
	}
	type response struct {
		Response string       `json:"response"`
		Data     []ItemDetail `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		resp, err := s.FindProductByIDService(req.ID)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeAddCategoryEndpoint(s Service) interface{} {
	type request struct {
		CategoryAdd
	}
	type response struct {
		Response string `json:"response"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		_, err := s.AddCategoryService(UserID, req.CategoryAdd)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success"}, nil
	}
}

func makeFavoriteProductEndpoint(s Service) interface{} {
	type response struct {
		Response string `json:"response"`
		Data     []Item `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		resp, err := s.FavoriteProductService()
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeAddItemEndpoint(s Service) interface{} {
	type request struct {
		ItemAdd
	}
	type response struct {
		Response string `json:"response"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		_, err := s.AddItemService(UserID, req.ItemAdd)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success"}, nil
	}
}

func makeFindPackageEndpoint(s Service) interface{} {
	// type request struct {
	// 	Limit int `json:"limit"`
	// }
	type response struct {
		Response string `json:"response"`
		Data     []Item `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		resp, err := s.FindPackageService()
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeFindYoutubeListEndpoint(s Service) interface{} {
	type response struct {
		Response string        `json:"response"`
		Data     []YoutubeList `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		resp, err := s.FindYoutubeList()
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeFindCampaignEndpoint(s Service) interface{} {
	type response struct {
		Response string     `json:"response"`
		Data     []Campaign `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		resp, err := s.FindCampaign()
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeFavoritePromotionEndpoint(s Service) interface{} {
	type response struct {
		Response string `json:"response"`
		Data     []Item `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		resp, err := s.FavoritePromotion()
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeItemListEndpoint(s Service) interface{} {
	type request struct {
		ID int64 `json:"ID"`
	}
	type response struct {
		Response string     `json:"response"`
		Data     []ItemList `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println("UserID:", UserID)
		resp, err := s.ItemList(req.ID)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeAddPackageEndpoint(s Service) interface{} {
	type request struct {
		ItemAdd
	}
	type response struct {
		Response string `json:"response"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println("UserID:", UserID)
		_, err := s.AddPackage(UserID, req.ItemAdd)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success"}, nil
	}
}

func makeAddPromotionEndpoint(s Service) interface{} {
	type request struct {
		ItemAdd
	}
	type response struct {
		Response string `json:"response"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println("UserID:", UserID)
		_, err := s.AddPromotion(UserID, req.ItemAdd)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success"}, nil
	}
}

func makeUnitListEndpoint(s Service) interface{} {
	type response struct {
		Response string     `json:"response"`
		Data     []ItemUnit `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		resp, err := s.UnitList()
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeItemHistoryEndpoint(s Service) interface{} {
	type request struct {
		ID int64 `json:"ID"`
	}
	type response struct {
		Response string        `json:"response"`
		Data     []ItemHistory `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		resp, err := s.ItemHistory(req.ID, UserID)
		fmt.Println("resp = ", resp)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeDeletePictureEndpoint(s Service) interface{} {
	type request struct {
		ID int64 `json:"ID"`
	}
	type response struct {
		Response string `json:"response"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetAdminID(ctx)
		_, err := s.DeletePicture(req.ID, UserID)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success"}, nil
	}
}

func makeDeleteProSubEndpoint(s Service) interface{} {
	type request struct {
		ID int64 `json:"ID"`
	}
	type response struct {
		Response string `json:"response"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetAdminID(ctx)
		_, err := s.DeleteProSub(req.ID, UserID)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success"}, nil
	}
}
