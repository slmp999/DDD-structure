package member

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/acoshift/hrpc"
	"gitlab.com/satit13/perfect_api/auth"
)

var (
	errMethodNotAllowed = errors.New("auth: method not allowed")
	errForbidden        = errors.New("auth: forbidden")
	errBadRequest       = errors.New("auth: bad request body")
	errUnauthorized     = errors.New("auth: Unauthorized")
)

type errorResponse struct {
	Error string `json:"error"`
}

// MakeMiddleware creates new auth middleware

// MakeHandler creates new auth handler
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
	mux.Handle("/get/user", m.Handler(FindProfileMemberEndpoint(s)))
	mux.Handle("/get/coupon", m.Handler(FindCouponMemberEndpoint(s)))
	mux.Handle("/check/coupon", m.Handler(CheckCouponByNoEndpoint(s)))
	mux.Handle("/update/profile", m.Handler(UpdateProfileMemberEndpoint(s)))
	mux.Handle("/update/member/admin", m.Handler(UpdateMemberBackEndpoint(s)))
	mux.Handle("/get/address", m.Handler(FindAddressMemberEndpoint(s)))
	mux.Handle("/get/address/id", m.Handler(FindAddressMemberByIdEndpoint(s)))
	mux.Handle("/add/address", m.Handler(makeAddProfileAddressEndpoint(s)))
	mux.Handle("/update/address", m.Handler(makeUpdateAddressEndpoint(s)))
	mux.Handle("/remove/user/admin", m.Handler(makeRemoveUserByAdmin(s)))
	mux.Handle("/get/user/admin", m.Handler(getUserByAdmin(s)))
	mux.Handle("/remove/address", m.Handler(makeDeleteProfileEndpoint(s)))
	mux.Handle("/upload/image", m.Handler(makeUploadimage(s)))
	mux.Handle("/get/users", m.Handler(FindAllMember(s)))

	return mustLogin(s)(mux)
}

// MakeAccountHandler creates new account handler
// func MakeAccountHandler(s Service) http.Handler {
// 	m := hrpc.Manager{
// 		Validate:     true,
// 		Decoder:      requestDecoder,
// 		Encoder:      responseEncoder,
// 		ErrorEncoder: errorEncoder,
// 	}
// 	mux := http.NewServeMux()
// 	//mux.Handle("/profile", m.Handler(makeGetProfileEndpoint(s)))
// 	// mux.Handle("/change-password", m.Handler(makeGetProfileEndpoint(s)))
// 	// mux.Handle("/hash", m.Handler(makeGetProfileEndpoint(s)))
// 	return mustLogin(s)(mux)
// }

func mustLogin(s Service) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			UserID := auth.GetUserID(r.Context())
			AdminID := auth.GetAdminID(r.Context())

			if UserID == "" && AdminID == "" {
				if !CheckUrlstring(r) {
					errorEncoder(w, r, errForbidden)
					fmt.Println("error mustLogin auth.transport.go")
					return
				}
			}
			enableCors(&w)
			h.ServeHTTP(w, r)
		})
	}
}

func CheckUrlstring(r *http.Request) bool {

	if !strings.Contains(r.URL.String(), "signup") || !strings.Contains(r.URL.String(), "signin") {
		return false
	}

	return true
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set(
		"Access-Control-Allow-Headers",
		"Accept, Access-Token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, access-control-allow-origin",
	)
}

func jsonDecoder(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		log.Printf("[Member] API request: %+v\n", v)
		log.Printf("[Member] error: %+v\n", err)
		return errBadRequest
	}
	return nil
}

func jsonEncoder(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Token,access-token,access-control-allow-origin")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(status)
	log.Printf("[Member] API response %d: %+v\n", status, v)
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
	log.Println(r.Body)
	switch err {
	case errMethodNotAllowed:
		status = http.StatusMethodNotAllowed
	case errForbidden:
		status = http.StatusForbidden
	case errBadRequest:
		status = http.StatusBadRequest
	case errUnauthorized:
		status = http.StatusUnauthorized
		// case ErrTokenExpired:
		// 	status = http.StatusUnauthorized
		// case ErrTokenNotFound:
		// 	status = http.StatusUnauthorized
	}
	if r.Method == http.MethodOptions {
		encoder(w, http.StatusNoContent, nil)
	} else {
		encoder(w, status, &errorResponse{err.Error()})
	}
}
