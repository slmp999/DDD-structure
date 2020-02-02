package manage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/gommon/log"
)

type service struct {
	repo Repository
}

// NewService creates new auth service
func NewManage(manage Repository) (Service, error) {
	s := service{manage}
	return &s, nil
}

type Service interface {
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
