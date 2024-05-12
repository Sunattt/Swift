package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"swift/pkg/helpers"
	"swift/pkg/models"
	"time"
)

const defaultProfilePhoto =`"D:\Swift-site(ssob)\foto_profile\Screenshot 2024-05-10 223659.png"`

func (h *Handler) LoginSignIn(w http.ResponseWriter, r *http.Request) {
	var u *models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		if errors.Is(err, io.EOF) {
			helpers.BadRequest(w, err, h.Logger)
			helpers.ResponseAnswer(w, "Incorrect data")
			return
		}
		log.Println("ERROR authorizate !!", err)
		helpers.InternalServerError(w, err, h.Logger)
		helpers.ResponseAnswer(w, "OPPSSS something went wrong")
		return
	}

	token, err := h.Service.GetLoginService(u)
	if err != nil {
		errAct := errors.New("User is bloced!")
		if errors.Is(err, errAct) {
			helpers.InternalServerError(w, err, h.Logger)
			err = helpers.ResponseAnswer(w, "User is not found!")
			if err != nil {
				helpers.InternalServerError(w, err, h.Logger)
				return
			}
			return
		}
		helpers.BadRequest(w, err, h.Logger)
		return
	}

	sendToken := models.SendToken{
		Date: time.Now(),
		Answer: "Authorization was successful!",
		Token: token,
	}
	err = helpers.SendToken(w, &sendToken)
	if err != nil {
	helpers.InternalServerError(w, err, h.Logger)
	return
	}

}

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	// декодируем json данные из тела запроса
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		return
	}

	err = h.Service.CheckNikName(user.NikName, user.UpdatedAt, nicknameChangeInterval)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect Username entered registration")
		return
	}

	err = h.Service.CheckPhone(user.Phone)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect Phone entered registration")
		return
	}

	err = h.Service.CheckEmail(user.Email)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect E-mail entered registration")
		return
	}

	err = h.Service.CheckPassword(user.Password)
	if err != nil {
		helpers.BadRequest(w, err, h.Logger)
		helpers.ResponseAnswer(w, "Incorrect Password entered registration")
		return
	}

	token, err := h.Service.RegistrationUser(user)
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		helpers.ResponseAnswer(w, "OPPSSS wrong while service")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	respBytes, err := json.Marshal(token)
	if err != nil {
		log.Printf("[ERROR] Can't Marshal Error JSON! Info: %v", err)
	}

	w.Write(respBytes)
	err = helpers.ResponseAnswer(w, "Registration complated successfully.")
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		return
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		return
	}

	err = helpers.ResponseAnswer(w, "Successful registration")
	if err != nil {
		helpers.InternalServerError(w, err, h.Logger)
		return
	}
}
