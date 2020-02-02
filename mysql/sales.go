package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"

	//"github.com/labstack/gommon/log"

	log "gitlab.com/satit13/perfect_api/logger"
	logg "gitlab.com/satit13/perfect_api/logger"
	"gitlab.com/satit13/perfect_api/sales"
	"upper.io/db.v3/lib/sqlbuilder"

	//"upper.io/db.v3/postgresql"
	"upper.io/db.v3/mysql"
)

// NewAuthRepository creates new auth repository
func NewSaleRepository(db *sql.DB) (sales.Repository, error) {
	// fmt.Println(1)
	pdb, err := mysql.New(db)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	dbx := sqlx.NewDb(db, "mysql")
	// fmt.Println(1)
	r := saleRepository{pdb, dbx}
	// fmt.Println(1)
	return &r, nil
}

type saleRepository struct {
	db  sqlbuilder.Database
	dbx *sqlx.DB
}

func (r *saleRepository) ApproveCommisionRepo(adminID string, docno string, status int64, slipapprove string) (interface{}, error) {
	if status == 1 {
		com, err := r.GetCommisionDocNo(docno)
		if err != nil {
			return nil, err
		}
		if com.Status > status {
			return nil, errors.New("ไม่สามารถยืนยันการร้องขอได้เนื่องจาก ได้ยกเลิกบิลไปแล้ว")
		}
		_, err = r.UpdateCommision(adminID, docno, status, slipapprove, 0)
		if err != nil {
			return nil, err
		}
		return com, nil
	} else if status == 2 {
		com, err := r.GetCommisionDocNo(docno)
		if err != nil {
			return nil, err
		}
		if com.Status > status {
			return nil, errors.New("ไม่สามารถยืนยันการร้องขอได้เนื่องจาก ได้ยกเลิกบิลไปแล้ว")
		}
		_, err = r.UpdateCommision(adminID, docno, status, slipapprove, 1)
		if err != nil {
			return nil, err
		}

		return com, nil
	} else {
		return nil, errors.New("ขณะนี้ยังไม่เปิดให้ ยกเลิกการค่า commison")
	}
}

func (r *saleRepository) UpdateCommision(adminID string, docno string, status int64, slipapprove string, confirm int64) (interface{}, error) {
	_, err := r.db.Exec(`update commisions
					set status = ?,confirm = ?,confirm_by = ?,
					confirm_time = ?,slip_approve = ?
					where doc_no = ?`,
		status,
		confirm,
		adminID,
		time.Now(),
		slipapprove,
		docno,
	)
	if err != nil {
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) ListSaleConfirmed(Limit int64) (interface{}, error) {
	sql1 := `SELECT a.id,
	ifnull(a.user_code,'') as user_id, ifnull(a.sales_code,'') as sales_code,
	ifnull(a.first_name,'') as first_name, ifnull(a.last_name,'') as last_name, 
	ifnull(b.role_id,0) as role_id,
	ifnull(c.role_name,'') as role_name,
	ifnull(a.card_id,'') card_id,
	ifnull(a.book_bank_id,'') as book_bank_id,
	ifnull(a.age,0) as age,
	ifnull(a.url_card_id,'') as url_card_id, 
	ifnull(a.url_book_bank,'') as url_book_bank,
	ifnull(a.url_slip_first_buy,'') as url_slip_first_buy,
	ifnull(a.file_url_3,'') as file_url_3,
	ifnull(b.telephone,'') as telephone,
	ifnull(b.invite_code,'') as invite_code,
	ifnull(a.confirm,0) as confirm,
	ifnull(a.bank_code,'') as bank_code,
	ifnull(a.bank_name,'') as bank_name,
	ifnull(DATE_FORMAT(a.create_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as create_date
	from sales_person a
	left join users b on a.user_code = b.user_id
	left join user_role c on b.role_id = c.id where a.confirm = 1 and a.active_status = 0  and b.active_status = 1  limit ` + fmt.Sprintln(Limit)
	rs1, err := r.db.Query(sql1)

	models := []sales.ModelSale{}
	for rs1.Next() {
		model := sales.ModelSale{}
		err = rs1.Scan(&model.ID,
			&model.UserID,
			&model.SalesCode,
			&model.FirstName,
			&model.LastName,
			&model.RoleID,
			&model.RoleName,
			&model.CardID,
			&model.BookBankID,
			&model.Age,
			&model.URLCardID,
			&model.URLBookBank,
			&model.URLFirstBuy,
			&model.FileURL3,
			&model.Telephone,
			&model.InviteCode,
			&model.Confirm,
			&model.BankCode,
			&model.BankName,
			&model.CreateDate)
		model.InvitePerson, _ = r.FindInviteCodeCount(model.InviteCode)
		logg.Println(model.InviteCode)
		logg.Println(model.InvitePerson)
		if err != nil {
			logg.Error(err.Error())
			return nil, err
		}
		models = append(models, model)
	}
	len, err := r.CoutSaleType("confirmed")
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	sale := map[string]interface{}{
		"type_list": "CONFIRMED",
		"length":    len,
		"sale_list": models,
	}
	// fmt.Println(models)
	// type Rsp struct {
	// 	Length   int               `json:"length"`
	// 	SaleList []sales.ModelSale `json:"sale_list"`
	// }
	// ln := len(models)

	// x := Rsp{ln, models}

	return sale, nil
}

func (r *saleRepository) ListSaleNoConfirm(Limit int64) (interface{}, error) {
	sql1 := `SELECT a.id,
	ifnull(a.user_code,'') as user_id, ifnull(a.sales_code,'') as sales_code,
	ifnull(a.first_name,'') as first_name, ifnull(a.last_name,'') as last_name, 
	ifnull(b.role_id,0) as role_id,
	ifnull(c.role_name,'') as role_name,
	ifnull(a.card_id,'') card_id,
	ifnull(a.book_bank_id,'') as book_bank_id,
	ifnull(a.age,0) as age,
	ifnull(a.url_card_id,'') as url_card_id, 
	ifnull(a.url_book_bank,'') as url_book_bank,
	ifnull(a.url_slip_first_buy,'') as url_slip_first_buy,
	ifnull(a.file_url_3,'') as file_url_3,
	ifnull(b.telephone,'') as telephone,
	ifnull(b.invite_code,'') as invite_code,
	ifnull(a.confirm,0) as confirm,
	ifnull(a.bank_code,'') as bank_code,
	ifnull(a.bank_name,'') as bank_name,
	ifnull(DATE_FORMAT(a.create_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as create_date
	from sales_person a
	left join users b on a.user_code = b.user_id 
	left join user_role c on b.role_id = c.id where a.confirm = 0 and a.active_status = 0 and b.active_status = 1 limit ` + fmt.Sprintln(Limit)
	rs1, err := r.db.Query(sql1)
	log.Println(sql1)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	models := []sales.ModelSale{}

	for rs1.Next() {
		model := sales.ModelSale{}
		err = rs1.Scan(&model.ID,
			&model.UserID,
			&model.SalesCode,
			&model.FirstName,
			&model.LastName,
			&model.RoleID,
			&model.RoleName,
			&model.CardID,
			&model.BookBankID,
			&model.Age,
			&model.URLCardID,
			&model.URLBookBank,
			&model.URLFirstBuy,
			&model.FileURL3,
			&model.Telephone,
			&model.InviteCode,
			&model.Confirm,
			&model.BankCode,
			&model.BankName,
			&model.CreateDate)
		model.InvitePerson, _ = r.FindInviteCodeCount(model.InviteCode)
		if err != nil {
			logg.Error(err.Error())
			return nil, err
		}
		models = append(models, model)
	}
	len, err := r.CoutSaleType("noconfirmed")
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	sale := map[string]interface{}{
		"type_list": "NOCONFIRMED",
		"length":    len,
		"sale_list": models,
	}
	// fmt.Println(models)
	// type Rsp struct {
	// 	Length   int               `json:"length"`
	// 	SaleList []sales.ModelSale `json:"sale_list"`
	// }
	// ln := len(models)

	// x := Rsp{ln, models}

	return sale, nil
}

func (repo *saleRepository) FindCommisionDocNo() (string, error) {
	like := "WDC" + GetYear() + GetMonth() + GetDay()
	var userid string
	sql := `select COALESCE(doc_no,'WDC62100300001') as userid from commisions where doc_no like '%` + like + `%' order by id desc LIMIT 1`
	rs, _ := repo.db.QueryRow(sql)
	fmt.Println(sql)
	rs.Scan(&userid)

	return userid, nil
}

func GetCommisionCode(user string) (string, error) {
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
		return "WDC" + GetYear() + GetMonth() + GetDay() + lastuserid, nil
	} else {
		return "WDC" + GetYear() + GetMonth() + GetDay() + "00001", nil
	}
	return "", nil
}

func (r *saleRepository) ListCommsionRepo(userID string) (interface{}, error) {
	comm, err := r.ListCommision(userID, 1)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	return comm, nil
}

