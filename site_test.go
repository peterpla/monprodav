package main

import "testing"

func TestNewSite(t *testing.T) {
	expected := "Amazon"
	s := NewSite(expected)

	got := s.Name
	if got != expected {
		t.Errorf("expected Name %s, got %s", expected, got)
	}
}

func TestSiteInStock(t *testing.T) {
	p := Product{}
	s := NewSite("Amazon")
	expected := true

	got := s.InStock(p)
	if got != expected {
		t.Errorf("expected InStock %t, got %t", expected, got)
	}
}

func TestSitePrice(t *testing.T) {
	// p := &Product{}
	s := NewSite("Amazon")
	expected := float32(0.00)

	got := s.Price
	if got != expected {
		t.Errorf("expected Price %f, got %f", expected, got)
	}
}

func TestSiteQuantity(t *testing.T) {
	// p := &Product{}
	s := NewSite("Amazon")
	expected := float32(0.00)

	got := s.Quantity
	if got != expected {
		t.Errorf("expected Quantity %f, got %f", expected, got)
	}
}
