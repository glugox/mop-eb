package controllers

import (
	"net/http"

	"github.com/glugox/mop/api/auth"
	"github.com/glugox/mop/api/models"
	"github.com/glugox/mop/api/responses"
)


func (server *Server) GetDashboard(w http.ResponseWriter, r *http.Request) {

	d := models.Dashboard{}
	wgt := models.Widget{}
	userID, err := auth.ExtractTokenID(r)

	wgts, err := wgt.FindAllWidgets(server.DB, 0)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	d.Widgets = *wgts


	ownWgts, err := wgt.FindAllWidgets(server.DB, userID)
	if err == nil {
		for _, element := range *ownWgts {
			d.Widgets = append(d.Widgets, element)
		}

	}

	responses.JsonResponse(w, http.StatusOK, d)
}

func (server *Server) GetPublicDashboard(w http.ResponseWriter, r *http.Request) {

	d := models.Dashboard{}
	wgt := models.Widget{}

	wgts, err := wgt.FindAllWidgets(server.DB, 0)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	d.Widgets = *wgts

	responses.JsonResponse(w, http.StatusOK, d)
}


func (server *Server) GetWidgets(w http.ResponseWriter, r *http.Request) {

	wgt := models.Widget{}

	wgts, err := wgt.FindAllWidgets(server.DB, 0)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	responses.JsonResponse(w, http.StatusOK, wgts)
}

func (server *Server) GetOwnWidgets(w http.ResponseWriter, r *http.Request) {

	wgt := models.Widget{}

	// Authentication is handled by the middleware
	userID, err := auth.ExtractTokenID(r)

	wgts, err := wgt.FindAllWidgets(server.DB, userID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	responses.JsonResponse(w, http.StatusOK, wgts)
}

