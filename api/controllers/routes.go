package controllers

import "github.com/glugox/mop/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/register", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/my/purchases", middlewares.SetMiddlewareAuthentication(s.GetOwnPurchases)).Methods("GET")

	//Products routes
	s.Router.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.GetProducts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetProduct)).Methods("GET")

	// Purchases
	// Make a purchase
	s.Router.HandleFunc("/purchases", middlewares.SetMiddlewareAuthentication(s.PurchaseProduct)).Methods("POST")
	// Get all purchases, for user own purchases , see /my/purchases
	s.Router.HandleFunc("/purchases", middlewares.SetMiddlewareAuthentication(s.GetAllPurchases)).Methods("GET")

	// Dashboard
	// TODO Create GetDashboard
	s.Router.HandleFunc("/dashboard", middlewares.SetMiddlewareJSON(s.GetPublicDashboard)).Methods("GET")
	s.Router.HandleFunc("/dashboard", middlewares.SetMiddlewareAuthentication(s.GetDashboard)).Methods("POST")
}