package upload

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"

	log "gitlab.com/satit13/perfect_api/logger"
	"gopkg.in/h2non/filetype.v1/matchers"
)

type service struct {
	repo Repository
	Mode string
}

// NewService creates new auth service
func NewService(member Repository, Mode string) (Service, error) {
	s := service{member, Mode}
	return &s, nil
}

type Service interface {
	UploadImageService(image multipart.File, Type string, handle *multipart.FileHeader, userID string) (interface{}, error)
}

func (s *service) UploadImageService(image multipart.File, Type string, handle *multipart.FileHeader, userID string) (interface{}, error) {
	_, err := createDirectory(Type, s.Mode, userID)
	if err != nil {
		return nil, err
	}

	resp, err := s.UploadProduct(image, Type, handle, userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
	// switch Type {
	// case "1":
	// 	fileHeader := make([]byte, 512)
	// 	if _, err := image.Read(fileHeader); err != nil {
	// 		return nil, err
	// 	}
	// 	resp, err := s.UploadProduct(image, Type)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return resp, nil
	// case "2":

	// }
	// return nil, nil
}

func CheckTypeImage(Type string) (string, error) {
	switch Type {
	case "image/jpeg", "image/jpg":
		return "jpg", nil
	case "image/gif":
		return "gif", nil
	case "image/png":
		return "png", nil
	case "application/pdf": // not image, but application !
		return "pdf", nil
	default:
		return "png", errors.New("กรุณา อัพโหลดไฟล์ jpg pdf png gif")
	}
}
func createDirectory(dirName string, mode string, userID string) (bool, error) {
	if mode == "Production" {
		src, err := os.Stat("/app/image/" + userID + "/" + dirName)

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("/app/image/"+userID+"/"+dirName, 0755)
			if errDir != nil {
				panic(err)
			}
			return true, nil
		}

		if src.Mode().IsRegular() {
			fmt.Println(dirName, "already exist as a file!")
			return false, nil
		}

		return false, nil
	} else {
		src, err := os.Stat("/app/image-demo/" + userID + "/" + dirName)

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("/app/image-demo/"+userID+"/"+dirName, 0755)
			if errDir != nil {
				panic(err)
			}
			return true, nil
		}

		if src.Mode().IsRegular() {
			fmt.Println(dirName, "already exist as a file!")
			return false, nil
		}

		return false, nil
	}
}

func checkimage(image []byte) (string, error) {

	if matchers.Jpeg(image) {
		return "jpg", nil
	} else if matchers.Png(image) {
		return "png", nil
	} else if matchers.Gif(image) {
		return "gif", nil
	} else if matchers.Pdf(image) {
		return "pdf", nil
	} else {
		return "", errors.New("กรุณา อัพโหลดไฟล์ jpg pdf png gif")
	}
	return "", nil
}

func (s *service) UploadProduct(image multipart.File, Type string, handle *multipart.FileHeader, userID string) (interface{}, error) {
	if s.Mode == "Production" {
		file, err := ioutil.ReadAll(image)
		if err != nil {
			return nil, err
		}
		types, err := checkimage(file)
		if err != nil {
			return nil, err
		}
		tempFile, err := ioutil.TempFile(`/app/image/`+userID+`/`+Type+`/`, Type+"-*."+types)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer tempFile.Close()
		tempFile.Write(file)

		var imagename string = ""

		filename := strings.Replace(tempFile.Name(), `\`, `/`, -1)
		files := strings.Split(filename, "/")
		for i, in := range files {
			if i > 1 {
				imagename += "/" + in
			}
		}
		URL := "https://api.pct2003.com"
		fmt.Println(filename, files)
		prodcut := UploadModel{
			Url:  URL + imagename,
			Type: Type,
		}
		return prodcut, nil
	} else {
		file, err := ioutil.ReadAll(image)
		if err != nil {
			return nil, err
		}
		types, err := checkimage(file)
		if err != nil {
			return nil, err
		}
		tempFile, err := ioutil.TempFile(`/app/image-demo/`+userID+`/`+Type+`/`, Type+"-*."+types)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer tempFile.Close()

		tempFile.Write(file)

		var imagename string = ""

		filename := strings.Replace(tempFile.Name(), `\`, `/`, -1)
		log.Println(filename)
		files := strings.Split(filename, "/")
		for i, in := range files {
			if i > 1 {
				imagename += "/" + in
			}
		}
		URL := "https://api-dev.pct2003.com"
		fmt.Println(filename, files)
		prodcut := UploadModel{
			Url:  URL + imagename,
			Type: Type,
		}
		return prodcut, nil
	}
}
