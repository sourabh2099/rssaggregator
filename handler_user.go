package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sourabh2099/rssaggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		Name string `json:name`
	}
	params := paramters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON", err))
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing Json", err))
		return
	}

	responseWithJSON(w, 200, databaseUserToUser(user))
}
func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	responseWithJSON(w, 200, databaseUserToUser(user))
}

func (apiConfig *apiConfig) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiConfig.DB.GetAllUsers(r.Context())
	if err != nil{
		responseWithError(w,400,fmt.Sprintf("Unable to fetch user data: %v",err))
	}
	responseWithJSON(w,200,users);
}
