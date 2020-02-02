package admin

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"gitlab.com/satit13/perfect_api/auth"
)

func makeUpdateStatusAdmin(s Service) interface{} {
	type request struct {
		AdminID      string `json:"admin_id"`
		AcitveStatus int64  `json:"active_status"`
	}
	type response struct {
		Result  string `json:"response"`
		Message string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		admincode := auth.GetAdminID(ctx)
		_, err := s.UpdateStatusAdminService(req.AdminID, req.AcitveStatus, admincode)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		message := ""
		if req.AcitveStatus != 1 {
			message = "ลบข้อมูลผู้ใช้งานสำเร็จ"
		} else {
			message = "เปิดใช้งานผู้ใช้งานสำเร็จ"
		}
		return &response{
			Result:  "success",
			Message: message,
		}, nil
	}
}

/**
 * @api {post} /admin/v1/remove/admin/adminid remove admin by admin id
 * @apiVersion 1.0.0
 * @apiName remove admin by admin id
 * @apiGroup Admin
 * @apiDescription สำรหับ remove admin by admin id
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 1ce991f5305d42bdb312d29107138cec (token ทดสอบ)
 * @apiParam (Parameter) {String} admin_id
 * @apiParamExample {json} Body request:
 * {
 * 		"admin_id":"BN62092400001"
 * 		"active_status":0
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "ลบข้อมูลผู้ใช้งานสำเร็จ",
 */

func makeAdminByAdminID(s Service) interface{} {
	type request struct {
		AdminID string `json:"admin_id"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {

		resp, err := s.GetAdminByUserID(req.AdminID)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result:  "success",
			Message: "บันทึกข้อมูลผู้ใช้งานสำเร็จ",
			Data:    resp,
		}, nil
	}
}

/**
 * @api {post} /admin/v1/get/admin/adminid get admin by admin id
 * @apiVersion 1.0.0
 * @apiName get admin by admin id
 * @apiGroup Admin
 * @apiDescription สำรหับ get admin by admin id
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 1ce991f5305d42bdb312d29107138cec (token ทดสอบ)
 * @apiParam (Parameter) {String} admin_id
 * @apiParamExample {json} Body request:
 * {
 * 		"admin_id":"BN62092400001"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "บันทึกข้อมูลผู้ใช้งานสำเร็จ",
 *     "data": {
 *         "admin_id": "BN62092400001",
 *         "company_id": 0,
 *         "role_id": 1,
 *         "role_name": "user",
 *         "prefix_name": "",
 *         "fname": "admin",
 *         "lname": "test",
 *         "active_status": 1,
 *         "phone": "0979679089",
 *         "email": "nziomin1i2322312@gmail.com",
 *         "password": ""
 *     }
 * }
 */
func makeEditAdmin(s Service) interface{} {
	type request UpdateUserProfileModel
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		adminID := auth.GetAdminID(ctx)
		User := UpdateUserProfileModel{
			AdminID:      req.AdminID,
			ProfixName:   req.ProfixName,
			RoleID:       req.RoleID,
			Fname:        req.Fname,
			Lname:        req.Lname,
			ActiveStatus: req.ActiveStatus,
			Phone:        req.Phone,
			Email:        req.Email,
			Password:     req.Password,
		}
		_, err := s.UpdateprofileBackEndService(&User, adminID)

		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result:  "success",
			Message: "บันทึกข้อมูลผู้ใช้งานสำเร็จ",
			Data:    User,
		}, nil
	}
}

/**
* @api {post} /admin/v1/update/admin Update Admin
* @apiVersion 1.0.0
* @apiName update profile
* @apiGroup Admin
* @apiDescription สำรหับ เพิ่มผูใช่งานหลังบ้าน
* @apiHeader Content-Type application/json
* @apiHeader Access-Token 1ce991f5305d42bdb312d29107138cec (token ทดสอบ)
* @apiParam (Parameter) {String} admin_id
* @apiParam (Parameter) {Int} company_id
* @apiParam (Parameter) {Int} role_id
* @apiParam (Parameter) {String} profix_name
* @apiParam (Parameter) {String} fname
* @apiParam (Parameter) {String} lname
* @apiParam (Parameter) {String} phone
* @apiParam (Parameter) {int} active_status
* @apiParam (Parameter) {String} email
* @apiParam (Parameter) {String} password
* @apiParamExample {json} Body request:
* {
 *
 *		  "admin_id":"BN62092400001",
 *	   	   "company_id": 0,
 *		   "role_id":1,
 *        "profix_name": "mr",
 *        "fname": "admin",
 *        "lname": "test",
 *        "active_status": 1,
 *        "phone": "0979679089",
 *        "email": "nziomin1i232@gmail.com",
 *		  "password" :"1234"
 *
* }
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} Message ข้อความตอบกลับ
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "message": "บันทึกข้อมูลผู้ใช้งานสำเร็จ",
*     "data": {
*         "company_id": 0,
*         "profix_name": "mr",
*         "fname": "admin",
*         "lname": "test",
*         "active_status": 1,
*         "phone": "0979679089",
*         "email": "nziomin1i@gmail.com"
*     }
* }
*/

func makeAddBackUser(s Service) interface{} {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleID   int64  `json:"role_id"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		adminID := auth.GetAdminID(ctx)
		resp, err := s.AddBackUserService(adminID, req.Email, req.Password, req.RoleID)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result:  "success",
			Message: "เพิ่มผู้ใช้งานสำเร็จ",
			Data:    resp,
		}, nil
	}
}

