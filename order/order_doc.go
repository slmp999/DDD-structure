package order

/**
* @api {post} /order/v1/cart/storage  CartGetBasket
* @apiVersion 1.0.0
* @apiName CartGetBasket
* @apiGroup Order
* @apiDescription ค้นหาข้อมูล Basket  / coupon_enabled = 1 ได้/0 ไม่ได้
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 7388b778df45484e84d99530054a88aa
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {Data} data ข้อมูล
* @apiSuccessExample {json} Success:
* {
*    "response": "success",
*    "data": [
*        {
*            "id": 5,
*            "user_id": 62100900011,
*            "sale_id": 0,
*            "doc_no": "62101291121",
*            "item_amount": 0,
*            "discount_amount": 0,
*            "net_amount": 0,
*            "my_description": "",
*            "create_time": "2019-10-12T09:11:21+07:00",
*            "basket_sub": [
*                {
*                    "basket_id": 5,
*                    "cust_id": 62100900011,
*                    "sale_id": 0,
*                    "item_id": 2,
*                    "item_name": "PERFECT HAIR STRAIGHT NO.B (น้ำยาตัว B ยืด สูตรใหม่)",
*                    "qty": 2,
*                    "unit_id": 1,
*                    "unit_code": "ถุง",
*                    "wh_id": 1,
*                    "wh_code": "1",
*                    "price": 380,
*                    "discount_amount": 0,
*                    "item_amount": 0,
*                    "net_amount": 0,
*                    "line_number": 0,
*                    "pic_file_name_1": "https://perfectapi.extensionsoft.biz/image/product/product-196221318.png",
*					 "coupon_enabled": 1
*                }
*            ]
*        }
*    ]
* }
 */

/**
* @api {post} /order/v1/cart/all/delete  CartClearBasket
* @apiVersion 1.0.0
* @apiName CartClearBasket
* @apiGroup Order
* @apiDescription ล้างข้อมูลในตระกร้า
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 7388b778df45484e84d99530054a88aa
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccessExample {json} Success:
* {
*    "response": "success",
* }
 */

/**
* @api {post} /order/v1/cart/item/delete  CartDeleteItem
* @apiVersion 1.0.0
* @apiName CartDeleteItem
* @apiGroup Order
* @apiDescription ลบสินค้าในตระกร้า
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 7388b778df45484e84d99530054a88aa
* @apiParam (Parameter) {int} item_id IDสินค้า
* @apiParamExample {json} Body request:
*{
*	"item_id": 2
*}
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccessExample {json} Success:
* {
*    "response": "success",
* }
 */

/**
* @api {post} /order/v1/cart/update/qty  CartUpdateQTY
* @apiVersion 1.0.0
* @apiName CartUpdateQTY
* @apiGroup Order
* @apiDescription ปรับจำนวนสินค้า
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 7388b778df45484e84d99530054a88aa
* @apiParam (Parameter) {int} item_id IDสินค้า
* @apiParam (Parameter) {int} qty จำนวนสินค้า
* @apiParamExample {json} Body request:
*{
*	"item_id": 1,
*	"qty": 2
*}
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccessExample {json} Success:
* {
*    "response": "success",
* }
 */

/**
* @api {post} /order/v1/cart/add  CartAdd
* @apiVersion 1.0.0
* @apiName CartAdd
* @apiGroup Order
* @apiDescription สินค้าโชว์ AllPackage
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiParam (Parameter) {int} item_id IDสินค้า
* @apiParam (Parameter) {int} qty จำนวนสินค้า
* @apiParamExample {json} Body request:
*{
*	"item_id": 2,
*	"qty": 2
*}
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {Data} data ข้อมูล
* @apiSuccessExample {json} Success:
*{
*    "response": "success",
*    "data": {
*        "item_id": 2,
*        "qty": 2
*    }
*}
 */

/**
* @api {post} /order/v1/cart/transfer  TransferCartToOrder
* @apiVersion 1.0.0
* @apiName TransferCartToOrder
* @apiGroup Order
* @apiDescription สินค้าโชว์ AllPackage
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiParam (Parameter) {[]int} item_id IDสินค้า
* @apiParam (Parameter) {float} discount_amount ราคาส่วนลด
* @apiParam (Parameter) {float} distance_amount ค่าขนส่ง
* @apiParam (Parameter) {string} coupon คูปอง
* @apiParamExample {json} Body request:
*{
*	"item_id": [30],
*	"discount_amount": 0,
*	"dilivery_code": "01",
*	"distance_amount": 0,
*	"coupon": "CPV6210087259"
*}
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {Data} data ข้อมูล
* @apiSuccessExample {json} Success:
*{
*    "response": "success",
*    "data": {
*        "id": 5
*    }
*}
 */

