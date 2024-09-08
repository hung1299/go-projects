package proxy

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func Forward(req *http.Request, url string) *http.Response {
	hr, err := http.NewRequest(req.Method, url+req.URL.Path, req.Body)
	if err != nil {
		fmt.Println("something wrong when create new request!!", err)
		return nil
	}
	hr.Header = req.Header

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	c := &http.Client{Transport: tr}

	res, err := c.Do(hr)
	if err != nil {
		fmt.Println("error occurred!!", err)
		return nil
	}

	return res
}