func makeAddAdminV2Endpoint(s Service) interface{} {
	type request AddadminBackEndModel
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		adminID := auth.GetAdminID(ctx)

		User := AddadminBackEndModel{
			ProfixName: req.ProfixName,
			RoleID:     req.RoleID,
			Fname:      req.Fname,
			Lname:      req.Lname,
			Phone:      req.Phone,
			Email:      req.Email,
			Password:   req.Password,
		}
		resp, err := s.AddAdminServiceV2(adminID, &User)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result:  "success",
			Message: "เพิ่มผู้ใช้งานสำเร็จ",
			Data:    resp,
		}, nil
	}
}

/**
 * @api {post} /admin/v1/add/admin  add admin
 * @apiVersion 1.0.0
 * @apiName add user admin
 * @apiGroup Admin
 * @apiDescription สำรหับ เพิ่มผูใช่งานหลังบ้าน
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 1ce991f5305d42bdb312d29107138cec (token ทดสอบ)
 * @apiParam (Parameter) {String} email อีเมล์
 * @apiParamExample {json} Body request:
 * {
 * 	   "company_id": 0,
 *         "profix_name": "mr",
 *         "fname": "admin",
 *         "lname": "test",
 *         "role_id": 5,
 *         "phone": "0979679089",
 *         "email": "admin2@gmail.com",
 *         "password":"1234"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "เพิ่มผู้ใช้งานสำเร็จ",
 *     "data": {
 *         "admin_code": "BN63010300002",
 *         "email": "admin2@gmail.com",
 *         "f_name": "admin",
 *         "l_name": "test",
 *         "password": "1234",
 *         "pre_fix": "",
 *         "role_id": 5
 *     }
 * }
 */

func makeAddAdminEndpoint(s Service) interface{} {
	type request AddadminBackEndModel
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		adminID := auth.GetAdminID(ctx)
		resp, err := s.AddAdminService(adminID, req.Email, req.Password, req.RoleID)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result:  "success",
			Message: "เพิ่มผู้ใช้งานสำเร็จ",
			Data:    resp,
		}, nil
	}
}

