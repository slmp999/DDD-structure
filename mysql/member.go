package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	//"github.com/labstack/gommon/log"
	"github.com/nyaruka/phonenumbers"
	log "gitlab.com/satit13/perfect_api/logger"
	logg "gitlab.com/satit13/perfect_api/logger"
	"gitlab.com/satit13/perfect_api/member"
	"gitlab.com/satit13/perfect_api/sales"
	"golang.org/x/crypto/bcrypt"
	"upper.io/db.v3/lib/sqlbuilder"

	//"upper.io/db.v3/postgresql"
	"upper.io/db.v3/mysql"
)

// NewAuthRepository creates new auth repository
func NewMemberRepository(db *sql.DB) (member.Repository, error) {
	// fmt.Println(1)
	pdb, err := mysql.New(db)
	if err != nil {
		return nil, err
	}
	dbx := sqlx.NewDb(db, "mysql")
	// fmt.Println(1)
	r := memberRepository{pdb, dbx}
	// fmt.Println(1)
	return &r, nil
}

type memberRepository struct {
	db  sqlbuilder.Database
	dbx *sqlx.DB
}

func hashAndSalt(pwd string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), 8)

	if err != nil {
		fmt.Println(err)
	}

	return string(hashedPassword)
}

func ConvertPhonetoCoutryCode(code string) string {
	defaultRegion := "TH"
	info, _ := phonenumbers.Parse(code, defaultRegion)
	formattedNum := phonenumbers.Format(info, phonenumbers.NANPA_COUNTRY_CODE)
	phone := strings.Replace(formattedNum, " ", "", -1)
	return phone
}
func (r *memberRepository) GetAddressByIdRepo(UserID string, AddressID int64) (*member.AddressProfileModel, error) {
	sql1 := `
	select 
		ifnull(id,0) as addr_id,
		ifnull(name,'') as addr_name,
		ifnull(addr_phone,'') as addr_phone,
		ifnull(addr_email,'') as addr_email,
		ifnull(addr_state,'') as addr_state,
		ifnull(addr_subarea,'') as addr_subarea,
		ifnull(addr_district,'') as addr_district,
		ifnull(addr_province,'') as addr_province,
		ifnull(addr_postal_code,'') as addr_postal_code,
		ifnull(main_address,0) as main_address from user_address where ref_id = ? and id = ? and status = 0
	`
	rs1, err := r.db.QueryRow(sql1, UserID, AddressID)
	if err != nil {
		fmt.Println(err.Error())
	}
	addr := member.AddressProfileModel{}

	err = rs1.Scan(&addr.AddressID, &addr.Name, &addr.Phone, &addr.Email, &addr.Address, &addr.SubArea, &addr.District, &addr.Province, &addr.PostalCode, &addr.MainAddress)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	return &addr, nil
}

func (r *memberRepository) GetAddressByUserRepo(UserID string) (*[]member.AddressProfileModel, error) {
	sql1 := `
	select 
		ifnull(id,0) as addr_id,
		ifnull(name,'') as addr_name,
		ifnull(addr_phone,'') as addr_phone,
		ifnull(addr_email,'') as addr_email,
		ifnull(addr_state,'') as addr_state,
		ifnull(addr_subarea,'') as addr_subarea,
		ifnull(addr_district,'') as addr_district,
		ifnull(addr_province,'') as addr_province,
		ifnull(addr_postal_code,'') as addr_postal_code,
		ifnull(main_address,0) as main_address from user_address where ref_id = ? and status = 0
	`
	rs1, err := r.db.Query(sql1, UserID)
	if err != nil {
		fmt.Println(err.Error())
	}
	addrs := []member.AddressProfileModel{}
	for rs1.Next() {
		addr := member.AddressProfileModel{}
		err = rs1.Scan(&addr.AddressID, &addr.Name, &addr.Phone, &addr.Email, &addr.Address, &addr.SubArea, &addr.District, &addr.Province, &addr.PostalCode, &addr.MainAddress)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, addr)
	}

	return &addrs, nil
}

