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

func (server *Server) CreateSegment(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	segment := models.Segment{}
	err = json.Unmarshal(body, &segment)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	segment.Prepare()
	err = segment.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	segmentCreated, err := segment.SaveSegment(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, segmentCreated.ID))
	responses.JSON(w, http.StatusCreated, segmentCreated)
}

func (server *Server) GetSegments(w http.ResponseWriter, r *http.Request) {

	segment := models.Segment{}

	segments, err := segment.FindAllSegments(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, segments)
}

func (server *Server) GetSegment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	segment := models.Segment{}

	segmentReceived, err := segment.FindSegmentByID(server.DB, sid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, segmentReceived)
}

func (server *Server) UpdateSegment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	segment := models.Segment{}
	err = json.Unmarshal(body, &segment)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	segment.Prepare()
	err = segment.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	segmentUpdated, err := segment.UpdateSegment(server.DB, sid)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, segmentUpdated)
}

func (server *Server) DeleteSegment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	segment := models.Segment{}

	sid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = segment.DeleteSegment(server.DB, sid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", sid))
	responses.JSON(w, http.StatusNoContent, "")
}
