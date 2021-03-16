package joker

import (
	"io"
	"net/http"
)
type Joker struct {
	apiUrl   string
	httpClient httpClient
}

func NewJoker(apiUrl string) Joker {
	return Joker{
		apiUrl,
		&http.Client{},
	}
}

func (j Joker) MakeJoke() string {
	//const response = j.httpClient.Get(j.apiUrl)
	//response
	return "nil"
}

type httpClient interface {
	Get(url string) (reader string, code int, err error)
}

type MyHttpClient struct {

}

func (c MyHttpClient) Get(url string) (reader string, code int, error error) {
	v := http.Client
	_, err := &v.Get("lol")
}
