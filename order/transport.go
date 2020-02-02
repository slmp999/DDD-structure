package order

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/acoshift/hrpc"
)

var (
	errMethodNotAllowed = errors.New("auth: method not allowed")
	errForbidden        = errors.New("auth: forbidden")
	errBadRequest       = errors.New("auth: bad request body")
)

type errorResponse struct {
	Error string `json:"error"`
}

// MakeMiddleware creates new auth middleware
// func MakeMiddleware(s Service) func(http.Handler) http.Handler {
// 	return func(h http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			// get token from header
// 			tokenID := r.Header.Get("X-Access-Token")

// 			fmt.Println("auth.transport token : ", tokenID)
// 			if len(tokenID) == 0 {
// 				h.ServeHTTP(w, r)
// 				return
// 			}

// 			tk, err := s.GetToken(tokenID)
// 			if err != nil {
// 				h.ServeHTTP(w, r)
// 				return
// 			}

// 			ctx := r.Context()
// 			ctx = context.WithValue(ctx, keyToken{}, tk)
// 			r = r.WithContext(ctx)
// 			h.ServeHTTP(w, r)
// 		})
// 	}
// }

// MakeHandler creates new device handler
func MakeHandler(s Service) http.Handler {
	//m := hrpc.New(hrpc.Config{
	//	Validate:        true,
	//	RequestDecoder:  requestDecoder,
	//	ResponseEncoder: responseEncoder,
	//	ErrorEncoder:    errorEncoder,
	//})

	m := hrpc.Manager{
		Validate:     true,
		Decoder:      requestDecoder,
		Encoder:      responseEncoder,
		ErrorEncoder: errorEncoder,
	}
	mux := http.NewServeMux()
	mux.Handle("/bank", m.Handler(makeBankEndpoint(s)))
	mux.Handle("/bank/list", m.Handler(makeBankListEndpoint(s)))
	mux.Handle("/cart/addbk", m.Handler(makeCartAddEndpoint(s)))
	mux.Handle("/cart/add", m.Handler(makeCartAddV2Endpoint(s)))
	mux.Handle("/cart/all/delete", m.Handler(makeCartDeleteEndpoint(s)))
	mux.Handle("/cart/item/delete", m.Handler(makeCartDeleteByItemEndpoint(s)))
	mux.Handle("/cart/update/qty", m.Handler(makeCartQTYEndpoint(s)))
	mux.Handle("/cart/storage", m.Handler(makeCartStorageEndpoint(s)))
	mux.Handle("/cart/transfer", m.Handler(makeCartToOrderEndpoint(s)))
	mux.Handle("/order/all", m.Handler(makeOrderAllEndpoint(s)))
	mux.Handle("/order/id", m.Handler(makeOrderByIDEndpoint(s)))
	mux.Handle("/order/confirm", m.Handler(makeOrderConfirmEndpoint(s)))
	mux.Handle("/order/cancel", m.Handler(makeOrderCancelEndpoint(s)))
	mux.Handle("/promptpay", m.Handler(makePromptPayEndpoint(s)))
	mux.Handle("/status/update", m.Handler(makeUpdateStatusEndpoint(s)))
	mux.Handle("/delivery/price", m.Handler(makeDeliveryPriceEndpoint(s)))
	mux.Handle("/delivery/all", m.Handler(makeDeliveryAllEndpoint(s)))
	mux.Handle("/buy/package", m.Handler(makeOrderPackageEndpoint(s)))
	//backend
	mux.Handle("/list/all", m.Handler(makeOrderListAllEndpoint(s)))
	mux.Handle("/list/send", m.Handler(makeOrderListSendEndpoint(s)))
	mux.Handle("/send/detail", m.Handler(makeSendDetailEndpoint(s)))
	//ckecktrack
	mux.Handle("/link/tracking", m.Handler(makeStatusOrderEndpoint(s)))

	return mustLogin()(mux)
}

func mustLogin() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			h.ServeHTTP(w, r)
		})
	}
}

func jsonDecoder(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		log.Printf("[Product] API request: %+v\n", v)
		return errBadRequest
	}
	return nil
}

func jsonEncoder(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Token,x-access-token")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(status)
	log.Printf("[Product] API response %d: %+v\n", status, v)
	return json.NewEncoder(w).Encode(v)
}

func requestDecoder(r *http.Request, v interface{}) error {
	if r.Method != http.MethodPost {
		return errMethodNotAllowed
	}
	return jsonDecoder(r, v)
}

func responseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) {
	jsonEncoder(w, http.StatusOK, v)
}

func errorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	encoder := jsonEncoder
	status := http.StatusInternalServerError
	switch err {
	case errMethodNotAllowed:
		if r.Method == http.MethodOptions {
			status = http.StatusOK
		} else {
			status = http.StatusMethodNotAllowed
		}
	case errForbidden:
		status = http.StatusForbidden
	case errBadRequest:
		status = http.StatusBadRequest
	}
	encoder(w, status, &errorResponse{err.Error()})
}
