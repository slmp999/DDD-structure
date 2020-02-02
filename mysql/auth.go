package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	logg "gitlab.com/satit13/perfect_api/logger"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/satit13/perfect_api/auth"
	"upper.io/db.v3/lib/sqlbuilder"

	//"upper.io/db.v3/postgresql"
	"upper.io/db.v3/mysql"
)

// NewAuthRepository creates new auth repository
func NewAuthRepository(db *sql.DB) (auth.Repository, error) {
	// fmt.Println(1)
	pdb, err := mysql.New(db)
	if err != nil {
		return nil, err
	}
	dbx := sqlx.NewDb(db, "mysql")

	r := authRepository{pdb, dbx}
	// fmt.Println(1)
	return &r, nil
}

type authRepository struct {
	db  sqlbuilder.Database
	dbx *sqlx.DB
}

func (repo *authRepository) FindUIDByEmail(email string) (string, error) {
	var user_id string = ""
	sql1 := `select user_id as user_id from users where email = ? and active_status = 1 limit 1`
	rs, _ := repo.db.QueryRow(sql1, email)
	rs.Scan(&user_id)
	if user_id == "" {
		return "", errors.New("ไม่พบ email ของผู้ใช้")
	}
	return user_id, nil
}

func (repo *authRepository) FindUIDByPhone(phone string) (string, error) {
	var user_id string = ""
	sql1 := `select user_id as user_id from users where  telephone = ? and active_status = 1 limit 1`
	rs, _ := repo.db.QueryRow(sql1, phone)
	rs.Scan(&user_id)
	if user_id == "" {
		return "", errors.New("ไม่พบ email ของผู้ใช้")
	}
	return user_id, nil
}

func (repo *authRepository) FindUIDAdmin(email string) (string, error) {
	var admin_code string = ""
	sql1 := `select admin_code as admin_code from user_admin where  email = ? limit 1`
	rs, _ := repo.db.QueryRow(sql1, email)
	rs.Scan(&admin_code)
	if admin_code == "" {
		return "", errors.New("ไม่พบ email ของผู้ใช้")
	}
	return admin_code, nil
}

func (repo *authRepository) FindUserByUID(userID string) (string, error) {
	var pwd string = ""
	sql := `select pwd from users where user_id = ? and active_status = 1 limit 1`
	rs, err := repo.db.QueryRow(sql, userID)
	if err != nil {
		return "", err
	}
	err = rs.Scan(&pwd)
	if err != nil {
		return "", err
	}
	return pwd, nil
}

