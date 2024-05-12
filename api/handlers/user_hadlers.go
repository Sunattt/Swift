package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"swift/internal/services"
	"swift/pkg/helpers"
	"swift/pkg/models"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)
const (
	nicknameChangeInterval = 14 * 24 * time.Hour // 14 дней
	nameChangeInterval     = 14 * 24 * time.Hour // 14 дней
   )

type Handler struct {
	Service *services.Service
	Logger  *zap.Logger
	Router  *mux.Router
}

func NewHandler(service *services.Service, logger *zap.Logger) *Handler {
	return &Handler{
		Service: service,
		Logger:  logger,
		Router:  mux.NewRouter(),
	}
}

func (h *Handler) GetProfileUser(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetUserIDFromContext(*r)

	// val := r.URL.Query()
	// val.Get("username")
	if err != nil {
		helpers.NotFoundErr(w, err, h.Logger)
		err = helpers.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helpers.InternalServerError(w, err, h.Logger)
			return
		}
		return
	}

	user, err := h.Service.GetUserFromService(id)
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Opps something went wrong!!")
		return
	}

	u, err := json.Marshal(user)
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		return
	}

	w.Write(u)
}

func(h *Handler) Hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello"))
}
// // Рендеринг HTML-шаблона с данными пользователя
// tmpl, err := template.ParseFiles("profile.html")
// if err != nil {
// 	helpers.InternalServerError(w, err, h.Logger)
// 	return
// }
// err = tmpl.Execute(w, user)
// if err != nil {
// 	helpers.InternalServerError(w, err, h.Logger)
// 	return
// }

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// Обработка формы для обновления данных профиля
	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			helpers.NotFoundErr(w, err, h.Logger)
			err = helpers.ResponseAnswer(w, "User is not found!")
			if err != nil {
				helpers.InternalServerError(w, err, h.Logger)
				return
			}
			return
		}
		var user *models.User
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			helpers.BadRequest(w, err, h.Logger)
			return 
		}
		
		
	err = h.Service.CheckNikName(user.NikName, user.UpdatedAt, nicknameChangeInterval)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect Username entered ")
		return
	}

	err = h.Service.ChackName(user.Name, user.UpdatedAt, nameChangeInterval)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "")
	}
	err = h.Service.CheckPhone(user.Phone)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect Phone entered")
		return
	}

	err = h.Service.CheckEmail(user.Email)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect E-mail entered ")
		return
	}

	err = h.Service.CheckPassword(user.Password)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect Password entered ")
		return
	}

	updateUser, err := h.Service.UpdateProfileUser(id, user)
		if err != nil {
			helpers.InternalServerError(w, err, h.Logger)
			helpers.ResponseAnswer(w, "OPPSSS wrong while service")
			return 
		}
		
		userBytes, err := json.Marshal(updateUser)
		if err != nil {
			helpers.InternalServerError(w, err, h.Logger)
			return
		}
	
		w.Write(userBytes)
		err = helpers.ResponseAnswer(w, "Registration complated successfully.")
		if err != nil {
			helpers.InternalServerError(w, err, h.Logger)
			return
		}
		
		return 
	}
	
}

// func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	
// 	f, fh, err := r.FormFile("file")
// 	if err != nil {
// 		helpers.InternalServerError(w, err, h.Logger)
// 		helpers.ResponseAnswer(w, "Error went getting file")
// 		return
// 	}
// 	defer f.Close()

// 	err = h.Service.UploadFileService(f, fh)
// 	if err != nil {

// 	}

// }

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	filename := r.Header.Get("filename")

	filepath := "files/" + filename

	f, err := os.ReadFile(filepath)
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filepath)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(f)

}

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.GetUserIDFromContext(*r)
	if err != nil {
		helpers.NotFoundErr(w, err, h.Logger)
		err = helpers.ResponseAnswer(w, "User is not found!")
		if err != nil {
			helpers.InternalServerError(w, err, h.Logger)
			return
		}
		return
	}

	err = h.Service.DeleteAccount(userID)
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		return
	}

	err = helpers.ResponseAnswer(w, "Account deleted successfully!")
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		return
	}

	return
}
