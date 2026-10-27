# ADR 0001: Offline Fuzzy Lexical Matcher & Terminology Matrix Engine

## Status
Accepted

## Konteks
Rumah sakit di Indonesia rata-rata memiliki 5.000 hingga 25.000 item master data lokal (obat farmasi, diagnosis penyakit, dan tindakan medis). Melakukan pemetaan manual ke standar Kemenkes (KFA, ICD-10, ICD-9-CM) menggunakan lembar kerja spreadsheet memakan waktu berbulan-bulan dan menimbulkan kelelahan verifikasi staf faskes. Pemetaan berbasis cloud API eksternal rentan downtime intranet RS dan melanggar kerahasiaan master data faskes.

## Keputusan
1. **Penyematan Kamus Terminologi Resmi Luring**:
   - Memasukkan dataset inti terminologi ICD-10, ICD-9-CM, dan KFA langsung di dalam binary aplikasi Go.
   - 100% offline tanpa perlu koneksi API eksternal saat proses pemetaan berjalan.
2. **Algoritma Pencocokan Hibrida (Token Jaccard + Levenshtein)**:
   - Normalisasi teks: pembersihan singkatan medis umum Indonesia (misal: "tab", "syr", "inj", "susp", "inf").
   - Menghitung nilai *Confidence Score*:
     - **Tinggi (>= 80%)**: Hijau — auto-matched, siap ekspor langsung.
     - **Sedang (50% - 79%)**: Kuning — perlu tinjauan cepat 1-klik dari daftar pilihan teratas.
     - **Rendah (< 50%)**: Merah — butuh pemilihan manual.
3. **Ekspor Serbaguna**:
   - Menghasilkan berkas CSV hasil pemetaan dan berkas script SQL (`INSERT/UPDATE INTO simrs_mapping ...`) yang siap diimpor ke database SIMRS faskes.

## Konsekuensi
- **Positif:** Menghemat ratusan jam kerja tim IT RS, tidak ada biaya server, kepatuhan SATUSEHAT tercapai lebih cepat.
- **Negatif:** Kamus terminologi versi baru memerlukan pembaruan berkas kamus atau pembaruan rilis aplikasi.
