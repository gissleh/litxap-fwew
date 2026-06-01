package litxapfwew

import (
	"slices"
	"testing"
)

func TestGlobal_SmokeTest(t *testing.T) {
	entries, err := Global().LookupEntries("tìfmetok")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) < 1 {
		t.Fatal("tìfmetok doesn't have entries")
	}

	entries, err = Global().LookupEntries("Änsìt")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) < 1 {
		t.Fatal("Änsìt doesn't have entries")
	}

	entries, err = Global().LookupEntries("pìwobe")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) < 1 {
		t.Fatal("pìwobe doesn't have entries")
	}

	entries, err = Global().LookupEntries("fpomtokxnga'")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) < 1 {
		t.Fatal("fpomtokxnga' doesn't have entries")
	}
	if !slices.Contains(entries[0].Syllables, "tokx") {
		t.Fatal("fpomtokxnga' has an error", entries[0].Syllables)
	}

	entries, err = Global().LookupEntries("tìmu'nitkan")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) < 1 {
		t.Fatal("tìmu'nitkan doesn't have entries")
	}
	if !slices.Equal(entries[0].Syllables, []string{"tì", "mu'", "nit", "kan"}) {
		t.Fatal("tìmu'nitkan has an error", entries[0].Syllables)
	}

	entries, err = Global().LookupEntries("syawm")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) < 1 {
		t.Fatal("syawm doesn't have entries")
	}
	if !slices.Equal(entries[0].Syllables, []string{"syawm"}) {
		t.Fatal("syawm has an error", entries[0].Syllables)
	}

	entries, err = Global().LookupEntries("oetsyìp")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) < 2 {
		t.Fatal("oetsyìp doesn't have both entries")
	}
	if !slices.Equal(entries[0].Syllables, []string{"oe", "tsyìp"}) {
		t.Fatal("oetsyìp has an error", entries[0].Syllables)
	}

	adpositions, err := Adpositions()
	if err != nil {
		t.Fatal(err)
	}
	if len(adpositions) == 0 {
		t.Fatal("Adpositions doesn't have entries")
	}

	mp := MultiWordPartDictionary()
	entries, err = mp.LookupEntries("tsaheyl")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) < 1 {
		t.Fatal("tsaheyl doesn't have entries in MultiWordPartDictionary")
	}
}
