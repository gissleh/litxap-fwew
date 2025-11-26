package litxapfwew

import "testing"

func TestGlobal_SmokeTest(t *testing.T) {
	entries, err := Global().LookupEntries("tìfmetok")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) < 0 {
		t.Fatal("tìfmetok doesn't have entries")
	}

	entries, err = Global().LookupEntries("pìwobe")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) < 0 {
		t.Fatal("pìwobe doesn't have entries")
	}

	adpositions, err := Adpositions()
	if err != nil {
		t.Fatal(err)
	}
	if len(adpositions) == 0 {
		t.Fatal("Adpositions doesn't have entries")
	}
}
