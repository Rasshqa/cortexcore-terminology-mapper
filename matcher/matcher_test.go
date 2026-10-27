package matcher

import (
	"testing"

	"terminology-mapper/dictionary"
)

func TestTokenizeAndFuzzyScore(t *testing.T) {
	tokens := Tokenize("Paracetamol 500 mg Kaplet (Oral)")
	if len(tokens) == 0 {
		t.Fatalf("expected tokens, got 0")
	}

	catalog := dictionary.GetDefaultCatalog()
	localItem := dictionary.LocalItem{
		LocalCode: "OBT-001",
		LocalName: "Paracetamol tab 500mg",
		Category:  dictionary.CategoryDrug,
	}

	matches := FindMatches(localItem, catalog, 3)
	if len(matches) == 0 {
		t.Fatalf("expected matches for Paracetamol, got 0")
	}

	if matches[0].StandardCode != "93001024" {
		t.Errorf("expected KFA code 93001024, got %s (%s)", matches[0].StandardCode, matches[0].StandardName)
	}

	if matches[0].Confidence < 50 {
		t.Errorf("expected high confidence for Paracetamol, got %d%%", matches[0].Confidence)
	}
}

func TestDiagnosisFuzzyMatch(t *testing.T) {
	catalog := dictionary.GetDefaultCatalog()
	localItem := dictionary.LocalItem{
		LocalCode: "DIAG-008",
		LocalName: "Darah Tinggi Primer",
		Category:  dictionary.CategoryDiagnosis,
	}

	matches := FindMatches(localItem, catalog, 3)
	if len(matches) == 0 {
		t.Fatalf("expected matches for Darah Tinggi, got 0")
	}

	if matches[0].StandardCode != "I10" {
		t.Errorf("expected ICD-10 I10 for Darah Tinggi, got %s", matches[0].StandardCode)
	}
}
