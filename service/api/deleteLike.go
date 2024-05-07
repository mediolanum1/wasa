package api

import (
	"NEWASA/service/api/reqcontext"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	photoId := ps.ByName("photoId")

	if photoId == "" {
		ctx.Logger.Warning("Error in deleteLike.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.db.DeleteLike(ctx.UserId, photoId)
	if err != nil {
		ctx.Logger.WithError(err).Warning("deleting like in deleteLike.go went wrong lol")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}
