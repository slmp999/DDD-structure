package mysql

import (
	"database/sql"

	"gitlab.com/satit13/perfect_api/product"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"

	//"github.com/ponzu-cms/ponzu/system/item"
	"fmt"
)

type status int

const (
	statusActive                      status = 0
	statusInactive                    status = 1
	statusDelete                      status = 9
	errVendingRoleDenied                     = "vending role & permission denied!"
	errDataLength                            = "data length error!"
	errDataNotEqualWithMasterCategory        = "your request data not equal to category database!"
	errDataNotEqualWithMasterItem            = "your request data not equal to category database!"
	errDataNotEqualWithMasterSKU             = "your request data not equal to sku database!"

	pettyOpenState    status = 0
	pettyCloseState   status = 1
	pettyConfirmState status = 2
)

// NewMobileRepository creates new mobile repository
func NewProductRepository(db *sql.DB) (product.Repository, error) {
	pdb, err := mysql.New(db)
	if err != nil {
		return nil, err
	}
	r := productRepo{pdb}
	return &r, nil
}

type productRepo struct {
	db sqlbuilder.Database
}

func (r *productRepo) FindByID(product int64) (interface{}, error) {

	return nil, nil
}

func (r *productRepo) StoreItem(tokenID string) error {
	return nil
}

