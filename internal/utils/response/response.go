package response

import (
	"encoding/json"
	"net/http"
)

type Response struct{
	Status string
	Error string
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error{
	
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

const(
	StatusOk = "OK"
	statusError = "Error"
)


func GeneralResponse(err error) Response{
	return Response{
		Status: statusError,
		Error: err.Error(),
	}
}