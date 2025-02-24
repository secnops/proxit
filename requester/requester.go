package requester

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"time"
)

func Request(method, url_, body string, headers map[string]string) string {

	//fmt.Println("method:", method)
	//fmt.Println("url:", url)
	//fmt.Println("body:", body)
	//fmt.Println("headers:", headers)

	/**/
	/*
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(&url.URL{
					Scheme: "http:",
					Host:   "127.0.0.1:8080",
				}),
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			//Timeout: 10 * time.Second,
		}
	*/

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, /* #nosec */
		},
	}
	var req *http.Request
	var err error
	if method == "GET" {

		req, err = http.NewRequest("GET", url_, nil)
		if err != nil {
			log.Fatal("ERROR -> ", "Couldn't initialize the request, check the parameters. Error: %+v", err)
			log.Fatalf("ERROR -> Couldn't initialize the requester")
		}

	} else {
		//fmt.Println("Request-> ", body)
		send_body := []byte(body)
		req, err = http.NewRequest(method, url_, bytes.NewBuffer(send_body))
		if err != nil {
			log.Fatal("ERROR -> ", "Couldn't initialize the request, check the parameters. Error: %+v", err)
		}
	}

	for name, content := range headers {
		req.Header.Add(name, content)
	}
	//check

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("ERROR -> ", "Couldn't send request. Error:%+v", err)
	}
	defer resp.Body.Close()
	body_response, err := io.ReadAll(resp.Body)
	//fmt.Println("Response -> ", string(body_response))
	if err != nil {
		log.Fatal("ERROR -> ", "Couldn't parse response body. Error: %+v", err)

	}
	/*
		var data_body []byte

		if err := json.Unmarshal([]byte(body_response), &data_body); err != nil {
			//log.Fatalf("Couldn't build the JSON. Error: %+v", err)
			return string(body_respvarse)
		}
	*/
	return string(body_response)
}
