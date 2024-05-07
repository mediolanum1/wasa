package api

import (
	"NEWASA/service/api/reqcontext"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	userToBan := ps.ByName("bannedId")
	if userToBan == "" {
		ctx.Logger.Error("Error in putBanned.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if userToBan == ctx.UserId {
		ctx.Logger.Error("Error in putBanned.go: can not ban yourself")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err1 := rt.db.PutBanned(ctx.UserId, userToBan)
	if err1 != nil {
		ctx.Logger.WithError(err1).Warning("error in putBanned.go: user already banned went wrong")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