/**
* @api {post} /order/v1/order/confirm  OrderConfirm
* @apiVersion 1.0.0
* @apiName OrderConfirm
* @apiGroup Order
* @apiDescription สินค้าโชว์ AllPackage
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiParam (Parameter) {String} id รหัสสินค้า
* @apiParam (Parameter) {url} id รูปภาพ
* @apiParam (Parameter) {address_id} id ทีอยู่ลูกค้า
* @apiParamExample {json} Body request:
*{
*	"id": 2,
*	"url": "http://perfectapi.extensionsoft.biz/image/product/product-325705219.png",
*	"address_id": 20
*}
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} Message ข้อความตอบกลับ
* @apiSuccessExample {json} Success:
*{
*    "response": "success",
*    "message": "การ confirm เรียบร้อย",
*    "data": {
*        "address_id": 20,
*        "order_id": 3,
*        "url": "http://perfectapi.extensionsoft.biz/image/product/product-325705219.png"
*    }
*}
 */

/**
* @api {post} /order/v1/status/update OrderUpdateStatus
* @apiVersion 1.0.0
* @apiName OrderUpdateStatus
* @apiGroup Order
* @apiDescription สินค้าโชว์ AllPackage
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiParam (Parameter) {String} id รหัสสินค้า
* @apiParam (Parameter) {url} status สถานะ order ที่จต้องการ่ง
* @apiParam (Parameter) {String} tracking_id code ตรวจสอบการส่งของ
* @apiParamExample {json} Body request:
*{
*	"id": 20,
*	"status": 2,
*	"tracking_id":"A2532"
*}
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccessExample {json} Success:
*{
*    "response": "success"
*}
 */

/**
* @api {post} /order/v1/promptpay  Promptpay
* @apiVersion 1.0.0
* @apiName Promptpay
* @apiGroup Order
* @apiDescription สินค้าโชว์ AllPackage
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiParam (Parameter) {float} amount ค่าเงิน
* @apiParamExample {json} Body request:
*{
*	"amount": 10
*}
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} promptpay tag29
* @apiSuccessExample {json} Success:
*{
*    "response": "success",
*    "promptpay": "00020101021129370016A000000677010111011300669391464745802TH530376454041.00630490D3"
*}
 */

/**
* @api {post} /order/v1/delivery/price  DeliveryPrice
* @apiVersion 1.0.0
* @apiName DeliveryPrice
* @apiGroup Order
* @apiDescription สินค้าโชว์ AllPackage
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiParam (Parameter) {[]int} item_id ค่าเงิน
* @apiParam (Parameter) {String} code ค่าเงิน
* @apiParamExample {json} Body request:
*{
*	"item_id": [7,6,1,2],
*	"code": "01"
*}
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} name ชื่อบริษัทขนส่ง
* @apiSuccess {float} amount จำนวนเงิน
* @apiSuccessExample {json} Success:
*{
*    "response": "success",
*    "name": "Flash Express",
*    "amount": 0
*}
 */

/**
* @api {post} /order/v1/buy/package  OrderBuyPackage
* @apiVersion 1.0.0
* @apiName OrderBuyPackage
* @apiGroup Order
* @apiDescription สินค้าโชว์ AllPackage
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiParam (Parameter) {int} item_id รหัสสินค้า
* @apiParamExample {json} Body request:
*{
*	"item_id": 30
*}
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} promptpay tag29
* @apiSuccessExample {json} Success:
*{
*    "response": "success",
*    "data": {
*        "id": 4
*    }
*}
 */

