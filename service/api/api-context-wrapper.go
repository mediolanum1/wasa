package api

import (
	"NEWASA/service/api/reqcontext"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// httpRouterHandler is the signature for functions that accepts a reqcontext.RequestContext in addition to those
// required by the httprouter package.
type httpRouterHandler func(http.ResponseWriter, *http.Request, httprouter.Params, reqcontext.RequestContext)

// wrap parses the request and adds a reqcontext.RequestContext instance related to the request.
func (rt *_router) wrap(fn httpRouterHandler, authRequired int) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		reqUUID, err := uuid.NewV4()
		if err != nil {
			rt.baseLogger.WithError(err).Error("can't generate a request UUID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var ctx = reqcontext.RequestContext{
			ReqUUID: reqUUID,
		}

		// Create a request-specific logger
		ctx.Logger = rt.baseLogger.WithFields(logrus.Fields{
			"reqid":     ctx.ReqUUID.String(),
			"remote-ip": r.RemoteAddr,
		})

		if authRequired == 1 {
			// check for the authorization
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				ctx.Logger.Warning("authorization error: empty authHeader")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)

			// Check if the Authorization header has the correct format
			if len(parts) != 2 || parts[0] != "Bearer" {
				ctx.Logger.Warning("authorization error: auth wrong format")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Extract the token from the Authorization header

			valid, err := rt.db.PostToken(parts[1])
			if err != nil {
				ctx.Logger.WithError(err).Warning("authorization error: coudnt login")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if !valid {
				ctx.Logger.WithError(err).Warning("authorization error: provided token is not valid")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx.UserId = parts[1]

		}
		// Call the next handler in chain (usually, the handler function for the path)
		fn(w, r, ps, ctx)

	}
}
