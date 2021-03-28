package main

import (
	"net/http"
	"priva.te/ziglu/lib"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {

	// Handle the get sources request, which are the sources of RSS feeds, i.e. Provider and Categories.
	router.GET("/sources", getSources)

	// Handle the set sources request, which are the sources of RSS feeds, i.e. Provider and Categories.
	router.POST("/sources", setSources)

	// Handle the RSS feed search request
	router.GET("/feeds", searchFeeds)

	// Handle the RSS feed search request for a specific feed item
	router.GET("/feeds/:feedID", searchFeed)

}

func getSources(c *gin.Context) {

	feedSources := lib.GetProviders()
	if feedSources == nil {
		c.Data(http.StatusOK, "application/json", emptyArray)
		return
	}
	c.JSON(http.StatusOK, feedSources)

}

func setSources(c *gin.Context) {

	newProviders := lib.Providers{}

	if err := c.BindJSON(&newProviders); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err) // nolint:errcheck
		return
	}

	lib.SetProviders(&newProviders)

	feedSources := lib.GetProviders()
	if feedSources == nil {
		c.Data(http.StatusNoContent, "application/json", emptyArray)
		return
	}
	c.JSON(http.StatusCreated, feedSources)

}

func searchFeeds(c *gin.Context) {

	sources := lib.GetNewsSources(c.Query("provider"), c.Query("category"))
	r := NewReader(sources, RSSReader)
	r.Read()

	if r.items == nil || len(r.items) == 0 {
		c.Data(http.StatusOK, "application/json", emptyArray)
		return
	}

	c.JSON(http.StatusOK, r.items)

}

func searchFeed(c *gin.Context) {

	searchKey := c.Param("feedID")

	if v, ok := itemCache.Get(searchKey); ok {
		item := v.(Item)
		displayHTMLItem(c, &item)
		return
	}

	sources := lib.GetNewsSources("", "")
	r := NewReader(sources, RSSReader)
	r.Read()

	for _, v := range r.items {
		if v.Key == searchKey {
			itemCache.Add(searchKey, *v)
			displayHTMLItem(c, v)
			return

		}
	}

	displayError(c)

}

func displayHTMLItem(c *gin.Context, item *Item) {

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the feed.html template
		"feedItem.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title":   "News Item",
			"payload": item,
		},
	)
}

func displayError(c *gin.Context) {

	c.HTML(
		// Set the HTTP status to 404 (Not Found)
		http.StatusNotFound,
		// Use the error.html template
		"error.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "404 Not Found",
		},
	)
}
