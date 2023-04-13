package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/otaviotech/quode/internal/application/app_errors"
	usecase "github.com/otaviotech/quode/internal/application/usecase"
)

type CreateSubjectHandler struct {
	method string
	path   string
	uc     usecase.CreateSubjectUseCaseInterface
}

func NewCreateSubjectHandler(uc usecase.CreateSubjectUseCaseInterface) *CreateSubjectHandler {
	return &CreateSubjectHandler{
		method: "POST",
		path:   "/subjects",
		uc:     uc,
	}
}

func (h *CreateSubjectHandler) GetMethod() string {
	return h.method
}

func (h *CreateSubjectHandler) GetPath() string {
	return h.path
}

func (h *CreateSubjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.uc.Execute(r.Context(), &usecase.CreateSubjectInput{
		Name:        data.Name,
		Description: data.Description,
	})

	if err != nil {
		var edv *app_errors.ErrDomainValidation
		if errors.As(err, &edv) {
			http.Error(w, "{\"error\": \""+edv.OriginalError.Error()+"\"}", http.StatusBadRequest)
			return
		}

		http.Error(w, "{\"error\": \"internal server error\"}", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
