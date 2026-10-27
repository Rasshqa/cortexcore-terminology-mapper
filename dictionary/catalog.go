package dictionary

// GetDefaultCatalog returns standard reference terms for SATUSEHAT integration.
func GetDefaultCatalog() []StandardTerm {
	return []StandardTerm{
		// ───── ICD-10 (Diagnosis) ─────
		{
			Code: "I10", DisplayID: "Hipertensi Esensial (Primer)", DisplayEN: "Essential (primary) hypertension",
			SystemURI: "http://hl7.org/fhir/sid/icd-10", Category: CategoryDiagnosis,
			SearchTerms: []string{"hipertensi", "esensial", "darah", "tinggi", "tekanan", "hypertension", "i10"},
		},
		{
			Code: "E11.9", DisplayID: "Diabetes Melitus Tipe 2 Tanpa Komplikasi", DisplayEN: "Type 2 diabetes mellitus without complications",
			SystemURI: "http://hl7.org/fhir/sid/icd-10", Category: CategoryDiagnosis,
			SearchTerms: []string{"diabetes", "melitus", "dm", "gula", "tipe 2", "t2dm", "e11"},
		},
		{
			Code: "A91", DisplayID: "Demam Berdarah Dengue (DBD)", DisplayEN: "Dengue haemorrhagic fever",
			SystemURI: "http://hl7.org/fhir/sid/icd-10", Category: CategoryDiagnosis,
			SearchTerms: []string{"demam", "berdarah", "dengue", "dbd", "dhf", "a91"},
		},
		{
			Code: "J00", DisplayID: "Nasofaringitis Akut (Common Cold / Flu)", DisplayEN: "Acute nasopharyngitis [common cold]",
			SystemURI: "http://hl7.org/fhir/sid/icd-10", Category: CategoryDiagnosis,
			SearchTerms: []string{"ispa", "flu", "pilek", "batuk", "nasofaringitis", "common cold", "j00"},
		},
		{
			Code: "A09", DisplayID: "Gastroenteritis dan Kolitis Menular (Diare)", DisplayEN: "Infectious gastroenteritis and colitis, unspecified",
			SystemURI: "http://hl7.org/fhir/sid/icd-10", Category: CategoryDiagnosis,
			SearchTerms: []string{"diare", "gastroenteritis", "muntaber", "mencret", "ge", "a09"},
		},
		{
			Code: "A15.0", DisplayID: "Tuberkulosis Paru Terkonfirmasi Bakteriologis", DisplayEN: "Tuberculosis of lung, confirmed by sputum microscopy",
			SystemURI: "http://hl7.org/fhir/sid/icd-10", Category: CategoryDiagnosis,
			SearchTerms: []string{"tbc", "tb", "tuberkulosis", "paru", "flek", "a15"},
		},
		{
			Code: "K35.8", DisplayID: "Apendisitis Akut Lainnya", DisplayEN: "Other and unspecified acute appendicitis",
			SystemURI: "http://hl7.org/fhir/sid/icd-10", Category: CategoryDiagnosis,
			SearchTerms: []string{"usus", "buntu", "apendisitis", "akut", "appendicitis", "k35"},
		},
		{
			Code: "J45.9", DisplayID: "Asma, Tidak Spesifik", DisplayEN: "Asthma, unspecified",
			SystemURI: "http://hl7.org/fhir/sid/icd-10", Category: CategoryDiagnosis,
			SearchTerms: []string{"asma", "sesak", "napas", "wheezing", "asthma", "j45"},
		},

		// ───── ICD-9-CM (Prosedur / Tindakan) ─────
		{
			Code: "89.52", DisplayID: "Elektrokardiogram (EKG)", DisplayEN: "Electrocardiogram",
			SystemURI: "http://hl7.org/fhir/sid/icd-9-cm", Category: CategoryProcedure,
			SearchTerms: []string{"ekg", "ecg", "elektrokardiogram", "rekam", "jantung", "89.52"},
		},
		{
			Code: "87.44", DisplayID: "Rontgen Dada Rutin (Foto Thorax)", DisplayEN: "Routine chest x-ray, so described",
			SystemURI: "http://hl7.org/fhir/sid/icd-9-cm", Category: CategoryProcedure,
			SearchTerms: []string{"rontgen", "thorax", "toraks", "dada", "xray", "chest", "87.44"},
		},
		{
			Code: "47.09", DisplayID: "Apendektomi Terbuka Lainnya", DisplayEN: "Other appendectomy",
			SystemURI: "http://hl7.org/fhir/sid/icd-9-cm", Category: CategoryProcedure,
			SearchTerms: []string{"operasi", "usus", "buntu", "apendektomi", "appendectomy", "47.09"},
		},
		{
			Code: "86.59", DisplayID: "Penjahitan Kulit dan Jaringan Bawah Kulit (Hecting)", DisplayEN: "Closure of skin and subcutaneous tissue of other sites",
			SystemURI: "http://hl7.org/fhir/sid/icd-9-cm", Category: CategoryProcedure,
			SearchTerms: []string{"jahit", "luka", "hecting", "heacting", "jahitan", "86.59"},
		},
		{
			Code: "39.95", DisplayID: "Hemodialisis (Cuci Darah)", DisplayEN: "Hemodialysis",
			SystemURI: "http://hl7.org/fhir/sid/icd-9-cm", Category: CategoryProcedure,
			SearchTerms: []string{"hemodialisa", "hemodialisis", "cuci", "darah", "hd", "39.95"},
		},
		{
			Code: "90.59", DisplayID: "Pemeriksaan Mikroskopis Darah (Darah Lengkap)", DisplayEN: "Microscopic examination of blood",
			SystemURI: "http://hl7.org/fhir/sid/icd-9-cm", Category: CategoryProcedure,
			SearchTerms: []string{"lab", "darah", "lengkap", "hematologi", "dl", "90.59"},
		},

		// ───── KFA (Kamus Farmasi Kemenkes) ─────
		{
			Code: "93001024", DisplayID: "Paracetamol 500 mg Tablet", DisplayEN: "Paracetamol 500 mg Tablet",
			SystemURI: "http://sys-ids.kemkes.go.id/kfa", Category: CategoryDrug,
			SearchTerms: []string{"paracetamol", "parasetamol", "pct", "500mg", "tablet", "tab", "panadol", "sanmol"},
		},
		{
			Code: "93001025", DisplayID: "Paracetamol 120 mg/5 ml Sirup", DisplayEN: "Paracetamol 120 mg/5 ml Syrup",
			SystemURI: "http://sys-ids.kemkes.go.id/kfa", Category: CategoryDrug,
			SearchTerms: []string{"paracetamol", "parasetamol", "sirup", "syr", "120mg", "anak"},
		},
		{
			Code: "93002110", DisplayID: "Amoxicillin 500 mg Kapsul/Kaplet", DisplayEN: "Amoxicillin 500 mg Capsule",
			SystemURI: "http://sys-ids.kemkes.go.id/kfa", Category: CategoryDrug,
			SearchTerms: []string{"amoxicillin", "amoksisilin", "amox", "500mg", "kaplet", "kapsul", "antibiotik"},
		},
		{
			Code: "93003445", DisplayID: "Amlodipine 5 mg Tablet", DisplayEN: "Amlodipine 5 mg Tablet",
			SystemURI: "http://sys-ids.kemkes.go.id/kfa", Category: CategoryDrug,
			SearchTerms: []string{"amlodipine", "amlodipin", "5mg", "tablet", "tensi", "antihipertensi"},
		},
		{
			Code: "93003446", DisplayID: "Amlodipine 10 mg Tablet", DisplayEN: "Amlodipine 10 mg Tablet",
			SystemURI: "http://sys-ids.kemkes.go.id/kfa", Category: CategoryDrug,
			SearchTerms: []string{"amlodipine", "amlodipin", "10mg", "tablet", "tensi"},
		},
		{
			Code: "93004882", DisplayID: "Metformin HCl 500 mg Tablet", DisplayEN: "Metformin hydrochloride 500 mg Tablet",
			SystemURI: "http://sys-ids.kemkes.go.id/kfa", Category: CategoryDrug,
			SearchTerms: []string{"metformin", "500mg", "tablet", "gula", "diabetes", "antidiabetes"},
		},
		{
			Code: "93005912", DisplayID: "Omeprazole 20 mg Kapsul", DisplayEN: "Omeprazole 20 mg Capsule",
			SystemURI: "http://sys-ids.kemkes.go.id/kfa", Category: CategoryDrug,
			SearchTerms: []string{"omeprazole", "omeprazol", "20mg", "kapsul", "lambung", "maag", "ppi"},
		},
		{
			Code: "93006118", DisplayID: "Cefixime 100 mg Kapsul", DisplayEN: "Cefixime 100 mg Capsule",
			SystemURI: "http://sys-ids.kemkes.go.id/kfa", Category: CategoryDrug,
			SearchTerms: []string{"cefixime", "kefiksim", "100mg", "kapsul", "antibiotik"},
		},
	}
}
