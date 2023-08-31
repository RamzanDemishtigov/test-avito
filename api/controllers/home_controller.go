package controllers

import (
	"net/http"

	"main/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "API для сегментирования пользователей")

}
