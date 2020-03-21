package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// p1 := NewProduct("Adam's Peanut Butter")
	// s := NewSite("Amazon", "https://smile.amazon.com/Adams-Natural-CRUNCHY-PEANUT-BUTTER/dp/B00K25O2EI/")
	// _ = p1.AddSite(s)

	// p2 := NewProduct("Clorox Disinfecting Wipes Lemon")
	// s = NewSite("Staples", "https://www.staples.com/Clorox-Disinfecting-Wipes-Lemon-Fresh-75-Count-Canister/product_616321")
	// _ = p2.AddSite(s)
	// s = NewSite("Costco", "https://www.costco.com/clorox-disinfecting-wipes%2c-variety-pack%2c-85-count%2c-5-pack.product.100534416.html")
	// _ = p2.AddSite(s)
	// s = NewSite("Amazon", "https://smile.amazon.com/Clorox-Disinfecting-Wipes-Bleach-Cleaning/dp/B00HSC9F2C/")
	// _ = p2.AddSite(s)

	// j1, err := json.Marshal(p1)
	// if err != nil {
	// 	msg := fmt.Sprintf("p1 json.Marshal err: %v", err)
	// 	panic(msg)
	// }
	// log.Printf("p1: %s\n", string(j1))

	// j2, err := json.Marshal(p2)
	// if err != nil {
	// 	msg := fmt.Sprintf("p2 json.Marshal err: %v", err)
	// 	panic(msg)
	// }
	// log.Printf("p2: %s\n", string(j2))

	pj := "./products.json"
	b, err := ioutil.ReadFile(pj)
	if err != nil {
		msg := fmt.Sprintf("ioutil.ReadFile(%s) err: %v", pj, err)
		panic(msg)
	}

	products := new([]Product)
	err = json.Unmarshal(b, products)
	if err != nil {
		msg := fmt.Sprintf("json.Unmarshal err: %v", err)
		panic(msg)
	}

	j, err := json.Marshal(products)
	if err != nil {
		msg := fmt.Sprintf("products json.Marshal err: %v", err)
		panic(msg)
	}
	log.Printf("products: %s\n", string(j))

	// if err := scrape(*s); err != nil {
	// 	log.Fatalf("main: %v\n", err)
	// }
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
