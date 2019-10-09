package request

import "net/http"

type Header struct {
	Name string
	Val string
}

func Get(url string, headers []Header) (req *http.Request) {
	return request("GET", url, headers)
}

func Post(url string, headers []Header) (req *http.Request) {
	return request("POST", url, headers)
}

func request(method, url string, headers []Header) (req *http.Request) {
	req, _ = http.NewRequest(method, url, nil)

	for _, h := range headers {
		req.Header.Set(h.Name, h.Val)
	}

	return
}
