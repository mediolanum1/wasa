package api

import (
	"NEWASA/service/api/reqcontext"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	// getting userId of banned user
	userToUnban := ps.ByName("bannedId")

	// check that bannedId is not empty
	if userToUnban == "" {
		ctx.Logger.Warning("Error in deleteBanned.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if userToUnban == ctx.UserId {
		ctx.Logger.Warning("Error in deleteBanned.go: can not unban yourself")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.db.DeleteBanned(ctx.UserId, userToUnban)
	if err != nil {
		ctx.Logger.WithError(err).Warning("error in deleteBanned.go: deleting user from banned list went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}
