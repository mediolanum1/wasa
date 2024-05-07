package api

import (
	"NEWASA/service/api/reqcontext"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	photoId := ps.ByName("photoId")

	if photoId == "" {
		ctx.Logger.Warning("Error in putlike.go: no photoId ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err1 := rt.db.PutLike(ctx.UserId, photoId)
	if err1 != nil {
		ctx.Logger.WithError(err1).Warning("error in putLike.go: putting like went wrong")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