func (r *memberRepository) GetCouponByUserPercent(UserID string) (*member.ModelCoupon, error) {
	sql1 := `SELECT 
	ifnull(id,0) as id,
	ifnull(doc_date,'') as doc_date,
	ifnull(coupon_no,'') as coupon_no,
	ifnull(coupon_type,0) as coupon_type,
	ifnull(user_id,'') as user_id,
	ifnull(name,'') as name,
	ifnull(value,0) as value,
	ifnull(remain,0) as remain,
	ifnull(expire_status,0) as expire_status,
	ifnull(begin_date,'') as begin_date,
	ifnull(expire_date,'') as expire_date
	from coupon where user_id = ? and coupon_type = 2 order by value desc limit 1
	`
	rs, err := r.db.QueryRow(sql1, UserID)
	if err != nil {
		return nil, err
	}

	coupons := member.ModelCoupon{}
	err = rs.Scan(
		&coupons.ID,
		&coupons.DocDate,
		&coupons.CouponNo,
		&coupons.CouponType,
		&coupons.UserID,
		&coupons.Name,
		&coupons.Value,
		&coupons.Remain,
		&coupons.ExpireStatus,
		&coupons.BeginDate,
		&coupons.ExpireDate)
	if err != nil {
		return nil, err
	}
	return &coupons, nil
}

func (r *memberRepository) GetCouponUsers(UserID string) ([]member.ModelCoupon, error) {
	sql1 := `SELECT 
	ifnull(id,0) as id,
	ifnull(doc_date,'') as doc_date,
	ifnull(coupon_no,'') as coupon_no,
	ifnull(coupon_type,0) as coupon_type,
	ifnull(user_id,'') as user_id,
	ifnull(name,'') as name,
	ifnull(value,0) as value,
	ifnull(remain,0) as remain,
	ifnull(expire_status,0) as expire_status,
	ifnull(begin_date,'') as begin_date,
	ifnull(expire_date,'') as expire_date
	from coupon where user_id = ? order by  create_time  desc
	`
	rs, err := r.db.Query(sql1, UserID)
	if err != nil {
		return nil, err
	}

	coupons := []member.ModelCoupon{}
	for rs.Next() {
		coupon := member.ModelCoupon{}
		err = rs.Scan(
			&coupon.ID,
			&coupon.DocDate,
			&coupon.CouponNo,
			&coupon.CouponType,
			&coupon.UserID,
			&coupon.Name,
			&coupon.Value,
			&coupon.Remain,
			&coupon.ExpireStatus,
			&coupon.BeginDate,
			&coupon.ExpireDate)
		if err != nil {
			return nil, err
		}
		coupons = append(coupons, coupon)
	}
	logg.Println(coupons)
	return coupons, nil
}

func (r *memberRepository) FindFirstPay(userID string) (int64, error) {
	var user_id int64 = 0
	sql1 := `select count(user_id) as user_id 
	from orders where user_id = ? and order_status > 1`
	rs, err := r.db.QueryRow(sql1, userID)
	if err != nil {
		return 0, err
	}
	rs.Scan(&user_id)
	if user_id > 0 {
		return 1, nil
	}
	return 0, nil
}

