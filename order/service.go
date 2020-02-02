package order

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	pp "github.com/Frontware/promptpay"
	"github.com/labstack/gommon/log"
)

type Service interface {
	CartAddService(req *Basket) (interface{}, error)
	CartAddServiceV2(UserID string, ItemID int64, Qty int64) (interface{}, error)
	CartStorageService(UserID string) ([]Basket, error)
	CartDeleteService(UserID string) (interface{}, error)
	CartToOrderService(UserID string, ItemID []int64, DiscountAmount float64, DeliveryCode string, DistanceAmount int64, Coupon string) (interface{}, error)
	OrderAllService(UserID string) ([]Order, error)
	OrderConfirmService(UserID string, ID int64, URL string, AddressID int64) (interface{}, error)
	CartDeleteByItemService(ItemID int64, UserID string) (interface{}, error)
	CartQTYService(ItemID int64, QTY int64, UserID string) (interface{}, error)
	OrderCancelService(UserID string, ID int64) (interface{}, error)
	OrderByIDService(ID int64) ([]Order, error)
	BankService() ([]Bank, error)
	PromptPayService(Amount float64) (interface{}, error)
	UpdateStatusService(UserID string, ID int64, Status int64, TrackingID string) (interface{}, error)
	DeliveryPriceService(UserID string, ItemID []int64, Code string) (name string, resp float64, err error)
	DeliveryAllService() ([]DeliveryAll, error)
	OrderListAllService(OrderStatus int64) ([]Order, error)
	OrderListSendService() ([]OrderSend, error)
	OrderSendDetailService(ID int64) ([]OrderSendDetail, error)
	OrderPackageService(UserID string, ItemID int64) (interface{}, error)
	BankList() ([]BankList, error)
	TrackingThaiPost(UserID string, TrackingID string) (interface{}, error)
	StatusOrder(DocNo string) (interface{}, error)
}

// NewService creates new device service
func NewService(order Repository) (Service, error) {
	s := service{order}
	return &s, nil
}

type service struct {
	repo Repository
}

