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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := requests.RequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(post))
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

func LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: "Invalid post ID"})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/like", config.APIURL, postID)
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

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: "Invalid post ID"})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/unlike", config.APIURL, postID)
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

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: "Invalid post ID"})
		return
	}

	r.ParseForm()
	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
	response, err := requests.RequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(post))
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
