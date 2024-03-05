package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)
// #1 The main function is the entry point, much like public class main in Java and def main(): in python
func main() {
    // set the router to the default provided by Gin. The router listens for incoming http requests and matches them to the routes we define in the following lines
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)
    // run the server locally
    router.Run("localhost:8080")
}

// #1 album represents data about a record album. The `json:"id"` is a struct tag, which is a mechanism to annotate the struct fields with metadata that can be used to reflect on the struct. In this case, it's used to specify the JSON key for the struct field. 
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice (Slices are a go datatype that reference an array)
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
// This takes one parameter which points to the struct gin.Context. *Pointers* in Go allow you to indirectly access and modify the value of a variable by referring to its memory address. By using a pointer, you can avoid making a copy of the entire object when passing it to a function, which can be more efficient, especially for large objects.
// The pointer is of the type gin context
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// post albums
func postAlbums(c *gin.Context) {
    // create the new variable to hold the received JSON using the pre-defined album struct
    var newAlbum album

    // Call BindJSON to bind the received JSON to newAlbum. if there is an error return nil! null! nothing!
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice. (In go, a slice is a 'view' of an underlying array. It's a reference to the array. So, when you append to a slice, you're actually appending to the underlying array.)
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// The handler to return a specific album
// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id") // This line retrieves the 'id' Parameter from the URL

    // Loop over the list of albums, looking for an album whose ID value matches the parameter.
    for _, a := range albums { // the _ is a blank identifier, which is used to tell the compiler that this variable is not going to be used, its placeholder so the second variable can be used. This is a placeholder for the initialization. For instance if creating an iterator you would use this to initialize the iterator (i :=0 ). 
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}