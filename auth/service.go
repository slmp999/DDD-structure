package auth

import (
	"bytes"
	"context"
	"encoding/base32"
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	logg "gitlab.com/satit13/perfect_api/logger"
	"gopkg.in/gomail.v2"

	//"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go"
	"github.com/nyaruka/phonenumbers"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	uuid "github.com/satori/go.uuid"
	"github.com/subosito/twilio"
	"golang.org/x/crypto/bcrypt"
)

type (
	keyToken struct{}
)

// Service is the auth service
type Service interface {
	Signout(token string) (interface{}, error)

	SigninService(Code string, password string) (interface{}, error)
	SignUpService(phone string, RefCode string, OTPvalidate string, password string) (interface{}, error)
	OTPRequest(phone string) (interface{}, error)
	ForgetPassword(code string) (*ForgetModel, error)
	ResetPassWordService(token string, newpassword string) (interface{}, error)
	OTPValidate(code string, passcode string) (bool, error)
	SwitchOtpRequest(code string) (interface{}, error)
	GetUserService(Code string) (*AuthenShoppingModel, error)
	UpdateProfileUser(auth *AuthProfileAddress) (interface{}, error)
	AddProfileAddressService(auth *AuthProfileAddress) (interface{}, error)
	GetProfileADdressByID(userID int64, addresID int64) (*GetProfileAddressModel, error)
	listProfileAddressService(userID int64) (*ListProfileAddressModel, error)
	DeleteProfileAddressService(userID int64, addressID int64) (interface{}, error)
	SwitchSaveToken(code string, token string) (interface{}, error)
	GetToken(tokenID string) (*Token, error)
	SwitchFogetEmail(model *ModelRequestForgotPassword) (interface{}, error)
	ChangePasswordService(code string, oldpassword string, newpassword string) (interface{}, error)

	//Getpartner(clientID int64) (*Endpoint, error)
	//GetpartnerV2(clientID int64, Name string) (*EndpointV2, error)
	//SigninV2(email string, password string, fcmtoken *string) (*AccountTokenV2, error)
}
type service struct {
	auths Repository
	Mode  string
}

// NewService creates new auth service
func NewService(auths Repository, Mode string) (Service, error) {
	s := service{auths, Mode}
	return &s, nil
}

func GetUserID(ctx context.Context) string {
	x, ok := ctx.Value(keyToken{}).(*Token)
	if !ok {
		return ""
	}
	return x.UserID
}

func GetAdminID(ctx context.Context) string {
	x, ok := ctx.Value(keyToken{}).(*Token)
	if !ok {
		return ""
	}
	return x.AdminCode
}

// GetAccountID returns account id from context
// func GetAccountID(ctx context.Context) int64 {
// 	x, ok := ctx.Value(keyToken{}).(*Token)
// 	if !ok {
// 		return -1
// 	}
// 	return x.AccountID
// }

// GetTokenID return access token
// func GetTokenID(ctx context.Context) string {
// 	x, ok := ctx.Value(keyToken{}).(*Token)
// 	if !ok {
// 		return ""
// 	}
// 	return x.TokenID
// }

func (s *service) SwitchOtpRequest(code string) (interface{}, error) {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	if strings.ContainsRune(code, '@') {
		_, err := s.OTPEmailREQuest(code)
		if err != nil {
			return nil, err
		}
		//return nil, errors.New("การยืนยัน ด้วย email ยังไม่เปิดใช้บริการ")
		// Validation failed
	} else if code[0] == '+' || re.MatchString(code) {

		phone := ConvertPhonetoCoutryCode(code)
		_, err := s.auths.FindUserTel(phone)
		if err != nil {
			return nil, err
		}
		if s.Mode == "Production" {
			_, err = s.OTPRequest(phone)
			if err != nil {
				return nil, err
			}
		}

	} else {
		return nil, errors.New("OTP invalid")
	}
	return nil, nil
}

func (s *service) DeleteProfileAddressService(userID int64, addressID int64) (interface{}, error) {
	return nil, nil
}

