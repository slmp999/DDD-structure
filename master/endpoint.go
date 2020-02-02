package master

import (
	"context"
	"fmt"
)

//  FindBankEndpoint
// for list of all bank
func FindBankEndpoint(s Service) interface{} {

	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		// UserID := auth.GetUserID(ctx)
		fmt.Println("get bank list: ")
		resp, err := s.BankList()
		if err != nil {
			fmt.Println("MakeBankListEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}

		return &response{
			Result: "success",
			Data:   resp,
		}, nil
	}
}

/**
* @api {post} master/v1/get/bank  check coupon
* @apiVersion 1.0.2
* @apiName bank list
* @apiGroup Master
* @apiDescription สำรหับ get bank list
* @apiHeader Content-Type application/json
* @apiHeader Access-Token cc7fcd56c99847bd9053a264c60527b4 (token ทดสอบ ของ user)
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} Message ข้อความตอบกลับ
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "message": "",
*     "data": [{
*         "id": 47,
*         "bank_code": "BBL",
*         "bank_name": "ธนาคารกรุง",
*     },
		{ ...}
		]
* }
* @apiErrorExample {json} coupon_expire:
* {
*     "response": "false",
*     "message": "bank list error ",
*     "data": null
* }
*/
