package api

import (
	"NEWASA/service/api/reqcontext"
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "image/png")

	photoId := ps.ByName("photoId")
	if photoId == "" {
		ctx.Logger.Warning("Error in getPhoto.go: empty json recieved")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := rt.db.GetPhoto(photoId)

	if err != nil {
		ctx.Logger.WithError(err).Warning("error in getPhoto.go: getting photo went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}
	path := "/tmp/photos/" + photoId + ".png"
	photofile, err := os.Open(path)
	if err != nil {
		ctx.Logger.WithError(err).Warning("error in getPhoto.go: opening photo went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}

	buf := bytes.NewBuffer(nil)

	_, err = io.Copy(buf, photofile)
	if err != nil {
		ctx.Logger.WithError(err).Warning("error in getPhoto.go: copying photo went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		ctx.Logger.WithError(err).Warning("error in getPhoto.go: listing photo went wrong")
		w.WriteHeader(http.StatusInternalServerError)
	}

}
