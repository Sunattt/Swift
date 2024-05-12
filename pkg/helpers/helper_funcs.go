package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"swift/pkg/models"
	"time"

	"go.uber.org/zap"
)

const (
	KeyUserId = "UserID"
)

func BadRequest(w http.ResponseWriter, err error, logger *zap.Logger) {
	logger.Info("the user entered incorrect data", zap.Error(err))
	http.Error(w, http.StatusText(http.StatusBadRequest), 400)
}

func InternalServerError(w http.ResponseWriter, err error, logger *zap.Logger) {
	logger.Error(http.StatusText(http.StatusInternalServerError), zap.Error(err))
	http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
}

func Unauthorized(w http.ResponseWriter, logger *zap.Logger) {
	logger.Error(http.StatusText(http.StatusUnauthorized), zap.Error(errors.New("User isn't auth")))
	http.Error(w, http.StatusText(http.StatusUnauthorized), 401)
}

func Forbidden(w http.ResponseWriter, err error, logger *zap.Logger) {
	logger.Info(http.StatusText(http.StatusForbidden), zap.Error(err))
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

func NotFoundErr(w http.ResponseWriter, err error, logger *zap.Logger) {
	logger.Info(http.StatusText(http.StatusNotFound), zap.Error(err))
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func ResponseAnswer(w http.ResponseWriter, text string) error {
	answer := models.ReplyToUser{
		Date:  time.Now(),
		Reply: text,
	}
	myAnswer, err := json.MarshalIndent(answer, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(myAnswer)
	if err != nil {
		return err
	}
	return nil
}

func GetUserIDFromContext(r http.Request) (id int, err error) {
	id, ok := r.Context().Value(KeyUserId).(int)
	if !ok {
		return id, fmt.Errorf("User isn't found!")
	}
	return id, nil
}
func SendToken(w http.ResponseWriter, sendToken *models.SendToken) error {
	b, err := json.MarshalIndent(sendToken, " ", "")
	if err != nil {
		return err
	}
	
	w.Write(b)
	return nil
}
