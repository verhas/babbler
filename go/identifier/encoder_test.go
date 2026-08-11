package identifier

import (
	"strings"
	"testing"
)

func TestNumberToIDZeroReturnsValidID(t *testing.T) {
	id, err := NumberToID(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	words := strings.Split(id, " ")
	if len(words) != 2 || len(words[0]) != 4 || len(words[1]) != 4 {
		t.Fatalf("expected two 4-letter words, got %q", id)
	}
}

func TestNumberToIDOneDiffersFromZero(t *testing.T) {
	id0, _ := NumberToID(0)
	id1, _ := NumberToID(1)
	if id0 == id1 {
		t.Fatalf("expected different ids, got %q for both", id0)
	}
}

func TestNumberToID10000MatchesReference(t *testing.T) {
	id, err := NumberToID(10000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != "Talo Buno" {
		t.Fatalf("expected %q, got %q", "Talo Buno", id)
	}
}

func TestNumberToID10001MatchesReference(t *testing.T) {
	id, err := NumberToID(10001)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != "Patu Luta" {
		t.Fatalf("expected %q, got %q", "Patu Luta", id)
	}
}

func TestNumberToIDMaxNumMatchesReference(t *testing.T) {
	id, err := NumberToID(int(MaxNum))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != "Dobu Zusa" {
		t.Fatalf("expected %q, got %q", "Dobu Zusa", id)
	}
}

func TestNumberToIDOutOfRangeErrors(t *testing.T) {
	if _, err := NumberToID(int(MaxNum) + 1); err == nil {
		t.Fatal("expected error for out-of-range input")
	}
}

func TestNumberToIDNegativeErrors(t *testing.T) {
	if _, err := NumberToID(-1); err == nil {
		t.Fatal("expected error for negative input")
	}
}

func TestCapitalizationFirstLetterOnly(t *testing.T) {
	id, _ := NumberToID(42)
	for _, word := range strings.Split(id, " ") {
		expected := strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		if word != expected {
			t.Fatalf("expected %q, got %q", expected, word)
		}
	}
}

func TestDeterminism(t *testing.T) {
	id1, _ := NumberToID(987654)
	id2, _ := NumberToID(987654)
	if id1 != id2 {
		t.Fatalf("expected deterministic output, got %q and %q", id1, id2)
	}
}

func TestUniquenessAndBlacklistAvoidanceAcrossEntireValidRange(t *testing.T) {
	seen := make(map[string]struct{}, MaxNum+1)
	for n := 0; int64(n) <= MaxNum; n++ {
		id, err := NumberToID(n)
		if err != nil {
			t.Fatalf("unexpected error for %d: %v", n, err)
		}
		if _, ok := seen[id]; ok {
			t.Fatalf("%q (from %d) was already issued for a different number", id, n)
		}
		seen[id] = struct{}{}

		for _, w := range strings.Split(strings.ToLower(id), " ") {
			if _, blacklisted := Blacklist[w]; blacklisted {
				t.Fatalf("%q from %d is blacklisted", w, n)
			}
		}
	}
	if int64(len(seen)) != MaxNum+1 {
		t.Fatalf("expected %d unique ids, got %d", MaxNum+1, len(seen))
	}
}
