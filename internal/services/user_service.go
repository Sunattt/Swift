package services

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"swift/pkg/models"
)

var (
	ErrRegistration        = errors.New("incorrect value entered for registration")
	ErrNikNameRegistration = errors.New("username is taken by another user")
)

func(s *Service) GetUserFromService(id int)(*models.User, error){
	u, err := s.Repository.GetUserFromDB(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UploadFileService(f multipart.File, fh *multipart.FileHeader) (error) {
	log.Println("file size", fh.Size)

	if fh.Size > 500_000 {
		return fmt.Errorf("file size exceeds 500_000") 
	}

	if !strings.Contains(fh.Filename, ".png") {
		return fmt.Errorf("Image type is not")
	}

	f2, err := os.Create("image/" + fh.Filename)
	if err != nil {
		return fmt.Errorf("error creating file2 %", err)
	}
	defer f2.Close()

	wr, err := io.Copy(f2, f)
	log.Println("Bytes copied", wr)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteAccount(id int) error {
	err := s.Repository.DeleteMyAccount(id)
	if err != nil {
		return err
	}

	return nil
}

func(s *Service) UpdateProfileUser(id int, user *models.User)( *models.User, error){
	user, err := s.Repository.UpdateProfileFromDB(id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func(s *Service) CheckProfilePhoto(photoURL *string) error { 
	// Проверка типа файла (допустимые типы: .jpg, .jpeg, .png)
	allowedTypes := []string{".jpg", ".jpeg", ".png"}
	//проверка типа файла url 
fileType := http.DetectContentType([]byte(*photoURL))
allowed := false
for _, t := range allowedTypes {
 if fileType == "image/"+t[1:] {
  allowed = true
  break
 }
}
if !allowed {
 return errors.New("invalid file type")
}

// Здесь можно добавить дополнительные проверки, например, максимального размера файла

return nil
}

func(s *Service) UploadFileService(f multipart.File , fh *multipart.FileHeader){
	
}
