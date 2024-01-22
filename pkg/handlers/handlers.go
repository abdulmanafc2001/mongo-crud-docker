package handlers

import (
	"crud/pkg/models"
	"crud/pkg/usecase"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Routes  *mux.Router
	UseCase usecase.UserUsecase
}

func NewRouter(usecase usecase.UserUsecase) *Router {
	mux := mux.NewRouter()

	return &Router{
		Routes:  mux,
		UseCase: usecase,
	}
}

func (r *Router) Run() error {
	r.Routes.HandleFunc("/create-user", r.CreateUser).Methods("POST")
	r.Routes.HandleFunc("/read-users", r.ReadUser).Methods("GET")
	r.Routes.HandleFunc("/delete-users", r.DeleteUser).Methods("DELETE")
	r.Routes.HandleFunc("/get-one-user/{user_id}", r.GetOneUser).Methods("GET")
	r.Routes.HandleFunc("/update-user/{user_id}", r.UpdateUser).Methods("PUT")
	r.Routes.HandleFunc("/delete-user/{user_id}", r.DeleteUserWithId).Methods("DELETE")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r.Routes,
	}
	return srv.ListenAndServe()
}

func (router *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload models.User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	id, err := router.UseCase.CreateUser(payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data := map[string]any{
		"id": id,
	}
	json.NewEncoder(w).Encode(data)
}

func (router *Router) ReadUser(w http.ResponseWriter, r *http.Request) {
	users, err := router.UseCase.ReadUser()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	data := map[string]any{
		"data": users,
	}
	json.NewEncoder(w).Encode(data)
}

func (router *Router) DeleteUser(w http.ResponseWriter, r *http.Request) {
	err := router.UseCase.DeleteUsers()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	data := map[string]any{
		"data": "Successfully deleted users",
	}
	json.NewEncoder(w).Encode(data)

}

func (router *Router) GetOneUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]

	usr, err := router.UseCase.GetOneUser(userId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := map[string]any{
		"data": usr,
	}
	json.NewEncoder(w).Encode(data)
}

func (router *Router) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["user_id"]

	var usr models.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = router.UseCase.UpdateUser(id, usr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := map[string]any{
		"data": "Successfully updated user data",
	}
	json.NewEncoder(w).Encode(data)

}

func (router *Router) DeleteUserWithId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["user_id"]

	err := router.UseCase.DeleteUser(id)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	data := map[string]any{
		"data": "Successfully deleted users",
	}
	json.NewEncoder(w).Encode(data)
}
