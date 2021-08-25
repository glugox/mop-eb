package controllers

import (
	"net/http"

	"github.com/glugox/mop/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JsonResponse(w, http.StatusOK, "Home OK!")

}