func makeChangePasswordAdmin(s Service) interface{} {
	type request struct {
		Email       string `json:"email"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		adminID := auth.GetAdminID(ctx)
		resp, err := s.GetProfileAdminService(adminID)
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

func makeListRole(s Service) interface{} {
	type request struct {
		Status int64 `json:"status"`
	}
	type response struct {
		Response string      `json:"response"`
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		log.Println("Endpooint MakeEndpointApprove :", req)
		resp, err := s.ListRoleService(req.Status)
		if err != nil {
			return &response{Response: "false", Message: err.Error()}, nil
		}
		return &response{
			Response: "success",
			Message:  "",
			Data:     resp,
		}, nil
	}
}

/**
 * @api {post} /admin/v1/list/role  list role
 * @apiVersion 1.0.0
 * @apiName list role
 * @apiGroup Admin
 * @apiDescription สำรหับ เพิ่มผูใช่งานหลังบ้าน
 * @apiHeader Content-Type application/json
 * @apiParam (Parameter) {String} status 0 -> หน้าบ้าน | 1 -> หลังบ้าน | 2 -> ทั้งหมด
 * @apiParamExample {json} Body request:
 * {
 * 	"status":1
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": [
 *         {
 *             "role_id": "1",
 *             "role_name": "user",
 *             "active_status": 0,
 *             "type": 0
 *         },
 *         {
 *             "role_id": "2",
 *             "role_name": "sales",
 *             "active_status": 0,
 *             "type": 0
 *         },
 *         {
 *             "role_id": "3",
 *             "role_name": "dealer",
 *             "active_status": 0,
 *             "type": 0
 *         }
 *     ]
 * }
 */
func MakeEndpointListAdmin(s Service) interface{} {
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
		len, resp, err := s.ListAdminService(req.Status, req.Keyword, req.Limit)
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
 * @api {post} /admin/v1/list/admin  list user admin
 * @apiVersion 1.0.0
 * @apiName list user admin
 * @apiGroup Admin
 * @apiDescription สำรหับ list admin
 * @apiHeader Content-Type application/json
 * @apiParam (Parameter) {String} keyword
 * @apiParam (Parameter) {String} limit
 * @apiParam (Parameter) {String} status status 0 -> ยกเลิก 1 -> ใช้งาน 2-> ทั้งหมด
 * @apiParamExample {json} Body request:
 *  {
 *  	"keyword":"",
 *  	"limit":5,
 *  	"status":1
 *  }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "length": 5,
 *     "data": [
 *         {
 *             "admin_id": "BN62092400001",
 *             "company_id": 0,
 *             "role_id": 1,
 *             "role_name": "user",
 *             "profix_name": "mr",
 *             "fname": "admin",
 *             "lname": "test",
 *             "active_status": 1,
 *             "phone": "0979679089",
 *             "email": "user_test@test.com",
 *             "menu": null
 *         },
 *         {
 *             "admin_id": "BN62122800001",
 *             "company_id": 0,
 *             "role_id": 1,
 *             "role_name": "user",
 *             "profix_name": "",
 *             "fname": "",
 *             "lname": "",
 *             "active_status": 1,
 *             "phone": "",
 *             "email": "nziomini1@gmail.com",
 *             "menu": null
 *         },
 *         {
 *             "admin_id": "BN62122800002",
 *             "company_id": 0,
 *             "role_id": 1,
 *             "role_name": "user",
 *             "profix_name": "",
 *             "fname": "",
 *             "lname": "",
 *             "active_status": 1,
 *             "phone": "",
 *             "email": "nziomini2@gmail.com",
 *             "menu": null
 *         },
 *         {
 *             "admin_id": "BN62122800003",
 *             "company_id": 0,
 *             "role_id": 1,
 *             "role_name": "user",
 *             "profix_name": "",
 *             "fname": "",
 *             "lname": "",
 *             "active_status": 1,
 *             "phone": "",
 *             "email": "nziomini3@gmail.com",
 *             "menu": null
 *         },
 *         {
 *             "admin_id": "BN62122800004",
 *             "company_id": 0,
 *             "role_id": 5,
 *             "role_name": "owner",
 *             "profix_name": "",
 *             "fname": "",
 *             "lname": "",
 *             "active_status": 1,
 *             "phone": "",
 *             "email": "nziomini4@gmail.com",
 *             "menu": null
 *         }
 *     ]
 * }
 */

func makeGetAdmin(s Service) interface{} {
	// type request struct {
	// 	UserCode string `json:"user_id"`
	// }
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		adminID := auth.GetAdminID(ctx)
		log.Println(adminID)
		resp, err := s.GetProfileAdminService(adminID)
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

/**
 * @api {post} /admin/v1/get/admin
 * @apiVersion 1.0.0
 * @apiName get Admin
 * @apiGroup Admin
 * @apiDescription ดึงข้อมูลขอ user หลังบ้าน
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 1ce991f5305d42bdb312d29107138cec (token ทดสอบ)
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} data ข้อมูลของ api ที่ตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 * "response": "success",
 * "message": "",
 * "data": {
 *       "company_id": 0,
 *       "role_id": 1,
 *       "role_name": "user",
 *       "profix_name": "",
 *       "fname": "",
 *       "lname": "",
 *       "active_status": 1,
 *       "phone": "",
 *       "email": "nziomin1i@gmail.com",
 *       "menu": [
 *           {
 *               "head_menu_id": 1,
 *               "head_manu_name": "การสั่งซื้อ",
 *               "head_menu_icon": "form",
 *               "head_menu_line_number": 1,
 * 				 "noti" : 0 ,
 *               "sub": [
 *                   {
 *                       "menu_id": 1,
 *                       "head_menu_id": 1,
 *                       "menu_name": "จัดการรายการสั่งซื้อ",
 *                       "menu_link": "/orders",
 *                       "menu_show": 0,
 *                       "menu_status": 0,
 *                       "menu_line_number": 1,
 * 						 "noti" : 0 ,
 *                   }
 *               ]
 *           }
 *       ]
 *   }
 * }
 */

func makeSigninAdmin(s Service) interface{} {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {

		resp, err := s.SignInAdminService(req.Email, req.Password)
		if err != nil {
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result: "success",
			Data:   resp,
		}, nil
	}
}

/**
 * @api {post} /admin/v1/signin  Login
 * @apiVersion 1.0.0
 * @apiName Signin
 * @apiGroup Admin
 * @apiDescription สำรหับ login ระบบหลังบ้าน
 * @apiHeader Content-Type application/json
 * @apiParam (Parameter) {String} email อีเมล์
 * @apiParam (Parameter) {string} password พาสเวิร์ด
 * @apiParamExample {json} Body request:
 * {
 *   "email": "user_test@test.com",
 *   "password": "123456"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *   "response": "success",
 *   "data": {
 *     "token_1": "efcefb98-2f4c-43b9-a7c4-e836d55a6bc4",
 *     "token_2": "efcefb98-2f4c-43b9-a7c4-e836d55a6bc2"
 *   }
 * }
 */

func makeSignOutAdmin(s Service) interface{} {
	type response struct {
		Result  string `json:"response"`
		Message string `json:"message"`
	}
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		fmt.Println("begin endpoint.makeSignoutEndpoint")
		xat := req.Header.Get("Access-Token")
		_, err := s.Signout(xat)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{Result: "success", Message: "Logout สำเร็จ"}, nil
	}
}

/**
 * @api {post} /admin/v1/signout SignOut
 * @apiVersion 1.0.2
 * @apiName SignOut
 * @apiGroup Admin
 * @apiDescription สำรหับ logout
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token ee26c312a27b4cc8ba403c3ddc53323b (token ทดสอบ ของ user)
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "logout สำเร็จ",
 * }
 */
