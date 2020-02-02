package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"gitlab.com/satit13/perfect_api/admin"
	logg "gitlab.com/satit13/perfect_api/logger"

	//"github.com/labstack/gommon/log"

	//log "gitlab.com/satit13/perfect_api/logger"

	"upper.io/db.v3/lib/sqlbuilder"

	//"upper.io/db.v3/postgresql"
	"upper.io/db.v3/mysql"
)

// NewAuthRepository creates new auth repository
func NewAdminRepository(db *sql.DB) (admin.Repository, error) {
	// fmt.Println(1)
	pdb, err := mysql.New(db)
	if err != nil {
		return nil, err
	}
	dbx := sqlx.NewDb(db, "mysql")
	// fmt.Println(1)
	r := adminRepository{pdb, dbx}
	// fmt.Println(1)
	return &r, nil
}

type adminRepository struct {
	db  sqlbuilder.Database
	dbx *sqlx.DB
}

func (r *adminRepository) RemoveAdmin(adminID string, status int64, adminCode string) (interface{}, error) {
	if status == 1 {
		ch, err := r.GetEmail(adminID)
		if err != nil {
			return nil, err
		}
		if ch.ActiveStatus != 1 {
			_, err = r.FindUserEmail(ch.Email)
			if err != nil {
				return nil, err
			}
		}
	}
	_, err := r.db.Exec(`update users_admin set active_status = ?,
	edit_by = ?,edit_time = ? 
	where admin_code = ?`,
		status, adminID, time.Now(), adminCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *adminRepository) AddadminUser(adminID string, password string, user *admin.AddadminBackEndModel) (interface{}, error) {
	iduser, _ := r.FindUserID()
	admincode, err := GetAdminID(iduser)
	if err != nil {
		return nil, err
	}
	q := r.db.
		InsertInto("users_admin").
		Values(map[string]interface{}{
			"company_id":  0,
			"role_id":     user.RoleID,
			"admin_code":  admincode,
			"pwd":         password,
			"profix_name": user.ProfixName,
			"user_fname":  user.Fname,
			"user_lname":  user.Lname,
			"phone":       user.Phone,
			"email":       user.Email,
			"create_by":   adminID,
			"create_time": time.Now(),
		})
	_, err = q.Exec()
	if err != nil {
		return nil, err
	}
	Account := map[string]interface{}{
		"admin_code": admincode,
		"pre_fix":    user.ProfixName,
		"f_name":     user.Fname,
		"l_name":     user.Lname,
		"email":      user.Email,
		"password":   user.Password,
		"role_id":    user.RoleID,
	}

	return Account, nil
}

func (r *adminRepository) FindAdminBack(email string) (bool, error) {
	var user int = 0
	sql := `select count(id) as user from users_admin where email = ?`
	rs, _ := r.db.QueryRow(sql, email)
	rs.Scan(&user)
	if user > 0 {
		return false, errors.New("email นี้ได้ทำการลงทะเบียน กับผู้ใช้อื่นแล้ว")
	}
	return true, nil
}
func (r *adminRepository) SigupUserBack(adminID, email, pass string, roleID int64) (interface{}, error) {
	iduser, _ := r.FindUserID()
	admincode, err := GetAdminID(iduser)
	if err != nil {
		return nil, err
	}
	_, err = r.dbx.NamedExec(
		`INSERT INTO
			users_admin
				(role_id,admin_code,pwd,create_time,create_by,email)
				VALUES
				(:role_id,:admin_code, :pwd,:create_time,:create_by, :email)
				`,
		map[string]interface{}{
			"role_id":     roleID,
			"admin_code":  admincode,
			"pwd":         pass,
			"create_time": time.Now(),
			"create_by":   adminID,
			"email":       email,
		},
	)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *adminRepository) GetEmail(adminID string) (*admin.UpdateUserProfileModel, error) {

	sqls := `select
	ifnull(email,'') as email,ifnull(active_status,0) as active_status from users_admin where admin_code = ? and active_status = 1 limit 1`
	rs, err := r.db.QueryRow(sqls, adminID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	us := admin.UpdateUserProfileModel{}
	err = rs.Scan(&us.Email, &us.ActiveStatus)
	if err != nil {
		return nil, err
	}
	return &us, nil
}

func (repo *adminRepository) UpdatePassword(user *admin.UpdateUserProfileModel, pass string, adminID string) (interface{}, error) {
	_, err := repo.db.Exec(`update users_admin set pwd = ?,edit_time = ?,edit_by =? where admin_code = ?`,
		pass, time.Now(), adminID, user.AdminID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (repo *adminRepository) UpdateProfileBackEndRepo(user *admin.UpdateUserProfileModel, adminID string) (interface{}, error) {
	ch, err := repo.GetEmail(user.AdminID)
	fmt.Println(ch)
	fmt.Println(user)
	if ch.Email != user.Email && user.Email != "" {
		_, err := repo.FindUserEmail(user.Email)
		if err != nil {
			return nil, err
		}
	}
	if ch.ActiveStatus != 1 {
		_, err := repo.FindUserEmail(user.Email)
		if err != nil {
			return nil, err
		}
	}
	_, err = repo.db.Exec(`update users_admin
					set company_id = ?,role_id = ?,profix_name = ?,user_fname = ?,user_lname = ?,
					active_status = ?,phone = ?,email = ?,edit_time = ?,edit_by = ?
					where admin_code = ?`,
		user.CompanyID,
		user.RoleID,
		user.ProfixName,
		user.Fname,
		user.Lname,
		user.ActiveStatus,
		user.Phone,
		user.Email,
		time.Now(),
		adminID,
		user.AdminID,
	)
	logg.Println(user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (repo *adminRepository) FindUserByEmail(email string) (*admin.UserModelEmail, error) {
	sql1 := `select email,pwd from users_admin where email = ? limit 1`

	rs, err := repo.db.QueryRow(sql1, email)
	if err != nil {
		return nil, err
	}
	user := admin.UserModelEmail{}
	logg.Println(1)
	err = rs.Scan(&user.Code, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("ไม่มี email นี้อยู่ในระบบ")
	}
	logg.Println(1)
	if err != nil {
		return nil, err
	}
	logg.Println(1)
	return &user, nil
}

func (repo *adminRepository) FindUIDAdmin(email string) (string, error) {
	var admin_code string = ""
	sql1 := `select admin_code as admin_code from users_admin where  email = ? limit 1`
	rs, _ := repo.db.QueryRow(sql1, email)
	rs.Scan(&admin_code)
	if admin_code == "" {
		return "", errors.New("ไม่พบ email ของผู้ใช้")
	}
	return admin_code, nil
}
func (repo *adminRepository) SaveToekenAdmin(email string, token string) (interface{}, error) {
	// _, err := repo.dbx.NamedExec(`UPDATE users_admin SET access_token = :token WHERE email = :email`,
	// 	map[string]interface{}{
	// 		"token": token,
	// 		"email": email,
	// 	},
	// )
	// if err != nil {
	// 	log.Error(err)
	// 	return nil, err
	// }
	// return nil, nil
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
			"type":         2,
			"access_token": token,
		})
	_, err = q.Exec()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (repo *adminRepository) FindUserEmailForActiveAdmin(email string) (bool, error) {
	var user int = 0
	sql := `select count(id) as user from users_admin where email = ? and active_status = 1`
	rs, _ := repo.db.QueryRow(sql, email)
	rs.Scan(&user)
	if user > 0 {
		return false, errors.New("user ที่กำลังปับสถานะนี้ ได้มีผู้ใช้อื่นได้ไช้ email นี้แล้ว กรุณาเปลี่ยน email หรือ ลบ ผู้ใช้ที่มี emailเดียวกันนี้")
	}
	return true, nil
}

func (repo *adminRepository) FindUserEmail(email string) (bool, error) {
	var user int = 0
	sql := `select count(id) as user from users_admin where email = ? and active_status = 1`
	rs, _ := repo.db.QueryRow(sql, email)
	rs.Scan(&user)
	if user > 0 {
		return false, errors.New("email นี้ได้ทำการลงทะเบียน กับผู้ใช้อื่นแล้ว")
	}
	return true, nil
}
func (repo *adminRepository) FindUserID() (string, error) {
	like := "BN" + GetYear() + GetMonth() + GetDay()
	var userid string
	sql := `select COALESCE(admin_code,'BN6208240000') as admin_code from users_admin where admin_code like '%` + like + `%' order by id desc LIMIT 1`
	rs, _ := repo.db.QueryRow(sql)
	fmt.Println(sql)
	rs.Scan(&userid)

	return userid, nil
}

func (repo *adminRepository) RemoveToken(tokenID string) (interface{}, error) {
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

func (r *adminRepository) ListRoleRepo(status int64) (interface{}, error) {
	search := ""
	if status == 0 {
		search += `where type='font'`
	} else if status == 1 {
		search += `where type='back'`
	} else {
		search += ""
	}

	sql1 := `select ifnull(code,0) as role_id, ifnull(role_name,'') as role_name,ifnull(active_status,0) as active_status,
	case when type = 'font' then 0 else 1 end as type
			from user_role ` + search
	rs, err := r.db.Query(sql1)
	if err != nil {
		return nil, err
	}
	models := []admin.ModelRole{}
	for rs.Next() {
		model := admin.ModelRole{}
		err := rs.Scan(&model.RoleID,
			&model.RoleName,
			&model.ActiveStatus,
			&model.Type)
		if err != nil {
			logg.Error(err.Error())
			return nil, err
		}
		models = append(models, model)
	}
	return models, nil
}

func (repo *adminRepository) AddUserAdminRepo(adminID string, email string, pass string, roleID int64) (interface{}, error) {
	_, err := repo.FindUserEmail(email)
	if err != nil {
		return nil, err
	}
	iduser, _ := repo.FindUserID()
	admincode, err := GetAdminID(iduser)
	if err != nil {
		return nil, err
	}
	_, err = repo.dbx.NamedExec(
		`INSERT INTO
			users_admin
				(role_id,admin_code,pwd,create_time,create_by,email)
				VALUES
				(:role_id,:admin_code, :pwd,:create_time,:create_by, :email)
				`,
		map[string]interface{}{
			"role_id":     roleID,
			"admin_code":  admincode,
			"pwd":         pass,
			"create_time": time.Now(),
			"create_by":   adminID,
			"email":       email,
		},
	)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return nil, nil
}
func GetAdminID(user string) (string, error) {
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
		return "BN" + GetYear() + GetMonth() + GetDay() + lastuserid, nil
	} else {
		return "BN" + GetYear() + GetMonth() + GetDay() + "00001", nil
	}
	return "", nil
}
func (r *adminRepository) CountAdminRepo(status int64, search string) (int64, error) {
	var TypeSearch string = ""
	var count int64 = 0
	if status == 0 {
		TypeSearch = " a.active_status = 0 and "
	} else if status == 1 {
		TypeSearch = " a.active_status = 1 and "
	} else {
		TypeSearch = " "
	}
	sql1 := `select count(*) as count
		from  users_admin a 
		where` + TypeSearch + ` (` + search + `)`
	rs, err := r.db.QueryRow(sql1)
	if err != nil {
		return 0, err
	}

	err = rs.Scan(&count)
	if err != nil {

		return 0, err
	}
	return count, nil
}

func (r *adminRepository) ListAdminRepo(status int64, search string, limit int64) (interface{}, error) {
	var TypeSearch string = ""
	var OrderBy string = ""
	if status == 0 {
		TypeSearch = " a.active_status = 0 and "
		OrderBy = "order by a.active_status asc"
	} else if status == 1 {
		TypeSearch = " a.active_status = 1 and "
		OrderBy = "order by a.active_status asc"
	} else {
		TypeSearch = " "
		OrderBy = "order by a.active_status desc"
	}
	sql1 := `select 
	ifnull(a.admin_code,'') as admin_code,
	ifnull(a.company_id,0) as company_id,
	ifnull(a.role_id,0) as role_id,
	ifnull(b.role_name,'') as role_name,
	ifnull(a.profix_name,'') as profix_name,
	ifnull(a.user_fname,'') as user_fname,
	ifnull(a.user_lname,'') as user_lname,
	ifnull(a.phone,'') as phone,
	ifnull(a.active_status,0) as active_status,
	ifnull(a.email,'') as email
	 from users_admin a 
	left join user_role b on a.role_id = b.id
 	where` + TypeSearch + `(` + search + `) ` + OrderBy + ` limit ` + fmt.Sprintln(limit)
	logg.Println(sql1)
	rs, err := r.db.Query(sql1)
	if err != nil {

		return nil, err
	}
	admins := []admin.AdminProfileModel{}
	for rs.Next() {
		admin := admin.AdminProfileModel{}
		err = rs.Scan(&admin.AdminID, &admin.CompanyID, &admin.RoleID, &admin.RoleName, &admin.ProfixName, &admin.Fname, &admin.Lname, &admin.Phone, &admin.ActiveStatus,
			&admin.Email)
		if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins, nil
}

func (repo *adminRepository) GetAdminByIDRepo(adminID string) (interface{}, error) {
	sql1 := `select 
	ifnull(a.admin_code,'') as admin_code,
	ifnull(a.company_id,0) as company_id,
	ifnull(a.role_id,0) as role_id,
	ifnull(b.role_name,'') as role_name,
	ifnull(a.profix_name,'') as profix_name,
	ifnull(a.user_fname,'') as user_fname,
	ifnull(a.user_lname,'') as user_lname,
	ifnull(a.phone,'') as phone,
	ifnull(a.active_status,0) as active_status,
	ifnull(a.email,'') as email
	 from users_admin a 
	left join user_role b on a.role_id = b.id
	where a.admin_code = ? limit 1`

	rs, err := repo.db.QueryRow(sql1, adminID)
	if err != nil {
		return nil, err
	}
	admin := admin.UpdateUserProfileModel{}
	err = rs.Scan(&admin.AdminID, &admin.CompanyID, &admin.RoleID, &admin.RoleName, &admin.ProfixName, &admin.Fname, &admin.Lname, &admin.Phone, &admin.ActiveStatus, &admin.Email)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo *adminRepository) GetProfileAdmin(AdminCode string) (*admin.AdminProfileModel, error) {
	sql := `select 
	ifnull(a.admin_code,'') as admin_code,
	ifnull(a.company_id,0) as company_id,
	ifnull(a.role_id,0) as role_id,
	ifnull(b.role_name,'') as role_name,
	ifnull(a.profix_name,'') as profix_name,
	ifnull(a.user_fname,'') as user_fname,
	ifnull(a.user_lname,'') as user_lname,
	ifnull(a.phone,'') as phone,
	ifnull(a.active_status,0) as active_status,
	ifnull(a.email,'') as email
	 from users_admin a 
	left join user_role b on a.role_id = b.id
	 where admin_code = ? limit 1`
	rs, err := repo.db.QueryRow(sql, AdminCode)
	if err != nil {
		return nil, err
	}
	admin := admin.AdminProfileModel{}
	err = rs.Scan(&admin.AdminID, &admin.CompanyID, &admin.RoleID, &admin.RoleName, &admin.ProfixName, &admin.Fname, &admin.Lname, &admin.Phone, &admin.ActiveStatus,
		&admin.Email)
	if err != nil {
		return nil, err
	}
	menu, err := repo.GetMenuPermision(AdminCode, admin.RoleID)
	if err != nil {
		return nil, err
	}
	admin.Menu = menu
	return &admin, nil
}

func (r *adminRepository) GetStatusOrder(status int64) (int64, error) {
	var count int64 = 0
	var keyword string = ""

	if status == 1 {
		keyword = `ifnull(pic_slip,'') <> '' and order_status = ` + fmt.Sprintln(status)
	} else {
		keyword = `order_status = ` + fmt.Sprintln(status)
	}
	sql1 := `select count(*) as count from orders where ` + keyword + `limit 1`
	rs, err := r.db.QueryRow(sql1)
	if err != nil {
		logg.Println(err.Error())

		return 0, err
	}
	err = rs.Scan(&count)
	if err != nil {
		logg.Println(err.Error())
		return 0, err
	}
	return count, nil
}

func (r *adminRepository) GetMenuPermision(UserID string, role int64) ([]admin.HeadMenuModel, error) {

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
	heads := []admin.HeadMenuModel{}
	for rs.Next() {
		head := admin.HeadMenuModel{}
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
		left JOIN bn_map_menu b on a.menu_id = b.menu_id and b.status = 0
		where a.head_menu_id = ? and b.role_id = ? order by a.menu_line_number asc`
		r2, err := r.db.Query(sql1, head.HeadMenuID, role)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		}
		subs := []admin.SubMenuModel{}
		for r2.Next() {
			sub := admin.SubMenuModel{}
			r2.Scan(
				&sub.MenuID,
				&sub.HeadMenuID,
				&sub.MenuName,
				&sub.MenuLink,
				&sub.MenuShow,
				&sub.MenuStatus,
				&sub.MenuLineNumber)
			if head.HeadMenuID == 1 {
				sub.Notify, _ = r.GetStatusOrder(sub.MenuLineNumber)
				head.Notify += sub.Notify
			}
			subs = append(subs, sub)
		}
		head.Sub = subs
		heads = append(heads, head)
	}

	return heads, nil
}
