package request

import "net/http"

func Get(url string, headers interface{}) (req *http.Request) {
	return request("GET", url, headers)
}

func Post(url string, headers interface{}) (req *http.Request) {
	return request("POST", url, headers)
}

type headersIterator interface {
	Iterate(f func(name, val string))
}

func request(method, url string, headers interface{}) (req *http.Request) {
	req, _ = http.NewRequest(method, url, nil)

	switch headers.(type) {
	case headersIterator:
		headers.(headersIterator).Iterate(func(name, val string) { req.Header.Set(name, val) })
	default:
		panic("Headers should satisfy interface headersIterator")
	}

	return
}
