package product

type Product struct {
	ItemID   int64  `json:"item_id"`
	ItemName string `json:"item_name"`
	Image    string `json:"image"`
	Degree   int64  `json:"degree"`
	Price    int64  `json:"price"`
	Qty      int64  `json:"qty"`
}

type Item struct {
	ID                 int64   `json:"id"`
	Code               string  `json:"code"`
	Name               string  `json:"name"`
	EngName            string  `json:"eng_name"`
	ShortName          string  `json:"short_name"`
	CategoryCode       string  `json:"category_code"`
	TypeCode           string  `json:"type_code"`
	MyDescription      string  `json:"my_description"`
	UnitID             int64   `json:"unit_id"`
	UnitCode           string  `json:"unit_code"`
	DefSaleWhID        int64   `json:"def_sale_wh_id"`
	DefSaleWhCode      string  `json:"def_sale_wh_code"`
	Price              int64   `json:"price"`
	MyGrade            string  `json:"my_grade"`
	UsePackage         int64   `json:"use_package"`
	IsPackge           int64   `json:"is_package"`
	IsSaler            int64   `json:"is_saler"`
	PicFileSaler       string  `json:"pic_file_saler"`
	StockType          int64   `json:"stock_type"`
	Degree             int64   `json:"degree"`
	PicFileName1       string  `json:"pic_file_name_1"`
	CategoryName       string  `json:"category_name"`
	CategoryNameEng    string  `json:"category_name_eng"`
	SalePrice1         float64 `json:"sale_price_1"`
	SalePrice2         float64 `json:"sale_price_2"`
	BeginDate          string  `json:"begin_date"`
	EndDate            string  `json:"end_date"`
	IsExpireDate       int64   `json:"is_expire_date"`
	PaymentDestination int64   `json:"payment_destination"`
	ItemPromoSub       []ItemPromoSub
}

type ItemPromoSub struct {
	ID         int64   `json:"id"`
	ParentCode string  `json:"parent_code"`
	ItemID     int64   `json:"item_id"`
	ItemCode   string  `json:"code"`
	Name       string  `json:"name"`
	Qty        int64   `json:"qty"`
	Price      float64 `json:"price"`
	Amount     float64 `json:"amount"`
	UnitID     int64   `json:"unit_id"`
	UnitCode   string  `json:"unit_code"`
	LineNumber int64   `json:"line_number"`
}
type ItemVolme struct {
	ID                int64   `json:"id"`
	ItemID            int64   `json:"item_id"`
	ItemName          string  `json:"item_name"`
	PromoCode         string  `json:"promo_code"`
	PromoType         int64   `json:"promo_type"`
	BeginDate         string  `json:"begin_date"`
	EndDate           string  `json:"end_date"`
	ProPrice          float64 `json:"pro_price"`
	ProDiscountAmount float64 `json:"discount_amount"`
}