func (r *saleRepository) CountListCommmisonBackEnd(status int64, search string) (int64, error) {
	var TypeSearch string = ""
	var count int64 = 0
	if status == 0 {
		TypeSearch = " ( a.status = 0 or a.status = 9 ) and a.confirm = 0 and "
	} else if status == 1 {
		TypeSearch = " a.status = 1 and a.confirm = 0 and"
	} else if status == 2 {
		TypeSearch = " a.status = 2 and a.confirm = 1 and"
	} else if status == 3 {
		TypeSearch = " a.status != 0 and"
	} else {
		TypeSearch = ""
	}

	sql1 := `select count(*) as count
		from  commisions a 
		left join users b on a.user_id = b.user_id
		where` + TypeSearch + ` (` + search + `)`
	rs, err := r.db.QueryRow(sql1)
	if err != nil {
		return 0, err
	}
	log.Println(sql1)
	err = rs.Scan(&count)
	if err != nil {

		return 0, err
	}
	return count, nil
}

func (r *saleRepository) ListCommsionBackendRepo(status int64, search string, limit int64) (interface{}, error) {
	// var TypeSearch string = ""
	// if Type == "NOCONFIRMED" {
	// 	TypeSearch = " and a.confirm = 0"
	// } else if Type == "CONFIRMED" {
	// 	TypeSearch = " and a.confirm = 1 "
	// } else {
	// 	TypeSearch = ""
	// }
	var TypeSearch string = ""
	var OrderBy string = ""
	if status == 0 {
		TypeSearch = " ( a.status = 0 or a.status = 9 ) and a.confirm = 0 and "
		OrderBy = "order by a.status asc"
	} else if status == 1 {
		TypeSearch = " a.status = 1 and a.confirm = 0 and "
	} else if status == 2 {
		TypeSearch = " a.status = 2 and a.confirm = 1 and "
	} else if status == 3 {
		TypeSearch = " a.status != 0 and "
		OrderBy = "order by a.confirm asc"
	} else {
		OrderBy = ""
		TypeSearch = ""
	}

	sql1 := `select a.id,
	ifnull(a.user_id,'') as user_id,
	ifnull(a.sales_code,'') as sales_code,
	ifnull(b.user_fname,'') as f_name,
	ifnull(b.user_lname,'') as l_name,
	ifnull(b.telephone,'') as telephone,
	ifnull(a.doc_no,'') as doc_no, 
	ifnull(DATE_FORMAT(a.doc_date, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as doc_date,
	ifnull(a.net_amount,0) as total_amount,
	ifnull(a.all_sale_commision,0) as all_sale_commision,
	ifnull(a.status,0) as status, 
	ifnull(a.my_description,'') as my_description,
	ifnull(a.slip_approve,'') as slip_approve,
	ifnull(a.confirm,0) as confirm,
	ifnull(DATE_FORMAT(a.confirm_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as confirm_time
	from commisions a
	left join users b on a.user_id = b.user_id
 	where` + TypeSearch + `(` + search + `) ` + OrderBy + ` limit ` + fmt.Sprintln(limit)
	log.Println(sql1)
	rs, err := r.db.Query(sql1)
	if err != nil {

		return nil, err
	}
	models := []sales.ModelCommision{}
	for rs.Next() {
		model := sales.ModelCommision{}
		err = rs.Scan(
			&model.ID,
			&model.UserID,
			&model.SalesCode,
			&model.FName,
			&model.LName,
			&model.Telephone,
			&model.DocNo,
			&model.DocDate,
			&model.TotalAmount,
			&model.ALLSaleCommision,
			&model.Status,
			&model.MyDescription,
			&model.SlipApprove,
			&model.Confirm,
			&model.ConfirmTime)
		if err != nil {
			return nil, err
		}
		binfo, err := r.GetBankInfoBySaleCode(model.SalesCode)
		if err != nil {
			return nil, err
		}
		model.BankInfo = sales.BankInfoModel{
			CardID:      binfo.CardID,
			URLCardID:   binfo.URLCardID,
			BookBankID:  binfo.BookBankID,
			URLBookBank: binfo.URLBookBank,
			BankCode:    binfo.BankCode,
			BankName:    binfo.BankName,
		}
		// sql2 := `select id,
		// 		ifnull(ref_doc,'') as order_doc,
		// 		ifnull(DATE_FORMAT(ref_doc_date, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as order_date,
		// 		ifnull(user_id,'') as user_id,
		// 		ifnull(f_name,'') as f_name,
		// 		ifnull(l_name,'') as l_name,
		// 		ifnull(net_amount,0) as total_amount
		// 		from commision_sub where doc_no = ?`
		// rs2, err := r.db.Query(sql2, model.DocNo)
		// if err != nil {
		// 	logg.Error(err.Error())
		// 	return nil, err
		// }
		// logg.Println(model)
		// subs := []sales.ModelCommisionSub{}
		// for rs2.Next() {
		// 	sub := sales.ModelCommisionSub{}
		// 	err = rs2.Scan(
		// 		&sub.ID,
		// 		&sub.OrderDoc,
		// 		&sub.OrderDate,
		// 		&sub.UserID,
		// 		&sub.FName,
		// 		&sub.LName,
		// 		&sub.TotalAmount)
		// 	if err != nil {
		// 		logg.Error(err.Error())
		// 		return nil, err
		// 	}
		// 	subs = append(subs, sub)
		// }
		// model.Sub = subs
		models = append(models, model)
	}
	logg.Println(models)
	return models, nil
}

func (r *saleRepository) ListCommision(userID string, status int64) ([]sales.ModelCommision, error) {
	var keyword string = ""
	if status == 0 {
		keyword = " and a.confirm = 0"
	}
	sql1 := `select a.id,
	ifnull(a.user_id,'') as user_id,
	ifnull(a.sales_code,'') as sales_code,
	ifnull(b.user_fname,'') as f_name,
	ifnull(b.user_lname,'') as l_name,
	ifnull(b.telephone,'') as telephone,
	ifnull(a.doc_no,'') as doc_no, 
	ifnull(DATE_FORMAT(a.doc_date, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as doc_date,
	ifnull(a.net_amount,0) as total_amount, 
	ifnull(a.all_sale_commision,0) as all_sale_commision,
	ifnull(a.status,0) as status, 
	ifnull(a.my_description,'') as my_description,
	ifnull(a.slip_approve,'') as slip_approve,
	ifnull(a.confirm,0) as confirm,
	ifnull(DATE_FORMAT(a.confirm_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as confirm_time
	from commisions a
	left join users b on a.user_id = b.user_id
	where a.user_id = ? ` + fmt.Sprintln(keyword)
	fmt.Println(sql1)
	rs, err := r.db.Query(sql1, userID)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	models := []sales.ModelCommision{}
	for rs.Next() {
		model := sales.ModelCommision{}
		err = rs.Scan(
			&model.ID,
			&model.UserID,
			&model.SalesCode,
			&model.FName,
			&model.LName,
			&model.Telephone,
			&model.DocNo,
			&model.DocDate,
			&model.TotalAmount,
			&model.ALLSaleCommision,
			&model.Status,
			&model.MyDescription,
			&model.SlipApprove,
			&model.Confirm,
			&model.ConfirmTime)
		if err != nil {
			return nil, err
		}
		binfo, err := r.GetBankInfoBySaleCode(model.SalesCode)
		if err != nil {
			return nil, err
		}
		model.BankInfo = sales.BankInfoModel{
			CardID:      binfo.CardID,
			URLCardID:   binfo.URLCardID,
			BookBankID:  binfo.BookBankID,
			URLBookBank: binfo.URLBookBank,
			BankCode:    binfo.BankCode,
			BankName:    binfo.BankName,
		}
		// sql2 := `select id,
		// 		ifnull(ref_doc,'') as order_doc,
		// 		ifnull(DATE_FORMAT(ref_doc_date, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as order_date,
		// 		ifnull(user_id,'') as user_id,
		// 		ifnull(f_name,'') as f_name,
		// 		ifnull(l_name,'') as l_name,
		// 		ifnull(net_amount,0) as total_amount
		// 		from commision_sub where doc_no = ?`
		// rs2, err := r.db.Query(sql2, model.DocNo)
		// if err != nil {
		// 	logg.Error(err.Error())
		// 	return nil, err
		// }
		// logg.Println(model)
		// subs := []sales.ModelCommisionSub{}
		// for rs2.Next() {
		// 	sub := sales.ModelCommisionSub{}
		// 	err = rs2.Scan(
		// 		&sub.ID,
		// 		&sub.OrderDoc,
		// 		&sub.OrderDate,
		// 		&sub.UserID,
		// 		&sub.FName,
		// 		&sub.LName,
		// 		&sub.TotalAmount)
		// 	if err != nil {
		// 		logg.Error(err.Error())
		// 		return nil, err
		// 	}
		// 	subs = append(subs, sub)
		// }
		// model.Sub = subs
		models = append(models, model)
	}
	logg.Println(models)
	return models, nil
}
func (r *saleRepository) FindOrderCommision(userID string) (interface{}, int64, error) {

	return nil, 0, nil
}

