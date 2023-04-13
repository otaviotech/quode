package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/otaviotech/quode/internal/application/app_errors"
	usecase "github.com/otaviotech/quode/internal/application/usecase"
)

type CreateAuthorHandler struct {
	method string
	path   string
	uc     usecase.CreateAuthorUseCaseInterface
}

func NewCreateAuthorHandler(uc usecase.CreateAuthorUseCaseInterface) *CreateAuthorHandler {
	return &CreateAuthorHandler{
		method: http.MethodPost,
		path:   "/authors",
		uc:     uc,
	}
}

func (h *CreateAuthorHandler) GetMethod() string {
	return h.method
}

func (h *CreateAuthorHandler) GetPath() string {
	return h.path
}

func (h *CreateAuthorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Name string `json:"name"`
		Bio  string `json:"bio"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.uc.Execute(r.Context(), &usecase.CreateAuthorInput{
		Name: data.Name,
		Bio:  data.Bio,
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
