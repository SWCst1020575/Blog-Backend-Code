package api

import "net/http"

func checkRequestNotfound(w http.ResponseWriter, r *http.Request) {
	response := ApiResponse{"404", "Not Found"}
	ResponseWithJson(w, http.StatusNotFound, response)
}
func checkRequestError(err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		response := ApiResponse{"400", "Bad Request"}
		ResponseWithJson(w, http.StatusBadRequest, response)
		panic(err)
	}
}
func checkServerError(err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		response := ApiResponse{"500", "Server Error"}
		ResponseWithJson(w, http.StatusInternalServerError, response)
		panic(err)
	}
}
