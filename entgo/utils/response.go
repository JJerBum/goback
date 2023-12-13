package utils

import (
	"encoding/json"
	"goback/entgo/model"
	"net/http"
)

func Return(w http.ResponseWriter, status bool, code int, err error, data any) {
	response := model.Response{
		Status: status,
		Code:   code,
		Error:  "",
		Data:   data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}
