package routes

import (
	"viewcount/handlers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/", handlers.GetVideos())
	r.POST("/add", handlers.PostVideo())
	r.GET("/get_views/:id", handlers.GetViews())
	r.POST("/inc_views/:id", handlers.IncViews())

	return r
}

// r := chi.NewRouter()

// r.Route("/view_count", func(r chi.Router) {
// 	r.Get("/{id}", handlers.ViewCount)
// })

// r.Route("/incr_count", func(r chi.Router) {
// 	r.Post("/{id}", handlers.IncrCount)
// })

// _ = http.ListenAndServe(portNumber, r)
