package api

import (
	"NEWASA/service/api/reqcontext"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	userId := ps.ByName("userId")
	if userId == "" {
		ctx.Logger.Warning("Error in getBanned.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// getting list of banned users
	bannedList, err := rt.db.GetBanned(userId)
	if err != nil {
		ctx.Logger.WithError(err).Warning("getting banned users in getBanned.go went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(bannedList)
	if err != nil {
		ctx.Logger.WithError(err).Warning("listing banned users in getBanned.go went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
