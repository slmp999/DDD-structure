package mysql

import (
	"fmt"
	"strconv"
	"time"

	"gitlab.com/satit13/perfect_api/product"
)

func (r *productRepo) GenCodeItem(item_type string) (string, error) {
	fmt.Println("step1")
	code, err := r.FindCodeItem(item_type)
	if err != nil {
		return "", err
	}
	fmt.Println("step2")
	Newcode, err := GenItemID(code)
	return item_type + Newcode, nil
}

func (r *productRepo) FindCodeItem(item_type string) (string, error) {
	var code string
	sqluser := `SELECT code FROM item WHERE code LIKE "%` + item_type + `%" ORDER BY id DESC LIMIT 1`
	sqluser1, err := r.db.QueryRow(sqluser)
	fmt.Println(sqluser)
	if err != nil {
		return "", err
	}
	sqluser1.Scan(&code)
	if err != nil {
		return "", err
	}
	itemProduct, err := GenItemID(code)
	return itemProduct, err
}

func GenItemID(code string) (string, error) {
	var codeID string
	id, _ := strconv.Atoi(code[len(code)-5:])
	codenumber := (strconv.Itoa(id + 1))
	fmt.Println(codenumber)
	if len(string(codenumber)) == 1 {
		codeID = "0000" + string(codenumber)
	}
	if len(string(codenumber)) == 2 {
		codeID = "000" + string(codenumber)
	}
	if len(string(codenumber)) == 3 {
		codeID = "00" + string(codenumber)
	}
	if len(string(codenumber)) == 4 {
		codeID = "0" + string(codenumber)
	}
	if len(string(codenumber)) == 5 {
		codeID = string(codenumber)
	}
	return codeID, nil
}

func (r *productRepo) AddCategoryRepo(UserID string, req product.CategoryAdd) (interface{}, error) {
	var check_exist_id int64
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now_time := time.Now().In(loc)

	fmt.Println("ID = ", req.ID)
	var UserFname string
	var UserLname string
	sqluser := `select user_fname,user_lname from users where user_id = ?`
	sqluser1, err := r.db.QueryRow(sqluser, UserID)
	if err != nil {
		return nil, err
	}
	sqluser1.Scan(&UserFname, &UserLname)
	if err != nil {
		return nil, err
	}
	Username := UserFname + " " + UserLname
	fmt.Println("Username :", Username)

	sqlexist := `select count(id) as check_exist from item_category where id = ?`
	sqlexist1, err := r.db.QueryRow(sqlexist, req.ID)
	if err != nil {
		return nil, err
	}
	sqlexist1.Scan(&check_exist_id)
	fmt.Println("check_exist_id = ", check_exist_id)
	if check_exist_id == 0 {
		var getcode string
		sql_code := `select code from item_category ORDER BY id DESC LIMIT 1`
		sql_code1, err := r.db.QueryRow(sql_code)
		if err != nil {
			return nil, err
		}
		sql_code1.Scan(&getcode)
		fmt.Println("getcode:", getcode)
		code_temp, err := strconv.Atoi(getcode)
		code_temp = code_temp + 1
		fmt.Println("code_temp:", code_temp)
		code := strconv.Itoa(code_temp)
		if len(code) == 1 {
			code = "0" + code
		}

		fmt.Println("code:", code)
		sql := `INSERT INTO item_category(company_id,code,name,eng_name,item_count,my_description,active_status,create_by,create_time) VALUES (?,?,?,?,?,?,?,?,?)`
		_, err = r.db.Exec(sql,
			req.CompanyID,
			code,
			req.Name,
			req.EngName,
			1,
			req.MyDescription,
			req.ActiveStatus,
			Username,
			now_time,
		)
		if err != nil {
			return "", err
		}
	} else {
		sql := `Update item_category set company_id=?,name=?,eng_name=?,item_count=?,my_description=?,active_status=?,edit_by=?,edit_time=? where id=?`
		fmt.Println("sql update = ", sql)
		_, err := r.db.Exec(sql,
			req.CompanyID,
			req.Name,
			req.EngName,
			1,
			req.MyDescription,
			req.ActiveStatus,
			Username,
			now_time,
			req.ID,
		)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}
	}
	return "success", nil
}

