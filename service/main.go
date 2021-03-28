package main

import (
	"github.com/gin-gonic/gin"
	"priva.te/ziglu/lib"
)

var (
	router     *gin.Engine
	emptyArray = []byte("[]")
)

func init() {
	lib.DefaultProviders()
}

func main() {

	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	router.Run() // nolint:errcheck
}
