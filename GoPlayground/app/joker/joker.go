package joker

import (
	. "GoPlayground/app/jsondecoder"
	"net/http"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type Joker interface {
	MakeJoke() (string, int)
}

type joker struct {
	apiUrl     string
	httpClient HttpClient
}

func (j joker) MakeJoke() (string, int) {
	joke, code := j.httpClient.Get(j.apiUrl)
	return joke, code
}

func NewJoker(apiUrl string, httpClient HttpClient) Joker {
	return joker{
		apiUrl,
		httpClient,
	}
}

//counterfeiter:generate . HttpClient
type HttpClient interface {
	Get(url string) (response string, code int)
}

type MyHttpClient struct {
}

func (c MyHttpClient) Get(url string) (response string, code int) {
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return "", 500
	}
	defer resp.Body.Close()

	d := MyJsonDecoder{}

	body, err := d.Decode(resp.Body)
	if err != nil {
		return "", 500
	}
	joke := body["value"]

	return joke, 200
}
