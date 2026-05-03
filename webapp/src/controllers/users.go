package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UnfollowUser allows a user to unfollow another user by their ID
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.APIURL, userID)
	response, err := requests.RequestWithAuth(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// FollowUser allows a user to follow another user by their ID
func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userID)
	response, err := requests.RequestWithAuth(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
