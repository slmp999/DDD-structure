package product

/**
 * @api {post} /product/v1/item  FindAllProduct
 * @apiVersion 1.0.0
 * @apiName FindAllProduct
 * @apiGroup Product
 * @apiDescription ค้นหาสินค้าทั้งหมด
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} data ข้อมูล
 * @apiSuccessExample {json} Success:
*{
*    "response": "success",
*    "data": [
*        {
*            "id": 1,
*            "code": "PCT00001",
*            "name": "PERFECT HAIR STRAIGHT & PERM NO.A (น้ำยาตัว A สูตรใหม่)",
*            "eng_name": "PERFECT HAIR STRAIGHT & PERM NO.A ",
*            "short_name": "",
*            "category_code": "01",
*            "type_code": "",
*            "my_description": "คำอธิบาย1",
*            "unit_id": 1,
*            "unit_code": "ถุง",
*            "def_sale_wh_id": 1,
*            "def_sale_wh_code": "1",
*            "price": 380,
*            "my_grade": "",
*            "is_packge": 0,
*            "degree": "4",
*            "pic_file_name_1": "https://perfectapi.extensionsoft.biz/image/product/product-325705219.png",
*            "category_name": "ผลิตภัณฑ์เกี่ยวกับเส้นผม",
*            "category_name_eng": "salon group",
*            "sale_price_1": 380,
*            "sale_price_2": 250,
*            "payment_destination": 0
*        },
*        {
* 			.
* 			.
* 			.
* 		 }
*    ]
*}
*/

/**
 * @api {post} /product/v1/item/category  FindItemCategoryByCode
 * @apiVersion 1.0.0
 * @apiName FindItemCategoryByCode
 * @apiGroup Product
 * @apiDescription ค้นหาสินค้าอ้างอิงตาม Category ทั้งหมด
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} data ข้อมูล
 * @apiSuccessExample {json} Success:
*{
*    "response": "success",
*    "data": [
*        {
*            "id": 1,
*            "code": "PCT00001",
*            "name": "PERFECT HAIR STRAIGHT & PERM NO.A (น้ำยาตัว A สูตรใหม่)",
*            "eng_name": "PERFECT HAIR STRAIGHT & PERM NO.A ",
*            "short_name": "",
*            "category_code": "01",
*            "type_code": "",
*            "my_description": "คำอธิบาย1",
*            "unit_id": 1,
*            "unit_code": "ถุง",
*            "def_sale_wh_id": 1,
*            "def_sale_wh_code": "1",
*            "price": 380,
*            "my_grade": "",
*            "is_packge": 0,
*            "degree": "4",
*            "pic_file_name_1": "https://perfectapi.extensionsoft.biz/image/product/product-325705219.png",
*            "category_name": "ผลิตภัณฑ์เกี่ยวกับเส้นผม",
*            "category_name_eng": "salon group",
*            "sale_price_1": 380,
*            "sale_price_2": 250,
*            "payment_destination": 0
*        },
*        {
* 			.
* 			.
* 			.
* 		 }
*    ]
*}
*/

/**
 * @api {post} /product/v1/category  FindCategory
 * @apiVersion 1.0.0
 * @apiName FindCategory
 * @apiGroup Product
 * @apiDescription ค้นหา Category ทั้งหมด
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} data ข้อมูล
 * @apiSuccessExample {json} Success:
*{
*    "response": "success",
*    "data": [
*        {
*            "id": 1,
*            "company_id": 0,
*            "code": "01",
*            "name": "ผลิตภัณฑ์เกี่ยวกับเส้นผม",
*            "eng_name": "salon group",
*            "item_count": 1,
*            "my_description": ""
*        },
*        {
*            "id": 2,
*            "company_id": 0,
*            "code": "02",
*            "name": "ผลิตภัณฑ์ความสวยความงาม",
*            "eng_name": "beauty group",
*            "item_count": 1,
*            "my_description": ""
*        }
*    ]
*}
*/

