package api

import (
	"NEWASA/service/api/reqcontext"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	// get the userId of user we want to follow and check that its not empty
	followToUser := ps.ByName("userId")
	if followToUser == "" {
		ctx.Logger.Error("Error in putFollowing.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if user is trying to follow himself
	if followToUser == ctx.UserId {
		ctx.Logger.Error("Error in putFollowing.go: can not follow yourself")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.db.PutFollowing(ctx.UserId, followToUser)
	if err != nil {
		ctx.Logger.WithError(err).Warning("following user in putFollowing.go went wrong")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
