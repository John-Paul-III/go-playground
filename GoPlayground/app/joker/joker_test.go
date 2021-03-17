package joker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoStuffWithTestServer(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/joke", req.URL.String())
		rw.Write([]byte(`{"value","some_joke"}}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	//api := API{server.Client(), server.URL}

	//joker := NewJoker("https://api.chucknorris.io/jokes/random")
	fmt.Println("SERVER URL:", server.URL)
	joker := NewJoker(server.URL + "/joke")
	joke, _ := joker.MakeJoke()

	assert.Equal(t, "something", joke)
	//assert.Nil(t, err)
}
