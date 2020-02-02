package auth

import (
	"context"
	"fmt"
	"net/http"
)

func makeSignUpEndpoint(s Service) interface{} {
	type request struct {
		Code        string `json:"code"`
		RefCode     string `json:"ref_code"`
		OTPvalidate string `json:"pass_code"`
		Password    string `json:"password"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeSigninEndpoint")
		resp, err := s.SignUpService(req.Code, req.RefCode, req.OTPvalidate, req.Password)
		fmt.Println("makeSigninEndpoint response: ", resp)
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

func makeOtpRequestEndpoint(s Service) interface{} {
	type request struct {
		Code string `json:"code"`
	}
	type response struct {
		Result  string `json:"response"`
		Message string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeOtpRequestEndpoint")
		_, err := s.SwitchOtpRequest(req.Code)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result:  "success",
			Message: "send otp to " + req.Code,
		}, nil
	}
}

func makeGetProfileAddressEndpoint(s Service) interface{} {
	type request struct {
		UserID    int64 `json:"user_id"`
		AddressID int64 `json:"address_id"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		resp, err := s.GetProfileADdressByID(req.UserID, req.AddressID)
		fmt.Println("makeSigninEndpoint response: ", resp)
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

func makeAddProfileAddressEndpoint(s Service) interface{} {
	type request AuthProfileAddress
	type response struct {
		Result  string `json:"response"`
		Message string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeSigninEndpoint")
		profile := AuthProfileAddress{
			Name:        req.Name,
			Phone:       req.Phone,
			Province:    req.Province,
			District:    req.District,
			PostalCode:  req.PostalCode,
			MainAddress: req.MainAddress,
		}

		resp, err := s.AddProfileAddressService(&profile)
		fmt.Println("makeSigninEndpoint response: ", resp)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result: "success",
		}, nil
	}
}

func makeListProfileEndpoint(s Service) interface{} {
	type request struct {
		UserID int64 `json:"user_id"`
	}
	type response struct {
		Result  string `json:"response"`
		Message string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeSigninEndpoint")

		resp, err := s.listProfileAddressService(req.UserID)
		fmt.Println("makeSigninEndpoint response: ", resp)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result: "success",
		}, nil
	}
}

func makeDeleteProfileEndpoint(s Service) interface{} {
	type request struct {
		UserID    int64 `json:"user_id"`
		AddressID int64 `json:"address_id"`
	}
	type response struct {
		Result  string `json:"response"`
		Message string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeSigninEndpoint")

		resp, err := s.DeleteProfileAddressService(req.UserID, req.AddressID)
		fmt.Println("makeSigninEndpoint response: ", resp)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result: "success",
		}, nil
	}
}

func makeUpdateProfileEndpoint(s Service) interface{} {
	type request AuthProfileAddress
	type response struct {
		Result  string `json:"response"`
		Message string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeSigninEndpoint")
		profile := AuthProfileAddress{
			AddressID:   req.AddressID,
			Name:        req.Name,
			Phone:       req.Phone,
			Province:    req.Province,
			District:    req.District,
			PostalCode:  req.PostalCode,
			MainAddress: req.MainAddress,
		}

		resp, err := s.UpdateProfileUser(&profile)
		fmt.Println("makeSigninEndpoint response: ", resp)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result: "success",
		}, nil
	}
}

func makeForgetPasswordEndpoint(s Service) interface{} {
	type request ModelRequestForgotPassword
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeForgetPasswordEndpoint")

		model := ModelRequestForgotPassword{
			Code: req.Code,
			Host: req.Host,
			Part: req.Part,
		}
		resp, err := s.SwitchFogetEmail(&model)
		fmt.Println("makeForgetPasswordEndpoint response: ", resp)
		if err != nil {
			fmt.Println("makeForgetPasswordEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result: "success",
			Data:   resp,
		}, nil
	}
}

/**
 * @api {post} /auth/v1/forget/password forget password
 * @apiVersion 1.0.2
 * @apiName  forget password
 * @apiGroup Auth
 * @apiDescription สำรหับ ลืมรหัสผ่าน ชอเปลี่ยนรหัส
 * @apiHeader Content-Type application/json
 * @apiHeader Content-Type application/json
 * @apiParam (Parameter) {string} code จะเป็น email หรือ เบอร์โทรก็ได้ กรณีเบอร์โทรใส่ +
 * @apiParam (Parameter) {string} host เป็น host ที่ไช้เช่น pct2003.com กรณีเป็น dev dev.pct2003.com
 * @apiParam (Parameter) {string} part กระณีเร้าไปอีก หน้า เช่น pct2003.com/forget แล้ว api จะส่ง token ตามมาเป็น ?token=?
 * @apiParamExample {json} Body request:
 * {
 *     "code": "nziomini@gmail.com",
 *     "host": "https://www.pct2003.com",
 *     "path":"forget"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *    "response": "success",
 *    "message": "",
 *    "data": "https://www.pct2003.com/forget/8baf127d030a4ff0adddb2502b077855"
 *}
 */

func makeResetPasswordEndpoint(s Service) interface{} {
	type request struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeResetPasswordEndpoint")
		resp, err := s.ResetPassWordService(req.Token, req.NewPassword)
		fmt.Println("makeResetPasswordEndpoint response: ", resp)
		if err != nil {
			fmt.Println("makeResetPasswordEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result:  "success",
			Message: "เปลี่ยนรหัสผ่านเส็จสิ้น",
			Data:    resp,
		}, nil
	}
}

