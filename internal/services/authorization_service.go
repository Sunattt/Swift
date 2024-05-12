package services

import (
	"errors"
	"log"
	"swift/internal/configs"
	"swift/internal/repositories"
	"swift/pkg/jwt_token"
	"swift/pkg/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository *repositories.Repository
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) CheckEmail(email string) error {
	//Если пользователь не ввёл данные
	if email == "" || !IsEmailValid(email) {
		return ErrRegistration
	}
	return nil
}

func (s *Service) CheckNikName(nikName string, lastUpdate time.Time, interval time.Duration) error {
	if nikName == "" || !IsNikNameValid(nikName) {
		return ErrRegistration
	}
	 // Проверка интервала для изменения никнейма
	 if time.Since(lastUpdate) < interval {
		return errors.New("nickname cannot be changed within the specified interval")
	   }
	  

	isFree, err := s.Repository.IsNikNameFree(nikName)
	if err != nil {
		return err
	}

	if isFree {
		return ErrNikNameRegistration
	}
	return nil
}

func(s *Service) ChackName(name string, lastUpdate time.Time, interval time.Duration) error{
	if len(name) < 4{
		return errors.New("invalid nickname length")
	}
	// Проверка интервала для изменения никнейма
	if time.Since(lastUpdate) < interval {
		return errors.New("nickname cannot be changed within the specified interval")
	}
	  
	return nil
}

func (s *Service) CheckPhone(phone string) error {
	if phone == "" || !IsPhoneValid(phone) {
		return ErrRegistration
	}
	return nil
}

func (s *Service) CheckPassword(password string) error {
	if password == "" || !IsPasswordValid(password) {
		return ErrRegistration
	}
	return nil
}

func (s *Service) RegistrationUser(u *models.User) (token string, err error) {
	//хеширование пароля (скрыть)
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	
	err = s.Repository.AddUserToDB(u, string(hash))
	if err != nil {
		log.Println("Error while add user to db")
		return "", err
	}

	token, eror := jwt_token.CreateToken(u.NikName, configs.Settings.JWTSecret)
	if eror != nil {
		return "", err
	}
	err = s.Repository.AddTokenToDb(u.Id, token)
	if err != nil {
		return "", err
	}

	return token, nil
}
//
func (s *Service) GetLoginService(u *models.User) (token string,err error) {
	active, err := s.Repository.DBCheckActiveById(u.Id)
	if err != nil {
		return "",err
	}

	if !active {
		err = errors.New("User is blocked!")
		return "", err
	}

	user, userCount, err := s.Repository.DBCheckNikName(u.NikName)
	if err != nil {
		return "", err
	}

	switch {
	case userCount <= 0:
		err = errors.New("User is not found")
		return "",err
	case userCount > 1:
		err = errors.New("Another same user found, conact tehnical support")
		return "",err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return  "",err
	}

	token, err = jwt_token.CreateToken(u.NikName, configs.Settings.JWTSecret)
	if err != nil {
		return  "" ,err
	}

	return token,	nil
}

