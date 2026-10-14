package preview

import "html/template"

var directoryTemplate = template.Must(template.New("directory").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}} · {{.Host}}</title>
  <style>
    :root {
      color-scheme: light dark;
      font: 15px/1.5 system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
      --page: Canvas;
      --surface: Canvas;
      --surface-subtle: color-mix(in srgb, CanvasText 4%, Canvas);
      --text: CanvasText;
      --muted: GrayText;
      --line: #8885;
      --accent: LinkText;
      --accent-hover: LinkText;
      --accent-soft: color-mix(in srgb, CanvasText 7%, Canvas);
      --focus: Highlight;
      --row-hover: color-mix(in srgb, CanvasText 4%, Canvas);
      --row-selected: color-mix(in srgb, Highlight 12%, Canvas);
      --badge-bg: color-mix(in srgb, CanvasText 7%, Canvas);
      --badge-text: GrayText;
      --shadow: 0 8px 28px rgb(0 0 0 / 10%);
    }
    * { box-sizing: border-box; }
    body { margin: 0; background: var(--page); color: var(--text); }
    a { color: inherit; }
    button, input, select { font: inherit; }
    button:focus-visible, a:focus-visible, input:focus-visible, select:focus-visible, label.upload-picker:focus-within { outline: 3px solid var(--focus); outline-offset: 2px; }
    [hidden] { display: none !important; }
    .page-shell { width: min(100%, 1120px); margin: 0 auto; padding: 30px 24px 56px; }
    .site-header { margin-bottom: 20px; }
    .brand { display: inline-flex; align-items: center; gap: .55rem; color: var(--muted); font-size: .82rem; font-weight: 720; letter-spacing: .07em; text-decoration: none; text-transform: lowercase; }
    .brand-mark { width: .55rem; height: .55rem; border-radius: 50%; background: var(--muted); }
    .breadcrumbs { margin-bottom: 22px; color: var(--muted); font-size: .88rem; overflow-wrap: anywhere; }
    .breadcrumbs a { color: LinkText; text-decoration: none; }
    .breadcrumbs a:visited { color: VisitedText; }
    .breadcrumbs a:hover { color: var(--accent); }
    .breadcrumbs span { padding: 0 .38rem; color: var(--line); }
    .directory-heading { display: flex; align-items: flex-end; justify-content: space-between; gap: 20px; margin-bottom: 20px; }
    .directory-heading h1 { margin: 0; font-size: clamp(1.6rem, 3vw, 2rem); font-weight: 680; letter-spacing: -.035em; line-height: 1.2; }
    .directory-path { margin: .42rem 0 0; color: var(--muted); font-size: .84rem; overflow-wrap: anywhere; }
    .directory-path code { font: inherit; }
    .directory-actions { display: flex; flex: 0 0 auto; align-items: center; flex-wrap: wrap; gap: 8px; }
    .button { min-height: 40px; display: inline-flex; align-items: center; justify-content: center; gap: .5rem; padding: 0 12px; border: 1px solid var(--line); border-radius: 8px; background: var(--surface); color: var(--text); cursor: pointer; font-size: .88rem; font-weight: 580; line-height: 1; text-decoration: none; white-space: nowrap; transition: background-color .15s ease, border-color .15s ease, color .15s ease; }
    .button:hover:not(:disabled) { background: var(--surface-subtle); border-color: var(--muted); }
    .button-primary { border-color: #8888; background: ButtonFace; color: ButtonText; }
    .button-primary:hover:not(:disabled) { border-color: var(--muted); background: ButtonFace; color: ButtonText; }
    .button-ghost { border-color: transparent; background: transparent; color: var(--muted); }
    .button-ghost:hover:not(:disabled) { background: var(--accent-soft); border-color: transparent; color: var(--text); }
    .button:disabled { cursor: not-allowed; opacity: .48; }
    .button svg { width: 16px; height: 16px; flex: 0 0 auto; }
    .listing-surface { overflow: visible; border: 1px solid var(--line); border-radius: 14px; background: var(--surface); box-shadow: var(--shadow); }
    .listing-toolbar { display: flex; align-items: center; flex-wrap: wrap; gap: 12px; padding: 16px 18px 14px; border-bottom: 1px solid var(--line); }
    .directory-filters { display: grid; flex: 1 1 600px; grid-template-columns: minmax(160px, 1fr) minmax(180px, 250px) auto; gap: 9px; min-width: 0; }
    .search-control { position: relative; display: flex; min-width: 0; height: 40px; align-items: center; border: 1px solid var(--line); border-radius: 8px; background: var(--surface); color: var(--muted); transition: border-color .15s ease, box-shadow .15s ease; }
    .search-control:focus-within { border-color: var(--accent); box-shadow: 0 0 0 3px color-mix(in srgb, var(--focus), transparent 82%); }
    .search-control > svg { width: 17px; height: 17px; margin-left: 12px; flex: 0 0 auto; }
    .search-control input { width: 100%; min-width: 0; height: 100%; padding: 0 11px; border: 0; outline: 0; background: transparent; color: var(--text); }
    .search-control input::placeholder { color: var(--muted); opacity: .88; }
    .type-filter { position: relative; min-width: 0; }
    .type-filter summary { display: flex; height: 40px; align-items: center; justify-content: space-between; gap: 8px; padding: 0 10px; border: 1px solid var(--line); border-radius: 8px; background: var(--surface); color: var(--muted); cursor: pointer; list-style: none; }
    .type-filter summary::-webkit-details-marker { display: none; }
    .type-filter summary:focus-visible { outline: 3px solid var(--focus); outline-offset: 2px; }
    .type-filter[open] summary { border-color: var(--accent); box-shadow: 0 0 0 3px color-mix(in srgb, var(--focus), transparent 82%); }
    .type-filter summary > span:first-child { flex: 0 0 auto; font-size: .84rem; }
    .type-filter summary .selected-types { min-width: 0; overflow: hidden; color: var(--text); font-size: .8rem; text-align: right; text-overflow: ellipsis; white-space: nowrap; }
    .type-filter summary > svg { width: 15px; height: 15px; flex: 0 0 auto; }
    .type-options { position: absolute; top: calc(100% + 8px); right: 0; z-index: 5; width: min(330px, calc(100vw - 48px)); max-height: min(360px, 65vh); overflow: auto; padding: 12px; border: 1px solid var(--line); border-radius: 10px; background: var(--surface); box-shadow: var(--shadow); }
    .type-options p { margin: 0 0 8px; color: var(--muted); font-size: .76rem; }
    .type-options-grid { display: grid; grid-template-columns: repeat(2, minmax(0, 1fr)); gap: 3px; }
    .type-option { min-height: 34px; display: flex; align-items: center; gap: 8px; padding: 0 7px; border-radius: 6px; color: var(--text); cursor: pointer; font-size: .82rem; }
    .type-option:hover { background: var(--surface-subtle); }
    .type-option input { width: 15px; height: 15px; margin: 0; }
    #clear-filters { min-width: 68px; padding: 0 9px; }
    #entry-count { margin: 0 0 0 auto; color: var(--muted); font-size: .82rem; white-space: nowrap; }
    .entry-table { width: 100%; border-collapse: collapse; table-layout: fixed; }
    .entry-table th, .entry-table td { padding: 10px 12px; border-bottom: 1px solid var(--line); text-align: left; vertical-align: middle; }
    .entry-table th { height: 42px; background: var(--surface-subtle); color: var(--muted); font-size: .74rem; font-weight: 650; letter-spacing: .045em; text-transform: uppercase; }
    .entry-table th:first-child, .entry-table td:first-child { width: 48px; padding-left: 18px; padding-right: 8px; }
    .entry-table th:nth-child(3), .entry-table td:nth-child(3) { width: 142px; }
    .entry-table th:last-child, .entry-table td:last-child { width: 94px; text-align: right; }
    .entry-table tbody tr:last-child td { border-bottom: 0; }
    .entry-row { transition: background-color .12s ease; }
    .entry-row:hover { background: var(--row-hover); }
    .entry-row.is-selected { background: var(--row-selected); }
    .entry-row[hidden] { display: none !important; }
    .entry-select, #select-all { width: 16px; height: 16px; margin: 0; vertical-align: middle; cursor: pointer; }
    .name-cell { min-width: 0; }
    .entry-name { display: flex; align-items: center; gap: 10px; min-width: 0; }
    .entry-icon { width: 19px; height: 19px; flex: 0 0 19px; color: var(--muted); }
    .entry-icon-folder { color: var(--muted); }
    .entry-link { min-width: 0; overflow-wrap: anywhere; color: LinkText; font-weight: 570; text-decoration: none; }
    .entry-link:visited, .download-link:visited { color: VisitedText; }
    .entry-link:hover { color: var(--accent); text-decoration: underline; text-underline-offset: 3px; }
    .kind-badge { display: inline-flex; max-width: 100%; align-items: center; padding: 3px 8px; border-radius: 999px; background: var(--badge-bg); color: var(--badge-text); font-size: .73rem; font-weight: 600; line-height: 1.4; white-space: nowrap; }
    .download-link { color: var(--accent); font-size: .82rem; font-weight: 600; text-decoration: none; }
    .download-link:hover { color: var(--accent-hover); text-decoration: underline; text-underline-offset: 3px; }
    #empty-directory td, #no-results td { padding: 30px 18px; color: var(--muted); text-align: center; }
    #no-results[hidden] { display: none !important; }
    .upload-panel { margin: 0 0 18px; padding: 16px; border: 1px solid var(--line); border-radius: 12px; background: var(--surface); }
    .upload-header, .upload-actions, .upload-progress-row { display: flex; align-items: center; flex-wrap: wrap; gap: 9px; }
    .upload-header { justify-content: space-between; margin-bottom: 12px; }
    .upload-header h2 { margin: 0; font-size: .96rem; font-weight: 650; }
    .upload-actions { justify-content: flex-end; }
    .upload-picker { position: relative; }
    .upload-picker input { position: absolute; width: 1px; height: 1px; padding: 0; margin: -1px; overflow: hidden; clip: rect(0, 0, 0, 0); white-space: nowrap; border: 0; }
    .drop-zone { min-height: 76px; display: flex; align-items: center; justify-content: center; padding: 16px; border: 1px dashed var(--line); border-radius: 9px; background: var(--surface-subtle); color: var(--muted); text-align: center; transition: background-color .15s ease, border-color .15s ease; }
    .drop-zone.dragover { border-color: var(--accent); background: var(--accent-soft); color: var(--text); }
    .upload-progress-row { margin: 12px 0 0; }
    progress { width: min(100%, 24rem); }
    .sr-only { position: absolute; width: 1px; height: 1px; padding: 0; margin: -1px; overflow: hidden; clip: rect(0, 0, 0, 0); white-space: nowrap; border: 0; }
    @media (max-width: 680px) {
      .page-shell { padding: 22px 16px 40px; }
      .directory-heading { align-items: flex-start; flex-direction: column; gap: 14px; }
      .directory-actions { width: 100%; }
      .directory-filters { flex-basis: 100%; }
      #entry-count { margin-left: auto; }
    }
    @media (max-width: 480px) {
      .page-shell { padding: 18px 12px 32px; }
      .listing-toolbar { gap: 8px; padding: 12px; }
      .directory-filters { grid-template-columns: minmax(0, 1fr) auto; gap: 8px; }
      .search-control { grid-column: 1 / -1; }
      .type-filter { min-width: 0; }
      .type-options { right: auto; left: 0; }
      #clear-filters { min-width: 66px; }
      #entry-count { width: 100%; margin: 0; text-align: right; }
      .entry-table th, .entry-table td { padding: 9px 7px; }
      .entry-table th:first-child, .entry-table td:first-child { width: 38px; padding-left: 10px; padding-right: 4px; }
      .entry-table th:nth-child(3), .entry-table td:nth-child(3) { width: 94px; }
      .entry-table th:last-child, .entry-table td:last-child { width: 72px; }
      .entry-name { gap: 7px; }
      .entry-icon { width: 17px; height: 17px; flex-basis: 17px; }
      .kind-badge { padding: 3px 6px; font-size: .68rem; }
      .download-link { font-size: .76rem; }
      .upload-header { align-items: flex-start; flex-direction: column; }
      .upload-actions { justify-content: flex-start; }
    }
  </style>
</head>
<body>
  <main class="page-shell">
    <header class="site-header"><a class="brand" href="/"><span class="brand-mark" aria-hidden="true"></span>ykview</a></header>
    <nav class="breadcrumbs" aria-label="Breadcrumb">{{.Breadcrumb}}</nav>
    <section class="directory-heading" aria-labelledby="directory-title">
      <div>
        <h1 id="directory-title">{{.Title}}</h1>
        <p class="directory-path"><code>{{.RemotePath}}</code></p>
      </div>
      <div class="directory-actions">
        {{if .Parent}}<a class="button" href="{{.Parent}}"><svg viewBox="0 0 20 20" fill="none" aria-hidden="true"><path d="M10 15.5V4.75M5.75 9l4.25-4.25L14.25 9" stroke="currentColor" stroke-width="1.7" stroke-linecap="round" stroke-linejoin="round"/></svg>Parent</a>{{end}}
        <button id="download-selected" class="button button-primary" type="button" disabled><svg viewBox="0 0 20 20" fill="none" aria-hidden="true"><path d="M10 3.5v8m0 0 3-3m-3 3-3-3M4.5 13.25v2.25c0 .55.45 1 1 1h9c.55 0 1-.45 1-1v-2.25" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round"/></svg>Download selected (.zip)</button>
      </div>
    </section>
    {{if .WriteEnabled}}
    <section class="upload-panel" aria-label="Upload">
      <div class="upload-header">
        <h2>Upload</h2>
        <div class="upload-actions">
          <label class="button upload-picker" for="file-picker"><input id="file-picker" type="file" multiple><svg viewBox="0 0 20 20" fill="none" aria-hidden="true"><path d="M10 13V4m0 0L6.75 7.25M10 4l3.25 3.25M4.5 12.5v2.75c0 .55.45 1 1 1h9c.55 0 1-.45 1-1V12.5" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round"/></svg>Choose files</label>
          <label class="button upload-picker" for="directory-picker"><input id="directory-picker" type="file" webkitdirectory directory multiple><svg viewBox="0 0 20 20" fill="none" aria-hidden="true"><path d="M2.75 5.75c0-.83.67-1.5 1.5-1.5h4l1.5 1.7h6c.83 0 1.5.67 1.5 1.5v7.3c0 .83-.67 1.5-1.5 1.5H4.25c-.83 0-1.5-.67-1.5-1.5v-9Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="round"/></svg>Choose directory</label>
        </div>
      </div>
      <div id="drop-zone" class="drop-zone">Drop files or a directory here</div>
      <div class="upload-progress-row"><progress id="upload-progress" max="100" value="0" hidden></progress><button id="upload-pause" class="button" type="button" hidden>Pause</button><span id="upload-status" role="status"></span></div>
    </section>
    {{end}}
    <section class="listing-surface" aria-label="Directory contents">
      <div class="listing-toolbar">
        <div class="directory-filters" role="group" aria-label="Directory filters">
          <label class="search-control" for="name-filter">
            <svg viewBox="0 0 20 20" fill="none" aria-hidden="true"><circle cx="8.8" cy="8.8" r="5.6" stroke="currentColor" stroke-width="1.6"/><path d="m13 13 3.75 3.75" stroke="currentColor" stroke-width="1.6" stroke-linecap="round"/></svg>
            <span class="sr-only">Name contains</span>
            <input id="name-filter" type="search" autocomplete="off" placeholder="Search names…">
          </label>
        <details id="type-filter" class="type-filter">
          <summary aria-controls="type-filter-options"><span>File type</span><span id="type-filter-value" class="selected-types">All types</span><svg viewBox="0 0 20 20" fill="none" aria-hidden="true"><path d="m5.5 7.5 4.5 4.5 4.5-4.5" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round"/></svg></summary>
          <div id="type-filter-options" class="type-options" role="group" aria-label="File types. Entries match any selected type.">
            <p>Show entries matching any selected type.</p>
            <div class="type-options-grid">
              <label class="type-option"><input class="type-filter-option" type="checkbox" value="directory" data-label="Directory"><span>Directory</span></label>
              <label class="type-option"><input class="type-filter-option" type="checkbox" value="markdown" data-label="Markdown"><span>Markdown</span></label>
              <label class="type-option"><input class="type-filter-option" type="checkbox" value="image" data-label="Image"><span>Image</span></label>
              <label class="type-option"><input class="type-filter-option" type="checkbox" value="html" data-label="HTML"><span>HTML</span></label>
              <label class="type-option"><input class="type-filter-option" type="checkbox" value="csv" data-label="CSV"><span>CSV</span></label>
              <label class="type-option"><input class="type-filter-option" type="checkbox" value="tsv" data-label="TSV"><span>TSV</span></label>
              <label class="type-option"><input class="type-filter-option" type="checkbox" value="json-lines" data-label="JSON Lines"><span>JSON Lines</span></label>
              <label class="type-option"><input class="type-filter-option" type="checkbox" value="audio" data-label="Audio"><span>Audio</span></label>
              <label class="type-option"><input class="type-filter-option" type="checkbox" value="video" data-label="Video"><span>Video</span></label>
              <label class="type-option"><input class="type-filter-option" type="checkbox" value="text" data-label="Text"><span>Text</span></label>
              <label class="type-option"><input class="type-filter-option" type="checkbox" value="other-file" data-label="Other file"><span>Other file</span></label>
            </div>
          </div>
        </details>
          <button id="clear-filters" class="button button-ghost" type="button" aria-label="Clear filters"><svg viewBox="0 0 20 20" fill="none" aria-hidden="true"><path d="m5.75 5.75 8.5 8.5m0-8.5-8.5 8.5" stroke="currentColor" stroke-width="1.6" stroke-linecap="round"/></svg><span>Clear</span></button>
        </div>
        <p id="entry-count" role="status" aria-live="polite"></p>
      </div>
      <table class="entry-table">
        <thead><tr><th><input id="select-all" type="checkbox" aria-label="Select all"></th><th>Name</th><th>Kind</th><th>Action</th></tr></thead>
        <tbody>
        {{range .Entries}}
          <tr class="entry-row" data-filter-type="{{.FilterKind}}">
            <td><input class="entry-select" type="checkbox" data-path="{{.SelectPath}}" aria-label="Select {{.Name}}"></td>
            <td class="name-cell"><div class="entry-name">
              {{if eq .FilterKind "directory"}}<svg class="entry-icon entry-icon-folder" viewBox="0 0 20 20" fill="none" aria-hidden="true" focusable="false"><path d="M2.75 5.5c0-.83.67-1.5 1.5-1.5h4l1.5 1.7h6c.83 0 1.5.67 1.5 1.5v7.3c0 .83-.67 1.5-1.5 1.5H4.25c-.83 0-1.5-.67-1.5-1.5v-9Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="round"/><path d="M2.9 8h14.2" stroke="currentColor" stroke-width="1.5"/></svg>{{else}}<svg class="entry-icon" viewBox="0 0 20 20" fill="none" aria-hidden="true" focusable="false"><path d="M5 2.75h6.2l3.8 3.8v10.7H5c-.83 0-1.5-.67-1.5-1.5v-11.5c0-.83.67-1.5 1.5-1.5Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="round"/><path d="M11 3v4h4M6.5 10h5.75M6.5 13h5.75" stroke="currentColor" stroke-width="1.4" stroke-linecap="round"/></svg>{{end}}
              <a class="entry-link" href="{{.Href}}">{{.Name}}</a>
            </div></td>
            <td><span class="kind-badge">{{.Kind}}</span></td>
            <td><a class="download-link" href="{{.DownloadURL}}">Download</a></td>
          </tr>
        {{else}}
          <tr id="empty-directory"><td colspan="4">Empty directory</td></tr>
        {{end}}
          <tr id="no-results" hidden><td colspan="4">No matching entries</td></tr>
        </tbody>
      </table>
    </section>
  </main>
  <div id="transfer-config" data-upload-url="{{.UploadURL}}" data-chunk-upload-url="{{.ChunkUploadURL}}"></div>
  <script>
    (() => {
      const rows = Array.from(document.querySelectorAll('.entry-row'));
      const nameFilter = document.getElementById('name-filter');
      const typeFilter = document.getElementById('type-filter');
      const typeFilterValue = document.getElementById('type-filter-value');
      const typeOptions = Array.from(document.querySelectorAll('.type-filter-option'));
      const clearFilters = document.getElementById('clear-filters');
      const entryCount = document.getElementById('entry-count');
      const noResults = document.getElementById('no-results');
      const selectAll = document.getElementById('select-all');
      const visibleRows = () => rows.filter((row) => !row.hidden);
      const visibleSelections = () => visibleRows().flatMap((row) => {
        const checkbox = row.querySelector('.entry-select');
        return checkbox?.checked ? [checkbox] : [];
      });
      const selectedTypes = () => typeOptions.filter((checkbox) => checkbox.checked);
      const updateTypeSummary = () => {
        if (!typeFilterValue) return;
        const labels = selectedTypes().map((checkbox) => checkbox.dataset.label || checkbox.value);
        if (labels.length === 0) typeFilterValue.textContent = 'All types';
        else if (labels.length === 1) typeFilterValue.textContent = labels[0];
        else if (labels.length === 2) typeFilterValue.textContent = labels[0] + ' or ' + labels[1];
        else typeFilterValue.textContent = labels.length + ' types selected';
        typeFilterValue.title = labels.length ? labels.join(' or ') : 'All types';
      };
      const updateSelectionState = () => {
        const visible = visibleRows();
        const checkedCount = visibleSelections().length;
        rows.forEach((row) => {
          const checkbox = row.querySelector('.entry-select');
          row.classList.toggle('is-selected', Boolean(checkbox?.checked));
        });
        if (selectAll) {
          selectAll.checked = visible.length > 0 && checkedCount === visible.length;
          selectAll.indeterminate = checkedCount > 0 && checkedCount < visible.length;
          selectAll.disabled = visible.length === 0;
        }
        const downloadSelected = document.getElementById('download-selected');
        if (downloadSelected) downloadSelected.disabled = checkedCount === 0;
      };
      const applyFilters = () => {
        const name = (nameFilter?.value || '').toLowerCase();
        const types = selectedTypes().map((checkbox) => checkbox.value);
        let visibleCount = 0;
        rows.forEach((row) => {
          const entryName = row.querySelector('.entry-link')?.textContent || '';
          const matches = entryName.toLowerCase().includes(name) &&
            (types.length === 0 || types.includes(row.dataset.filterType));
          row.hidden = !matches;
          if (matches) visibleCount++;
          else {
            const checkbox = row.querySelector('.entry-select');
            if (checkbox) checkbox.checked = false;
          }
        });
        if (entryCount) entryCount.textContent = 'Showing ' + visibleCount + ' of ' + rows.length + ' entries';
        if (noResults) noResults.hidden = rows.length === 0 || visibleCount !== 0;
        updateSelectionState();
      };
      nameFilter?.addEventListener('input', applyFilters);
      typeOptions.forEach((checkbox) => checkbox.addEventListener('change', () => {
        updateTypeSummary();
        applyFilters();
      }));
      typeFilter?.addEventListener('keydown', (event) => {
        if (event.key === 'Escape' && typeFilter.open) {
          typeFilter.open = false;
          typeFilter.querySelector('summary')?.focus();
        }
      });
      document.addEventListener('click', (event) => {
        if (typeFilter?.open && !typeFilter.contains(event.target)) typeFilter.open = false;
      });
      clearFilters?.addEventListener('click', () => {
        if (nameFilter) nameFilter.value = '';
        typeOptions.forEach((checkbox) => { checkbox.checked = false; });
        updateTypeSummary();
        applyFilters();
        if (typeFilter) typeFilter.open = false;
        nameFilter?.focus();
      });
      document.querySelectorAll('.entry-select').forEach((checkbox) => {
        checkbox.addEventListener('change', updateSelectionState);
      });
      updateTypeSummary();
      applyFilters();
      selectAll?.addEventListener('change', () => {
        visibleRows().forEach((row) => {
          const checkbox = row.querySelector('.entry-select');
          if (checkbox) checkbox.checked = selectAll.checked;
        });
        updateSelectionState();
      });
      document.getElementById('download-selected')?.addEventListener('click', () => {
        const paths = visibleSelections().map((checkbox) => checkbox.dataset.path).filter(Boolean);
        if (!paths.length) {
          window.alert('Select at least one file or directory.');
          return;
        }
        const query = new URLSearchParams();
        paths.forEach((entryPath) => query.append('path', entryPath));
        window.location.href = '/_ykview/transfer/download-zip?' + query.toString();
      });

      const config = document.getElementById('transfer-config');
      const filePicker = document.getElementById('file-picker');
      const directoryPicker = document.getElementById('directory-picker');
      const dropZone = document.getElementById('drop-zone');
      const progress = document.getElementById('upload-progress');
      const pauseButton = document.getElementById('upload-pause');
      const status = document.getElementById('upload-status');
      if (!config || !filePicker || !directoryPicker || !dropZone) return;

      const pickerRecords = (input) => Array.from(input.files || []).map((file) => ({
        file,
        path: file.webkitRelativePath || file.name,
      }));
      const readEntries = (reader) => new Promise((resolve, reject) => {
        const all = [];
        const read = () => reader.readEntries((entries) => {
          if (!entries.length) { resolve(all); return; }
          all.push(...entries);
          read();
        }, reject);
        read();
      });
      const walkEntry = async (entry, prefix, output) => {
        if (entry.isFile) {
          await new Promise((resolve, reject) => entry.file((file) => {
            output.push({ file, path: prefix + file.name });
            resolve();
          }, reject));
          return;
        }
        if (entry.isDirectory) {
          const children = await readEntries(entry.createReader());
          for (const child of children) await walkEntry(child, prefix + entry.name + '/', output);
        }
      };
      const dropRecords = async (dataTransfer) => {
        const items = Array.from(dataTransfer.items || []);
        const entries = items.map((item) => item.webkitGetAsEntry && item.webkitGetAsEntry()).filter(Boolean);
        if (!entries.length) return Array.from(dataTransfer.files || []).map((file) => ({ file, path: file.name }));
        const output = [];
        for (const entry of entries) await walkEntry(entry, '', output);
        return output;
      };
      let paused = false;
      let uploading = false;
      let resumeWaiters = [];
      const waitForResume = () => paused ? new Promise((resolve) => resumeWaiters.push(resolve)) : Promise.resolve();
      const updateProgress = (uploaded, total) => {
        progress.value = total ? Math.min(100, Math.round(uploaded / total * 100)) : 100;
      };
      const requestChunk = (record, uploadID, offset, chunk) => new Promise((resolve, reject) => {
        const url = new URL(config.dataset.chunkUploadUrl, window.location.href);
        url.searchParams.set('path', record.path);
        url.searchParams.set('total', String(record.file.size));
        url.searchParams.set('offset', String(offset));
        if (uploadID) url.searchParams.set('upload_id', uploadID);
        const xhr = new XMLHttpRequest();
        xhr.open('PATCH', url.toString());
        xhr.onload = () => {
          let response = {};
          try { response = JSON.parse(xhr.responseText); } catch (_) {}
          if (xhr.status >= 200 && xhr.status < 300) resolve(response);
          else reject(new Error(response.error || ('Upload failed (' + xhr.status + ').')));
        };
        xhr.onerror = () => reject(new Error('Upload connection failed.'));
        xhr.send(chunk);
      });
      const uploadOne = async (record, progressBytes, totalBytes) => {
        let uploadID = '';
        let offset = 0;
        do {
          await waitForResume();
          const end = Math.min(offset + 8 * 1024 * 1024, record.file.size);
          const response = await requestChunk(record, uploadID, offset, record.file.slice(offset, end));
          uploadID = response.upload_id || uploadID;
          offset = response.offset;
          progressBytes[record.index] = offset;
          updateProgress(progressBytes.reduce((sum, value) => sum + value, 0), totalBytes);
          if (response.complete) return;
        } while (offset < record.file.size || record.file.size === 0);
      };
      const upload = async (records) => {
        if (!records.length) return;
        if (uploading) return;
        if (!window.confirm('Upload ' + records.length + ' file(s) to this directory? Existing regular files will be replaced.')) {
          status.textContent = 'Upload cancelled.';
          return;
        }
        uploading = true;
        paused = false;
        records = records.map((record, index) => ({ ...record, index }));
        const total = records.reduce((sum, record) => sum + record.file.size, 0);
        const progressBytes = records.map(() => 0);
        progress.hidden = false;
        progress.value = 0;
        status.textContent = 'Uploading…';
        pauseButton.hidden = false;
        pauseButton.textContent = 'Pause';
        try {
          let next = 0;
          const worker = async () => {
            while (next < records.length) {
              const index = next++;
              await uploadOne(records[index], progressBytes, total);
            }
          };
          await Promise.all(Array.from({ length: Math.min(3, records.length) }, worker));
          progress.value = 100;
          status.textContent = 'Uploaded ' + records.length + ' file(s). Reloading…';
          window.setTimeout(() => window.location.reload(), 300);
        } catch (error) {
          status.textContent = error.message || 'Upload failed.';
        } finally {
          uploading = false;
          pauseButton.hidden = true;
          paused = false;
          resumeWaiters.forEach((resolve) => resolve());
          resumeWaiters = [];
        }
      };
      pauseButton.addEventListener('click', () => {
        if (!uploading) return;
        paused = !paused;
        pauseButton.textContent = paused ? 'Resume' : 'Pause';
        status.textContent = paused ? 'Paused after the current chunk.' : 'Resuming…';
        if (!paused) {
          resumeWaiters.forEach((resolve) => resolve());
          resumeWaiters = [];
        }
      });
      filePicker.addEventListener('change', () => upload(pickerRecords(filePicker)));
      directoryPicker.addEventListener('change', () => upload(pickerRecords(directoryPicker)));
      dropZone.addEventListener('dragover', (event) => { event.preventDefault(); dropZone.classList.add('dragover'); });
      dropZone.addEventListener('dragleave', () => dropZone.classList.remove('dragover'));
      dropZone.addEventListener('drop', async (event) => {
        event.preventDefault();
        dropZone.classList.remove('dragover');
        try { upload(await dropRecords(event.dataTransfer)); }
        catch (_) { status.textContent = 'Could not read dropped files.'; }
      });
    })();
  </script>
</body>
</html>`))

var markdownTemplate = template.Must(template.New("markdown").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Name}} · {{.Host}}</title>
  <style>
    :root { color-scheme: light dark; font: 16px/1.5 system-ui, sans-serif; }
    :root { --code-bg: #e9eff6; --code-border: #cbd5e1; --code-text: #1e293b; }
    @media (prefers-color-scheme: dark) {
      :root { --code-bg: #202b3b; --code-border: #3b4a60; --code-text: #e5edf7; }
    }
    body { max-width: 1000px; margin: 0 auto; padding: 1rem; }
    nav { margin-bottom: 1rem; overflow-wrap: anywhere; }
    nav a { text-decoration: none; }
    .toolbar { display: flex; gap: .8rem; margin-bottom: 1rem; }
    #rendered, #source, #rendered pre { overflow-x: auto; }
    #rendered pre:not(:has(code.language-mermaid)), #source {
      box-sizing: border-box;
      padding: 1rem 1.25rem;
      border: 1px solid var(--code-border);
      border-radius: .5rem;
      background: var(--code-bg);
      color: var(--code-text);
    }
    #source { white-space: pre-wrap; font: 13px/1.45 ui-monospace, monospace; }
    .notice { padding: .8rem; border: 1px solid #8888; border-radius: .4rem; }
  </style>
</head>
<body>
  <nav aria-label="Breadcrumb">{{.Breadcrumb}}</nav>
  <h1>{{.Name}}</h1>
  <p><code>{{.RemotePath}}</code></p>
  <div class="toolbar"><a href="{{.RawURL}}">Raw</a><button id="source-toggle" type="button">Source</button></div>
  <div id="notice" class="notice" hidden></div>
  <article id="rendered"></article>
  <pre id="source" hidden></pre>
  <script src="/_ykview/assets/v1/marked.js"></script>
  <script src="/_ykview/assets/v1/mermaid.js"></script>
  <script>
    const source = {{.SourceJSON}};
    const rendered = document.getElementById('rendered');
    const sourceView = document.getElementById('source');
    const notice = document.getElementById('notice');
    const showSource = (message) => {
      rendered.hidden = true;
      sourceView.hidden = false;
      sourceView.textContent = source;
      if (message) { notice.hidden = false; notice.textContent = message; }
    };
    document.getElementById('source-toggle').addEventListener('click', () => {
      const isHidden = sourceView.hidden;
      sourceView.hidden = !isHidden;
      rendered.hidden = isHidden;
    });
    try {
      if (!window.marked || !window.mermaid) {
        throw new Error('preview assets unavailable');
      }
      rendered.innerHTML = window.marked.parse(source, { gfm: true, breaks: false });
      window.mermaid.initialize({ startOnLoad: false });
      Promise.resolve(window.mermaid.run({ querySelector: 'pre code.language-mermaid' }))
        .catch(() => showSource('Rich preview could not be loaded. Showing Markdown source.'));
    } catch (_) {
      showSource('Rich preview could not be loaded. Showing Markdown source.');
    }
  </script>
</body>
</html>`))

