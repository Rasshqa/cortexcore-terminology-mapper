package dictionary

// Category represents the domain of terminology.
type Category string

const (
	CategoryDiagnosis Category = "ICD10"
	CategoryProcedure Category = "ICD9CM"
	CategoryDrug      Category = "KFA"
)

// StandardTerm is a verified official reference code from WHO or Kemenkes RI.
type StandardTerm struct {
	Code        string   `json:"code"`         // e.g. "I10", "89.52", "93001024"
	DisplayID   string   `json:"display_id"`   // Bahasa Indonesia standard label
	DisplayEN   string   `json:"display_en"`   // English standard label
	SystemURI   string   `json:"system_uri"`   // Official FHIR system URL
	Category    Category `json:"category"`     // ICD10 | ICD9CM | KFA
	SearchTerms []string `json:"search_terms"` // Cleaned tokens for indexing
}

// LocalItem is an unmapped entry exported from hospital SIMRS master data.
type LocalItem struct {
	LocalCode string   `json:"local_code"` // e.g. "OBT-0012", "DIAG-HT"
	LocalName string   `json:"local_name"` // e.g. "Paracetamol Syr 60ml", "Darah Tinggi"
	Category  Category `json:"category"`   // ICD10 | ICD9CM | KFA
}

// MappingMatch holds the suggested standard term and confidence score.
type MappingMatch struct {
	StandardCode string   `json:"standard_code"`
	StandardName string   `json:"standard_name"`
	SystemURI    string   `json:"system_uri"`
	Category     Category `json:"category"`
	Confidence   int      `json:"confidence"` // 0 - 100%
	Status       string   `json:"status"`     // "AUTO_MATCHED" | "NEEDS_REVIEW" | "UNMAPPED"
}

// MappedRow is the final unified output row.
type MappedRow struct {
	LocalCode    string        `json:"local_code"`
	LocalName    string        `json:"local_name"`
	Category     Category      `json:"category"`
	BestMatch    *MappingMatch `json:"best_match,omitempty"`
	Candidates   []MappingMatch`json:"candidates,omitempty"`
	IsApproved   bool          `json:"is_approved"`
	ManualCode   string        `json:"manual_code,omitempty"`
	ManualName   string        `json:"manual_name,omitempty"`
}
