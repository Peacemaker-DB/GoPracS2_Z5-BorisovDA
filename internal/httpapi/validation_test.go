package httpapi

import "testing"

func TestParsePositiveID(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		ok   bool
	}{
		{name: "valid", raw: "1", ok: true},
		{name: "zero", raw: "0", ok: false},
		{name: "negative", raw: "-1", ok: false},
		{name: "sql injection", raw: "1 OR 1=1", ok: false},
		{name: "letters", raw: "abc", ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := parsePositiveID(tt.raw)
			if ok != tt.ok {
				t.Fatalf("expected %v, got %v", tt.ok, ok)
			}
		})
	}
}

func TestNormalizeEmail(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		ok   bool
	}{
		{name: "valid", raw: "Ivanov@Example.com", ok: true},
		{name: "without at", raw: "ivanov.example.com", ok: false},
		{name: "sql injection", raw: "a@example.com' OR '1'='1", ok: false},
		{name: "empty", raw: "", ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := normalizeEmail(tt.raw)
			if ok != tt.ok {
				t.Fatalf("expected %v, got %v", tt.ok, ok)
			}
		})
	}
}