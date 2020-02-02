package order

import (
	"context"
	"fmt"

	"gitlab.com/satit13/perfect_api/auth"
)

//"github.com/labstack/gommon/log"

func makeCartAddEndpoint(s Service) interface{} {
	type response struct {
		Response string `json:"response"`
		// Data     []Payment `json:"data"`
	}
	return func(ctx context.Context, req *Basket) (*response, error) {
		fmt.Println("enpoint req =", req)
		_, err := s.CartAddService(req)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success"}, nil
	}
}
func makeCartAddV2Endpoint(s Service) interface{} {
	type request struct {
		ItemID int64 `json:"item_id"`
		Qty    int64 `json:"qty"`
	}
	type response struct {
		Response interface{} `json:"response"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println("enpoint req =", req)
		resp, err := s.CartAddServiceV2(UserID, req.ItemID, req.Qty)
		if err != nil {
			return &response{Response: "error"}, err
		}
		a := map[string]interface{}{
			"item_id": req.ItemID,
			"qty":     req.Qty,
		}
		return &response{Response: resp, Data: a}, nil
	}
}

func makeCartStorageEndpoint(s Service) interface{} {
	type response struct {
		Response string   `json:"response"`
		Data     []Basket `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetUserID(ctx)
		resq, err := s.CartStorageService(UserID)
		if err != nil {
			return &response{Response: "error"}, err
		}
		return &response{Response: "success", Data: resq}, nil
	}
}

