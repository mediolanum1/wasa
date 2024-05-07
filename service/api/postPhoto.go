package api

import (
	"NEWASA/service/api/reqcontext"

	"time"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Accessing files
	file, _, err := r.FormFile("photo")
	if err != nil {
		ctx.Logger.Warning("Error in postPhoto.go: wrong file format ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	description := r.FormValue("caption")
	currentTime := time.Now()
	defer file.Close()

	// Getting current date
	formattedDate := currentTime.Format("02/01/2006 15:04:05")

	err = rt.db.PostPhoto(ctx.UserId, file, description, formattedDate)
	if err != nil {
		ctx.Logger.WithError(err).Warning("Error in postPhoto.go: error saving photo to database")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
