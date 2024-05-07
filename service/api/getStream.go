package api

import (
	"NEWASA/service/api/reqcontext"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	stream, err := rt.db.GetStream(ctx.UserId)
	if err != nil {
		ctx.Logger.WithError(err).Warning("error in getStream: getting stream went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(stream)
	if err != nil {
		ctx.Logger.WithError(err).Warning("listing stream went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