var textTemplate = template.Must(template.New("text").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Name}} · {{.Host}}</title>
  <style>
    :root { color-scheme: light dark; font: 16px/1.5 system-ui, sans-serif; }
    body { max-width: 1200px; margin: 0 auto; padding: 1rem; }
    nav { margin-bottom: 1rem; overflow-wrap: anywhere; }
    nav a { text-decoration: none; }
    .toolbar { display: flex; gap: .8rem; margin-bottom: 1rem; }
    pre { overflow: auto; padding: 1rem; border-radius: .4rem; background: #8882; }
    .notice { padding: .8rem; margin: 1rem 0; border: 1px solid #8888; border-radius: .4rem; }
  </style>
  <link rel="stylesheet" href="/_ykview/assets/v1/highlight.css">
</head>
<body>
  <nav aria-label="Breadcrumb">{{.Breadcrumb}}</nav>
  <h1>{{.Name}}</h1>
  <p><code>{{.RemotePath}}</code></p>
  <div class="toolbar"><a href="{{.RawURL}}">Raw</a></div>
  {{if .Notice}}<p class="notice" role="status">{{.Notice}}</p>{{end}}
  <pre><code id="source-code" class="language-{{.Language}}">{{.Source}}</code></pre>
  <script src="/_ykview/assets/v1/highlight.js"></script>
  <script>
    const source = {{.SourceJSON}};
    const language = {{.LanguageJSON}};
    const sourceCode = document.getElementById('source-code');
    sourceCode.textContent = source;
    if (language && window.hljs && window.hljs.getLanguage(language)) {
      try {
        sourceCode.innerHTML = window.hljs.highlight(source, { language }).value;
      } catch (_) {
        sourceCode.textContent = source;
      }
    }
  </script>
</body>
</html>`))

var delimitedTemplate = template.Must(template.New("delimited").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Name}} · {{.Host}}</title>
  <style>
    :root { color-scheme: light dark; font: 16px/1.5 system-ui, sans-serif; }
    body { max-width: 1200px; margin: 0 auto; padding: 1rem; }
    nav { margin-bottom: 1rem; overflow-wrap: anywhere; }
    nav a { text-decoration: none; }
    .toolbar { display: flex; gap: .8rem; margin-bottom: 1rem; }
    .notice { padding: .8rem; margin: 1rem 0; border: 1px solid #8888; border-radius: .4rem; }
    .table-scroll { overflow-x: auto; }
    table { border-collapse: collapse; width: 100%; }
    th, td { border: 1px solid #8885; padding: .45rem .6rem; text-align: left; vertical-align: top; }
    th { position: sticky; top: 0; background: Canvas; }
    td { max-width: 32rem; white-space: pre-wrap; overflow-wrap: anywhere; }
    pre { overflow: auto; padding: 1rem; border-radius: .4rem; background: #8882; }
  </style>
</head>
<body>
  <nav aria-label="Breadcrumb">{{.Breadcrumb}}</nav>
  <h1>{{.Name}}</h1>
  <p><code>{{.RemotePath}}</code></p>
  <div class="toolbar">
    <a href="{{.RawURL}}">Raw</a>
    {{if .SourceView}}<a href="{{.TableURL}}">Table</a>{{else}}<a href="{{.SourceURL}}">Source</a>{{end}}
  </div>
  {{if .LimitMessage}}<p class="notice" role="status">{{.LimitMessage}} <a href="{{.SourceURL}}">View full source</a></p>{{end}}
  {{if .SourceView}}
    <pre>{{.Source}}</pre>
  {{else if .Empty}}
    <p>Empty CSV/TSV file.</p>
  {{else}}
    <div class="table-scroll">
      <table>
        <thead><tr>{{range .Headers}}<th scope="col">{{.}}</th>{{end}}</tr></thead>
        <tbody>
          {{range .Rows}}<tr>{{range .}}<td>{{.}}</td>{{end}}</tr>{{else}}<tr><td colspan="{{len .Headers}}">No records.</td></tr>{{end}}
        </tbody>
      </table>
    </div>
  {{end}}
</body>
</html>`))

