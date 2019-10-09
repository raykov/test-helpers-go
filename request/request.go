package request

import "net/http"

type header struct {
	Name string
	Val string
}

func Get(url string, headers interface{}) (req *http.Request) {
	return request("GET", url, headers)
}

func Post(url string, headers interface{}) (req *http.Request) {
	return request("POST", url, headers)
}

func request(method, url string, headers interface{}) (req *http.Request) {
	var checkedHeaders []header
	checkedHeaders, ok := headers.([]header)
	if !ok {
		panic("Headers should be an array of struct with Name and Val fields")
	}

	req, _ = http.NewRequest(method, url, nil)

	for _, h := range checkedHeaders {
		req.Header.Set(h.Name, h.Val)
	}

	return
}
