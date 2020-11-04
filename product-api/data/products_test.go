package data

import "testing"

func TestChecksvalidation(t *testing.T) {
	p := &Product{
		Name:  "something",
		Price: 1.00,
		SKU:   "abc-def-hij",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
