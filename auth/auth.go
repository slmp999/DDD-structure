package auth

import (
	"errors"
	"time"
)

// Errors
var (
	ErrTokenNotFound = errors.New("auth: token not found")
	ErrTokenExpired  = errors.New("token_expired")
)

// Token domain-model
// type Token struct {
// 	ID          string
// 	ClientID    int64
// 	AccountID   int64
// 	VendingID   int64
// 	VendingUUID string
// 	TokenID     string
// 	PartNer     Partner
// }
// type Endpoint struct {
// 	URL   string `json:"url"`
// 	Token string `json:"token"`
// }
// type Partner struct {
// 	Endpoint Endpoint
// }

// type PartnerV2 struct {
// 	Endpoint EndpointV2
// }

// type EndpointV2 struct {
// 	Name  string `json:"partner_name"`
// 	URL   string `json:"url"`
// 	Token string `json:"token"`
// }

// Account domain-model
// type Account struct {
// 	ID           int64
// 	ClientID     int64
// 	FullName     string
// 	Email        string
// 	PasswordHash string
// 	Status       bool
// }

// Access domain-model
// type Access struct {
// 	TokenID   string
// 	AccountID int64
// 	ClientID  int64
// 	// VendingID   int64
// 	// Meta        string
// 	// VendingUUID string
// }

// type Meta struct {
// 	Id   int64          `json:"id" db:"id"`
// 	Meta sql.NullString `json:"meta,omitempty" db:"meta,omitempty"`
// }
// type MetaPartner struct {
// 	ID          int64          `json:"id" db:"id"`
// 	PartnerName string         `json:"name" db:"name"`
// 	Meta        sql.NullString `json:"meta,omitempty" db:"meta,omitempty"`
// 	EndpointURl string         `json:"endpoint_url" db:"endpoint_url"`
// }

// AccountToken domain-model
// type AccountToken struct {
// 	Token        string
// 	AccountID    int64
// 	FullName     string
// 	Email        string
// 	IsRegistered bool
// 	MQTTKey      string
// 	MQTTChannel  string
// }

// type AccountTokenV2 struct {
// 	Token        string `json:"token,omitempty"`
// 	AccountID    int64  `json:"account_id,omitempty"`
// 	FullName     string `json:"full_name,omitempty"`
// 	RoleID       int64  `json:"role_id,omitempty"`
// 	RoleName     string `json:"role_name,omitempty"`
// 	Email        string `json:"email,omitempty"`
// 	IsRegistered bool   `json:"is_registered,omitempty"`
// 	MQTTKey      string `json:"mqtt_key,omitempty"`
// 	MQTTChannel  string `json:"mqtt_channel,omitempty"`
// }

type AuthProfileAddress struct {
	AddressID   int64  `json:"address_id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Province    string `json:"province"`
	District    string `json:"district"`
	PostalCode  string `json:"postal_code"`
	MainAddress int    `json:"main_address"`
}

type ListProfileAddressModel struct {
	UserID int64                `json:"user_id"`
	Data   []AuthProfileAddress `json:"data"`
}
type GetProfileAddressModel struct {
	UserID  int64              `json:"user_id"`
	Address AuthProfileAddress `json:"address"`
}
type AuthenShoppingModel struct {
	ID       int64  `json:"id"`
	Code     string `json:"Code"`
	Token    string `json:"token_1"`
	Token2   string `json:"token_2"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Phone1   string `json:"phone_1"`
	Phone2   string `json:"phone_2"`
	Email    string `json:"email"`
	Address1 string `json:"address_1"`
	Address2 string `json:"address_2"`
	Zipcode  int64  `json:"zipcode"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Gender   string `json:"gender"`
	RoleID   int64  `json:"role_id"`
	RoleName string `json:"role_name"`
}

type ForgetModel struct {
	Code string `json:"code"`
}

// type AccountV2 struct {
// 	ID           int64
// 	ClientID     int64
// 	FullName     string
// 	Email        string
// 	PasswordHash string
// 	Status       bool
// }

// Profile domain model
// type Profile struct {
// 	AccountID  int64
// 	FullName   string
// 	Email      string
// 	PictureURL *string
// 	Status     int
// }

// type AccountClientRole struct {
// 	ID           int64 `json:"id" db:"id"`
// 	AccountID    int64 `json:"account_id" db:"account_id"`
// 	ClientRoleID int64 `json:"client_role_id" db:"client_role_id"`
// }

// type ClientRole struct {
// 	ID   int64  `json:"role_id" db:"id"`
// 	Name string `json:"role_name" db:"name"`
// }

// Repository is the auth storage
type Repository interface {
	SignUpPhoneRepo(code string, refcode string, invcode string, password string) (interface{}, error)
	SignUpEmailRepo(code string, refcode string, invcode string, password string) (interface{}, error)
	SignInWithEmailRepo(code string, password string) (interface{}, error)
	SignInWithPhoneRepo(code string, password string) (interface{}, error)
	FindUserID() (string, error)
	FindUserTel(code string) (bool, error)
	FindUserEmail(code string) (bool, error)
	SaveTokenEmail(code string, token string) (interface{}, error)
	SaveTokenPhone(code string, token string) (interface{}, error)
	FindUserByPhone(code string) (*UserModelPhone, error)
	FindUserByEmail(code string) (*UserModelEmail, error)
	GetToken(tokenID string, mode string) (*Token, error)
	SaveToekenAdmin(email string, token string) (interface{}, error)
	GetInviteCode() string
	RemoveToken(tokenID string) (interface{}, error)
	GetResetTokenEmail(email string) (string, error)
	GetResetTokenPhone(phone string) (string, error)
	ResetPasswordRepo(token string, newpassword string) (interface{}, error)
	ChangePasswordEmailRepo(email string, password string) (interface{}, error)
	ChangePasswordPhoneRepo(phone string, password string) (interface{}, error)
	SaveTokenByTokenReset(reset_token string, token string) (interface{}, error)
	FindUserByUID(userID string) (string, error)
	ChangePasswordUserID(UserID string, pass string) (interface{}, error)
	// GetAccount(email string) (*Account, error)
	// GetProfile(token string) (*Profile, error)
}

type Token struct {
	UserID     string    `json:"user_id" db:"user_id"`
	AdminCode  string    `json:"admin_code" db:"admin_code"`
	Token      string    `json:"access_token" db:"access_token"`
	AccessTime time.Time `db:"access_time"`
}

// Generated by https://quicktype.io

type ModelSMS struct {
	MessageCount string    `json:"message-count"`
	Messages     []Message `json:"messages"`
}

type Message struct {
	To               string `json:"to"`
	MessageID        string `json:"message-id"`
	Status           string `json:"status"`
	RemainingBalance string `json:"remaining-balance"`
	MessagePrice     string `json:"message-price"`
	Network          string `json:"network"`
	Error            string `json:"error-text"`
}

type UserModelPhone struct {
	Code     string `json:"telephone" db:"telephone"`
	Password string `json:"pwd" db:"pwd"`
}

type UserModelEmail struct {
	Code     string `json:"email" db:"email"`
	Password string `json:"pwd" db:"pwd"`
}

type UserModel struct {
	Code     string `json:"code"`
	Password string `json:"password"`
}

// Generated by https://quicktype.io

type ModelRequestForgotPassword struct {
	Code string `json:"code"`
	Host string `json:"host"`
	Part string `json:"path"`
}