/**
* @api {post} /order/v1/list/all  OrderListAll
* @apiVersion 1.0.0
* @apiName OrderListAll
* @apiGroup Order
* @apiDescription ค้นหาออเดอร์ตามสถานะ
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiParam (Parameter) {int} order_status สถานะออเดอร์
* @apiParamExample {json} Body request:
* {
* 	"order_status": 1
* }
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} data ข้อมูล
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "data": [
* 	{
* 		"id": 1,
* 		"user_id": "62100900002",
* 		"user_name": "62100900002",
* 		"sale_id": 0,
* 		"sale_name": "",
* 		"doc_no": "62100933830",
* 		"sum_of_item_amount": 760,
* 		"discount_amount": 810,
* 		"after_discount_amount": 0,
* 		"before_tax": 0,
* 		"tax_amount": 0,
* 		"total_amount": 0,
* 		"sale_type": 0,
* 		"sum_cash_amount": 0,
* 		"sum_credit_amount": 0,
* 		"sum_deposit_amount": 0,
* 		"sum_coupon_amount": 0,
* 		"sum_bank_amount": 0,
* 		"change_amount": 0,
* 		"net_debt_amount": 0,
* 		"my_description": "",
* 		"order_status": 1,
* 		"delivery_date": "2019-09-03T18:23:44+07:00",
* 		"distance": 0,
* 		"distance_amount": 50,
* 		"delivery_link": "",
* 		"delivery_id": "0",
* 		"referral_id": "0",
* 		"pic_slip": "",
* 		"send_text": [
* 			"สร้างใหม่"
* 		],
* 		"create_time": "2019-10-09T03:38:30+07:00",
* 		"tracking_id": "",
* 		"is_package": 0,
* 		"OrderSub": [
* 			{
* 				"id": 1,
* 				"order_id": 1,
* 				"item_id": 2,
* 				"item_name": "PERFECT HAIR STRAIGHT NO.B (น้ำยาตัว B ยืด สูตรใหม่)",
* 				"wh_id": 1,
* 				"wh_code": "1",
* 				"shelf_id": 0,
* 				"shelf_code": "",
* 				"qty": 1,
* 				"cn_qty": 0,
* 				"unit_id": 1,
* 				"unit_code": "ถุง",
* 				"price": 380,
* 				"discount_amount": 0,
* 				"item_amount": 0,
* 				"average_cost": 0,
* 				"sum_of_cost": 0,
* 				"rate_1": 0,
* 				"stock_type": 0,
* 				"basket_id": 2,
* 				"item_sub_cat": "0",
* 				"item_group": "0",
* 				"type_code": "0",
* 				"point": 0,
* 				"is_cancel": 0,
* 				"ref_line_number": 0,
* 				"pic_file_name_1": "https://perfectapi.extensionsoft.biz/image/product/product-196221318.png",
* 				"line_number": 0
* 			}
* 		]
* 	},
* }
 */

