package handlers

import (
	"net/http"
	"viewcount/platform"
	"viewcount/services"

	"github.com/gin-gonic/gin"
)

func GetVideos() gin.HandlerFunc {
	return func(c *gin.Context) {
		results := services.GetVideos()
		c.JSON(http.StatusOK, results)
	}
}

// type postVideorequest struct {
// 	Id    string `json:"id"`
// 	Title string `json:"title"`
// 	Views int    `json:"views"`
// }

func PostVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := platform.Video{}
		c.Bind(&requestBody)

		id := services.PostVideo(&requestBody)

		c.JSON(http.StatusOK, id)
	}
}

func GetViews() gin.HandlerFunc {
	return func(c *gin.Context) {
		vars := c.Params
		id, _ := vars.Get("id")

		result, ok := services.GetViews(id)
		if !ok {
			c.JSON(http.StatusBadRequest, "video not found")
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

func IncViews() gin.HandlerFunc {
	return func(c *gin.Context) {
		vars := c.Params
		id, _ := vars.Get("id")

		ok := services.IncViews(id)

		if !ok {
			c.JSON(http.StatusBadRequest, "video not found")
		} else {
			c.Status(http.StatusNoContent)
		}
	}
}

// func ViewCount(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")

// 	fmt.Fprintf(w, "no of views on this video is: %d", countMap[id])
// }

// func IncrCount(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	countMap[id]++
// }
