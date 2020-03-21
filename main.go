package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func main() {
	p := NewProduct("Adam's Peanut Butter")
	p.ID = uuid.New()
	s := NewSite("Amazon")
	_ = p.AddSite(s)

	if err := scrape(*s); err != nil {
		log.Fatalf("main: %v\n", err)
	}
}

func scrape(s Site) error {

	apiKey := os.Getenv("SCRAPESTACK_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("scrape: SCRAPESTACK_API_KEY not defined")
	}

	httpClient := http.Client{}

	req, err := http.NewRequest("GET", "http://api.scrapestack.com/scrape", nil)
	if err != nil {
		return fmt.Errorf("scrape: %v", err)
	}

	q := req.URL.Query()
	q.Add("access_key", apiKey)
	q.Add("url", s.URL)
	req.URL.RawQuery = q.Encode()

	res, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("scrape: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("scrape: %v", err)
		}
		websiteContent := string(bodyBytes)
		fmt.Println(websiteContent)
	}

	return nil
}
