package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
