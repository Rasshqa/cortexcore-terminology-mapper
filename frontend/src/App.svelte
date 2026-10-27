<script>
  import { onMount } from 'svelte'
  import { IngestCSV, OpenCSVFileDialog, ExportSQLFileDialog, ExportCSVFileDialog, GetCatalogSummary } from '../wailsjs/go/main/App.js'

  let stats = $state({ total: 0, auto_matched: 0, needs_review: 0, unmapped: 0 })
  let rows = $state([])
  let catalogCounts = $state({ ICD10: 0, ICD9CM: 0, KFA: 0, TOTAL: 0 })
  let activeFilter = $state('ALL')
  let statusMsg = $state('')

  const demoCSV = `local_code,local_name,category
OBT-PCT-01,Paracetamol 500mg tab,DRUG
OBT-AMX-02,Amoxicillin 500 mg kaplet,DRUG
OBT-AML-05,Amlodipine 5mg tablet,DRUG
OBT-OMZ-20,Omeprazole caps 20mg,DRUG
DIAG-HT-99,Darah Tinggi Esensial,DIAGNOSIS
DIAG-DBD-01,Demam Berdarah Dengue (DBD),DIAGNOSIS
DIAG-FLU-02,Batuk pilek flu ispa,DIAGNOSIS
TND-EKG-01,Pemeriksaan EKG Rekam Jantung,PROCEDURE
TND-THX-02,Foto Rontgen Dada Thorax,PROCEDURE
TND-JHT-03,Jahit luka hecting kulit,PROCEDURE`

  onMount(async () => {
    try {
      const summary = await GetCatalogSummary()
      if (summary) catalogCounts = summary
      loadDemoData()
    } catch (e) {
      console.error(e)
    }
  })

  function loadDemoData() {
    processCSVString(demoCSV)
    statusMsg = "Data demo SIMRS dimuat (10 item)"
  }

  async function handleLoadCSV() {
    try {
      const res = await OpenCSVFileDialog()
      if (res.error) {
        statusMsg = res.error
        return
      }
      applyResult(res)
      statusMsg = `Berhasil memuat berkas (${res.total} item)`
    } catch (e) {
      statusMsg = "Gagal membuka berkas: " + e
    }
  }

  async function processCSVString(content) {
    try {
      const res = await IngestCSV(content)
      if (res.error) {
        statusMsg = res.error
        return
      }
      applyResult(res)
    } catch (e) {
      statusMsg = "Gagal memproses CSV: " + e
    }
  }

  function applyResult(res) {
    stats = {
      total: res.total || 0,
      auto_matched: res.auto_matched || 0,
      needs_review: res.needs_review || 0,
      unmapped: res.unmapped || 0
    }
    rows = res.rows || []
  }

  async function handleExportSQL() {
    try {
      const res = await ExportSQLFileDialog("simrs_satusehat_mapping")
      if (res.error) {
        statusMsg = res.error
      } else {
        statusMsg = "Script SQL diekspor ke: " + res.saved
      }
    } catch (e) {
      statusMsg = "Gagal ekspor SQL: " + e
    }
  }

  async function handleExportCSV() {
    try {
      const res = await ExportCSVFileDialog()
      if (res.error) {
        statusMsg = res.error
      } else {
        statusMsg = "CSV diekspor ke: " + res.saved
      }
    } catch (e) {
      statusMsg = "Gagal ekspor CSV: " + e
    }
  }

  let filteredRows = $derived(rows.filter(r => {
    if (activeFilter === 'AUTO') return r.best_match && r.best_match.confidence >= 75
    if (activeFilter === 'REVIEW') return r.best_match && r.best_match.confidence >= 40 && r.best_match.confidence < 75
    if (activeFilter === 'UNMAPPED') return !r.best_match || r.best_match.confidence < 40
    return true
  }))
</script>

