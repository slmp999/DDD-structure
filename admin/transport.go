package admin

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/acoshift/hrpc"
	"gitlab.com/satit13/perfect_api/auth"
	logg "gitlab.com/satit13/perfect_api/logger"
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

	mux.Handle("/list/role", m.Handler(makeListRole(s)))
	mux.Handle("/list/admin", m.Handler(MakeEndpointListAdmin(s)))
	mux.Handle("/", m.Handler(MakeAccountHandler(s, &m)))
	return (mux)
}

// MakeAccountHandler creates new account handler
func MakeAccountHandler(s Service, m *hrpc.Manager) http.Handler {

	mux := http.NewServeMux()
	mux.Handle("/signin", m.Handler(makeSigninAdmin(s)))
	mux.Handle("/signout", m.Handler(makeSignOutAdmin(s)))
	mux.Handle("/get/admin", m.Handler(makeGetAdmin(s)))
	mux.Handle("/update/admin", m.Handler(makeEditAdmin(s)))
	mux.Handle("/change/password", m.Handler(makeChangePasswordAdmin(s)))
	mux.Handle("/add/admin", m.Handler(makeAddAdminV2Endpoint(s)))
	mux.Handle("/add/back/user", m.Handler(makeAddBackUser(s)))
	mux.Handle("/get/admin/adminid", m.Handler(makeAdminByAdminID(s)))
	mux.Handle("/remove/admin/adminid", m.Handler(makeUpdateStatusAdmin(s)))
	return mustLogin(s)(mux)
}

func mustLogin(s Service) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			AdminID := auth.GetAdminID(r.Context())
			if AdminID == "" {
				if !CheckUrlstring(r) {
					errorEncoder(w, r, errForbidden)
					fmt.Println("error mustLogin admin.transport.go")
					return
				}
			}
			enableCors(&w)
			h.ServeHTTP(w, r)
		})
	}
}

func CheckUrlstring(r *http.Request) bool {

	if strings.Contains(r.URL.String(), "signup") || strings.Contains(r.URL.String(), "/signin") {
		logg.Println(strings.Contains(r.URL.String(), "signin"))
		return true
	}

	return false
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
		encoder(w, status, &errorResponse{Response: "false", Message: err.Error()})
	}
}
