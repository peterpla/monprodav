package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

// Site holds site-specific properties of a Product
type Site struct {
	Name     string    `json:"name"`    // Friendly name of the site
	ID       uuid.UUID `json:"id"`      // unique ID for this Site
	URL      string    `json:"url"`     // this site's URL for product
	Body     string    `json:"body"`    // HTML body returned for URL
	Quantity float32   `json:"qty"`     // product Quantity (in Units) for URL
	Units    string    `json:"units"`   // units of measure, e.g., ounce, pound, count
	Price    float32   `json:"price"`   // product Price for URL
	Product  *Product  `json:"product"` // reference to product
}

// NewSite creates a new Site structure
func NewSite(name string, url string) *Site {
	s := &Site{
		Name: name,
		URL:  url,
	}
	s.ID = uuid.New()
	log.Printf("NewSite: %+v\n", *s)
	return s
}

func (s *Site) AddProduct(p *Product) error {
	s.Product = p
	return nil
}

// InStock dispatches to the site-specific InStock method
func (s Site) InStock(p Product) bool {
	switch s.Name {
	case "Amazon":
		return AmazonInStock(p)

	default:
		msg := fmt.Sprintf("InStock: unsupported site %s", s.Name)
		panic(msg)
	}
}

func AmazonInStock(p Product) bool {
	return true
}

func AmazonPrice(p Product) float32 {
	return 1.00
}

func AmazonUnitPrice(p Product) float32 {
	return 0.10
}