func (repo *authRepository) ChangePasswordUserID(UserID string, pass string) (interface{}, error) {
	_, err := repo.db.Exec(`update users
	set pwd = ?
	where user_id = ?`,
		pass,
		UserID,
	)
	if err != nil {
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}

func (repo *authRepository) SaveTokenPhone(code string, token string) (interface{}, error) {
	// _, err := repo.dbx.NamedExec(`UPDATE users SET access_token = :token,
	// expire_date_token WHERE telephone = :telephone`,
	// 	map[string]interface{}{
	// 		"token":     token,
	// 		"exp_token": time.Now(),
	// 		"telephone": code,
	// 	},
	// )
	// if err != nil {
	// 	log.Error(err)
	// 	return nil, err
	// }
	// return nil, nil
	userID, err := repo.FindUIDByPhone(code)
	if err != nil {
		return nil, err
	}
	q := repo.db.
		InsertInto("user_access").
		Values(map[string]interface{}{
			"company_id":   1,
			"user_id":      userID,
			"access_time":  time.Now(),
			"type":         1,
			"access_token": token,
		})
	_, err = q.Exec()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (repo *authRepository) FindUIDBYResetToken(token string) (string, error) {
	var user_id string = ""
	sql1 := `select user_id as user_id from users where reset_token = ? and active_status = 1 limit 1`
	rs, _ := repo.db.QueryRow(sql1, token)
	rs.Scan(&user_id)

	if user_id == "" {
		return "", errors.New("ไม่พบ email ของผู้ใช้")
	}
	return user_id, nil
}

func (repo *authRepository) SaveTokenByTokenReset(reset_token string, token string) (interface{}, error) {
	userID, err := repo.FindUIDBYResetToken(reset_token)
	if err != nil {
		return nil, err
	}
	q := repo.db.
		InsertInto("user_access").
		Values(map[string]interface{}{
			"company_id":   1,
			"user_id":      userID,
			"access_time":  time.Now(),
			"type":         1,
			"access_token": token,
		})
	_, err = q.Exec()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (repo *authRepository) SaveTokenEmail(code string, token string) (interface{}, error) {
	// _, err := repo.dbx.NamedExec(`UPDATE users SET access_token = :token,
	// expire_date_token = :exp_token WHERE email = :email`,
	// 	map[string]interface{}{
	// 		"token":     token,
	// 		"exp_token": time.Now(),
	// 		"email":     code,
	// 	},
	// )
	// if err != nil {
	// 	log.Error(err)
	// 	return nil, err
	// }
	// return nil, nil
	userID, err := repo.FindUIDByEmail(code)
	if err != nil {
		return nil, err
	}
	q := repo.db.
		InsertInto("user_access").
		Values(map[string]interface{}{
			"company_id":   1,
			"user_id":      userID,
			"access_time":  time.Now(),
			"type":         1,
			"access_token": token,
		})
	_, err = q.Exec()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (repo *authRepository) SaveToekenAdmin(email string, token string) (interface{}, error) {
	// _, err := repo.dbx.NamedExec(`UPDATE users SET access_token = :token ,
	// expire_date_token = :exp_token WHERE email = :email`,
	// 	map[string]interface{}{
	// 		"token":     token,
	// 		"exp_token": time.Now(),
	// 		"email":     email,
	// 	},
	// )
	// if err != nil {
	// 	log.Error(err)
	// 	return nil, err
	// }
	// userID, err := repo.FindUIDByPhone(code)
	// if err != nil {
	// 	return nil, err
	// }
	admin_code, err := repo.FindUIDAdmin(email)
	if err != nil {
		return nil, err
	}
	q := repo.db.
		InsertInto("user_access").
		Values(map[string]interface{}{
			"company_id":   1,
			"user_id":      admin_code,
			"access_time":  time.Now(),
			"type":         1,
			"access_token": token,
		})
	_, err = q.Exec()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (repo *authRepository) FindUserID() (string, error) {
	like := GetYear() + GetMonth() + GetDay()
	var userid string
	sql := `select COALESCE(user_id,'6208240000') as user_id from users where user_id like '%` + like + `%' order by id desc LIMIT 1`
	rs, _ := repo.db.QueryRow(sql)
	fmt.Println(sql)
	rs.Scan(&userid)

	return userid, nil
}

func (repo *authRepository) FindResetToekn(code string) (bool, error) {
	var reset_token int64 = 0
	sql := `select count(reset_token) as reset_token from users where reset_token = ? and active_status = 1`
	rs, _ := repo.db.QueryRow(sql, code)
	rs.Scan(&reset_token)
	if reset_token > 0 {
		return false, errors.New("email นี้ได้ทำการลงทะเบียนแล้ว")
	}
	return true, nil
}

func (repo *authRepository) GetResetToten() string {
	token, _ := uuid.NewV4()
	reset := strings.Replace(token.String(), "-", "", -1)
	ch, _ := repo.FindResetToekn(reset)
	if ch != true {
		reset = repo.GetResetToten()
	}
	return reset
}
func (repo *authRepository) ResetPasswordRepo(token string, newpassword string) (interface{}, error) {
	newResetToekn := repo.GetResetToten()
	_, err := repo.db.Exec(`update users
	set pwd = ?,reset_token = ?
	where reset_token = ?`,
		newpassword,
		newResetToekn,
		token,
	)
	if err != nil {
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil

}

func (repo *authRepository) updateResetTokenPhone(phone string, token string) (interface{}, error) {
	_, err := repo.db.Exec(`update users
	set reset_token = ?
	where telephone = ?`,
		token,
		phone,
	)
	if err != nil {
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}

func (repo *authRepository) updateResetTokenEmail(email string, token string) (interface{}, error) {
	_, err := repo.db.Exec(`update users
	set reset_token = ?
	where email = ?`,
		token,
		email,
	)
	if err != nil {
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}
func (repo *authRepository) GetResetTokenEmail(email string) (string, error) {
	var reset_token string = ""
	sql1 := `select (reset_token,'') as reset_token from users where email = ? and active_status = 1 limit 1`
	rs, err := repo.db.QueryRow(sql1, email)
	if err != nil {
		logg.Println(err.Error())
		return "", err
	}
	err = rs.Scan(&reset_token)
	if err != nil {
		logg.Println(err.Error())
		if err == sql.ErrNoRows {
			return "", errors.New("ไม่พบบัญชี ที่ตรงกับข้อมูลที่คุณระบุ")
		}
	}

	reset_token = repo.GetResetToten()
	_, err = repo.updateResetTokenEmail(email, reset_token)
	if err != nil {
		return "", err
	}

	return reset_token, nil
}
func (repo *authRepository) ChangePasswordEmailRepo(email string, password string) (interface{}, error) {
	_, err := repo.db.Exec(`update users
	set pwd = ?
	where email = ?`,
		password,
		email,
	)
	if err != nil {
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil

}
func (repo *authRepository) ChangePasswordPhoneRepo(phone string, password string) (interface{}, error) {
	_, err := repo.db.Exec(`update users
	set pwd = ?
	where telephone = ?`,
		password,
		phone,
	)
	if err != nil {
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}

func (repo *authRepository) GetResetTokenPhone(phone string) (string, error) {
	var reset_token string = ""
	sql1 := `select ifnull(reset_token,'') as reset_token from users where telephone = ? and active_status = 1 limit 1`
	rs, _ := repo.db.QueryRow(sql1, phone)
	err := rs.Scan(&reset_token)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("ไม่พบบัญชี ที่ตรงกับข้อมูลที่คุณระบุ")
		}
		return "", err
	}

	reset_token = repo.GetResetToten()
	_, err = repo.updateResetTokenPhone(phone, reset_token)
	if err != nil {
		return "", err
	}

	return reset_token, nil
}
func GenUserID(user string) (string, error) {
	var lastuserid string
	if user != "" {
		userid, _ := strconv.Atoi(user[len(user)-5:])
		lastnumber := (strconv.Itoa(userid + 1))
		fmt.Println(lastnumber)
		if len(string(lastnumber)) == 1 {
			lastuserid = "0000" + string(lastnumber)
		}
		if len(string(lastnumber)) == 2 {
			lastuserid = "000" + string(lastnumber)
		}
		if len(string(lastnumber)) == 3 {
			lastuserid = "00" + string(lastnumber)
		}
		if len(string(lastnumber)) == 4 {
			lastuserid = "0" + string(lastnumber)
		}
		if len(string(lastnumber)) == 5 {
			lastuserid = string(lastnumber)
		}
		return GetYear() + GetMonth() + GetDay() + lastuserid, nil
	} else {
		return GetYear() + GetMonth() + GetDay() + "00001", nil
	}
	return "", nil
}

func (repo *authRepository) SignInWithEmailRepo(code string, password string) (interface{}, error) {

	return nil, nil
}
func (repo *authRepository) SignInWithPhoneRepo(code string, password string) (interface{}, error) {

	return nil, nil
}

func (repo *authRepository) FindUserByPhone(code string) (*auth.UserModelPhone, error) {

	sql := `select telephone,pwd from users where telephone = ? and active_status = 1 limit 1`

	rs, err := repo.db.QueryRow(sql, code)
	if err != nil {
		return nil, err
	}
	user := auth.UserModelPhone{}
	rs.Scan(&user.Code, &user.Password)
	return &user, nil
}
func (repo *authRepository) FindUserByEmail(code string) (*auth.UserModelEmail, error) {
	sql := `select email,pwd from users where email = ? and active_status = 1 limit 1`

	rs, err := repo.db.QueryRow(sql, code)
	if err != nil {
		return nil, err
	}
	user := auth.UserModelEmail{}
	rs.Scan(&user.Code, &user.Password)
	return &user, nil
}

func (repo *authRepository) GetUserId(tokenID string, mode string) (*auth.Token, bool, error) {
	sql1 := `select user_id,access_time  from user_access where access_token = ? and type = 1 and status = 0  order by id desc limit 1`
	rs, err := repo.db.QueryRow(sql1, tokenID)
	if err != nil {
		return nil, false, err
	}
	user := auth.Token{}
	err = rs.Scan(&user.UserID, &user.AccessTime)
	expireTime := time.Now().Add(-(2 * 24 * time.Hour))

	logg.Println(expireTime)
	if err == sql.ErrNoRows {
		return nil, false, auth.ErrTokenNotFound
	}
	if user.AccessTime.Before(expireTime) && mode == "Production" {
		repo.UpdateExpirToken(tokenID)
		return nil, false, auth.ErrTokenExpired
	}
	if err != nil {
		return nil, false, err
	}
	user.Token = tokenID
	return &user, true, nil
}

func (repo *authRepository) RemoveToken(tokenID string) (interface{}, error) {
	_, err := repo.db.Exec(`update user_access
						set status = ? 
						where access_token = ?`,
		1,
		tokenID,
	)
	if err != nil {
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}

func (repo *authRepository) UpdateExpirToken(tokenID string) (interface{}, error) {

	_, err := repo.db.Exec(`update user_access
						set expire_status = ? 
						where access_token = ?`,
		1,
		tokenID,
	)
	if err != nil {
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}
func (repo *authRepository) GetTokenAdmin(tokenID string, mode string) (*auth.Token, bool, error) {
	sql1 := `select user_id,access_time 
	from user_access where access_token = ? and type = 2 and status = 0 order by id desc limit 1`
	rs, err := repo.db.QueryRow(sql1, tokenID)
	if err != nil {
		return nil, false, err
	}
	user := auth.Token{}
	err = rs.Scan(&user.AdminCode, &user.AccessTime)

	logg.Println(user)
	if err == sql.ErrNoRows {
		return nil, false, auth.ErrTokenNotFound
	}
	expireTime := time.Now().Add(-(2 * 24 * time.Hour))
	if user.AccessTime.Before(expireTime) && mode == "Production" {
		repo.UpdateExpirToken(tokenID)
		return nil, false, auth.ErrTokenExpired
	}

	if err != nil {
		return nil, false, err
	}
	user.Token = tokenID
	return &user, true, nil
}

func (repo *authRepository) GetToken(tokenID string, mode string) (*auth.Token, error) {
	uid, checkuid, err1 := repo.GetUserId(tokenID, mode)
	adminid, checkadmin, err2 := repo.GetTokenAdmin(tokenID, mode)
	if checkuid != true && checkadmin != true {
		if err1.Error() == "token_expired" || err2.Error() == "token_expired" {
			return nil, auth.ErrTokenExpired
		}
		if err1 == sql.ErrNoRows || err2 == sql.ErrNoRows {
			return nil, auth.ErrTokenNotFound
		}
		if err1 != nil || err2 != nil {
			if err2 != nil {
				return nil, err2
			} else {
				return nil, err1
			}
		}
		return nil, err1
	}
	logg.Println(adminid)
	if checkuid == false {
		user := auth.Token{
			UserID:    "",
			AdminCode: adminid.AdminCode,
			Token:     adminid.Token,
		}
		return &user, nil
	} else {
		user := auth.Token{
			UserID:    uid.UserID,
			AdminCode: "",
			Token:     uid.Token,
		}
		return &user, nil
	}
	return nil, nil
}

func GetYear() string {
	var intyear int
	if time.Now().Year() >= 2560 {
		intyear = time.Now().Year()
	} else {
		intyear = time.Now().Year() + 543
	}
	year1 := strconv.Itoa(intyear)
	year2 := year1[2:len(year1)]
	return year2
}

func GetMonth() string {

	var vmonth1 string
	mounth1 := int(time.Now().Month())
	mounth2 := int(mounth1)
	vmonth := strconv.Itoa(mounth2)

	if len(vmonth) == 1 {
		vmonth1 = "0" + vmonth
	} else {
		vmonth1 = vmonth
	}
	return vmonth1
}

func GetDay() string {
	day := time.Now().Day()
	days := strconv.Itoa(day)
	if len(days) == 1 {
		days = "0" + strconv.Itoa(day)
	}
	return days
}

func (repo *authRepository) FindUserTel(code string) (bool, error) {
	var user int = 0
	sql := `select count(id) as user from users where telephone = ? and active_status = 1`
	rs, _ := repo.db.QueryRow(sql, code)
	rs.Scan(&user)
	if user > 0 {
		return false, errors.New("หมายเลขนี้ได้ทำการลงทะเบียนแล้ว")
	}
	return true, nil
}

func (repo *authRepository) FindUserEmail(code string) (bool, error) {
	var user int = 0
	sql := `select count(id) as user from users where email = ? and active_status = 1`
	rs, _ := repo.db.QueryRow(sql, code)
	rs.Scan(&user)
	if user > 0 {
		return false, errors.New("email นี้ได้ทำการลงทะเบียนแล้ว")
	}
	return true, nil
}
func (repo *authRepository) FindInviteCode(code string) (bool, error) {
	var invite_code int = 0
	sql := `select count(invite_code) as invite_code from users where invite_code = ? and active_status = 1`
	rs, _ := repo.db.QueryRow(sql, code)
	rs.Scan(&invite_code)
	if invite_code > 0 {
		return false, errors.New("email นี้ได้ทำการลงทะเบียนแล้ว")
	}
	return true, nil
}

func (repo *authRepository) GetInviteCode() string {
	Invite, _ := uuid.NewV4()
	Inv := strings.Replace(Invite.String(), "-", "", -1)
	ch, _ := repo.FindInviteCode(Inv)
	if ch != true {
		Inv = repo.GetInviteCode()
	}
	return Inv
}

func (repo *authRepository) SignUpPhoneRepo(code string, refcode string, invcode string, password string) (interface{}, error) {
	_, err := repo.FindUserTel(code)
	if err != nil {
		return nil, err
	}
	// fmt.Println(1)
	iduser, _ := repo.FindUserID()
	// fmt.Println(1)
	userid, err := GenUserID(iduser)
	if err != nil {
		return nil, err
	}

	_, err = repo.dbx.NamedExec(
		`INSERT INTO
			users
				(user_id,pwd,create_time,telephone,ref_code,invite_code)
				VALUES
				(:user_id, :pwd,:create_time, :telephone, :ref_code, :invite_code)
				`,
		map[string]interface{}{
			"user_id":     userid,
			"pwd":         password,
			"create_time": time.Now(),
			"telephone":   code,
			"ref_code":    refcode,
			"invite_code": invcode,
		},
	)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return nil, nil
}

func (repo *authRepository) SignUpEmailRepo(code string, refcode string, invcode string, password string) (interface{}, error) {
	_, err := repo.FindUserEmail(code)
	if err != nil {
		return nil, err
	}
	iduser, _ := repo.FindUserID()
	userid, err := GenUserID(iduser)
	if err != nil {
		return nil, err
	}

	_, err = repo.dbx.NamedExec(
		`INSERT INTO
			users
				(user_id,pwd,create_time,email,ref_code,invite_code)
				VALUES
				(:user_id, :pwd,:create_time, :email, :ref_code,:invite_code)
				`,
		map[string]interface{}{
			"user_id":     userid,
			"pwd":         password,
			"create_time": time.Now(),
			"email":       code,
			"ref_code":    refcode,
			"invite_code": invcode,
		},
	)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return nil, nil
}

// func (repo *authRepository) GetToken(tokenID string) (*auth.Token, error) {
// 	var m struct {
// 		ClientID    sql.NullInt64  `db:"client_id"`
// 		AccountID   sql.NullInt64  `db:"account_id"`
// 		VendingID   sql.NullInt64  `db:"vending_id"`
// 		VendingUUID sql.NullString `db:"vending_uuid"`
// 		TokenID     sql.NullString `db:"token"`
// 		Created     time.Time      `db:"created"`
// 	}
// 	//fmt.Println("start repo *authRepository.GetToken")
// 	rows, err := repo.db.QueryRow(`
// 		select
// 			accounts.client_id, account_id,  null as vending_id, null as vending_uuid,accesses.id,accesses.created
// 			from accesses inner join accounts on accesses.account_id = accounts.id
// 			where accesses.id = $1
// 		union select
// 			client_id, null, vending_id, vending_uuid, token, created
// 		from vending_access
// 		where token = ?
// 	`, tokenID)
// 	err = rows.Scan(
// 		&m.ClientID, &m.AccountID, &m.VendingID, &m.VendingUUID, &m.TokenID, &m.Created,
// 	)
// 	if err == sql.ErrNoRows {
// 		return nil, auth.ErrTokenNotFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	expireTime := time.Now().Add(-(365 * 24 * time.Hour))
// 	if m.Created.Before(expireTime) {
// 		return nil, auth.ErrTokenExpired
// 	}

// 	tk := auth.Token{ID: tokenID}
// 	//fmt.Println("postgres.auth.go -> auth.Token.ID = ", tokenID)
// 	if m.ClientID.Valid {
// 		tk.ClientID = m.ClientID.Int64
// 	} else {
// 		tk.ClientID = -1
// 	}
// 	if m.AccountID.Valid {
// 		tk.AccountID = m.AccountID.Int64
// 	} else {
// 		tk.AccountID = -1
// 	}
// 	if m.VendingID.Valid {
// 		tk.VendingID = m.VendingID.Int64
// 	} else {
// 		tk.VendingID = -1
// 	}
// 	if m.VendingUUID.Valid {
// 		tk.VendingUUID = m.VendingUUID.String
// 	}
// 	if m.TokenID.Valid {
// 		tk.TokenID = m.TokenID.String
// 	}
// 	// sql := `select id,meta from partner_master where id = (select partner_id from clients where id = ? limit 1)`
// 	// sqls, err := repo.db.QueryRow(sql, tk.ClientID)
// 	// if err != nil {
// 	// 	fmt.Printf("error %v",err.Error())
// 	// 	return nil, err
// 	// }

// 	// mt := &auth.Partner{}
// 	// meta := auth.Meta{}
// 	// sqls.Scan(&meta.Id, &meta.Meta)
// 	// fmt.Println((meta.Meta))
// 	// err = json.Unmarshal([]byte(meta.Meta.String), &mt)
// 	// if err != nil {
// 	// 	fmt.Println("error unmarshal: ",err.Error())
// 	// 	return nil, err
// 	// }
// 	// var validate *validator.Validate
// 	// validate = validator.New()
// 	// err = validate.Struct(mt)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// end := auth.Endpoint {
// 	// 	URL:mt.Endpoint.URL,
// 	// 	Token:mt.Endpoint.Token,
// 	// }
// 	// tk.PartNer.Endpoint = end
// 	//fmt.Println("return postgres.auth.GetToken : ", tk)
// 	return &tk, nil
// }

// func (repo *authRepository) GetAccount(email string) (*auth.Account, error) {
// 	var account struct {
// 		ID           int64  `db:"id"`
// 		ClientID     int64  `db:"client_id"`
// 		FullName     string `db:"full_name"`
// 		Email        string `db:"email"`
// 		PasswordHash string `db:"password_hash"`
// 		Status       int    `db:"account_status"`
// 	}
// 	err := repo.db.
// 		SelectFrom("accounts").
// 		Where("email ilike ?", email).
// 		One(&account)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ac := auth.Account{
// 		ID:           account.ID,
// 		ClientID:     account.ClientID,
// 		FullName:     account.FullName,
// 		Email:        account.Email,
// 		PasswordHash: account.PasswordHash,
// 	}
// 	if account.Status == 1 {
// 		ac.Status = true
// 	}
// 	// fmt.Println("return postgres.auth.GetAccount")
// 	return &ac, nil
// }

// func (repo *authRepository) GetAccountClientReleID(accountID int64) (*auth.AccountClientRole, error) {
// 	sql := `select id,account_id,client_role_id from accounts_client_role where account_id = ?`
// 	pb, err := repo.db.QueryRow(sql, accountID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	accounts := auth.AccountClientRole{}
// 	err = pb.Scan(&accounts.ID, &accounts.AccountID, &accounts.ClientRoleID)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	ac := auth.AccountClientRole{
// 		ID:           accounts.ID,
// 		AccountID:    accounts.AccountID,
// 		ClientRoleID: accounts.ClientRoleID,
// 	}
// 	return &ac, nil
// }

// func (repo *authRepository) GetClientRole(roleid int64) (*auth.ClientRole, error) {
// 	sql := `select id,Name from client_role where id = ?`
// 	fmt.Println(roleid)
// 	pb, err := repo.db.QueryRow(sql, roleid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	accounts := auth.ClientRole{}
// 	err = pb.Scan(&accounts.ID, &accounts.Name)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	ac := auth.ClientRole{
// 		ID:   accounts.ID,
// 		Name: accounts.Name,
// 	}
// 	return &ac, nil
// }

// func (repo *authRepository) GetAccountV2(email string) (*auth.AccountV2, error) {
// 	var account struct {
// 		ID           int64  `db:"id"`
// 		ClientID     int64  `db:"client_id"`
// 		FullName     string `db:"full_name"`
// 		Email        string `db:"email"`
// 		PasswordHash string `db:"password_hash"`
// 		Status       int    `db:"account_status"`
// 	}
// 	err := repo.db.
// 		SelectFrom("accounts").
// 		Where("email ilike ?", email).
// 		One(&account)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ac := auth.AccountV2{
// 		ID:           account.ID,
// 		ClientID:     account.ClientID,
// 		FullName:     account.FullName,
// 		Email:        account.Email,
// 		PasswordHash: account.PasswordHash,
// 	}
// 	if account.Status == 1 {
// 		ac.Status = true
// 	}
// 	// fmt.Println("return postgres.auth.GetAccount")
// 	return &ac, nil
// }

// func (repo *authRepository) SaveAccess(acc *auth.Access) error {
// 	_, err := repo.db.
// 		InsertInto("accesses").
// 		Values(map[string]interface{}{
// 			"id":         acc.TokenID,
// 			"account_id": acc.AccountID,
// 			"client_id":  acc.ClientID,
// 		}).Exec()
// 	if err != nil {
// 		return err
// 	}
// 	// fmt.Println("return postgres.auth.GetAccount")
// 	return nil
// }

// func (repo *authRepository) SaveDeviceToken(accountID int64, token string) error {
// 	_, err := repo.db.
// 		InsertInto("devices").
// 		Values(map[string]interface{}{
// 			"fcm_token":  token,
// 			"account_id": accountID,
// 		}).Exec()
// 	if err != nil {
// 		return err
// 	}
// 	// fmt.Println("return postgres.auth.SaveDeviceToken")
// 	return nil
// }

// func (repo *authRepository) DeleteAccess(token string) error {
// 	_, err := repo.db.
// 		DeleteFrom("accesses").
// 		Where("id ilike ?", token).
// 		Exec()
// 	if err != nil {
// 		return err
// 	}
// 	// fmt.Println("return postgres.auth.DeleteAccess")
// 	return nil
// }

// func (repo *authRepository) GetClientMQTTKey(clientID int64) (*string, error) {
// 	var m struct {
// 		MQTTKey sql.NullString `db:"mqtt_key"`
// 	}

// 	err := repo.db.SelectFrom("clients").
// 		Where("id = ?", clientID).
// 		One(&m)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var key *string
// 	if m.MQTTKey.Valid || m.MQTTKey.String != "" {
// 		s := m.MQTTKey.String
// 		key = &s
// 	}
// 	return key, nil
// }

// func (repo *authRepository) GetProfile(token string) (*auth.Profile, error) {
// 	var ac struct {
// 		ID         int64          `db:"id"`
// 		FullName   string         `db:"full_name"`
// 		Email      string         `db:"email"`
// 		PictureURL sql.NullString `db:"picture_url"`
// 		Status     int            `db:"account_status"`
// 	}
// 	rows, err := repo.db.QueryRow(`
// 		select
// 			accounts.id, accounts.full_name, accounts.email, accounts.picture_url, accounts.account_status
// 			from accounts inner join accesses on accounts.id = accesses.account_id
// 			where accesses.id = ?
// 	`, token)
// 	err = rows.Scan(
// 		&ac.ID, &ac.FullName, &ac.Email, &ac.PictureURL, &ac.Status,
// 	)
// 	if err == sql.ErrNoRows {
// 		return nil, auth.ErrTokenNotFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	p := auth.Profile{
// 		AccountID: ac.ID,
// 		FullName:  ac.FullName,
// 		Email:     ac.Email,
// 		Status:    ac.Status,
// 	}
// 	if ac.PictureURL.Valid {
// 		p.PictureURL = &ac.PictureURL.String
// 	}
// 	return &p, nil
// }