type ItemDetail struct {
	ID               int64   `json:"id"`
	Code             string  `json:"code"`
	Name             string  `json:"name"`
	EngName          string  `json:"eng_name"`
	ShortName        string  `json:"short_name"`
	GroupCode        string  `json:"group_code"`
	TypeCode         string  `json:"type_code"`
	MyDescription    string  `json:"my_description"`
	StockType        int64   `json:"stock_type"`
	UnitID           int64   `json:"unit_id"`
	UnitCode         string  `json:"unit_code"`
	Price            float64 `json:"price"`
	BuyUnitID        int64   `json:"buy_unit_id"`
	BuyUnitCode      string  `json:"buy_unit_code"`
	SaleUnitID       int64   `json:"sale_unit_id"`
	SaleUnitCode     string  `json:"sale_unit_code"`
	Weight           float64 `json:"weight"`
	Hight            float64 `json:"hight"`
	Wide             float64 `json:"wide"`
	CategoryCode     string  `json:"category_code"`
	SubCatCode       string  `json:"sub_cat_code"`
	MyGrade          string  `json:"my_grade"`
	BrandCode        string  `json:"brand_code"`
	ColorCode        string  `json:"color_code"`
	MyClass          string  `json:"my_class"`
	UnitType         int64   `json:"unit_type"`
	DeliveryCharge   int64   `json:"delivery_charge"`
	InstallCharge    int64   `json:"install_charge"`
	ReturnStatus     int64   `json:"return_status"`
	ItemStatus       int64   `json:"item_status"`
	LastPrice        float64 `json:"last_price"`
	AverageCost      float64 `json:"average_cost"`
	CostType         int64   `json:"cost_type"`
	StockQty         float64 `json:"stock_qty"`
	StockMin         float64 `json:"stock_min"`
	StockMax         float64 `json:"stock_max"`
	PicFileName1     string  `json:"pic_file_name_1"`
	PicFileName2     string  `json:"pic_file_name_2"`
	PicFileName3     string  `json:"pic_file_name_3"`
	PicFileName4     string  `json:"pic_file_name_4"`
	PicFileName5     string  `json:"pic_file_name_5"`
	DefSaleWhID      int64   `json:"def_sale_wh_id"`
	DefSaleWhCode    string  `json:"def_sale_wh_code"`
	DefSaleShelfID   int64   `json:"def_sale_shelf_id"`
	DefSaleShelfCode string  `json:"def_sale_shelf_code"`
	DefBuyWhID       int64   `json:"def_buy_wh_id"`
	DefBuyWhCode     string  `json:"def_buy_wh_code"`
	DefBuyShelfID    int64   `json:"def_buy_shelf_id"`
	DefBuyShelfCode  string  `json:"def_buy_shelf_code"`
	UsePackage       int64   `json:"use_package"`
	IsPackage        int64   `json:"is_package"`
	FavoritePromo    int64   `json:"favorite_promo"`
	Degree           int64   `json:"degree"`
	PackageDiscount  float64 `json:"package_discount"`
	CouponAmount     float64 `json:"coupon_amount"`
	CampaignCode     string  `json:"campaign_code"`
	IsPromotion      int64   `json:"is_promotion"`
	ActiveStatus     int64   `json:"active_status"`
	//////////////////item_price///////////////////////
	CategoryName       string  `json:"category_name"`
	CategoryNameEng    string  `json:"category_name_eng"`
	SalePrice1         float64 `json:"sale_price_1"`
	SalePrice2         float64 `json:"sale_price_2"`
	BeginDate          string  `json:"begin_date"`
	EndDate            string  `json:"end_date"`
	IsExpireDate       int64   `json:"is_expire_date"`
	PaymentDestination int64   `json:"payment_destination"`
	IsSaler            int64   `json:"is_saler"`
	PicFileSaler       string  `json:"pic_file_saler"`
	ItemPromoSub       []ItemPromoSub
	PicSub             []PicSub
}

