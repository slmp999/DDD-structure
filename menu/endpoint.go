package menu

import "context"

// import (
// 	"context"

// 	"gitlab.com/satit13/perfect_api/auth"
// )

func makeHeadMenuEndpoint(s Service) interface{} {
	type request struct {
		HeadName string `json:"head_menu_name"`
	}
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {

		resp, err := s.ShowlistConfirmSales(10)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     resp,
		}, nil
	}
}

// func makeSubmenuEndpoint(s Service) interface{} {
// 	type request struct {
// 		HeadMemuID int64  `json:"head_menu_id"`
// 		SubName    string `json:"sub_menu_name"`
// 	}
// 	type response struct {
// 		Response string      `json:"response"`
// 		Message  string      `json:"message"`
// 		Data     interface{} `json:"data"`
// 	}
// 	return func(ctx context.Context, req *request) (*response, error) {

// 		resp, err := s.AddSubMenuService(req.HeadMemuID, req.SubName)
// 		if err != nil {
// 			return &response{Response: "false", Message: err.Error()}, nil
// 		}
// 		return &response{
// 			Response: "success",
// 			Data:     resp,
// 		}, nil
// 	}
// }

// func makeListMenuEnpoint(s Service) interface{} {

// 	type response struct {
// 		Response string      `json:"response"`
// 		Message  string      `json:"message"`
// 		Data     interface{} `json:"data"`
// 	}
// 	return func(ctx context.Context) (*response, error) {

// 		resp, err := s.ConfirmSaleService(AdminID, req.SaleCode, req.UserID)
// 		if err != nil {
// 			return &response{Response: "false", Message: err.Error()}, nil
// 		}
// 		return &response{
// 			Response: "success",
// 			Data:     resp,
// 		}, nil
// 	}
// }

// func makeListHeadMenuEndpoint(s Service) interface{} {
// 	type request struct {
// 		Limit int64 `json:"limit"`
// 	}
// 	type response struct {
// 		Response string      `json:"response"`
// 		Message  string      `json:"message"`
// 		Data     interface{} `json:"data"`
// 	}
// 	return func(ctx context.Context, req *request) (*response, error) {
// 		resp, err := s.ShowlistConfirmSales(req.Limit)
// 		if err != nil {
// 			return &response{Response: "false", Message: err.Error()}, nil
// 		}
// 		return &response{
// 			Response: "success",
// 			Data:     resp,
// 		}, nil
// 	}
// }
