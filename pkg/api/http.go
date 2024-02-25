package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/erkylima/hexagonal/hexagonal/internal/beneficiary"
	js "github.com/erkylima/hexagonal/hexagonal/internal/beneficiary/serializer/json"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type BeneficiaryHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	beneficiaryService beneficiary.BeneficiaryService
}

func NewHandler(beneficiaryService beneficiary.BeneficiaryService) BeneficiaryHandler {
	return &handler{beneficiaryService: beneficiaryService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) beneficiary.BeneficiarySerializer {
	if contentType == "application/x-msgpack" {
	}
	return &js.Beneficiary{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "name")
	benefic, err := h.beneficiaryService.Find(code)
	if err != nil {
		if errors.Cause(err) == beneficiary.ErrBeneficiaryNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("hi %v", benefic.Name)))
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	benefic, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = h.beneficiaryService.Store(benefic)
	if err != nil {
		if errors.Cause(err) == beneficiary.ErrBeneficiaryInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).Encode(benefic)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}