func (s *service) listProfileAddressService(userID int64) (*ListProfileAddressModel, error) {
	auth := AuthProfileAddress{
		AddressID:   1,
		Name:        "นาย ชารัต",
		Phone:       "+66979679089",
		Province:    "เชียงใหม่",
		District:    "สารภี",
		PostalCode:  "50140",
		MainAddress: 0,
	}
	Address := []AuthProfileAddress{}
	Address = append(Address, auth)
	auth = AuthProfileAddress{
		AddressID:   1,
		Name:        "นาย ชารัต",
		Phone:       "+66979679089",
		Province:    "เชียงใหม่",
		District:    "สารภี",
		PostalCode:  "50140",
		MainAddress: 0,
	}
	Address = append(Address, auth)
	list := ListProfileAddressModel{
		UserID: userID,
		Data:   Address,
	}
	return &list, nil
}

func (s *service) GetProfileADdressByID(userID int64, addresID int64) (*GetProfileAddressModel, error) {
	auth := AuthProfileAddress{
		AddressID:   1,
		Name:        "นาย ชารัต",
		Phone:       "+66979679089",
		Province:    "เชียงใหม่",
		District:    "สารภี",
		PostalCode:  "50140",
		MainAddress: 0,
	}
	get := GetProfileAddressModel{
		UserID:  userID,
		Address: auth,
	}
	return &get, nil
}
func (s *service) AddProfileAddressService(auth *AuthProfileAddress) (interface{}, error) {
	return nil, nil
}
func (s *service) UpdateProfileUser(auth *AuthProfileAddress) (interface{}, error) {
	return nil, nil
}

func (s *service) GetUserService(Code string) (*AuthenShoppingModel, error) {
	newUUID, err := uuid.NewV4()
	newUUID2, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	randomKey := strings.Replace(newUUID.String(), "-", "", -1)
	Account := AuthenShoppingModel{
		ID:       1,
		Code:     Code,
		Token:    randomKey,
		Token2:   strings.Replace(newUUID2.String(), "-", "", -1),
		Fname:    "admin",
		Lname:    "administrator",
		Phone1:   "+66974520603",
		Phone2:   "+66974520602",
		Email:    "test@gmail.com",
		Address1: "123213",
		Address2: "12312312",
		Zipcode:  50140,
		City:     "เชียงใหม่",
		Country:  "thailand",
		Gender:   "ชาย",
		RoleID:   1,
		RoleName: "Customer",
	}
	return &Account, nil
}

func sendMail(email, subject, body string) error {

	//mailTmpl := template.Must(template.ParseFiles())
	mailTmpl, err := template.New("webpage").Parse(htmls)
	if err != nil {

		return err
	}

	data := struct {
		OTP string
	}{
		body,
	}
	buf := new(bytes.Buffer)
	err = mailTmpl.Execute(buf, data)
	if err != nil {

		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "Perfect <no-reply@Perfect.com>")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", buf.String())
	//d := gomail.NewDialer("smtp.mailtrap.io", 2525, "398070e236371be32", "ab7e34e242c89e")
	d := gomail.NewDialer("smtp.gmail.com", 587, "crackergametv@gmail.com", "0979679089")
	err = d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) OTPEmailREQuest(code string) (interface{}, error) {

	t := time.Now()
	var vdOption totp.ValidateOpts
	vdOption.Digits = otp.DigitsSix
	vdOption.Algorithm = otp.AlgorithmSHA1
	vdOption.Period = 360
	vdOption.Skew = 1
	telBase32 := base32.StdEncoding.EncodeToString([]byte(code))
	// Gen Passcode for change for Period 30 sec -> 60 sec
	pass, err := totp.GenerateCodeCustom(telBase32, t, vdOption)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = sendMail(code, "ยืนยันบัญชีของคุณ", pass)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) SendMAilService(email string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "Perfect <no-reply@Perfect.com>")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//d := gomail.NewDialer("smtp.mailtrap.io", 2525, "398070e236371be32", "ab7e34e242c89e")
	d := gomail.NewDialer("smtp.gmail.com", 587, "crackergametv@gmail.com", "0979679089")
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SendSms(phone string, passcode string) (bool, error) {
	var (
		AccountSid = "AC26da1afd834d89427c9aec4f5000f622"
		AuthToken  = "2f697644d064e692409ce4679aa17dd8"
		// To         = "" Your phone number, including the +N zone number.
	)
	c := twilio.NewClient(AccountSid, AuthToken, nil)

	params := twilio.MessageParams{
		Body: "OTP for validate Perfect Shopping is " + passcode,
	}

	p, resp, err := c.Messages.Send("+12563048791", phone, params)
	log.Println("Send:", p)
	log.Println("Response:", resp)
	log.Println("Err:", err)
	return false, nil
}

