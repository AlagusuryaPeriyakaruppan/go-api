// simple RESTful web service

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Post struct {
	UserId int    `json:"userId"`
	ID     int    `json:"id"`
	Tname  string `json:"title"`
	Body   string `json:"body"`
}

type NewPost struct {
	UserId int    `json:"userId"`
	ID     int    `json:"id"`
	Tname  string `json:"tname"`
	Body   string `json:"body"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if error := c.BindJSON(&newAlbum); error != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, value := range albums {
		if value.ID == id {
			c.IndentedJSON(http.StatusOK, value)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ALBUM NOT FOUND!"})
}
func getPosts(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	var posts []Post
	
	error := json.Unmarshal(body, &posts)
	if error != nil {
		fmt.Println("ERROR UMARSHLINH", error)
	}

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	var newPosts = []NewPost{

		{UserId: posts[0].UserId,
			ID:    posts[0].ID,
			Tname: posts[0].Tname,
			Body:  posts[0].Body},
	}

	c.IndentedJSON(http.StatusOK, newPosts)

}

func main() {
	router := gin.Default()
	// router.GET("/albums",getAlbums)
	// router.POST("/albums",postAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	router.GET("/posts", getPosts)
	router.Run("localhost:8080")
}

