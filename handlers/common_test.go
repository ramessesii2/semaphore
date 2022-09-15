package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/RAMESSESII2/test-go-gin/models"
	"github.com/gin-gonic/gin"
)

var tmpArticleList []models.Article

// This function is used for setup before executing the test function
func TestMain(m *testing.M) {
	// Get Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		path, _ := os.UserHomeDir()
		r.LoadHTMLGlob(path + "/RAMESSESII2/test-go-gin/templates/*")
	}
	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This function is used to store the main list into a temporary one
func saveLists() {
	tmpArticleList = models.Articles
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	models.Articles = tmpArticleList
}