func (s *service) OTPValidate(code string, passcode string) (bool, error) {
	if len(passcode) != 6 {
		return false, errors.New("กรุณา กรอกรหัส OTP ให้ถูกต้อง")
	}
	//re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	if !strings.ContainsRune(code, '@') {
		code = ConvertPhonetoCoutryCode(code)
	}
	phoneBase32 := base32.StdEncoding.EncodeToString([]byte(code))
	var vdOption totp.ValidateOpts
	vdOption.Digits = otp.DigitsSix
	vdOption.Algorithm = otp.AlgorithmSHA1
	vdOption.Period = 360
	vdOption.Skew = 1
	valid, err := totp.ValidateCustom(passcode, phoneBase32, time.Now(), vdOption)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	if !valid {

		return false, errors.New("กรุณา กรอกรหัส OTP ให้ถูกต้อง")
	}
	return true, nil
}

func (s *service) ResetPassWordService(token string, newpassword string) (interface{}, error) {
	if token == "" {
		return nil, errors.New("token นี้ไม่สามารถใช้ได้")
	}
	if len(newpassword) < 6 {
		return nil, errors.New("ท่านไม่สามารถใช้รหัสผ่านนี้ได้ รหัสขอท่าน ต่ำกว่า 6 ตัว")
	}
	pass := hashAndSalt(newpassword)
	newUUID, err := uuid.NewV4()
	newtoken2, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	randomKey := strings.Replace(newUUID.String(), "-", "", -1)
	_, err = s.auths.SaveTokenByTokenReset(token, randomKey)
	if err != nil {
		return nil, err
	}
	Account := map[string]string{
		"token_1": randomKey,
		"token_2": strings.Replace(newtoken2.String(), "-", "", -1),
	}
	_, err = s.auths.ResetPasswordRepo(token, pass)
	logg.Println(1)
	if err != nil {
		return nil, err
	}
	return &Account, nil
}
func (s *service) ForgetPassword(code string) (*ForgetModel, error) {
	forget := ForgetModel{
		Code: code,
	}
	return &forget, nil
}

func (s *service) ChangePasswordService(UserID string, oldpassword string, newpassword string) (interface{}, error) {
	HasPass, err := s.auths.FindUserByUID(UserID)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(HasPass), []byte(oldpassword))
	if err != nil {
		return nil, errors.New("รหัสผ่านเก่าไม่ตรงกับรหัสผ่าน")
	}
	pass := hashAndSalt(newpassword)
	_, err = s.auths.ChangePasswordUserID(UserID, pass)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) SwitchChangePassword(code string, password string) (interface{}, error) {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	if strings.ContainsRune(code, '@') {
		_, err := s.auths.ChangePasswordEmailRepo(code, password)
		if err != nil {
			return nil, err
		}
	} else if code[0] == '+' || re.MatchString(code) {
		phone := ConvertPhonetoCoutryCode(code)
		_, err := s.auths.ChangePasswordPhoneRepo(phone, password)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Email or phone number")
	}
	return nil, nil
}

