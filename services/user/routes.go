package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rupeshmahanta/socialmedia-go/types"
	"github.com/rupeshmahanta/socialmedia-go/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/login", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get the json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	//check the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}
	//create new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	utils.WriteJson(w, http.StatusCreated, nil)
}
