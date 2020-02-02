package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"
	"gitlab.com/satit13/perfect_api/order"
	"gitlab.com/satit13/perfect_api/product"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

type orderRepo struct {
	db sqlbuilder.Database
}

var (
	orderPendingState int = 1
)

func NewOrderRepository(db *sql.DB) (order.Repository, error) {
	pdb, err := mysql.New(db)
	if err != nil {
		return nil, err
	}
	r := orderRepo{pdb}
	return &r, nil
}

func GenDoc() (string, error) {
	return Year() + Month() + Day() + Sec(), nil
}

func Sec() string {
	var timeHour int
	var timeSecond int
	var timeMinute int
	loc, _ := time.LoadLocation("Asia/Bangkok")
	timeHour = int(time.Now().In(loc).Hour())
	Hour := strconv.Itoa(timeHour)
	if len(Hour) == 1 {
		Hour = "0" + Hour
	}
	timeMinute = time.Now().In(loc).Minute()
	Minute := strconv.Itoa(timeMinute)
	if len(Minute) == 1 {
		Minute = "0" + Minute
	}
	timeSecond = time.Now().In(loc).Second()
	Second := strconv.Itoa(timeSecond)
	if len(Second) == 1 {
		Second = "0" + Second
	}
	fmt.Println("gendoc:", Hour+Hour+Second)
	return Hour + Minute + Second
}

func Year() string {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	var intyear int
	if time.Now().In(loc).Year() >= 2560 {
		intyear = time.Now().In(loc).Year()
	} else {
		intyear = time.Now().In(loc).Year() + 543
	}
	year1 := strconv.Itoa(intyear)
	year2 := year1[2:len(year1)]
	return year2
}

func Month() string {
	var vmonth1 string
	loc, _ := time.LoadLocation("Asia/Bangkok")
	mounth1 := int(time.Now().In(loc).Month())
	mounth2 := int(mounth1)
	vmonth := strconv.Itoa(mounth2)

	if len(vmonth) == 1 {
		vmonth1 = "0" + vmonth
	} else {
		vmonth1 = vmonth
	}
	return vmonth1
}

func Day() string {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	day := time.Now().In(loc).Day()
	days := strconv.Itoa(day)
	if len(days) == 1 {
		days = "0" + strconv.Itoa(day)
	}
	return days
}

func (r *orderRepo) CartAddRepo(req *order.Basket) (resq interface{}, err error) {
	var check_exist_id int64
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now_time := time.Now().In(loc)

	doc_no, err := GenDoc()
	log.Printf("ID: %v", req.ID)
	sqlexist := `select count(user_id) as check_exist from basket where id = ?`
	sqlexist1, err := r.db.QueryRow(sqlexist, req.ID)
	if err != nil {
		return nil, err
	}
	sqlexist1.Scan(&check_exist_id)
	log.Printf("check_exist_id: %v", check_exist_id)

	if check_exist_id == 0 {
		sql := `INSERT INTO basket(user_id,sale_id,doc_no,item_amount,discount_amount,net_amount,my_description,create_time) VALUES (?,?,?,?,?,?,?,?)`
		ressql, err := r.db.Exec(sql,
			req.UserID,
			req.SaleID,
			doc_no,
			req.ItemAmount,
			req.DiscountAmount,
			req.NetAmount,
			req.MyDescription,
			now_time,
		)
		if err != nil {
			return "", err
		}
		resqID, _ := ressql.LastInsertId()
		req.ID = resqID
	} else {
		sql := `Update basket set item_amount=?,discount_amount=?,net_amount=?,my_description=?,edit_time=? where id=?`
		fmt.Println("sql update = ", sql)
		id, err := r.db.Exec(sql,
			req.ItemAmount,
			req.DiscountAmount,
			req.NetAmount,
			req.MyDescription,
			now_time,
			req.ID,
		)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}

		rowAffect, err := id.RowsAffected()
		fmt.Println("Row Affect = ", rowAffect)

		sql_del_sub := `delete from basket_sub where basket_id = ?`
		_, err = r.db.Exec(sql_del_sub, req.ID)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}
	}

	for _, sub := range req.BasketSub {
		sqlsub := `INSERT INTO basket_sub(basket_id,cust_id,sale_id,item_id,item_name,qty,unit_id,unit_code,wh_id,wh_code,price,discount_amount,item_amount,net_amount,line_number) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := r.db.Exec(sqlsub,
			req.ID,
			sub.CustID,
			sub.SaleID,
			sub.ItemID,
			sub.ItemName,
			sub.Qty,
			sub.UnitID,
			sub.UnitCode,
			sub.WhID,
			sub.WhCode,
			sub.Price,
			sub.DiscountAmount,
			sub.ItemAmount,
			sub.NetAmount,
			sub.LineNumber,
		)
		if err != nil {
			return nil, err
		}
	}
	return req.ID, nil
}

func (r *orderRepo) CartAddRepoV2(UserID string, ItemID int64, Qty int64) (interface{}, error) {
	var check_exist_id int64
	var busID int64
	var check_item int64
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now_time := time.Now().In(loc)
	doc_no, err := GenDoc()
	fmt.Println("UserID :", UserID)
	fmt.Println("ItemID :", ItemID)

	sqlitem := `select a.id,a.code,a.name,a.price,a.pic_file_name_1,a.unit_id,a.unit_code,a.def_sale_wh_id,a.def_sale_wh_code,a.use_package
	from item a left JOIN item_category b on a.category_code = b.code
	where a.id = ?`
	rs, err := r.db.QueryRow(sqlitem, ItemID)
	if err != nil {
		return nil, err
	}
	item := product.Item{}
	rs.Scan(&item.ID, &item.Code, &item.Name, &item.Price, &item.PicFileName1, &item.UnitID, &item.UnitCode, &item.DefSaleWhID, &item.DefSaleWhCode, &item.UsePackage)
	if err != nil {
		return nil, err
	}
	// fmt.Println(rs)
	// fmt.Println(item)

	sqlexist := `select count(user_id) as check_exist from basket where user_id = ?`
	sqlexist1, err := r.db.QueryRow(sqlexist, UserID)
	if err != nil {
		return nil, err
	}
	sqlexist1.Scan(&check_exist_id)
	log.Printf("check_exist_id: %v", check_exist_id)

	if check_exist_id == 0 {
		sql := `INSERT INTO basket(user_id,doc_no,create_time) VALUES (?,?,?)`
		ressql, err := r.db.Exec(sql,
			UserID,
			doc_no,
			now_time,
		)
		if err != nil {
			return "", err
		}
		resqID, _ := ressql.LastInsertId()
		busID = resqID
	} else {
		sql := `Update basket set edit_time=? where user_id=?`
		fmt.Println("sql update = ", sql)
		_, err := r.db.Exec(sql,
			now_time,
			UserID,
		)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}
		sqlupdate := `select id from basket where user_id = ?`
		sqlupdate1, err := r.db.QueryRow(sqlupdate, UserID)
		if err != nil {
			return nil, err
		}
		sqlupdate1.Scan(&busID)
		log.Printf("update_id: %v", busID)

		// rowAffect, err := id.RowsAffected()
		// fmt.Println("Row Affect = ", rowAffect)

		// sql_del_sub := `delete from basket_sub where basket_id = ?`
		// _, err = r.db.Exec(sql_del_sub, busID)
		// if err != nil {
		// 	fmt.Println("Error = ", err.Error())
		// 	return nil, err
		// }
	}
	////////////////////////////////////////////////// update sub ////////////////////////////////////////////////////
	sql_check_item := `select count(item_id) as check_exist from basket_sub where item_id = ? and cust_id = ?`
	sql_check_item1, err := r.db.QueryRow(sql_check_item, ItemID, UserID)
	if err != nil {
		return nil, err
	}
	sql_check_item1.Scan(&check_item)
	log.Printf("check_item: %v", check_item)

	if check_item == 0 {
		sqlsub1 := `INSERT INTO basket_sub(basket_id,cust_id,item_id,item_name,qty,unit_id,unit_code,wh_id,wh_code,price,pic_file_name_1,coupon_enabled) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := r.db.Exec(sqlsub1,
			busID,
			UserID,
			item.ID,
			item.Name,
			Qty,
			item.UnitID,
			item.UnitCode,
			item.DefSaleWhID,
			item.DefSaleWhCode,
			item.Price,
			item.PicFileName1,
			item.UsePackage,
		)
		if err != nil {
			return nil, err
		}
	} else {
		sql := `Update basket_sub set qty=(qty+?) where item_id = ? and cust_id = ?`
		// fmt.Println("sql update = ", sql)
		_, err := r.db.Exec(sql,
			Qty,
			ItemID,
			UserID,
		)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}
	}
	return "success", nil
}

