package exception

import (
	"fmt"
	"net/http"

	"github.com/MrBista/crud-go/helper"
	"github.com/MrBista/crud-go/model/web"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

type ValidateErrRes struct {
	Type    string      `json:"type"`
	Message interface{} `json:"message"`
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, errorOccurs := err.(validator.ValidationErrors)
	if errorOccurs {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		var messageErrors []*ValidateErrRes

		for _, fieldErr := range exception {
			var singelErr ValidateErrRes
			singelErr.Message = fieldErr.Value()
			singelErr.Type = fieldErr.Tag()

			messageErrors = append(messageErrors, &singelErr)
			fmt.Println(fieldErr)
		}

		webResponse := web.WebReponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   messageErrors,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebReponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