/**
* @api {post} /product/v1/item/id  FindItemByID
* @apiVersion 1.0.0
* @apiName FindItemByID
* @apiGroup Product
* @apiDescription ค้นหาสินค้าตาม รหัสสินค้า ทั้งหมด
* @apiHeader Content-Type application/json
* @apiParam (Parameter) {int} id รหัสสินค้า
* @apiParamExample {json} Body request:
* {
* 	"id": 1
*	}
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} data ข้อมูล
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "data": [
*         {
*             "id": 69,
*             "code": "PCT254251",
*             "name": "เทสส1",
*             "eng_name": "test1",
*             "short_name": "tt",
*             "group_code": "",
*             "type_code": "",
*             "my_description": "description",
*             "stock_type": 0,
*             "unit_id": 1,
*             "unit_code": "ถุง",
*             "price": 1200,
*             "buy_unit_id": 0,
*             "buy_unit_code": "",
*             "sale_unit_id": 1,
*             "sale_unit_code": "ถุง",
*             "weight": 12,
*             "hight": 0,
*             "wide": 0,
*             "category_code": "02",
*             "sub_cat_code": "00",
*             "my_grade": "",
*             "brand_code": "",
*             "color_code": "",
*             "my_class": "",
*             "unit_type": 0,
*             "delivery_charge": 0,
*             "install_charge": 0,
*             "return_status": 0,
*             "item_status": 0,
*             "last_price": 0,
*             "average_cost": 0,
*             "cost_type": 0,
*             "stock_qty": 0,
*             "stock_min": 0,
*             "stock_max": 0,
*             "pic_file_name_1": "https://mainpic",
*             "pic_file_name_2": "",
*             "pic_file_name_3": "",
*             "pic_file_name_4": "",
*             "pic_file_name_5": "",
*             "def_sale_wh_id": 0,
*             "def_sale_wh_code": "",
*             "def_sale_shelf_id": 0,
*             "def_sale_shelf_code": "",
*             "def_buy_wh_id": 0,
*             "def_buy_wh_code": "",
*             "def_buy_shelf_id": 0,
*             "def_buy_shelf_code": "",
*             "use_package": 0,
*             "is_package": 0,
*             "favorite_promo": 0,
*             "degree": 1,
*             "package_discount": 0,
*             "coupon_amount": 0,
*             "campaign_code": "",
*             "is_promotion": 0,
*             "active_status": 1,
*             "category_name": "ผลิตภัณฑ์ความสวยความงาม",
*             "category_name_eng": "beauty",
*             "sale_price_1": 0,
*             "sale_price_2": 0,
*             "begin_date": "2019-10-31",
*             "end_date": "2019-10-08",
*             "is_expire_date": 0,
*             "payment_destination": 0,
*             "is_saler": 1,
*             "pic_file_saler": "sss",
*             "ItemPromoSub": [
*                  {
*                      "id": 8,
*                      "parent_code": "PMT00008",
*                      "item_id": 1,
*                      "code": "PCT00001",
*		               "name": "เพอเฟคท์ แฮร์เสตร็ด No.A (น้ำยาตัว A สูตรใหม่)",
*                      "qty": 1,
*                      "price": 200,
*                      "amount": 200,
*                      "unit_id": 0,
*                      "unit_code": "",
*                      "line_number": 0
*                  }
*             ],
*             "PicSub": [
*                 {
*                     "id": 7,
*                     "url": "https://testadd"
*                 },
*                 {
*                     "id": 8,
*                     "url": "https://test11"
*                 }
*             ]
*         }
*     ]
* }
 */

