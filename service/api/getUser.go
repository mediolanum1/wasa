package api

import (
	"NEWASA/service/api/reqcontext"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	username := query.Get("username")
	if username == "" {
		ctx.Logger.Warning("Error in getUser.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)

	}
	// getting users from db
	user, err := rt.db.GetUser(username)
	if err != nil {
		ctx.Logger.WithError(err).Warning("user not found")
		w.WriteHeader(http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		ctx.Logger.WithError(err).Warning("listing user went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}

}