var mediaTemplate = template.Must(template.New("media").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Name}} · {{.Host}}</title>
  <style>
    :root { color-scheme: light dark; font: 16px/1.5 system-ui, sans-serif; }
    body { max-width: 1000px; margin: 0 auto; padding: 1rem; }
    nav { margin-bottom: 1rem; overflow-wrap: anywhere; }
    nav a { text-decoration: none; }
    .toolbar { display: flex; gap: .8rem; margin-bottom: 1rem; }
    audio, video { display: block; width: min(100%, 60rem); margin: 1rem 0; }
    video { max-height: 75vh; background: #000; }
  </style>
</head>
<body>
  <nav aria-label="Breadcrumb">{{.Breadcrumb}}</nav>
  <h1>{{.Name}}</h1>
  <p><code>{{.RemotePath}}</code></p>
  <div class="toolbar"><a href="{{.RawURL}}">Raw</a><a href="{{.DownloadURL}}">Download</a></div>
  {{if eq .Kind "audio"}}
    <audio controls preload="metadata">
      <source src="{{.RawURL}}" type="{{.ContentType}}">
      Your browser cannot play this audio format. Use Raw or Download.
    </audio>
  {{else}}
    <video controls preload="metadata">
      <source src="{{.RawURL}}" type="{{.ContentType}}">
      Your browser cannot play this video format. Use Raw or Download.
    </video>
  {{end}}
</body>
</html>`))
