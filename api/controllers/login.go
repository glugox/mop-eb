package controllers

import (
	"encoding/json"
	"errors"
	"github.com/glugox/mop/api/requests"
	"io/ioutil"
	"net/http"

	"github.com/glugox/mop/api/auth"
	"github.com/glugox/mop/api/models"
	"github.com/glugox/mop/api/responses"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	loginReq := requests.LoginRequest{}
	err = json.Unmarshal(body, &loginReq)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = loginReq.Validate()
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(loginReq.Email, loginReq.Password)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JsonResponse(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error
	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", errors.New("invalid credentials")
	}
	return auth.CreateToken(user.ID)
}