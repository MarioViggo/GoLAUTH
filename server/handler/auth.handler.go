package handler

import (
	"encoding/json"
	"golauth/dados"
	"golauth/model"
	"net/http"

	"github.com/gorilla/mux"
)

type loginResponse struct {
	Token string
	User  *model.User
}

func (a *App) AssignProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assigned, err := a.d.AssignProfile(vars["user_uuid"], vars["profile_uuid"])
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}

	if !assigned {
		w.Write(dados.JSONMessage("Perfil não foi relacionado!"))
		return
	}

	w.Write(dados.JSONMessage("Perfil relacionado com sucesso!"))
	return
}

func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {
	var auth *model.Auth
	_ = json.NewDecoder(r.Body).Decode(&auth)

	created, err, id := a.d.Register(auth)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}

	user, err := a.d.GetUserByID(id)
	jsonResponse, err := json.Marshal(user.Uuid)

	if !created {
		w.Write(dados.JSONMessage("Usuário não foi cadastrado!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
	return
}

func (a *App) SignIn(w http.ResponseWriter, r *http.Request) {
	var (
		auth      *model.Auth
		logged    bool
		user      *model.User
		userSpare *model.Auth
		_         = json.NewDecoder(r.Body).Decode(&auth)
	)

	logged, user = a.d.MatchEmailAndPassword(auth.Email, auth.Password)

	_ = userSpare
	if logged {
		token := dados.GenerateToken(user.Email, user.Password)
		response := loginResponse{
			Token: token,
			User:  user,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(jsonResponse)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(dados.JSONMessage("Credênciais Incorretas!"))
	}

}