func (s *service) SMS(phone string, message string) (interface{}, error) {
	jsonData := map[string]string{
		"to":        phone,
		"from":      "Perfect",
		"text":      message,
		"apiKey":    "49351bdd-cb72-4d3b-9bd8-39667840f36e",
		"apiSecret": "e392f445-9746-417d-9116-cd67446ccf26",
	}
	jsonValue, err := json.Marshal(jsonData)
	response, err := http.Post("https://api.apitel.co/sms", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	log.Println(string(data))
	var smsResponse struct {
		Status string `json:"status"`
	}
	json.Unmarshal(data, &smsResponse)
	if smsResponse.Status != "ACCEPTED" {
		log.Println(err)
		return nil, err
	}
	return "", nil
}

func (s *service) OTPRequest(phone string) (interface{}, error) {
	// generate passcode from phone number
	t := time.Now()
	var vdOption totp.ValidateOpts
	vdOption.Digits = otp.DigitsSix
	vdOption.Algorithm = otp.AlgorithmSHA1
	vdOption.Period = 360
	vdOption.Skew = 1
	telBase32 := base32.StdEncoding.EncodeToString([]byte(phone))
	// Gen Passcode for change for Period 30 sec -> 60 sec
	pass, err := totp.GenerateCodeCustom(telBase32, t, vdOption)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	jsonData := map[string]string{
		"to":        phone,
		"from":      "Perfect",
		"text":      "OTP for validation" + pass,
		"apiKey":    "49351bdd-cb72-4d3b-9bd8-39667840f36e",
		"apiSecret": "e392f445-9746-417d-9116-cd67446ccf26",
	}
	jsonValue, err := json.Marshal(jsonData)
	response, err := http.Post("https://api.apitel.co/sms", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	log.Println(string(data))
	var smsResponse struct {
		Status string `json:"status"`
	}
	json.Unmarshal(data, &smsResponse)
	if smsResponse.Status != "ACCEPTED" {
		log.Println(err)
		return nil, err
	}
	return "", nil
	// _, err = s.SendSms(phone, pass)
	// if err != nil {
	// 	return nil, err
	// }
	// jsonData := map[string]string{
	// 	"from":      "NEXMO",
	// 	"text":      "OTP for validate register" + pass,
	// 	"to":        phone,
	// 	"apiKey":    "49351bdd-cb72-4d3b-9bd8-39667840f36e",
	// 	"apiSecret": "e392f445-9746-417d-9116-cd67446ccf26",
	// }
	// jsonValue, err := json.Marshal(jsonData)
	// if err != nil {
	// 	return nil, err
	// }

	// data := url.Values{}
	//text := "OTP for validate register" + pass
	// data.Set("api_key", "996a8ccc")
	// data.Add("apiSecret", "Mux38nTbsiFJmZ1S")
	// data.Add("to", phone)
	// data.Add("from", "NEXMO")
	// data.Add("text", text)
	// urlapi := "https://rest.nexmo.com/sms/json"

	// client := &http.Client{}

	// r, _ := http.NewRequest("POST", urlapi, strings.NewReader(data.Encode())) // URL-encoded payload
	// r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// r.Header.Add("Cache-Control", "no-cache")
	// r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	// response, err := client.Do(r)

	// if err != nil {
	// 	return nil, err
	// }
	// log.Println(response.Body)
	// datas, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// sms := ModelSMS{}
	// json.Unmarshal(datas, &sms)
	// log.Println(sms)
	// url := "https://rest.nexmo.com/sms/json"

	// payload := strings.NewReader("api_key=996a8ccc&api_secret=Mux38nTbsiFJmZ1S&to=" + phone + "&from=NEXMO&text=" + text)

	// req, _ := http.NewRequest("POST", url, payload)

	// req.Header.Add("User-Agent", "PostmanRuntime/7.15.2")
	// req.Header.Add("Accept", "*/*")
	// req.Header.Add("Cache-Control", "no-cache")
	// req.Header.Add("Postman-Token", "84290f72-1d83-416d-a3ba-2805e5056cab,74a5188a-d585-49ec-9d4a-0dad6a19b139")
	// req.Header.Add("Host", "rest.nexmo.com")
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("Accept-Encoding", "gzip, deflate")
	// req.Header.Add("Content-Length", "145")
	// req.Header.Add("Connection", "keep-alive")
	// req.Header.Add("cache-control", "no-cache")

	// res, _ := http.DefaultClient.Do(req)

	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body), pass)
	//return nil, nil
}

func (s *service) GetTemplateHtmlFogetEmail(model *ModelRequestForgotPassword, token string) (string, error) {
	mailTmpl, err := template.New("webpage").Parse(templateSentEmailResetPassword)
	if err != nil {
		return "", err
	}

	data := struct {
		LINK string
	}{
		model.Host + "/" + model.Part + "/" + token,
	}
	buf := new(bytes.Buffer)
	err = mailTmpl.Execute(buf, data)
	if err != nil {

		return "", err
	}
	return buf.String(), nil
}

