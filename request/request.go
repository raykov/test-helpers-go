package request

import "net/http"

func Get(url string, headers interface{}) (req *http.Request) {
	return request("GET", url, headers)
}

func Post(url string, headers interface{}) (req *http.Request) {
	return request("POST", url, headers)
}

func request(method, url string, headers interface{}) (req *http.Request) {
	var checkedHeaders []struct {Name string; Val string}
	checkedHeaders, ok := headers.([]struct {Name string; Val string})
	if !ok {
		panic("Headers should be an array of struct with Name and Val fields")
	}

	req, _ = http.NewRequest(method, url, nil)

	for _, h := range checkedHeaders {
		req.Header.Set(h.Name, h.Val)
	}

	return
}
