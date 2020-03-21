package main

import (
	"github.com/google/uuid"
)

// Product holds the sites to check for a product
type Product struct {
	Name  string    `json:"name"`  // Friendly product name
	ID    uuid.UUID `json:"id"`    // unique ID for this product
	Sites []Site    `json:"sites"` // sites to monitor for availability of this product
}

// NewProduct creates a new Product structure
func NewProduct(Name string) *Product {
	return &Product{
		Name: Name,
	}
}

func (p Product) AddSite(s *Site) error {
	// if s.Name does not already exist
	p.Sites = append(p.Sites, *s)
	return nil
}