func (s *service) SwitchFogetEmail(model *ModelRequestForgotPassword) (interface{}, error) {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	if strings.ContainsRune(model.Code, '@') {
		token, err := s.auths.GetResetTokenEmail(model.Code)
		if err != nil {
			return nil, err
		}
		template, _ := s.GetTemplateHtmlFogetEmail(model, token)
		err = s.SendMAilService(model.Code, "ยืนยันบัญการขอเปลี่ยนรหัสผ่าน", template)
		if err != nil {
			return nil, err
		}
		message := model.Host + "/" + model.Part + "/" + token
		return message, nil
		//return nil, errors.New("การยืนยัน ด้วย email ยังไม่เปิดใช้บริการ")
		// Validation failed
	} else if model.Code[0] == '+' || re.MatchString(model.Code) {
		phone := ConvertPhonetoCoutryCode(model.Code)
		token, err := s.auths.GetResetTokenPhone(phone)
		if err != nil {
			return nil, err
		}

		message := "url สำหรับการเปลี่ยนรหัสผ่าน \n" + model.Host + "/" + model.Part + "/" + token
		_, err = s.SMS(phone, message)
		if err != nil {
			return nil, err
		}
		messages := model.Host + "/" + model.Part + "/" + token
		return messages, nil
	} else {
		return nil, errors.New("Email or phone number")
	}
	return nil, nil
}

func (s *service) SwitchSignIn(code string, password string) (interface{}, error) {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	if strings.ContainsRune(code, '@') {
		_, err := s.auths.SignInWithEmailRepo(code, password)
		if err != nil {
			return nil, err
		}
		//return nil, errors.New("การยืนยัน ด้วย email ยังไม่เปิดใช้บริการ")
		// Validation failed
	} else if code[0] == '+' || re.MatchString(code) {
		phone := ConvertPhonetoCoutryCode(code)
		_, err := s.auths.SignInWithPhoneRepo(phone, password)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Email or phone number")
	}
	return nil, nil
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *service) SigninService(Code string, password string) (interface{}, error) {

	// set some claims
	co, err := s.SwitchFindUser(Code)

	if err != nil {
		return nil, err
	}
	// if co.Password != password {
	// 	return nil, errors.New("your password not math")
	// }
	err = bcrypt.CompareHashAndPassword([]byte(co.Password), []byte(password))
	if err != nil {
		return nil, errors.New("รหัสผู้ใช้งาน หรือ รหัสผ่านไม่ถูกต้อง กรุณาลองอีกครั้ง")
	}
	newUUID, err := uuid.NewV4()
	newtoken2, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	// expireTime := time.Now().Add(-(5 * 24 * time.Hour))
	// if co.ExpireDateToken.Before(expireTime) {
	// 	randomKey := strings.Replace(newUUID.String(), "-", "", -1)
	// 	_, err = s.SwitchSaveToken(Code, randomKey)
	// 	Account := map[string]string{
	// 		"token_1": randomKey,
	// 		"token_2": strings.Replace(newtoken2.String(), "-", "", -1),
	// 	}
	// 	return &Account, nil
	// }
	randomKey := strings.Replace(newUUID.String(), "-", "", -1)
	_, err = s.SwitchSaveToken(Code, randomKey)
	Account := map[string]string{
		"token_1": randomKey,
		"token_2": strings.Replace(newtoken2.String(), "-", "", -1),
	}
	return &Account, nil
}

func ConvertPhonetoCoutryCode(code string) string {
	defaultRegion := "TH"
	info, _ := phonenumbers.Parse(code, defaultRegion)
	formattedNum := phonenumbers.Format(info, phonenumbers.NANPA_COUNTRY_CODE)
	phone := strings.Replace(formattedNum, " ", "", -1)
	return phone
}
func (s *service) SwitchFindUser(code string) (*UserModel, error) {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	if strings.ContainsRune(code, '@') {
		user, err := s.auths.FindUserByEmail(code)
		if err != nil {
			return nil, err
		}
		ap := UserModel{
			Code:     user.Code,
			Password: user.Password,
		}
		return &ap, nil
		//return nil, errors.New("การยืนยัน ด้วย email ยังไม่เปิดใช้บริการ")
		// Validation failed
	} else if code[0] == '+' || re.MatchString(code) {
		phone := ConvertPhonetoCoutryCode(code)
		user, err := s.auths.FindUserByPhone(phone)
		if err != nil {
			return nil, err
		}
		ap := UserModel{
			Code:     user.Code,
			Password: user.Password,
		}
		return &ap, nil
	} else {
		return nil, errors.New("Email or phone number")
	}
	return nil, nil
}

