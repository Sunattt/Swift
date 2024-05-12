package models

import (
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	User struct {
		Id            int    `json:"-"`
		NikName       string `json:"nik_name"`
		Name          string `json:"name"`
		Birthday      string `json:"age"`
		Phone         string `json:"phone"`
		Country       string `json:"country"`
		Email         string `json:"email"`
		Password      string `json:"password"`
		FotoProfile   *MultipartFile
		Bio           string    `json:"bio"`
		RoleId        string    `json:"role_id"`
		Active        bool      `json:"active"`
		Followers     int       `json:"followers"`
		Subscriptions int       `json:"subscriptions"`
		Publications  int       `json:"publications"`
		CreatedAt     time.Time `json:"creared_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		DelectedAt    time.Time `json:"deleted_at"`
	}

	MultipartFile struct {
		File     io.Reader
		FileName string
	}

	SavedVideo struct {
		Id      int
		SaverId int
		SavedId int
	}

	Follows struct {
		FollowerId int
		FollowedId int
	}

	Notification struct {
		Id         int       `json:"id"`
		Message    string    `json:"notification"`
		AdresseeId int       `json:"adressee_id"`
		Status     bool      `json:"active"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		DeletedAt  time.Time `json:"deleted_at"`
	}

	Message struct {
		Id          int
		SenderId    int
		RecipientId int
		MessageText string
		CreatedAt   time.Time
		Date        time.Time
	}

	MessageReels struct {
		Id         int
		SenderId   int
		ReelsId    int
		CreatedAt  time.Time
		UpdatedAt  time.Time
		DelectedAt time.Time
	}

	Reels struct {
		Id       int
		NiName   string
		Duration string `json:"duration"`
	}

	SavedPublications struct {
		Id      int
		User_id int
		Name    string
	}

	BirthdayDayFriends struct {
		Id          int
		User_id     int
		NameFreinds string
		Birthday    time.Time
	}
)

type (
	TokenClaims struct {
		NikName string    `json:"nik_name"`
		Expire  time.Time `json:"expire"`
		jwt.StandardClaims
	}

	ConfigModel struct {
		Server    Server `json:"server"`
		Db        DB     `json:"db"`
		JWTSecret string `json:"secret_key"`
	}

	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	DB struct {
		User     string `json:"user"`
		DBname   string `json:"dbname"`
		Password string `json:"password"`
	}
	Token struct {
		Id     int       `json:"id"`
		Token  string    `json:"token"`
		UserId int       `json:"user_id"`
		Expire time.Time `json:"expire"`
	}

	SendToken struct{
		Date time.Time
		Answer string
		Token string
	}

	ReplyToUser struct {
		Date  time.Time
		Reply string
	}
)
