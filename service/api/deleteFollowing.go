package api

import (
	"NEWASA/service/api/reqcontext"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	userToUnfollow := ps.ByName("userId")
	if userToUnfollow == "" {
		ctx.Logger.Warning("Error in deleteFollowing.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if userToUnfollow == ctx.UserId {
		ctx.Logger.Warning("error in deleteFollowing: can not unfollow yourself")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err1 := rt.db.DeleteFollowing(ctx.UserId, userToUnfollow)
	if err1 != nil {
		ctx.Logger.WithError(err1).Warning("unfollowing user in deleteFollowing.go went wrong lol")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}
