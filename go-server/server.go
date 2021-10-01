package main

import (
	"awesomeproject/goserver/api"
	"awesomeproject/goserver/wiki"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// RUN GIN SERVER
	router := gin.Default()
	router.GET("/albums", api.GetAlbums)
	router.GET("/albums/:id", api.GetAlbumById)
	router.POST("/albums", api.PostAlbums)
	router.Run("localhost:8080")

	// RUN GO DEFAULT SERVER
	http.HandleFunc("/view/", wiki.ViewHandler)
	http.HandleFunc("/edit/", wiki.EditHandler)
	http.HandleFunc("/save/", wiki.SaveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