func makeCartToOrderEndpoint(s Service) interface{} {
	type request struct {
		ItemID         []int64 `json:"item_id"`
		DiscountAmount float64 `json:"discount_amount"`
		DeliveryCode   string  `json:"dilivery_code"`
		DistanceAmount int64   `json:"distance_amount"`
		Coupon         string  `json:"coupon"`
	}
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		resp, err := s.CartToOrderService(UserID, req.ItemID, req.DiscountAmount, req.DeliveryCode, req.DistanceAmount, req.Coupon)
		if err != nil {
			fmt.Println("CartToOrder error ", err.Error())
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeOrderAllEndpoint(s Service) interface{} {
	type response struct {
		Response string  `json:"response"`
		Data     []Order `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetUserID(ctx)
		resq, err := s.OrderAllService(UserID)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success", Data: resq}, nil
	}
}

func makeOrderByIDEndpoint(s Service) interface{} {
	type request struct {
		ID int64 `json:"id"`
	}
	type response struct {
		Response string  `json:"response"`
		Data     []Order `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println(UserID)
		resq, err := s.OrderByIDService(req.ID)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success", Data: resq}, nil
	}
}

func makeOrderCancelEndpoint(s Service) interface{} {
	type request struct {
		ID int64 `json:"id"`
	}
	type response struct {
		Response interface{} `json:"response"`
		// Data     []Payment `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		resq, err := s.OrderCancelService(UserID, req.ID)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: resq}, nil
	}
}

func makeCartDeleteEndpoint(s Service) interface{} {
	type response struct {
		Response string `json:"response"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetUserID(ctx)
		_, err := s.CartDeleteService(UserID)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success"}, nil
	}
}

func makeCartDeleteByItemEndpoint(s Service) interface{} {
	type request struct {
		ItemID int64 `json:"item_id"`
	}
	type response struct {
		Response string `json:"response"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		_, err := s.CartDeleteByItemService(req.ItemID, UserID)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success"}, nil
	}
}

func makeCartQTYEndpoint(s Service) interface{} {
	type request struct {
		ItemID int64 `json:"item_id"`
		QTY    int64 `json:"qty"`
	}
	type response struct {
		Response string `json:"response"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		_, err := s.CartQTYService(req.ItemID, req.QTY, UserID)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success"}, nil
	}
}

func makeOrderConfirmEndpoint(s Service) interface{} {
	type request struct {
		ID        int64  `json:"id"`
		URL       string `json:"url"`
		AddressID int64  `json:"address_id"`
	}
	type response struct {
		Response interface{} `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		resp, err := s.OrderConfirmService(UserID, req.ID, req.URL, req.AddressID)
		if err != nil {
			return &response{Response: "false"}, err
		}
		a := map[string]interface{}{
			"order_id":   req.ID,
			"url":        req.URL,
			"address_id": req.AddressID,
		}
		return &response{Response: resp, Message: "การ confirm เรียบร้อย", Data: a}, nil
	}
}

func makeBankEndpoint(s Service) interface{} {
	type response struct {
		Response string `json:"response"`
		Data     []Bank `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println(UserID)
		resp, err := s.BankService()
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makePromptPayEndpoint(s Service) interface{} {
	type request struct {
		// Phone  string  `json:"phone"`
		Amount float64 `json:"amount"`
	}
	type response struct {
		Response  string      `json:"response"`
		PromptPay interface{} `json:"promptpay"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println(UserID)
		resp, err := s.PromptPayService(req.Amount)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success", PromptPay: resp}, nil
	}
}

func makeUpdateStatusEndpoint(s Service) interface{} {
	type request struct {
		ID         int64  `json:"id"`
		Status     int64  `json:"status"`
		TrackingID string `json:"tracking_id"`
	}
	type response struct {
		Response interface{} `json:"response"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		// UserID := auth.GetUserID(ctx)
		UserID := auth.GetAdminID(ctx)
		fmt.Println(UserID)
		resp, err := s.UpdateStatusService(UserID, req.ID, req.Status, req.TrackingID)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: resp}, nil
	}

}

func makeDeliveryPriceEndpoint(s Service) interface{} {
	type request struct {
		ItemID []int64 `json:"item_id"`
		Code   string  `json:"code"`
	}
	type response struct {
		Response string  `json:"response"`
		Name     string  `json:"name"`
		Amount   float64 `json:"amount"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println(UserID)
		name, resp, err := s.DeliveryPriceService(UserID, req.ItemID, req.Code)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success", Name: name, Amount: resp}, nil
	}
}

func makeDeliveryAllEndpoint(s Service) interface{} {
	type response struct {
		Response string        `json:"response"`
		Data     []DeliveryAll `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println(UserID)
		resp, err := s.DeliveryAllService()
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeOrderListAllEndpoint(s Service) interface{} {
	type request struct {
		OrderStatus int64 `json:"order_status"`
	}
	type response struct {
		Response string  `json:"response"`
		Data     []Order `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println(UserID)
		resq, err := s.OrderListAllService(req.OrderStatus)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success", Data: resq}, nil
	}
}

func makeOrderListSendEndpoint(s Service) interface{} {
	type response struct {
		Response string      `json:"response"`
		Data     []OrderSend `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println(UserID)
		resq, err := s.OrderListSendService()
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success", Data: resq}, nil
	}
}

func makeSendDetailEndpoint(s Service) interface{} {
	type request struct {
		ID int64 `json:"id"`
	}
	type response struct {
		Response string            `json:"response"`
		Data     []OrderSendDetail `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println(UserID)
		resq, err := s.OrderSendDetailService(req.ID)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success", Data: resq}, nil
	}
}

func makeOrderPackageEndpoint(s Service) interface{} {
	type request struct {
		ItemID int64 `json:"item_id"`
	}
	type response struct {
		Response string      `json:"response"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		resp, err := s.OrderPackageService(UserID, req.ItemID)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeBankListEndpoint(s Service) interface{} {
	type response struct {
		Response string     `json:"response"`
		Message  string     `json:"message"`
		Data     []BankList `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		resp, err := s.BankList()
		if err != nil {
			fmt.Println("makeBankListEndpoint error ", err.Error())
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{Response: "success", Data: resp}, nil
	}
}

func makeStatusOrderEndpoint(s Service) interface{} {
	type request struct {
		DocNo string `json:"doc_no"`
	}
	type response struct {
		Response string      `json:"response"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		resp, err := s.StatusOrder(req.DocNo)
		if err != nil {
			return &response{Response: "false"}, err
		}
		return &response{Response: "success", Data: resp}, nil
	}
}