func (r *productRepo) AddItemRepo(UserID string, req product.ItemAdd) (interface{}, error) {
	var check_exist_id int64
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now_time := time.Now().In(loc)
	NewCode, _ := r.GenCodeItem("PCT")

	sqlexist := `select count(id) as check_exist from item where id = ?`
	sqlexist1, err := r.db.QueryRow(sqlexist, req.ID)
	if err != nil {
		return nil, err
	}
	sqlexist1.Scan(&check_exist_id)

	if check_exist_id == 0 {
		fmt.Println("No ID")
		sql := `INSERT INTO item(code,name,eng_name,short_name,group_code,type_code,my_description,stock_type,unit_id,unit_code,
			price,buy_unit_id,buy_unit_code,sale_unit_id,sale_unit_code,weight,hight,wide,category_code,sub_cat_code,
			my_grade,brand_code,color_code,my_class,unit_type,delivery_charge,install_charge,return_status,item_status,last_price,
			average_cost,cost_type,pic_file_name_1,favorite_promo,degree,favorite,
			package_discount,coupon_amount,is_promotion,is_saler,pic_file_saler,active_status,create_by,create_time) 
			VALUES (?,?,?,?,?,?,?,?,?,?,
				?,?,?,?,?,?,?,?,?,?,
				?,?,?,?,?,?,?,?,?,?,
				?,?,?,?,?,?,
				?,?,?,?,?,?,?,?)`
		id, err := r.db.Exec(sql,
			NewCode,
			req.Name,
			req.EngName,
			req.ShortName,
			"", //req.GroupCode,
			"", //req.TypeCode,
			req.MyDescription,
			0, //req.StockType,
			req.UnitID,
			req.UnitCode,
			req.Price,
			0,            //req.BuyUnitID,
			"",           //req.BuyUnitCode,
			req.UnitID,   //req.SaleUnitID,
			req.UnitCode, //req.SaleUnitCode,
			req.Weight,
			0, //req.Hight,
			0, //req.Wide,
			req.CategoryCode,
			"00", //req.SubCatCode,
			"",   //req.MyGrade,
			"",   //req.BrandCode,
			"",   //req.ColorCode,
			"",   //req.MyClass,
			0,    //req.UnitType,
			0,    //req.DeliveryCharge,
			0,    //req.InstallCharge,
			0,    //req.ReturnStatus,
			0,    //req.ItemStatus,
			req.LastPrice,
			0, //req.AverageCost,
			0, //req.CostType,
			req.PicFileName1,
			0, //req.FavoritePromo,
			req.Degree,
			req.Favorite,
			0, //req.PackageDiscount,
			0, //req.CouponAmount,
			0, //req.IsPromotion,
			req.IsSaler,
			req.PicFileSaler,
			req.ActiveStatus,
			UserID,
			now_time,
		)
		if err != nil {
			return "", err
		}
		var GetID int64
		resqID, _ := id.LastInsertId()
		GetID = resqID

		for _, SubPic := range req.PicSub {
			sql_sub := `INSERT INTO pic_item_sub(item_id,url) VALUES (?,?)`
			_, err := r.db.Exec(sql_sub,
				GetID,
				SubPic.URL,
			)
			if err != nil {
				return "", err
			}
		}

		sql3 := `INSERT INTO item_price(item_id,item_code,sale_price_1,sale_price_2,unit_id,unit_code,per_discount,begin_date,end_date,create_by,create_time,active_status) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
		id_sub, err := r.db.Exec(sql3,
			GetID,
			NewCode,
			req.SalePrice1,
			req.SalePrice2,
			req.UnitID,
			req.UnitCode,
			req.PerDiscount,
			now_time, // req.BeginDate,
			now_time, // req.EndDate,
			UserID,
			now_time,
			req.ActiveStatus,
		)
		if err != nil {
			return "", err
		}
		fmt.Println("id_sub =", id_sub)
	} else {
		fmt.Println("have ID")
		sql := `Update item set name=?,eng_name=?,short_name=?,my_description=?,unit_id=?,unit_code=?,price=?,sale_unit_id=?,sale_unit_code=?,
		weight=?,category_code=?,last_price=?,pic_file_name_1=?,
		degree=?,favorite=?,is_saler=?,pic_file_saler=?,active_status=?,edit_time=?,edit_by=?
		where id=?`
		// fmt.Println("sql update = ", sql)
		_, err := r.db.Exec(sql,
			req.Name,
			req.EngName,
			req.ShortName,
			req.MyDescription,
			req.UnitID,
			req.UnitCode,
			req.Price,
			req.UnitID,   //req.SaleUnitID,
			req.UnitCode, //req.SaleUnitCode,
			req.Weight,
			req.CategoryCode,
			req.LastPrice,
			req.PicFileName1,
			req.Degree,
			req.Favorite,
			req.IsSaler,
			req.PicFileSaler,
			req.ActiveStatus,
			now_time,
			UserID,
			req.ID,
		)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}

		for _, SubPic := range req.PicSub {
			if SubPic.ID == 0 {
				sql_sub := `INSERT INTO pic_item_sub(item_id,url) VALUES (?,?)`
				_, err := r.db.Exec(sql_sub,
					req.ID,
					SubPic.URL,
				)
				if err != nil {
					return "", err
				}
			} else {
				sql_sub := `Update pic_item_sub set url=? where id=?`
				_, err := r.db.Exec(sql_sub,
					SubPic.URL,
					SubPic.ID,
				)
				if err != nil {
					return "", err
				}
			}
		}

		sql3 := `Update item_price set sale_price_1=?,sale_price_2=?,unit_id=?,unit_code=?,per_discount=?,begin_date=?,end_date=?,edit_time=?,edit_by=? where item_id=?`
		id_sub, err := r.db.Exec(sql3,
			req.SalePrice1,
			req.SalePrice2,
			req.UnitID,
			req.UnitCode,
			req.PerDiscount,
			req.BeginDate,
			req.EndDate,
			now_time,
			UserID,
			req.ID,
		)
		if err != nil {
			return "", err
		}
		fmt.Println("id_sub:", id_sub)
	}
	return nil, nil
}

func (r *productRepo) AddPromotion(UserID string, req product.ItemAdd) (interface{}, error) {
	var check_exist_id int64
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now_time := time.Now().In(loc)
	NewCode, _ := r.GenCodeItem("PMT")
	sqlexist := `select count(id) as check_exist from item where id = ?`
	sqlexist1, err := r.db.QueryRow(sqlexist, req.ID)
	if err != nil {
		return nil, err
	}
	sqlexist1.Scan(&check_exist_id)

	if check_exist_id == 0 {
		sql := `INSERT INTO item(code,name,eng_name,short_name,group_code,type_code,my_description,stock_type,unit_id,unit_code,
			price,buy_unit_id,buy_unit_code,sale_unit_id,sale_unit_code,weight,hight,wide,category_code,sub_cat_code,
			my_grade,brand_code,color_code,my_class,unit_type,delivery_charge,install_charge,return_status,item_status,last_price,
			average_cost,cost_type,pic_file_name_1,favorite_promo,degree,favorite,
			package_discount,coupon_amount,is_promotion,create_time,edit_by,active_status) 
			VALUES (?,?,?,?,?,?,?,?,?,?,
				?,?,?,?,?,?,?,?,?,?,
				?,?,?,?,?,?,
				?,?,?,?,?,?,?,?,?,?,
				?,?,?,?,?,?)`
		id, err := r.db.Exec(sql,
			NewCode,
			req.Name,
			req.EngName,
			req.ShortName,
			req.GroupCode,
			"", //req.TypeCode,
			req.MyDescription,
			0, //req.StockType,
			req.UnitID,
			req.UnitCode,
			req.Price,
			0,            //req.BuyUnitID,
			"",           //req.BuyUnitCode,
			req.UnitID,   //req.SaleUnitID,
			req.UnitCode, //req.SaleUnitCode,
			req.Weight,
			0,         //req.Hight,
			0,         //req.Wide,
			"05",      //req.CategoryCode,
			"00",      //req.SubCatCode,
			"",        //req.MyGrade,
			"",        //req.BrandCode,
			"",        //req.ColorCode,
			"",        //req.MyClass,
			0,         //req.UnitType,
			0,         //req.DeliveryCharge,
			0,         //req.InstallCharge,
			0,         //req.ReturnStatus,
			0,         //req.ItemStatus,
			req.Price, //req.LastPrice,
			0,         //req.AverageCost,
			0,         //req.CostType,
			req.PicFileName1,
			req.FavoritePromo,
			req.Degree,
			0, //req.Favorite,
			0, //req.PackageDiscount,
			0, //req.CouponAmount,
			1, //req.IsPromotion,
			now_time,
			UserID,
			req.ActiveStatus,
		)
		if err != nil {
			return "", err
		}
		var GetID int64
		resqID, _ := id.LastInsertId()
		GetID = resqID
		fmt.Println("GetID:", GetID)
		for _, SubPic := range req.PicSub {
			sql_sub := `INSERT INTO pic_item_sub(item_id,url) VALUES (?,?)`
			_, err := r.db.Exec(sql_sub,
				GetID,
				SubPic.URL,
			)
			if err != nil {
				return "", err
			}
		}

		sql3 := `INSERT INTO item_price(item_id,item_code,sale_price_1,sale_price_2,unit_id,unit_code,begin_date,end_date,create_time,active_status) VALUES (?,?,?,?,?,?,?,?,?,?)`
		id_sub, err := r.db.Exec(sql3,
			GetID,
			NewCode,
			req.SalePrice1,
			req.SalePrice2,
			req.UnitID,
			req.UnitCode,
			req.BeginDate,
			req.EndDate,
			now_time,
			req.ActiveStatus,
		)
		if err != nil {
			return "", err
		}
		fmt.Println("id_sub:", id_sub)

		for _, sub := range req.ItemPromoSub {
			sql_sub := `INSERT INTO item_components(parent_code,item_id,item_code,qty,price,amount,unit_id,unit_code,line_number) 
			VALUES (?,?,?,?,?,?,?,?,?)`
			_, err := r.db.Exec(sql_sub,
				NewCode,
				sub.ItemID,
				sub.ItemCode,
				sub.Qty,
				sub.Price,
				sub.Amount,
				sub.UnitID,
				sub.UnitCode,
				sub.LineNumber,
			)
			if err != nil {
				return "", err
			}
		}
	} else {
		fmt.Println("have ID")
		sql := `Update item set name=?,eng_name=?,short_name=?,group_code=?,my_description=?,unit_id=?,unit_code=?,price=?,sale_unit_id=?,sale_unit_code=?,
		weight=?,category_code=?,last_price=?,pic_file_name_1=?,degree=?,favorite=?,active_status=?
		where id=?`
		fmt.Println("sql update = ", sql)
		_, err := r.db.Exec(sql,
			//req.Code,
			req.Name,
			req.EngName,
			req.ShortName,
			req.GroupCode,
			req.MyDescription,
			req.UnitID,
			req.UnitCode,
			req.Price,
			req.UnitID,   //req.SaleUnitID,
			req.UnitCode, //req.SaleUnitCode,
			req.Weight,
			req.CategoryCode,
			req.Price, //req.LastPrice,
			req.PicFileName1,
			req.Degree,
			req.Favorite,
			req.ActiveStatus,
			req.ID,
		)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}

		for _, SubPic := range req.PicSub {
			if SubPic.ID == 0 {
				sql_sub := `INSERT INTO pic_item_sub(item_id,url) VALUES (?,?)`
				_, err := r.db.Exec(sql_sub,
					req.ID,
					SubPic.URL,
				)
				if err != nil {
					return "", err
				}
			} else {
				sql_sub := `Update pic_item_sub set url=? where id=?`
				_, err := r.db.Exec(sql_sub,
					SubPic.URL,
					SubPic.ID,
				)
				if err != nil {
					return "", err
				}
			}
		}

		sql3 := `Update item_price set sale_price_1=?,sale_price_2=?,unit_id=?,unit_code=?,begin_date=?,end_date=?,edit_time=? where item_id=?`
		id_sub, err := r.db.Exec(sql3,
			req.SalePrice1,
			req.SalePrice2,
			req.UnitID,
			req.UnitCode,
			req.BeginDate,
			req.EndDate,
			now_time,
			req.ID,
		)
		if err != nil {
			return "", err
		}
		fmt.Println("id_sub:", id_sub)

		for _, sub := range req.ItemPromoSub {
			if sub.ID == 0 {
				sql_sub := `INSERT INTO item_components(parent_code,item_id,item_code,name,qty,price,amount,unit_id,unit_code,line_number) 
				VALUES (?,?,?,?,?,?,?,?,?,?)`
				_, err := r.db.Exec(sql_sub,
					req.Code,
					sub.ItemID,
					sub.ItemCode,
					sub.Name,
					sub.Qty,
					sub.Price,
					sub.Price,
					sub.UnitID,
					sub.UnitCode,
					sub.LineNumber,
				)
				if err != nil {
					return "", err
				}
			} else {
				sql_sub := `Update item_components set parent_code=?,item_id=?,item_code=?,name=?,qty=?,price=?,amount=?,unit_id=?,unit_code=?,line_number=? where id=?`
				_, err := r.db.Exec(sql_sub,
					req.Code,
					sub.ItemID,
					sub.ItemCode,
					sub.Name,
					sub.Qty,
					sub.Price,
					sub.Price,
					sub.UnitID,
					sub.UnitCode,
					sub.LineNumber,
					sub.ID,
				)
				if err != nil {
					return "", err
				}
			}
		}
	}
	return nil, nil
}

func (r *productRepo) AddPackage(UserID string, req product.ItemAdd) (interface{}, error) {
	var check_exist_id int64
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now_time := time.Now().In(loc)
	NewCode, _ := r.GenCodeItem("PKS")
	sqlexist := `select count(id) as check_exist from item where id = ?`
	sqlexist1, err := r.db.QueryRow(sqlexist, req.ID)
	if err != nil {
		return nil, err
	}
	sqlexist1.Scan(&check_exist_id)

	if check_exist_id == 0 {
		fmt.Println("No ID")
		sql := `INSERT INTO item(code,name,eng_name,short_name,group_code,type_code,my_description,stock_type,unit_id,unit_code,
			price,buy_unit_id,buy_unit_code,sale_unit_id,sale_unit_code,weight,hight,wide,category_code,sub_cat_code,
			my_grade,brand_code,color_code,my_class,unit_type,delivery_charge,install_charge,return_status,item_status,last_price,
			average_cost,cost_type,pic_file_name_1,favorite_promo,degree,favorite,
			package_discount,coupon_amount,is_package,is_promotion,active_status,create_by,create_time) 
			VALUES (?,?,?,?,?,?,?,?,?,?,
				?,?,?,?,?,?,?,?,?,?,
				?,?,?,?,?,?,?,?,?,?,
				?,?,?,?,?,?,
				?,?,?,?,?,?,?)`
		id, err := r.db.Exec(sql,
			NewCode,
			req.Name,
			req.EngName,
			req.ShortName,
			"", //req.GroupCode,
			"", //req.TypeCode,
			req.MyDescription,
			9,         //req.StockType,
			4,         //req.UnitID,
			"package", //req.UnitCode,
			req.Price,
			0,         //req.BuyUnitID,
			"",        //req.BuyUnitCode,
			4,         //req.SaleUnitID,
			"package", //req.SaleUnitCode,
			0,         //req.Weight,
			0,         //req.Hight,
			0,         //req.Wide,
			"04",      //req.CategoryCode,
			"051",     //req.SubCatCode,
			"",        //req.MyGrade,
			"",        //req.BrandCode,
			"",        //req.ColorCode,
			"",        //req.MyClass,
			0,         //req.UnitType,
			0,         //req.DeliveryCharge,
			0,         //req.InstallCharge,
			0,         //req.ReturnStatus,
			0,         //req.ItemStatus,
			req.LastPrice,
			0, //req.AverageCost,
			0, //req.CostType,
			req.PicFileName1,
			0, //req.FavoritePromo,
			5, //req.Degree,
			0, //req.Favorite,
			req.PackageDiscount,
			req.CouponAmount,
			1, //is_package
			0, //req.IsPromotion,
			req.ActiveStatus,
			UserID,
			now_time,
		)
		if err != nil {
			return "", err
		}
		var GetID int64
		resqID, _ := id.LastInsertId()
		GetID = resqID
		fmt.Println("id_sub:", id)

		for _, SubPic := range req.PicSub {
			sql_sub := `INSERT INTO pic_item_sub(item_id,url) VALUES (?,?)`
			_, err := r.db.Exec(sql_sub,
				GetID,
				SubPic.URL,
			)
			if err != nil {
				return "", err
			}
		}

		sql3 := `INSERT INTO item_price(item_id,item_code,sale_price_1,sale_price_2,unit_id,unit_code,create_by,create_time,active_status) VALUES (?,?,?,?,?,?,?,?,?)`
		id_sub, err := r.db.Exec(sql3,
			GetID,
			NewCode,
			req.Price,
			req.Price,
			4,
			"package",
			UserID,
			now_time,
			req.ActiveStatus,
		)
		if err != nil {
			return "", err
		}
		fmt.Println("id_sub:", id_sub)
	} else {
		fmt.Println("have ID")
		sql := `Update item set name=?,eng_name=?,short_name=?,my_description=?,price=?,last_price=?,pic_file_name_1=?,
		package_discount=?,coupon_amount=?,active_status=?,edit_by=?,edit_time=?
		where id=?`
		fmt.Println("sql update = ", sql)
		_, err := r.db.Exec(sql,
			req.Name,
			req.EngName,
			req.ShortName,
			req.MyDescription,
			req.Price,
			req.LastPrice,
			req.PicFileName1,
			req.PackageDiscount,
			req.CouponAmount,
			req.ActiveStatus,
			UserID,
			now_time,
			req.ID,
		)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}

		for _, SubPic := range req.PicSub {
			if SubPic.ID == 0 {
				sql_sub := `INSERT INTO pic_item_sub(item_id,url) VALUES (?,?)`
				_, err := r.db.Exec(sql_sub,
					req.ID,
					SubPic.URL,
				)
				if err != nil {
					return "", err
				}
			} else {
				sql_sub := `Update pic_item_sub set url=? where id=?`
				_, err := r.db.Exec(sql_sub,
					SubPic.URL,
					SubPic.ID,
				)
				if err != nil {
					return "", err
				}
			}
		}

		sql3 := `Update item_price set sale_price_1=?,sale_price_2=?,edit_by=?,edit_time=? where item_id=?`
		id_sub, err := r.db.Exec(sql3,
			req.Price,
			req.Price,
			UserID,
			now_time,
			req.ID,
		)
		if err != nil {
			return "", err
		}
		fmt.Println("id_sub:", id_sub)
	}
	return nil, nil
}

func (r *productRepo) DeletePicture(ID int64, UserID string) (interface{}, error) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now_time := time.Now().In(loc)

	sql := `update pic_item_sub set active_status=0, delete_by=?, delete_time=? where id=?`
	_, err := r.db.Exec(sql, UserID, now_time, ID)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	return "success", nil
}

func (r *productRepo) DeleteProSub(ID int64, UserID string) (interface{}, error) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now_time := time.Now().In(loc)

	sql := `update item_components set active_status=0, delete_by=?, delete_time=? where id=?`
	_, err := r.db.Exec(sql, UserID, now_time, ID)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	return "success", nil
}
