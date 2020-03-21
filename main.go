package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	apiKey := os.Getenv("SCRAPESTACK_API_KEY")
	if apiKey == "" {
		panic("SCRAPESTACK_API_KEY not defined")
	}

	httpClient := http.Client{}

	req, err := http.NewRequest("GET", "http://api.scrapestack.com/scrape", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("access_key", apiKey)
	q.Add("url", "http://scrapestack.com")
	req.URL.RawQuery = q.Encode()

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		websiteContent := string(bodyBytes)
		fmt.Println(websiteContent)
	}
}
