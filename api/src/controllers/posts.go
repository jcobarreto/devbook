package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID
	if err = post.Prepare(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	post.ID, err = repository.Create(post)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	posts, err := repository.Get(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	post, err := repository.GetByID(postID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if post.ID == 0 {
		responses.Erro(w, http.StatusNotFound, errors.New("post not found"))
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserID extraído do token:")
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	postInDatabase, err := repository.GetByID(postID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if postInDatabase.ID == 0 {
		responses.Erro(w, http.StatusNotFound, errors.New("post not found"))
		return
	}

	if postInDatabase.AuthorID != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("you don't have permission to update this post"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	post.ID = postID
	post.AuthorID = userID
	if err = post.Prepare(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = repository.Update(postID, post); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	postInDatabase, err := repository.GetByID(postID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if postInDatabase.ID == 0 {
		responses.Erro(w, http.StatusNotFound, errors.New("post not found"))
		return
	}

	if postInDatabase.AuthorID != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("you don't have permission to delete this post"))
		return
	}

	if err = repository.Delete(postID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func GetPostsByUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	posts, err := repository.GetByUser(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	if err = repository.Like(postID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	if err = repository.Unlike(postID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