func (s *service) CartAddService(req *Basket) (interface{}, error) {
	resp, err := s.repo.CartAddRepo(req)
	println(resp)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) CartAddServiceV2(UserID string, ItemID int64, Qty int64) (interface{}, error) {
	resp, err := s.repo.CartAddRepoV2(UserID, ItemID, Qty)
	println(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) CartDeleteService(UserID string) (interface{}, error) {
	resp, err := s.repo.CartDeleteRepo(UserID)
	println(resp)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) CartStorageService(UserID string) ([]Basket, error) {
	resp, err := s.repo.CartStorageRepo(UserID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *service) CartToOrderService(UserID string, ItemID []int64, DiscountAmount float64, DeliveryCode string, DistanceAmount int64, Coupon string) (interface{}, error) {
	if len(ItemID) == 0 {
		return nil, fmt.Errorf("ไม่มีข้อมูลสินค้า")
	}
	resp, err := s.repo.CartToOrderRepo(UserID, ItemID, DiscountAmount, DeliveryCode, DistanceAmount, Coupon)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) OrderAllService(UserID string) ([]Order, error) {
	resp, err := s.repo.OrderAllRepo(UserID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *service) OrderByIDService(ID int64) ([]Order, error) {
	resp, err := s.repo.OrderByIDRepo(ID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *service) OrderConfirmService(UserID string, ID int64, URL string, AddressID int64) (interface{}, error) {
	resp, Fname, Lname, OrAmount, DocNo, err := s.repo.OrderConfirmRepo(UserID, ID, URL, AddressID)
	if err != nil {
		return nil, err
	}
	//init the loc
	loc, _ := time.LoadLocation("Asia/Bangkok")
	// now := time.Now().In(loc)

	if resp == "success" {
		message := fmt.Sprintf(
			"Confirm Order %v \nราคา %v\nชื่อ %v %v\nวันที่ %v เวลา %v",
			DocNo,
			OrAmount,
			Fname,
			Lname,
			time.Now().In(loc).Format("2006-01-02"),
			time.Now().In(loc).Format("15:04:05"),
		)
		go s.sendLine(message, 1)
	}
	return resp, nil
}

func (s *service) CartDeleteByItemService(ItemID int64, UserID string) (interface{}, error) {
	resp, err := s.repo.CartDeleteByItemRepo(ItemID, UserID)
	println(resp)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) CartQTYService(ItemID int64, QTY int64, UserID string) (interface{}, error) {
	resp, err := s.repo.CartQTYRepo(ItemID, QTY, UserID)
	println(resp)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) OrderCancelService(UserID string, ID int64) (interface{}, error) {
	resq, err := s.repo.OrderCancelRepo(UserID, ID)
	if err != nil {
		return nil, err
	}
	return resq, nil
}

func (s *service) BankService() ([]Bank, error) {
	resq, err := s.repo.BankRepo()
	if err != nil {
		return nil, err
	}
	return resq, nil
}

func (s *service) PromptPayService(Amount float64) (interface{}, error) {
	// if Phone[0:1] == "0" {
	// 	Phone = "66" + Phone[1:len(Phone)]
	// }

	payment := pp.PromptPay{
		PromptPayID: "66939146474", // Tax-ID/ID Card/E-Wallet
		Amount:      Amount,        // Positive amount
	}

	qrcode, _ := payment.Gen() // Generate string to be use in QRCode
	fmt.Println(qrcode)        // Print string
	return qrcode, nil
}

func (s *service) UpdateStatusService(UserID string, ID int64, Status int64, TrackingID string) (interface{}, error) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	// now := time.Now().In(loc)

	resq, err := s.repo.UpdateStatusRepo(UserID, ID, Status, TrackingID)
	if err != nil {
		return nil, err
	}
	if resq == "success" && Status == 2 {
		message := fmt.Sprintf(
			"Pack Order %d สถานะ %d\nวันที่ %v เวลา %v",
			ID,
			Status,
			time.Now().In(loc).Format("2006-01-02"),
			time.Now().In(loc).Format("15:04:05"),
		)
		go s.sendLine(message, 2)
	}
	return resq, nil
}

func (s *service) DeliveryPriceService(UserID string, ItemID []int64, Code string) (name string, resp float64, err error) {
	name, resq, err := s.repo.DeliveryPriceRepo(UserID, ItemID, Code)
	if err != nil {
		return "", resp, err
	}
	return name, resq, nil
}

func (s *service) DeliveryAllService() ([]DeliveryAll, error) {
	resp, err := s.repo.DeliveryAllRepo()
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *service) sendLine(message string, ID int64) (string, error) {
	clientToken, err := s.repo.CallTokenLine(ID)
	client := &http.Client{}
	url := "https://notify-api.line.me/api/notify"
	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"message\"\r\n\r\n" + message + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")
	cReq, err := http.NewRequest("POST", url, payload)
	cReq.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	// cReq.Header.Add("authorization", "Bearer djp2CpuGYUb8slBPE5YFxMQ4s1u5u4DU0Fpoe25gbZS")
	cReq.Header.Add("authorization", "Bearer "+clientToken)
	cReq.Header.Add("cache-control", "no-cache")
	response, err := client.Do(cReq)
	if err != nil {
		log.Error("send line notify error : ", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("send line notify error : ", err)
	}
	var returnBody struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	err = json.Unmarshal(body, &returnBody)
	if err != nil {
		log.Error("send line notify error : ", err)
	}
	if returnBody.Status != 200 {
		log.Error("send line notify error : ", returnBody)
	}
	return returnBody.Message, nil
}

func (s *service) OrderListAllService(OrderStatus int64) ([]Order, error) {
	resp, err := s.repo.OrderListAllRepo(OrderStatus)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *service) OrderListSendService() ([]OrderSend, error) {
	resp, err := s.repo.OrderListSendAllRepo()
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *service) OrderSendDetailService(ID int64) ([]OrderSendDetail, error) {
	resp, err := s.repo.OrderSendDetailRepo(ID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *service) OrderPackageService(UserID string, ItemID int64) (interface{}, error) {
	resp, err := s.repo.OrderPackageRepo(UserID, ItemID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) BankList() ([]BankList, error) {
	resq, err := s.repo.BankList()
	if err != nil {
		return nil, err
	}
	return resq, nil
}

func (s *service) TrackingThaiPost(UserID string, TrackingID string) (interface{}, error) {

	type TrackingIDRequest struct {
		TrackingID []ItemsPost `json:"EF582568151TH"`
	}
	type Responses struct {
		Items      TrackingIDRequest `json:"items"`
		TrackCount TrackCount        `json:"track_count"`
	}
	type PostResponse struct {
		Response Responses `json:"response"`
		Message  string    `json:"message"`
		Status   bool      `json:"status"`
	}

	jsonData := map[string]interface{}{
		"status":   "all",
		"language": "TH",
		"barcode":  []string{TrackingID},
	}
	fmt.Println("jsonData :", jsonData)
	jsonValue, _ := json.Marshal(jsonData)
	client := &http.Client{}
	respo, err := http.NewRequest("POST", "https://trackapi.thailandpost.co.th/post/api/v1/track", bytes.NewReader(jsonValue))
	respo.Header.Add("Authorization", "Token eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJzZWN1cmUtYXBpIiwiYXVkIjoic2VjdXJlLWFwcCIsInN1YiI6IkF1dGhvcml6YXRpb24iLCJleHAiOjE1Nzg4MTM5MTIsInJvbCI6WyJST0xFX1VTRVIiXSwiZCpzaWciOnsicCI6InpXNzB4IiwicyI6bnVsbCwidSI6ImM0YjIyZmE2ZDViNzQ1NWViZTYwMjhlNGIzY2M5ZjQ4IiwiZiI6InhzeiM5In19.FTZVufNLPQOVdyKq8qmIh6jSMLF3Glhw_B7Os5lvUf-ioFbapr-z1-47mtq1EpHFw2Dzut7iyJm9fZRccOFdHQ")
	respo.Header.Add("Content-Type", "application/json")
	response, err := client.Do(respo)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	var respons PostResponse
	err = json.Unmarshal(data, &respons)
	if err != nil {
		return nil, err
	}
	if respons.Status == true {
		if respons.Message != "successful" {
			fmt.Println("case no successful")
			respnData := PostFail{
				Message: respons.Message,
				Status:  respons.Status,
			}
			return &respnData, nil
		} else {
			fmt.Println("successful")
			respnData := PostResponse{
				Response: respons.Response,
				Message:  respons.Message,
				Status:   respons.Status,
			}
			return &respnData, nil
		}
	} else {
		return nil, errors.New("unknown")
	}
	return nil, nil
}

func (s *service) StatusOrder(DocNo string) (interface{}, error) {
	resp, err := s.repo.StatusOrder(DocNo)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
