package auth

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	logg "gitlab.com/satit13/perfect_api/logger"

	"github.com/acoshift/hrpc"
)

var (
	errMethodNotAllowed = errors.New("auth: method not allowed")
	errForbidden        = errors.New("auth: forbidden")
	errBadRequest       = errors.New("auth: bad request body")
	errUnauthorized     = errors.New("auth: Unauthorized")
)

type errorResponse struct {
	Response string `json:"response"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}

// MakeMiddleware creates new auth middleware
func MakeMiddleware(s Service) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get token from header
			tokenID := r.Header.Get("Access-Token")

			log.Println("auth.transport token : ", tokenID)
			if len(tokenID) == 0 {
				h.ServeHTTP(w, r)
				return
			}

			//fmt.Println("s.GetToken")

			tk, err := s.GetToken(tokenID)
			logg.Println(tk)
			if err != nil {
				// h.ServeHTTP(w, r)
				log.Println(err.Error())
				if !strings.Contains(r.URL.String(), "signup") || !strings.Contains(r.URL.String(), "signin") || !strings.Contains(r.URL.String(), "forget") || !strings.Contains(r.URL.String(), "reset") {
					errorEncoder(w, r, err)
					return
				}
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, keyToken{}, tk)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

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

	mux.Handle("/signup", m.Handler(makeSignUpEndpoint(s)))
	mux.Handle("/signin", m.Handler(makeSigninEndpoint(s)))

	mux.Handle("/get/user", m.Handler(makeGetProfile(s)))

	mux.Handle("/profile/address/add", m.Handler(makeAddProfileAddressEndpoint(s)))
	mux.Handle("/profile/address/id", m.Handler(makeGetProfileAddressEndpoint(s)))
	mux.Handle("/profile/address/list", m.Handler(makeListProfileEndpoint(s)))

	mux.Handle("/update/profileaddress", m.Handler(makeUpdateProfileEndpoint(s)))
	mux.Handle("/delete/profileaddress", m.Handler(makeDeleteProfileEndpoint(s)))

	mux.Handle("/forget/password", m.Handler(makeForgetPasswordEndpoint(s)))

	mux.Handle("/otp/request", m.Handler(makeOtpRequestEndpoint(s)))

	mux.Handle("/otp/validate", m.Handler(makeOtpValidateEndpoint(s)))

	mux.Handle("/reset/password", m.Handler(makeResetPasswordEndpoint(s)))
	//mux.Handle("/signin/facebook", m.Handler(makeSigninEndpoint(s)))
	//mux.Handle("/signin/gmail", m.Handler(makeSigninEndpoint(s)))

	// mux.Handle("/signinv2", m.Handler(makeSigninV2Endpoint(s)))
	mux.Handle("/signout", m.Handler(makeSignoutEndpoint(s)))

	mux.Handle("/admin/user", m.Handler(makeGetProfile(s)))
	mux.Handle("/change/password", m.Handler(makeChangePassword(s)))

	return mux
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
			// clientID := GetClientID(r.Context())
			// if clientID < 0 {
			// 	errorEncoder(w, r, errForbidden)
			// 	fmt.Println("error mustLogin auth.transport.go")
			// 	return
			// }
			enableCors(&w)
			h.ServeHTTP(w, r)
		})
	}
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
		log.Printf("[Sale] API request: %+v\n", v)
		log.Printf("[Sale] error: %+v\n", err)
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
	log.Printf("[Auth] API response %d: %+v\n", status, v)
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
		status = http.StatusMethodNotAllowed
	case errForbidden:
		status = http.StatusForbidden
	case errBadRequest:
		status = http.StatusBadRequest
	case errUnauthorized:
		status = http.StatusUnauthorized
	case ErrTokenExpired:
		status = http.StatusUnauthorized
	case ErrTokenNotFound:
		status = http.StatusUnauthorized
	}
	if r.Method == http.MethodOptions {
		encoder(w, http.StatusNoContent, nil)
	} else {
		encoder(w, status, &errorResponse{Response: "false", Message: err.Error(), Error: err.Error()})
	}
}
