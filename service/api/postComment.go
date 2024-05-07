package api

import (
	"NEWASA/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var b struct {
		Comment string `json:"comment"`
	}
	err := json.NewDecoder(r.Body).Decode(&b)

	if b.Comment == "" {
		ctx.Logger.WithError(err).Warning("Error in postComment.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		ctx.Logger.WithError(err).Warning("Error in postComment.go: wrong json recieved")
		w.WriteHeader(http.StatusBadRequest)

	}
	photoId := ps.ByName("photoId")
	currentTime := time.Now()

	// Extract day, month, and year
	date := currentTime.Format("02/01/2006 15:04:05")

	err = rt.db.PostComment(ctx.UserId, photoId, b.Comment, date)
	if err != nil {
		ctx.Logger.WithError(err).Warning("Error in postComment.go: error saving comment to database")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusCreated)
}
