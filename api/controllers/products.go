package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/glugox/mop/api/models"
	"github.com/glugox/mop/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
	}
	p := models.Product{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = p.Validate()
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	createProduct, err := p.SaveProduct(server.DB)

	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, createProduct.ID))
	responses.JsonResponse(w, http.StatusCreated, createProduct)
}

func (server *Server) GetProducts(w http.ResponseWriter, r *http.Request) {

	p := models.Product{}

	ps, err := p.FindAllProducts(server.DB)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	responses.JsonResponse(w, http.StatusOK, ps)
}

func (server *Server) GetProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	p := models.Product{}
	product, err := p.FindProductByID(server.DB, uint32(uid))
	if err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	responses.JsonResponse(w, http.StatusOK, product)
}