package product

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
	// search
	mux.Handle("/item", m.Handler(makeFindAllProductEndpoint(s)))
	mux.Handle("/item/id", m.Handler(makeFindProductByIDEndpoint(s)))
	mux.Handle("/item/favorite", m.Handler(makeFavoriteProductEndpoint(s)))
	mux.Handle("/item/category", m.Handler(makeFindCategoryByCodeEndpoint(s)))
	mux.Handle("/category", m.Handler(makeFindCategoryEndpoint(s)))
	mux.Handle("/package", m.Handler(makeFindPackageEndpoint(s)))
	mux.Handle("/youtube", m.Handler(makeFindYoutubeListEndpoint(s)))
	mux.Handle("/campaign", m.Handler(makeFindCampaignEndpoint(s)))
	mux.Handle("/promotion/favorite", m.Handler(makeFavoritePromotionEndpoint(s)))
	mux.Handle("/item/list", m.Handler(makeItemListEndpoint(s)))
	mux.Handle("/item/unitlist", m.Handler(makeUnitListEndpoint(s)))
	mux.Handle("/item/history", m.Handler(makeItemHistoryEndpoint(s)))
	// add/delete/update
	mux.Handle("/add/category", m.Handler(makeAddCategoryEndpoint(s)))
	mux.Handle("/add/item", m.Handler(makeAddItemEndpoint(s)))
	mux.Handle("/add/promotion", m.Handler(makeAddPromotionEndpoint(s)))
	mux.Handle("/add/package", m.Handler(makeAddPackageEndpoint(s)))
	mux.Handle("/delete/picture/item", m.Handler(makeDeletePictureEndpoint(s)))
	mux.Handle("/delete/promo/item", m.Handler(makeDeleteProSubEndpoint(s)))
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