func (r *orderRepo) CartStorageRepo(UserID string) (resq []order.Basket, err error) {
	sql := `select a.id,a.user_id,a.sale_id,a.doc_no,a.item_amount,a.discount_amount,a.net_amount,ifnull(a.my_description,'') as my_description ,a.create_time
	from basket a
	where a.user_id = ?`
	items := []order.Basket{}
	rs, err := r.db.Query(sql, UserID)
	if err != nil {
		return resq, err
	}
	for rs.Next() {
		item := order.Basket{}
		err = rs.Scan(&item.ID, &item.UserID, &item.SaleID, &item.DocNo, &item.ItemAmount, &item.DiscountAmount, &item.NetAmount,
			&item.MyDescription, &item.CreateTime)
		if err != nil {
			return resq, err
		}
		fmt.Println(UserID)

		sql1 := `select a.basket_id,a.cust_id,a.sale_id,a.item_id,a.item_name,a.qty,a.unit_id,a.unit_code,a.wh_id ,a.wh_code,a.price,a.discount_amount,
		a.item_amount,a.net_amount,a.pic_file_name_1,a.coupon_enabled
		from basket_sub a
		where a.cust_id = ?`
		rs1, err := r.db.Query(sql1, UserID)
		if err != nil {
			return resq, err
		}
		for rs1.Next() {
			item_sub := order.BasketSub{}
			err = rs1.Scan(&item_sub.BasketID, &item_sub.CustID, &item_sub.SaleID, &item_sub.ItemID, &item_sub.ItemName, &item_sub.Qty, &item_sub.UnitID,
				&item_sub.UnitCode, &item_sub.WhID, &item_sub.WhCode, &item_sub.Price, &item_sub.DiscountAmount, &item_sub.ItemAmount,
				&item_sub.NetAmount, &item_sub.PicFileName1, &item_sub.CouponEnabled)
			if err != nil {
				return resq, err
			}
			item.BasketSub = append(item.BasketSub, item_sub)
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *orderRepo) CartDeleteRepo(UserID string) (resq interface{}, err error) {
	fmt.Println("UserID :", UserID)
	sql_del_sub := `delete from basket_sub where cust_id = ?`
	_, err = r.db.QueryRow(sql_del_sub, UserID)
	if err != nil {
		return "fail", err
	}
	return "success", nil
}

func (r *orderRepo) CartDeleteByItemRepo(ItemID int64, UserID string) (resq interface{}, err error) {
	fmt.Println("UserID :", UserID)
	sql_del_sub := `delete from basket_sub where item_id = ? and cust_id = ?`
	_, err = r.db.QueryRow(sql_del_sub, ItemID, UserID)
	if err != nil {
		return "fail", err
	}
	return "success", nil
}

func (r *orderRepo) CartQTYRepo(ItemID int64, QTY int64, UserID string) (interface{}, error) {
	fmt.Println("UserID :", UserID)
	sql := `Update basket_sub set qty = ? where item_id = ? and cust_id = ?`
	fmt.Println("sql update = ", sql)
	_, err := r.db.Exec(sql,
		QTY,
		ItemID,
		UserID,
	)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	return "success", nil
}

func (r *orderRepo) CalPriceBK(ItemID []int64, DiscountAmount float64, DistanceAmount int64, busID int64) (TotolSumPrice float64, err error) {
	var SumPrice float64
	for _, sub := range ItemID {
		fmt.Println("busket.ID:", busID)
		sql_cal_busket := `select qty,price,discount_amount from basket_sub where item_id = ? and basket_id = ?`
		cal_rs, err := r.db.QueryRow(sql_cal_busket, sub, busID)
		if err != nil {
			return TotolSumPrice, err
		}
		sum_busket_sub := order.BasketSub{}
		err = cal_rs.Scan(&sum_busket_sub.Qty, &sum_busket_sub.Price, &sum_busket_sub.DiscountAmount)
		fmt.Println("sub:", sub)
		fmt.Println("sum_busket_sub.Price :", sum_busket_sub.Price)
		SumPrice = SumPrice + (sum_busket_sub.Price * sum_busket_sub.Qty)
	}
	fmt.Println("SumPrice :", SumPrice)
	TotolSumPrice = SumPrice - DiscountAmount + float64(DistanceAmount)
	fmt.Println("TotolSumPrice :", TotolSumPrice)

	return TotolSumPrice, nil
}

func (r *orderRepo) CartToOrderRepo(UserID string, ItemID []int64, DiscountAmount float64, DeliveryCode string, DistanceAmount int64, Coupon string) (interface{}, error) {
	// var check_exist_id int64
	var GetID int64
	var SumPrice float64
	var CalDiscount float64
	var CalSumPrice float64
	var TotolSumPrice float64
	var IsPackage int
	var chackUsePack int
	pushPackage1 := 0
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now_time := time.Now().In(loc)
	doc_no, err := GenDoc()
	doc_no = "ORD" + doc_no
	fmt.Println("UserID =", UserID)
	ItemCou := order.CouponDetail{}
	if Coupon != "" {
		fmt.Println("Coupon =", Coupon)
		sqlTypePk := `select id,doc_date,coupon_no,coupon_type,user_id,name,value,remain,expire_status,begin_date,expire_date from coupon where coupon_no = ?`
		sqlTypePk1, err := r.db.QueryRow(sqlTypePk, Coupon)
		if err != nil {
			return nil, err
		}
		err = sqlTypePk1.Scan(&ItemCou.ID, &ItemCou.DocDate, &ItemCou.CouponNo, &ItemCou.CouponType, &ItemCou.UserID, &ItemCou.Name, &ItemCou.Value,
			&ItemCou.Remain, &ItemCou.ExpireStatus, &ItemCou.BeginDate, &ItemCou.ExpireDate,
		)
		// fmt.Println("getcp")
		// if ItemCou.CouponType == 1 {
		for _, sub1 := range ItemID {
			sql_pk := `select use_package from item where id = ?`
			checkPK, err := r.db.QueryRow(sql_pk, sub1)
			if err != nil {
				return nil, err
			}
			err = checkPK.Scan(&chackUsePack)
			if chackUsePack == 0 {
				return nil, errors.New("มีสินค้าที่ไม่ตรงตามเงื่อนไขคูปอง")
			}
		}
		// }
		// fmt.Println("getcp1")
	}

	for _, sub1 := range ItemID {
		sql_pk := `select is_package from item where id = ?`
		checkPK, err := r.db.QueryRow(sql_pk, sub1)
		if err != nil {
			return nil, err
		}
		err = checkPK.Scan(&IsPackage)
		if IsPackage == 1 {
			pushPackage1 = 1
			break
		}
	}

	sqluser := `select user_fname,user_lname from users where user_id = ?`
	sqluser1, err := r.db.QueryRow(sqluser, UserID)
	if err != nil {
		return nil, err
	}
	UseDetail := order.UseOrder{}
	sqluser1.Scan(&UseDetail.UserFname, &UseDetail.UserLname)
	if err != nil {
		return nil, err
	}
	Username := UseDetail.UserFname + " " + UseDetail.UserLname
	// fmt.Println("Username :", Username)
	// fmt.Println("Username :", UserID)
	sqlitem := `select a.id,a.user_id,a.sale_id,a.doc_no,a.item_amount,a.discount_amount,a.net_amount,ifnull(a.my_description,'') as my_description
	from basket a
	where a.user_id = ?`
	rs, err := r.db.QueryRow(sqlitem, UserID)
	if err != nil {
		return nil, err
	}
	busket := order.Basket{}
	err = rs.Scan(&busket.ID, &busket.UserID, &busket.SaleID, &busket.DocNo, &busket.ItemAmount, &busket.DiscountAmount, &busket.NetAmount, &busket.MyDescription)
	if err != nil {
		return nil, err
	}
	fmt.Println("busket:", busket)
	/////////////////////////////cal price////////////////////////////////
	for _, sub := range ItemID {
		// fmt.Println("busket.ID:", busket.ID)
		sql_cal_busket := `select qty,price,discount_amount from basket_sub where item_id = ? and basket_id = ?`
		cal_rs, err := r.db.QueryRow(sql_cal_busket, sub, busket.ID)
		if err != nil {
			return nil, err
		}
		sum_busket_sub := order.BasketSub{}
		err = cal_rs.Scan(&sum_busket_sub.Qty, &sum_busket_sub.Price, &sum_busket_sub.DiscountAmount)
		// fmt.Println("sum_busket_sub.Qty", sum_busket_sub.Qty)
		// fmt.Println("sum_busket_sub.Price =", sum_busket_sub.Price)
		SumPrice = SumPrice + (sum_busket_sub.Price * sum_busket_sub.Qty)
	}
	fmt.Println("SumPrice =", SumPrice)
	if ItemCou.CouponType == 2 {
		CalDiscount = (SumPrice * (ItemCou.Value / 100))
		CalSumPrice = SumPrice - CalDiscount
	} else {
		CalSumPrice = SumPrice
	}
	fmt.Println("CalSumPrice =", CalSumPrice)

	if ItemCou.CouponType == 1 {
		CalDiscount = SumPrice
		TotolSumPrice = float64(DistanceAmount)
	} else {
		// TotolSumPrice = SumPrice - DiscountAmount + float64(DistanceAmount)
		TotolSumPrice = CalSumPrice + float64(DistanceAmount)
	}
	fmt.Println("TotolSumPrice =", TotolSumPrice)
	fmt.Println("CouponType", ItemCou.CouponType)
	fmt.Println("CalDiscount", CalDiscount)
	fmt.Println("ItemCou.Remain", ItemCou.Remain)
	fmt.Println("TotolSumPrice", TotolSumPrice)
	//////////////////////////////////////////////////////////////////////Push/////////////////////////////////////////////////////////////////////////
	fmt.Println("pushPackage1 :", pushPackage1)
	sql := `INSERT INTO orders(user_id,user_name,sale_id,sale_name,doc_no,doc_date,sum_of_item_amount,discount_amount,after_discount_amount,before_tax,
		tax_amount,total_amount,sum_cash_amount,sum_credit_amount,sum_deposit_amount,sum_coupon_amount,sum_bank_amount,change_amount,
		net_debt_amount,my_description,delivery_date,distance,distance_amount,delivery_id,referral_id,create_by,create_time,is_package,coupon,order_status)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1)`
	sql1, err := r.db.Exec(sql,
		UserID,         //user_id
		Username,       //user_name
		busket.SaleID,  //sale_id
		"",             //sale_name
		doc_no,         //doc_no
		now_time,       //doc_date
		SumPrice,       // SumOfItemAmount,
		CalDiscount,    // DiscountAmount,
		TotolSumPrice,  // AfterDiscountAmount,
		0,              // BeforeTax,
		0,              // TaxAmount,
		TotolSumPrice,  // TotalAmount,
		0,              // SumCashAmount,
		0,              // SumCreditAmount,
		0,              // SumDepositAmount,
		CalDiscount,    // SumCouponAmount,
		TotolSumPrice,  // SumBankAmount,
		0,              // ChangeAmount,
		0,              // NetDebtDmount,
		"",             // MyDescription,
		now_time,       // DeliveryDate,
		0,              // Distance,
		DistanceAmount, // DistanceAmount,
		DeliveryCode,   // DeliveryID,
		0,              // ReferralID,
		UserID,         //create_by,
		now_time,       // create_time,
		pushPackage1,   // is_package,
		Coupon,
	)
	// log.Printf("sql1:", sql1)
	if err != nil {
		return nil, err
	}
	resqID, _ := sql1.LastInsertId()
	GetID = resqID

	for _, sub := range ItemID {
		sql_sub_busket := `select a.basket_id,a.cust_id,a.sale_id,a.item_id,a.item_name,a.qty,a.unit_id,a.unit_code,a.wh_id ,a.wh_code,a.price,a.discount_amount,
		a.item_amount,a.net_amount,a.pic_file_name_1
		from basket_sub a
		where a.item_id = ? and a.cust_id = ?`
		rs1, err := r.db.QueryRow(sql_sub_busket, sub, UserID)
		if err != nil {
			return nil, err
		}
		busket_sub := order.BasketSub{}
		err = rs1.Scan(&busket_sub.BasketID, &busket_sub.CustID, &busket_sub.SaleID, &busket_sub.ItemID, &busket_sub.ItemName, &busket_sub.Qty, &busket_sub.UnitID,
			&busket_sub.UnitCode, &busket_sub.WhID, &busket_sub.WhCode, &busket_sub.Price, &busket_sub.DiscountAmount, &busket_sub.ItemAmount,
			&busket_sub.NetAmount, &busket_sub.PicFileName1)
		if err != nil {
			return nil, err
		}
		// log.Printf("busket_sub:", busket_sub)

		sqlsub1 := `INSERT INTO order_sub(order_id,item_id,item_name,wh_id,wh_code,shelf_id,shelf_code,qty,cn_qty,unit_id,unit_code,price,discount_amount,item_amount,
		average_cost,sum_of_cost,rate_1,stock_type,basket_id,item_sub_cat,item_group,type_code,point,is_cancel,ref_line_number,pic_file_name_1,line_number) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = r.db.Exec(sqlsub1,
			GetID,
			busket_sub.ItemID,
			busket_sub.ItemName,
			busket_sub.WhID,
			busket_sub.WhCode,
			0,  //busket_sub.ShelfID,
			"", //busket_sub.ShelfCode,
			busket_sub.Qty,
			0, //busket_sub.CnQty,
			busket_sub.UnitID,
			busket_sub.UnitCode,
			busket_sub.Price,
			busket_sub.DiscountAmount,
			busket_sub.ItemAmount,
			0, //busket_sub.AverageCost,
			0, //busket_sub.SumOfCost,
			0, //busket_sub.Rate1,
			0, //busket_sub.StockType,
			busket_sub.BasketID,
			0, //busket_sub.Itembusket_subCat,
			0, //busket_sub.ItemGroup,
			0, //busket_sub.TypeCode,
			0, //busket_sub.Point,
			0, //busket_sub.IsCancel,
			0, //busket_sub.RefLineNumber,
			busket_sub.PicFileName1,
			busket_sub.LineNumber,
		)
		// log.Printf("sqlbksub:", sqlsub)
		if err != nil {
			return nil, err
		}
	}

	if Coupon != "" {
		// sqlTypePk := `select id,doc_date,coupon_no,coupon_type,user_id,name,value,remain,expire_status,begin_date,expire_date from coupon where coupon_no = ?`
		// sqlTypePk1, err := r.db.QueryRow(sqlTypePk, Coupon)
		// if err != nil {
		// 	return nil, err
		// }

		// ItemCou := order.CouponDetail{}
		// err = sqlTypePk1.Scan(&ItemCou.ID, &ItemCou.DocDate, &ItemCou.CouponNo, &ItemCou.CouponType, &ItemCou.UserID, &ItemCou.Name, &ItemCou.Value,
		// 	&ItemCou.Remain, &ItemCou.ExpireStatus, &ItemCou.BeginDate, &ItemCou.ExpireDate,
		// )

		if ItemCou.CouponType == 1 {
			sqladdC := `INSERT INTO coupon_receive(inv_id,inv_no,user_id,coupon_id,coupon_no,doc_date,coupon_value,is_cancel,line_number) VALUES (?,?,?,?,?,?,?,?,?)`
			_, err := r.db.Exec(sqladdC,
				GetID,
				doc_no,
				UserID,
				ItemCou.ID,
				ItemCou.CouponNo,
				now_time,
				CalSumPrice,
				0,
				1,
			)
			// log.Printf("sqlbksub:", sqladdC1)
			if err != nil {
				return nil, err
			}
			////////////////////////update coupon////////////////////////////////
			RemainCP := ItemCou.Remain - CalSumPrice
			sqlupcp := `Update coupon set remain=?,edit_time=? where coupon_no=?`
			_, err = r.db.Exec(sqlupcp,
				RemainCP,
				now_time,
				Coupon,
			)
			if err != nil {
				fmt.Println("Error = ", err.Error())
				return nil, err
			}
		} else {
			sqladdC := `INSERT INTO coupon_receive(inv_id,inv_no,user_id,coupon_id,coupon_no,doc_date,coupon_value,is_cancel,line_number) VALUES (?,?,?,?,?,?,?,?,?)`
			_, err = r.db.Exec(sqladdC,
				GetID,
				doc_no,
				UserID,
				ItemCou.ID,
				ItemCou.CouponNo,
				now_time,
				TotolSumPrice,
				0,
				1,
			)
			// log.Printf("sqlbksub:", sqladdC1)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, BkSubDelete := range ItemID {
		sql_del_crd := `delete from basket_sub where item_id=? and cust_id=?`
		_, err = r.db.Exec(sql_del_crd, BkSubDelete, UserID)
		if err != nil {
			return nil, err
		}
	}

	return map[string]interface{}{
		"id": GetID,
	}, nil
}

func (r *orderRepo) CalItem(ItemID []int64, busID int64) (SumPrice float64, err error) {

	for _, sub := range ItemID {
		sql_sub_busket := `select a.qty,a.price,a.discount_amount
		from basket_sub a
		where a.item_id = ? and a.basket_id = ?`
		rs1, err := r.db.QueryRow(sql_sub_busket, sub, busID)
		if err != nil {
			return SumPrice, err
		}
		busket_sub := order.BasketSub{}
		err = rs1.Scan(&busket_sub.Qty, &busket_sub.Price, &busket_sub.DiscountAmount)
		if err != nil {
			return SumPrice, err
		}
		log.Printf("busket_sub:", busket_sub)

		SumPrice = SumPrice + (busket_sub.Price * busket_sub.Qty)
	}
	fmt.Println("SumPrice", SumPrice)

	return SumPrice, nil
}

func (r *orderRepo) OrderAllRepo(UserID string) (resq []order.Order, err error) {
	fmt.Println(UserID)
	items := []order.Order{}
	sql := `select a.id, a.user_id, a.user_name, a.sale_id, a.sale_name, a.doc_no, a.sum_of_item_amount, a.discount_amount, a.after_discount_amount, 
	a.before_tax, a.tax_amount, a.total_amount, a.sum_cash_amount, a.sum_credit_amount, a.sum_deposit_amount, a.sum_coupon_amount, 
	a.sum_bank_amount, a.change_amount, a.net_debt_amount, a.my_description, a.order_status,a.delivery_date, a.distance, a.distance_amount, 
	ifnull(a.delivery_link,'') as delivery_link , a.delivery_id, a.referral_id, ifnull(a.pic_slip,'') as pic_slip, a.create_time,  ifnull(a.tracking_id,'') as tracking_id 
	from orders a
	where a.user_id = ? `
	rs, err := r.db.Query(sql, UserID)
	if err != nil {
		return resq, err
	}

	for rs.Next() {
		item := order.Order{}
		err = rs.Scan(&item.ID, &item.UserID, &item.UserName, &item.SaleID, &item.SaleName, &item.DocNo, &item.SumOfItemAmount, &item.DiscountAmount, &item.AfterDiscountAmount,
			&item.BeforeTax, &item.TaxAmount, &item.TotalAmount,
			&item.SumCashAmount, &item.SumCreditAmount, &item.SumDepositAmount, &item.SumCouponAmount, &item.SumBankAmount, &item.ChangeAmount,
			&item.NetDebtDmount, &item.MyDescription, &item.OrderStatus, &item.DeliveryDate,
			&item.Distance, &item.DistanceAmount, &item.DeliveryLink, &item.DeliveryID, &item.ReferralID, &item.PicSlip, &item.CreateTime, &item.TrackingID,
		)
		if err != nil {
			return resq, err
		}
		// fmt.Println("OrderStatus :", item.OrderStatus)
		// sql_send_stauus := `select a.send_text
		// from send_status a
		// where a.id = ?`
		// rs_send_stauus, err := r.db.QueryRow(sql_send_stauus, &item.OrderStatus)
		// if err != nil {
		// 	return resq, err
		// }
		// err = rs_send_stauus.Scan(&item.SendText)
		// if err != nil {
		// 	return resq, err
		// }

		for i := 1; i <= int(item.OrderStatus); i++ {
			var SubStatus string

			sql_sub_status := `select a.send_text from send_status a where a.id = ?`
			sql_sub_status1, err := r.db.QueryRow(sql_sub_status, int(i))
			if err != nil {
				return nil, err
			}
			err = sql_sub_status1.Scan(&SubStatus)
			item.SendText = append(item.SendText, SubStatus)
			// fmt.Println("i:", i)
			// fmt.Println("SubStatus:", SubStatus)
		}

		sql1 := `select a.id ,a.order_id ,a.item_id ,a.item_name ,a.wh_id ,a.wh_code ,a.shelf_id ,a.shelf_code ,a.qty ,a.cn_qty ,a.unit_id ,a.unit_code ,
		a.price ,a.discount_amount ,a.item_amount ,a.average_cost ,a.sum_of_cost ,a.rate_1 ,a.stock_type ,a.basket_id ,a.item_sub_cat ,a.item_group ,
		a.type_code ,a.point ,a.is_cancel ,a.ref_line_number , pic_file_name_1, a.line_number
		from order_sub a
		where a.order_id = ?`
		rs1, err := r.db.Query(sql1, item.ID)
		if err != nil {
			return resq, err
		}
		for rs1.Next() {
			item_sub := order.OrderSub{}
			err = rs1.Scan(&item_sub.ID, &item_sub.OrderID, &item_sub.ItemID, &item_sub.ItemName, &item_sub.WhID, &item_sub.WhCode, &item_sub.ShelfID, &item_sub.ShelfCode,
				&item_sub.Qty, &item_sub.CnQty, &item_sub.UnitID, &item_sub.UnitCode, &item_sub.Price, &item_sub.DiscountAmount, &item_sub.ItemAmount,
				&item_sub.AverageCost, &item_sub.SumOfCost, &item_sub.Rate1, &item_sub.StockType, &item_sub.BasketID, &item_sub.ItemSubCat, &item_sub.ItemGroup,
				&item_sub.TypeCode, &item_sub.Point, &item_sub.IsCancel, &item_sub.RefLineNumber, &item_sub.PicFileName1, &item_sub.LineNumber,
			)
			if err != nil {
				return resq, err
			}
			item.OrderSub = append(item.OrderSub, item_sub)
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *orderRepo) OrderByIDRepo(ID int64) (resq []order.Order, err error) {
	fmt.Println("ID", ID)
	sql := `select a.id, a.user_id, a.user_name, a.sale_id, a.sale_name, a.doc_no, a.sum_of_item_amount, a.discount_amount, a.after_discount_amount, 
	a.before_tax, a.tax_amount, a.total_amount, a.sum_cash_amount, 
	a.sum_credit_amount, a.sum_deposit_amount, a.sum_coupon_amount, a.sum_bank_amount, a.change_amount, a.net_debt_amount, a.my_description, a.order_status,
	a.delivery_date, a.distance, a.distance_amount, ifnull(a.delivery_link,'') as delivery_link , a.delivery_id, 
	a.referral_id,  ifnull(a.pic_slip,'') as pic_slip, a.create_time, ifnull(a.tracking_id,'') as tracking_id
	from orders a
	where a.id = ?`
	items := []order.Order{}
	rs, err := r.db.QueryRow(sql, ID)
	if err != nil {
		return resq, err
	}
	item := order.Order{}
	err = rs.Scan(&item.ID, &item.UserID, &item.UserName, &item.SaleID, &item.SaleName, &item.DocNo, &item.SumOfItemAmount, &item.DiscountAmount, &item.AfterDiscountAmount,
		&item.BeforeTax, &item.TaxAmount, &item.TotalAmount,
		&item.SumCashAmount, &item.SumCreditAmount, &item.SumDepositAmount, &item.SumCouponAmount, &item.SumBankAmount, &item.ChangeAmount,
		&item.NetDebtDmount, &item.MyDescription, &item.OrderStatus, &item.DeliveryDate,
		&item.Distance, &item.DistanceAmount, &item.DeliveryLink, &item.DeliveryID, &item.ReferralID, &item.PicSlip, &item.CreateTime, &item.TrackingID,
	)
	if err != nil {
		return resq, err
	}
	fmt.Println("OrderStatus :", item.OrderStatus)
	// sql_send_stauus := `select a.send_text
	// from send_status a
	// where a.id = ?`
	// rs_send_stauus, err := r.db.QueryRow(sql_send_stauus, &item.OrderStatus)
	// if err != nil {
	// 	return resq, err
	// }
	// err = rs_send_stauus.Scan(&item.SendText)
	// if err != nil {
	// 	return resq, err
	// }
	if item.OrderStatus == 99 {
		var SubStatus string

		sql_sub_status := `select a.send_text from send_status a where a.id = ?`
		sql_sub_status1, err := r.db.QueryRow(sql_sub_status, item.OrderStatus)
		if err != nil {
			return nil, err
		}
		err = sql_sub_status1.Scan(&SubStatus)
		item.SendText = append(item.SendText, SubStatus)
		fmt.Println("SubStatus:", SubStatus)
	} else {
		for i := 1; i <= int(item.OrderStatus); i++ {
			var SubStatus string

			sql_sub_status := `select a.send_text from send_status a where a.id = ?`
			sql_sub_status1, err := r.db.QueryRow(sql_sub_status, int(i))
			if err != nil {
				return nil, err
			}
			err = sql_sub_status1.Scan(&SubStatus)
			item.SendText = append(item.SendText, SubStatus)
			fmt.Println("i:", i)
			fmt.Println("SubStatus:", SubStatus)
		}
	}

	sql1 := `select a.id ,a.order_id ,a.item_id ,a.item_name ,a.wh_id ,a.wh_code ,a.shelf_id ,a.shelf_code ,a.qty ,a.cn_qty ,a.unit_id ,a.unit_code ,
		a.price ,a.discount_amount ,a.item_amount ,a.average_cost ,a.sum_of_cost ,a.rate_1 ,a.stock_type ,a.basket_id ,a.item_sub_cat ,a.item_group ,
		a.type_code ,a.point ,a.is_cancel ,a.ref_line_number , pic_file_name_1, a.line_number
		from order_sub a
		where a.order_id = ?`
	rs1, err := r.db.Query(sql1, item.ID)
	if err != nil {
		return resq, err
	}
	for rs1.Next() {
		item_sub := order.OrderSub{}
		err = rs1.Scan(&item_sub.ID, &item_sub.OrderID, &item_sub.ItemID, &item_sub.ItemName, &item_sub.WhID, &item_sub.WhCode, &item_sub.ShelfID, &item_sub.ShelfCode,
			&item_sub.Qty, &item_sub.CnQty, &item_sub.UnitID, &item_sub.UnitCode, &item_sub.Price, &item_sub.DiscountAmount, &item_sub.ItemAmount,
			&item_sub.AverageCost, &item_sub.SumOfCost, &item_sub.Rate1, &item_sub.StockType, &item_sub.BasketID, &item_sub.ItemSubCat, &item_sub.ItemGroup,
			&item_sub.TypeCode, &item_sub.Point, &item_sub.IsCancel, &item_sub.RefLineNumber, &item_sub.PicFileName1, &item_sub.LineNumber,
		)
		if err != nil {
			return resq, err
		}
		item.OrderSub = append(item.OrderSub, item_sub)
	}
	items = append(items, item)

	return items, nil
}

func (r *orderRepo) OrderCancelRepo(UserID string, ID int64) (interface{}, error) {
	var check_confirm int64
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(loc)

	sql_get := `select order_status, a.id, a.doc_no, ifnull(a.tracking_id,'') as tracking_id
	from orders a
	where a.id = ?`
	rs1, err := r.db.QueryRow(sql_get, ID)
	item := order.Order{}
	err = rs1.Scan(&check_confirm, &item.ID, &item.DocNo, &item.TrackingID)
	if err != nil {
		return nil, err
	}

	// sqlexist := `select order_status from orders where id = ?`
	// sqlexist1, err := r.db.QueryRow(sqlexist, ID)
	// if err != nil {
	// 	return nil, err
	// }
	// sqlexist1.Scan(&check_confirm)

	if check_confirm != 1 {
		return "รายการนี้ถูกนำไปใช้งานแล้ว", nil
	}

	sql := `Update orders set order_status=99, cancel_by=?, cancel_time=? where id=?`
	id, err := r.db.Exec(sql,
		UserID,
		now,
		ID,
	)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	fmt.Println("id", id)

	sql_track := `INSERT INTO track_order_status(order_id,doc_no,tracking_id,status_id,update_status_by,status_date) VALUES (?,?,?,?,?,?)`
	id_track, err := r.db.Exec(sql_track,
		item.ID,
		item.DocNo,
		item.TrackingID,
		99,
		UserID,
		now,
	)
	fmt.Println("id_track", id_track)

	return "success", nil
}

func (r *orderRepo) OrderConfirmRepo(UserID string, ID int64, URL string, AddressID int64) (resp string, Fname string, Lname string, OrAmount float64, DocNo string, err error) {
	var check_confirm int64
	// ConfirmCP := 0
	// var Fname string
	// var Lname string
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(loc)

	// fmt.Println("sql update = ", UserID, ID, AddressID, URL)
	sqlexist := `select order_status, id,total_amount,doc_no, ifnull(tracking_id,'') as tracking_id from orders where id = ?`
	sqlexist1, err := r.db.QueryRow(sqlexist, ID)
	item := order.Order{}
	if err != nil {
		return "", "", "", OrAmount, "", err
	}
	sqlexist1.Scan(&check_confirm, &item.ID, &OrAmount, &DocNo, &item.TrackingID)

	sqluser := `select user_fname,user_lname from users where user_id = ?`
	sqluser1, err := r.db.QueryRow(sqluser, UserID)
	if err != nil {
		return "", "", "", OrAmount, "", err
	}
	sqluser1.Scan(&Fname, &Lname)

	if check_confirm >= 2 {
		return "", "", "รายการนี้ถูกนำไปใช้งานแล้ว", OrAmount, "", nil
	}

	sql := `update orders set pic_slip=?, order_status=?, address_id = ?, edit_by=?, edit_time=? where id=?`
	id, err := r.db.Exec(sql, URL, orderPendingState, AddressID, UserID, now, ID)
	if err != nil {
		fmt.Println("Error1= ", err.Error())
		return "", "", "", OrAmount, "", err
	}
	fmt.Println("id = ", id)

	sql_track := `INSERT INTO track_order_status(order_id,doc_no,tracking_id,status_id,update_status_by,status_date) VALUES (?,?,?,?,?,?)`
	id_track, err := r.db.Exec(sql_track,
		item.ID,
		DocNo,
		item.TrackingID,
		orderPendingState,
		UserID,
		now,
	)
	fmt.Println("id_track", id_track)

	return "success", Fname, Lname, OrAmount, DocNo, nil
}

func (r *orderRepo) BankRepo() ([]order.Bank, error) {
	fmt.Println("start repo FindCategoryRepo")
	sql := `select a.id,a.name,a.code,b.bank_code,b.bank_name,b.code as code_name,b.image
	from define_bank a left JOIN bank_code b ON a.bank_id = b.id`
	banks := []order.Bank{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		bank := order.Bank{}

		err = rs.Scan(&bank.ID, &bank.Name, &bank.Code, &bank.BankCode, &bank.BankName, &bank.CodeName, &bank.Image)
		if err != nil {
			return nil, err
		}
		banks = append(banks, bank)
	}
	if err != nil {
		return nil, err
	}
	fmt.Println("return ", banks)
	return banks, nil
}

func (r *orderRepo) BankList() ([]order.BankList, error) {
	sql := `select id,ifnull(bank_code,'') as bank_code , ifnull(bank_name,'') as bank_name , ifnull(code,'') as code, ifnull(description,'') as description, ifnull(image,'') as image 
	from bank_code
	where active = 1`
	banks := []order.BankList{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		bank := order.BankList{}

		err = rs.Scan(&bank.ID, &bank.BankCode, &bank.BankName, &bank.Code, &bank.Description, &bank.Image)
		if err != nil {
			return nil, err
		}
		banks = append(banks, bank)
	}
	if err != nil {
		return nil, err
	}
	fmt.Println("return ", banks)
	return banks, nil
}

func (r *orderRepo) UpdateStatusRepo(UserID string, ID int64, Status int64, TrackingID string) (interface{}, error) {
	var StatusPackge []int64
	var OrderUser string
	var pricePK float64
	var namePK string
	var package_discount float64
	fmt.Println("UserID = ", UserID)
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(loc)
	after := now.AddDate(1, 0, 0)
	doc_no, err := GenDoc()
	if err != nil {
		return nil, err
	}
	doc_no_cp := "CPV" + doc_no
	doc_no_cp1 := "CPP" + doc_no

	sqlexist := `select id, doc_no from orders where id = ?`
	sqlexist1, err := r.db.QueryRow(sqlexist, ID)
	item := order.Order{}
	if err != nil {
		return nil, err
	}
	sqlexist1.Scan(&item.ID, &item.DocNo)

	if Status == 2 || Status == 5 {
		sql_detail := `select a.user_id,c.coupon_amount,c.name,c.package_discount 
		from orders a left join order_sub b on a.id=b.order_id left join item c on b.item_id=c.id 
		where a.id=?`
		OrderUser_rs, err := r.db.QueryRow(sql_detail, ID)
		if err != nil {
			return nil, err
		}
		OrderUser_rs.Scan(&OrderUser, &pricePK, &namePK, &package_discount)

		sql_item := `select is_package from orders where id=?`
		cal_rs, err := r.db.Query(sql_item, ID)
		if err != nil {
			return nil, err
		}
		for cal_rs.Next() {
			var StatusPackgesSub int64
			err = cal_rs.Scan(&StatusPackgesSub)
			if err != nil {
				return nil, err
			}
			StatusPackge = append(StatusPackge, StatusPackgesSub)
		}
		if err != nil {
			return nil, err
		}
		fmt.Println("StatusPackge:", StatusPackge)

		for _, sub := range StatusPackge {
			if sub == 1 {
				sql_up := `update users set dealer=1 where user_id=?`
				iduser, err := r.db.Exec(sql_up, UserID)
				if err != nil {
					return nil, err
				}
				fmt.Println("iduser = ", iduser)

				sqlcoupon := `INSERT INTO coupon(doc_date,coupon_no,coupon_type,user_id,name,value,remain,expire_status,begin_date,expire_date,create_by,create_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
				ressql, err := r.db.Exec(sqlcoupon,
					now,       //doc_date
					doc_no_cp, //coupon_no
					1,         //coupon_type
					OrderUser, //user_id
					namePK,    //name
					pricePK,   //value
					pricePK,   //remain
					0,         //expire_status
					now,       //begin_date
					after,     //expire_date
					UserID,    //create_by
					now,       //create_time
				)
				if err != nil {
					return "", err
				}
				fmt.Println("ressql", ressql)

				sqlcoupon1 := `INSERT INTO coupon(doc_date,coupon_no,coupon_type,user_id,name,value,remain,expire_status,begin_date,expire_date,create_by,create_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
				ressql1, err := r.db.Exec(sqlcoupon1,
					now,              //doc_date
					doc_no_cp1,       //coupon_no
					2,                //coupon_type
					OrderUser,        //user_id
					namePK,           //name
					package_discount, //value
					package_discount, //remain
					0,                //expire_status
					now,              //begin_date
					after,            //expire_date
					UserID,           //create_by
					now,              //create_time
				)
				if err != nil {
					return "", err
				}
				fmt.Println("ressql", ressql1)
				break
			}
			fmt.Println("sub", sub)
		}
		fmt.Println("step")
	}

	sql := `update orders set order_status=?, edit_time=?, tracking_id=? where id=?`
	id, err := r.db.Exec(sql, Status, now, TrackingID, ID)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	fmt.Println("id = ", id)
	fmt.Println("UserIdUpdate = ", UserID)
	sql_track := `INSERT INTO track_order_status(order_id,doc_no,tracking_id,status_id,update_status_by,status_date) VALUES (?,?,?,?,?,?)`
	id_track, err := r.db.Exec(sql_track,
		item.ID,
		item.DocNo,
		TrackingID,
		Status,
		UserID,
		now,
	)
	fmt.Println("id_track", id_track)

	return "success", nil
}

func (r *orderRepo) DeliveryPriceRepo(UserID string, ItemID []int64, Code string) (name string, resp float64, err error) {
	var sumWeight float64
	// var sumPrice float64
	// var levelPrice string
	var MaxWeight float64
	fmt.Println("UserID:", UserID)

	delivery_max := `SELECT MAX(end_dtn) AS max_weight
	FROM cost_delivery WHERE code = ?`
	delivery_max1, err := r.db.QueryRow(delivery_max, Code)
	if err != nil {
		return "", resp, err
	}
	DeliveryCost := order.Delivery{}
	err = delivery_max1.Scan(&MaxWeight)

	for _, sub := range ItemID {
		// fmt.Println("sub:", sub)
		sql_item := `select a.weight,a.payment_destination,b.qty from basket_sub b left JOIN item a on a.id = b.item_id where a.id = ? and b.cust_id = ?`
		cal_rs, err := r.db.QueryRow(sql_item, sub, UserID)
		if err != nil {
			return "", resp, err
		}
		item := product.ItemDetail{}
		var qty float64
		err = cal_rs.Scan(&item.Weight, &item.PaymentDestination, &qty)
		sumWeight = sumWeight + item.Weight*qty
		// fmt.Println("Weight: ", item.Weight)
		// fmt.Println("qty: ", qty)
		// if item.PaymentDestination == 1 && DeliveryID == 5 {
		// 	return resp, fmt.Errorf("have item payment_destination = 1")
		// }
	}
	fmt.Println("sumWeight:", sumWeight)

	if sumWeight > MaxWeight {
		delivery_sql := `SELECT name, cost
		FROM cost_delivery
		WHERE code = ? ORDER BY id DESC LIMIT 1`
		delivery_sql1, err := r.db.QueryRow(delivery_sql, Code)
		if err != nil {
			return "", resp, err
		}
		err = delivery_sql1.Scan(&DeliveryCost.Name, &DeliveryCost.Cost)
		return DeliveryCost.Name, DeliveryCost.Cost, nil
	}
	if sumWeight <= MaxWeight {
		delivery_sql := `SELECT name, cost
		FROM cost_delivery
		WHERE ? between begin_dtn and end_dtn and code = ?`
		delivery_sql1, err := r.db.QueryRow(delivery_sql, sumWeight, Code)
		if err != nil {
			return "", resp, err
		}
		err = delivery_sql1.Scan(&DeliveryCost.Name, &DeliveryCost.Cost)
	}
	if sumWeight == 0 {
		return DeliveryCost.Name, 0.00, nil
	}
	// fmt.Println("Name:", DeliveryCost.Name)
	// fmt.Println("Cost:", DeliveryCost.Cost)
	return DeliveryCost.Name, DeliveryCost.Cost, nil
}

func (r *orderRepo) CallTokenLine(ID int64) (token string, err error) {
	fmt.Println("repo", ID)
	rs, err := r.db.QueryRow(`SELECT line_token from line_group where id = ?`, ID)
	if err != nil {
		fmt.Println(err)
		return token, err
	}
	rs.Scan(&token)
	fmt.Println("token:", token)
	return token, nil
}

func (r *orderRepo) DeliveryAllRepo() ([]order.DeliveryAll, error) {
	sql := `select DISTINCT name,code from cost_delivery`
	items := []order.DeliveryAll{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	// fmt.Println(clientID)
	for rs.Next() {
		item := order.DeliveryAll{}

		err = rs.Scan(&item.Name, &item.Code)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *orderRepo) OrderListAllRepo(OrderStatus int64) ([]order.Order, error) {
	var sql string

	if OrderStatus == 1 {
		fmt.Println("Case1 =", OrderStatus)
		sql = `select a.id, a.user_id, a.user_name, a.sale_id, a.sale_name, a.doc_no, a.sum_of_item_amount, a.discount_amount, a.after_discount_amount, 
	a.before_tax, a.tax_amount, a.total_amount, a.sum_cash_amount, a.sum_credit_amount, a.sum_deposit_amount, a.sum_coupon_amount, 
	a.sum_bank_amount, a.change_amount, a.net_debt_amount, a.my_description, a.order_status,a.delivery_date, a.distance, a.distance_amount, 
	ifnull(a.delivery_link,'') as delivery_link , a.delivery_id, a.referral_id, ifnull(a.pic_slip,'') as pic_slip, a.create_time, ifnull(a.tracking_id,'') as tracking_id,
	a.is_package
	from orders a 
	where pic_slip !='' and order_status = 1 ORDER BY a.id DESC;`
	} else if OrderStatus != 0 {
		fmt.Println("Case2 =", OrderStatus)
		sql = `select a.id, a.user_id, a.user_name, a.sale_id, a.sale_name, a.doc_no, a.sum_of_item_amount, a.discount_amount, a.after_discount_amount, 
	a.before_tax, a.tax_amount, a.total_amount, a.sum_cash_amount, a.sum_credit_amount, a.sum_deposit_amount, a.sum_coupon_amount, 
	a.sum_bank_amount, a.change_amount, a.net_debt_amount, a.my_description, a.order_status,a.delivery_date, a.distance, a.distance_amount, 
	ifnull(a.delivery_link,'') as delivery_link , a.delivery_id, a.referral_id, ifnull(a.pic_slip,'') as pic_slip, a.create_time, ifnull(a.tracking_id,'') as tracking_id,
	a.is_package
	from orders a 
	where order_status = ` + fmt.Sprintln(OrderStatus) + "ORDER BY a.id DESC"
	} else if OrderStatus == 0 {
		fmt.Println("Case0 =", OrderStatus)
		sql = `select a.id, a.user_id, a.user_name, a.sale_id, a.sale_name, a.doc_no, a.sum_of_item_amount, a.discount_amount, a.after_discount_amount, 
	a.before_tax, a.tax_amount, a.total_amount, a.sum_cash_amount, a.sum_credit_amount, a.sum_deposit_amount, a.sum_coupon_amount, 
	a.sum_bank_amount, a.change_amount, a.net_debt_amount, a.my_description, a.order_status, a.delivery_date, a.distance, a.distance_amount, 
	ifnull(a.delivery_link,'') as delivery_link , a.delivery_id, a.referral_id, ifnull(a.pic_slip,'') as pic_slip, a.create_time, ifnull(a.tracking_id,'') as tracking_id,
	a.is_package
	from orders a ORDER BY a.id DESC;`
	}
	items := []order.Order{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		item := order.Order{}
		err = rs.Scan(&item.ID, &item.UserID, &item.UserName, &item.SaleID, &item.SaleName, &item.DocNo, &item.SumOfItemAmount, &item.DiscountAmount, &item.AfterDiscountAmount,
			&item.BeforeTax, &item.TaxAmount, &item.TotalAmount, &item.SumCashAmount, &item.SumCreditAmount, &item.SumDepositAmount, &item.SumCouponAmount,
			&item.SumBankAmount, &item.ChangeAmount, &item.NetDebtDmount, &item.MyDescription, &item.OrderStatus, &item.DeliveryDate, &item.Distance, &item.DistanceAmount,
			&item.DeliveryLink, &item.DeliveryID, &item.ReferralID, &item.PicSlip, &item.CreateTime, &item.TrackingID, &item.IsPackage,
		)
		if err != nil {
			return nil, err
		}
		// fmt.Println("itemID = ", item.UserID)
		sqlP := `select ifnull(a.telephone,'') as telephone
		from users a 
		where user_id = ?`
		rsP, err := r.db.QueryRow(sqlP, item.UserID)
		err = rsP.Scan(&item.Talephone)
		// if err != nil {
		// 	return nil, err
		// }

		if item.OrderStatus == 99 {
			var SubStatus string

			sql_sub_status := `select a.send_text from send_status a where a.id = ?`
			sql_sub_status1, err := r.db.QueryRow(sql_sub_status, item.OrderStatus)
			if err != nil {
				return nil, err
			}
			err = sql_sub_status1.Scan(&SubStatus)
			item.SendText = append(item.SendText, SubStatus)
			// fmt.Println("SubStatus:", SubStatus)
		} else {
			for i := 1; i <= int(item.OrderStatus); i++ {
				var SubStatus string

				sql_sub_status := `select a.send_text from send_status a where a.id = ?`
				sql_sub_status1, err := r.db.QueryRow(sql_sub_status, int(i))
				if err != nil {
					return nil, err
				}
				err = sql_sub_status1.Scan(&SubStatus)
				item.SendText = append(item.SendText, SubStatus)
				// fmt.Println("i:", i)
				// fmt.Println("SubStatus:", SubStatus)
			}
		}

		// for i := 1; i <= int(item.OrderStatus); i++ {
		// 	var SubStatus string

		// 	sql_sub_status := `select a.send_text from send_status a where a.id = ?`
		// 	sql_sub_status1, err := r.db.QueryRow(sql_sub_status, int(i))
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	err = sql_sub_status1.Scan(&SubStatus)
		// 	item.SendText = append(item.SendText, SubStatus)
		// 	// fmt.Println("i:", i)
		// 	// fmt.Println("SubStatus:", SubStatus)
		// }
		// fmt.Println("itemID = ", item.ID)
		sql1 := `select a.id ,a.order_id ,a.item_id ,a.item_name ,a.wh_id ,a.wh_code ,a.shelf_id ,a.shelf_code ,a.qty ,a.cn_qty ,a.unit_id ,a.unit_code ,
		a.price ,a.discount_amount ,a.item_amount ,a.average_cost ,a.sum_of_cost ,a.rate_1 ,a.stock_type ,a.basket_id ,a.item_sub_cat ,a.item_group ,
		a.type_code ,a.point ,a.is_cancel ,a.ref_line_number , pic_file_name_1, a.line_number
		from order_sub a
		where a.order_id = ?`
		rs1, err := r.db.Query(sql1, item.ID)
		if err != nil {
			return nil, err
		}
		for rs1.Next() {
			item_sub := order.OrderSub{}
			err = rs1.Scan(&item_sub.ID, &item_sub.OrderID, &item_sub.ItemID, &item_sub.ItemName, &item_sub.WhID, &item_sub.WhCode, &item_sub.ShelfID, &item_sub.ShelfCode,
				&item_sub.Qty, &item_sub.CnQty, &item_sub.UnitID, &item_sub.UnitCode, &item_sub.Price, &item_sub.DiscountAmount, &item_sub.ItemAmount,
				&item_sub.AverageCost, &item_sub.SumOfCost, &item_sub.Rate1, &item_sub.StockType, &item_sub.BasketID, &item_sub.ItemSubCat, &item_sub.ItemGroup,
				&item_sub.TypeCode, &item_sub.Point, &item_sub.IsCancel, &item_sub.RefLineNumber, &item_sub.PicFileName1, &item_sub.LineNumber,
			)
			if err != nil {
				return nil, err
			}
			item.OrderSub = append(item.OrderSub, item_sub)
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *orderRepo) OrderListSendAllRepo() ([]order.OrderSend, error) {
	sql := `select a.id, a.order_status, ifnull(a.delivery_link,'') as delivery_link , a.delivery_id, a.address_id, a.doc_no, ifnull(a.tracking_id,'') as tracking_id
	from orders a where order_status = 3`
	items := []order.OrderSend{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		item := order.OrderSend{}
		err = rs.Scan(&item.ID, &item.OrderStatus, &item.DeliveryLink, &item.DeliveryID, &item.AddressID, &item.DocNo, &item.TrackingID)
		if err != nil {
			return nil, err
		}

		sql1 := `select a.id ,a.order_id ,a.item_id ,a.item_name ,a.wh_id ,a.wh_code ,a.shelf_id ,a.shelf_code ,a.qty ,a.cn_qty ,a.unit_id ,a.unit_code ,
		a.price ,a.discount_amount ,a.item_amount ,a.average_cost ,a.sum_of_cost ,a.rate_1 ,a.stock_type ,a.basket_id ,a.item_sub_cat ,a.item_group ,
		a.type_code ,a.point ,a.is_cancel ,a.ref_line_number , pic_file_name_1, a.line_number
		from order_sub a
		where a.order_id = ?`
		rs1, err := r.db.Query(sql1, item.ID)
		if err != nil {
			return nil, err
		}
		for rs1.Next() {
			item_sub := order.OrderSub{}
			err = rs1.Scan(&item_sub.ID, &item_sub.OrderID, &item_sub.ItemID, &item_sub.ItemName, &item_sub.WhID, &item_sub.WhCode, &item_sub.ShelfID, &item_sub.ShelfCode,
				&item_sub.Qty, &item_sub.CnQty, &item_sub.UnitID, &item_sub.UnitCode, &item_sub.Price, &item_sub.DiscountAmount, &item_sub.ItemAmount,
				&item_sub.AverageCost, &item_sub.SumOfCost, &item_sub.Rate1, &item_sub.StockType, &item_sub.BasketID, &item_sub.ItemSubCat, &item_sub.ItemGroup,
				&item_sub.TypeCode, &item_sub.Point, &item_sub.IsCancel, &item_sub.RefLineNumber, &item_sub.PicFileName1, &item_sub.LineNumber,
			)
			if err != nil {
				return nil, err
			}
			item.OrderSub = append(item.OrderSub, item_sub)
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *orderRepo) OrderSendDetailRepo(ID int64) ([]order.OrderSendDetail, error) {
	fmt.Println("ID", ID)
	sql := `select a.id,a.company_id, a.user_id, a.user_name, a.sale_id, a.sale_name, a.doc_no, a.sum_of_item_amount, a.discount_amount, a.after_discount_amount, 
	a.before_tax, a.tax_amount, a.total_amount, a.sum_cash_amount, 
	a.sum_credit_amount, a.sum_deposit_amount, a.sum_coupon_amount, a.sum_bank_amount, a.change_amount, a.net_debt_amount, a.my_description, a.order_status,
	a.delivery_date, a.distance, a.distance_amount, ifnull(a.delivery_link,'') as delivery_link , a.delivery_id, 
	a.referral_id,  ifnull(a.pic_slip,'') as pic_slip, a.create_time, ifnull(a.tracking_id,'') as tracking_id, a.address_id, a.is_package
	from orders a
	where a.id = ?`
	items := []order.OrderSendDetail{}
	rs, err := r.db.QueryRow(sql, ID)
	if err != nil {
		return nil, err
	}
	item := order.OrderSendDetail{}
	err = rs.Scan(&item.ID, &item.CompanyID, &item.UserID, &item.UserName, &item.SaleID, &item.SaleName, &item.DocNo, &item.SumOfItemAmount, &item.DiscountAmount, &item.AfterDiscountAmount,
		&item.BeforeTax, &item.TaxAmount, &item.TotalAmount,
		&item.SumCashAmount, &item.SumCreditAmount, &item.SumDepositAmount, &item.SumCouponAmount, &item.SumBankAmount, &item.ChangeAmount,
		&item.NetDebtDmount, &item.MyDescription, &item.OrderStatus, &item.DeliveryDate,
		&item.Distance, &item.DistanceAmount, &item.DeliveryLink, &item.DeliveryID, &item.ReferralID, &item.PicSlip, &item.CreateTime, &item.TrackingID, &item.AddressID, &item.IsPackage,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println("OrderStatus =", item.OrderStatus)
	fmt.Println("DeliveryID =", item.DeliveryID)
	sqlDelivery := `select name FROM cost_delivery  where code = ? LIMIT 1`
	sqlDelivery1, err := r.db.QueryRow(sqlDelivery, item.DeliveryID)
	if err != nil {
		return nil, err
	}
	err = sqlDelivery1.Scan(&item.DeliveryName)
	////////////////////////SubStatus////////////////////////
	var SubStatus string
	sql_sub_status := `select a.send_text from send_status a where a.id = ?`
	sql_sub_status1, err := r.db.QueryRow(sql_sub_status, item.OrderStatus)
	if err != nil {
		return nil, err
	}
	err = sql_sub_status1.Scan(&SubStatus)
	item.SendText = append(item.SendText, SubStatus)

	////////////////////////CompanyID////////////////////////
	sql_company := `select id,ifnull(name,'') as name,ifnull(address,'') as address,ifnull(tex_id,'') as tex_id,
	ifnull(phone_mobile,'') as phone_mobile,ifnull(phone_home,'') as phone_home
	from define_company 
	where id = ?`
	rs_sql_company, err := r.db.QueryRow(sql_company, item.CompanyID)
	if err != nil {
		fmt.Println(err.Error())
	}
	company := order.OrderCompanyDetail{}
	err = rs_sql_company.Scan(&company.ID, &company.Name, &company.Address, &company.TexID, &company.PhoneMobile, &company.PhoneHome)
	if err != nil {
		return nil, err
	}
	item.OrderCompanyDetail = company

	////////////////////////AddressID////////////////////////
	if item.AddressID != 0 {
		sql_addr := `select ifnull(id,0) as addr_id,ifnull(name,'') as addr_name,ifnull(addr_phone,'') as addr_phone,ifnull(addr_email,'') as addr_email,
	ifnull(addr_state,'') as addr_state,ifnull(addr_subarea,'') as addr_subarea,ifnull(addr_district,'') as addr_district,ifnull(addr_province,'') as addr_province,
	ifnull(addr_postal_code,'') as addr_postal_code,ifnull(main_address,0) as main_address 
	from user_address 
	where id = ?`
		rs_sql_address, err := r.db.QueryRow(sql_addr, item.AddressID)
		if err != nil {
			fmt.Println(err.Error())
		}
		addr := order.OrderSendAddress{}
		err = rs_sql_address.Scan(&addr.AddressID, &addr.Name, &addr.Phone, &addr.Email, &addr.Address, &addr.SubArea, &addr.District, &addr.Province, &addr.PostalCode, &addr.MainAddress)
		if err != nil {
			return nil, err
		}
		item.OrderSendAddress = addr
	} else {

	}

	////////////////////////OrderSub////////////////////////
	sql1 := `select a.id ,a.order_id ,a.item_id ,a.item_name ,a.wh_id ,a.wh_code ,a.shelf_id ,a.shelf_code ,a.qty ,a.cn_qty ,a.unit_id ,a.unit_code ,
		a.price ,a.discount_amount ,a.item_amount ,a.average_cost ,a.sum_of_cost ,a.rate_1 ,a.stock_type ,a.basket_id ,a.item_sub_cat ,a.item_group ,
		a.type_code ,a.point ,a.is_cancel ,a.ref_line_number , pic_file_name_1, a.line_number
		from order_sub a
		where a.order_id = ?`
	rs1, err := r.db.Query(sql1, item.ID)
	if err != nil {
		return nil, err
	}
	for rs1.Next() {
		item_sub := order.OrderSub{}
		err = rs1.Scan(&item_sub.ID, &item_sub.OrderID, &item_sub.ItemID, &item_sub.ItemName, &item_sub.WhID, &item_sub.WhCode, &item_sub.ShelfID, &item_sub.ShelfCode,
			&item_sub.Qty, &item_sub.CnQty, &item_sub.UnitID, &item_sub.UnitCode, &item_sub.Price, &item_sub.DiscountAmount, &item_sub.ItemAmount,
			&item_sub.AverageCost, &item_sub.SumOfCost, &item_sub.Rate1, &item_sub.StockType, &item_sub.BasketID, &item_sub.ItemSubCat, &item_sub.ItemGroup,
			&item_sub.TypeCode, &item_sub.Point, &item_sub.IsCancel, &item_sub.RefLineNumber, &item_sub.PicFileName1, &item_sub.LineNumber,
		)
		if err != nil {
			return nil, err
		}
		item.OrderSub = append(item.OrderSub, item_sub)
	}
	items = append(items, item)

	return items, nil
}

func (r *orderRepo) OrderPackageRepo(UserID string, ItemID int64) (interface{}, error) {
	// var check_exist_id int64
	var GetID int64
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now_time := time.Now().In(loc)
	doc_no, err := GenDoc()
	doc_no = "ORD" + doc_no
	fmt.Println("UserID :", UserID)
	//////////////////////////////////////////////////////////////////////getuser/////////////////////////////////////////////////////////////////////////
	sqluser := `select user_fname,user_lname
	from users
	where user_id = ?`
	sqluser1, err := r.db.QueryRow(sqluser, UserID)
	if err != nil {
		return nil, err
	}
	UseDetail := order.UseOrder{}
	sqluser1.Scan(&UseDetail.UserFname, &UseDetail.UserLname)
	if err != nil {
		return nil, err
	}
	Username := UseDetail.UserFname + UseDetail.UserLname
	fmt.Println("Username :", Username)
	///////////////////////////////////////////////////////////////////////getitem//////////////////////////////////////////////////////////////////////////
	sqlitem := `select a.id,a.code,a.name,a.eng_name,a.short_name,a.category_code,a.type_code,a.my_description,a.price,a.my_grade,a.pic_file_name_1,a.unit_id,
	a.unit_code,a.def_sale_wh_id,a.def_sale_wh_code,a.degree,b.name,b.eng_name
	from item a left JOIN item_category b on a.category_code = b.code
	where a.id = ?`
	rs, err := r.db.QueryRow(sqlitem, ItemID)
	if err != nil {
		return nil, err
	}
	item := product.Item{}
	rs.Scan(&item.ID, &item.Code, &item.Name, &item.EngName, &item.ShortName, &item.CategoryCode, &item.TypeCode,
		&item.MyDescription, &item.Price, &item.MyGrade, &item.PicFileName1, &item.UnitID, &item.UnitCode, &item.DefSaleWhID, &item.DefSaleWhCode,
		&item.Degree, &item.CategoryName, &item.CategoryNameEng,
	)
	if err != nil {
		return nil, err
	}

	//////////////////////////////////////////////////////////////////////
	sql := `INSERT INTO orders(user_id,user_name,sale_id,sale_name,doc_no,doc_date,sum_of_item_amount,discount_amount,after_discount_amount,before_tax,
		tax_amount,total_amount,sum_cash_amount,sum_credit_amount,sum_deposit_amount,sum_coupon_amount,sum_bank_amount,change_amount,
		net_debt_amount,my_description,delivery_date,distance,distance_amount,delivery_id,referral_id,create_time,is_package,order_status)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1)`
	sql1, err := r.db.Exec(sql,
		UserID,     // user_id
		Username,   // user_name
		0,          // sale_id
		"",         // sale_name
		doc_no,     // doc_no
		now_time,   // doc_date
		item.Price, // SumOfItemAmount,
		0,          // DiscountAmount,
		item.Price, // AfterDiscountAmount,
		0,          // BeforeTax,
		0,          // TaxAmount,
		item.Price, // TotalAmount,
		0,          // SumCashAmount,
		0,          // SumCreditAmount,
		0,          // SumDepositAmount,
		0,          // SumCouponAmount,
		0,          // SumBankAmount,
		0,          // ChangeAmount,
		0,          // NetDebtDmount,
		"",         // MyDescription,
		now_time,   // DeliveryDate,
		0,          // Distance,
		0,          // DistanceAmount,
		0,          // DeliveryID,
		0,          // ReferralID,
		now_time,   // create_time
		1,          // is_package
	)
	log.Printf("sql1:", sql1)
	if err != nil {
		return nil, err
	}
	resqID, _ := sql1.LastInsertId()
	GetID = resqID

	sqlsub1 := `INSERT INTO order_sub(order_id,item_id,item_name,wh_id,wh_code,shelf_id,shelf_code,qty,cn_qty,unit_id,unit_code,price,discount_amount,item_amount,
		average_cost,sum_of_cost,rate_1,stock_type,basket_id,item_sub_cat,item_group,type_code,point,is_cancel,ref_line_number,pic_file_name_1,line_number)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	sqlsub, err := r.db.Exec(sqlsub1,
		GetID,
		item.ID,
		item.Name,
		item.DefSaleWhID,
		item.DefSaleWhCode,
		0,  //busket_sub.ShelfID,
		"", //busket_sub.ShelfCode,
		1,
		0, //busket_sub.CnQty,
		item.UnitID,
		item.UnitCode,
		item.Price,
		0,
		0,
		0, //busket_sub.AverageCost,
		0, //busket_sub.SumOfCost,
		0, //busket_sub.Rate1,
		0, //busket_sub.StockType,
		0,
		0, //busket_sub.Itembusket_subCat,
		0, //busket_sub.ItemGroup,
		0, //busket_sub.TypeCode,
		0, //busket_sub.Point,
		0, //busket_sub.IsCancel,
		0, //busket_sub.RefLineNumber,
		item.PicFileName1,
		1,
	)
	log.Printf("sqlbksub:", sqlsub)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id": GetID,
	}, nil
}

func (r *orderRepo) StatusOrder(DocNo string) (interface{}, error) {
	fmt.Println("Track", DocNo)
	sql := `select a.order_status, ifnull(a.tracking_id,'') as tracking_id, ifnull(a.delivery_id,'') as delivery_id from orders a where a.doc_no = ?`
	rs, err := r.db.QueryRow(sql, DocNo)
	if err != nil {
		return nil, err
	}
	item := order.OrderTrack{}
	err = rs.Scan(&item.OrderStatus, &item.TrackingID, &item.DeliveryID)
	if err != nil {
		return nil, err
	}
	if item.OrderStatus == 99 {
		var SubStatus string
		sql_sub_status := `select a.send_text from send_status a where a.id = ?`
		sql_sub_status1, err := r.db.QueryRow(sql_sub_status, item.OrderStatus)
		if err != nil {
			return nil, err
		}
		err = sql_sub_status1.Scan(&SubStatus)
		item.SendText = append(item.SendText, SubStatus)
		fmt.Println("SubStatus:", SubStatus)
	} else {
		for i := 1; i <= int(item.OrderStatus); i++ {
			var SubStatus string
			sql_sub_status := `select a.send_text from send_status a where a.id = ?`
			sql_sub_status1, err := r.db.QueryRow(sql_sub_status, int(i))
			if err != nil {
				return nil, err
			}
			err = sql_sub_status1.Scan(&SubStatus)
			if SubStatus == "ทำการจัดส่ง" {
				sqlDelivery := `select name FROM cost_delivery  where code = ? LIMIT 1`
				sqlDelivery1, err := r.db.QueryRow(sqlDelivery, item.DeliveryID)
				if err != nil {
					return nil, err
				}
				err = sqlDelivery1.Scan(&item.DeliveryName)
				switch item.DeliveryID {
				case "01":
					SubStatus = "กำลังขนส่งโดย " + item.DeliveryName + " ติดตามได้ที่ " + " https://www.flashexpress.co.th/tracking/"
					break
				case "02":
					SubStatus = "กำลังขนส่งโดย " + item.DeliveryName + " ติดตามได้ที่ " + "https://www.nimexpress.com/web/p/tracking?i=" + item.TrackingID
					break
				case "03":
					SubStatus = "กำลังขนส่งโดย " + item.DeliveryName + " ติดตามได้ที่ " + "https://th.kerryexpress.com/en/track/?track" + item.TrackingID
					break
				case "04":
					SubStatus = "กำลังขนส่งโดย " + item.DeliveryName + " ติดตามได้ที่ " + "https://track.thailandpost.co.th/?trackNumber=" + item.TrackingID
					break
				}
			}
			item.SendText = append(item.SendText, SubStatus)
			fmt.Println("i:", i)
			fmt.Println("SubStatus:", SubStatus)
		}
	}

	return item, nil
}
