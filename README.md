# CortexCore :: TerminologyMapper Studio

> **Sistem Pemetaan Terminologi Standar Nasional SATUSEHAT Kemenkes RI (PMK No. 24/2022)**  
> Bagian dari inisiatif arsitektur **AetherCortex** (Klaster 01: Health & SATUSEHAT).  
> **Arsitek & Pengembang Utama:** Rashqa Andrean Fitrah Sulaeman

---

## 🏥 Latar Belakang Masalah
Rumah sakit dan fasilitas pelayanan kesehatan (faskes) di seluruh Indonesia menghadapi kendala regulasi besar: seluruh rekam medis lokal wajib dipetakan ke standar terminologi nasional sebelum dapat dikirimkan ke server SATUSEHAT:
1. **Diagnosis**: ICD-10 (WHO / Kemenkes)
2. **Prosedur / Tindakan**: ICD-9-CM
3. **Katalog Obat & BMHP**: KFA (Kamus Farmasi dan Alat Kesehatan Kemenkes RI)

Melakukan pemetaan puluhan ribu item master SIMRS secara manual lewat spreadsheet memakan waktu berbulan-bulan dan sarat salah ketik. **TerminologyMapper Studio** menyediakan solusi desktop berkecepatan tinggi, 100% luring (tanpa kebocoran data ke luar intranet RS), dengan mesin pencocokan token leksikal (Jaccard + Levenshtein) dan ekspor langsung ke script SQL/CSV.

---

## ✨ Fitur Unggulan
- **Mesin Fuzzy Match Luring (Offline)**: Memetakan nama lokal SIMRS (misal `Paracetamol 500mg tab`, `Darah Tinggi Primer`, `Rekam Jantung Rutin`) ke kode KFA/ICD secara otomatis dengan skor keyakinan (*confidence score*).
- **Kamus Resmi Bawaan**: Bundled dataset kode resmi KFA, ICD-10, dan ICD-9-CM yang siap digunakan tanpa koneksi internet.
- **Matrix Review Studio**: Visualisasi status pemetaan berkode warna:
  - 🟢 **Auto-Matched (≥75%)**: Terpetakan otomatis, siap pakai.
  - 🟡 **Perlu Review (40–74%)**: Rekomendasi teratas siap dikonfirmasi dalam 1 klik.
  - 🔴 **Unmapped (<40%)**: Butuh seleksi manual.
- **Ekspor Serbaguna**: Menghasilkan script SQL `INSERT/UPDATE ON CONFLICT` untuk database SIMRS lokal faskes (PostgreSQL / MySQL) atau berkas CSV terstruktur.

---

## 🛠️ Arsitektur & Teknologi
- **Backend**: Go 1.26 + Wails v2 (Webkit2GTK 4.1)
- **Frontend**: Svelte 5 (Runes) + Terminology Matrix UI (`#0B0F19` Matrix Charcoal & Sky Blue)
- **Keamanan & Kepatuhan**: 100% offline, zero cloud data exfiltration, sesuai UU Perlindungan Data Pribadi (UU No. 27/2022).

---

## 🚀 Menjalankan Aplikasi
```bash
# Clone repositori
git clone https://github.com/Rasshqa/cortexcore-terminology-mapper.git
cd cortexcore-terminology-mapper

# Menjalankan dalam mode pengembangan
wails dev

# Membangun binary produksi
wails build
```

---

## ⚖️ Lisensi
Dilisensikan di bawah lisensi **MIT**. Bebas digunakan oleh seluruh RS, vendor SIMRS, dan faskes di Indonesia.