/**
 * @api {post} /product/v1/item/favorite  FineFavoriteItem
 * @apiVersion 1.0.0
 * @apiName FavoriteItem
 * @apiGroup Product
 * @apiDescription สินค้าโชว์หน้าแรก
 * @apiHeader Content-Type application/json
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
*{
*    "response": "success",
*    "data": [
*        {
*            "id": 2,
*            "code": "PCT00002",
*            "name": "PERFECT HAIR STRAIGHT NO.B (น้ำยาตัว B ยืด สูตรใหม่)",
*            "eng_name": "PERFECT HAIR STRAIGHT NO.B",
*            "short_name": "",
*            "category_code": "01",
*            "type_code": "",
*            "my_description": "คำอธิบาย2",
*            "unit_id": 1,
*            "unit_code": "ถุง",
*            "def_sale_wh_id": 1,
*            "def_sale_wh_code": "1",
*            "price": 380,
*            "my_grade": "",
*            "degree": "5",
*            "pic_file_name_1": "https://perfectapi.extensionsoft.biz/image/product/product-196221318.png",
*            "category_name": "ผลิตภัณฑ์เกี่ยวกับเส้นผม",
*            "category_name_eng": "salon group",
*            "sale_price_1": 380,
*            "sale_price_2": 250,
*            "payment_destination": 0
*        },
*        {
* 			.
* 			.
* 			.
* 		 }
*    ]
*}
*/

/**
 * @api {post} /product/v1/package  AllPackage
 * @apiVersion 1.0.0
 * @apiName AllPackage
 * @apiGroup Product
 * @apiDescription สินค้าโชว์ AllPackage
 * @apiHeader Content-Type application/json
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 *{
 *    "response": "success",
 *    "data": [
 *        {
 *           "id": 30,
 *           "code": "PKS00001",
 *           "name": "Package Size S",
 *           "eng_name": "Package Size S",
 *           "short_name": "",
 *           "category_code": "04",
 *           "type_code": "",
 *           "my_description": "",
 *           "unit_id": 4,
 *           "unit_code": "package",
 *           "def_sale_wh_id": 1,
 *           "def_sale_wh_code": "1",
 *           "price": 10000,
 *           "my_grade": "",
 *           "is_packge": 1,
 *           "degree": 5,
 *           "pic_file_name_1": "https://perfectapi.extensionsoft.biz/image/bill/bill-644382371.png",
 *           "category_name": "แพ็คเกจ",
 *           "category_name_eng": "distribute",
 *           "sale_price_1": 10000,
 *           "sale_price_2": 10000,
 *           "payment_destination": 0
 *       },
 *       {
 *			.
 *			.
 *			.
 *		 }
 *    ]
 *}
 */

/**
 * @api {post} /product/v1/add/category  AddCategory
 * @apiVersion 1.0.0
 * @apiName AddCategory
 * @apiGroup ProductBackend
 * @apiDescription เพิ่มหมวดสินค้า
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
 * @apiParam (Parameter) {int} id รหัสสินค้า
 * @apiParam (Parameter) {int} company_id รหัสบริษัท
 * @apiParam (Parameter) {String} name ชื่อสินค้า
 * @apiParam (Parameter) {String} eng_name ชื่อภาษาอังกฤษ
 * @apiParam (Parameter) {String} my_description รายละเอียด
 * @apiParamExample {json} Body request:
 *{
 *	"id": 0,
 *	"company_id": 1,
 *	"name": "เทส",
 *	"eng_name": "test1",
 *	"active_status": 1,
 *	"my_description": "sasx"
 *}
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 *{
 *    "response": "success"
 *}
 */

/**
 * @api {post} /product/v1/youtube  FineYoutube
 * @apiVersion 1.0.0
 * @apiName Youtube
 * @apiGroup Product
 * @apiDescription Youtube list
 * @apiHeader Content-Type application/json
 * @apiParam (Parameter) {int} id ลำดับ
 * @apiParam (Parameter) {int} link_address ที่อยู่ลิงค์
 * @apiParam (Parameter) {String} my_description คำอธิบาย
 * @apiParamExample {json} Body request:
 *{
 *    "response": "success",
 *    "data": [
 *        {
 *            "id": 1,
 *            "link_address": "https://www.youtube.com/watch?v=nXGg2bxy7i4",
 *            "my_description": "open web"
 *        },
 *        {
 *            "id": 2,
 *            "link_address": "https://www.youtube.com/watch?v=mVJdAcR6qLQ",
 *            "my_description": "open web"
 *        }
 *    ]
 *}
 */