<div class="app-shell">
  <!-- Topbar -->
  <header class="topbar">
    <div class="topbar-brand">
      <div class="brand-dot"></div>
      <div>
        <span class="brand-name">CortexCore :: TerminologyMapper Studio</span>
        <span class="brand-sub"> · SATUSEHAT Kemenkes RI (PMK 24/2022)</span>
      </div>
    </div>
    <div style="display:flex; gap:8px;">
      <button class="btn btn-ghost" onclick={loadDemoData}>Demo SIMRS</button>
      <button class="btn btn-ghost" onclick={handleLoadCSV}>Buka CSV SIMRS</button>
      <button class="btn btn-sky" onclick={handleExportSQL}>Ekspor SQL</button>
      <button class="btn btn-emerald" onclick={handleExportCSV}>Ekspor CSV</button>
    </div>
  </header>

  <!-- Stat Strip -->
  <div class="stat-strip">
    <div class="stat-group">
      <div class="stat-pill">
        <span>Total Master:</span>
        <span class="num num-sky">{stats.total}</span>
      </div>
      <div class="stat-pill">
        <span>Auto-Matched (≥75%):</span>
        <span class="num num-green">{stats.auto_matched}</span>
      </div>
      <div class="stat-pill">
        <span>Perlu Review (40-74%):</span>
        <span class="num num-amber">{stats.needs_review}</span>
      </div>
      <div class="stat-pill">
        <span>Unmapped (&lt;40%):</span>
        <span class="num num-red">{stats.unmapped}</span>
      </div>
    </div>

    <!-- Filter Buttons -->
    <div style="display:flex; gap:6px;">
      <button class="btn btn-ghost" style={activeFilter==='ALL'?'border-color:var(--tm-sky)':''} onclick={() => activeFilter = 'ALL'}>Semua ({stats.total})</button>
      <button class="btn btn-ghost" style={activeFilter==='AUTO'?'border-color:var(--tm-emerald)':''} onclick={() => activeFilter = 'AUTO'}>Auto ({stats.auto_matched})</button>
      <button class="btn btn-ghost" style={activeFilter==='REVIEW'?'border-color:var(--tm-amber)':''} onclick={() => activeFilter = 'REVIEW'}>Review ({stats.needs_review})</button>
      <button class="btn btn-ghost" style={activeFilter==='UNMAPPED'?'border-color:var(--tm-crimson)':''} onclick={() => activeFilter = 'UNMAPPED'}>Unmapped ({stats.unmapped})</button>
    </div>
  </div>

  <!-- Workspace Matrix -->
  <main class="workspace">
    {#if statusMsg}
      <div style="font-size:11px; margin-bottom:8px; color:var(--tm-sky); font-family:monospace;">
        ❯ {statusMsg}
      </div>
    {/if}

    <table class="table-matrix">
      <thead>
        <tr>
          <th style="width:120px;">Kode SIMRS</th>
          <th>Nama Item Lokal SIMRS</th>
          <th style="width:90px;">Kategori</th>
          <th style="width:120px;">Standar Kemenkes</th>
          <th>Nama Resmi Standar</th>
          <th style="width:90px; text-align:center;">Confidence</th>
          <th style="width:100px; text-align:center;">Status</th>
        </tr>
      </thead>
      <tbody>
        {#each filteredRows as row}
          <tr>
            <td>
              <span class="code-badge">{row.local_code}</span>
            </td>
            <td style="font-weight:600;">{row.local_name}</td>
            <td>
              <span style="font-size:10px; color:var(--tm-muted); text-transform:uppercase;">
                {row.category}
              </span>
            </td>
            <td>
              {#if row.best_match}
                <span class="code-badge" style="color:var(--tm-emerald); border-color:var(--tm-emerald);">
                  {row.best_match.standard_code}
                </span>
              {:else}
                <span style="color:var(--tm-muted); font-size:11px;">-</span>
              {/if}
            </td>
            <td>
              {#if row.best_match}
                <div>{row.best_match.standard_name}</div>
                <div style="font-size:10px; color:var(--tm-muted); font-family:monospace;">
                  {row.best_match.system_uri}
                </div>
              {:else}
                <span style="color:var(--tm-crimson); font-size:11px;">Belum ada kecocokan</span>
              {/if}
            </td>
            <td style="text-align:center;">
              {#if row.best_match}
                <span class="conf-pill {row.best_match.confidence >= 75 ? 'conf-hi' : (row.best_match.confidence >= 40 ? 'conf-md' : 'conf-lo')}">
                  {row.best_match.confidence}%
                </span>
              {:else}
                <span class="conf-pill conf-lo">0%</span>
              {/if}
            </td>
            <td style="text-align:center;">
              {#if row.best_match && row.best_match.confidence >= 75}
                <span style="color:var(--tm-emerald); font-size:10px; font-weight:700;">TERPETAKAN</span>
              {:else if row.best_match && row.best_match.confidence >= 40}
                <span style="color:var(--tm-amber); font-size:10px; font-weight:700;">PERIKSA</span>
              {:else}
                <span style="color:var(--tm-crimson); font-size:10px; font-weight:700;">MANUAL</span>
              {/if}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </main>

  <!-- Footer -->
  <footer class="footer-attribution">
    <div>
      <span class="footer-brand">AETHERCORTEX</span>
      <span> · Architected by Rashqa Andrean Fitrah Sulaeman</span>
    </div>
    <div style="display:flex; gap:12px;">
      <span>Kamus Terpasang: {catalogCounts.TOTAL} entri resmi (KFA · ICD-10 · ICD-9-CM)</span>
      <span>100% Offline · Standar PMK 24/2022</span>
    </div>
  </footer>
</div>
