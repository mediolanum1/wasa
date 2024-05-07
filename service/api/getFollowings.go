package api

import (
	"NEWASA/service/api/reqcontext"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	userId := ps.ByName("userId")
	if userId == "" {
		ctx.Logger.Warning("Error in getFollowings.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users, err := rt.db.GetFollowings(userId)
	if err != nil {
		ctx.Logger.WithError(err).Warning("error in getFollowings: getting users list went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		ctx.Logger.WithError(err).Warning("listing users went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