/**
 * @api {post} /product/v1/campaign  FineCampaign
 * @apiVersion 1.0.0
 * @apiName Campaign
 * @apiGroup Product
 * @apiDescription Campaign list
 * @apiHeader Content-Type application/json
 * @apiParam (Parameter) {int} id ลำดับ
 * @apiParam (Parameter) {data} data ข้อมูล
 * @apiParamExample {json} Body request:
 *{
 *    "response": "success",
 *    "data": [
 *        {
 *            "id": 1,
 *            "code": "CP6210001",
 *            "my_description": "ssss",
 *            "pic_file_name_1": "https://perfectapi.extensionsoft.biz/image/bill/bill-015762854.png",
 *            "begin_date": "2019-10-01",
 *            "end_date": "2019-12-31"
 *		},
 *		{
 *            "id": 2,
 *            "code": "CP6210002",
 *            "my_description": "aaaa",
 *            "pic_file_name_1": "https://perfectapi.extensionsoft.biz/image/bill/bill-015762854.png",
 *            "begin_date": "2019-10-01",
 *            "end_date": "2019-12-31"
 *        }
 *    ]
 *}
 */

/**
 * @api {post} /product/v1/promotion/favorite  FineFavoritePromotion
 * @apiVersion 1.0.0
 * @apiName FineFavoritePromotion
 * @apiGroup Product
 * @apiDescription FineFavoritePromotion list
 * @apiHeader Content-Type application/json
 * @apiParam (Parameter) {int} id รหัสสินค้า
 * @apiParam (Parameter) {data} data ข้อมูล
 * @apiParamExample {json} Body request:
{
    "response": "success",
    "data": [
        {
            "id": 41,
            "code": "PMT00001",
            "name": "โปรโมชั่น Perfect Gold 1 ขวด เหลือ 200 บาท",
            "eng_name": "",
            "short_name": "",
            "category_code": "05",
            "type_code": "",
            "my_description": "",
            "unit_id": 3,
            "unit_code": "ขวด",
            "def_sale_wh_id": 1,
            "def_sale_wh_code": "1",
            "price": 200,
            "my_grade": "",
            "use_package": 0,
            "is_package": 0,
            "stock_type": 0,
            "degree": 5,
            "pic_file_name_1": "https://api.pct2003.com/image/bill/bill-707470248.png",
            "category_name": "โปรโมชั่น",
            "category_name_eng": "promotion",
            "sale_price_1": 200,
            "sale_price_2": 200,
            "begin_date": "",
            "end_date": "",
            "is_expire_date": 0,
            "payment_destination": 0,
            "ItemPromoSub": null
		},
		{
			.
			.
			.
		}
    ]
}
*/

