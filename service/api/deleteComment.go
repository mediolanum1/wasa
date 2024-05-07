package api

import (
	"NEWASA/service/api/reqcontext"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	commentId := ps.ByName("commentId")
	if commentId == "" {
		ctx.Logger.Warning("Error in deleteComment.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := rt.db.DeleteComment(commentId, ctx.UserId)
	if err != nil {
		ctx.Logger.WithError(err).Warning("deleting comment in deleteComment.go went wrong lol")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}
