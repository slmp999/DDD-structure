package upload

import (
	"context"
	"fmt"
	"net/http"

	"gitlab.com/satit13/perfect_api/auth"
)

/**
 * @api {post} /upload/v1/upload/image  upload image
 * @apiVersion 1.0.0
 * @apiName upload image
 * @apiGroup Upload
 * @apiDescription สำรหับ upload image
 * @apiHeader Content-Type application/json
 * @apiHeader Access-Token c0c05dd13ac8490e934bd361e24b724a (token ทดสอบ ของ admin หรือ token ของ user)
 * @apiParam (Parameter) {SelectFile} myFile ไฟร์รูปภาพ content type image/jpeg
 * @apiParam (Parameter) {string} type sale || product || bank || bill
 * @apiParamExample {form-data} Body request:
 * {
 * 	"myFile":"ไฟล์รูป",
 * 	"type":"sales"
 * }
 * @apiSuccess {String} response สถานะตอบกลับ
 * @apiSuccess {String} Message ข้อความตอบกลับ
 * @apiSuccessExample {json} Success:
 * {
 *     "response": "success",
 *     "message": "",
 *     "data": {
 *         "url": "https://api-dev.pct2003.com/image-demo/sales/sales-993089593.png",
 *         "type": "sales"
 *     }
 * }
 */
func makeUploadimage(s Service) interface{} {
	type response struct {
		Result  string      `json:"response"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	return func(ctx context.Context, r *http.Request) (*response, error) {
		var ID string = ""
		UserID := auth.GetUserID(ctx)
		AdminID := auth.GetAdminID(ctx)
		if UserID != "" {
			ID = UserID
		} else {
			ID = AdminID
		}
		print(r)
		r.ParseMultipartForm(10 << 30)
		file, handler, err := r.FormFile("myFile")
		Type := r.FormValue("type")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return nil, err
		}

		resp, err := s.UploadImageService(file, Type, handler, ID)
		if err != nil {
			fmt.Println("makeSigninEndpoint error ", err.Error())
			return &response{Result: "false", Message: err.Error()}, nil
		}
		defer file.Close()

		fmt.Printf("Uploaded File: %+v\n", handler.Header)

		return &response{
			Result: "success",
			Data:   resp,
		}, nil

	}
}