/**
 * @api {post} /product/v1/add/item  AddItem
 * @apiVersion 1.0.0
 * @apiName AddItem
 * @apiGroup ProductBackend
 * @apiDescription id == 0 เพิ่มสินค้า id != 0 อัพเดตตาม id // PicSub id จะใช้ในการอัพเดต ถ้า id == 0 เพิ่มรูป และ id == 1 แก้ไขรูป
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
 * @apiParam (Parameter) {int} id รหัสสินค้า
 * @apiParamExample {json} Body request:
 * {
 * 	"id":0,
 * 	"code": "PCT254251",
 * 	"name": "เทสส1",
 * 	"eng_name": "test1",
 * 	"short_name": "tt",
 * 	"my_description": "description",
 * 	"unit_id": 1,
 * 	"unit_code": "ถุง",
 * 	"price": 1200,
 * 	"weight": 12,
 * 	"category_code": "10",
 * 	"last_price":0,
 * 	"pic_file_name_1": "https://mainpic",
 * 	"PicSub":[
 * 		{
 * 		"id":0,
 * 		"url":"https://testadd"
 * 		},
 * 		{
 * 		"id":1,
 * 		"url":"https://test11"
 * 		},
 * 		{
 * 		"id":2,
 * 		"url":"https://test2"
 * 		}
 * 	],
 * 	"sale_price_1":0,
 * 	"sale_price_2":0,
 * 	"degree":1,
 * 	"favorite":1,
 * 	"is_saler": 1,
 * 	"pic_file_saler": "sss",
 * 	"begin_date": "2019-10-31",
 * 	"end_date": "2019-10-08",
 * 	"per_discount": 3,
 * 	"active_status":1
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *    "response": "success"
 * }
 */

/**
 * @api {post} /product/v1/item/list  FindItemList
 * @apiVersion 1.0.0
 * @apiName FindItemList
 * @apiGroup ProductBackend
 * @apiDescription ค้นหาสินค้า  id 1 = สินค้าทั่วไป,id 2 = สินค้าโปร,id 3 = สินค้าแพคเกจ
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
 * @apiParam (Parameter) {int} id ประเภทสินค้า
 * @apiParamExample {json} Body request:
 * {
 *	"id":1
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccessExample {json} Success:
 *  {
 *      "response": "success",
 *      "data": [
 *          {
 *              "id": 41,
 *              "code": "PMT00001",
 *              "name": "โปรโมชั่น Perfect Gold 1 ขวด เหลือ 200 บาท",
 *              "eng_name": "",
 *              "category_code": "05",
 *              "price": 200,
 *              "sale_price_1": 200,
 *              "sale_price_2": 200,
 *              "pic_file_name_1": "https://api.pct2003.com/image/bill/bill-707470248.png",
 *              "category_name": "โปรโมชั่น",
 *              "category_name_eng": "promotion"
 *          },
 *          {....}
 *      ]
 *  }
 */

/**
 * @api {post} /product/v1/add/promotion  AddPromotion
 * @apiVersion 1.0.0
 * @apiName AddPromotion
 * @apiGroup ProductBackend
 * @apiDescription เพิ่ม item promotion id == 0 เพิ่มสินค้า id != 0 อัพเดตตาม id // PicSub id จะใช้ในการอัพเดต ถ้า id == 0 เพิ่มรูป และ id == 1 แก้ไขรูป  // ItemPromoSub  ถ้า id == 0 เพิ่ม และ id !=0 แก้ไขรูปตาม id
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
 * @apiParam (Parameter) {int} id รหัสสินค้า
 * @apiParamExample {json} Body request:
 * {
 * 	"id":0,
 * 	"code": "PCT254251",
 * 	"name": "promoTest",
 * 	"eng_name": "promoTest",
 * 	"short_name": "pt",
 * 	"my_description": "promoTestpromoTestpromoTest",
 * 	"unit_id": 1,
 * 	"unit_code": "ถุง",
 * 	"price": 1200,
 * 	"sale_price_1":1200,
 * 	"sale_price_2":800,
 * 	"weight": 2,
 * 	"pic_file_name_1": "",
 * 	"PicSub":[
 * 		{
 * 		"id":0,
 * 		"url":"https://test11"
 * 		}
 * 	],
 * 	"degree":4,
 * 	"favorite_promo":1,
 * 	"begin_date": "2019-10-31",
 * 	"end_date": "2019-10-08",
 * 	"active_status":1,
 * 	"ItemPromoSub":[
 * 		{
 *		 "id": 0,
 *		 "item_id": 1,
 *		 "code": "PCT00001",
 *		 "name": "เพอเฟคท์ แฮร์เสตร็ด No.A (น้ำยาตัว A สูตรใหม่)",
 *		 "qty": 1,
 *		 "price": 380,
 *		 "unit_id": 1,
 *		 "unit_code": "บาน",
 *		 "line_number": 1
 *		}
 * 	]
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *    "response": "success"
 * }
 */

