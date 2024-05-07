package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/sessions", rt.wrap(rt.postLogin, 0))
	rt.router.GET("/users/", rt.wrap(rt.getUser, 1))
	rt.router.PUT("/username", rt.wrap(rt.putUsername, 1))
	rt.router.POST("/photos", rt.wrap(rt.postPhoto, 1))
	rt.router.GET("/photos/:photoId", rt.wrap(rt.getPhoto, 1))
	rt.router.DELETE("/photos/:photoId", rt.wrap(rt.deletePhoto, 1))
	rt.router.GET("/followings/:userId", rt.wrap(rt.getFollowings, 1))
	rt.router.PUT("/followings/:userId", rt.wrap(rt.putFollowing, 1))
	rt.router.DELETE("/followings/:userId", rt.wrap(rt.deleteFollowing, 1))
	rt.router.GET("/followers/:userId", rt.wrap(rt.getFollowers, 1))
	rt.router.PUT("/banned/:bannedId", rt.wrap(rt.putBanned, 1))
	rt.router.DELETE("/banned/:bannedId", rt.wrap(rt.deleteBanned, 1))
	rt.router.GET("/banned/:userId", rt.wrap(rt.getBanned, 1))
	rt.router.PUT("/photos/:photoId/likes/self", rt.wrap(rt.putLike, 1))
	rt.router.DELETE("/photos/:photoId/likes/self", rt.wrap(rt.deleteLike, 1))
	rt.router.POST("/photos/:photoId/comments", rt.wrap(rt.postComment, 1))
	rt.router.DELETE("/comments/:commentId", rt.wrap(rt.deleteComment, 1))
	rt.router.GET("/stream", rt.wrap(rt.getStream, 1))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
