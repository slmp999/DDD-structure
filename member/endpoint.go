package member

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gitlab.com/satit13/perfect_api/auth"
)

func CheckCouponByNoEndpoint(s Service) interface{} {
	type request struct {
		CouponNo string `json:"coupon_no"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println("getCouPn response: ", UserID)
		resp, err := s.CheckCouponService(UserID, req.CouponNo)

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
* @api {post} member/v1/check/coupon  check coupon
* @apiVersion 1.0.2
* @apiName check coupon
* @apiGroup Member
* @apiDescription สำรหับ get coupon ของ users
* @apiHeader Content-Type application/json
* @apiHeader Access-Token cc7fcd56c99847bd9053a264c60527b4 (token ทดสอบ ของ user)
 * @apiParam (Parameter) {String} coupon_no รหัสcoupon
* @apiParamExample {json} Body request:
* {
* 	 	 "coupon_no":"CPV621021115524",
* }
* @apiSuccess {String} response สถานะตอบกลับ
* @apiSuccess {String} Message ข้อความตอบกลับ
* @apiSuccessExample {json} Success:
* {
*     "response": "success",
*     "message": "",
*     "data": {
*         "id": 47,
*         "doc_date": "2019-10-21",
*         "coupon_no": "CPV621021115524",
*         "coupon_type": 1,
*         "user_id": "62100900005",
*         "name": "Package Size S",
*         "value": 12000,
*         "remain": 12000,
*         "expire_status": 0,
*         "begin_date": "2019-10-21",
*         "expire_date": "2020-11-21"
*     }
* }
* @apiErrorExample {json} coupon_expire:
* {
*     "response": "false",
*     "message": "coupon หมดอายุการใช้งาน",
*     "data": null
* }
*/

func FindCouponMemberEndpoint(s Service) interface{} {
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println("getCouPn response: ", UserID)
		resp, err := s.GetCouponService(UserID)

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
 * @api {post} member/v1/get/coupon  get coupon
 * @apiVersion 1.0.2
 * @apiName get coupon
 * @apiGroup Member
 * @apiDescription สำรหับ get coupon ของ users
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token cc7fcd56c99847bd9053a264c60527b4 (token ทดสอบ ของ user)
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": [
 *         {
 *             "id": 1,
 *             "doc_date": "2019-10-07",
 *             "coupon_no": "CPV62100700001",
 *             "coupon_type": 2,
 *             "user_id": "62100900005",
 *             "name": "Package Size S",
 *             "value": 15,
 *             "remain": 15,
 *             "expire_status": 1,
 *             "begin_date": "2019-10-07",
 *             "expire_date": "2020-10-07"
 *         },
 *         {
 *             "id": 2,
 *             "doc_date": "2019-10-07",
 *             "coupon_no": "CPP62100700002",
 *             "coupon_type": 1,
 *             "user_id": "62100900005",
 *             "name": "Package Size S",
 *             "value": 12000,
 *             "remain": 12000,
 *             "expire_status": 1,
 *             "begin_date": "2019-10-07",
 *             "expire_date": "2010-10-07"
 *         }
 *     ]
 * }
 */

func getUserByAdmin(s Service) interface{} {
	type request struct {
		UserID string `json:"user_id"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {

		resp, err := s.GetUserByAdmin(req.UserID)

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
 * @api {post} /member/v1/get/user/admin Get user By Admin
 * @apiVersion 1.0.2
 * @apiName Get user By Admin
 * @apiGroup Member
 * @apiDescription สำรหับ Get ข้อมูล ของ user สำหรับเว็บหลังบ้าน
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token c0c05dd13ac8490e934bd361e24b724a (token ทดสอบ ของ Admin)
 * @apiParam (Parameter) {string} user_id รหัส ของ user 62083000013
 * @apiParamExample {json} Body request:
 * {
 * 	"user_id":"62083000013"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": {
 *         "company_id": 0,
 *         "role_id": 1,
 *         "role_name": "user",
 *         "dealer": 0,
 *         "profix_name": "",
 *         "sales_status": 0,
 *         "fname": "นาย",
 *         "lname": "เทียมแย้ม",
 *         "first_buy": 0,
 *         "active_status": 1,
 *         "telephone": "+66987727968",
 *         "email": "",
 *         "discount_person": "",
 *         "discount_expridate": "",
 *         "ref_code": "",
 *         "invite_code": "c6557aae878d46a2adfe6dfb593ccbac",
 *         "invite_person": 0,
 *         "create_date": "2019-11-14 05:40:01"
 *     }
 * }
 * @apiErrorExample {json} token_expire:
 * {
 *   "response": "false"
 *   "message":"token_expire"
 *   "error":"token_expire"
 * }
 */
func makeRemoveUserByAdmin(s Service) interface{} {
	type request struct {
		UserID string `json:"user_id"`
	}
	type response struct {
		Result  string `json:"response"`
		Message string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		AdminID := auth.GetAdminID(ctx)
		fmt.Println("makeRemoveUserByAdmin response: ", AdminID)
		_, err := s.RemoveUserServiceAdmin(AdminID, req.UserID)

		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}

		return &response{
			Result:  "success",
			Message: "ลบข้อมูลสำเร็จ",
		}, nil
	}
}

/**
 * @api {post} /member/v1/remove/user/admin  remove user by admin
 * @apiVersion 1.0.0
 * @apiName remove user by admin
 * @apiGroup Member
 * @apiDescription สำรหับ remove user by admin
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token c0c05dd13ac8490e934bd361e24b724a (token ทดสอบ ของ admin)
 * @apiParam (Parameter) {string} user_id รหัส ของ user 62083000013
 * @apiParamExample {json} Body request:
 * {
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

func FindProfileMemberEndpoint(s Service) interface{} {
	// type request struct {
	// 	UserCode string `json:"user_id"`
	// }
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println("GetProfileMemberService response: ", UserID)
		resp, err := s.GetProfileMemberService(UserID)

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
 * @api {post} /member/v1/get/user Get user
 * @apiVersion 1.0.2
 * @apiName Get user
 * @apiGroup Member
 * @apiDescription สำรหับ Get ข้อมูล ของ user
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e (token ทดสอบ ของ user)
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": {
 *         "company_id": 0,
 *         "role_id": 2,
 *         "role_name": "sales",
 *         "dealer": 0,
 *         "profix_name": "",
 *         "sales_status": 0,
 *         "fname": "test",
 *         "lname": "test",
 *         "first_buy": 0,
 *         "active_status": 1,
 *         "telephone": "+66987727968",
 *         "email": "sfe@fefsf",
 *         "discount_person": "15%",
 *         "discount_expridate": "2020-10-07",
 *         "ref_code": "",
 *         "invite_code": "77d6825fab454600a1301586569b4df4",
 *         "invite_person": 2,
 *         "invite": [
 *             {
 *                 "company_id": 0,
 *                 "role_id": 2,
 *                 "role_name": "sales",
 *                 "profix_name": "",
 *                 "fname": "nathaphol",
 *                 "lname": "wichonit",
 *                 "active_status": 1,
 *                 "telephone": "+66848703805",
 *                 "email": "nathaphol008@gmail.com",
 *                 "ref_code": "77d6825fab454600a1301586569b4df4",
 *                 "invite_code": "2ea4b6de37934bdfbd7ca595235af7b6",
 *                 "invite_person": 0
 *             },
 *             {
 *                 "company_id": 0,
 *                 "role_id": 2,
 *                 "role_name": "sales",
 *                 "profix_name": "",
 *                 "fname": "satit",
 *                 "lname": "chomwattana",
 *                 "active_status": 1,
 *                 "telephone": "+66816398388",
 *                 "email": "satit@extensionsoft.biz",
 *                 "ref_code": "77d6825fab454600a1301586569b4df4",
 *                 "invite_code": "c7c501c8edf346c08c510f807e1e166d",
 *                 "invite_person": 1,
 *                 "invite": [
 *                     {
 *                         "company_id": 0,
 *                         "role_id": 2,
 *                         "role_name": "sales",
 *                         "profix_name": "",
 *                         "fname": "tanakorn",
 *                         "lname": "phuntulap",
 *                         "active_status": 1,
 *                         "telephone": "+66613351889",
 *                         "email": "destiny00005@gmail.com",
 *                         "ref_code": "c7c501c8edf346c08c510f807e1e166d",
 *                         "invite_code": "b0c71303730daas795a9707c7a8a73f1",
 *                         "invite_person": 0
 *                     }
 *                 ]
 *             }
 *         ],
 *         "member_address": [],
 *         "menu": [
 *             {
 *                 "head_menu_id": 1,
 *                 "head_manu_name": "การสั่งซื้อ",
 *                 "head_menu_icon": "form",
 *                 "head_menu_line_number": 1,
 *                 "sub": []
 *             },
 *             {
 *                 "head_menu_id": 2,
 *                 "head_manu_name": "ตัวแทนจำหน่าย",
 *                 "head_menu_icon": "form",
 *                 "head_menu_line_number": 2,
 *                 "sub": []
 *             }
 *         ],
 *         "list_coupon": [
 *             {
 *                 "id": 1,
 *                 "doc_date": "2019-10-07",
 *                 "coupon_no": "CPV62100700001",
 *                 "coupon_type": 2,
 *                 "user_id": "62100900005",
 *                 "name": "Package Size S",
 *                 "value": 15,
 *                 "remain": 15,
 *                 "expire_status": 1,
 *                 "begin_date": "2019-10-07",
 *                 "expire_date": "2020-10-07"
 *             },
 *             {
 *                 "id": 2,
 *                 "doc_date": "2019-10-07",
 *                 "coupon_no": "CPP62100700002",
 *                 "coupon_type": 1,
 *                 "user_id": "62100900005",
 *                 "name": "Package Size S",
 *                 "value": 12000,
 *                 "remain": 12000,
 *                 "expire_status": 1,
 *                 "begin_date": "2019-10-07",
 *                 "expire_date": "2010-10-07"
 *             }
 *         ]
 *     }
 * }
 * @apiErrorExample {json} token_expire:
 * {
 *   "response": "false"
 *   "message":"token_expire"
 *   "error":"token_expire"
 * }
 */

func UpdateProfileMemberEndpoint(s Service) interface{} {
	type request UpdateUserProfileModel
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		User := UpdateUserProfileModel{
			RoleID:     req.RoleID,
			ProfixName: req.ProfixName,
			Fname:      req.Fname,
			Lname:      req.Lname,
			// // UserID:       UserID,
			// UserName:     req.UserName,
			// EngName:      req.EngName,
			ActiveStatus: req.ActiveStatus,
			Telephone:    req.Telephone,
			Email:        req.Email,
		}
		resp, err := s.UpdatePRofileUserService(&User, UserID)
		fmt.Println("makeSigninEndpoint response: ", resp)
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

func UpdateMemberBackEndpoint(s Service) interface{} {
	type request UpdateUserProfileModelByAdmin
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		AdminID := auth.GetAdminID(ctx)
		User := UpdateUserProfileModelByAdmin{
			CompanyID:  req.CompanyID,
			UserID:     req.UserID,
			RoleID:     req.RoleID,
			ProfixName: req.ProfixName,
			Fname:      req.Fname,
			Lname:      req.Lname,
			// // UserID:       UserID,
			// UserName:     req.UserName,
			// EngName:      req.EngName,
			ActiveStatus: req.ActiveStatus,
			Telephone:    req.Telephone,
			Email:        req.Email,
		}
		resp, err := s.UpdateMemberByAdmin(&User, AdminID)
		fmt.Println("makeSigninEndpoint response: ", resp)
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
 * @api {post} /member/v1/update/member/admin  update user by admin
 * @apiVersion 1.0.0
 * @apiName update user by admin
 * @apiGroup Member
 * @apiDescription สำรหับ remove user by admin
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token c0c05dd13ac8490e934bd361e24b724a (token ทดสอบ ของ admin)
 * @apiParam (Parameter) {string} user_id รหัส ของ user 62083000013
 * @apiParamExample {json} Body request:
 * {
 * 	"company_id": 0,
 * 	"user_id":"62111400003",
 * 	"role_id":1,
 *    "profix_name": "mr",
 * 	 "fname": "admin",
 * 	 "lname": "test",
 * 	 "active_status": 1,
 * 	 "phone": "0987727968",
 * 	 "email": "nziomini@gmail.com"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 *{
 *    "response": "success",
 *    "message": "บันทึกข้อมูลผู้ใช้งานสำเร็จ",
 *    "data": {
 *        "user_id": "62111400003",
 *        "company_id": 0,
 *        "role_id": 1,
 *        "profix_name": "mr",
 *        "fname": "admin",
 *        "lname": "test",
 *        "active_status": 1,
 *        "telephone": "+66987727968",
 *        "email": "nziomini@gmail.com"
 *    }
 *}
 */

func FindAddressMemberEndpoint(s Service) interface{} {

	// type request struct {
	// 	UserCode string `json:"user_id"`
	// }
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetUserID(ctx)

		resp, err := s.GetAddressByUserService(UserID)
		if err != nil {
			return &response{Result: "false", Message: err.Error()}, nil
		}

		return &response{
			Result: "success",
			Data:   resp,
		}, nil
	}
}

func FindAddressMemberByIdEndpoint(s Service) interface{} {
	type request struct {
		AddressID int64 `json:"addr_id"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		resp, err := s.GetAddressByIdService(UserID, req.AddressID)
		if err != nil {
			return &response{Result: "false", Message: err.Error()}, nil
		}

		return &response{
			Result: "success",
			Data:   resp,
		}, nil
	}
}

func makeAddProfileAddressEndpoint(s Service) interface{} {
	type request AddressProfileModel
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println("begin endpoint.makeSigninEndpoint")
		profile := AddAddressModel{
			UserID: UserID,
			Address: AddressProfileModel{
				AddressID:   req.AddressID,
				Name:        req.Name,
				Phone:       req.Phone,
				Email:       req.Email,
				Address:     req.Address,
				SubArea:     req.SubArea,
				Province:    req.Province,
				District:    req.District,
				PostalCode:  req.PostalCode,
				MainAddress: req.MainAddress,
			},
		}
		resp, err := s.AddProfileAddressService(&profile)
		fmt.Println("makeSigninEndpoint response: ", resp)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result: "success",
			Data:   profile,
		}, nil
	}
}

func makeUploadimage(s Service) interface{} {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r.ParseMultipartForm(10 << 30)
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, handler, err := r.FormFile("myFile")

		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return nil, err
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		tempFile, err := ioutil.TempFile("/app/temp-images", "upload-*.png")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		// return that we have successfully uploaded our file!
		return nil, nil
	}

}

func makeUpdateAddressEndpoint(s Service) interface{} {
	type request AddressProfileModel
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeSigninEndpoint")

		UserID := auth.GetUserID(ctx)
		log.Println(UserID)
		profile := AddAddressModel{
			UserID: UserID,
			Address: AddressProfileModel{
				AddressID:   req.AddressID,
				Name:        req.Name,
				Phone:       req.Phone,
				Email:       req.Email,
				Address:     req.Address,
				SubArea:     req.SubArea,
				Province:    req.Province,
				District:    req.District,
				PostalCode:  req.PostalCode,
				MainAddress: req.MainAddress,
			},
		}

		resp, err := s.UpdateProfileAddress(&profile)
		fmt.Println("makeSigninEndpoint response: ", resp)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result: "success",
			Data:   profile,
		}, nil
	}
}

func makeDeleteProfileEndpoint(s Service) interface{} {
	type request struct {
		AddressID int64 `json:"addr_id"`
	}
	type response struct {
		Result  string `json:"response"`
		Message string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		UserID := auth.GetUserID(ctx)
		fmt.Println("begin endpoint.makeSigninEndpoint")

		resp, err := s.DeleteProfileAddressService(UserID, req.AddressID)
		fmt.Println("makeSigninEndpoint response: ", resp)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result:  "success",
			Message: "ลบข้อมูลสำเร็จ",
		}, nil
	}
}

func FindAllMember(s Service) interface{} {
	// type request struct {
	// 	UserCode string `json:"user_id"`
	// }
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context) (*response, error) {
		UserID := auth.GetAdminID(ctx)
		fmt.Println("Start GetAllMemberService response: ", UserID)
		resp, err := s.GetAllMember()

		if err != nil {
			fmt.Println("make GetAllMember error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}

		return &response{
			Result: "success",
			Data:   resp,
		}, nil
	}
}

/**
 * @api {post} /member/v1/get/users  Get All user
 * @apiVersion 1.0.2
 * @apiName Get All user
 * @apiGroup Member
 * @apiDescription สำรหับ Get ข้อมูล ของ user ทั้งหมด
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token 46a824c366ad40f8a1436b790e98f45e (token ทดสอบ ของ admin)
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": [
 *         {
 *             "company_id": 0,
 *             "user_id": "62100900004",
 *             "role_id": 1,
 *             "role_name": "user",
 *             "dealer": 0,
 *             "profix_name": "",
 *             "sales_status": 0,
 *             "fname": "เกศสุรีย์",
 *             "lname": "เทียนทอง",
 *             "first_buy": 0,
 *             "active_status": 1,
 *             "telephone": "+66925919266",
 *             "email": "gunnie.kes2708@gmail.com",
 *             "discount_person": "",
 *             "discount_expridate": "",
 *             "ref_code": "",
 *             "invite_code": "47a9cdccb8654ae4918886830796c12b",
 *             "invite_person": 0,
 *             "create_date": "2019-10-09 05:20:42"
 *         },
 *         {
 *             "company_id": 0,
 *             "user_id": "62100900006",
 *             "role_id": 1,
 *             "role_name": "user",
 *             "dealer": 0,
 *             "profix_name": "",
 *             "sales_status": 0,
 *             "fname": "จินตนา",
 *             "lname": "ไกรสำโรง",
 *             "first_buy": 0,
 *             "active_status": 1,
 *             "telephone": "+66903320778",
 *             "email": "Yod1599@hotmail.com",
 *             "discount_person": "",
 *             "discount_expridate": "",
 *             "ref_code": "",
 *             "invite_code": "1b4d395e95ce4b42a1564d0ec73d4e5b",
 *             "invite_person": 0,
 *             "create_date": "2019-10-09 07:18:54"
 *         },
 *         {
 *             "company_id": 0,
 *             "user_id": "62100900008",
 *             "role_id": 1,
 *             "role_name": "user",
 *             "dealer": 0,
 *             "profix_name": "",
 *             "sales_status": 0,
 *             "fname": "ลัชชา",
 *             "lname": "แก้วสมตัว",
 *             "first_buy": 0,
 *             "active_status": 1,
 *             "telephone": "+66959490825",
 *             "email": "Latcha.kaewsomtua@gmail.com",
 *             "discount_person": "",
 *             "discount_expridate": "",
 *             "ref_code": "",
 *             "invite_code": "56ec148b9de74368ad5bb9a4992f000e",
 *             "invite_person": 0,
 *             "create_date": "2019-10-09 07:23:08"
 *         }
 *  ]
 * }
 */
