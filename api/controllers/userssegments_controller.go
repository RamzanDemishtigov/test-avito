package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main/api/models"
	"main/api/responses"
	"main/api/utils/formaterror"
)

func (server *Server) AddUserToSegments(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userssegment := models.UsersSegmentsDTO{}
	err = json.Unmarshal(body, &userssegment)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userssegmentCreated, err := userssegment.AddUserToSegments(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.URL.Path))
	responses.JSON(w, http.StatusCreated, userssegmentCreated)
}

func (server *Server) GetAllUsersSegments(w http.ResponseWriter, r *http.Request) {

	userssegment := models.UsersSegments{}

	segments, err := userssegment.FindAllUsersSegments(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, segments)
}

func (server *Server) GetUsersSegments(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userssegments := models.UsersSegments{}

	userssegmentsReceived, err := userssegments.FindUsersSegmentsByID(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userssegmentsReceived)
}