/**
 * @api {post} /auth/v1/reset/password reset password
 * @apiVersion 1.0.2
 * @apiName  reset password
 * @apiGroup Auth
 * @apiDescription สำรหับ รีเช็ทรหัสผ่าน
 * @apiHeader Content-Type application/json
 * @apiParam (Parameter) {string} token จะเป็น token ที่ได้จากการส่งของendpoint
 * @apiParam (Parameter) {string} new_password รหัสผ่านไหม่
 * @apiParamExample {json} Body request:
 * {
 *     "token":"8baf127d030a4ff0adddb2502b077855",
 *     "new_password": "1234567890",
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "เปลี่ยนรหัสผ่านเส็จสิ้น",
 *     "data": {
 *         "token_1": "44197ea8847a436595282307f31b30b8",
 *         "token_2": "4a7c62e6f1384b8cbef1beaf975a9f11"
 *     }
 * }
 */

func makeOtpValidateEndpoint(s Service) interface{} {
	type request struct {
		Code     string `json:"code"`
		PassCode string `json:"pass_code"`
	}
	type response struct {
		Result  string `json:"response"`
		Message string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeSigninEndpoint")
		if len(req.PassCode) != 6 {
			return &response{Result: "false", Message: "OTP length unexpected"}, nil
		}
		resp, err := s.OTPValidate(req.Code, req.PassCode)
		fmt.Println("makeSigninEndpoint response: ", resp)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		if resp != true {
			return &response{Result: "false", Message: "validate falied"}, nil
		}
		return &response{
			Result: "success",
		}, nil
	}
}

func makeChangePassword(s Service) interface{} {
	type request struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	type response struct {
		Result  string `json:"response"`
		Message string `json:"message"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeSigninEndpoint")
		if req.NewPassword == req.OldPassword {
			return &response{Result: "false", Message: "รหัสผ่านใหม่ กับ ช้ำกับรหัสผ่านเก่า"}, nil
		}
		UserID := GetUserID(ctx)
		_, err := s.ChangePasswordService(UserID, req.OldPassword, req.NewPassword)
		if err != nil {
			return &response{Result: "false", Message: err.Error()}, nil
		}
		return &response{
			Result:  "success",
			Message: "เปลี่ยนรหัสผ่านสำเร็จ",
		}, nil

	}
}

/**
 * @api {post} /auth/v1/change/password change password
 * @apiVersion 1.0.2
 * @apiName  change password
 * @apiGroup Auth
 * @apiDescription สำรหับ เปลี่ยน รหัสผ่าน
 * @apiHeader Content-Type application/json
 * @apiHeader Access-token 61ebaf100b4e41109516b1f45f5586c4 token ของuser
 * @apiParam (Parameter) {string} old_password รหัสผ่านเก่า กรณีรหัสช้ำกันจะไม่สามารถเปลี่ยนได้
 * @apiParam (Parameter) {string} new_password รหัสผ่านใหม่ กรณีรหัสช้ำกันจะไม่สามารถเปลี่ยนได้
 * @apiParamExample {json} Body request:
{
    "old_password": "123456",
    "new_password": "1234567"
}
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "เปลี่ยนรหัสผ่านเส็จสิ้น"
 * }
*/

func makeGetProfile(s Service) interface{} {
	type request struct {
		Code string `json:"code"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeSigninEndpoint")
		resp, err := s.GetUserService(req.Code)
		fmt.Println("makeSigninEndpoint response: ", resp)
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

func makeSigninEndpoint(s Service) interface{} {
	type request struct {
		Code     string `json:"code"`
		Password string `json:"password"`
	}
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, req *request) (*response, error) {
		fmt.Println("begin endpoint.makeSigninEndpoint")
		resp, err := s.SigninService(req.Code, req.Password)
		fmt.Println("makeSigninEndpoint response: ", resp)
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

// func makeSigninV2Endpoint(s Service) interface{} {
// 	type request struct {
// 		Email    string  `json:"email"`
// 		Password string  `json:"password"`
// 		FCMToken *string `json:"fcm_token,omitempty"`
// 	}
// 	type response struct {
// 		Result  string      `json:"result"`
// 		Message string      `json:"message"`
// 		Data    interface{} `json:"data"`
// 	}
// 	return func(ctx context.Context, req *request) (interface{}, error) {
// 		fmt.Println(1)
// 		resp, err := s.SigninV2(req.Email, req.Password, req.FCMToken)
// 		fmt.Println(resp)
// 		if err != nil {
// 			return &response{Result: "failed", Message: ""}, err
// 		}
// 		return &response{Result: "success", Data: resp}, nil
// 	}
// }

func makeSignoutEndpoint(s Service) interface{} {
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
 * @api {post} /auth/v1/signout SignOut
 * @apiVersion 1.0.2
 * @apiName SignOut
 * @apiGroup Auth
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
// func makeGetProfileEndpoint(s Service) interface{} {
// 	type response struct {
// 		Result     string  `json:"result"`
// 		AccountID  int64   `json:"account_id"`
// 		FullName   string  `json:"full_name"`
// 		Email      string  `json:"email"`
// 		PictureURL *string `json:"picture_url"`
// 		Status     string  `json:"status"`
// 	}
// 	return func(ctx context.Context, req *http.Request) (interface{}, error) {
// 		fmt.Println("begin endpoint.makeGetProfileEndpoint")
// 		xat := req.Header.Get("X-Access-Token")
// 		p, err := s.Profile(xat)
// 		if err != nil {
// 			fmt.Println("makeGetProfileEndpoint error ", err.Error())
// 			return &response{Result: "failed"}, err
// 		}
// 		res := response{
// 			Result:     "success",
// 			AccountID:  p.AccountID,
// 			FullName:   p.FullName,
// 			Email:      p.Email,
// 			PictureURL: p.PictureURL,
// 			Status:     "Unverified",
// 		}
// 		if p.Status == 1 {
// 			res.Status = "Verified"
// 		}
// 		return &res, nil
// 	}
// }
