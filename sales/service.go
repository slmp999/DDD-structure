package sales

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
)

type service struct {
	repo Repository
}

// NewService creates new auth service
func NewSales(sales Repository) (Service, error) {
	s := service{sales}
	return &s, nil
}

type Service interface {
	RegisterSaleService(UserID string, sale *RegisterSalesModel) (interface{}, error)
	ConfirmSaleService(AdminID string, SaleCode string, UserID string) (interface{}, error)
	ShowlistConfirmSales(Type string, Limit int64) (interface{}, error)
	GetSalesCodeService(SaleCode string) (interface{}, error)
	GetSaleByUser(UserID string) (interface{}, error)
	UpdateSaleService(UserID string, sale *UpdateSaleModel) (interface{}, error)
	GetSalesTeamService(userID string) (interface{}, error)
	RemoveSales(AdminID string, SaleCode string, UserID string) (interface{}, error)
	GetSaleByUserV2(UserID string) (interface{}, error)
	UpdateSaleServiceAdmin(AdminID string, sale *UpdateSaleModelAdmin) (interface{}, error)
	ListCommisionService(userID string) (interface{}, error)
	GenarateComisionService(userID string) (interface{}, error)
	GetCommisionByDocNoService(userID string, docno string) (interface{}, error)
	CancelCommisionDocNoService(userID string, docno string) (interface{}, error)
	ListCommisonService(status int64, keyword string, limit int64) (int64, interface{}, error)
	ApproveCommision(adminID string, docno string, status int64, slipapprove string) (interface{}, error)
	GetCommisionByDocNo(docno string) (interface{}, error)
	CancelCommisionFontService(userID string, docno string) (interface{}, error)
}

func (s *service) CancelCommisionDocNoService(userID string, docno string) (interface{}, error) {
	resp, err := s.repo.CancelCommisionRepo(userID, docno)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (s *service) CancelCommisionFontService(userID string, docno string) (interface{}, error) {
	resp, err := s.repo.CancelCommisionFontRepo(userID, docno)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (s *service) ApproveCommision(adminID string, docno string, status int64, slipapprove string) (interface{}, error) {
	resp, err := s.repo.ApproveCommisionRepo(adminID, docno, status, slipapprove)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) GetCommisionByDocNo(docno string) (interface{}, error) {
	resp, err := s.repo.GetCommisionDocNo(docno)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) ListCommisonService(status int64, keyword string, limit int64) (int64, interface{}, error) {
	var search string = ""
	words := strings.Fields(keyword)
	if len(words) > 0 {
		likescode := ` a.doc_no like '%` + words[0] + `%'`
		likesname := ` a.user_id like '%` + words[0] + `%'`
		likesalec := ` a.sales_code like '%` + words[0] + `%'`
		likename := ` b.user_fname like '%` + words[0] + `%'`
		likelname := ` b.user_lname like '%` + words[0] + `%'`
		likephone := ` b.telephone like '%` + words[0] + `%'`
		for i := range words {
			if i > 0 {
				likescode += ` and a.doc_no like '%` + words[i] + `%'`
				likesname += ` and a.user_id like '%` + words[i] + `%'`
				likesalec += ` and a.sales_code like '%` + words[i] + `%'`
				likename += ` and b.user_fname like '%` + words[i] + `%'`
				likelname += ` and b.user_lname like '%` + words[i] + `%'`
				likephone += ` and b.telephone like '%` + words[i] + `%'`
			}
		}
		search = likescode + ` or ` + likesname + ` or ` + likesalec + ` or ` + likename + ` or ` + likelname + ` or ` + likephone
	} else {
		likescode := ` a.doc_no like '%` + keyword + `%'`
		likesname := ` a.user_id like '%` + keyword + `%'`
		likesalec := ` a.sales_code like '%` + keyword + `%'`
		likename := ` b.user_fname like '%` + keyword + `%'`
		likelname := ` b.user_lname like '%` + keyword + `%'`
		likephone := ` b.telephone like '%` + keyword + `%'`
		search = likescode + ` or ` + likesname + ` or ` + likesalec + ` or ` + likename + ` or ` + likelname + ` or ` + likephone
	}

	resp, err := s.repo.ListCommsionBackendRepo(status, search, limit)
	if err != nil {
		return 0, nil, err
	}
	len, err := s.repo.CountListCommmisonBackEnd(status, search)
	if err != nil {
		return 0, nil, err
	}

	return len, resp, nil
}

func (s *service) GetCommisionByDocNoService(userID string, docno string) (interface{}, error) {
	resp, err := s.repo.GetCommisionDocNoRepo(userID, docno)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) RemoveSales(AdminID string, SaleCode string, UserID string) (interface{}, error) {
	resp, err := s.repo.RemoveSalesRepo(AdminID, SaleCode, UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) GetSalesTeamService(userID string) (interface{}, error) {
	resp, err := s.repo.GetSalesTeamRepo(userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) GenarateComisionService(userID string) (interface{}, error) {
	resp, err := s.repo.GenarateCommisionRepo(userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) UpdateSaleServiceAdmin(AdminID string, sale *UpdateSaleModelAdmin) (interface{}, error) {
	resp, err := s.repo.UpdateSalebyAdminRepo(AdminID, sale)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (s *service) ListCommisionService(userID string) (interface{}, error) {
	resp, err := s.repo.ListCommsionRepo(userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) UpdateSaleService(UserID string, sale *UpdateSaleModel) (interface{}, error) {
	resp, err := s.repo.UpdateSaleRepo(UserID, sale)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) RegisterSaleService(UserID string, sale *RegisterSalesModel) (interface{}, error) {
	resp, err := s.repo.RegisterSaleRepo(UserID, sale)
	if err != nil {
		return nil, err
	}
	loc, _ := time.LoadLocation("Asia/Bangkok")
	message := fmt.Sprintf(
		"Sale Register \nชื่อ %v %v\nวันที่ %v เวลา %v",
		sale.Fname,
		sale.Lname,
		time.Now().In(loc).Format("2006-01-02"),
		time.Now().In(loc).Format("15:04:05"),
	)
	go s.sendLine(message, 4)

	return resp, nil
}

func (s *service) GetSaleByUser(UserID string) (interface{}, error) {
	resp, err := s.repo.GetProfileSalesbyUser(UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) GetSaleByUserV2(UserID string) (interface{}, error) {
	resp, err := s.repo.GetProfileSalesbyUserV2(UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (s *service) GetSalesCodeService(SaleCode string) (interface{}, error) {
	resp, err := s.repo.GetProfileSales(SaleCode)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) ConfirmSaleService(AdminID string, SaleCode string, UserID string) (interface{}, error) {
	resp, err := s.repo.SaleConfirmRepo(AdminID, SaleCode, UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) ShowlistConfirmSales(Type string, Limit int64) (interface{}, error) {
	if Type == "NOCONFIRMED" {
		resp, err := s.repo.ListSaleNoConfirm(Limit)

		if err != nil {
			return nil, err
		}

		return resp, nil
	} else if Type == "CONFIRMED" {
		resp, err := s.repo.ListSaleConfirmed(Limit)
		if err != nil {
			return nil, err
		}

		return resp, nil
	} else {
		resp, err := s.repo.ListSaleALL(Limit)

		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, nil
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