func (r *productRepo) FindAllProductRepo() ([]product.Item, error) {

	sql := `select a.id, 
	ifnull(a.code,'') as code,
	ifnull(a.name,'') as name, 
	ifnull(a.eng_name,'') as eng_name, 
	ifnull(a.short_name,'') as short_name,
	ifnull(a.category_code,'') as category_code,
	ifnull(a.type_code,'') as type_code,
	ifnull(a.my_description,'') as my_description, 
	ifnull(a.price,0) as price,
	ifnull(a.my_grade,'') as my_grade,
	ifnull(a.pic_file_name_1,'') as pic_file_name_1,
	ifnull(a.unit_id,0) as unit_id,
	ifnull(a.unit_code,'') as unit_code,
	ifnull(a.def_sale_wh_id,0) as def_sale_wh_id, 
	ifnull(a.def_sale_wh_code,'') as def_sale_wh_code, 
	ifnull(a.degree, 0) as degree, 
	ifnull(a.use_package, 0) as use_package,
	ifnull(a.is_package, 0) as is_package,
	ifnull(a.is_saler, 0) as is_saler,
	ifnull(a.pic_file_saler, '') as pic_file_saler,
	ifnull(b.name,'') as name, 
	ifnull(b.eng_name,'') as eng_name ,
	ifnull(c.sale_price_1,0) as sale_price_1 ,
	ifnull(c.sale_price_2,0) as sale_price_2,
	ifnull(a.payment_destination, 0) as payment_destination
	from item a left JOIN item_category b on a.category_code = b.code left JOIN item_price c on a.code = c.item_code
	where a.active_status = 1 and a.is_package = 0`

	// sql := `select a.id,a.code,a.name,a.eng_name,a.short_name,a.category_code,a.type_code,a.my_description,a.price,a.my_grade,a.pic_file_name_1,a.unit_id,
	// a.unit_code,a.def_sale_wh_id,a.def_sale_wh_code,a.degree,b.name,b.eng_name
	// from item a left JOIN item_category b on a.category_code = b.code`
	items := []product.Item{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	// fmt.Println(clientID)
	for rs.Next() {
		item := product.Item{}

		err = rs.Scan(&item.ID, &item.Code, &item.Name, &item.EngName, &item.ShortName, &item.CategoryCode, &item.TypeCode,
			&item.MyDescription, &item.Price, &item.MyGrade, &item.PicFileName1, &item.UnitID, &item.UnitCode, &item.DefSaleWhID, &item.DefSaleWhCode,
			&item.Degree, &item.UsePackage, &item.IsPackge, &item.IsSaler, &item.PicFileSaler, &item.CategoryName, &item.CategoryNameEng, &item.SalePrice1, &item.SalePrice2, &item.PaymentDestination,
		)
		if err != nil {
			return nil, err
		}
		if item.IsSaler == 1 {
			item.PicFileName1 = item.PicFileSaler
		}

		items = append(items, item)
	}
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *productRepo) FindProductByIDRepo(ID int64) ([]product.ItemDetail, error) {
	sql := `select a.id,
	ifnull(a.code,'') as code,
	ifnull(a.name,'') as name,
	ifnull(a.eng_name,'') as eng_name,
	ifnull(a.group_code,'') as group_code,
	ifnull(a.short_name,'') as short_name,
	ifnull(a.type_code,'') as type_code,
	ifnull(a.my_description,'') as my_description,
	ifnull(a.stock_type,'') as stock_type,
	ifnull(a.unit_id,0) as unit_id,
	ifnull(a.unit_code,'') as unit_code,
	ifnull(a.price,0) as price,
	ifnull(a.buy_unit_id,0) as buy_unit_id,
	ifnull(a.buy_unit_code,'') as buy_unit_code,
	ifnull(a.sale_unit_id,0) as sale_unit_id,
	ifnull(a.sale_unit_code,'') as sale_unit_code,
	ifnull(a.weight,0) as weight,
	ifnull(a.hight,0) as hight,
	ifnull(a.wide,0) as wide,
	ifnull(a.category_code,'') as category_code,
	ifnull(a.sub_cat_code,'') as sub_cat_code,
	ifnull(a.my_grade,'') as my_grade,
	ifnull(a.brand_code,'') as brand_code,
	ifnull(a.color_code,'') as color_code,
	ifnull(a.my_class,'') as my_class,
	ifnull(a.unit_type,0) as unit_type,
	ifnull(a.delivery_charge,0) as delivery_charge,
	ifnull(a.install_charge,0) as install_charge,
	ifnull(a.return_status,0) as return_status,
	ifnull(a.item_status,0) as item_status,
	ifnull(a.last_price,0) as last_price,
	ifnull(a.average_cost,0) as average_cost,
	ifnull(a.cost_type,0) as cost_type,
	ifnull(a.stock_qty,0) as stock_qty,
	ifnull(a.stock_min,0) as stock_min,
	ifnull(a.stock_max,0) as stock_max,
	ifnull(a.pic_file_name_1,'') as pic_file_name_1,
	ifnull(a.pic_file_name_2,'') as pic_file_name_2,
	ifnull(a.pic_file_name_3,'') as pic_file_name_3,
	ifnull(a.pic_file_name_4,'') as pic_file_name_4,
	ifnull(a.pic_file_name_5,'') as pic_file_name_5,
	ifnull(a.def_sale_wh_id,0) as def_sale_wh_id,
	ifnull(a.def_sale_wh_code,'') as def_sale_wh_code,
	ifnull(a.def_sale_shelf_id,0) as def_sale_shelf_id,
	ifnull(a.def_sale_shelf_code,'') as def_sale_shelf_code,
	ifnull(a.def_buy_wh_id,0) as def_buy_wh_id,
	ifnull(a.def_buy_wh_code,'') as def_buy_wh_code,
	ifnull(a.def_buy_shelf_id,0) as def_buy_shelf_id,
	ifnull(a.def_buy_shelf_code,'') as def_buy_shelf_code,
	ifnull(a.use_package,0) as use_package,
	ifnull(a.is_package,0) as is_package,
	ifnull(a.favorite_promo,0) as favorite_promo,
	ifnull(a.degree,0) as degree,
	ifnull(a.package_discount,0) as package_discount,
	ifnull(a.coupon_amount,0) as coupon_amount,
	ifnull(a.campaign_code,'') as campaign_code,
	ifnull(a.is_promotion,0) as is_promotion,
	a.active_status,
	ifnull(b.name,'') as name,
	ifnull(b.eng_name,'') as eng_name,
	ifnull(c.sale_price_1,0) as sale_price_1,
	ifnull(c.sale_price_2,0) as sale_price_2,
	ifnull(c.begin_date,0) as begin_date,
	ifnull(c.end_date,0) as end_date,
	c.is_expire_date,
	a.payment_destination,
	ifnull(a.is_saler,0) as is_saler,
	ifnull(a.pic_file_saler,'') as pic_file_saler
	from item a left JOIN item_category b on a.category_code = b.code left JOIN item_price c on a.code = c.item_code
	where a.id = ?`
	items := []product.ItemDetail{}
	rs, err := r.db.Query(sql, ID)
	if err != nil {
		return nil, err
	}
	// fmt.Println(clientID)
	for rs.Next() {
		item := product.ItemDetail{}
		err = rs.Scan(&item.ID, &item.Code, &item.Name, &item.EngName, &item.GroupCode, &item.ShortName, &item.TypeCode,
			&item.MyDescription, &item.StockType, &item.UnitID, &item.UnitCode, &item.Price, &item.BuyUnitID, &item.BuyUnitCode, &item.SaleUnitID,
			&item.SaleUnitCode, &item.Weight, &item.Hight, &item.Wide, &item.CategoryCode, &item.SubCatCode, &item.MyGrade, &item.BrandCode, &item.ColorCode,
			&item.MyClass, &item.UnitType, &item.DeliveryCharge, &item.InstallCharge, &item.ReturnStatus, &item.ItemStatus, &item.LastPrice, &item.AverageCost,
			&item.CostType, &item.StockQty, &item.StockMin, &item.StockMax, &item.PicFileName1, &item.PicFileName2, &item.PicFileName3, &item.PicFileName4,
			&item.PicFileName5, &item.DefSaleWhID, &item.DefSaleWhCode, &item.DefSaleShelfID, &item.DefSaleShelfCode, &item.DefBuyWhID, &item.DefBuyWhCode,
			&item.DefBuyShelfID, &item.DefBuyShelfCode, &item.UsePackage, &item.IsPackage, &item.FavoritePromo, &item.Degree, &item.PackageDiscount,
			&item.CouponAmount, &item.CampaignCode, &item.IsPromotion, &item.ActiveStatus,
			&item.CategoryName, &item.CategoryNameEng, &item.SalePrice1, &item.SalePrice2, &item.BeginDate, &item.EndDate, &item.IsExpireDate, &item.PaymentDestination,
			&item.IsSaler, &item.PicFileSaler,
		)
		if err != nil {
			return nil, err
		}

		sql1 := `select  a.id,
		a.item_id,
		ifnull(a.parent_code,'') as parent_code,
		ifnull(a.price,0) as price,
		ifnull(a.amount,0) as amount,
		ifnull(a.item_code,'') as item_code,
		ifnull(a.qty,0) as qty,
		ifnull(a.unit_id,0) as unit_id,
		ifnull(a.unit_code,'') as unit_code,
		ifnull(a.line_number,0) as line_number,
		ifnull(b.name,'') as name
		from item_components a left JOIN item b on a.item_id = b.id
		where a.parent_code = ? and a.active_status = 1`
		rs1, err := r.db.Query(sql1, &item.Code)
		if err != nil {
			return nil, err
		}
		for rs1.Next() {
			itemProsub := product.ItemPromoSub{}
			err = rs1.Scan(&itemProsub.ID, &itemProsub.ItemID, &itemProsub.ParentCode, &itemProsub.Price, &itemProsub.Amount, &itemProsub.ItemCode, &itemProsub.Qty,
				&itemProsub.UnitID, &itemProsub.UnitCode, &itemProsub.LineNumber, &itemProsub.Name,
			)
			if err != nil {
				return nil, err
			}
			item.ItemPromoSub = append(item.ItemPromoSub, itemProsub)
		}
		println("setper =", item.ID)

		sql2 := `select id,
		ifnull(url,'') as url
		from pic_item_sub
		where item_id = ? and active_status = 1`
		rs2, err := r.db.Query(sql2, item.ID)
		if err != nil {
			return nil, err
		}
		for rs2.Next() {
			itemPicSub := product.PicSub{}
			err = rs2.Scan(&itemPicSub.ID, &itemPicSub.URL)
			if err != nil {
				return nil, err
			}
			item.PicSub = append(item.PicSub, itemPicSub)
		}
		items = append(items, item)
	}
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *productRepo) FindCategoryRepo() ([]product.Category, error) {
	fmt.Println("start repo FindCategoryRepo")
	sql := `select a.id,a.code,a.name,a.eng_name,a.item_count,ifnull(a.my_description,0) as my_description
	from item_category a`
	cats := []product.Category{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		cat := product.Category{}

		err = rs.Scan(&cat.ID, &cat.Code, &cat.Name, &cat.EngName, &cat.ItemCount, &cat.MyDescription)
		if err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	if err != nil {
		return nil, err
	}
	fmt.Println("return ", cats)
	return cats, nil
}

func (r *productRepo) FindCategoryByCodeRepo(Code string) ([]product.Item, error) {
	sql := `select a.id,
	ifnull(a.code,'' ) as code,
	ifnull(a.name,'' ) as name,
	ifnull(a.eng_name,'' ) as eng_name,
	ifnull(a.short_name,'' ) as short_name,
	ifnull(a.category_code,'' ) as category_code,
	ifnull(a.type_code,'' ) as type_code,
	ifnull(a.my_description,'' ) as my_description,
	ifnull(a.price,0 ) as price,
	ifnull(a.my_grade,'' ) as my_grade,
	ifnull(a.pic_file_name_1,'' ) as pic_file_name_1,
	ifnull(a.unit_id,0) as unit_id,
	ifnull(a.unit_code,'' ) as unit_code,
	ifnull(a.def_sale_wh_id,0 ) as def_sale_wh_id,
	ifnull(a.def_sale_wh_code,0 ) as def_sale_wh_code,
	ifnull(a.degree,0 ) as degree,
	ifnull(a.use_package,0 ) as use_package,
	ifnull(a.is_package,0 ) as is_package,
	ifnull(a.stock_type,'' ) as stock_type,
	ifnull(a.is_saler, 0) as is_saler,
	ifnull(a.pic_file_saler, '') as pic_file_saler,
	ifnull(b.name,'' ) as name,
	ifnull(b.eng_name,'' ) as eng_name,
	ifnull(c.sale_price_1,0) as sale_price_1,
	ifnull(c.sale_price_2,0) as sale_price_2,
	ifnull(c.begin_date,'') as begin_date,
	ifnull(c.end_date,'') as end_date,
	ifnull(c.is_expire_date,0) as is_expire_date,
	a.payment_destination
	from item a left JOIN item_category b on a.category_code = b.code left JOIN item_price c on a.code = c.item_code
	where a.active_status = 1 and category_code = ? and a.is_package = 0`
	items := []product.Item{}
	rs, err := r.db.Query(sql, Code)
	if err != nil {
		return nil, err
	}
	// fmt.Println(clientID)
	for rs.Next() {
		item := product.Item{}
		err = rs.Scan(&item.ID, &item.Code, &item.Name, &item.EngName, &item.ShortName, &item.CategoryCode, &item.TypeCode,
			&item.MyDescription, &item.Price, &item.MyGrade, &item.PicFileName1, &item.UnitID, &item.UnitCode, &item.DefSaleWhID, &item.DefSaleWhCode,
			&item.Degree, &item.UsePackage, &item.IsPackge, &item.StockType, &item.IsSaler, &item.PicFileSaler, &item.CategoryName, &item.CategoryNameEng, &item.SalePrice1,
			&item.SalePrice2, &item.BeginDate, &item.EndDate, &item.IsExpireDate, &item.PaymentDestination,
		)
		if err != nil {
			return nil, err
		}

		if item.IsSaler == 1 {
			item.PicFileName1 = item.PicFileSaler
		}
		sql1 := `select  id,
		item_id,
		ifnull(parent_code,'') as parent_code,
		ifnull(price,0) as price,
		ifnull(amount,0) as amount
		from item_components
		where item_code = ?`
		rs1, err := r.db.Query(sql1, &item.Code)
		if err != nil {
			return nil, err
		}
		for rs1.Next() {
			itemProsub := product.ItemPromoSub{}
			err = rs1.Scan(&itemProsub.ID, &itemProsub.ItemID, &itemProsub.ParentCode, &itemProsub.Price, &itemProsub.Amount)
			if err != nil {
				return nil, err
			}
			item.ItemPromoSub = append(item.ItemPromoSub, itemProsub)
		}

		items = append(items, item)
	}

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *productRepo) FavoriteProductRepo() ([]product.Item, error) {
	sql := `select a.id,
	ifnull(a.code,'') as code,
	ifnull(a.name,'') as name,
	ifnull(a.eng_name,'') as eng_name,
	ifnull(a.short_name,'') as short_name,
	ifnull(a.category_code,'') as category_code,
	ifnull(a.type_code,'') as type_code,
	ifnull(a.my_description,'') as my_description,
	ifnull(a.price,0) as price,
	ifnull(a.my_grade,'') as my_grade,
	ifnull(a.pic_file_name_1,'') as pic_file_name_1,
	ifnull(a.unit_id,0) as unit_id,
	ifnull(a.unit_code,'') as unit_code,
	ifnull(a.def_sale_wh_id,0) as def_sale_wh_id,
	ifnull(a.def_sale_wh_code,'') as def_sale_wh_code,
	ifnull(a.degree,0) as degree,
	ifnull(b.name,'') as name,
	ifnull(b.eng_name,'') as eng_name,
	ifnull(c.sale_price_1,0) as sale_price_1,
	ifnull(c.sale_price_2,0) as sale_price_2,
	a.payment_destination
	from item a left JOIN item_category b on a.category_code = b.code left JOIN item_price c on a.code = c.item_code
	where a.active_status = 1 and a.favorite = 1`
	items := []product.Item{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	// fmt.Println(clientID)
	for rs.Next() {
		item := product.Item{}

		err = rs.Scan(&item.ID, &item.Code, &item.Name, &item.EngName, &item.ShortName, &item.CategoryCode, &item.TypeCode,
			&item.MyDescription, &item.Price, &item.MyGrade, &item.PicFileName1, &item.UnitID, &item.UnitCode, &item.DefSaleWhID, &item.DefSaleWhCode,
			&item.Degree, &item.CategoryName, &item.CategoryNameEng, &item.SalePrice1, &item.SalePrice2, &item.PaymentDestination,
		)
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

func (r *productRepo) FindPackageRepo() ([]product.Item, error) {

	sql := `select a.id,
	ifnull(a.code,'') as code,
	ifnull(a.name,'') as name,
	ifnull(a.eng_name,'') as eng_name,
	ifnull(a.short_name,'') as short_name,
	ifnull(a.category_code,'') as category_code,
	ifnull(a.type_code,'') as type_code,
	ifnull(a.my_description,'') as my_description,
	ifnull(a.price,0) as price,
	ifnull(a.my_grade,'') as my_grade,
	ifnull(a.pic_file_name_1,'') as pic_file_name_1,
	ifnull(a.unit_id,0) as unit_id,
	ifnull(a.unit_code,'') as unit_code,
	ifnull(a.def_sale_wh_id,0) as def_sale_wh_id,
	ifnull(a.def_sale_wh_code,'') as def_sale_wh_code,
	ifnull(a.degree, 0) as degree,
	ifnull(a.is_package,0) as is_package,
	ifnull(b.name,'') as name,
	ifnull(b.eng_name,'') as eng_name,
	ifnull(c.sale_price_1,0) as sale_price_1,
	ifnull(c.sale_price_2,0) as sale_price_2,
	a.payment_destination
	from item a left JOIN item_category b on a.category_code = b.code left JOIN item_price c on a.code = c.item_code
	where a.active_status = 1 and is_package = 1`

	// sql := `select a.id,a.code,a.name,a.eng_name,a.short_name,a.category_code,a.type_code,a.my_description,a.price,a.my_grade,a.pic_file_name_1,a.unit_id,
	// a.unit_code,a.def_sale_wh_id,a.def_sale_wh_code,a.degree,b.name,b.eng_name
	// from item a left JOIN item_category b on a.category_code = b.code`
	items := []product.Item{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	// fmt.Println(clientID)
	for rs.Next() {
		item := product.Item{}

		err = rs.Scan(&item.ID, &item.Code, &item.Name, &item.EngName, &item.ShortName, &item.CategoryCode, &item.TypeCode,
			&item.MyDescription, &item.Price, &item.MyGrade, &item.PicFileName1, &item.UnitID, &item.UnitCode, &item.DefSaleWhID, &item.DefSaleWhCode,
			&item.Degree, &item.IsPackge, &item.CategoryName, &item.CategoryNameEng, &item.SalePrice1, &item.SalePrice2, &item.PaymentDestination,
		)
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

func (r *productRepo) FavoritePromotion() ([]product.Item, error) {

	sql := `select a.id,
	 ifnull(a.code,'') as code,
	 ifnull(a.name,'') as name,
	 ifnull(a.eng_name,'') as eng_name,
	 ifnull(a.short_name,'') as short_name,
	 ifnull(a.category_code,'') as category_code,
	 ifnull(a.type_code,'') as type_code,
	 ifnull(a.my_description,'') as my_description,
	 ifnull(a.price,0) as price,
	 ifnull(a.my_grade,'') as my_grade,
	 ifnull(a.pic_file_name_1,'') as pic_file_name_1,
	 ifnull(a.unit_id,0) as unit_id,
	 ifnull(a.unit_code,'') as unit_code,
	 ifnull(a.def_sale_wh_id,0) as def_sale_wh_id,
	 ifnull(a.def_sale_wh_code,'') as def_sale_wh_code,
	 ifnull(a.degree, 0) as degree,
	 ifnull(a.use_package,0) as use_package,
	 ifnull(a.is_package,0) as is_package,
	 ifnull(b.name,'') as name,
	 ifnull(b.eng_name,'') as eng_name ,
	 ifnull(c.sale_price_1,0) as sale_price_1 ,
	 ifnull(c.sale_price_2,0) as sale_price_2,
	 a.payment_destination
	 from item a left JOIN item_category b on a.category_code = b.code left JOIN item_price c on a.code = c.item_code
	 where a.active_status = 1 and a.is_package = 0 and a.favorite_promo = 1`

	// sql := `select a.id,a.code,a.name,a.eng_name,a.short_name,a.category_code,a.type_code,a.my_description,a.price,a.my_grade,a.pic_file_name_1,a.unit_id,
	// a.unit_code,a.def_sale_wh_id,a.def_sale_wh_code,a.degree,b.name,b.eng_name
	// from item a left JOIN item_category b on a.category_code = b.code`
	items := []product.Item{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	// fmt.Println(clientID)
	for rs.Next() {
		item := product.Item{}

		err = rs.Scan(&item.ID, &item.Code, &item.Name, &item.EngName, &item.ShortName, &item.CategoryCode, &item.TypeCode,
			&item.MyDescription, &item.Price, &item.MyGrade, &item.PicFileName1, &item.UnitID, &item.UnitCode, &item.DefSaleWhID, &item.DefSaleWhCode,
			&item.Degree, &item.UsePackage, &item.IsPackge, &item.CategoryName, &item.CategoryNameEng, &item.SalePrice1, &item.SalePrice2, &item.PaymentDestination,
		)
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

func (r *productRepo) FindYoutubeList() ([]product.YoutubeList, error) {

	sql := `select a.id, 
	ifnull(a.link_address,'') as link_address, 
	ifnull(a.my_description,'') as my_description
	from youtube_link a
	where a.active_status = 1`

	items := []product.YoutubeList{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		item := product.YoutubeList{}
		err = rs.Scan(&item.ID, &item.LinkAddress, &item.MyDescription)
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

func (r *productRepo) FindCampaign() ([]product.Campaign, error) {

	sql := `select a.id, 
	ifnull(a.code,'') as code, 
	ifnull(a.my_description,'') as my_description,
	ifnull(a.pic_file_name_1,'') as pic_file_name_1,
	ifnull(a.begin_date,'') as begin_date,
	ifnull(a.end_date,'') as end_date
	from promotion_campaign a
	where a.active_status = 1`

	items := []product.Campaign{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		item := product.Campaign{}
		err = rs.Scan(&item.ID, &item.Code, &item.MyDescription, &item.PicFileName1, &item.BeginDate, &item.EndDate)
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

func (r *productRepo) ItemList(ID int64) ([]product.ItemList, error) {
	var sql string
	switch ID {
	case 1:
		sql = `select a.id,
			ifnull(a.code,'') as code,
			ifnull(a.name,'') as name,
			ifnull(a.eng_name,'') as eng_name,
			ifnull(a.category_code,'') as category_code,
			ifnull(a.price,0) as price,
			ifnull(a.pic_file_name_1,"") as pic_file_name_1,
			ifnull(b.name,'') as name,
			ifnull(b.eng_name,'') as eng_name ,
			ifnull(c.sale_price_1,0) as sale_price_1 ,
			ifnull(c.sale_price_2,0) as sale_price_2
			from item a left JOIN item_category b on a.category_code = b.code left JOIN item_price c on a.code = c.item_code
			where a.is_package = 0 and a.is_promotion = 0`
	case 2:
		sql = `select a.id,
			ifnull(a.code,'') as code,
			ifnull(a.name,'') as name,
			ifnull(a.eng_name,'') as eng_name,
			ifnull(a.category_code,'') as category_code,
			ifnull(a.price,0) as price,
			ifnull(a.pic_file_name_1,"") as pic_file_name_1,
			ifnull(b.name,'') as name,
			ifnull(b.eng_name,'') as eng_name ,
			ifnull(c.sale_price_1,0) as sale_price_1 ,
			ifnull(c.sale_price_2,0) as sale_price_2
			from item a left JOIN item_category b on a.category_code = b.code left JOIN item_price c on a.code = c.item_code
			where a.is_promotion = 1`
	case 3:
		sql = `select a.id,
			ifnull(a.code,'') as code,
			ifnull(a.name,'') as name,
			ifnull(a.eng_name,'') as eng_name,
			ifnull(a.category_code,'') as category_code,
			ifnull(a.price,0) as price,
			ifnull(a.pic_file_name_1,"") as pic_file_name_1,
			ifnull(b.name,'') as name,
			ifnull(b.eng_name,'') as eng_name ,
			ifnull(c.sale_price_1,0) as sale_price_1 ,
			ifnull(c.sale_price_2,0) as sale_price_2
			from item a left JOIN item_category b on a.category_code = b.code left JOIN item_price c on a.code = c.item_code
			where a.is_package = 1`
	}
	items := []product.ItemList{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	// fmt.Println(clientID)
	for rs.Next() {
		item := product.ItemList{}

		err = rs.Scan(&item.ID, &item.Code, &item.Name, &item.EngName, &item.CategoryCode, &item.Price, &item.PicFileName1,
			&item.CategoryName, &item.CategoryNameEng, &item.SalePrice1, &item.SalePrice2,
		)
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

func (r *productRepo) UnitList() ([]product.ItemUnit, error) {
	sql := `select a.id, 
	ifnull(a.company_id,'') as company_id,
	ifnull(a.code,'') as code, 
	ifnull(a.name,'') as name,
	ifnull(a.rate_1,'') as rate_1
	from item_unit a
	where a.active_status = 1`

	items := []product.ItemUnit{}
	rs, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		item := product.ItemUnit{}
		err = rs.Scan(&item.ID, &item.CompanyID, &item.Code, &item.Name, &item.Rate1)
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

func (r *productRepo) ItemHistory(ID int64, UserID string) ([]product.ItemHistory, error) {
	sql := `select
	ifnull(sum(a.qty),0) as qty_all,
	ifnull(sum(a.price * a.qty),0) as price
	from order_sub a left JOIN orders b on a.order_id = b.id
	where a.item_id = ? and b.order_status > 1`
	rs, err := r.db.Query(sql, ID)
	items := []product.ItemHistory{}
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		item := product.ItemHistory{}
		err = rs.Scan(&item.SumQTY, &item.SumAllAmount)
		if err != nil {
			return nil, err
		}

		sql1 := `select
		ifnull(b.doc_date,'') as doc_date,
		ifnull(b.user_name,'') as user_name,
		ifnull(b.doc_no,'') as doc_no,
		ifnull(a.qty,'') as qty,
		ifnull(a.price,'') as price,
		a.qty * a.price as sum_price
		from order_sub a left JOIN orders b on a.order_id = b.id
		where a.item_id = ? and b.order_status > 1`
		rs1, err := r.db.Query(sql1, ID)
		if err != nil {
			return nil, err
		}
		for rs1.Next() {
			item_sub := product.ItemHistoryDetail{}
			err = rs1.Scan(&item_sub.Date, &item_sub.Name, &item_sub.DocNo, &item_sub.QTY, &item_sub.Price, &item_sub.SumAmount)
			if err != nil {
				return nil, err
			}
			item.ItemHistoryDetail = append(item.ItemHistoryDetail, item_sub)
		}
		items = append(items, item)
	}
	if err != nil {
		return nil, err
	}
	return items, nil
}
