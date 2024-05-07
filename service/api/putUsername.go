package api

import (
	"NEWASA/service/api/reqcontext"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var b struct {
		UsernameNew string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&b)
	if b.UsernameNew == "" {
		ctx.Logger.Warning("Error in changeUsername.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		ctx.Logger.WithError(err).Warning("Error in changeUsername.go: wrong json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.ChangeUsername(ctx.UserId, b.UsernameNew)
	if err != nil {
		ctx.Logger.WithError(err).Warning("Error in changeUsername.go: error changing username")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
