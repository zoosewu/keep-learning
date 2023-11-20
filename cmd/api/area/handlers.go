package area

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"createTime"`
}

// albums slice to seed record album data.
var albums = []album{
	{Id: 1, Name: "All Train"},
	{Id: 2, Name: "Blue Train"},
	{Id: 3, Name: "Clue Train"},
}

// getAlbums responds with the list of all albums as JSON.
func GetAreas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func PostArea(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAreaByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err == nil {
		for _, a := range albums {
			if a.Id == id {
				context.IndentedJSON(http.StatusOK, a)
				return
			}
		}
	}
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
