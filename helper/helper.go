package helper

import (
	"encoding/json"
	httpresponse "medsearch/helper/httpResponse"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, res interface{}, statusCode int) {
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, statusCode int, res httpresponse.Response) {
	RespondWithJSON(w, res, statusCode)
}
