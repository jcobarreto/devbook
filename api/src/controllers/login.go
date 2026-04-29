package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userSavedInDB, err := repository.GetByEmail(user.Email)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if erro = security.VerifyPassword(userSavedInDB.Password, user.Password); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New("invalid email or password"))
		return
	}

	token, err := authentication.CreateToken(userSavedInDB.ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(userSavedInDB.ID, 10)
	responses.JSON(w, http.StatusOK, models.AuthenticationData{
		ID:    userID,
		Token: token,
	})
}
