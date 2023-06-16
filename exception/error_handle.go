package exception

import (
	"encoding/json"
	"go_restfulapi/data/response"
	"go_restfulapi/helpers"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundHandle(w, r, err) {
		return
	}
	internalServerError(w, r, err)

}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := response.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	encoder := json.NewEncoder(w)
	errResponse := encoder.Encode(webResponse)
	helpers.PanicIfError(errResponse)
}

func notFoundHandle(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := response.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		encoder := json.NewEncoder(w)
		errResponse := encoder.Encode(webResponse)
		helpers.PanicIfError(errResponse)
		return true
	} else {
		return false
	}
}
