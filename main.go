package main

import (
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type version struct {
	ID string `json:"id"`
	Url string `json:"download_url"`
	Contributor string `json:"contributor"`
	Version int16 `json:"version"`
}

var versions = []version{
	{ID: "1", Url: "https://web-service-gin-0z21.onrender.com/static/assets/version_11.apk", Contributor: "Sarah Vaughan", Version: 1},
}

func getVersions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, versions)
}

func getLatestVersion(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, versions[len(versions)-1])
}

func postVersions(c *gin.Context) {
	var newVersion version

	if err := c.ShouldBindJSON(&newVersion); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}

    router.Use(cors.New(config))
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
            c.HTML(http.StatusOK, "index.html", gin.H{
                "Title":   "Versions",
                "Message": versions,
            })
        })
	router.GET("/versions", getVersions)
	router.GET("/latest", getLatestVersion)
	router.GET("/version/:id", getVersionByID)
	router.POST("/versions", postVersions)

	router.Run()
}
