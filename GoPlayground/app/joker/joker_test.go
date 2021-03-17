package joker_test

import (
	. "GoPlayground/app/joker"
	"GoPlayground/app/joker/jokerfakes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJoker_MakeJoke_HandleError(t *testing.T) {
	const expectedResponse = `{"value"}`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/joke", req.URL.String())
		_, _ = rw.Write([]byte(expectedResponse))
	}))
	defer server.Close()

	j := NewJoker(server.URL + "/joke", MyHttpClient{})
	joke, status := j.MakeJoke()

	assert.Equal(t, 500, status)
	assert.Equal(t, "", joke)
}

func TestJoker_MakeJoke_HandleResponse(t *testing.T) {
	const expectedResponse = "Chuck Norris fährt in England auf der rechten Seite!"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/joke", req.URL.String())
		_, _ = rw.Write([]byte(`{"value": "` + expectedResponse + `"}"`))
	}))
	defer server.Close()

	j := NewJoker(server.URL + "/joke", MyHttpClient{})
	joke, status := j.MakeJoke()

	assert.Equal(t, 200, status)
	assert.Equal(t, expectedResponse, joke)
}

func TestJoker_MakeJoke_HandleMockedResponse(t *testing.T) {
	const expectedResponse = "Chuck Norris fährt in England auf der rechten Seite!"
	const expectedHttpCode = 200
	mock := &jokerfakes.FakeHttpClient{}
	mock.GetReturns(expectedResponse, expectedHttpCode)

	j := NewJoker("", mock)
	joke, status := j.MakeJoke()

	assert.Equal(t, 200, status)
	assert.Equal(t, expectedResponse, joke)
}
