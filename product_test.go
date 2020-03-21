package main

import "testing"

func TestProductName(t *testing.T) {
	expected := "Adam's Peanut Butter"
	p := NewProduct(expected)

	got := p.Name
	if got != expected {
		t.Errorf("expected name %s, got %s", expected, got)
	}
}