func (r *saleRepository) GetCommisionDocNo(docno string) (*sales.ModelCommision, error) {

	sql1 := `select a.id,
	ifnull(a.user_id,'') as user_id,
	ifnull(a.sales_code,'') as sales_code,
	ifnull(b.user_fname,'') as f_name,
	ifnull(b.user_lname,'') as l_name,
	ifnull(b.telephone,'') as telephone,
	ifnull(a.doc_no,'') as doc_no, 
	ifnull(DATE_FORMAT(a.doc_date, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as doc_date,
	ifnull(a.net_amount,0) as total_amount, 
	ifnull(a.all_sale_commision,0) as all_sale_commision,
	ifnull(a.status,0) as status, 
	ifnull(a.my_description,'') as my_description,
	ifnull(a.slip_approve,'') as slip_approve,
	ifnull(a.confirm,0) as confirm,
	ifnull(DATE_FORMAT(a.confirm_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as confirm_time
	from commisions a 
	left join users b on a.user_id = b.user_id
	where a.doc_no = ? limit 1`
	fmt.Println(sql1)
	rs, err := r.db.QueryRow(sql1, docno)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	model := sales.ModelCommision{}
	err = rs.Scan(
		&model.ID,
		&model.UserID,
		&model.SalesCode,
		&model.FName,
		&model.LName,
		&model.Telephone,
		&model.DocNo,
		&model.DocDate,
		&model.TotalAmount,
		&model.ALLSaleCommision,
		&model.Status,
		&model.MyDescription,
		&model.SlipApprove,
		&model.Confirm,
		&model.ConfirmTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ไม่พบ รหัสเอกสารที่ระบุ")
		}
		return nil, err
	}
	binfo, err := r.GetBankInfoBySaleCode(model.SalesCode)
	if err != nil {
		return nil, err
	}
	model.BankInfo = sales.BankInfoModel{
		CardID:      binfo.CardID,
		URLCardID:   binfo.URLCardID,
		BookBankID:  binfo.BookBankID,
		URLBookBank: binfo.URLBookBank,
		BankCode:    binfo.BankCode,
		BankName:    binfo.BankName,
	}
	sql2 := `select id,
				ifnull(ref_doc,'') as order_doc,
				ifnull(DATE_FORMAT(ref_doc_date, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as order_date,
				ifnull(user_id,'') as user_id,
				ifnull(f_name,'') as f_name,
				ifnull(l_name,'') as l_name,
				ifnull(net_amount,0) as total_amount
				from commision_sub where doc_no = ?`
	rs2, err := r.db.Query(sql2, model.DocNo)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	logg.Println(model)
	subs := []sales.ModelCommisionSub{}
	for rs2.Next() {
		sub := sales.ModelCommisionSub{}
		err = rs2.Scan(
			&sub.ID,
			&sub.OrderDoc,
			&sub.OrderDate,
			&sub.UserID,
			&sub.FName,
			&sub.LName,
			&sub.TotalAmount)
		if err != nil {
			logg.Error(err.Error())
			return nil, err
		}
		subs = append(subs, sub)
	}
	model.Sub = subs
	return &model, nil
}

func (r *saleRepository) GetBankInfoBySaleCode(saleCode string) (*sales.BankInfoModel, error) {
	sq1 := `select ifnull(a.card_id,'') as card_id,
		ifnull(a.url_card_id,'') as url_card_id,
		ifnull(a.book_bank_id,'') as book_bank_id,
		ifnull(a.url_book_bank,'') as url_book_bank,
		ifnull(a.bank_code,'') as bank_code,
		ifnull(a.bank_name,'') as bank_name
		from sales_person a
		where a.sales_code = ? limit 1 `
	rs, err := r.db.QueryRow(sq1, saleCode)
	if err != nil {
		return nil, err
	}
	model := sales.BankInfoModel{}
	err = rs.Scan(
		&model.CardID,
		&model.URLCardID,
		&model.BookBankID,
		&model.URLBookBank,
		&model.BankCode,
		&model.BankName)
	if err != nil {
		return nil, err
		log.Error(err.Error())
	}
	return &model, nil
}

func (r *saleRepository) GetCommisionDocNoRepo(userID string, docno string) (*sales.ModelCommision, error) {

	sql1 := `select a.id,
	ifnull(a.user_id,'') as user_id,
	ifnull(a.sales_code,'') as sales_code,
	ifnull(b.user_fname,'') as f_name,
	ifnull(b.user_lname,'') as l_name,
	ifnull(b.telephone,'') as telephone,
	ifnull(a.doc_no,'') as doc_no, 
	ifnull(DATE_FORMAT(a.doc_date, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as doc_date,
	ifnull(a.net_amount,0) as total_amount, 
	ifnull(a.all_sale_commision,0) as all_sale_commision,
	ifnull(a.status,0) as status, 
	ifnull(a.my_description,'') as my_description,
	ifnull(a.slip_approve,'') as slip_approve,
	ifnull(a.confirm,0) as confirm,
	ifnull(DATE_FORMAT(a.confirm_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as confirm_time
	from commisions a 
	left join users b on a.user_id = b.user_id
	where a.user_id = ? and a.doc_no = ? limit 1`
	fmt.Println(sql1)
	rs, err := r.db.QueryRow(sql1, userID, docno)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	model := sales.ModelCommision{}

	err = rs.Scan(
		&model.ID,
		&model.UserID,
		&model.SalesCode,
		&model.FName,
		&model.LName,
		&model.Telephone,
		&model.DocNo,
		&model.DocDate,
		&model.TotalAmount,
		&model.ALLSaleCommision,
		&model.Status,
		&model.MyDescription,
		&model.SlipApprove,
		&model.Confirm,
		&model.ConfirmTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ไม่พบ รหัสเอกสารที่ระบุ")
		}
		return nil, err
	}
	binfo, err := r.GetBankInfoBySaleCode(model.SalesCode)
	if err != nil {
		return nil, err
	}
	model.BankInfo = sales.BankInfoModel{
		CardID:      binfo.CardID,
		URLCardID:   binfo.URLCardID,
		BookBankID:  binfo.BookBankID,
		URLBookBank: binfo.URLBookBank,
		BankCode:    binfo.BankCode,
		BankName:    binfo.BankName,
	}
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	sql2 := `select id,
				ifnull(ref_doc,'') as order_doc,
				ifnull(DATE_FORMAT(ref_doc_date, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as order_date,
				ifnull(user_id,'') as user_id,
				ifnull(f_name,'') as f_name,
				ifnull(l_name,'') as l_name,
				ifnull(net_amount,0) as total_amount
				from commision_sub where doc_no = ?`
	rs2, err := r.db.Query(sql2, model.DocNo)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	logg.Println(model)
	subs := []sales.ModelCommisionSub{}
	for rs2.Next() {
		sub := sales.ModelCommisionSub{}
		err = rs2.Scan(
			&sub.ID,
			&sub.OrderDoc,
			&sub.OrderDate,
			&sub.UserID,
			&sub.FName,
			&sub.LName,
			&sub.TotalAmount)
		if err != nil {
			logg.Error(err.Error())
			return nil, err
		}
		subs = append(subs, sub)
	}
	model.Sub = subs
	return &model, nil
}
func (r *saleRepository) CountcommisionOrder(userID string) (int64, error) {
	var count int64 = 0
	var total_commision = 0
	sql1 := `select sum(sum_cash_amount+sum_credit_amount+sum_bank_amount) as total_commision,count(*) as count from orders where user_id = ? and commision = 0 and order_status <> 1 and order_status <> 99 limit 1`
	rs, err := r.db.QueryRow(sql1, userID)
	if err != nil {
		logg.Error(err.Error())
		return 0, err
	}
	rs.Scan(&count, &total_commision)
	if count <= 0 {
		return 0, errors.New("ไม่พบ รายการสั่งชื้อของลูกทีม")
	}
	return count, nil
}
func (r *saleRepository) FindallcommsisionByUserID(UserID string) (float64, error) {
	var all_sale float64 = 0

	sql1 := `select sum(sum_cash_amount+sum_credit_amount+sum_bank_amount) as all_sale from orders where user_id = ? and commision = 0 and order_status <> 1 and order_status <> 99 limit 1`
	rs1, _ := r.db.QueryRow(sql1, UserID)
	rs1.Scan(&all_sale)
	return all_sale, nil
}

func (r *saleRepository) FindAllSaleByUserID(UserID string, status int64) (float64, float64, error) {
	var all_sale float64 = 0
	var commision float64 = 0
	sql1 := `select sum(sum_cash_amount+sum_credit_amount+sum_bank_amount) as all_sale from orders where user_id = ? and order_status <> 1 and order_status <> 99 limit 1`
	rs1, _ := r.db.QueryRow(sql1, UserID)
	rs1.Scan(&all_sale)

	if status == 2 {
		sql2 := `select sum(sum_cash_amount+sum_credit_amount+sum_bank_amount) as commision from orders where user_id = ? and commision = 0 and order_status <> 1 and order_status <> 99 limit 1`
		rs2, _ := r.db.QueryRow(sql2, UserID)
		rs2.Scan(&commision)
		return all_sale, commision, nil
	}
	return all_sale, 0, nil
}

