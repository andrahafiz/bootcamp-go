package employee

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services Service
}

func NewHandler(service Service) Handler {
	return Handler{
		services: service,
	}
}

func (h Handler) Create(rw http.ResponseWriter, r *http.Request) {
	var req CreateEmployeeRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Create Fail",
		}
		WriteJsonResponse(rw, resp)
		return
	}
	employee := New(req.Name, req.Nip, req.Address)
	fmt.Print("masuk")
	err = h.services.Save(employee)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}
	resp := APIResponse{
		Status:  http.StatusCreated,
		Message: "Create Success",
		Payload: req,
	}

	WriteJsonResponse(rw, resp)
}


func (h Handler) GetAll(rw http.ResponseWriter, r *http.Request) {
	employees, err := h.services.repo.GetAll()
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}

	if employees == nil {
		resp := APIResponse{
			Status:  http.StatusNotFound,
			Message: "DATA NOT FOUND",
			Error:   "data not found",
		}
		WriteJsonResponse(rw, resp)
		return
	}


	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: employees,
	}
	WriteJsonResponse(rw, resp)
}

func (h Handler) Delete(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   "id is required",
		}
		WriteJsonResponse(rw, resp)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}
	_, err = h.services.repo.GetById(idInt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			resp := APIResponse{
				Status:  http.StatusNotFound,
				Message: "DATA NOT FOUND",
				Error:   "data not found",
			}
			WriteJsonResponse(rw, resp)
			return
		}
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "GET DATA FAILED",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}

	err = h.services.repo.Delete(idInt)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Delete Fail",
		}
		WriteJsonResponse(rw, resp)
		return
	}
	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
	}
	WriteJsonResponse(rw, resp)
}
