package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/glugox/mop/api/auth"
	"github.com/glugox/mop/api/models"
	"github.com/glugox/mop/api/responses"
	"github.com/stripe/stripe-go/v72"
)

func (server *Server) PurchaseProduct(w http.ResponseWriter, r *http.Request) {

	// Read request data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
	}
	p := models.Purchase{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	userID, err := auth.ExtractTokenID(r)
	p.UserID = userID

	// Validate
	err = p.Validate()
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Create DB record
	createPurchase, err := p.SavePurchase(server.DB)

	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Get info about user and product
	product := models.Product{}
	productFound, err := product.FindProductByID(server.DB, createPurchase.ProductID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	user := models.User{}
	userFound, err := user.FindUserByID(server.DB, createPurchase.UserID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Stripe
	stripe.Key = os.Getenv("STRIPE_KEY")
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(productFound.Price)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		ReceiptEmail: stripe.String(userFound.Email),
	}
	pi, _ := paymentintent.New(params)



	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, createPurchase.ID))
	responses.JsonResponse(w, http.StatusCreated, pi)
}

func (server *Server) GetAllPurchases(w http.ResponseWriter, r *http.Request) {

	p := models.Purchase{}

	ps, err := p.FindAllPurchases(server.DB, 0)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	responses.JsonResponse(w, http.StatusOK, ps)
}

func (server *Server) GetOwnPurchases(w http.ResponseWriter, r *http.Request) {

	p := models.Purchase{}

	// Authentication is handled by the middleware
	userID, _ := auth.ExtractTokenID(r)

	ps, err := p.FindAllPurchases(server.DB, userID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	responses.JsonResponse(w, http.StatusOK, ps)
}