type ItemCategory struct {
	ID           int64  `json:"id"`
	CategoryName string `json:"category_name"`
	Image        string `json:"image"`
}
type Category struct {
	ID            int64  `json:"id"`
	CompanyID     int64  `json:"company_id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	EngName       string `json:"eng_name"`
	ItemCount     int64  `json:"item_count"`
	MyDescription string `json:"my_description"`
}

// type ItemComponant struct {
// 	ID         int64   `json:"id"`
// 	ParentCode string  `json:"parent_code"`
// 	ItemID     int64   `json:"item_id"`
// 	ItemCode   string  `json:"item_code"`
// 	Qty        int64   `json:"qty"`
// 	Price      float64 `json:"price"`
// 	Amount     float64 `json:"amount"`
// 	UnitID     int64   `json:"unit_id"`
// 	UnitCode   string  `json:"unit_code"`
// 	LineNumber int64   `json:"line_number"`
// }

type ItemAdd struct {
	ID              int64   `json:"id"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	EngName         string  `json:"eng_name"`
	ShortName       string  `json:"short_name"`
	GroupCode       string  `json:"group_code"`
	TypeCode        string  `json:"type_code"`
	MyDescription   string  `json:"my_description"`
	StockType       int64   `json:"stock_type"`
	UnitID          int64   `json:"unit_id"`
	UnitCode        string  `json:"unit_code"`
	Price           float64 `json:"price"`
	BuyUnitID       int64   `json:"buy_unit_id"`
	BuyUnitCode     string  `json:"buy_unit_code"`
	SaleUnitID      int64   `json:"sale_unit_id"`
	SaleUnitCode    string  `json:"sale_unit_code"`
	Weight          float64 `json:"weight"`
	Hight           float64 `json:"hight"`
	Wide            float64 `json:"wide"`
	CategoryCode    string  `json:"category_code"`
	SubCatCode      string  `json:"sub_cat_code"`
	MyGrade         string  `json:"my_grade"`
	BrandCode       string  `json:"brand_code"`
	ColorCode       string  `json:"color_code"`
	MyClass         string  `json:"my_class"`
	UnitType        int64   `json:"unit_type"`
	DeliveryCharge  int64   `json:"delivery_charge"`
	InstallCharge   int64   `json:"install_charge"`
	ReturnStatus    int64   `json:"return_status"`
	ItemStatus      int64   `json:"item_status"`
	LastPrice       float64 `json:"last_price"`
	AverageCost     float64 `json:"average_cost"`
	CostType        int64   `json:"cost_type"`
	PicFileName1    string  `json:"pic_file_name_1"`
	PicFileName2    string  `json:"pic_file_name_2"`
	PicFileName3    string  `json:"pic_file_name_3"`
	PicFileName4    string  `json:"pic_file_name_4"`
	PicFileName5    string  `json:"pic_file_name_5"`
	PicSub          []PicSub
	SalePrice1      float64 `json:"sale_price_1"`
	SalePrice2      float64 `json:"sale_price_2"`
	FavoritePromo   int64   `json:"favorite_promo"`
	Degree          int64   `json:"degree"`
	Favorite        int64   `json:"favorite"`
	PackageDiscount float64 `json:"package_discount"`
	CouponAmount    float64 `json:"coupon_amount"`
	IsPromotion     int64   `json:"is_promotion"`
	IsSaler         int64   `json:"is_saler"`
	PicFileSaler    string  `json:"pic_file_saler"`
	ActiveStatus    int64   `json:"active_status"`
	PerDiscount     int64   `json:"per_discount"`
	BeginDate       string  `json:"begin_date"`
	EndDate         string  `json:"end_date"`
	ItemPromoSub    []ItemPromoSub
}

type PicSub struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}

type CategoryAdd struct {
	ID            int64  `json:"id"`
	CompanyID     int64  `json:"company_id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	EngName       string `json:"eng_name"`
	ItemCount     int64  `json:"item_count"`
	ActiveStatus  int64  `json:"active_status"`
	MyDescription string `json:"my_description"`
}

type YoutubeList struct {
	ID            int64  `json:"id"`
	LinkAddress   string `json:"link_address"`
	MyDescription string `json:"my_description"`
}

type Campaign struct {
	ID            int64  `json:"id"`
	Code          string `json:"code"`
	MyDescription string `json:"my_description"`
	PicFileName1  string `json:"pic_file_name_1"`
	BeginDate     string `json:"begin_date"`
	EndDate       string `json:"end_date"`
}

type ItemList struct {
	ID              int64   `json:"id"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	EngName         string  `json:"eng_name"`
	CategoryCode    string  `json:"category_code"`
	Price           float64 `json:"price"`
	SalePrice1      float64 `json:"sale_price_1"`
	SalePrice2      float64 `json:"sale_price_2"`
	PicFileName1    string  `json:"pic_file_name_1"`
	CategoryName    string  `json:"category_name"`
	CategoryNameEng string  `json:"category_name_eng"`
}

type ItemUnit struct {
	ID        int64  `json:"id"`
	CompanyID int64  `json:"company_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Rate1     int64  `json:"rate_1"`
}

type ItemHistory struct {
	SumAllAmount      float64 `json:"sum_all_amount"`
	SumQTY            float64 `json:"sum_qty"`
	ItemHistoryDetail []ItemHistoryDetail
}
type ItemHistoryDetail struct {
	Date      string  `json:"date"`
	Name      string  `json:"name"`
	DocNo     string  `json:"cod_no"`
	QTY       float64 `json:"qty"`
	Price     float64 `json:"price"`
	SumAmount float64 `json:"sum_amount"`
}