func (s *service) SwitchSaveToken(code string, token string) (interface{}, error) {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	if strings.ContainsRune(code, '@') {
		_, err := s.auths.SaveTokenEmail(code, token)
		if err != nil {
			return nil, err
		}
		//return nil, errors.New("การยืนยัน ด้วย email ยังไม่เปิดใช้บริการ")
		// Validation failed
	} else if code[0] == '+' || re.MatchString(code) {
		phone := ConvertPhonetoCoutryCode(code)
		_, err := s.auths.SaveTokenPhone(phone, token)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Email or phone number")
	}
	return nil, nil
}

func (s *service) SwitchSignUp(code string, RefCode string, password string) (interface{}, error) {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	if strings.ContainsRune(code, '@') {
		inv := s.auths.GetInviteCode()
		_, err := s.auths.SignUpEmailRepo(code, RefCode, inv, password)
		if err != nil {
			return nil, err
		}
		//return nil, errors.New("การยืนยัน ด้วย email ยังไม่เปิดใช้บริการ")
		// Validation failed
		ap := map[string]interface{}{
			"email":      code,
			"ref_code":   RefCode,
			"invte_code": inv,
		}
		return ap, nil
	} else if code[0] == '+' || re.MatchString(code) {
		phone := ConvertPhonetoCoutryCode(code)
		inv := s.auths.GetInviteCode()
		_, err := s.auths.SignUpPhoneRepo(phone, RefCode, inv, password)
		if err != nil {
			return nil, err
		}
		ap := map[string]interface{}{
			"phone":       code,
			"ref_code":    RefCode,
			"invite_code": inv,
		}
		return ap, nil
	} else {
		return nil, errors.New("Email or phone number")
	}
	return nil, nil
}

func (s *service) SignUpService(phone string, RefCode string, OTPvalidate string, password string) (interface{}, error) {
	cphone := ConvertPhonetoCoutryCode(phone)
	if s.Mode == "Production" {
		_, err := s.OTPValidate(cphone, OTPvalidate)
		if err != nil {
			return nil, err
		}
		if len(RefCode) <= 2 || RefCode == "" {
			RefCode = "f2958c905e5b4bf68631204b9372c0ff"
		}
	}

	pass := hashAndSalt(password)
	resp, err := s.SwitchSignUp(cphone, RefCode, pass)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func hashAndSalt(pwd string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), 8)

	if err != nil {
		log.Println(err)
	}

	return string(hashedPassword)
}

// func (s *service) Signin(email, password string, fcmtoken *string) (*AccountToken, error) {
// 	// account, err := s.auths.GetAccount(email)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// err = bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(password))
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// newUUID, err := uuid.NewV4()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// randomKey := strings.Replace(newUUID.String(), "-", "", -1)
// 	// // access := Access{
// 	// // 	TokenID:   randomKey,
// 	// // 	AccountID: account.ID,
// 	// // 	ClientID:  account.ClientID,
// 	// // }
// 	// // err = s.auths.SaveAccess(&access)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// at := AccountToken{
// 	// 	Token:        randomKey,
// 	// 	AccountID:    account.ID,
// 	// 	FullName:     account.FullName,
// 	// 	Email:        account.Email,
// 	// 	IsRegistered: account.Status,
// 	// }
// 	newUUID, err := uuid.NewV4()
// 	if err != nil {
// 		return nil, err
// 	}
// 	randomKey := strings.Replace(newUUID.String(), "-", "", -1)
// 	at := AccountToken{
// 		Token:        randomKey,
// 		AccountID:    1,
// 		FullName:     email,
// 		Email:        password,
// 		IsRegistered: true,
// 	}
// 	return &at, nil
// }

func (s *service) Signout(token string) (interface{}, error) {
	_, err := s.auths.RemoveToken(token)
	if err != nil {
		return nil, err
	}
	return nil, err
}

// func (s *service) Profile(xAccessToken string) (*Profile, error) {
// 	p, err := s.auths.GetProfile(xAccessToken)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p, nil
// }

func (s *service) GetToken(tokenID string) (*Token, error) {
	tk, err := s.auths.GetToken(tokenID, s.Mode)

	if err != nil {
		return nil, err
	}
	return tk, nil
}
