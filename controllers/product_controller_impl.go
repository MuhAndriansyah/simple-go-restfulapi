package controllers

import (
	"encoding/json"
	"go_restfulapi/data/request"
	"go_restfulapi/data/response"
	"go_restfulapi/helpers"
	"go_restfulapi/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	productRequest := request.ProductCreateRequest{}
	err := decoder.Decode(&productRequest)
	helpers.PanicIfError(err)

	productResponse := controller.ProductService.Create(r.Context(), productRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   productResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helpers.PanicIfError(err)
}

func (controller *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	productRequest := request.ProductUpdateRequest{}
	err := decoder.Decode(&productRequest)
	helpers.PanicIfError(err)

	id, err := strconv.Atoi(ps.ByName("productId"))
	helpers.PanicIfError(err)

	productRequest.Id = id

	productResponse := controller.ProductService.Update(r.Context(), productRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helpers.PanicIfError(err)
}

func (controller *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("productId"))
	helpers.PanicIfError(err)

	controller.ProductService.Delete(r.Context(), id)

	webResponse := response.WebResponse{
		Code:   http.StatusNoContent,
		Status: "OK",
	}

	w.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helpers.PanicIfError(err)
}

func (controller *ProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("productId"))
	helpers.PanicIfError(err)

	productResponse := controller.ProductService.FindById(r.Context(), id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helpers.PanicIfError(err)
}

func (controller *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	productResponses := controller.ProductService.FindAll(r.Context())

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponses,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	helpers.PanicIfError(err)
}
