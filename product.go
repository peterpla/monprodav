package main

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

// Product holds the sites to check for a product
type Product struct {
	Name  string    `json:"name"`  // Friendly product name
	ID    uuid.UUID `json:"id"`    // unique ID for this product
	Sites *[]Site   `json:"sites"` // sites to monitor for availability of this product
}

// NewProduct creates a new Product structure
func NewProduct(Name string) *Product {
	s := new([]Site)
	p := &Product{
		Name:  Name,
		Sites: s,
	}
	p.ID = uuid.New()
	log.Printf("NewProduct, p: %+v, Sites: %+v\n", p, p.Sites)
	return p
}

func (p *Product) AddSite(s *Site) error {
	for _, rs := range *p.Sites {
		if rs.Name == s.Name {
			return ErrDuplicateSite
		}
	}
	*p.Sites = append(*p.Sites, *s)
	log.Printf("AddSite: %+v\n", p.Sites)
	return nil
}

var ErrDuplicateSite = errors.New("Duplicate Site")
