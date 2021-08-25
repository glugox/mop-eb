package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/glugox/mop/api/requests"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/glugox/mop/api/models"
	"github.com/glugox/mop/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
	}
	regReq := requests.RegisterRequest{}
	err = json.Unmarshal(body, &regReq)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = regReq.Validate();
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{
	  Email: regReq.Email,
	  FirstName: regReq.FirstName,
	  LastName: regReq.LastName,
	  Password: regReq.Password,
	}
	user.Prepare()
	err = user.Validate("register")
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JsonResponse(w, http.StatusCreated, userCreated)
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	users, err := user.FindAll(server.DB)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	responses.JsonResponse(w, http.StatusOK, users)
}


func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	responses.JsonResponse(w, http.StatusOK, userGotten)
}