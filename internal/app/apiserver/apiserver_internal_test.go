// internal is for inner methods and values testing.
package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiServer_HandleHell(t *testing.T) {
	s := New(NewConfig())
	responseWriter := httptest.NewRecorder()
	// 
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(responseWriter, request)

	// assert values from res
	assert.Equal(t, responseWriter.Body.String(), "Hello")
}