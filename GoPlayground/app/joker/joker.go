package joker

import (
	. "GoPlayground/app/jsondecoder"
	"net/http"
)

type Joker struct {
	apiUrl     string
	httpClient httpClient
}

func (j Joker) getApiUrl() string {
	return j.apiUrl
}

func (j Joker) setApiUrl(s string) {
	j.apiUrl = s
}

func NewJoker(apiUrl string) Joker {
	return Joker{
		apiUrl,
		MyHttpClient{},
	}
}

func (j Joker) MakeJoke() (string, int) {
	joke, code := j.httpClient.Get(j.apiUrl)
	return *joke, code
}

type httpClient interface {
	Get(url string) (response *string, code int)
}

type MyHttpClient struct {
}

func (c MyHttpClient) Get(url string) (response *string, code int) {
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, 500
	}
	defer resp.Body.Close()

	d := MyJsonDecoder{}

	body, err := d.Decode(resp.Body)
	if err != nil {
		return nil, 500
	}
	joke := body["value"]

	return &joke, 200
}
