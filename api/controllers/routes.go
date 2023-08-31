package controllers

import "main/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.DeleteUser)).Methods("DELETE")

	//Segments routes
	s.Router.HandleFunc("/segments", middlewares.SetMiddlewareJSON(s.CreateSegment)).Methods("POST")
	s.Router.HandleFunc("/segments", middlewares.SetMiddlewareJSON(s.GetSegments)).Methods("GET")
	s.Router.HandleFunc("/segments/{id}", middlewares.SetMiddlewareJSON(s.GetSegment)).Methods("GET")
	s.Router.HandleFunc("/segments/{id}", middlewares.SetMiddlewareJSON(s.UpdateSegment)).Methods("PUT")
	s.Router.HandleFunc("/segments/{id}", middlewares.SetMiddlewareJSON(s.DeleteSegment)).Methods("DELETE")

	//UsersSegments routes
	s.Router.HandleFunc("/usersSegments", middlewares.SetMiddlewareJSON(s.AddUserToSegments)).Methods("POST")
	s.Router.HandleFunc("/usersSegments/{id}", middlewares.SetMiddlewareJSON(s.GetUsersSegments)).Methods("GET")
	s.Router.HandleFunc("/usersSegments",middlewares.SetMiddlewareJSON(s.GetAllUsersSegments)).Methods("GET")
}
