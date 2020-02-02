package admin

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"math/rand"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	logg "gitlab.com/satit13/perfect_api/logger"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type service struct {
	repo Repository
}

// NewService creates new auth service
func NewService(admin Repository) (Service, error) {
	s := service{admin}
	return &s, nil
}

type Service interface {
	SignInAdminService(email string, password string) (interface{}, error)
	GetProfileAdminService(admincode string) (interface{}, error)
	AddAdminService(adminID string, email string, password string, roleID int64) (interface{}, error)
	UpdateprofileBackEndService(user *UpdateUserProfileModel, adminID string) (interface{}, error)
	Signout(token string) (interface{}, error)
	AddBackUserService(adminID string, email string, password string, roleID int64) (interface{}, error)
	ListAdminService(status int64, keyword string, limit int64) (int64, interface{}, error)
	ListRoleService(status int64) (interface{}, error)
	AddAdminServiceV2(adminID string, user *AddadminBackEndModel) (interface{}, error)
	GetAdminByUserID(adminID string) (interface{}, error)
	UpdateStatusAdminService(adminID string, status int64, adminCode string) (interface{}, error)
}

func (s *service) UpdateStatusAdminService(adminID string, status int64, adminCode string) (interface{}, error) {
	if status == 0 {
		resp, err := s.repo.RemoveAdmin(adminID, 0, adminCode)
		if err != nil {
			return nil, err
		}
		return resp, nil
	} else if status == 1 {
		resp, err := s.repo.RemoveAdmin(adminID, 1, adminCode)
		if err != nil {
			return nil, err
		}
		return resp, nil
	} else {
		return nil, errors.New("กรุณา ปรับ status 1 หรือ 0")
	}

	return nil, nil
}
func (s *service) GetAdminByUserID(adminID string) (interface{}, error) {
	resp, err := s.repo.GetAdminByIDRepo(adminID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) ListAdminService(status int64, keyword string, limit int64) (int64, interface{}, error) {
	var search string = ""
	words := strings.Fields(keyword)
	if len(words) > 0 {
		likescode := ` a.admin_code like '%` + words[0] + `%'`
		likesname := ` a.user_fname like '%` + words[0] + `%'`
		llkelname := ` a.user_lname like '%` + words[0] + `%'`
		likephone := ` a.phone like '%` + words[0] + `%'`
		likeemail := ` a.email like '%` + words[0] + `%'`
		for i := range words {
			if i > 0 {
				likescode += ` and a.admin_code like '%` + words[i] + `%'`
				likesname += ` and a.user_fname like '%` + words[i] + `%'`
				llkelname += ` and a.user_lname like '%` + words[i] + `%'`
				likephone += ` and a.phone like '%` + words[i] + `%'`
				likeemail += ` and a.email like '%` + words[i] + `%'`

			}
		}
		search = likescode + ` or ` + likesname + ` or ` + llkelname + ` or ` + likephone + ` or ` + likeemail
	} else {
		likescode := ` a.admin_code like '%` + keyword + `%'`
		likesname := ` a.user_fname like '%` + keyword + `%'`
		llkelname := ` a.user_lname like '%` + keyword + `%'`
		likephone := ` a.phone like '%` + keyword + `%'`
		likeemail := ` a.email like '%` + keyword + `%'`
		search = likescode + ` or ` + likesname + ` or ` + llkelname + ` or ` + likephone + ` or ` + likeemail
	}

	resp, err := s.repo.ListAdminRepo(status, search, limit)
	if err != nil {
		return 0, nil, err
	}
	len, err := s.repo.CountAdminRepo(status, search)
	if err != nil {
		return 0, nil, err
	}

	return len, resp, nil
}

func (s *service) ListRoleService(status int64) (interface{}, error) {
	resp, err := s.repo.ListRoleRepo(status)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) AddAdminServiceV2(adminID string, user *AddadminBackEndModel) (interface{}, error) {
	if len(user.Password) <= 4 {
		return nil, errors.New("กรุณา กรอกรหัสผ่าน มากกว่า 4 ตัว")
	}
	_, err := s.repo.FindAdminBack(user.Email)
	if err != nil {
		return nil, err
	}
	pass := hashAndSalt(user.Password)
	resp, err := s.repo.AddadminUser(adminID, pass, user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) AddBackUserService(adminID string, email string, password string, roleID int64) (interface{}, error) {

	_, err := s.repo.FindAdminBack(email)
	if err != nil {
		return nil, err
	}

	pass := hashAndSalt(password)
	_, err = s.repo.SigupUserBack(adminID, email, pass, roleID)
	if err != nil {
		return nil, err
	}
	Account := map[string]interface{}{
		"email":   email,
		"role_id": roleID,
	}

	return Account, nil
}
func (s *service) Signout(token string) (interface{}, error) {
	_, err := s.repo.RemoveToken(token)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func randomstring() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // E.g. "ExcbsVQs"
	return str
}

func (s *service) UpdateprofileBackEndService(user *UpdateUserProfileModel, adminID string) (interface{}, error) {
	_, err := s.repo.UpdateProfileBackEndRepo(user, adminID)
	if err != nil {
		return nil, err
	}
	if user.Password != "" {
		pass := hashAndSalt(user.Password)
		_, err = s.repo.UpdatePassword(user, pass, adminID)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (s *service) AddAdminService(adminID string, email string, password string, roleID int64) (interface{}, error) {

	pass := hashAndSalt(password)
	_, err := s.repo.AddUserAdminRepo(adminID, email, pass, roleID)
	if err != nil {
		return nil, err
	}
	Account := map[string]interface{}{
		"email":   email,
		"role_id": roleID,
	}

	return Account, nil
}
func sendMail(email string, subject, password string) error {

	//mailTmpl := template.Must(template.ParseFiles())
	mailTmpl, err := template.New("webpage").Parse(htmls)
	if err != nil {

		return err
	}

	data := struct {
		Email    string
		Password string
	}{
		Email:    email,
		Password: password,
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
func hashAndSalt(pwd string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), 8)

	if err != nil {
		log.Println(err)
	}

	return string(hashedPassword)
}

func (s *service) FindUserAdminbyEmail(email string) (*UserModel, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		logg.Println(err.Error())
		return nil, err
	}
	ap := UserModel{
		Code:     user.Code,
		Password: user.Password,
	}
	return &ap, nil
}

func (s *service) GetProfileAdminService(admincode string) (interface{}, error) {
	profile, err := s.repo.GetProfileAdmin(admincode)
	if err != nil {
		return nil, err
	}
	// pr := ProfileMemberModel{
	// 	Email:     "admin@gmail.com",
	// 	Phone:     "+66121312345",
	// 	ShopName:  "admin",
	// 	Gender:    0,
	// 	BrithDate: "1994/04/15",
	// }
	// ad := []AddressProfileModel{
	// 	AddressProfileModel{
	// 		AddressID:   1,
	// 		Name:        "admin",
	// 		Phone:       "+661213131231",
	// 		Address:     "123/7 บ้านตำหนัก หมู่3 ต.ดอนแก้ว",
	// 		Province:    "chiang mai",
	// 		District:    "อ.สารภี",
	// 		PostalCode:  "50140",
	// 		MainAddress: 1,
	// 	},
	// 	AddressProfileModel{
	// 		AddressID:   2,
	// 		Name:        "admin",
	// 		Phone:       "+66123456789",
	// 		Address:     "123/7 บ้านตำหนัก หมู่3 ต.ดอนแก้ว",
	// 		Province:    "chiang mai",
	// 		District:    "อ.สารภี",
	// 		PostalCode:  "50140",
	// 		MainAddress: 0,
	// 	},
	// }

	// profile := ProfileModel{
	// 	UserID:  "1232131",
	// 	Profile: pr,
	// 	Address: ad,
	// }
	return profile, nil
}

func (s *service) SignInAdminService(email string, password string) (interface{}, error) {

	co, err := s.FindUserAdminbyEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(co.Password), []byte(password))
	if err != nil {
		return nil, errors.New("your password not math")
	}
	newUUID, err := uuid.NewV4()
	newtoken2, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	randomKey := strings.Replace(newUUID.String(), "-", "", -1)
	_, err = s.repo.SaveToekenAdmin(email, randomKey)
	if err != nil {
		return nil, err
	}
	Account := map[string]string{
		"token_1": randomKey,
		"token_2": strings.Replace(newtoken2.String(), "-", "", -1),
	}
	return &Account, nil
}
