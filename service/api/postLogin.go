package api

import (
	"NEWASA/service/api/reqcontext"
	"encoding/json"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var a struct {
		LoginUsername string `json:"username"`
	}
	err := json.NewDecoder(r.Body).Decode(&a)

	if a.LoginUsername == "" {
		ctx.Logger.WithError(err).Warning("Error in postLogin.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		ctx.Logger.WithError(err).Warning("Error in postLogin.go: wrong json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err = rt.db.GetUser(a.LoginUsername); err != nil {
		ctx.Logger.Println("Error in postLogin.go: username not found, creating a new user")
		err = rt.db.InsertUser(a.LoginUsername)
		if err != nil {
			ctx.Logger.WithError(err).Warning("Error in postLogin.go: error inserting user in db")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx.Logger.Printf("User  %s not found, creating a new user", a.LoginUsername)
	}
	userId, err1 := rt.db.PostLogin(a.LoginUsername)

	if err1 != nil {
		ctx.Logger.WithError(err).Warning("Error in postLogin.go: username not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	http.ErrBodyReadAfterClose = json.NewEncoder(w).Encode(userId)
	if err != nil {
		ctx.Logger.WithError(err).Warning("error returning userId")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
