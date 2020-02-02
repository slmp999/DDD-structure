package report

/**
* @api {post} /report/v1/report/sale  reportsale
* @apiVersion 1.0.0
* @apiName reportsale
* @apiGroup report
* @apiDescription สำรหับ report
* @apiHeader Content-Type application/json
* @apiHeader Access-Token c0c05dd13ac8490e934bd361e24b724a (token ทดสอบ)
* @apiParam (Parameter) {String} doc_no รหัส เอกสาร
* @apiParamExample {json} Body request:
*  {
*   	"start_date":"2019-11-05",
*   	"end_date":"2019-11-05"
*  }
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} Message ข้อความตอบกลับ
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "message": "",
*     "data": {
* 		"date": "2019-11-05 --> 2019-11-05",
*         "sum_amount": 110829
*         "length": 8,
*         "list": [
*             {
*                 "doc_no": "ORD621105090718",
*                 "user_id": "62101600002",
*                 "user_name": "สมรถ หลักฐาน",
*                 "create_time": "2019-11-05 02:07:18",
*                 "commision_percent": 10,
*                 "commision_amount": 1259.9,
*                 "amount": 12599,
*                 "sale_type": "ธนาคาร"
*             },
*             {
*                 "doc_no": "ORD621105091418",
*                 "user_id": "62102500003",
*                 "user_name": "สมปอง หลักฐาน",
*                 "create_time": "2019-11-05 02:14:19",
 *                 "commision_percent": 10,
*                 "commision_amount": 2874.7,
*                 "amount": 28747,
*                 "sale_type": "ธนาคาร"
*             },
*             {
*                 "doc_no": "ORD621105094241",
*                 "user_id": "62102500002",
*                 "user_name": "สมรัก หลักฐาน",
*                 "create_time": "2019-11-05 02:42:42",
 *                 "commision_percent": 10,
*                 "commision_amount": 1535.0,
*                 "amount": 15350,
*                 "sale_type": "ธนาคาร"
*             },
*             {
*                 "doc_no": "ORD621105215648",
*                 "user_id": "62103100001",
*                 "user_name": "satit chomwattana",
*                 "create_time": "2019-11-05 14:56:49",
 *                 "commision_percent": 10,
*                 "commision_amount": 3000,
*                 "amount": 30000,
*                 "sale_type": "ธนาคาร"
*             },
*             {
*                 "doc_no": "ORD621105220741",
*                 "user_id": "62110400001",
*                 "user_name": "jidapa chomwattana",
*                 "create_time": "2019-11-05 15:07:41",
 *                 "commision_percent": 10,
*                 "commision_amount": 725,
*                 "amount": 7250,
*                 "sale_type": "ธนาคาร"
*             },
*             {
*                 "doc_no": "ORD621105221426",
*                 "user_id": "62102500005",
*                 "user_name": "ปองศักดิ์ หลักฐาน",
*                 "create_time": "2019-11-05 15:14:26",
 *                 "commision_percent": 10,
*                 "commision_amount": 1165,
*                 "amount": 11650,
*                 "sale_type": "ธนาคาร"
*             },
*             {
*                 "doc_no": "ORD621105221629",
*                 "user_id": "62102500005",
*                 "user_name": "ปองศักดิ์ หลักฐาน",
*                 "create_time": "2019-11-05 15:16:29",
 *                 "commision_percent": 10,
*                 "commision_amount": 365.4,
*                 "amount": 3654,
*                 "sale_type": "ธนาคาร"
*             },
*             {
*                 "doc_no": "ORD621105224655",
*                 "user_id": "62101500001",
*                 "user_name": "nathaphol wichonit",
*                 "create_time": "2019-11-05 15:46:56",
 *                 "commision_percent": 10,
*                 "commision_amount": 157.9,
*                 "amount": 1579,
*                 "sale_type": "ธนาคาร"
*             }
*         ],
*     }
* }
*/