func (r *memberRepository) RemoveMember(adminID string, userID string, status int64) (interface{}, error) {
	_, err := r.db.Exec(`update users set active_status = ?,
	edit_by = ?,edit_time = ? 
	where user_id = ?`,
		status, adminID, time.Now(), userID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *memberRepository) RemoveUserByAdmin(adminID string, userID string) (interface{}, error) {
	_, err := r.RemoveMember(adminID, userID, 0)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (r *memberRepository) GetStatusSales(userID string) (int64, error) {
	var user int = 0
	var confirm int64 = 0
	sql := `select count(user_code) as user from sales_person where user_code = ? and active_status = 0`
	rs, _ := r.db.QueryRow(sql, userID)
	rs.Scan(&user)
	if user > 0 {
		sql2 := `select confirm from sales_person where user_code =  ? and active_status = 0 limit 1`
		rs2, _ := r.db.QueryRow(sql2, userID)
		rs2.Scan(&confirm)
		if confirm == 0 {
			return 1, nil
		} else {
			return 2, nil
		}
	}
	return 0, nil
}

func (r *memberRepository) GetUserByAdmin(userID string) (*member.UserProfileModelBackEnd, error) {
	sqll := `select 
	ifnull(a.company_id,0) as company_id,
	ifnull(a.user_id,'') as user_id,
	ifnull(a.role_id,0) as role_id,
	ifnull(b.role_name,'') as role_name,
	ifnull(a.dealer,0) as dealer,
	ifnull(a.profix_name,'') as profix_name,
	ifnull(a.user_fname,'') as user_fname,
	ifnull(a.user_lname,'') as user_lname,
	ifnull(a.telephone,'') as telephone,
	ifnull(a.active_status,0) as active_status,
	ifnull(a.email,'') as email,
	ifnull(a.ref_code,'') as ref_code,
	ifnull(a.invite_code,'') as invite_code,
	ifnull(DATE_FORMAT(a.create_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as create_date
	 from users a 
	left join user_role b on a.role_id = b.id
	 where a.user_id = ?  limit 1`
	rs, err := r.db.QueryRow(sqll, userID)
	if err != nil {
		return nil, err
	}
	user := member.UserProfileModelBackEnd{}
	err = rs.Scan(&user.CompanyID, &user.UserID, &user.RoleID, &user.RoleName, &user.Dealer, &user.ProfixName, &user.Fname, &user.Lname, &user.Telephone, &user.ActiveStatus,
		&user.Email, &user.RefCode, &user.InviteCode, &user.CreateDate)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *memberRepository) UpdateMemberByadmin(user *member.UpdateUserProfileModelByAdmin, AdminID string) (interface{}, error) {
	ch, err := r.GetEmailAndPhone(user.UserID)
	logg.Println(ch)
	if ch.Email != user.Email && user.Email != "" {
		_, err := r.FindUserEmail(user.Email)
		if err != nil {
			return nil, err
		}
	}
	user.Telephone = ConvertPhonetoCoutryCode(user.Telephone)
	if ch.Telephone != user.Telephone && user.Telephone != "" {

		_, err := r.FindUserTel(user.Telephone)
		if err != nil {
			return nil, err
		}
	}

	_, err = r.db.Exec(`update users
					set company_id = ?,profix_name = ?,user_fname = ?,user_lname = ?,
					active_status = ?,telephone = ?,email = ?,edit_time = ?
					where user_id = ?`,
		user.CompanyID,
		user.ProfixName,
		user.Fname,
		user.Lname,
		user.ActiveStatus,
		user.Telephone,
		user.Email,
		time.Now(),
		user.UserID,
	)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *memberRepository) GeTProfileMemberRepo(UserID string) (*member.UserProfileModel, error) {
	sqll := `select 
	ifnull(a.company_id,0) as company_id,
	ifnull(a.role_id,0) as role_id,
	ifnull(b.role_name,'') as role_name,
	ifnull(a.dealer,0) as dealer,
	ifnull(a.profix_name,'') as profix_name,
	ifnull(a.user_fname,'') as user_fname,
	ifnull(a.user_lname,'') as user_lname,
	ifnull(a.telephone,'') as telephone,
	ifnull(a.active_status,0) as active_status,
	ifnull(a.email,'') as email,
	ifnull(a.ref_code,'') as ref_code,
	ifnull(a.invite_code,'') as invite_code,
	ifnull(DATE_FORMAT(a.create_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as create_date
	 from users a 
	left join user_role b on a.role_id = b.id
	 where a.user_id = ? and a.active_status = 1 limit 1`
	rs, err := r.db.QueryRow(sqll, UserID)
	if err != nil {
		return nil, err
	}
	user := member.UserProfileModel{}
	user.FirstBuy, _ = r.FindFirstPay(UserID)
	coupon, err := r.GetCouponByUserPercent(UserID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			user.DicountPerson = fmt.Sprintf("%.0f", 0.00) + "%"
			user.DiscountExpirDate = ""
		}
	} else {

		user.DicountPerson = fmt.Sprintf("%.0f", coupon.Value) + "%"
		user.DiscountExpirDate = coupon.ExpireDate
	}

	rs.Scan(&user.CompanyID, &user.RoleID, &user.RoleName, &user.Dealer, &user.ProfixName, &user.Fname, &user.Lname, &user.Telephone, &user.ActiveStatus,
		&user.Email, &user.RefCode, &user.InviteCode, &user.CreateDate)
	user.SalesStatus, _ = r.GetStatusSales(UserID)
	user.Coupon, err = r.GetCouponUsers(UserID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	user.InvitePerson, err = r.FindInviteCodeCount(user.InviteCode)
	if err != nil {
		return nil, err
	}
	if user.InvitePerson > 0 && user.RoleID == 2 {
		user.Invite, err = r.FindUerInviteCode(user.InviteCode, 1)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(1)
	// fmt.Println(1)
	sql1 := `select 
		ifnull(id,0) as addr_id,
		ifnull(name,'') as addr_name,
		ifnull(addr_phone,'') as addr_phone,
		ifnull(addr_email,'') as addr_email,
		ifnull(addr_state,'') as addr_state,
		ifnull(addr_subarea,'') as addr_subarea,
		ifnull(addr_district,'') as addr_district,
		ifnull(addr_province,'') as addr_province,
		ifnull(addr_postal_code,'') as addr_postal_code,
		ifnull(main_address,0) as main_address from user_address where ref_id = ? and status = 0
	`
	rs1, err := r.db.Query(sql1, UserID)
	if err != nil {
		fmt.Println(err.Error())
	}
	addrs := []member.AddressProfileModel{}
	for rs1.Next() {
		addr := member.AddressProfileModel{}
		err = rs1.Scan(&addr.AddressID, &addr.Name, &addr.Phone, &addr.Email, &addr.Address, &addr.SubArea, &addr.District, &addr.Province, &addr.PostalCode, &addr.MainAddress)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, addr)
	}
	user.Address = addrs
	menu, err := r.GetMenuPermision(UserID, user.RoleID)
	if err != nil {
		return nil, err
	}
	user.Menu = menu
	return &user, nil
}
func (r *memberRepository) GetNumberInvitePersonRefCode(ref_code string) (int64, error) {
	var invite_person int64 = 10
	sql1 := `select ifnull(invite_person,10) as invite_person  from users where invite_code = ? limit 1`
	rs, _ := r.db.QueryRow(sql1, ref_code)
	rs.Scan(&invite_person)
	if invite_person <= 0 {
		return 10, nil
	}
	return invite_person, nil
}

func (r *memberRepository) FindUerInviteCode(ref_code string, level int64) ([]sales.ModelStructureSales, error) {
	Limit, _ := r.GetNumberInvitePersonRefCode(ref_code)
	sql1 := `select 
	ifnull(a.company_id,0) as company_id,
	ifnull(a.role_id,0) as role_id,
	ifnull(b.role_name,'') as role_name,
	ifnull(a.profix_name,'') as profix_name,
	ifnull(a.user_fname,'') as user_fname,
	ifnull(a.user_lname,'') as user_lname,
	ifnull(a.telephone,'') as telephone,
	ifnull(a.active_status,0) as active_status,
	ifnull(a.email,'') as email,
	ifnull(a.ref_code,'') as ref_code,
	ifnull(a.invite_code,'') as invite_code
	 from users a 
	left join user_role b on a.role_id = b.id
	 where  a.ref_code = ? and a.active_status = 1 order by a.create_time asc limit ` + fmt.Sprintln(Limit)
	rs, err := r.db.Query(sql1, ref_code)
	if err != nil {
		return nil, err
	}
	users := []sales.ModelStructureSales{}
	for rs.Next() {
		user := sales.ModelStructureSales{}
		rs.Scan(&user.CompanyID, &user.RoleID, &user.RoleName, &user.ProfixName, &user.Fname, &user.Lname, &user.Telephone, &user.ActiveStatus,
			&user.Email, &user.RefCode, &user.InviteCode)
		user.InvitePerson, _ = r.FindInviteCodeCount(user.InviteCode)
		if level <= 5 && user.InvitePerson > 0 && user.RoleID == 2 {
			user.Invite, _ = r.FindUerInviteCode(user.InviteCode, level+1)
		}
		users = append(users, user)

	}
	return users, nil

}

func (r *memberRepository) FindInviteCodeCount(inviteCode string) (int64, error) {

	var ref_code int64 = 0
	sql := `select count(ref_code) as ref_code from users where ref_code = ? and active_status = 1`
	rs, _ := r.db.QueryRow(sql, inviteCode)
	rs.Scan(&ref_code)
	if ref_code <= 0 {
		return 0, nil
	}
	return ref_code, nil
}

func (r *memberRepository) AddProfileAddRepo(auth *member.AddAddressModel) (interface{}, error) {
	if auth.Address.MainAddress == 1 {
		_, err := r.UpdateMainAddress(auth.UserID)
		if err != nil {
			return nil, err
		}
	}
	_, err := r.dbx.NamedExec(
		`INSERT INTO
			user_address
				(ref_id,
					status,
					name,
					addr_phone,
					addr_email,
					addr_state,
					addr_subarea,
					addr_district,
					addr_province,
					addr_postal_code,
					main_address,
					create_time)
				VALUES
				(	:ref_id,
					:status,
					:name,
					:addr_phone,
					:addr_email,
					:addr_state,
					:addr_subarea,
					:addr_district,
					:addr_province,
					:addr_postal_code,
					:main_address,
					:create_time)
				`,
		map[string]interface{}{
			"ref_id":           auth.UserID,
			"status":           0,
			"name":             auth.Address.Name,
			"addr_phone":       auth.Address.Phone,
			"addr_email":       auth.Address.Email,
			"addr_state":       auth.Address.Address,
			"addr_subarea":     auth.Address.SubArea,
			"addr_district":    auth.Address.District,
			"addr_province":    auth.Address.Province,
			"addr_postal_code": auth.Address.PostalCode,
			"main_address":     auth.Address.MainAddress,
			"create_time":      time.Now(),
		},
	)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *memberRepository) GetCouponByNo(userID string, CouponNo string) (*member.ModelCoupon, error) {
	sql1 := `SELECT 
	ifnull(id,0) as id,
	ifnull(doc_date,'') as doc_date,
	ifnull(coupon_no,'') as coupon_no,
	ifnull(coupon_type,0) as coupon_type,
	ifnull(user_id,'') as user_id,
	ifnull(name,'') as name,
	ifnull(value,0) as value,
	ifnull(remain,0) as remain,
	ifnull(expire_status,0) as expire_status,
	ifnull(begin_date,'') as begin_date,
	ifnull(expire_date,'') as expire_date
	from coupon where user_id = ? and coupon_no = ? limit 1
	`
	rs, err := r.db.QueryRow(sql1, userID, CouponNo)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ท่านไม่สามารถใช้ coupon นี้ได้")
		}
		return nil, err
	}

	coupons := member.ModelCoupon{}
	err = rs.Scan(
		&coupons.ID,
		&coupons.DocDate,
		&coupons.CouponNo,
		&coupons.CouponType,
		&coupons.UserID,
		&coupons.Name,
		&coupons.Value,
		&coupons.Remain,
		&coupons.ExpireStatus,
		&coupons.BeginDate,
		&coupons.ExpireDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ท่านไม่สามารถใช้ coupon นี้ได้")
		}
		return nil, err
	}
	return &coupons, nil
}