/**
* @api {post} /order/v1/send/detail OrderSendDetail
* @apiVersion 1.0.0
* @apiName OrderSendDetail
* @apiGroup Order
* @apiDescription ค้นหาออเดอร์ตามสถานะ
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiParam (Parameter) {int} order_status สถานะออเดอร์
* @apiParamExample {json} Body request:
* {
* 	"id": 1
* }
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} data ข้อมูล
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "data": [
*         {
*             "id": 49,
*             "company_id": 1,
*             "user_id": "62100900011",
*             "user_name": "tanakorn phuntulap",
*             "sale_id": 0,
*             "sale_name": "",
*             "doc_no": "ORD621112122038",
*             "sum_of_item_amount": 380,
*             "discount_amount": 0,
*             "after_discount_amount": 470,
*             "before_tax": 0,
*             "tax_amount": 0,
*             "total_amount": 470,
*             "sale_type": 0,
*             "sum_cash_amount": 0,
*             "sum_credit_amount": 0,
*             "sum_deposit_amount": 0,
*             "sum_coupon_amount": 0,
*             "sum_bank_amount": 470,
*             "change_amount": 0,
*             "net_debt_amount": 0,
*             "my_description": "",
*             "order_status": 1,
*             "delivery_date": "2019-11-12T05:20:38Z",
*             "distance": 0,
*             "distance_amount": 90,
*             "delivery_link": "",
*             "delivery_id": "02",
*             "delivery_name": "Nim Express",
*             "referral_id": "0",
*             "pic_slip": "https://api-dev.pct2003.com/image/bill/bill-911830497.png",
*             "send_text": [
*                 "รอยืนยัน"
*             ],
*             "create_time": "2019-11-12T05:20:38Z",
*             "tracking_id": "",
*             "address_id": 36,
*             "is_package": 0,
*             "OrderSub": [
*                 {
*                     "id": 124,
*                     "order_id": 49,
*                     "item_id": 1,
*                     "item_name": "PERFECT HAIR STRAIGHT & PERM NO.A (น้ำยาตัว A สูตรใหม่)",
*                     "wh_id": 1,
*                     "wh_code": "1",
*                     "shelf_id": 0,
*                     "shelf_code": "",
*                     "qty": 1,
*                     "cn_qty": 0,
*                     "unit_id": 1,
*                     "unit_code": "ถุง",
*                     "price": 380,
*                     "discount_amount": 0,
*                     "item_amount": 0,
*                     "average_cost": 0,
*                     "sum_of_cost": 0,
*                     "rate_1": 0,
*                     "stock_type": 0,
*                     "basket_id": 8,
*                     "item_sub_cat": "0",
*                     "item_group": "0",
*                     "type_code": "0",
*                     "point": 0,
*                     "is_cancel": 0,
*                     "ref_line_number": 0,
*                     "pic_file_name_1": "https://api.pct2003.com/image/product/product-325705219.png",
*                     "line_number": 0
*                 }
*             ],
*             "OrderSendAddress": {
*                 "addr_id": 36,
*                 "addr_name": "tanakorn",
*                 "addr_phone": "0861803046",
*                 "addr_email": "destiny00005@gmail.com",
*                 "addr_state": "tanakorn",
*                 "addr_subarea": "เด่นใหญ่",
*                 "addr_district": "หันคา",
*                 "addr_province": "ชัยนาท",
*                 "addr_postal_code": 17130,
*                 "main_address": 1
*             },
*             "OrderCompanyDetail": {
*                 "id": 1,
*                 "name": "บริษัท เพอร์เฟคท์ แชมป์2003 (ไทยแลนด์) จำกัด (สำนักงานใหญ่)",
*                 "address": "221 หมู่ 6 ต.ฟ้าฮ่าม อ.เมือง จ.เชียงใหม่ 50000",
*                 "tex_id": "0505560006519",
*                 "phone_mobile": "0613414888",
*                 "phone_home": "052005539"
*             }
*         }
*     ]
* }
 */

/**
* @api {post}  /order/v1/order/id OrderById
* @apiVersion 1.0.0
* @apiName OrderById
* @apiGroup Order
* @apiDescription ค้นหาออเดอร์ตามID
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiParam (Parameter) {int} id รหัสสินค้า
* @apiParamExample {json} Body request:
* {
* 	"id": 1
* }
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} data ข้อมูล
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "data": [
*         {
*             "id": 2,
*             "user_id": "62100900003",
*             "user_name": "62100900003",
*             "sale_id": 0,
*             "sale_name": "",
*             "doc_no": "621009101237",
*             "sum_of_item_amount": 760,
*             "discount_amount": 810,
*             "after_discount_amount": 0,
*             "before_tax": 0,
*             "tax_amount": 0,
*             "total_amount": 0,
*             "sale_type": 0,
*             "sum_cash_amount": 0,
*             "sum_credit_amount": 0,
*             "sum_deposit_amount": 0,
*             "sum_coupon_amount": 0,
*             "sum_bank_amount": 0,
*             "change_amount": 0,
*             "net_debt_amount": 0,
*             "my_description": "",
*             "order_status": 99,
*             "delivery_date": "2019-09-03T18:23:44+07:00",
*             "distance": 0,
*             "distance_amount": 50,
*             "delivery_link": "",
*             "delivery_id": "0",
*             "referral_id": "0",
*             "pic_slip": "https://perfectapi.extensionsoft.biz/image/bill/bill-188056024.png",
*             "send_text": [
*                 "ยกเลิก order"
*             ],
*             "create_time": "2019-10-09T10:12:38+07:00",
*             "tracking_id": "",
*             "is_package": 0,
*             "OrderSub": [
*                 {
*                     "id": 3,
*                     "order_id": 2,
*                     "item_id": 1,
*                     "item_name": "PERFECT HAIR STRAIGHT & PERM NO.A (น้ำยาตัว A สูตรใหม่)",
*                     "wh_id": 1,
*                     "wh_code": "1",
*                     "shelf_id": 0,
*                     "shelf_code": "",
*                     "qty": 2,
*                     "cn_qty": 0,
*                     "unit_id": 1,
*                     "unit_code": "ถุง",
*                     "price": 380,
*                     "discount_amount": 0,
*                     "item_amount": 0,
*                     "average_cost": 0,
*                     "sum_of_cost": 0,
*                     "rate_1": 0,
*                     "stock_type": 0,
*                     "basket_id": 3,
*                     "item_sub_cat": "0",
*                     "item_group": "0",
*                     "type_code": "0",
*                     "point": 0,
*                     "is_cancel": 0,
*                     "ref_line_number": 0,
*                     "pic_file_name_1": "https://perfectapi.extensionsoft.biz/image/product/product-325705219.png",
*                     "line_number": 0
*                 }
*             ]
*         }
*     ]
* }
 */

