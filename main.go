package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type version struct {
	ID string `json:"id"`
	Url string `json:"Url"`
	Contributor string `json:"contributor"`
	Version float64 `json:"Version"`
}

var versions = []version{
	{ID: "1", Url: "Blue Train", Contributor: "John Coltrane", Version: 1.0},
	{ID: "2", Url: "Jeru", Contributor: "Gerry Mulligan", Version: 1.001},
	{ID: "3", Url: "Sarah Vaughan and Clifford Brown", Contributor: "Sarah Vaughan", Version: 1.002},
}

func getVersions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, versions)
}

func postVersions(c *gin.Context) {
	var newVersion version 

	if err := c.BindJSON(&newVersion); err != nil {
		return
	}

	versions = append(versions, newVersion)
	c.IndentedJSON(http.StatusCreated, newVersion)
}

func getVersionByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range versions {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "version not found"})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
            c.HTML(http.StatusOK, "index.html", gin.H{
                "Title":   "Gin Template Example",
                "Message": "Hello, World!",
            })
        })
	router.GET("/versions", getVersions)
	router.GET("/version/:id", getVersionByID)
	router.POST("/versions", postVersions)

	router.Run("0.0.0.0:8080")
}