func (r *memberRepository) CheckCouponRepo(userID, couponNO string) (interface{}, error) {
	coupon, err := r.GetCouponByNo(userID, couponNO)
	if err != nil {
		return nil, err
	}
	_, err = CheckExpireDate(coupon.ExpireDate)
	if err != nil {
		return nil, err
	}
	return &coupon, nil
}

func CheckExpireDate(expdate string) (bool, error) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().Format("2006-01-02")

	layOut := "2006-01-02"
	nowfomat, _ := time.Parse(layOut, now)
	dateStamp, _ := time.Parse(layOut, expdate)
	log.Println(nowfomat, dateStamp)
	if dateStamp.In(loc).Before(nowfomat) {
		return false, errors.New("coupon หมดอายุการใช้งาน")
	}
	return true, nil
}

func (r *memberRepository) UpdateMainAddress(code string) (interface{}, error) {
	_, err := r.db.Exec(`update user_address set main_address = 0 where ref_id = ?`,
		code)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *memberRepository) UpdateProFileAddrRepo(auth *member.AddAddressModel) (interface{}, error) {
	if auth.Address.MainAddress == 1 {
		_, err := r.UpdateMainAddress(auth.UserID)
		if err != nil {
			return nil, err
		}
	}
	_, err := r.db.Exec(`update user_address
					set status = ?,name = ?,addr_phone = ?,addr_email = ?,addr_state = ?,addr_subarea = ?,
					addr_district = ?,addr_province = ?,addr_postal_code = ?,main_address = ?,edit_time = ?
					where ref_id=? and id = ?`,
		0,
		auth.Address.Name,
		auth.Address.Phone,
		auth.Address.Email,
		auth.Address.Address,
		auth.Address.SubArea,
		auth.Address.District,
		auth.Address.Province,
		auth.Address.PostalCode,
		auth.Address.MainAddress,
		time.Now(),
		auth.UserID,
		auth.Address.AddressID,
	)
	log.Println(auth)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *memberRepository) DeleteProfileAddressRepo(userID string, addressID int64) (interface{}, error) {
	_, err := r.db.Exec(`update user_address set status = ? where ref_id = ? and id = ?`, 9, userID, addressID)
	fmt.Println(userID, addressID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *memberRepository) GetEmailAndPhone(code string) (*member.UpdateUserProfileModel, error) {

	sqls := `select ifnull(telephone,'') as telephone,
	ifnull(email,'') as email from users where user_id = ?  and active_status = 1 limit 1`
	rs, err := r.db.QueryRow(sqls, code)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	us := member.UpdateUserProfileModel{}
	rs.Scan(&us.Telephone, &us.Email)
	return &us, nil
}

func (repo *memberRepository) FindUserTel(code string) (bool, error) {
	var user int = 0
	sql := `select count(id) as user from users where telephone = ? and active_status = 1`
	rs, _ := repo.db.QueryRow(sql, code)
	rs.Scan(&user)
	if user > 0 {
		return false, errors.New("เบอร์โทรศัพนี้ได้ลงทะเบียนไว้แล้วหากท่านต้องการเปลี่ยนเบอร์หรือนำ Account นี้กลับมาใช้กรุณาลบ accout ที่มีเบอร์นี้ผูกอยู่")
	}
	return true, nil
}

func (repo *memberRepository) FindUserEmail(code string) (bool, error) {
	var user int = 0
	sql := `select count(id) as user from users where email = ? and active_status = 1`
	rs, _ := repo.db.QueryRow(sql, code)
	rs.Scan(&user)
	if user > 0 {
		return false, errors.New("Emailนี้ได้ลงทะเบียนไว้แล้วหากท่านต้องการเปลี่ยนEmail หรือนำ Account นี้กลับมาใช้กรุณาลบ accout ที่มีEmail นี้ผูกอยู่")
	}
	return true, nil
}

func (r *memberRepository) GetMenuPermision(UserID string, role int64) ([]member.HeadMenuModel, error) {

	sq := `SELECT head_menu_id, 
	ifnull(head_menu_name,'') as head_menu_name,
		ifnull(head_menu_icon,'') as head_menu_icon, 
		ifnull(head_menu_line_number,0) as head_menu_line_number
	from bn_head_menu order by head_menu_line_number asc
`
	rs, err := r.db.Query(sq)
	if err != nil {
		return nil, err
	}
	heads := []member.HeadMenuModel{}
	for rs.Next() {
		head := member.HeadMenuModel{}
		rs.Scan(&head.HeadMenuID,
			&head.HeadManuName,
			&head.HeadMenuIcon,
			&head.HeadMenuLineNumber)
		sql1 := `SELECT 
		ifnull(a.menu_id,'') as menu_id,
		ifnull(a.head_menu_id,'') as head_menu_id,
		ifnull(a.menu_name,'') as menu_name,
		ifnull(a.menu_link,'') as menu_link,
		ifnull(a.menu_show,0) as menu_show,
		ifnull(a.menu_status,0) as menu_status, 
		ifnull(a.menu_line_number,1) as menu_line_number
		from bn_menu a 
		left JOIN bn_map_menu b on a.menu_id = b.menu_id 
		where a.head_menu_id = ? and b.role_id = ? order by a.menu_line_number asc`
		r2, err := r.db.Query(sql1, head.HeadMenuID, role)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		}
		subs := []member.SubMenuModel{}
		for r2.Next() {
			sub := member.SubMenuModel{}
			r2.Scan(
				&sub.MenuID,
				&sub.HeadMenuID,
				&sub.MenuName,
				&sub.MenuLink,
				&sub.MenuShow,
				&sub.MenuStatus,
				&sub.MenuLineNumber)
			subs = append(subs, sub)
		}
		head.Sub = subs
		heads = append(heads, head)
	}

	return heads, nil
}

