package scraper

import "testing"

func TestReadCompanyNames(t *testing.T) {
	names, err := ReadCompanyNames("data/input.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(names) != 10 {
		t.Fatalf("expected 3 names, got %d", len(names))
	}

	expected := "Kampala Digital Hub"
	if names[0] != expected {
		t.Errorf("expected %s, got %s", expected, names[0])
	}
}
