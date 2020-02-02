package report

import (
	"context"

	log "gitlab.com/satit13/perfect_api/logger"
)

func MakereportSale(s Service) interface{} {
	type request struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}
	type response struct {
		Response string `json:"response"`
		Message  string `json:"message"`

		Data interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpooint MakeEndpointApprove :", req)
		resp, err := s.ListReportSale(req.StartDate, req.EndDate)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Message:  "",
			Data:     resp,
		}, nil
	}
}