func (r *memberRepository) UpdatePRofileRepo(user *member.UpdateUserProfileModel, userID string) (interface{}, error) {
	ch, err := r.GetEmailAndPhone(userID)
	if ch.Email != user.Email && user.Email != "" {
		_, err := r.FindUserEmail(user.Email)
		if err != nil {
			return nil, err
		}
	}
	user.Telephone = ConvertPhonetoCoutryCode(user.Telephone)
	if ch.Telephone != user.Telephone && user.Telephone != "" {
		_, err := r.FindUserTel(user.Telephone)
		if err != nil {
			return nil, err
		}
	}
	_, err = r.db.Exec(`update users
					set company_id = ?,profix_name = ?,user_fname = ?,user_lname = ?,
					active_status = ?,telephone = ?,email = ?,edit_time = ?
					where user_id = ?`,
		user.CompanyID,
		user.ProfixName,
		user.Fname,
		user.Lname,
		user.ActiveStatus,
		user.Telephone,
		user.Email,
		time.Now(),
		userID,
	)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *memberRepository) GetAllMember() (interface{}, error) {
	sqll := `select 
	ifnull(a.company_id,0) as company_id,
	ifnull(a.user_id,'') as user_id,
	ifnull(a.role_id,0) as role_id,
	ifnull(b.role_name,'') as role_name,
	ifnull(a.dealer,0) as dealer,
	ifnull(a.profix_name,'') as profix_name,
	ifnull(a.user_fname,'') as user_fname,
	ifnull(a.user_lname,'') as user_lname,
	ifnull(a.telephone,'') as telephone,
	ifnull(a.active_status,0) as active_status,
	ifnull(a.email,'') as email,
	ifnull(a.ref_code,'') as ref_code,
	ifnull(a.invite_code,'') as invite_code,
	ifnull(DATE_FORMAT(a.create_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as create_date
	 from users a 
	left join user_role b on a.role_id = b.id`
	rs, err := r.db.Query(sqll)
	if err != nil {
		return nil, err
	}
	Users := []member.UserProfileModelBackEnd{}

	// user.FirstBuy, _ = r.FindFirstPay(UserID)
	// coupon, err := r.GetCouponByUserPercent(UserID)
	// if err != nil {
	// 	if err != sql.ErrNoRows {
	// 		return nil, err
	// 	} else {
	// 		user.DicountPerson = fmt.Sprintf("%.0f", 0.00) + "%"
	// 		user.DiscountExpirDate = ""
	// 	}
	// } else {

	// 	user.DicountPerson = fmt.Sprintf("%.0f", coupon.Value) + "%"
	// 	user.DiscountExpirDate = coupon.ExpireDate
	// }

	for rs.Next() {
		user := member.UserProfileModelBackEnd{}
		rs.Scan(&user.CompanyID, &user.UserID, &user.RoleID, &user.RoleName,
			&user.Dealer, &user.ProfixName, &user.Fname,
			&user.Lname, &user.Telephone, &user.ActiveStatus,
			&user.Email, &user.RefCode, &user.InviteCode, &user.CreateDate)
		fmt.Println("user : ", user)
		Users = append(Users, user)
	}
	fmt.Println(Users)
	// user.SalesStatus, _ = r.GetStatusSales(UserID)
	// user.Coupon, err = r.GetCouponUsers(UserID)
	// if err != nil && err != sql.ErrNoRows {
	// 	return nil, err
	// }
	// user.InvitePerson, err = r.FindInviteCodeCount(user.InviteCode)
	// if err != nil {
	// 	return nil, err
	// }
	// if user.InvitePerson > 0 && user.RoleID == 2 {
	// 	user.Invite, err = r.FindUerInviteCode(user.InviteCode, 1)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// fmt.Println(1)
	// sql1 := `select
	// 	ifnull(id,0) as addr_id,
	// 	ifnull(name,'') as addr_name,
	// 	ifnull(addr_phone,'') as addr_phone,
	// 	ifnull(addr_email,'') as addr_email,
	// 	ifnull(addr_state,'') as addr_state,
	// 	ifnull(addr_subarea,'') as addr_subarea,
	// 	ifnull(addr_district,'') as addr_district,
	// 	ifnull(addr_province,'') as addr_province,
	// 	ifnull(addr_postal_code,'') as addr_postal_code,
	// 	ifnull(main_address,0) as main_address from user_address where ref_id = ? and status = 0
	// `
	// rs1, err := r.db.Query(sql1, UserID)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// addrs := []member.AddressProfileModel{}
	// for rs1.Next() {
	// 	addr := member.AddressProfileModel{}
	// 	err = rs1.Scan(&addr.AddressID, &addr.Name, &addr.Phone, &addr.Email, &addr.Address, &addr.SubArea, &addr.District, &addr.Province, &addr.PostalCode, &addr.MainAddress)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	addrs = append(addrs, addr)
	// }
	// user.Address = addrs
	// menu, err := r.GetMenuPermision(UserID, user.RoleID)
	// if err != nil {
	// 	return nil, err
	// }
	// user.Menu = menu
	return &Users, nil
}
