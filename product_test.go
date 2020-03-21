package main

import (
	"testing"

	"github.com/google/uuid"
)

func TestProductName(t *testing.T) {
	expected := "Adam's Peanut Butter"
	p := NewProduct(expected)

	got := p.Name
	if got != expected {
		t.Errorf("expected name %s, got %s", expected, got)
	}
}

func TestProductAddSite(t *testing.T) {
	pn := "Adam's Peanut Butter"
	p := NewProduct(pn)

	sn := "Amazon"
	s1 := &Site{Name: sn}
	s1.ID = uuid.New()
	if err := p.AddSite(s1); err != nil {
		t.Errorf("shouldn't happen, AddSite(%s) err: %v", s1.Name, err)
	}

	s2 := &Site{Name: sn}
	s2.ID = uuid.New()
	err := p.AddSite(s2)
	if err != ErrDuplicateSite {
		t.Errorf("expected %v, got %v", ErrDuplicateSite, err)
	}

	sn = "Costco"
	s3 := &Site{Name: sn}
	s3.ID = uuid.New()
	err = p.AddSite(s3)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	// log.Printf("p: %+v, Sites: %+v\n", p, p.Sites)
}
