package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/reward-rabieth/b2b/util"
	"net/http"
)

var ErrRecordNotFound = pgx.ErrNoRows

type createUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *App) Register(w http.ResponseWriter, r *http.Request) {
	var newUser createUserRequest
	fmt.Println(newUser)
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to decode body"), http.StatusBadRequest, w)
		return
	}

	// Attempt to retrieve user by email
	_, err = app.repos.GetUserByEmail(r.Context(), newUser.Email)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			// No user found with this email, proceed with registration
			createdUser, err := app.UserSessionComponent.CreateUser(r.Context(), newUser.Email, newUser.Password, newUser.Username, app.repos, newUser.Role)
			if err != nil {
				app.HandleAPIError(fmt.Errorf("failed to sign up user: %w", err), http.StatusInternalServerError, w)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(createdUser); err != nil {
				app.HandleAPIError(fmt.Errorf("failed to encode the response: %w", err), http.StatusInternalServerError, w)
			}
			return
		} else {
			// An unexpected error occurred
			app.HandleAPIError(fmt.Errorf("unexpected error checking email: %w", err), http.StatusInternalServerError, w)
			return
		}
	}

	// If the email is found, return a conflict response
	app.HandleAPIError(fmt.Errorf("email already exists"), http.StatusConflict, w)
	return
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	var user loginUserRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to decode body: %v", err), http.StatusBadRequest, w)
		return
	}

	dbUser, err := app.repos.GetUserByEmail(r.Context(), user.Email)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to get user by email: %v", err), http.StatusInternalServerError, w)
		return
	}

	fmt.Println("DB User:", dbUser) // Log the retrieved user for debugging

	fmt.Println("DB Password:", dbUser.Password)
	fmt.Println("Request Password:", user.Password)

	if err != nil {
		return
	}
	err = util.CheckPassword(user.Password, dbUser.Password)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("passwords do not match: %v", err), http.StatusBadRequest, w)
		return
	}

	AuthUser, err := app.UserSessionComponent.LoginUser(r.Context(), user.Email, user.Password, app.repos)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("error: %v", err), http.StatusInternalServerError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(AuthUser); err != nil {
		app.HandleAPIError(fmt.Errorf("failed to encode the response: %v", err), http.StatusInternalServerError, w)
	}
}