func (r *saleRepository) FindOrderCommisions(userID string, model *[]sales.ModeOrder) (interface{}, error) {
	sql1 :=
		`select 
		ifnull(a.id,'') as id,
		ifnull(a.user_id,'') as user_id,
		ifnull(a.doc_no,'') as doc_no,
		ifnull(DATE_FORMAT(a.create_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as ref_doc_date,
		ifnull(b.user_fname,'') as f_name,
		ifnull(b.user_lname,'') as l_name,
		(a.sum_cash_amount+a.sum_credit_amount+a.sum_bank_amount) as total_amount 
		from orders a
		inner JOIN users b on a.user_id = b.user_id 
		where a.user_id = ? and a.commision = 0 and a.order_status <> 1 and a.order_status <> 99
	`
	rs, err := r.db.Query(sql1, userID)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	for rs.Next() {
		models := sales.ModeOrder{}
		rs.Scan(&models.ID, &models.UserID, &models.DocNo, &models.RefDocDate,
			&models.FName, &models.LName, &models.TotalAmount)
		*model = append(*model, models)
	}
	return nil, nil
}

func (r *saleRepository) CalAllSaleTeamForGen(ref_code string) (float64, []sales.ModeOrder, error) {
	var all_sale float64 = 0
	sql1 := `select user_id,role_id from users where ref_code = ?`
	rs, _ := r.db.Query(sql1, ref_code)
	//models := []sales.ModelLinkTeamsale{}
	Orders := []sales.ModeOrder{}
	for rs.Next() {
		model := sales.ModelLinkTeamsale{}
		rs.Scan(&model.UserID, &model.RoleID)
		model.AllTeamSale, _ = r.FindallcommsisionByUserID(model.UserID)
		all_sale += model.AllTeamSale
		r.FindOrderCommisions(model.UserID, &Orders)
	}

	return all_sale, Orders, nil
}
func (r *saleRepository) CalAllSaleTeam(ref_code string) (float64, float64, error) {
	var all_sale, commision float64 = 0, 0
	sql1 := `select user_id,role_id from users where ref_code = ?`
	rs, _ := r.db.Query(sql1, ref_code)
	//models := []sales.ModelLinkTeamsale{}
	for rs.Next() {
		model := sales.ModelLinkTeamsale{}
		rs.Scan(&model.UserID, &model.RoleID)
		model.AllTeamSale, commision, _ = r.FindAllSaleByUserID(model.UserID, 2)
		all_sale += model.AllTeamSale
		commision += commision
	}
	return all_sale, commision, nil
}
func (r *saleRepository) CalAllSaleTeamCommision(ref_code string) (float64, error) {
	var all_sale float64 = 0
	sql1 := `select user_id,role_id from users where ref_code = ?`
	rs, _ := r.db.Query(sql1, ref_code)
	//models := []sales.ModelLinkTeamsale{}
	for rs.Next() {
		model := sales.ModelLinkTeamsale{}
		rs.Scan(&model.UserID, &model.RoleID)
		model.AllTeamSale, _ = r.FindallcommsisionByUserID(model.UserID)
		all_sale += model.AllTeamSale
	}
	return all_sale, nil
}

func (r *saleRepository) GenarateCommisionRepo(userID string) (interface{}, error) {
	pro, err := r.GetProfileSalesUser(userID)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	inviteSale, order, err := r.CalAllSaleTeamForGen(pro.InviteCode)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	commision, config, err := r.GetCommisonSale(inviteSale)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	if commision < 3000 {
		return nil, errors.New("ค่าคอมถอนขึ้นต่ำ 3000 บาท")
	}
	com, _ := r.FindCommisionDocNo()
	comdocno, err := GetCommisionCode(com)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	_, err = r.updateCommisionTeam(pro.InviteCode, userID)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	_, err = r.createBillCommision(userID, pro.SalesCode, inviteSale, order, comdocno, commision, config)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}

	list, err := r.ListCommision(userID, 1)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	logg.Println(list)
	return list, nil
}

func (r *saleRepository) updateCommisionTeam(ref_code string, userID string) (interface{}, error) {
	sql1 := `select user_id,role_id from users where ref_code = ?`
	rs, _ := r.db.Query(sql1, ref_code)
	//models := []sales.ModelLinkTeamsale{}
	for rs.Next() {
		model := sales.ModelLinkTeamsale{}
		rs.Scan(&model.UserID, &model.RoleID)
		_, err := r.UpdateCommisionOrder(model.UserID, userID, 1)
		if err != nil {
			logg.Error(err.Error())
			return nil, err
		}
	}
	return nil, nil
}

