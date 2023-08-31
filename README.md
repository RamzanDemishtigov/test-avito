Чтобы запустить:
docker-compose up

Руты для запросов:
s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET") - главная страница
Руты для юзеров:
s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.DeleteUser)).Methods("DELETE")
Руты для сегментов:
s.Router.HandleFunc("/segments", middlewares.SetMiddlewareJSON(s.CreateSegment)).Methods("POST")
s.Router.HandleFunc("/segments", middlewares.SetMiddlewareJSON(s.GetSegments)).Methods("GET")
s.Router.HandleFunc("/segments/{id}", middlewares.SetMiddlewareJSON(s.GetSegment)).Methods("GET")
s.Router.HandleFunc("/segments/{id}", middlewares.SetMiddlewareJSON(s.UpdateSegment)).Methods("PUT")
s.Router.HandleFunc("/segments/{id}", middlewares.SetMiddlewareJSON(s.DeleteSegment)).Methods("DELETE")
Руты для сегментов пользователей:
s.Router.HandleFunc("/usersSegments", middlewares.SetMiddlewareJSON(s.AddUserToSegments)).Methods("POST")
s.Router.HandleFunc("/usersSegments/{id}", middlewares.SetMiddlewareJSON(s.GetUsersSegments)).Methods("GET")
s.Router.HandleFunc("/usersSegments",middlewares.SetMiddlewareJSON(s.GetAllUsersSegments)).Methods("GET")