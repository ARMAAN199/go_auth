package controllers

// controllers are basically repositories i guess.

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/ARMAAN199/Go_EcomApi/models"
	"github.com/ARMAAN199/Go_EcomApi/stores"
	"github.com/ARMAAN199/Go_EcomApi/utils"
	"github.com/ARMAAN199/Go_EcomApi/utils/types"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	store stores.UserStore
}

func NewUserController(store stores.UserStore) *UserController {
	controller := UserController{
		store: store,
	}
	return &controller
}

func (ctrl *UserController) LoginUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var userPayload types.UserLoginPayload
		err = json.Unmarshal(body, &userPayload)
		if err != nil {
			log.Print(err)
		}

		if err := utils.Validator.Struct(userPayload); err != nil {
			error := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("bad payload %v", error))
			return
		}

		user, err := ctrl.store.GetUser(userPayload.Username)
		if err != nil {
			http.Error(w, "Can't find user with this username", http.StatusBadRequest)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userPayload.Password)); err != nil {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("incorrect password %v", err))
			return
		}

		token, err := utils.GenerateJWT(user.Username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		refreshToken, err := utils.GenerateRefreshToken(user.Username)
		if err != nil {
			http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			Expires:  time.Now().Add(50 * time.Minute),
			HttpOnly: true,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Expires:  time.Now().Add(2 * time.Minute),
			HttpOnly: true,
		})

		loginToken := models.Token{
			AccessToken:  token,
			RefreshToken: refreshToken,
		}

		utils.WriteJSON(w, 200, loginToken)

	}
}

func (ctrl *UserController) RegisterUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var userPayload types.UserRegisterPayload
		err = json.Unmarshal(body, &userPayload)
		if err != nil {
			log.Print(err)
		}

		if err := utils.Validator.Struct(userPayload); err != nil {
			error := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("bad payload %v", error))
			return
		}

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(userPayload.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		DBUser := models.User{
			Username:     userPayload.Username,
			Email:        userPayload.Email,
			PasswordHash: string(hashedPass),
		}

		id, err := ctrl.store.Register(&DBUser)
		if err != nil || id == 0 {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to add to db %v", err))
			return
		}

		utils.WriteJSON(w, 200, fmt.Sprintf("created user with id %d", id))
	}
}

func (ctrl *UserController) GetUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r
		utils.WriteJSON(w, 200, "Test logged in")
	}
}
