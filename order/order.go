package order

import "time"

type Payment struct {
	ItemID   int64  `json:"item_id"`
	ItemName string `json:"item_name"`
	Price    int64  `json:"price"`
	Qty      int64  `json:"qty"`
}

type Cart struct {
	ItemID int64 `json:"item_id"`
	Qty    int64 `json:"qty"`
}
type Basket struct {
	ID             int64       `json:"id"`
	UserID         string      `json:"user_id"`
	SaleID         int64       `json:"sale_id"`
	DocNo          string      `json:"doc_no"`
	ItemAmount     float64     `json:"item_amount"`
	DiscountAmount float64     `json:"discount_amount"`
	NetAmount      float64     `json:"net_amount"`
	MyDescription  string      `json:"my_description"`
	CreateTime     time.Time   `json:"create_time"`
	BasketSub      []BasketSub `json:"basket_sub"`
}

type BasketSub struct {
	BasketID       int64   `json:"basket_id"`
	CustID         int64   `json:"cust_id"`
	SaleID         int64   `json:"sale_id"`
	ItemID         int64   `json:"item_id"`
	ItemName       string  `json:"item_name"`
	Qty            float64 `json:"qty"`
	UnitID         int64   `json:"unit_id"`
	UnitCode       string  `json:"unit_code"`
	WhID           int64   `json:"wh_id"`
	WhCode         string  `json:"wh_code"`
	Price          float64 `json:"price"`
	DiscountAmount float64 `json:"discount_amount"`
	ItemAmount     float64 `json:"item_amount"`
	NetAmount      float64 `json:"net_amount"`
	LineNumber     float64 `json:"line_number"`
	PicFileName1   string  `json:"pic_file_name_1"`
	CouponEnabled  int64   `json:"coupon_enabled"`
}

type Order struct {
	ID                  int64     `json:"id"`
	UserID              string    `json:"user_id"`
	UserName            string    `json:"user_name"`
	Talephone           string    `json:"telephone"`
	SaleID              int64     `json:"sale_id"`
	SaleName            string    `json:"sale_name"`
	DocNo               string    `json:"doc_no"`
	SumOfItemAmount     float64   `json:"sum_of_item_amount"`
	DiscountAmount      float64   `json:"discount_amount"`
	AfterDiscountAmount float64   `json:"after_discount_amount"`
	BeforeTax           float64   `json:"before_tax"`
	TaxAmount           float64   `json:"tax_amount"`
	TotalAmount         float64   `json:"total_amount"`
	SaleType            int64     `json:"sale_type"`
	SumCashAmount       float64   `json:"sum_cash_amount"`
	SumCreditAmount     float64   `json:"sum_credit_amount"`
	SumDepositAmount    float64   `json:"sum_deposit_amount"`
	SumCouponAmount     float64   `json:"sum_coupon_amount"`
	SumBankAmount       float64   `json:"sum_bank_amount"`
	ChangeAmount        float64   `json:"change_amount"`
	NetDebtDmount       float64   `json:"net_debt_amount"`
	MyDescription       string    `json:"my_description"`
	OrderStatus         int64     `json:"order_status"`
	DeliveryDate        time.Time `json:"delivery_date"`
	Distance            float64   `json:"distance"`
	DistanceAmount      float64   `json:"distance_amount"`
	DeliveryLink        string    `json:"delivery_link"`
	DeliveryID          string    `json:"delivery_id"`
	ReferralID          string    `json:"referral_id"`
	PicSlip             string    `json:"pic_slip"`
	SendText            []string  `json:"send_text"`
	CreateTime          time.Time `json:"create_time"`
	TrackingID          string    `json:"tracking_id"`
	IsPackage           int64     `json:"is_package"`
	OrderSub            []OrderSub
}

type OrderSub struct {
	ID             int64   `json:"id"`
	OrderID        int64   `json:"order_id"`
	ItemID         int64   `json:"item_id"`
	ItemName       string  `json:"item_name"`
	WhID           int64   `json:"wh_id"`
	WhCode         string  `json:"wh_code"`
	ShelfID        int64   `json:"shelf_id"`
	ShelfCode      string  `json:"shelf_code"`
	Qty            float64 `json:"qty"`
	CnQty          float64 `json:"cn_qty"`
	UnitID         int64   `json:"unit_id"`
	UnitCode       string  `json:"unit_code"`
	Price          float64 `json:"price"`
	DiscountAmount float64 `json:"discount_amount"`
	ItemAmount     float64 `json:"item_amount"`
	AverageCost    float64 `json:"average_cost"`
	SumOfCost      float64 `json:"sum_of_cost"`
	Rate1          int64   `json:"rate_1"`
	StockType      int64   `json:"stock_type"`
	BasketID       int64   `json:"basket_id"`
	ItemSubCat     string  `json:"item_sub_cat"`
	ItemGroup      string  `json:"item_group"`
	TypeCode       string  `json:"type_code"`
	Point          float64 `json:"point"`
	IsCancel       int64   `json:"is_cancel"`
	RefLineNumber  int64   `json:"ref_line_number"`
	PicFileName1   string  `json:"pic_file_name_1"`
	LineNumber     int64   `json:"line_number"`
}

type Bank struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	BankCode string `json:"bank_code"`
	BankName string `json:"bank_name"`
	CodeName string `json:"code_name"`
	Image    string `json:"image"`
}

