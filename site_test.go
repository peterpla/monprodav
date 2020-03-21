package main

import (
	"testing"
)

func TestNewSite(t *testing.T) {
	u := "https://smile.amazon.com/Adams-Natural-CRUNCHY-PEANUT-BUTTER/dp/B00K25O2EI/"
	expected := "Amazon"
	s := NewSite(expected, u)

	got := s.Name
	if got != expected {
		t.Errorf("expected Name %s, got %s", expected, got)
	}
}

func TestSiteAddProduct(t *testing.T) {
	u := "https://smile.amazon.com/Adams-Natural-CRUNCHY-PEANUT-BUTTER/dp/B00K25O2EI/"
	sn := "Amazon"
	s := NewSite(sn, u)

	pn := "Adam's Peanut Butter"
	p := NewProduct(pn)

	if err := s.AddProduct(p); err != nil {
		t.Errorf("unexpected err %v", err)
	}
	if s.Product.Name != pn {
		t.Errorf("expected Name %s, got %s", pn, s.Product.Name)
	}
	// log.Printf("Site: %+v\n", s)
}

func TestSiteInStock(t *testing.T) {
	p := Product{}
	u := "https://smile.amazon.com/Adams-Natural-CRUNCHY-PEANUT-BUTTER/dp/B00K25O2EI/"
	s := NewSite("Amazon", u)
	expected := true

	got := s.InStock(p)
	if got != expected {
		t.Errorf("expected InStock %t, got %t", expected, got)
	}
}

func TestSitePrice(t *testing.T) {
	u := "https://smile.amazon.com/Adams-Natural-CRUNCHY-PEANUT-BUTTER/dp/B00K25O2EI/"
	s := NewSite("Amazon", u)
	expected := float32(0.00)

	got := s.Price
	if got != expected {
		t.Errorf("expected Price %f, got %f", expected, got)
	}
}

func TestSiteQuantity(t *testing.T) {
	u := "https://smile.amazon.com/Adams-Natural-CRUNCHY-PEANUT-BUTTER/dp/B00K25O2EI/"
	s := NewSite("Amazon", u)
	expected := float32(0.00)

	got := s.Quantity
	if got != expected {
		t.Errorf("expected Quantity %f, got %f", expected, got)
	}
}
