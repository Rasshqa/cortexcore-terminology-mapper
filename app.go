package main

import (
	"context"
	"encoding/csv"
	"os"
	"strings"

	"terminology-mapper/dictionary"
	"terminology-mapper/mapper"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx     context.Context
	catalog []dictionary.StandardTerm
	rows    []dictionary.MappedRow
}

func NewApp() *App {
	return &App{
		catalog: dictionary.GetDefaultCatalog(),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetCatalogSummary returns count of bundled official terms by category.
func (a *App) GetCatalogSummary() map[string]int {
	counts := map[string]int{"ICD10": 0, "ICD9CM": 0, "KFA": 0, "TOTAL": len(a.catalog)}
	for _, t := range a.catalog {
		counts[string(t.Category)]++
	}
	return counts
}

// IngestCSV parses hospital master data and runs auto-mapping.
func (a *App) IngestCSV(content string) map[string]interface{} {
	r := csv.NewReader(strings.NewReader(content))
	records, err := r.ReadAll()
	if err != nil || len(records) < 2 {
		return map[string]interface{}{"error": "CSV tidak valid atau kosong"}
	}

	headers := records[0]
	idxCode, idxName, idxCat := -1, -1, -1
	for i, h := range headers {
		hl := strings.ToLower(strings.TrimSpace(h))
		if strings.Contains(hl, "code") || strings.Contains(hl, "kode") || strings.Contains(hl, "id") {
			if idxCode == -1 { idxCode = i }
		}
		if strings.Contains(hl, "name") || strings.Contains(hl, "nama") || strings.Contains(hl, "item") || strings.Contains(hl, "deskripsi") {
			if idxName == -1 { idxName = i }
		}
		if strings.Contains(hl, "category") || strings.Contains(hl, "kategori") || strings.Contains(hl, "jenis") {
			if idxCat == -1 { idxCat = i }
		}
	}

	if idxCode == -1 || idxName == -1 {
		idxCode, idxName = 0, 1 // Fallback to col 0 and 1
	}

	var items []dictionary.LocalItem
	for _, row := range records[1:] {
		if len(row) <= idxCode || len(row) <= idxName {
			continue
		}
		cat := dictionary.CategoryDrug
		if idxCat >= 0 && idxCat < len(row) {
			cv := strings.ToUpper(strings.TrimSpace(row[idxCat]))
			if strings.Contains(cv, "DIAG") || strings.Contains(cv, "ICD10") {
				cat = dictionary.CategoryDiagnosis
			} else if strings.Contains(cv, "PROC") || strings.Contains(cv, "TIND") || strings.Contains(cv, "ICD9") {
				cat = dictionary.CategoryProcedure
			}
		}

		items = append(items, dictionary.LocalItem{
			LocalCode: strings.TrimSpace(row[idxCode]),
			LocalName: strings.TrimSpace(row[idxName]),
			Category:  cat,
		})
	}

	a.rows = mapper.ProcessItems(items, a.catalog)

	autoCount, reviewCount, unmappedCount := 0, 0, 0
	for _, r := range a.rows {
		if r.BestMatch != nil {
			if r.BestMatch.Confidence >= 75 {
				autoCount++
			} else if r.BestMatch.Confidence >= 40 {
				reviewCount++
			} else {
				unmappedCount++
			}
		} else {
			unmappedCount++
		}
	}

	return map[string]interface{}{
		"total":          len(a.rows),
		"auto_matched":   autoCount,
		"needs_review":   reviewCount,
		"unmapped":       unmappedCount,
		"rows":           a.rows,
	}
}

// OpenCSVFileDialog prompts user to choose a CSV file and loads it.
func (a *App) OpenCSVFileDialog() map[string]interface{} {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:   "Pilih Berkas CSV Master Data SIMRS",
		Filters: []runtime.FileFilter{{DisplayName: "CSV Files", Pattern: "*.csv"}},
	})
	if err != nil || path == "" {
		return map[string]interface{}{"error": "file tidak dipilih"}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return map[string]interface{}{"error": "gagal membaca: " + err.Error()}
	}

	return a.IngestCSV(string(data))
}

// ExportSQLFileDialog saves SQL INSERT/UPDATE script to disk.
func (a *App) ExportSQLFileDialog(tableName string) map[string]interface{} {
	if len(a.rows) == 0 {
		return map[string]interface{}{"error": "Belum ada data pemetaan"}
	}

	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "Ekspor Script SQL Pemetaan",
		DefaultFilename: "satusehat_mapping.sql",
		Filters:         []runtime.FileFilter{{DisplayName: "SQL Script", Pattern: "*.sql"}},
	})
	if err != nil || path == "" {
		return map[string]interface{}{"error": "ekspor dibatalkan"}
	}

	sql := mapper.GenerateSQLScript(a.rows, tableName)
	if err := os.WriteFile(path, []byte(sql), 0644); err != nil {
		return map[string]interface{}{"error": "gagal menulis file: " + err.Error()}
	}

	return map[string]interface{}{"saved": path}
}

// ExportCSVFileDialog saves final mapped CSV to disk.
func (a *App) ExportCSVFileDialog() map[string]interface{} {
	if len(a.rows) == 0 {
		return map[string]interface{}{"error": "Belum ada data pemetaan"}
	}

	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "Ekspor CSV Hasil Pemetaan",
		DefaultFilename: "satusehat_mapped.csv",
		Filters:         []runtime.FileFilter{{DisplayName: "CSV Files", Pattern: "*.csv"}},
	})
	if err != nil || path == "" {
		return map[string]interface{}{"error": "ekspor dibatalkan"}
	}

	csvStr := mapper.ExportCSV(a.rows)
	if err := os.WriteFile(path, []byte(csvStr), 0644); err != nil {
		return map[string]interface{}{"error": "gagal menulis file: " + err.Error()}
	}

	return map[string]interface{}{"saved": path}
}
