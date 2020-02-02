package sales

import (
	"context"
	"fmt"

	"gitlab.com/satit13/perfect_api/auth"
	log "gitlab.com/satit13/perfect_api/logger"
)

func MakeEnpointGetSalesTeam(s Service) interface{} {
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetUserID(ctx)
		log.Println("MakeEnpointGetSalesTeam response: ", UserID)
		resp, err := s.GetSalesTeamService(UserID)

		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}

		return &response{
			Result: "success",
			Data:   resp,
		}, nil
	}
}

func MakeEnpointGetCommsionDocNoBackEnd(s Service) interface{} {
	type request struct {
		DocNo string `json:"doc_no"`
	}
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpoint makeGetSaleByUser")

		resp, err := s.GetCommisionByDocNo(req.DocNo)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/get/commision/admin/docno get commision docno
 * @apiVersion 1.0.2
 * @apiName Commision get commision admin docno
 * @apiGroup CommisionAdmin
 * @apiDescription สำรหับ get บิล ค่าคอม
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token eb109a1cd4d04aeda41c8aad2f7997f3 (token ทดสอบ ของ admin)
 * @apiParam (Parameter) {String} doc_no รหัสเลขที่เอกสาร
 * @apiParamExample {json} Body request:
 * {
 * 	"doc_no":"WDC62120001"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *    "response": "success",
 *    "message": "",
 *    "data": [
 *        {
 *            "id": 29,
 *            "user_id": "62110700001",
 *            "sales_code": "SAL62103100001",
 *            "f_name": "satit",
 *            "l_name": "chomwattana",
 *            "telephone": "+66816398388",
 *            "doc_no": "WDC62113000001",
 *            "doc_date": "2019-11-30 00:00:00",
 *            "total_amount": 8805.96,
 *            "all_sale_commision": 88059.6,
 *            "bank_info": {
 *                "card_id": "1709800118596",
 *                "url_card_id": "url-",
 *                "book_bank_id": "123772710202",
 *                "url_book_bank": "url2",
 *                "bank_code": "",
 *                "bank_name": ""
 *            },
 *            "status": "0",
 *            "my_description": "",
 *            "slip_approve": "",
 *            "confirm": 0,
 *            "confirm_time": "1990-01-01 00:00:00",
 *            "sub": [
 *                {
 *                    "id": 600,
 *                    "order_doc": "ORD621108091438",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 1372
 *                },
 *                {
 *                    "id": 601,
 *                    "order_doc": "ORD621108092240",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 646
 *                },
 *                {
 *                    "id": 602,
 *                    "order_doc": "ORD621108094613",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 646
 *                },
 *                {
 *                    "id": 603,
 *                    "order_doc": "ORD621108140702",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 380
 *                },
 *                {
 *                    "id": 604,
 *                    "order_doc": "ORD621108140843",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 430
 *                },
 *                {
 *                    "id": 605,
 *                    "order_doc": "ORD621108141237",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 810
 *                },
 *                {
 *                    "id": 606,
 *                    "order_doc": "ORD621108141310",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 810
 *                },
 *                {
 *                    "id": 607,
 *                    "order_doc": "ORD621108145948",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 0
 *                },
 *                {
 *                    "id": 608,
 *                    "order_doc": "ORD621108153656",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 0
 *                },
 *                {
 *                    "id": 609,
 *                    "order_doc": "ORD621109170059",
 *                    "order_date": "2019-11-09 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 430
 *                },
 *                {
 *                    "id": 610,
 *                    "order_doc": "ORD621112122038",
 *                    "order_date": "2019-11-12 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 470
 *                },
 *                {
 *                    "id": 611,
 *                    "order_doc": "ORD621115153457",
 *                    "order_date": "2019-11-15 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 30000
 *                },
 *                {
 *                    "id": 612,
 *                    "order_doc": "ORD621115154034",
 *                    "order_date": "2019-11-15 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 30000
 *                },
 *                {
 *                    "id": 613,
 *                    "order_doc": "ORD621123150758",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 2780
 *                },
 *                {
 *                    "id": 614,
 *                    "order_doc": "ORD621123150857",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 450
 *                },
 *                {
 *                    "id": 615,
 *                    "order_doc": "ORD621123151147",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 450
 *                },
 *                {
 *                    "id": 616,
 *                    "order_doc": "ORD621123153419",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 450
 *                },
 *                {
 *                    "id": 617,
 *                    "order_doc": "ORD621105090718",
 *                    "order_date": "2019-11-05 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 12599
 *                },
 *                {
 *                    "id": 618,
 *                    "order_doc": "ORD621106223705",
 *                    "order_date": "2019-11-06 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 5070
 *                },
 *                {
 *                    "id": 619,
 *                    "order_doc": "ORD621106223816",
 *                    "order_date": "2019-11-06 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 430
 *                },
 *                {
 *                    "id": 620,
 *                    "order_doc": "ORD621108142955",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 4396.6
 *                },
 *                {
 *                    "id": 621,
 *                    "order_doc": "ORD621108151356",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 1600
 *                },
 *                {
 *                    "id": 622,
 *                    "order_doc": "ORD621108151652",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 1600
 *                }
 *            ]
 *        }
 *    ]
 *}
 */
func MakeEndpointListCommision(s Service) interface{} {
	type request struct {
		Keyword string `json:"keyword"`
		Limit   int64  `json:"limit"`
		Status  int64  `json:"status"`
	}
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Length   int64       `json:"length"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpooint MakeEndpointApprove :", req)
		len, resp, err := s.ListCommisonService(req.Status, req.Keyword, req.Limit)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Message:  "",
			Length:   len,
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/list/commision/all list commision all
 * @apiVersion 1.0.2
 * @apiName Commision list commision all
 * @apiGroup CommisionAdmin
 * @apiDescription list ค่าคอม สำหรับหลังบ้าน
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token eb109a1cd4d04aeda41c8aad2f7997f3 (token ทดสอบ ของ admin)
 * @apiParam (Parameter) {String} keyword -> search docno กับ name
 * @apiParam (Parameter) {Int} limit
 * @apiParam (Parameter) {String} status 0 -> ร้องขอการถอน กับสถานะยกเลิก 9  | 1 -> ยืนยันการร้องขอ สถานะการอนุมัติ รอการโอนเงิน | 2 -> สถานะการโอนเงินสำเร็จ | 3 - > สถานะ 1 กับ 2 | 4 -> ทุก สถานะ
 * @apiParamExample {json} Body request:
 * {
 * 	"keyword":"WDC 6212 0001"
 *  "limit":1,
 *  "status" :0
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *    "response": "success",
 *    "message": "",
 * 	  "length": 2,
 *    "data": [
 *        {
 *            "id": 29,
 *            "user_id": "62110700001",
 *            "sales_code": "SAL62103100001",
 *            "f_name": "satit",
 *            "l_name": "chomwattana",
 *            "telephone": "+66816398388",
 *            "doc_no": "WDC62113000001",
 *            "doc_date": "2019-11-30 00:00:00",
 *            "total_amount": 8805.96,
 *           "all_sale_commision": 90080,
 *            "bank_info": {
 *                "card_id": "898989898",
 *                "url_card_id": "https://api-dev.pct2003.com/image/sales/sales-490783200.png",
 *                "book_bank_id": "888888",
 *                "url_book_bank": "https://api-dev.pct2003.com/image/sales/sales-454223433.png",
 *                "bank_code": "002",
 *                "bank_name": "ธ. กรุงเทพ จำกัด (มหาชน)"
 *            },
 *            "status": 0,
 *            "my_description": "",
 *            "slip_approve": "",
 *            "confirm": 0,
 *            "confirm_time": "1990-01-01 00:00:00",
 *        },
 *        {
 *            "id": 29,
 *            "user_id": "62110700001",
 *            "user_id": "62110700001",
 *            "sales_code": "SAL62103100002",
 *            "f_name": "สุริยา",
 *            "l_name": "เทียมแย้ม",
 *            "telephone": "+66816398388",
 *            "doc_no": "WDC62113000001",
 *            "doc_date": "2019-11-30 00:00:00",
 *            "total_amount": 8805.96,
 *            "all_sale_commision": 88059.6,
 *            "bank_info": {
 *                "card_id": "1709800118596",
 *                "url_card_id": "url-",
 *                "book_bank_id": "123772710202",
 *                "url_book_bank": "url2",
 *                "bank_code": "",
 *                "bank_name": ""
 *            },
 *            "status": 0,
 *            "my_description": "",
 *            "slip_approve": "",
 *            "confirm": 0,
 *            "confirm_time": "1990-01-01 00:00:00",
 *        }
 *    ]
 *}
 */

func MakeEndpointApprove(s Service) interface{} {
	type request struct {
		DocNo       string `json:"doc_no"`
		Status      int64  `json:"status"`
		SlipApprove string `json:"slip_approve"`
	}
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpooint MakeEndpointApprove :", req)
		AdminID := auth.GetAdminID(ctx)
		data, err := s.ApproveCommision(AdminID, req.DocNo, req.Status, req.SlipApprove)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Message:  "ลบข้อมูล sales สำเร็จ",
			Data:     data,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/update/approve/commision update approve commision docno
 * @apiVersion 1.0.2
 * @apiName Commision approve commision docno
 * @apiGroup CommisionAdmin
 * @apiDescription ยืนยัน ค่าคอม
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token eb109a1cd4d04aeda41c8aad2f7997f3 (token ทดสอบ ของ aadmin)
 * @apiParam (Parameter) {String} doc_no รหัสเลขที่เอกสาร
 * @apiParam (Parameter) {String} status สถานะบิล 1 -> ยืนยันการร้องขอ ->  สถานะการอนุมัติ รอการโอนเงิน  | 2 -> อนุมัติการโอนเงิน พร้อมแนบสลิป
 * @apiParam (Parameter) {String} slip_approve แนปสลิป แนบสลิปตอนอนุมัตโอนเงิน status = 2
 * @apiParamExample {json} Body request:
 * {
 * 	"doc_no":"WDC62120001"
 *  "status":1,
 *  "slip_approve" :"url"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *    "response": "success",
 *    "message": "ยืนยัน สำเร็จ",
 *    "data": [
 *        {
 *            "id": 29,
 *            "user_id": "62110700001",
 *            "sales_code": "SAL62103100001",
 *            "f_name": "satit",
 *            "l_name": "chomwattana",
 *            "telephone": "+66816398388",
 *            "doc_no": "WDC62113000001",
 *            "doc_date": "2019-11-30 00:00:00",
 *            "total_amount": 8805.96,
 *            "bank_info": {
 *                "card_id": "1709800118596",
 *                "url_card_id": "url-",
 *                "book_bank_id": "123772710202",
 *                "url_book_bank": "url2",
 *                "bank_code": "",
 *                "bank_name": ""
 *            },
 *            "status": 1,
 *            "my_description": "",
 *            "slip_approve": "",
 *            "confirm": 0,
 *            "confirm_time": "1990-01-01 00:00:00",
 *            "sub": [
 *                {
 *                    "id": 600,
 *                    "order_doc": "ORD621108091438",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 1372
 *                },
 *                {
 *                    "id": 601,
 *                    "order_doc": "ORD621108092240",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 646
 *                },
 *                {
 *                    "id": 602,
 *                    "order_doc": "ORD621108094613",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 646
 *                },
 *                {
 *                    "id": 603,
 *                    "order_doc": "ORD621108140702",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 380
 *                },
 *                {
 *                    "id": 604,
 *                    "order_doc": "ORD621108140843",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 430
 *                },
 *                {
 *                    "id": 605,
 *                    "order_doc": "ORD621108141237",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 810
 *                },
 *                {
 *                    "id": 606,
 *                    "order_doc": "ORD621108141310",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 810
 *                },
 *                {
 *                    "id": 607,
 *                    "order_doc": "ORD621108145948",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 0
 *                },
 *                {
 *                    "id": 608,
 *                    "order_doc": "ORD621108153656",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 0
 *                },
 *                {
 *                    "id": 609,
 *                    "order_doc": "ORD621109170059",
 *                    "order_date": "2019-11-09 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 430
 *                },
 *                {
 *                    "id": 610,
 *                    "order_doc": "ORD621112122038",
 *                    "order_date": "2019-11-12 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 470
 *                },
 *                {
 *                    "id": 611,
 *                    "order_doc": "ORD621115153457",
 *                    "order_date": "2019-11-15 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 30000
 *                },
 *                {
 *                    "id": 612,
 *                    "order_doc": "ORD621115154034",
 *                    "order_date": "2019-11-15 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 30000
 *                },
 *                {
 *                    "id": 613,
 *                    "order_doc": "ORD621123150758",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 2780
 *                },
 *                {
 *                    "id": 614,
 *                    "order_doc": "ORD621123150857",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 450
 *                },
 *                {
 *                    "id": 615,
 *                    "order_doc": "ORD621123151147",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 450
 *                },
 *                {
 *                    "id": 616,
 *                    "order_doc": "ORD621123153419",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 450
 *                },
 *                {
 *                    "id": 617,
 *                    "order_doc": "ORD621105090718",
 *                    "order_date": "2019-11-05 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 12599
 *                },
 *                {
 *                    "id": 618,
 *                    "order_doc": "ORD621106223705",
 *                    "order_date": "2019-11-06 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 5070
 *                },
 *                {
 *                    "id": 619,
 *                    "order_doc": "ORD621106223816",
 *                    "order_date": "2019-11-06 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 430
 *                },
 *                {
 *                    "id": 620,
 *                    "order_doc": "ORD621108142955",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 4396.6
 *                },
 *                {
 *                    "id": 621,
 *                    "order_doc": "ORD621108151356",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 1600
 *                },
 *                {
 *                    "id": 622,
 *                    "order_doc": "ORD621108151652",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 1600
 *                }
 *            ]
 *        }
 *    ]
 *}
 */

func DeleteSalesEnpdoint(s Service) interface{} {

	type request struct {
		SalesCode string `json:"sales_code"`
		UserID    string `json:"user_id"`
	}
	type response struct {
		Response string `json:"response"`
		Message  string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpooint makeRegisterSaleEndpoint :", req)

		AdminID := auth.GetAdminID(ctx)
		_, err := s.RemoveSales(AdminID, req.SalesCode, req.UserID)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Message:  "ลบข้อมูล sales สำเร็จ",
		}, nil
	}

}

/**
 * @api {post} /sales/v1/remove/sales  remove sales
 * @apiVersion 1.0.0
 * @apiName remove sales
 * @apiGroup Sales
 * @apiDescription สำรหับ remove sales
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token c0c05dd13ac8490e934bd361e24b724a (token ทดสอบ ของ admin)
 * @apiParam (Parameter) {String} sales_code รหัส sales
 * @apiParam (Parameter) {string} user_id รหัส ของ user 62083000013
 * @apiParamExample {json} Body request:
 * {
 * 	"sales_code":"SAL62100500001",
 * 	"user_id":"62083000013"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "ลบข้อมูล sales สำเร็จ"
 * }
 */

func makeRegisterSaleEndpoint(s Service) interface{} {
	type request RegisterSalesModel
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpooint makeRegisterSaleEndpoint :", req)

		UserID := auth.GetUserID(ctx)
		sale := RegisterSalesModel{
			CardID:      req.CardID,
			Fname:       req.Fname,
			Lname:       req.Lname,
			URLCardID:   req.URLCardID,
			URLBookBank: req.URLBookBank,
			URLFirstBuy: req.URLFirstBuy,
			File3:       req.File3,
			BookBankID:  req.BookBankID,
			BankCode:    req.BankCode,
			BankName:    req.BankName,
		}
		resp, err := s.RegisterSaleService(UserID, &sale)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/sales/register  sales register
 * @apiVersion 1.0.2
 * @apiName Sales Register
 * @apiGroup Sales
 * @apiDescription สำรหับ ลง ทะเบียน sales
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token ee26c312a27b4cc8ba403c3ddc53323b (token ทดสอบ ของ user)
 * @apiParam (Parameter) {String} card_id รหัสบัตรปชช
 * @apiParam (Parameter) {string} f_name ชื่อ
 * @apiParam (Parameter) {string} l_name นามสกุล
 * @apiParam (Parameter) {string} url_card_id urlรูปcard_id
 * @apiParam (Parameter) {string} url_book_bank urlรูปbook_bank
 * @apiParam (Parameter) (string) url_slip_first_buy รูปสลิปที่แนบ
 * @apiParam (Parameter) {string} file_url_3 urlรูป 3
 * @apiParam (Parameter) {string} book_bank_id รหัส bookbank
 * @apiParamExample {json} Body request:
 * {
 *     "card_id": "1709800118596",
 *     "f_name": "สุริยา",
 *     "l_name": "เทียมแย้ม",
 *     "url_card_id": "url-",
 *     "url_book_bank": "url2",
 *     "file_url_3": "12312312",
 *     "url_slip_first_buy":"124124124",
 *     "book_bank_id": "123772710202",
 *     "bank_code":"bkb",
 *     "bank_name":"ธนาคารกรุงเทพ"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": {
 *         "id": 69,
 *         "user_id": "62110700001",
 *         "sales_code": "SAL62111300001",
 *         "first_name": "สุริยา",
 *         "invite_person": 0,
 *         "last_name": "เทียมแย้ม",
 *         "role_id": 1,
 *         "role_name": "user",
 *         "card_id": "1709800118596",
 *         "book_bank_id": "123772710202",
 *         "age": 0,
 *         "url_card_id": "url-",
 *         "url_book_bank": "url2",
 *         "url_slip_first_buy": "124124124",
 *         "file_url_3": "12312312",
 *         "invite_code": "2c6fae0e6e6443b487a4e66066d45566",
 *         "confirm": 0,
 *         "bank_code": "bkb",
 *         "bank_name": "ธนาคารกรุงเทพ",
 *         "create_date": "2019-11-13 20:33:05"
 *     }
 * }
 */

func makeUpdateSaleByAdmin(s Service) interface{} {
	type request UpdateSaleModelAdmin
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpoint Update Sale By Admin", req)
		AdminID := auth.GetAdminID(ctx)
		sale := UpdateSaleModelAdmin{
			SaleCode:    req.SaleCode,
			CardID:      req.CardID,
			Fname:       req.Fname,
			Lname:       req.Lname,
			URLCardID:   req.URLCardID,
			URLBookBank: req.URLBookBank,
			URLFirstBuy: req.URLFirstBuy,
			File3:       req.File3,
			BookBankID:  req.BookBankID,
			BankCode:    req.BankCode,
			BankName:    req.BankName,
			Confirm:     req.Confirm,
		}
		_, err := s.UpdateSaleServiceAdmin(AdminID, &sale)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     sale,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/update/sales/byadmin sales update by Admin
 * @apiVersion 1.0.2
 * @apiName Sales Update by Admin
 * @apiGroup Sales
 * @apiDescription สำรหับ update sales
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token c0947e79bb604eabad865b93aa7f333c (token ทดสอบ ของ Admin)
 * @apiParam (Parameter) {String} sales_code รหัส sales
 * @apiParam (Parameter) {String} card_id รหัสบัตรปชช
 * @apiParam (Parameter) {string} f_name ชื่อ
 * @apiParam (Parameter) {string} l_name นามสกุล
 * @apiParam (Parameter) {string} url_card_id urlรูปcard_id
 * @apiParam (Parameter) {string} url_book_bank urlรูปbook_bank
 * @apiParam (Parameter) {string} file_url_3 urlรูป 3
 * @apiParam (Parameter) {string} book_bank_id รหัส bookbank
 * @apiParam (Parameter) {int} confirm confirm
 * @apiParamExample {json} Body request:
 * {
 *     "sales_code":"SAL62100500001",
 *     "card_id": "1709800118596",
 *     "f_name": "สุริยา",
 *     "l_name": "เทียมแย้ม",
 *     "url_card_id": "url-",
 *     "url_book_bank": "url2",
 *     "file_url_3": "12312312",
 *     "url_slip_first_buy":"124124124",
 *     "book_bank_id": "123772710202",
 *     "bank_code":"002",
 *     "bank_name":"ธ. กรุงเทพ จำกัด (มหาชน)",
 *     "confirm":1
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": {
 *         "sales_code":"SAL62100500001",
 *         "card_id": "1709800118596",
 *         "f_name": "สุริยา",
 *         "l_name": "เทียมแย้ม",
 *         "url_card_id": "url-",
 *         "url_book_bank": "url2",
 *         "url_slip_first_buy": "124124124",
 *         "file_url_3": "12312312",
 *         "book_bank_id": "123772710202",
 *         "bank_code": "bkb",
 *         "bank_name": "ธนาคารกรุงเท1พ",
 *         "confirm": 1
 *     }
 * }
 */

func makeCancelCommision(s Service) interface{} {
	type request struct {
		DocNo string `json:"doc_no"`
	}
	type response struct {
		Response string `json:"response"`
		Message  string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		_, err := s.CancelCommisionDocNoService(UserID, req.DocNo)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Message:  "ยกเลิกบิล สำเร็จ",
		}, nil
	}
}

/**
 * @api {post} /sales/v1/cancel/commision/user/docno  cancel commision
 * @apiVersion 1.0.0
 * @apiName cancel commision
 * @apiGroup Commision
 * @apiDescription สำรหับ cancel commisionสำหรับ user
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token c0c05dd13ac8490e934bd361e24b724a (token ทดสอบ ของ user)
 * @apiParam (Parameter) {String} doc_no รหัส เอกสาร
 * @apiParamExample {json} Body request:
 * {
 * 	"doc_no":"WDC62113000001",
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "ยกเลิกบิล สำเร็จ"
 * }
 */
func makeCancelCommisionFont(s Service) interface{} {
	type request struct {
		DocNo string `json:"doc_no"`
	}
	type response struct {
		Response string `json:"response"`
		Message  string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		AdminID := auth.GetAdminID(ctx)
		_, err := s.CancelCommisionFontService(AdminID, req.DocNo)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Message:  "ยกเลิกบิลสำเร็จ",
		}, nil
	}
}

/**
 * @api {post} /sales/v1/cancel/commision/admin/docno  cancel commision admin
 * @apiVersion 1.0.0
 * @apiName cancel commision admin
 * @apiGroup CommisionAdmin
 * @apiDescription สำรหับ cancel commision admin
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token c0c05dd13ac8490e934bd361e24b724a (token ทดสอบ ของ admin)
 * @apiParam (Parameter) {String} doc_no รหัส เอกสาร
 * @apiParamExample {json} Body request:
 * {
 * 	"doc_no":"WDC62113000001",
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "ยกเลิกบิล สำเร็จ"
 * }
 */

func makeUpdateSaleEndpoint(s Service) interface{} {
	type request UpdateSaleModel
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpoint Update Sale By Admin", req)
		UserID := auth.GetUserID(ctx)
		sale := UpdateSaleModel{
			CardID:      req.CardID,
			Fname:       req.Fname,
			Lname:       req.Lname,
			URLCardID:   req.URLCardID,
			URLBookBank: req.URLBookBank,
			URLFirstBuy: req.URLFirstBuy,
			File3:       req.File3,
			BookBankID:  req.BookBankID,
			BankCode:    req.BankCode,
			BankName:    req.BankName,
			Confirm:     req.Confirm,
		}
		_, err := s.UpdateSaleService(UserID, &sale)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     sale,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/update/sales sales update
 * @apiVersion 1.0.2
 * @apiName Sales Update Sales
 * @apiGroup Sales
 * @apiDescription สำรหับ update sales
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token ee26c312a27b4cc8ba403c3ddc53323b (token ทดสอบ ของ user)
 * @apiParam (Parameter) {String} card_id รหัสบัตรปชช
 * @apiParam (Parameter) {string} f_name ชื่อ
 * @apiParam (Parameter) {string} l_name นามสกุล
 * @apiParam (Parameter) {string} url_card_id urlรูปcard_id
 * @apiParam (Parameter) {string} url_book_bank urlรูปbook_bank
 * @apiParam (Parameter) {string} file_url_3 urlรูป 3
 * @apiParam (Parameter) {string} book_bank_id รหัส bookbank
 * @apiParam (Parameter) {int} confirm confirm
 * @apiParamExample {json} Body request:
 * {
 *     "card_id": "1709800118596",
 *     "f_name": "สุริยา",
 *     "l_name": "เทียมแย้ม",
 *     "url_card_id": "url-",
 *     "url_book_bank": "url2",
 *     "file_url_3": "12312312",
 *     "url_slip_first_buy":"124124124",
 *     "book_bank_id": "123772710202",
 *     "bank_code":"bkb",
 *     "bank_name":"ธนาคารกรุงเท1พ",
 *     "confirm":1
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": {
 *         "card_id": "1709800118596",
 *         "f_name": "สุริยา",
 *         "l_name": "เทียมแย้ม",
 *         "url_card_id": "url-",
 *         "url_book_bank": "url2",
 *         "url_slip_first_buy": "124124124",
 *         "file_url_3": "12312312",
 *         "book_bank_id": "123772710202",
 *         "bank_code": "bkb",
 *         "bank_name": "ธนาคารกรุงเท1พ",
 *         "confirm": 1
 *     }
 * }
 */

func makeGetSaleByUserV2Endpoint(s Service) interface{} {
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		log.Println("Endpoint makeGetSaleByUserV2Endpoint")
		UserID := auth.GetUserID(ctx)
		resp, err := s.GetSaleByUserV2(UserID)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/get/sale/commision get sale commision
 * @apiVersion 1.0.2
 * @apiName Get  sale commision
 * @apiGroup Commision
 * @apiDescription สำรหับ Get ข้อมูลการสมัค sales commison
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token ee26c312a27b4cc8ba403c3ddc53323b (token ทดสอบ ของ user)
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 *  {
 *     "response": "success",
 *     "message": "",
 *     "data": {
 *         "id": 21,
 *         "user_id": "62101600002",
 *         "sales_code": "SAL62101600001",
 *         "first_name": "สมรถ",
 *         "invite_person": 4,
 *         "last_name": "หลักฐาน",
 *         "role_id": 2,
 *         "role_name": "sales",
 *         "card_id": "5",
 *         "book_bank_id": "6",
 *         "age": 0,
 *         "sale": 22495.6,
 *         "invite_sale": 344097,
 *         "commission": 34409.7,
 *         "all_sale": 765330.6,
 *         "url_card_id": "https://api.pct2003.com/image/sales/sales-500355709.png",
 *         "url_book_bank": "https://api.pct2003.com/image/sales/sales-146691512.png",
 *         "url_slip_first_buy": "https://api.pct2003.com/image/sales/sales-146691512.png",
 *         "file_url_3": "",
 *         "invite_code": "a5beab81fcb547f5ad2817a242fd5e96",
 *         "incentive": 15306.612,
 *         "confirm": 1,
 *         "create_date": "2019-10-16 05:02:02",
 *         "all_level": [
 *             {
 *                 "level": 0,
 *                 "all_sale": 22495.6
 *             },
 *             {
 *                 "level": 1,
 *                 "all_sale": 344097
 *             },
 *             {
 *                 "level": 2,
 *                 "all_sale": 15304
 *             },
 *             {
 *                 "level": 3,
 *                 "all_sale": 313804
 *             },
 *             {
 *                 "level": 4,
 *                 "all_sale": 69630
 *             }
 *         ],
 *         "invite": [
 *             {
 *                 "company_id": 0,
 *                 "role_id": 1,
 *                 "user_id": "62102500002",
 *                 "role_name": "user",
 *                 "profix_name": "",
 *                 "fname": "สมรัก",
 *                 "lname": "หลักฐาน",
 *                 "active_status": 1,
 *                 "telephone": "+66910732880",
 *                 "email": "",
 *                 "ref_code": "a5beab81fcb547f5ad2817a242fd5e96",
 *                 "invite_code": "0c3dc71eb4e344778985db7d22f137da",
 *                 "invite_person": 0,
 *                 "sale": 15350,
 *                 "invite_sale": 0,
 *                 "commission": 0
 *             },
 *             {
 *                 "company_id": 0,
 *                 "role_id": 2,
 *                 "user_id": "62102500003",
 *                 "role_name": "sales",
 *                 "profix_name": "",
 *                 "fname": "สมปอง",
 *                 "lname": "หลักฐาน",
 *                 "active_status": 1,
 *                 "telephone": "+66910732781",
 *                 "email": "",
 *                 "ref_code": "a5beab81fcb547f5ad2817a242fd5e96",
 *                 "invite_code": "b12282a1288b40c78c75f23e8c53db47",
 *                 "invite_person": 8,
 *                 "sale": 28747,
 *                 "invite_sale": 15304,
 *                 "commission": 1530.4,
 *                 "invite": [
 *                     {
 *                         "company_id": 0,
 *                         "role_id": 2,
 *                         "user_id": "62102500005",
 *                         "role_name": "sales",
 *                         "profix_name": "",
 *                         "fname": "ปองศักดิ์",
 *                         "lname": "หลักฐาน",
 *                         "active_status": 1,
 *                         "telephone": "+66910732800",
 *                         "email": "",
 *                         "ref_code": "b12282a1288b40c78c75f23e8c53db47",
 *                         "invite_code": "823dc08260e84b9fae21695a39efa06e",
 *                         "invite_person": 1,
 *                         "sale": 15304,
 *                         "invite_sale": 3260,
 *                         "commission": 326,
 *                         "invite": [
 *                             {
 *                                 "company_id": 0,
 *                                 "role_id": 2,
 *                                 "user_id": "62102500007",
 *                                 "role_name": "sales",
 *                                 "profix_name": "",
 *                                 "fname": "ศักดิ์กรินทร์",
 *                                 "lname": "หลักฐาน",
 *                                 "active_status": 1,
 *                                 "telephone": "+66910732801",
 *                                 "email": "",
 *                                 "ref_code": "823dc08260e84b9fae21695a39efa06e",
 *                                 "invite_code": "77d2150eb42d4b3ca7f74c8e4133de0d",
 *                                 "invite_person": 2,
 *                                 "sale": 3260,
 *                                 "invite_sale": 69630,
 *                                 "commission": 6963,
 *                                 "invite": [
 *                                     {
 *                                         "company_id": 0,
 *                                         "role_id": 2,
 *                                         "user_id": "62102500008",
 *                                         "role_name": "sales",
 *                                         "profix_name": "",
 *                                         "fname": "รินณภัชร์",
 *                                         "lname": "หลักฐาน",
 *                                         "active_status": 1,
 *                                         "telephone": "+66910738010",
 *                                         "email": "",
 *                                         "ref_code": "77d2150eb42d4b3ca7f74c8e4133de0d",
 *                                         "invite_code": "ebf72913d3924cd0844f9a6cd1c2a67d",
 *                                         "invite_person": 0,
 *                                         "sale": 33450,
 *                                         "invite_sale": 0,
 *                                         "commission": 0
 *                                     },
 *                                     {
 *                                         "company_id": 0,
 *                                         "role_id": 2,
 *                                         "user_id": "62102500009",
 *                                         "role_name": "sales",
 *                                         "profix_name": "",
 *                                         "fname": "รินฤทธิ์",
 *                                         "lname": "หลักฐาน",
 *                                         "active_status": 1,
 *                                         "telephone": "+66910732811",
 *                                         "email": "",
 *                                         "ref_code": "77d2150eb42d4b3ca7f74c8e4133de0d",
 *                                         "invite_code": "5f4ebf0180964a2cb6f03c73d2feb83c",
 *                                         "invite_person": 1,
 *                                         "sale": 36180,
 *                                         "invite_sale": 130799,
 *                                         "commission": 13079.9,
 *                                         "invite": [
 *                                             {
 *                                                 "company_id": 0,
 *                                                 "role_id": 2,
 *                                                 "user_id": "62110700002",
 *                                                 "role_name": "sales",
 *                                                 "profix_name": "",
 *                                                 "fname": "ฤทธิไกร",
 *                                                 "lname": "หลักฐาน",
 *                                                 "active_status": 1,
 *                                                 "telephone": "+66910732870",
 *                                                 "email": "",
 *                                                 "ref_code": "5f4ebf0180964a2cb6f03c73d2feb83c",
 *                                                 "invite_code": "ff82c168125e4d1cb42a995c63336e6e",
 *                                                 "invite_person": 1,
 *                                                 "sale": 130799,
 *                                                 "invite_sale": 0,
 *                                                 "commission": 0
 *                                             }
 *                                         ]
 *                                     }
 *                                 ]
 *                             }
 *                         ]
 *                     },
 *                     {
 *                         "company_id": 0,
 *                         "role_id": 2,
 *                         "user_id": "62102900001",
 *                         "role_name": "sales",
 *                         "profix_name": "",
 *                         "fname": "ปองพล",
 *                         "lname": "หลักฐาน",
 *                         "active_status": 1,
 *                         "telephone": "+66910732820",
 *                         "email": "",
 *                         "ref_code": "b12282a1288b40c78c75f23e8c53db47",
 *                         "invite_code": "d2a2ed1a2bf842eeac9ded3bf8d74763",
 *                         "invite_person": 0,
 *                         "sale": 0,
 *                         "invite_sale": 0,
 *                         "commission": 0
 *                     },
 *                     {
 *                         "company_id": 0,
 *                         "role_id": 2,
 *                         "user_id": "62102900002",
 *                         "role_name": "sales",
 *                         "profix_name": "",
 *                         "fname": "ปองสุข",
 *                         "lname": "หลักฐาน",
 *                         "active_status": 1,
 *                         "telephone": "+66910732830",
 *                         "email": "",
 *                         "ref_code": "b12282a1288b40c78c75f23e8c53db47",
 *                         "invite_code": "801d5e1fd2a94f659d6c6d65515f394a",
 *                         "invite_person": 1,
 *                         "sale": 0,
 *                         "invite_sale": 310544,
 *                         "commission": 31054.4,
 *                         "invite": [
 *                             {
 *                                 "company_id": 0,
 *                                 "role_id": 1,
 *                                 "user_id": "62102900003",
 *                                 "role_name": "user",
 *                                 "profix_name": "",
 *                                 "fname": "สุขสวัสดิ์",
 *                                 "lname": "หลักฐาน",
 *                                 "active_status": 1,
 *                                 "telephone": "+66910732831",
 *                                 "email": "",
 *                                 "ref_code": "801d5e1fd2a94f659d6c6d65515f394a",
 *                                 "invite_code": "49f7de01706f400bb49bea1be039c5ba",
 *                                 "invite_person": 0,
 *                                 "sale": 310544,
 *                                 "invite_sale": 0,
 *                                 "commission": 0
 *                             }
 *                         ]
 *                     },
 *                     {
 *                         "company_id": 0,
 *                         "role_id": 2,
 *                         "user_id": "62102900004",
 *                         "role_name": "sales",
 *                         "profix_name": "",
 *                         "fname": "ปองกูล",
 *                         "lname": "หลักฐาน",
 *                         "active_status": 1,
 *                         "telephone": "+66910732840",
 *                         "email": "",
 *                         "ref_code": "b12282a1288b40c78c75f23e8c53db47",
 *                         "invite_code": "7ce0f25b767a4d66a1815e5df01e9d96",
 *                         "invite_person": 0,
 *                         "sale": 0,
 *                         "invite_sale": 0,
 *                         "commission": 0
 *                     },
 *                     {
 *                         "company_id": 0,
 *                         "role_id": 1,
 *                         "user_id": "62102900005",
 *                         "role_name": "user",
 *                         "profix_name": "",
 *                         "fname": "ปองจิตต์",
 *                         "lname": "หลักฐาน",
 *                         "active_status": 1,
 *                         "telephone": "+66910732785",
 *                         "email": "",
 *                         "ref_code": "b12282a1288b40c78c75f23e8c53db47",
 *                         "invite_code": "7133ed2fb4bd4c0baf1568fab87f2719",
 *                         "invite_person": 0,
 *                         "sale": 0,
 *                         "invite_sale": 0,
 *                         "commission": 0
 *                     },
 *                     {
 *                         "company_id": 0,
 *                         "role_id": 1,
 *                         "user_id": "62102900006",
 *                         "role_name": "user",
 *                         "profix_name": "",
 *                         "fname": "ปองเกียรติ",
 *                         "lname": "หลักฐาน",
 *                         "active_status": 1,
 *                         "telephone": "+66910732786",
 *                         "email": "",
 *                         "ref_code": "b12282a1288b40c78c75f23e8c53db47",
 *                         "invite_code": "8ebc30ac5fa14ffcb557232b010c12d7",
 *                         "invite_person": 0,
 *                         "sale": 0,
 *                         "invite_sale": 0,
 *                         "commission": 0
 *                     },
 *                     {
 *                         "company_id": 0,
 *                         "role_id": 2,
 *                         "user_id": "62102900007",
 *                         "role_name": "sales",
 *                         "profix_name": "",
 *                         "fname": "ปองชัย",
 *                         "lname": "หลักฐาน",
 *                         "active_status": 1,
 *                         "telephone": "+66910732850",
 *                         "email": "",
 *                         "ref_code": "b12282a1288b40c78c75f23e8c53db47",
 *                         "invite_code": "9563940e1f8d4185b02c6f787e15bc3c",
 *                         "invite_person": 0,
 *                         "sale": 0,
 *                         "invite_sale": 0,
 *                         "commission": 0
 *                     },
 *                     {
 *                         "company_id": 0,
 *                         "role_id": 2,
 *                         "user_id": "62102900008",
 *                         "role_name": "sales",
 *                         "profix_name": "",
 *                         "fname": "ปองศรี",
 *                         "lname": "หลักฐาน",
 *                         "active_status": 1,
 *                         "telephone": "+66910732860",
 *                         "email": "",
 *                         "ref_code": "b12282a1288b40c78c75f23e8c53db47",
 *                         "invite_code": "ff4acbb191f04634b2fd11ab165f50bd",
 *                         "invite_person": 0,
 *                         "sale": 0,
 *                         "invite_sale": 0,
 *                         "commission": 0
 *                     }
 *                 ]
 *             },
 *             {
 *                 "company_id": 0,
 *                 "role_id": 1,
 *                 "user_id": "62102500004",
 *                 "role_name": "user",
 *                 "profix_name": "",
 *                 "fname": "สมหมาย",
 *                 "lname": "หลักฐาน",
 *                 "active_status": 1,
 *                 "telephone": "+66910732782",
 *                 "email": "",
 *                 "ref_code": "a5beab81fcb547f5ad2817a242fd5e96",
 *                 "invite_code": "310e30ab2a254a4daad39da8b2a06be1",
 *                 "invite_person": 0,
 *                 "sale": 0,
 *                 "invite_sale": 0,
 *                 "commission": 0
 *             },
 *             {
 *                 "company_id": 0,
 *                 "role_id": 2,
 *                 "user_id": "62102500006",
 *                 "role_name": "sales",
 *                 "profix_name": "",
 *                 "fname": "สมบูรณ์",
 *                 "lname": "หลักฐาน",
 *                 "active_status": 1,
 *                 "telephone": "+66910732783",
 *                 "email": "",
 *                 "ref_code": "a5beab81fcb547f5ad2817a242fd5e96",
 *                 "invite_code": "2829911f17a742b7a09006e636f49d4e",
 *                 "invite_person": 0,
 *                 "sale": 300000,
 *                 "invite_sale": 0,
 *                 "commission": 0
 *             }
 *         ]
 *     }
 * }
 */

func makeGenarateCommision(s Service) interface{} {
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetUserID(ctx)
		resp, err := s.GenarateComisionService(UserID)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/genarte/commision genarate commision
 * @apiVersion 1.0.2
 * @apiName Commision genarate commision
 * @apiGroup Commision
 * @apiDescription สำรหับ genarate ค่าคอม
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token eb109a1cd4d04aeda41c8aad2f7997f3 (token ทดสอบ ของ user)
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *    "response": "success",
 *    "message": "",
 *    "data": [
 *            "id": 29,
 *            "user_id": "62110700001",
 *            "sales_code": "SAL62103100002",
 *            "f_name": "สุริยา",
 *            "l_name": "เทียมแย้ม",
 *            "telephone": "+66987727968",
 *            "doc_no": "WDC62113000001",
 *            "doc_date": "2019-11-30 00:00:00",
 *            "total_amount": 8805.96,
 *            "all_sale_commision": 88059.6,
 *            "bank_info": {
 *                "card_id": "1709800118596",
 *                "url_card_id": "url-",
 *                "book_bank_id": "123772710202",
 *                "url_book_bank": "url2",
 *                "bank_code": "",
 *                "bank_name": ""
 *            },
 *            "status": 0,
 *            "my_description": "",
 *            "slip_approve": "",
 *            "confirm": 0,
 *            "confirm_time": "1990-01-01 00:00:00",
 *        }
 *    ]
 *}
 * @apiErrorExample Error-Response:
 * {
 *     "response": "false",
 *     "message": "ค่าคอมถอนขึ้นต่ำ 3000 บาท",
 *     "data": null
 * }
 */

func makeGetCommision(s Service) interface{} {
	type request struct {
		DocNo string `json:"doc_no"`
	}
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpoint makeGetSaleByUser")
		UserID := auth.GetUserID(ctx)
		resp, err := s.GetCommisionByDocNoService(UserID, req.DocNo)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/get/commision/docno get commision docno
 * @apiVersion 1.0.2
 * @apiName Commision get commision docno
 * @apiGroup Commision
 * @apiDescription สำรหับ get บิล ค่าคอม
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token eb109a1cd4d04aeda41c8aad2f7997f3 (token ทดสอบ ของ user)
 * @apiParam (Parameter) {String} doc_no รหัสเลขที่เอกสาร
 * @apiParamExample {json} Body request:
 * {
 * 	"doc_no":"WDC62120001"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *    "response": "success",
 *    "message": "",
 *    "data": [
 *        {
 *            "id": 29,
 *            "user_id": "62110700001",
 *            "sales_code": "SAL62103100002",
 *            "f_name": "สุริยา",
 *            "l_name": "เทียมแย้ม",
 *            "telephone": "+66987727968",
 *            "doc_no": "WDC62113000001",
 *            "doc_date": "2019-11-30 00:00:00",
 *            "total_amount": 8805.96,
 *            "all_sale_commision": 88059.6,
 *            "bank_info": {
 *                "card_id": "1709800118596",
 *                "url_card_id": "url-",
 *                "book_bank_id": "123772710202",
 *                "url_book_bank": "url2",
 *                "bank_code": "",
 *                "bank_name": ""
 *            },
 *            "status": 0,
 *            "my_description": "",
 *            "slip_approve": "",
 *            "confirm": 0,
 *            "confirm_time": "1990-01-01 00:00:00",
 *            "sub": [
 *                {
 *                    "id": 600,
 *                    "order_doc": "ORD621108091438",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 1372
 *                },
 *                {
 *                    "id": 601,
 *                    "order_doc": "ORD621108092240",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 646
 *                },
 *                {
 *                    "id": 602,
 *                    "order_doc": "ORD621108094613",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 646
 *                },
 *                {
 *                    "id": 603,
 *                    "order_doc": "ORD621108140702",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 380
 *                },
 *                {
 *                    "id": 604,
 *                    "order_doc": "ORD621108140843",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 430
 *                },
 *                {
 *                    "id": 605,
 *                    "order_doc": "ORD621108141237",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 810
 *                },
 *                {
 *                    "id": 606,
 *                    "order_doc": "ORD621108141310",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 810
 *                },
 *                {
 *                    "id": 607,
 *                    "order_doc": "ORD621108145948",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 0
 *                },
 *                {
 *                    "id": 608,
 *                    "order_doc": "ORD621108153656",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 0
 *                },
 *                {
 *                    "id": 609,
 *                    "order_doc": "ORD621109170059",
 *                    "order_date": "2019-11-09 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 430
 *                },
 *                {
 *                    "id": 610,
 *                    "order_doc": "ORD621112122038",
 *                    "order_date": "2019-11-12 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 470
 *                },
 *                {
 *                    "id": 611,
 *                    "order_doc": "ORD621115153457",
 *                    "order_date": "2019-11-15 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 30000
 *                },
 *                {
 *                    "id": 612,
 *                    "order_doc": "ORD621115154034",
 *                    "order_date": "2019-11-15 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 30000
 *                },
 *                {
 *                    "id": 613,
 *                    "order_doc": "ORD621123150758",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 2780
 *                },
 *                {
 *                    "id": 614,
 *                    "order_doc": "ORD621123150857",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 450
 *                },
 *                {
 *                    "id": 615,
 *                    "order_doc": "ORD621123151147",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 450
 *                },
 *                {
 *                    "id": 616,
 *                    "order_doc": "ORD621123153419",
 *                    "order_date": "2019-11-23 00:00:00",
 *                    "user_id": "62100900011",
 *                    "f_name": "tanakorn",
 *                    "l_name": "phuntulap",
 *                    "total_amount": 450
 *                },
 *                {
 *                    "id": 617,
 *                    "order_doc": "ORD621105090718",
 *                    "order_date": "2019-11-05 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 12599
 *                },
 *                {
 *                    "id": 618,
 *                    "order_doc": "ORD621106223705",
 *                    "order_date": "2019-11-06 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 5070
 *                },
 *                {
 *                    "id": 619,
 *                    "order_doc": "ORD621106223816",
 *                    "order_date": "2019-11-06 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 430
 *                },
 *                {
 *                    "id": 620,
 *                    "order_doc": "ORD621108142955",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 4396.6
 *                },
 *                {
 *                    "id": 621,
 *                    "order_doc": "ORD621108151356",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 1600
 *                },
 *                {
 *                    "id": 622,
 *                    "order_doc": "ORD621108151652",
 *                    "order_date": "2019-11-08 00:00:00",
 *                    "user_id": "62101600002",
 *                    "f_name": "สมรถ",
 *                    "l_name": "หลักฐาน",
 *                    "total_amount": 1600
 *                }
 *            ]
 *        }
 *    ]
 *}
 */
func makeListCommision(s Service) interface{} {
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		log.Println("Endpoint makeGetSaleByUser")
		UserID := auth.GetUserID(ctx)
		resp, err := s.ListCommisionService(UserID)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/list/commision list commision
 * @apiVersion 1.0.2
 * @apiName Commision list commision
 * @apiGroup Commision
 * @apiDescription สำรหับ list บิล ค่าคอม
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token eb109a1cd4d04aeda41c8aad2f7997f3 (token ทดสอบ ของ user)
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *    "response": "success",
 *    "message": "",
 *    "data": [
 *        {
 *            "id": 29,
 *            "user_id": "62110700001",
 *            "sales_code": "SAL62103100002",
 *            "f_name": "สุริยา",
 *            "l_name": "เทียมแย้ม",
 *            "telephone": "+66987727968",
 *            "doc_no": "WDC62113000001",
 *            "doc_date": "2019-11-30 00:00:00",
 *            "total_amount": 8805.96,
 *            "all_sale_commision": 88059.6,
 *            "bank_info": {
 *                "card_id": "1709800118596",
 *                "url_card_id": "url-",
 *                "book_bank_id": "123772710202",
 *                "url_book_bank": "url2",
 *                "bank_code": "",
 *                "bank_name": ""
 *            },
 *            "status": 0,
 *            "my_description": "",
 *            "slip_approve": "",
 *            "confirm": 0,
 *            "confirm_time": "1990-01-01 00:00:00",
 *        }
 *    ]
 *}
 */

func makeGetSaleByUser(s Service) interface{} {
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		log.Println("Endpoint makeGetSaleByUser")
		UserID := auth.GetUserID(ctx)
		resp, err := s.GetSaleByUser(UserID)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/get/sale/byuser get sale by user
 * @apiVersion 1.0.2
 * @apiName Get Sale by user
 * @apiGroup Sales
 * @apiDescription สำรหับ Get ข้อมูลการสมัค sales
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token ee26c312a27b4cc8ba403c3ddc53323b (token ทดสอบ ของ user)
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": {
 *         "id": 1,
 *         "user_id": "62100700002",
 *         "sales_code": "SAL62100700001",
 *         "first_name": "สุริยา",
 *         "invite_person": 1,
 *         "last_name": "เทียมแย้ม",
 *         "role_id": 2,
 *         "role_name": "sales",
 *         "card_id": "1709800118596",
 *         "book_bank_id": "123772710202",
 *         "age": 0,
 *         "url_card_id": "url-",
 *         "url_book_bank": "url2",
 *		   "url_slip_first_buy":"url3",
 *         "file_url_3": "",
 *         "invite_code": "4c3f831ba4aa47f8840ef64285b28e13",
 *         "confirm": 1,
 * 		   "bank_code": "bkb",
 * 		   "bank_name": "ธนาคารกรุงเทพ",
 *         "invite": [
 *             {
 *                 "company_id": 0,
 *                 "role_id": 1,
 *                 "role_name": "user",
 *                 "profix_name": "",
 *                 "fname": "satit",
 *                 "lname": "Chomwattana",
 *                 "active_status": 1,
 *                 "telephone": "+66816398388",
 *                 "email": "satit@extensionsoft.biz",
 *                 "ref_code": "4c3f831ba4aa47f8840ef64285b28e13",
 *                 "invite_code": "5a40d7a444f846de8c16d3668a75e576",
 *                 "invite_person": 2
 *             }
 *         ]
 *     }
 * }
 */

func makeGetSalesPersonEndpoint(s Service) interface{} {
	type request struct {
		SalesCode string `json:"sales_code"`
	}
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpoint makeGetSaleByUser", req)
		resp, err := s.GetSalesCodeService(req.SalesCode)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/get/sales  get sales
 * @apiVersion 1.0.0
 * @apiName Get sales
 * @apiGroup Sales
 * @apiDescription สำรหับ Get sales
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token c0c05dd13ac8490e934bd361e24b724a (token ทดสอบ ของ admin)
 * @apiParam (Parameter) {String} sales_code รหัส sales
 * @apiParamExample {json} Body request:
 * {
 * 	"sales_code":"SAL62111300001"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": {
 *         "id": 69,
 *         "user_id": "62110700001",
 *         "sales_code": "SAL62111300001",
 *         "first_name": "สุริยา",
 *         "invite_person": 0,
 *         "last_name": "เทียมแย้ม",
 *         "role_id": 2,
 *         "role_name": "sales",
 *         "card_id": "1709800118596",
 *         "book_bank_id": "123772710202",
 *         "age": 0,
 *         "url_card_id": "url-",
 *         "url_book_bank": "url2",
 *         "url_slip_first_buy": "124124124",
 *         "file_url_3": "12312312",
 *         "invite_code": "2c6fae0e6e6443b487a4e66066d45566",
 *         "confirm": 1,
 *         "bank_code": "bkb",
 *         "bank_name": "ธนาคารกรุงเทพ",
 *         "create_date": "2019-11-13 20:33:05"
 *     }
 * }
 */

func makeConfrimSaleEndpoint(s Service) interface{} {
	type request struct {
		SaleCode string `json:"sales_code"`
		UserID   string `json:"user_id"`
	}
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpoint makeConfrimSaleEndpoint", req)
		AdminID := auth.GetAdminID(ctx)
		resp, err := s.ConfirmSaleService(AdminID, req.SaleCode, req.UserID)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/confirm/sales  confirm sales
 * @apiVersion 1.0.0
 * @apiName Confirm sales
 * @apiGroup Sales
 * @apiDescription สำรหับ confirm sales
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token c0c05dd13ac8490e934bd361e24b724a (token ทดสอบ ของ admin)
 * @apiParam (Parameter) {String} sales_code รหัส sales
 * @apiParam (Parameter) {string} user_id รหัส ของ user 62083000013
 * @apiParamExample {json} Body request:
 * {
 * 	"sales_code":"SAL62100500001",
 * 	"user_id":"62083000013"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": {
 *         "id": 20,
 *         "user_id": "62083000013",
 *         "sales_code": "SAL62100500001",
 *         "first_name": "สุริยา",
 *         "last_name": "เทียมแย้ม",
 *         "role_id": 2,
 *         "role_name": "sales",
 *         "card_id": "1709800118596",
 *         "book_bank_id": "123772710202",
 *         "age": 0,
 *         "url_card_id": "url-",
 *         "url_book_bank": "url2",
 *		   "url_slip_first_buy":"url3"
 *         "file_url_3": "",
 *         "invite_code": "261cf6fefbf64c209853ba1718066b50",
 *         "confirm": 1
 *     }
 * }
 */

func makeSaleConfirmList(s Service) interface{} {
	type request struct {
		Type  string `json:"type"`
		Limit int64  `json:"limit"`
	}
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpooint makeSaleConfirmList :", req)
		resp, err := s.ShowlistConfirmSales(req.Type, req.Limit)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /sales/v1/list/sales  list Sales
 * @apiVersion 1.0.0
 * @apiName Sales list
 * @apiGroup Sales
 * @apiDescription List sales
 * @apiHeader Content-Type application/json
 * @apiParam (Parameter) {String} type ประเภทการ search ALL=ทั้งหมด|CONFIRMED=confirm แล้ว |NOCONFIRMED = ที่ยังไม่ confirm
 * @apiParam (Parameter) {string} limit จำนวนการ search
 * @apiParamExample {json} Body request:
 * {
 *	"type":"ALL",
 *	"limit":2
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *    "response": "success",
 *    "message": "",
 *    "data": {
 *        "length": 13,
 *        "sale_list": [
 *            {
 *                "id": 7,
 *                "user_id": "62101000006",
 *                "sales_code": "SAL62101000002",
 *                "first_name": "ธันยนันท์",
 *                "invite_person": 13,
 *                "last_name": "เตียวเจริญโสภา",
 *                "role_id": 1,
 *                "role_name": "user",
 *                "card_id": "3329900003880",
 *                "book_bank_id": "0141277774",
 *                "age": 0,
 *                "url_card_id": "Choose file",
 *                "url_book_bank": "https://perfectapi.extensionsoft.biz/image/sales/sales-532137107.png",
 *                "url_slip_first_buy": "",
 *                "file_url_3": "",
 *				  "telephone":"+6651231231",
 *                "invite_code": "d55ec6357e3048c0aff18760c9982b6d",
 *                "confirm": 0,
 *				  "bank_code": "bkb",
 *				  "bank_name": "ธนาคารกรุงเทพ",
 * 				  "create_date": "2019-10-31 07:55:57"
 *            },
 *            {
 *                "id": 1,
 *                "user_id": "62100900003",
 *                "sales_code": "SAL62100900001",
 *                "first_name": "ชื่อจริง",
 *                "invite_person": 0,
 *                "last_name": "นามสกุลจริง",
 *                "role_id": 2,
 *                "role_name": "sales",
 *                "card_id": "1234567891325",
 *                "book_bank_id": "1234671483",
 *                "age": 0,
 *                "url_card_id": "Choose file",
 *                "url_book_bank": "Choose file",
 *                "url_slip_first_buy": "",
 *                "file_url_3": "",
 *				  "telephone":"+6651231231",
 *                "invite_code": "3f28bd9e951543a583c0ea02e73a1e38",
 *                "confirm": 1,
 *				  "bank_code": "bkb",
 *				  "bank_name": "ธนาคารกรุงเทพ",
 * 				  "create_date": "2019-10-31 07:55:57"
 *            }
 *        ],
 *        "type_list": "ALL"
 *    }
 *
 */
