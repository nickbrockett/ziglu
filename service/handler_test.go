package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var testFeedItems Items


// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func TestFeedItemNotFound(t *testing.T) {

	assert.Nil(t, json.Unmarshal(testItems, &testFeedItems))

	r := getRouter(true)

	r.GET("/feeds/:feedID", searchFeed)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/feeds/abc", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {

		return w.Code == http.StatusNotFound
	})
}

func TestFeedItemFound(t *testing.T) {
	r := getRouter(true)

	r.GET("/feeds/:feedID", searchFeed)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/feeds/abc", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {

		t.Logf("njb we have %d ", w.Code)
		//
		//statusOK := w.Code == http.StatusOK
		//
		//// Test that the page title is "Home Page"
		//// You can carry out a lot more detailed tests using libraries that can
		//// parse and process HTML pages
		//p, err := ioutil.ReadAll(w.Body)
		//pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0
		//
		//return statusOK && pageOK
		// Test that the http status code is 404
		return w.Code == http.StatusNotFound
	})
}

var testItems = []byte(`
    {
        "Title": "Easter celebrations set to rival Christmas - even down to the tree",
        "Link": "https://www.bbc.co.uk/news/business-56541002",
        "Desc": "Easter trees and garden furniture are being sought out by shoppers keen to make the most of lockdown easing.",
        "PubDate": "Sun, 28 Mar 2021 00:03:58 GMT",
        "Key": "7vj8dXoxvlaV_lHEjqktVaxlNvrVtuUQdlozvZmD0fs="
    },
    {
        "Title": "How the 'world's worst sniffer dog' is helping the NHS",
        "Link": "https://www.bbc.co.uk/news/uk-england-london-56375874",
        "Desc": "Dexter used to work as a sniffer dog for the Met but he was \"too sociable\" for the job.",
        "PubDate": "Sun, 28 Mar 2021 00:00:57 GMT",
        "Key": "xUmj7E0DXrEFqdp_MEzM27rnNzirPBlrRSZ0v8bkPlA="
    },
`)