func (r *saleRepository) UpdateCommisionOrder(userID string, editby string, status int64) (interface{}, error) {
	_, err := r.db.Exec(`update orders
	set commision = ?,
	edit_by = ?,edit_time = ?
	where user_id = ? and commision = 0 and order_status <> 1 and order_status <> 99`,
		status,
		editby,
		time.Now(),
		userID,
	)
	if err != nil {

		logg.Error(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) CreateBillCommisionSub(userID string, order []sales.ModeOrder, docno string) (interface{}, error) {

	for _, res := range order {
		q := r.db.
			InsertInto("commision_sub").
			Values(map[string]interface{}{
				"user_id":      res.UserID,
				"doc_no":       docno,
				"f_name":       res.FName,
				"l_name":       res.LName,
				"ref_doc":      res.DocNo,
				"ref_doc_date": res.RefDocDate,
				"total_amount": res.TotalAmount,
				"net_amount":   res.TotalAmount,
				"status":       0,
				"create_by":    userID,
				"create_time":  time.Now(),
				"confirm":      0,
			})
		_, err := q.Exec()
		if err != nil {
			logg.Error(err.Error())
			return nil, err
		}

	}
	return nil, nil
}

func (r *saleRepository) createBillCommision(userID string, saleCode string, invite_sale float64, order []sales.ModeOrder, docno string, commision float64, config *sales.CommisionModel) (interface{}, error) {
	var taxRate float64 = 0
	totalAmount := commision
	if config.TaxStatus == 1 {
		commision = commision * float64(config.TaxRate) / 100.00
		taxRate = config.TaxRate
	}
	q := r.db.
		InsertInto("commisions").
		Values(map[string]interface{}{
			"user_id":            userID,
			"sales_code":         saleCode,
			"doc_no":             docno,
			"doc_date":           time.Now(),
			"all_sale_commision": invite_sale,
			"total_amount":       totalAmount,
			"tax_rate":           taxRate,
			"net_amount":         commision,
			"status":             0,
			"create_by":          userID,
			"create_time":        time.Now(),
			"confirm":            0,
			"before_tax_amount":  commision,
		})
	_, err := q.Exec()
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	_, err = r.CreateBillCommisionSub(userID, order, docno)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	return nil, nil

}

func (r *saleRepository) CoutSaleType(Type string) (int64, error) {
	logg.Println(1)
	var word string = ""
	var count int64 = 0
	if Type != "all" {
		if Type == "confirmed" {
			word = "and a.confirm = 1"
		} else if Type == "noconfirmed" {
			word = "and a.confirm = 0"
		} else {
			word = ""
		}
	}
	sql1 := `select count(*) as count from sales_person a where a.active_status = 0 ` + word
	logg.Println(sql1)
	rs, err := r.db.QueryRow(sql1)
	if err != nil {
		logg.Error(err.Error())
		return 0, err
	}
	err = rs.Scan(&count)
	if err != nil {
		logg.Error(err.Error())
		return 0, err
	}

	return count, nil
}

func (r *saleRepository) ListSaleALL(Limit int64) (interface{}, error) {

	sql1 := `SELECT a.id,
	ifnull(a.user_code,'') as user_id, ifnull(a.sales_code,'') as sales_code,
	ifnull(a.first_name,'') as first_name, ifnull(a.last_name,'') as last_name, 
	ifnull(b.role_id,0) as role_id,
	ifnull(c.role_name,'') as role_name,
	ifnull(a.card_id,'') card_id,
	ifnull(a.book_bank_id,'') as book_bank_id,
	ifnull(a.age,0) as age,
	ifnull(a.url_card_id,'') as url_card_id, 
	ifnull(a.url_book_bank,'') as url_book_bank,
	ifnull(a.url_slip_first_buy,'') as url_slip_first_buy,
	ifnull(a.file_url_3,'') as file_url_3,
	ifnull(b.telephone,'') as telephone,
	ifnull(b.invite_code,'') as invite_code,
	ifnull(a.confirm,0) as confirm,
	ifnull(a.bank_code,'') as bank_code,
	ifnull(a.bank_name,'') as bank_name,
	ifnull(DATE_FORMAT(a.create_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as create_date
	from sales_person a
	left join users b on a.user_code = b.user_id 
	left join user_role c on b.role_id = c.id where a.active_status = 0 and b.active_status = 1 limit ` + fmt.Sprintln(Limit)
	rs1, err := r.db.Query(sql1)
	if err != nil {
		logg.Error(err.Error())
		logg.Println(err.Error())
		return nil, err
	}
	models := []sales.ModelSale{}
	for rs1.Next() {
		model := sales.ModelSale{}
		err = rs1.Scan(&model.ID,
			&model.UserID,
			&model.SalesCode,
			&model.FirstName,
			&model.LastName,
			&model.RoleID,
			&model.RoleName,
			&model.CardID,
			&model.BookBankID,
			&model.Age,
			&model.URLCardID,
			&model.URLBookBank,
			&model.URLFirstBuy,
			&model.FileURL3,
			&model.Telephone,
			&model.InviteCode,
			&model.Confirm,
			&model.BankCode,
			&model.BankName,
			&model.CreateDate)
		model.InvitePerson, _ = r.FindInviteCodeCount(model.InviteCode)

		if err != nil {
			logg.Error(err.Error())
			return nil, err
		}
		models = append(models, model)
	}
	logg.Println("all")
	len, err := r.CoutSaleType("all")
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	sale := map[string]interface{}{
		"type_list": "ALL",
		"length":    len,
		"sale_list": models,
	}
	// fmt.Println(models)
	// type Rsp struct {
	// 	Length   int               `json:"length"`
	// 	SaleList []sales.ModelSale `json:"sale_list"`
	// }
	// ln := len(models)

	// x := Rsp{ln, models}

	return sale, nil
}
func (r *saleRepository) CancelCommisionFontRepo(userID string, docno string) (interface{}, error) {
	comm, err := r.GetCommisionDocNoRepo(userID, docno)
	if err != nil {
		return nil, err
	}
	if comm.Status != 0 {
		return nil, errors.New("ไม่สามารถยกเลิก commision ที่ ยืนยันการร้องขอไปแล้วกรุณาติดต่อเจ้าหน้าที่")
	}
	_, err = r.ChangeOrderCommisionDocNo(0, userID, comm)
	if err != nil {
		return nil, err
	}
	_, err = r.CancelCommision(9, userID, docno)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) CancelCommisionRepo(userID string, docno string) (interface{}, error) {
	comm, err := r.GetCommisionDocNo(docno)
	if err != nil {
		return nil, err
	}
	if comm.Status != 0 {
		return nil, errors.New("ไม่สามารถยกเลิก commision ที่ ยืนยันการร้องขอไปแล้วกรุณาติดต่อเจ้าหน้าที่")
	}
	_, err = r.ChangeOrderCommisionDocNo(0, userID, comm)
	if err != nil {
		return nil, err
	}
	_, err = r.CancelCommision(9, userID, docno)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) CancelCommision(status int64, userID string, docno string) (interface{}, error) {
	_, err := r.db.Exec(`update commisions
	set status = ?,
	edit_by = ?,edit_time = ?
	where doc_no = ?`, status, userID, time.Now(), docno)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) ChangeOrderCommisionDocNo(status int64, userID string, model *sales.ModelCommision) (interface{}, error) {
	for _, i := range model.Sub {
		_, err := r.db.Exec(`update orders
	set commision = ?,
	edit_by = ?,edit_time = ?
	where doc_no = ?`, status, userID, time.Now(), i.OrderDoc)
		if err != nil {
			logg.Error(err.Error())
			return nil, err
		}
	}
	return nil, nil
}
func (r *saleRepository) RemoveSales(status int64, adminID string, saleCode string) (interface{}, error) {

	_, err := r.db.Exec(`update sales_person
		set active_status = ?,edit_time = ?,edit_by = ?
		where sales_code = ?`,
		status,
		time.Now(),
		adminID,
		saleCode,
	)
	if err != nil {
		logg.Error(err.Error())

		return nil, err
	}
	return nil, nil

}

func (r *saleRepository) RemoveSalesRepo(AdminID string, SaleCode string, USerID string) (interface{}, error) {
	_, err := r.RemoveSales(9, AdminID, SaleCode)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	_, err = r.UpdateRoleUser(1, USerID, AdminID)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) FindNameSale(saleCode string, userID string) (string, string, error) {
	var f_name string = ""
	var l_name string = ""
	sql1 := `select ifnull(first_name,'') as f_name, ifnull(last_name,'') as l_name from sales_person
		where user_code = ? and sales_code = ? limit 1`
	rs, err := r.db.QueryRow(sql1, userID, saleCode)
	if err != nil {
		logg.Error(err.Error())
		return "", "", err
	}
	err = rs.Scan(&f_name, &l_name)
	if err != nil {
		logg.Error(err.Error())
		return "", "", err
	}
	return f_name, l_name, nil
}

func (r *saleRepository) UpdateuserSale(f_name string, l_name string, adminID string, userID string) (interface{}, error) {

	_, err := r.db.Exec(`update users
		set user_fname = ?,user_lname = ?,edit_by = ?,edit_time = ?
		where user_id = ?`,
		f_name,
		l_name,
		adminID,
		time.Now(),
		userID,
	)
	if err != nil {
		logg.Error(err.Error())
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) UpdateUserAffterRegister(AdminID string, saleCode string, userID string) (interface{}, error) {
	f_name, l_name, err := r.FindNameSale(saleCode, userID)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	_, err = r.UpdateuserSale(f_name, l_name, AdminID, userID)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) SaleConfirmRepo(AdminID string, SaleCode string, UserID string) (interface{}, error) {
	_, err := r.UpdateUserAffterRegister(AdminID, SaleCode, UserID)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	_, err = r.UpdateSaleConfirm(AdminID, SaleCode)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	_, err = r.UpdateRoleUser(2, UserID, AdminID)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	sale, err := r.GetProfileSales(SaleCode)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	return sale, nil
}
func (r *saleRepository) GetProfileSalesbyUserV2(UserID string) (interface{}, error) {
	sql1 := `SELECT a.id,
	ifnull(a.user_code,'') as user_id, 
	ifnull(a.sales_code,'') as sales_code,
	ifnull(a.first_name,'') as first_name, 
	ifnull(a.last_name,'') as last_name, 
	ifnull(b.role_id,0) as role_id,
	ifnull(c.role_name,'') as role_name,
	ifnull(a.card_id,'') card_id,
	ifnull(a.book_bank_id,'') as book_bank_id,
	ifnull(a.age,0) as age,
	ifnull(a.url_card_id,'') as url_card_id, 
	ifnull(a.url_book_bank,'') as url_book_bank,
	ifnull(a.url_slip_first_buy,'') as url_slip_first_buy,
	ifnull(a.file_url_3,'') as file_url_3,
	ifnull(b.invite_code,'') as invite_code,
	ifnull(a.confirm,0) as confirm,
	ifnull(DATE_FORMAT(a.create_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as create_date
	from sales_person a
	left join users b on a.user_code = b.user_id 
	left join user_role c on b.role_id = c.id
	where b.user_id = ? and a.active_status = 0 and b.active_status = 1 limit 1
	`
	rs1, err := r.db.QueryRow(sql1, UserID)
	if err != nil {
		logg.Error(err.Error())
		logg.Println(err.Error())
		if err == sql.ErrNoRows {
			return nil, errors.New("ไม่มีข้อมูลการสมัค sales ของผู้ใช้")
		}
		return nil, err
	}

	models := sales.ModelSaleV2{}
	err = rs1.Scan(&models.ID,
		&models.UserID,
		&models.SalesCode,
		&models.FirstName,
		&models.LastName,
		&models.RoleID,
		&models.RoleName,
		&models.CardID,
		&models.BookBankID,
		&models.Age,
		&models.URLCardID,
		&models.URLBookBank,
		&models.URLFirstBuy,
		&models.FileURL3,
		&models.InviteCode,
		&models.Confirm,
		&models.CreateDate)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("ไม่มีข้อมูลการสมัค sales ของผู้ใช้")
		}
		logg.Error(err.Error())
		return nil, err
	}
	models.InvitePerson, err = r.FindInviteCodeCount(models.InviteCode)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	node := sales.Node{}
	if models.InvitePerson > 0 && models.RoleID == 2 {
		models.Invite, _, _, _, err = r.FindUerInviteCodeV2(models.InviteCode, &node, 1, 0)
		slice := make([]sales.ModelallSalelevel, 5)
		if err != nil {
			logg.Error(err.Error())
			logg.Println(err.Error())
		}
		models.InvitePerson = node.InvitePerson
		models.Level, _ = r.FindAllsaleLevel(slice, 1, models.InviteCode)
		var commision float64 = 0
		models.Sale, _, _ = r.FindAllSaleByUserID(UserID, 1)
		// user.Sale, commision, _ = r.FindAllSaleByUserID(user.UserID)
		models.Level[0].Allsale = models.Sale
		models.InviteSale, commision, _ = r.CalAllSaleTeam(models.InviteCode)

		models.AllSale, _ = r.SumAllSale(models.Level)
		// commision, _ := r.CalAllSaleTeamCommision(models.InviteCode)
		models.Commission, _, _ = r.GetCommisonSale(commision)
		models.Incentive, _ = r.FindIncentive(models.AllSale)
	}
	return models, nil
}
func (r *saleRepository) CallTokenLine(ID int64) (token string, err error) {
	fmt.Println("repo", ID)
	rs, err := r.db.QueryRow(`SELECT line_token from line_group where id = ?`, ID)
	if err != nil {
		logg.Error(err.Error())
		fmt.Println(err)
		return token, err
	}
	rs.Scan(&token)
	fmt.Println("token:", token)
	return token, nil
}

func (r *saleRepository) SumAllSale(model []sales.ModelallSalelevel) (float64, error) {
	var allsale float64 = 0
	for _, i := range model {
		allsale += i.Allsale
	}
	return allsale, nil
}

func (r *saleRepository) FindAllsaleLevel(model []sales.ModelallSalelevel, level int64, ref_code string) ([]sales.ModelallSalelevel, error) {
	limit, _ := r.GetNumberInvitePersonRefCode(ref_code)
	sql1 := `select ifnull(a.role_id,0) as role_id,
	ifnull(a.user_id,'') as user_id,
	ifnull(a.ref_code,'') as ref_code,
	ifnull(a.invite_code,'') as invite_code from users a 
	left join user_role b on a.role_id = b.id
	where a.ref_code = ? and a.active_status = 1 order by  a.create_time asc limit ` + fmt.Sprintln(limit)
	rs, err := r.db.Query(sql1, ref_code)
	if err != nil {
		logg.Error(err.Error())
		logg.Println("ไม่มีลูก")
		if err == sql.ErrNoRows {

			return nil, nil
		}
		return nil, err
	}
	if level <= 4 {
		model[level].Level = level
	}

	for rs.Next() {
		if level <= 4 {
			user := sales.ModelStructureSalesV2{}
			err = rs.Scan(&user.RoleID, &user.UserID,
				&user.RefCode, &user.InviteCode)
			if err != nil {
				logg.Error(err.Error())
				logg.Println("ไม่มีลูก")
				logg.Println(err.Error())
			}

			allsale, _, _ := r.FindAllSaleByUserID(user.UserID, 1)
			model[level].Allsale += allsale
			r.FindAllsaleLevel(model, level+1, user.InviteCode)

		}
	}
	return model, nil
}

func (r *saleRepository) FindLevelSale(ref_code string, level int64, all_sale float64) (*sales.ModelLevel, float64, float64, error) {

	Limit, _ := r.GetNumberInvitePersonRefCode(ref_code)
	sql1 := `select ifnull(a.role_id,0) as role_id,
		ifnull(a.user_id,'') as user_id,
		ifnull(a.ref_code,'') as ref_code,
		ifnull(a.invite_code,'') as invite_code from users a 
		left join user_role b on a.role_id = b.id
		where a.ref_code = ? and a.active_status = 1 order by  a.create_time asc limit ` + fmt.Sprintln(Limit)
	rs, err := r.db.Query(sql1, ref_code)
	logg.Println(level)
	if err != nil {
		logg.Error(err.Error())
		logg.Println("ไม่มีลูก")
		if err == sql.ErrNoRows {
			logg.Println("ไม่มีลูก")
			return nil, 0, all_sale, nil
		}
		return nil, 0, 0, err
	}

	model := sales.ModelLevel{}
	model.Level = level
	for rs.Next() {
		if level <= 5 {

			user := sales.ModelStructureSalesV2{}
			err = rs.Scan(&user.RoleID, &user.UserID,
				&user.RefCode, &user.InviteCode)
			if err != nil {
				logg.Error(err.Error())

			}
			user.InvitePerson, _ = r.FindInviteCodeCount(user.InviteCode)

			teamall, _, _ := r.FindAllSaleByUserID(user.UserID, 1)
			model.AllSaleLevel += teamall

			if level <= 4 {
				log.Println(model.NextLevel)
				model.NextLevel, model.AllSaleNextLevel, _, err = r.FindLevelSale(user.InviteCode, level+1, model.SumSale)
				if err != nil {
					logg.Error(err.Error())
					logg.Println("ไม่มีลูก")
					return nil, 0, 0, err
				}
			}
			model.SumSale = model.AllSaleNextLevel + model.AllSaleLevel

		}
	}
	all_sale += model.SumSale + all_sale
	return &model, model.SumSale, all_sale, nil

}

func (r *saleRepository) FindUerInviteCodeV2(ref_code string, node *sales.Node, level int64, all_sale float64) ([]sales.ModelStructureSalesV2, *sales.ModelLevel, float64, float64, error) {
	Limit, _ := r.GetNumberInvitePersonRefCode(ref_code)
	sql1 := `select 
	ifnull(a.company_id,0) as company_id,
	ifnull(a.role_id,0) as role_id,
	ifnull(a.user_id,'') as user_id,
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
		logg.Error(err.Error())
		return nil, nil, 0, 0, err
	}
	users := []sales.ModelStructureSalesV2{}

	for rs.Next() {
		if level <= 5 {
			node.InvitePerson += 1
			user := sales.ModelStructureSalesV2{}
			rs.Scan(&user.CompanyID, &user.RoleID, &user.UserID, &user.RoleName, &user.ProfixName, &user.Fname, &user.Lname, &user.Telephone, &user.ActiveStatus,
				&user.Email, &user.RefCode, &user.InviteCode)
			user.InvitePerson, _ = r.FindInviteCodeCount(user.InviteCode)

			// teamall, _ := r.FindAllSaleByUserID(user.UserID)
			// model.AllSaleLevel += teamall
			var commision float64 = 0
			user.Sale, _, _ = r.FindAllSaleByUserID(user.UserID, 1)
			if level <= 4 && user.InvitePerson > 0 {
				user.Invite, _, _, _, _ = r.FindUerInviteCodeV2(user.InviteCode, node, level+1, 0)

				// user.Commission, _, _ = r.GetCommisonSale(user.InviteSale)
				// commision, _ := r.CalAllSaleTeamCommision(user.InviteCode)
				user.InviteSale, commision, _ = r.CalAllSaleTeam(user.InviteCode)
				user.Commission, _, _ = r.GetCommisonSale(commision)
			}

			users = append(users, user)
			// log.Println(level, model)
		}
	}
	// all_sale += model.SumSale + all_sale

	return users, nil, 0, 0, nil

}

func (r *saleRepository) FindIncentive(allsale float64) (float64, error) {
	sql1 := `select id,name_rank,commision_percent from commision where ` + fmt.Sprintln(allsale) + ` >= all_sale_before  and ` + fmt.Sprintln(allsale) + ` <= all_sale_affter limit 1`
	rs1, err := r.db.QueryRow(sql1)
	if err != nil {
		logg.Error(err.Error())
		return 0, err
	}
	model := sales.CommisionModel{}
	rs1.Scan(&model.ID,
		&model.NameRank,
		&model.CommisionPercent)
	commision := allsale * float64(model.CommisionPercent) / 100.00
	return commision, nil
}

// func (r *saleRepository) GetCommisonSale(allsale float64) (float64, error) {
// 	sql1 := `select id,name_rank,commision_percent from commision where ` + fmt.Sprintln(allsale) + ` >= all_sale_before  and ` + fmt.Sprintln(allsale) + ` <= all_sale_affter limit 1`
// 	rs1, err := r.db.QueryRow(sql1)
// 	if err != nil {

// 		return 0, err
// 	}
// 	model := sales.CommisionModel{}
// 	rs1.Scan(&model.ID,
// 		&model.NameRank,
// 		&model.CommisionPercent)
// 	commision := allsale * float64(model.CommisionPercent) / 100.00
// 	return commision, nil
// }

func (r *saleRepository) GetCommisonSale(invite_sale float64) (float64, *sales.CommisionModel, error) {
	sql1 := `select id,name_rank,commision_percent,tax_rate,tax_status from commision_config where id = 1 limit 1`
	rs1, err := r.db.QueryRow(sql1)
	if err != nil {
		logg.Error(err.Error())
		return 0, nil, err
	}
	model := sales.CommisionModel{}
	rs1.Scan(&model.ID,
		&model.NameRank,
		&model.CommisionPercent, &model.TaxRate, &model.TaxStatus)
	commision := (invite_sale * float64(model.CommisionPercent)) / 100.00
	return commision, &model, nil
	//	return (invite_sale * 5) / 100.00, nil
}

func (r *saleRepository) GetProfileSalesbyUser(UserID string) (interface{}, error) {
	sql1 := `SELECT a.id,
	ifnull(a.user_code,'') as user_id, ifnull(a.sales_code,'') as sales_code,
	ifnull(a.first_name,'') as first_name, ifnull(a.last_name,'') as last_name, 
	ifnull(b.role_id,0) as role_id,
	ifnull(c.role_name,'') as role_name,
	ifnull(a.card_id,'') card_id,
	ifnull(a.book_bank_id,'') as book_bank_id,
	ifnull(a.age,0) as age,
	ifnull(a.url_card_id,'') as url_card_id, 
	ifnull(a.url_book_bank,'') as url_book_bank,
	ifnull(a.url_slip_first_buy,'') as url_slip_first_buy,
	ifnull(a.file_url_3,'') as file_url_3,
	ifnull(b.invite_code,'') as invite_code,
	ifnull(a.confirm,0) as confirm,
	ifnull(a.bank_code,'') as bank_code,
	ifnull(a.bank_name,'') as bank_name,
	ifnull(DATE_FORMAT(a.create_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as create_date
	from sales_person a
	left join users b on a.user_code = b.user_id 
	left join user_role c on b.role_id = c.id
	where b.user_id = ? and a.active_status = 0 and b.active_status = 1 limit 1
	`
	rs1, err := r.db.QueryRow(sql1, UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ไม่มีข้อมูลการสมัค sales ของผู้ใช้")
		}
		logg.Error(err.Error())
		return nil, err
	}

	models := sales.ModelSale{}
	err = rs1.Scan(&models.ID,
		&models.UserID,
		&models.SalesCode,
		&models.FirstName,
		&models.LastName,
		&models.RoleID,
		&models.RoleName,
		&models.CardID,
		&models.BookBankID,
		&models.Age,
		&models.URLCardID,
		&models.URLBookBank,
		&models.URLFirstBuy,
		&models.FileURL3,
		&models.InviteCode,
		&models.Confirm,
		&models.BankCode,
		&models.BankName,
		&models.CreateDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ไม่มีข้อมูลการสมัค sales ของผู้ใช้")
		}
		logg.Error(err.Error())
		return nil, err
	}
	models.InvitePerson, err = r.FindInviteCodeCount(models.InviteCode)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	if models.InvitePerson > 0 && models.RoleID == 2 {
		models.Invite, err = r.FindUerInviteCode(models.InviteCode, 1)
		if err != nil {
			logg.Error(err.Error())
			return nil, err
		}
	}
	return models, nil
}

func (r *saleRepository) GetNumberInvitePerson(userID string) (int64, error) {
	var invite_person int64 = 0
	sql1 := `select ifnull(invite_person,10) as invite_person from users where user_id = ? limit 1`
	rs, _ := r.db.QueryRow(sql1, userID)
	rs.Scan(&invite_person)
	if invite_person <= 0 {
		return 10, nil
	}
	return invite_person, nil
}

func (r *saleRepository) GetNumberInvitePersonRefCode(ref_code string) (int64, error) {
	var invite_person int64 = 10
	sql1 := `select ifnull(invite_person,10) as invite_person from users where invite_code = ? limit 1`
	rs, _ := r.db.QueryRow(sql1, ref_code)
	rs.Scan(&invite_person)
	if invite_person <= 0 {
		return 10, nil
	}
	return invite_person, nil
}

func (r *saleRepository) FindUerInviteCode(ref_code string, level int64) ([]sales.ModelStructureSales, error) {
	limit, _ := r.GetNumberInvitePersonRefCode(ref_code)

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
	 where  a.ref_code = ? and a.active_status = 1 order by a.create_time asc limit ` + fmt.Sprintln(limit)
	rs, err := r.db.Query(sql1, ref_code)
	if err != nil {
		logg.Error(err.Error())
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

func (r *saleRepository) FindInviteCodeCount(inviteCode string) (int64, error) {

	var ref_code int64 = 0
	sql := `select count(ref_code) as ref_code from users where ref_code = ? and active_status = 1`
	rs, _ := r.db.QueryRow(sql, inviteCode)
	rs.Scan(&ref_code)
	if ref_code <= 0 {
		return 0, nil
	}
	return ref_code, nil
}

func (r *saleRepository) GetProfileSalesUser(usercode string) (*sales.ModelSale, error) {
	sql1 := `SELECT a.id,
	ifnull(a.user_code,'') as user_id, ifnull(a.sales_code,'') as sales_code,
	ifnull(a.first_name,'') as first_name, ifnull(a.last_name,'') as last_name, 
	ifnull(b.role_id,0) as role_id,
	ifnull(c.role_name,'') as role_name,
	ifnull(a.card_id,'') card_id,
	ifnull(a.book_bank_id,'') as book_bank_id,
	ifnull(a.age,0) as age,
	ifnull(a.url_card_id,'') as url_card_id, 
	ifnull(a.url_book_bank,'') as url_book_bank,
	ifnull(a.url_slip_first_buy,'') as url_slip_first_buy,
	ifnull(a.file_url_3,'') as file_url_3,
	ifnull(b.invite_code,'') as invite_code,
	ifnull(a.confirm,0) as confirm,
	ifnull(a.bank_code,'') as bank_code,
	ifnull(a.bank_name,'') as bank_name,
	ifnull(DATE_FORMAT(a.create_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as create_date
	from sales_person a
	left join users b on a.user_code = b.user_id 
	left join user_role c on b.role_id = c.id
	where b.user_id = ? and a.active_status = 0 and b.active_status = 1 limit 1
	`
	rs1, err := r.db.QueryRow(sql1, usercode)
	if err != nil {
		logg.Error(err.Error())
		logg.Println(err.Error())
		return nil, err
	}
	models := sales.ModelSale{}
	err = rs1.Scan(&models.ID,
		&models.UserID,
		&models.SalesCode,
		&models.FirstName,
		&models.LastName,
		&models.RoleID,
		&models.RoleName,
		&models.CardID,
		&models.BookBankID,
		&models.Age,
		&models.URLCardID,
		&models.URLBookBank,
		&models.URLFirstBuy,
		&models.FileURL3,
		&models.InviteCode,
		&models.Confirm,
		&models.BankCode,
		&models.BankName,
		&models.CreateDate)
	if err != nil {
		logg.Error(err.Error())
		logg.Println(err.Error())
		return nil, err
	}
	return &models, nil
}

func (r *saleRepository) GetProfileSales(SaleCode string) (interface{}, error) {
	sql1 := `SELECT a.id,
	ifnull(a.user_code,'') as user_id, ifnull(a.sales_code,'') as sales_code,
	ifnull(a.first_name,'') as first_name, ifnull(a.last_name,'') as last_name, 
	ifnull(b.role_id,0) as role_id,
	ifnull(c.role_name,'') as role_name,
	ifnull(a.card_id,'') card_id,
	ifnull(a.book_bank_id,'') as book_bank_id,
	ifnull(a.age,0) as age,
	ifnull(a.url_card_id,'') as url_card_id, 
	ifnull(a.url_book_bank,'') as url_book_bank,
	ifnull(a.url_slip_first_buy,'') as url_slip_first_buy,
	ifnull(a.file_url_3,'') as file_url_3,
	ifnull(b.invite_code,'') as invite_code,
	ifnull(a.confirm,0) as confirm,
	ifnull(a.bank_code,'') as bank_code,
	ifnull(a.bank_name,'') as bank_name,
	ifnull(DATE_FORMAT(a.create_time, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as create_date
	from sales_person a
	left join users b on a.user_code = b.user_id 
	left join user_role c on b.role_id = c.id
	where a.sales_code = ? and a.active_status = 0 and b.active_status = 1 limit 1
	`
	rs1, err := r.db.QueryRow(sql1, SaleCode)
	if err != nil {
		logg.Error(err.Error())
		logg.Println(err.Error())
		return nil, err
	}
	models := sales.ModelSale{}
	err = rs1.Scan(&models.ID,
		&models.UserID,
		&models.SalesCode,
		&models.FirstName,
		&models.LastName,
		&models.RoleID,
		&models.RoleName,
		&models.CardID,
		&models.BookBankID,
		&models.Age,
		&models.URLCardID,
		&models.URLBookBank,
		&models.URLFirstBuy,
		&models.FileURL3,
		&models.InviteCode,
		&models.Confirm,
		&models.BankCode,
		&models.BankName,
		&models.CreateDate)
	if err != nil {
		logg.Error(err.Error())
		logg.Println(err.Error())
		return nil, err
	}
	return &models, nil

}

// func (r *saleRepository) GeTProfileMemberRepo(UserID string) (interface{}, error) {
// 	sql := `select
// 	ifnull(a.company_id,0) as company_id,
// 	ifnull(a.role_id,0) as role_id,
// 	ifnull(b.role_name,'') as role_name,
// 	ifnull(a.profix_name,'') as profix_name,
// 	ifnull(a.user_fname,'') as user_fname,
// 	ifnull(a.user_lname,'') as user_lname,
// 	ifnull(a.telephone,'') as telephone,
// 	ifnull(a.active_status,0) as active_status,
// 	ifnull(a.email,'') as email,
// 	ifnull(a.ref_code,'') as ref_code,
// 	ifnull(a.invite_code,'') as invite_code
// 	 from users a
// 	left join user_role b on a.role_id = b.id
// 	 where user_id = ? limit 1`
// 	rs, err := r.db.QueryRow(sql, UserID)
// 	if err != nil {

// 		return nil, err
// 	}
// 	fmt.Println(UserID)
// 	user := member.UserProfileModel{}
// 	rs.Scan(&user.CompanyID, &user.RoleID, &user.RoleName, &user.ProfixName, &user.Fname, &user.Lname, &user.Telephone, &user.ActiveStatus,
// 		&user.Email, &user.RefCode, &user.InviteCode)
// 	fmt.Println(1)
// 	sql1 := `
// 	select
// 		ifnull(id,0) as addr_id,
// 		ifnull(name,'') as addr_name,
// 		ifnull(addr_phone,'') as addr_phone,
// 		ifnull(addr_email,'') as addr_email,
// 		ifnull(addr_state,'') as addr_state,
// 		ifnull(addr_subarea,'') as addr_subarea,
// 		ifnull(addr_district,'') as addr_district,
// 		ifnull(addr_province,'') as addr_province,
// 		ifnull(addr_postal_code,'') as addr_postal_code,
// 		ifnull(main_address,0) as main_address from user_address where ref_id = ? and status = 0
// 	`
// 	rs1, err := r.db.Query(sql1, UserID)
// 	if err != nil {

// 		fmt.Println(err.Error())
// 	}
// 	addrs := []member.AddressProfileModel{}
// 	for rs1.Next() {
// 		addr := member.AddressProfileModel{}
// 		err = rs1.Scan(&addr.AddressID, &addr.Name, &addr.Phone, &addr.Email, &addr.Address, &addr.SubArea, &addr.District, &addr.Province, &addr.PostalCode, &addr.MainAddress)
// 		if err != nil {

// 			return nil, err
// 		}
// 		addrs = append(addrs, addr)
// 	}
// 	user.Address = addrs

// 	if err != nil {

// 		return nil, err
// 	}
// 	return user, nil
// }

func (r *saleRepository) UpdateSaleConfirm(AdminID string, SaleCode string) (interface{}, error) {
	_, err := r.db.Exec(`update sales_person
	set confirm = ?,confirm_time = ?,confirm_by = ?
	where sales_code = ?`,
		1,
		time.Now(),
		AdminID,
		SaleCode,
	)
	if err != nil {
		logg.Error(err.Error())
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}

// func (r *saleRepository) FindAllSaleByUserID(UserID string) (float64, error) {
// 	var all_sale float64 = 0
// 	sql1 := `select sum(total_amount) as all_sale where user_id =? limit 1`
// 	rs1, _ := r.db.QueryRow(sql1, UserID)
// 	rs1.Scan(&all_sale)
// 	return all_sale, nil
// }

func (r *saleRepository) FindInviteUser(UserID string) (string, error) {
	var invite_code string = ""
	sql := `select invite_code from users where user_id = ? and active_status = 1`
	rs, _ := r.db.QueryRow(sql, UserID)
	rs.Scan(&invite_code)
	return invite_code, nil
}

func (r *saleRepository) GetSalesTeamRepo(UserID string) (interface{}, error) {

	return nil, nil

}

func (r *saleRepository) FindInviteCode(inviteCode string) (bool, error) {

	var ref_code int = 0
	sql := `select count(ref_code) as ref_code from users where ref_code = ? and active_status = 1`
	rs, _ := r.db.QueryRow(sql, inviteCode)
	rs.Scan(&ref_code)
	if ref_code > 0 {
		return false, nil
	}
	return true, nil
}

func (r *saleRepository) UpdateSaleRepo(UserID string, sale *sales.UpdateSaleModel) (interface{}, error) {
	_, err := r.UpdateSale(UserID, sale)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) UpdateSale(UserID string, sale *sales.UpdateSaleModel) (interface{}, error) {
	_, err := r.db.Exec(`update sales_person
	set first_name = ?,
	last_name = ?,
	card_id = ?,
	book_bank_id = ?,
	url_card_id = ?,
	url_book_bank = ?,
	url_slip_first_buy = ?,
	file_url_3 = ?,
	confirm = ?,
	bank_code = ?,
	bank_name = ?
	where user_code = ?`,
		sale.Fname,
		sale.Lname,
		sale.CardID,
		sale.BookBankID,
		sale.URLBookBank,
		sale.URLCardID,
		sale.URLFirstBuy,
		sale.File3,
		sale.Confirm,
		sale.BankCode,
		sale.BankName,
		UserID)
	if err != nil {
		logg.Error(err.Error())
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) FindUserdata(UserID string, sale *sales.RegisterSalesModel) (interface{}, error) {
	var user_fname string = ""
	var user_lname string = ""
	sql1 := `select ifnull(user_fname,'') as user_fname,user_lname from users where user_id = ?`
	rs, err := r.db.QueryRow(sql1, UserID)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	err = rs.Scan(&user_fname, &user_lname)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	if user_fname == "" || user_lname == "" {
		return nil, errors.New("ข้อมูล user กรุณากรองข้อมูล ส่วนตัวใน ในช่องข้อมูลส่วนตัว")
	}

	return nil, nil
}

func (r *saleRepository) RegisterSaleRepo(UserID string, sale *sales.RegisterSalesModel) (interface{}, error) {
	_, err := r.FindSales(UserID)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	// _, err = r.FindUserdata(UserID, sale)
	// if err != nil {

	// 	return nil, err
	// }

	fnsale, _ := r.FindSaleCode()
	saleCode, err := GEtSaleCode(fnsale)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	_, err = r.CreateSalePerson(UserID, sale, saleCode)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	sales, err := r.GetProfileSales(saleCode)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}

	return sales, nil
}

func (r *saleRepository) UpdateSalebyAdminRepo(AdminID string, sale *sales.UpdateSaleModelAdmin) (interface{}, error) {
	_, err := r.db.Exec(`update sales_person
	set first_name = ?,
	last_name = ?,
	card_id = ?,
	book_bank_id = ?,
	url_card_id = ?,
	url_book_bank = ?,
	url_slip_first_buy = ?,
	file_url_3 = ?,
	confirm = ?,
	edit_by =?,
	bank_code = ?,
	bank_name = ?
	where sales_code = ?`,
		sale.Fname,
		sale.Lname,
		sale.CardID,
		sale.BookBankID,
		sale.URLBookBank,
		sale.URLCardID,
		sale.URLFirstBuy,
		sale.File3,
		sale.Confirm,
		AdminID,
		sale.BankCode,
		sale.BankName,
		sale.SaleCode)
	if err != nil {
		logg.Error(err.Error())
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) CreateSalePerson(UserID string, sale *sales.RegisterSalesModel, saleCode string) (interface{}, error) {

	log.Println(sale)
	q := r.db.
		InsertInto("sales_person").
		Values(map[string]interface{}{
			"user_code":          UserID,
			"first_name":         sale.Fname,
			"last_name":          sale.Lname,
			"card_id":            sale.CardID,
			"url_card_id":        sale.URLCardID,
			"url_book_bank":      sale.URLBookBank,
			"file_url_3":         sale.File3,
			"url_slip_first_buy": sale.URLFirstBuy,
			"confirm":            0,
			"create_time":        time.Now(),
			"sales_code":         saleCode,
			"book_bank_id":       sale.BookBankID,
			"bank_code":          sale.BankCode,
			"bank_name":          sale.BankName,
		})
	_, err := q.Exec()
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	return nil, nil
}

func (r *saleRepository) UpdateRoleUser(role int64, UserID string, Updateby string) (interface{}, error) {
	_, err := r.db.Exec(`update users
					set role_id = ?,edit_time = ?,
					edit_by = ?
					where user_id = ?`,
		role,
		time.Now(),
		Updateby,
		UserID,
	)
	if err != nil {
		logg.Error(err.Error())
		logg.Println(err.Error())
		return nil, err
	}
	return nil, nil
}

func (repo *saleRepository) FindSales(code string) (bool, error) {
	var user int = 0
	sql := `select count(user_code) as user from sales_person where user_code = ? and active_status = 0`
	rs, _ := repo.db.QueryRow(sql, code)
	rs.Scan(&user)
	if user > 0 {
		return false, errors.New("sales is registed")
	}
	return true, nil
}

func (repo *saleRepository) FindSaleCode() (string, error) {
	like := "SAL" + GetYear() + GetMonth() + GetDay()
	var userid string
	sql := `select COALESCE(sales_code,'SAL62100300001') as userid from sales_person where sales_code like '%` + like + `%' order by id desc LIMIT 1`
	rs, _ := repo.db.QueryRow(sql)
	fmt.Println(sql)
	rs.Scan(&userid)

	return userid, nil
}

func GEtSaleCode(user string) (string, error) {
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
		return "SAL" + GetYear() + GetMonth() + GetDay() + lastuserid, nil
	} else {
		return "SAL" + GetYear() + GetMonth() + GetDay() + "00001", nil
	}
	return "", nil
}