/**
* @api {post}  /order/v1/order/cancel OrderCancel
* @apiVersion 1.0.0
* @apiName OrderCancel
* @apiGroup Order
* @apiDescription ยกเลิกออเดอร์
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiParam (Parameter) {int} id รหัสสินค้า
* @apiParamExample {json} Body request:
* {
* 	"id": 1
* }
* @apiSuccess {String} response สถานะตอบกลับล
* @apiSuccessExample {json} Success:
*{
*    "response": "success"
*}
 */

/**
* @api {post}  /order/v1/bank Bank
* @apiVersion 1.0.0
* @apiName Bank
* @apiGroup Order
* @apiDescription ค้าหาธนาคารผู้ประกอบการ
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiSuccess {String} response สถานะตอบกลับล
* @apiSuccess {String} data ข้อมูล
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "data": [
*         {
*             "id": 1,
*             "name": "ปทิตตา  อุ่นเตียม",
*             "code": "9860586527",
*             "bank_code": "006",
*             "bank_name": "ธ. กรุงไทย จำกัด (มหาชน)",
*             "code_name": "KTB",
*             "image": "http://perfectapi.extensionsoft.biz/image/product/product-834525452.png"
*         },
*         {
*             "id": 2,
*             "name": "อุดร  อุ่นเตียม",
*             "code": "0939146474",
*             "bank_code": "004",
*             "bank_name": "ธ. กสิกรไทย จำกัด (มหาชน)",
*             "code_name": "KBANK",
*             "image": "http://perfectapi.extensionsoft.biz/image/product/product-436002337.png"
*         }
*     ]
* }
 */

/**
* @api {post}  /order/v1/bank/list BankList
* @apiVersion 1.0.0
* @apiName BankList
* @apiGroup Order
* @apiDescription ค้าหาธนาคาร
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 7388b778df45484e84d99530054a88aa
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} message แจ้ง error
* @apiSuccess {String} data ข้อมูล
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "message": "",
*     "data": [
*         {
*             "id": 2,
*             "bank_code": "001",
*             "bank_name": "ธนาคารแห่งประเทศไทย",
*             "code": "ธนาคารแห่งประเทศไทย",
*             "description": "",
*             "image": ""
* 		},
* 		{....}
* 	]
* }
 */

/**
* @api {post}  /order/v1/delivery/all DeliveryAll
* @apiVersion 1.0.0
* @apiName DeliveryAll
* @apiGroup Order
* @apiDescription บริษัทขนส่ง
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} data ข้อมูล
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "data": [
*         {
*             "code": "01",
*             "name": "Flash Express"
*         },
*         {
*             "code": "02",
*             "name": "Nim Express"
*         },
*         {
*             "code": "03",
*             "name": "Kerry Express"
*         },
*         {
*             "code": "04",
*             "name": "Thai Post Express"
*         }
*     ]
* }
 */

/**
* @api {post}  /order/v1/link/tracking Checktracking
* @apiVersion 1.0.0
* @apiName CheckTracking
* @apiGroup Order
* @apiDescription Checktracking
* @apiHeader Content-Type application/json
* @apiParam (Parameter) {string} doc_no รหัส order
* @apiParamExample {json} Body request:
* {
* 	"doc_no": "ORD621117224740"
* }
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} data ข้อมูล
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "data": {
*         "order_status": 5,
*         "delivery_id": "01",
*         "delivery_name": "แฟลช เอ็กซ์เพรส",
*         "tracking_id": "89495859606958",
*         "send_text": [
*             "รอยืนยัน",
*             "ชำระแล้ว",
*             "จัดของ",
*             "กำลังขนส่งโดย นิ่ม เอ็กซ์เพรส ติดตามได้ที่ https://www.nimexpress.com/web/p/tracking?i=EY145587896TH",
*             "สำเร็จ"
*         ]
*     }
* }
 */
