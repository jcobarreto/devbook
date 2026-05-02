package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadSignUpUserPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signup.html", nil)
}

func LoadMainPage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	var posts []models.Posts
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Posts
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}

func LoadEditPostPage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: "Invalid post ID"})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	var post models.Posts
	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "update-post.html", post)
}
