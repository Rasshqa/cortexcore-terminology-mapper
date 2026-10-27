package mapper

import (
	"strings"
	"testing"

	"terminology-mapper/dictionary"
)

func TestProcessAndExportScript(t *testing.T) {
	catalog := dictionary.GetDefaultCatalog()
	items := []dictionary.LocalItem{
		{LocalCode: "OBT-001", LocalName: "Paracetamol 500mg tablet", Category: dictionary.CategoryDrug},
		{LocalCode: "TND-002", LocalName: "Rekam Jantung Rutin", Category: dictionary.CategoryProcedure},
	}

	results := ProcessItems(items, catalog)
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	sql := GenerateSQLScript(results, "simrs_satusehat_mapping")
	if !strings.Contains(sql, "93001024") {
		t.Errorf("expected SQL to contain KFA code 93001024, got:\n%s", sql)
	}

	csvOut := ExportCSV(results)
	if !strings.Contains(csvOut, "93001024") {
		t.Errorf("expected CSV to contain KFA code 93001024, got:\n%s", csvOut)
	}
}
