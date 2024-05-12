package handlers

import (
	"log"
	"net/http"
	"swift/internal/configs"
	"swift/pkg/db"
	"swift/pkg/helpers"
	"swift/pkg/jwt_token"
	"swift/pkg/models"
)
const (
	KeyUserId = "userID"
)

// и возвращает обработчик, который будет выполнять проверку аутентификации клиента перед вызовом следующего обработчика.
func (h *Handler) Authentication(hand http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//заголовка запросов 
		token := r.Header.Get("Authorization")
		
		nikName,valid , err := jwt_token.ValidToken(token, configs.Settings.JWTSecret)
		if err != nil {
			helpers.InternalServerError(w, err, h.Logger)
			helpers.ResponseAnswer(w, "OOPPSS somethong went wrong!!!")
			return
		}

		if !valid{
			log.Println(valid)
			helpers.Unauthorized(w, h.Logger)
			helpers.ResponseAnswer(w, "Token has expired!!!")
			return 
		}

		// ctx := context.WithValue(r.Context(),KeyUserId,nikName)
		// r = r.WithContext(ctx)

		r.Header.Set("nik_name", nikName)
		hand.ServeHTTP(w, r)	
		log.Println("User Authentication was successfully!")	
	})
}

func(h *Handler) UserRole(hand http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ur struct {
			Get, Post, Put, Del bool
		}
		nikName := r.Header.Get("nik_name")

		var user *models.User
		err := h.Service.CheckNikName(nikName, user.UpdatedAt, nicknameChangeInterval)
		if err != nil {
			helpers.BadRequest(w, err, h.Logger)
			return 
		}

		row := db.DB.QueryRow(`SELECT r.get, r.post, r.put, r.del FROM users u JOIN user_roles r ON u.role_id = r.id WHERE u.nik_name = $1`, nikName)
		err = row.Scan(&ur.Get, &ur.Post, &ur.Put, &ur.Del)
		if err != nil {
			log.Println("[ERROR] Can't scan user roles!", err)
			helpers.InternalServerError(w, err, h.Logger)
		}

		switch r.Method {
		case "GET":
			if !ur.Get {
				helpers.Forbidden(w, err, h.Logger)
				return
			}
		case "POST":
			if !ur.Post {
				helpers.Forbidden(w,err, h.Logger)
				return
			}
		case "PUT":
			if !ur.Put {
				helpers.Forbidden(w, err, h.Logger)
				return
			}
		case "DELETE":
			if !ur.Del {
				helpers.Forbidden(w, err, h.Logger)
				return
			}
		default:
			helpers.NotFoundErr(w, err, h.Logger)
			return

		}

		//Если ошибок не обнаружено, идём к необходимому роуту для продолжения обработки запроса 
		hand.ServeHTTP(w, r)
	})
}
