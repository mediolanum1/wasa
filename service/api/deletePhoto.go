package api

import (
	"NEWASA/service/api/reqcontext"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	photoId := ps.ByName("photoId")
	if photoId == "" {
		ctx.Logger.Warning("Error in deletePhoto.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.db.DeletePhoto(ctx.UserId, photoId)
	if err != nil {
		ctx.Logger.WithError(err).Warning("deleting photo user in deletePhoto.go went wrong lol")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}