type BankList struct {
	ID          int64  `json:"id"`
	BankCode    string `json:"bank_code"`
	BankName    string `json:"bank_name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
type Delivery struct {
	ID   int64   `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

type DeliveryAll struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type OrderSend struct {
	ID           int64  `json:"id"`
	OrderStatus  int64  `json:"order_status"`
	DeliveryLink string `json:"delivery_link"`
	DeliveryID   string `json:"delivery_id"`
	AddressID    string `json:"address_id"`
	DocNo        string `json:"doc_no"`
	TrackingID   string `json:"tracking_id"`
	OrderSub     []OrderSub
}

type OrderSendDetail struct {
	ID                  int64     `json:"id"`
	CompanyID           int64     `json:"company_id"`
	UserID              string    `json:"user_id"`
	UserName            string    `json:"user_name"`
	SaleID              int64     `json:"sale_id"`
	SaleName            string    `json:"sale_name"`
	DocNo               string    `json:"doc_no"`
	SumOfItemAmount     float64   `json:"sum_of_item_amount"`
	DiscountAmount      float64   `json:"discount_amount"`
	AfterDiscountAmount float64   `json:"after_discount_amount"`
	BeforeTax           float64   `json:"before_tax"`
	TaxAmount           float64   `json:"tax_amount"`
	TotalAmount         float64   `json:"total_amount"`
	SaleType            int64     `json:"sale_type"`
	SumCashAmount       float64   `json:"sum_cash_amount"`
	SumCreditAmount     float64   `json:"sum_credit_amount"`
	SumDepositAmount    float64   `json:"sum_deposit_amount"`
	SumCouponAmount     float64   `json:"sum_coupon_amount"`
	SumBankAmount       float64   `json:"sum_bank_amount"`
	ChangeAmount        float64   `json:"change_amount"`
	NetDebtDmount       float64   `json:"net_debt_amount"`
	MyDescription       string    `json:"my_description"`
	OrderStatus         int64     `json:"order_status"`
	DeliveryDate        time.Time `json:"delivery_date"`
	Distance            float64   `json:"distance"`
	DistanceAmount      float64   `json:"distance_amount"`
	DeliveryLink        string    `json:"delivery_link"`
	DeliveryID          string    `json:"delivery_id"`
	DeliveryName        string    `json:"delivery_name"`
	ReferralID          string    `json:"referral_id"`
	PicSlip             string    `json:"pic_slip"`
	SendText            []string  `json:"send_text"`
	CreateTime          time.Time `json:"create_time"`
	TrackingID          string    `json:"tracking_id"`
	AddressID           int64     `json:"address_id"`
	IsPackage           int64     `json:"is_package"`
	OrderSub            []OrderSub
	OrderSendAddress    OrderSendAddress
	OrderCompanyDetail  OrderCompanyDetail
}

type OrderSendAddress struct {
	AddressID   int64  `json:"addr_id,omitempty" db:"addr_id,omitempty"`
	Name        string `json:"addr_name" db:"addr_name"`
	Phone       string `json:"addr_phone" db:"addr_phone"`
	Email       string `json:"addr_email" db:"addr_email"`
	Address     string `json:"addr_state" db:"addr_state"`
	SubArea     string `json:"addr_subarea" db:"addr_subarea"`
	District    string `json:"addr_district" db:"addr_district"`
	Province    string `json:"addr_province" db:"addr_province"`
	PostalCode  int64  `json:"addr_postal_code" db:"addr_postal_code"`
	MainAddress int    `json:"main_address" db:"main_address"`
}

type OrderCompanyDetail struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	TexID       string `json:"tex_id"`
	PhoneMobile string `json:"phone_mobile"`
	PhoneHome   string `json:"phone_home"`
}

type CouponDetail struct {
	ID           int64     `json:"id"`
	DocDate      time.Time `json:"doc_date"`
	CouponNo     string    `json:"coupon_no"`
	CouponType   int64     `json:"coupon_type"`
	UserID       string    `json:"user_id"`
	Name         string    `json:"name"`
	Value        float64   `json:"value"`
	Remain       float64   `json:"remain"`
	ExpireStatus int64     `json:"expire_status"`
	BeginDate    time.Time `json:"begin_date"`
	ExpireDate   time.Time `json:"expire_date"`
}

type UseOrder struct {
	UserFname string `json:"user_fname"`
	UserLname string `json:"user_lname"`
}

// type PostResponse struct {
// 	Response Response `json:"response"`
// 	Message  string   `json:"message"`
// 	Status   bool     `json:"status"`
// }

type PostFail struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

// type Response struct {
// 	Items      EF582568151TH `json:"items"`
// 	TrackCount TrackCount    `json:"track_count"`
// }

// type EF582568151TH struct {
// 	EF582568151TH []ItemsPost `json:"EF582568151TH"`
// }
type ItemsPost struct {
	Barcode             string `json:"barcode"`
	Status              string `json:"status"`
	StatusDescription   string `json:"status_description"`
	StatusDate          string `json:"status_date"`
	Location            string `json:"location"`
	Postcode            string `json:"postcode"`
	DeliveryStatus      string `json:"delivery_status"`
	DeliveryDescription string `json:"delivery_description"`
	DeliveryDatetime    string `json:"delivery_datetime"`
	ReceiverName        string `json:"receiver_name"`
	Signature           string `json:"signature"`
}

type TrackCount struct {
	TrackDate       string `json:"track_date"`
	CountNumber     int    `json:"count_number"`
	TrackCountLimit int    `json:"track_count_limit"`
}

type OrderTrack struct {
	OrderStatus  int64    `json:"order_status"`
	DeliveryID   string   `json:"delivery_id"`
	DeliveryName string   `json:"delivery_name"`
	TrackingID   string   `json:"tracking_id"`
	SendText     []string `json:"send_text"`
}