/**
 * @api {post} /product/v1/add/package  AddPackage
 * @apiVersion 1.0.0
 * @apiName AddPackage
 * @apiGroup ProductBackend
 * @apiDescription เพิ่ม item Package id == 0 เพิ่มสินค้า id != 0 อัพเดตตาม id // PicSub id จะใช้ในการอัพเดต ถ้า id == 0 เพิ่มรูป และ id == 1 แก้ไขรูป
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
 * @apiParam (Parameter) {int} id รหัสสินค้า
 * @apiParamExample {json} Body request:
 * {
 * {
 * 	"id":0,
 * 	"code": "PCT254251",
 * 	"name": "เทสส2",
 * 	"eng_name": "test2",
 * 	"short_name": "tt",
 * 	"my_description": "description",
 * 	"price": 1,
 * 	"last_price":1,
 * 	"pic_file_name_1": "",
 * 	"PicSub":[
 * 		{
 * 		"id":0,
 * 		"url":"https://test11"
 * 		}
 * 	],
 * 	"sale_price_1":12000,
 * 	"sale_price_2":12000,
 * 	"package_discount":100,
 * 	"coupon_amount":100,
 * 	"active_status":1
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *    "response": "success"
 * }
 */

/**
 * @api {post} /product/v1/item/unitlist  UnitList
 * @apiVersion 1.0.0
 * @apiName UnitList
 * @apiGroup ProductBackend
 * @apiDescription UnitList
 * @apiHeader Content-Type application/json
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} data ข้อมูล
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "data": [
 *         {
 *             "id": 1,
 *             "company_id": 1,
 *             "code": "ถุง",
 *             "name": "ถุง",
 *             "rate_1": 1
 *         },
 *         {....}
 *     ]
 * }
 */

/**
 * @api {post} /product/v1/item/history  ItemHistory
 * @apiVersion 1.0.0
 * @apiName ItemHistory
 * @apiGroup ProductBackend
 * @apiDescription ItemHistory
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e
 * @apiParam (Parameter) {int} id รหัสสินค้า
 * @apiParamExample {json} Body request:
 * {
 *	"id":1
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} data ข้อมูล
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "data": [
 *         {
 *             "sum_all_amount": 6840,
 *             "sum_qty": 18,
 *             "ItemHistoryDetail": [
 *                 {
 *                     "date": "2019-11-06 15:37:06",
 *                     "name": "สมรถ หลักฐาน",
 *                     "cod_no": "ORD621106223705",
 *                     "qty": 3,
 *                     "price": 380,
 *                     "sum_amount": 1140
 *                 },
 *                 {
 *                     ....
 *                 }
 *             ]
 *         }
 *     ]
 * }
 */

/**
 * @api {post} /product/v1/delete/picture/item  DeleteItemPicSub
 * @apiVersion 1.0.0
 * @apiName DeleteItemPicSub
 * @apiGroup ProductBackend
 * @apiDescription DeleteItemPicSub
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token e97ca15ddba94f808369606043f17d08
 * @apiParam (Parameter) {int} id รหัสสินค้า
 * @apiParamExample {json} Body request:
 * {
 *	"id":7
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success"
 * }
 */

/**
 * @api {post} /product/v1/delete/promo/item  DeleteItemPromotionPicSub
 * @apiVersion 1.0.0
 * @apiName DeleteItemPromotionPicSub
 * @apiGroup ProductBackend
 * @apiDescription DeleteItemPromotionPicSub
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token e97ca15ddba94f808369606043f17d08
 * @apiParam (Parameter) {int} id รหัส
 * @apiParamExample {json} Body request:
 * {
 *	"id": 39
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success"
 * }
 */
