/* ============================================================
   MER Viewer – app.js
   Complete, interactive MER diagram renderer (SVG) with
   pan, zoom, theme toggle and tab switching.
   ============================================================ */

// ── Theme Toggle ──────────────────────────────────────────────
const themeToggle = document.getElementById('theme-toggle');
const rootEl = document.documentElement;
const themeIcon = themeToggle.querySelector('i');

themeToggle.addEventListener('click', () => {
    const isDark = rootEl.getAttribute('data-theme') === 'dark';
    if (isDark) {
        rootEl.removeAttribute('data-theme');
        themeIcon.classList.remove('fa-moon');
        themeIcon.classList.add('fa-sun');
    } else {
        rootEl.setAttribute('data-theme', 'dark');
        themeIcon.classList.remove('fa-sun');
        themeIcon.classList.add('fa-moon');
    }
});

// ── Tab Logic ─────────────────────────────────────────────────
const tabBtns = document.querySelectorAll('.tab-btn');
const caseContents = document.querySelectorAll('.case-content');
const editorTitle = document.getElementById('editor-title');

tabBtns.forEach(btn => {
    btn.addEventListener('click', () => {
        tabBtns.forEach(b => b.classList.remove('active'));
        caseContents.forEach(c => c.classList.remove('active'));
        btn.classList.add('active');
        const target = btn.getAttribute('data-target');
        document.getElementById(target).classList.add('active');

        if (target === 'case1') {
            editorTitle.textContent = 'caso_1_tiquetes.mer';
            renderDiagram(case1Data);
        } else {
            editorTitle.textContent = 'caso_2_veterinaria.mer';
            renderDiagram(case2Data);
        }
        resetZoom();
    });
});

// ── Pan & Zoom ────────────────────────────────────────────────
let zoom = 1;
let panX = 0;
let panY = 0;
let dragging = false;
let dragStartX, dragStartY;

const viewport = document.getElementById('viewport');
const svgContainer = document.getElementById('svg-container');

document.getElementById('zoom-in').addEventListener('click', () => {
    zoom = Math.min(zoom + 0.15, 3);
    applyTransform();
});
document.getElementById('zoom-out').addEventListener('click', () => {
    zoom = Math.max(zoom - 0.15, 0.25);
    applyTransform();
});
document.getElementById('reset-view').addEventListener('click', resetZoom);

// Mouse wheel zoom
viewport.addEventListener('wheel', (e) => {
    e.preventDefault();
    const delta = e.deltaY > 0 ? -0.08 : 0.08;
    zoom = Math.max(0.25, Math.min(3, zoom + delta));
    applyTransform();
}, { passive: false });

function resetZoom() {
    zoom = 0.72;
    panX = 40;
    panY = 20;
    applyTransform();
}

function applyTransform() {
    svgContainer.style.transform = `translate(${panX}px, ${panY}px) scale(${zoom})`;
}

viewport.addEventListener('mousedown', (e) => {
    dragging = true;
    dragStartX = e.clientX - panX;
    dragStartY = e.clientY - panY;
    viewport.style.cursor = 'grabbing';
});
window.addEventListener('mouseup', () => {
    dragging = false;
    viewport.style.cursor = '';
});
window.addEventListener('mousemove', (e) => {
    if (!dragging) return;
    e.preventDefault();
    panX = e.clientX - dragStartX;
    panY = e.clientY - dragStartY;
    applyTransform();
});
viewport.addEventListener('dragstart', e => e.preventDefault());

// ── Data Models ───────────────────────────────────────────────

/*
  Layout strategy — Case 1:
  Row 1 (y≈170):  PASAJERO ── realiza ── RESERVA ── se_cubre_con ── PAGO
  Row 2 (y≈540):  EQUIPAJE ── registra ── TIQUETE ── recibe ── ASIGNACION ── para ── VUELO ═══ AEROPUERTO
  Row 3 (y≈910):                         EMBARQUE              ASIENTO ── contiene ── AERONAVE
  Vertical: RESERVA↓genera↓TIQUETE  TIQUETE↓habilita↓EMBARQUE  ASIG↓ocupa↓ASIENTO  VUELO↓usa↓AERONAVE
*/

const case1Data = {
    canvasW: 2650,
    canvasH: 1120,
    entities: [
        // ─── Row 1 ───
        { id:'E_PAS', name:'PASAJERO', x:250, y:170, type:'strong', attrs:[
            {name:'documento',        isPk:true,  dx:-105, dy:-78},
            {name:'nombre',           isPk:false, dx:15,   dy:-100},
            {name:'fecha_nacimiento', isPk:false, dx:140,  dy:-68}
        ]},
        { id:'E_RES', name:'RESERVA', x:750, y:170, type:'strong', attrs:[
            {name:'código', isPk:true,  dx:-105, dy:-78},
            {name:'fecha',  isPk:false, dx:15,   dy:-100},
            {name:'estado', isPk:false, dx:120,  dy:-72}
        ]},
        { id:'E_PAG', name:'PAGO', x:1250, y:170, type:'strong', attrs:[
            {name:'referencia', isPk:true,  dx:0,    dy:-90},
            {name:'fecha',      isPk:false, dx:105,  dy:-62},
            {name:'valor',      isPk:false, dx:-105, dy:-62}
        ]},

        // ─── Row 2 ───
        { id:'E_EQU', name:'EQUIPAJE', x:250, y:540, type:'strong', attrs:[
            {name:'etiqueta', isPk:true,  dx:0,   dy:82},
            {name:'peso',     isPk:false, dx:-95, dy:62},
            {name:'estado',   isPk:false, dx:95,  dy:62}
        ]},
        { id:'E_TIQ', name:'TIQUETE', x:750, y:540, type:'strong', attrs:[
            {name:'número',         isPk:true,  dx:-165, dy:-68},
            {name:'fecha_emisión',  isPk:false, dx:-168, dy:0},
            {name:'clase_servicio', isPk:false, dx:-165, dy:68}
        ]},
        { id:'E_ASIG', name:'ASIGNACION_ASIENTO', x:1200, y:540, type:'associative', attrs:[
            {name:'tiquete+vuelo+asiento', isPk:true, dx:0, dy:-82}
        ]},
        { id:'E_VUE', name:'VUELO', x:1660, y:540, type:'strong', attrs:[
            {name:'número',          isPk:true,  dx:-92, dy:-88},
            {name:'fecha_salida',    isPk:true,  dx:0,   dy:-108},
            {name:'hora_programada', isPk:false, dx:100, dy:-82}
        ]},
        { id:'E_AERO', name:'AEROPUERTO', x:2150, y:540, type:'strong', attrs:[
            {name:'código', isPk:true,  dx:120, dy:-68},
            {name:'nombre', isPk:false, dx:135, dy:0},
            {name:'ciudad', isPk:false, dx:120, dy:68}
        ]},

        // ─── Row 3 ───
        { id:'E_EMB', name:'EMBARQUE', x:750, y:910, type:'weak', attrs:[
            {name:'número_tiquete',          isPk:true,  dx:-170, dy:-25},
            {name:'hora_ingreso',            isPk:false, dx:-155, dy:55},
            {name:'puerta',                  isPk:false, dx:0,    dy:92},
            {name:'condición_presentación',  isPk:false, dx:175,  dy:52}
        ]},
        { id:'E_ASIE', name:'ASIENTO', x:1200, y:910, type:'weak', attrs:[
            {name:'número',    isPk:true,  dx:-95, dy:78},
            {name:'matrícula', isPk:true,  dx:0,   dy:98},
            {name:'fila',      isPk:false, dx:95,  dy:78},
            {name:'ubicación', isPk:false, dx:140, dy:30}
        ]},
        { id:'E_NAVE', name:'AERONAVE', x:1660, y:910, type:'strong', attrs:[
            {name:'matrícula', isPk:true,  dx:0,    dy:85},
            {name:'modelo',    isPk:false, dx:-105, dy:65},
            {name:'capacidad', isPk:false, dx:105,  dy:65}
        ]}
    ],
    relationships: [
        // Row 1 horizontal
        { id:'R_REALIZA', name:'realiza',       x:500,  y:170 },
        { id:'R_CUBRE',   name:'se_cubre_con',  x:1000, y:170 },
        // Vertical Row 1↔2
        { id:'R_GENERA',  name:'genera',         x:750,  y:355 },
        // Row 2 horizontal
        { id:'R_REGISTRA',name:'registra',       x:500,  y:540 },
        { id:'R_RECIBE',  name:'recibe',         x:975,  y:540 },
        { id:'R_PARA',    name:'para',           x:1430, y:540 },
        // VUELO↔AEROPUERTO (two roles, offset vertically)
        { id:'R_SALE',    name:'sale_de',         x:1905, y:440 },
        { id:'R_LLEGA',   name:'llega_a',         x:1905, y:640 },
        // Vertical Row 2↔3
        { id:'R_HABILITA',name:'habilita',       x:750,  y:725 },
        { id:'R_OCUPA',   name:'ocupa',          x:1200, y:725 },
        { id:'R_USA',     name:'usa',            x:1660, y:725 },
        // Row 3 horizontal
        { id:'R_CONTIENE',name:'contiene',       x:1430, y:910 }
    ],
    links: [
        // Row 1
        { source:'E_PAS',     target:'R_REALIZA',  card:'1',    cardDx:12,  cardDy:-18 },
        { source:'R_REALIZA', target:'E_RES',      card:'0..N', cardDx:-38, cardDy:-18 },
        { source:'E_RES',     target:'R_CUBRE',    card:'1',    cardDx:12,  cardDy:-18 },
        { source:'R_CUBRE',   target:'E_PAG',      card:'0..1', cardDx:-38, cardDy:-18 },
        // Vertical
        { source:'E_RES',     target:'R_GENERA',   card:'1',    cardDx:14,  cardDy:-14 },
        { source:'R_GENERA',  target:'E_TIQ',      card:'0..1', cardDx:14,  cardDy:-18 },
        // Row 2
        { source:'E_TIQ',     target:'R_REGISTRA', card:'1',    cardDx:-18, cardDy:-18 },
        { source:'R_REGISTRA',target:'E_EQU',      card:'0..N', cardDx:22,  cardDy:-18 },
        { source:'E_TIQ',     target:'R_RECIBE',   card:'1',    cardDx:14,  cardDy:-18 },
        { source:'R_RECIBE',  target:'E_ASIG',     card:'1',    cardDx:-28, cardDy:-18 },
        { source:'E_ASIG',    target:'R_PARA',     card:'N',    cardDx:14,  cardDy:-18 },
        { source:'R_PARA',    target:'E_VUE',      card:'1',    cardDx:-28, cardDy:-18 },
        // VUELO↔AEROPUERTO double role
        { source:'E_VUE',     target:'R_SALE',     card:'0..N', cardDx:-14, cardDy:-22 },
        { source:'R_SALE',    target:'E_AERO',     card:'1',    cardDx:-14, cardDy:-22 },
        { source:'E_VUE',     target:'R_LLEGA',    card:'0..N', cardDx:-14, cardDy:14 },
        { source:'R_LLEGA',   target:'E_AERO',     card:'1',    cardDx:-14, cardDy:14 },
        // Vertical Row 2↔3
        { source:'E_TIQ',     target:'R_HABILITA', card:'1',    cardDx:14,  cardDy:-14 },
        { source:'R_HABILITA',target:'E_EMB',      card:'0..1', cardDx:14,  cardDy:-20 },
        { source:'E_ASIG',    target:'R_OCUPA',    card:'N',    cardDx:14,  cardDy:-14 },
        { source:'R_OCUPA',   target:'E_ASIE',     card:'1',    cardDx:14,  cardDy:-20 },
        { source:'E_VUE',     target:'R_USA',      card:'0..N', cardDx:14,  cardDy:-14 },
        { source:'R_USA',     target:'E_NAVE',     card:'1',    cardDx:14,  cardDy:-20 },
        // Row 3
        { source:'E_NAVE',    target:'R_CONTIENE', card:'1',    cardDx:-28, cardDy:-18 },
        { source:'R_CONTIENE',target:'E_ASIE',     card:'1..N', cardDx:28,  cardDy:-18 }
    ]
};

/*
  Layout strategy — Case 2:
  Row 1 (y≈170):  PROPIETARIO ── posee ── MASCOTA   ···   PRESCRIPCION ── indica ── MEDICAMENTO
  Row 2 (y≈530):  VETERINARIO ── atiende ── CITA    ···   TRATAMIENTO              LOTE
  Row 3 (y≈880):                            CONSULTA ── determina ── DIAGNOSTICO
  Vertical: MASCOTA↓agenda↓CITA  CITA↓deriva↓CONSULTA  DIAG↑origina↑TRAT  TRAT↑incluye↑PRESC  MED↓dispone↓LOTE
*/

const case2Data = {
    canvasW: 2100,
    canvasH: 1080,
    entities: [
        // ─── Row 1 ───
        { id:'E_PROP', name:'PROPIETARIO', x:250, y:170, type:'strong', attrs:[
            {name:'documento', isPk:true,  dx:-108, dy:-75},
            {name:'nombre',    isPk:false, dx:12,   dy:-100},
            {name:'teléfono',  isPk:false, dx:118,  dy:-70}
        ]},
        { id:'E_MASC', name:'MASCOTA', x:720, y:170, type:'strong', attrs:[
            {name:'código_clínico',   isPk:true,  dx:-115, dy:-80},
            {name:'nombre',           isPk:false, dx:12,   dy:-102},
            {name:'fecha_nacimiento', isPk:false, dx:145,  dy:-62}
        ]},
        { id:'E_PRES', name:'PRESCRIPCION', x:1280, y:170, type:'associative', attrs:[
            {name:'cod_tratamiento+cod_medicamento', isPk:true, dx:0, dy:-92},
            {name:'dosis',               isPk:false, dx:-160, dy:-25},
            {name:'frecuencia',          isPk:false, dx:-155, dy:38},
            {name:'vía_administración',  isPk:false, dx:-120, dy:88}
        ]},
        { id:'E_MED', name:'MEDICAMENTO', x:1740, y:170, type:'strong', attrs:[
            {name:'código',        isPk:true,  dx:0,    dy:-90},
            {name:'nombre',        isPk:false, dx:-110, dy:-62},
            {name:'concentración', isPk:false, dx:115,  dy:-62}
        ]},

        // ─── Row 2 ───
        { id:'E_VET', name:'VETERINARIO', x:250, y:530, type:'strong', attrs:[
            {name:'tarjeta_profesional', isPk:true,  dx:-142, dy:-38},
            {name:'nombre',              isPk:false, dx:-125, dy:22},
            {name:'especialidad',        isPk:false, dx:-108, dy:75}
        ]},
        { id:'E_CITA', name:'CITA', x:720, y:530, type:'strong', attrs:[
            {name:'código', isPk:true,  dx:135, dy:-40},
            {name:'fecha',  isPk:false, dx:138, dy:18},
            {name:'motivo', isPk:false, dx:118, dy:72}
        ]},
        { id:'E_TRAT', name:'TRATAMIENTO', x:1280, y:530, type:'weak', attrs:[
            {name:'código_diagnóstico', isPk:true,  dx:155, dy:-50},
            {name:'fecha_inicio',       isPk:false, dx:162, dy:8},
            {name:'duración',           isPk:false, dx:142, dy:58},
            {name:'instrucciones',      isPk:false, dx:108, dy:100}
        ]},
        { id:'E_LOTE', name:'LOTE', x:1740, y:530, type:'weak', attrs:[
            {name:'número',              isPk:true,  dx:115, dy:-52},
            {name:'código_medicamento',  isPk:true,  dx:148, dy:0},
            {name:'fecha_vencimiento',   isPk:false, dx:128, dy:52},
            {name:'cantidad_disponible', isPk:false, dx:82,  dy:98}
        ]},

        // ─── Row 3 ───
        { id:'E_CONS', name:'CONSULTA', x:720, y:880, type:'strong', attrs:[
            {name:'consecutivo', isPk:true,  dx:-132, dy:-28},
            {name:'fecha',       isPk:false, dx:-128, dy:30},
            {name:'peso',        isPk:false, dx:-105, dy:80}
        ]},
        { id:'E_DIAG', name:'DIAGNOSTICO', x:1280, y:880, type:'strong', attrs:[
            {name:'código',          isPk:true,  dx:0,    dy:88},
            {name:'descripción',     isPk:false, dx:-115, dy:65},
            {name:'nivel_gravedad',  isPk:false, dx:120,  dy:65}
        ]}
    ],
    relationships: [
        // Row 1 horizontal
        { id:'R_POSEE',  name:'posee',   x:485,  y:170 },
        { id:'R_INDICA', name:'indica',  x:1510, y:170 },
        // Vertical Row 1↔2
        { id:'R_AGENDA',  name:'agenda',    x:720,  y:350 },
        { id:'R_INCLUYE', name:'incluye',   x:1280, y:350 },
        { id:'R_DISPONE', name:'dispone_de',x:1740, y:350 },
        // Row 2 horizontal
        { id:'R_ATIENDE', name:'atiende', x:485, y:530 },
        // Vertical Row 2↔3
        { id:'R_DERIVA',    name:'deriva_en', x:720,  y:705 },
        { id:'R_ORIGINA',   name:'origina',   x:1280, y:705 },
        // Row 3 horizontal
        { id:'R_DETERMINA', name:'determina', x:1000, y:880 }
    ],
    links: [
        // Row 1
        { source:'E_PROP',  target:'R_POSEE',    card:'1',    cardDx:12,  cardDy:-18 },
        { source:'R_POSEE', target:'E_MASC',     card:'1..N', cardDx:-38, cardDy:-18 },
        { source:'E_PRES',  target:'R_INDICA',   card:'0..N', cardDx:14,  cardDy:-18 },
        { source:'R_INDICA',target:'E_MED',      card:'1',    cardDx:-28, cardDy:-18 },
        // Vertical Row 1↔2
        { source:'E_MASC',   target:'R_AGENDA',  card:'1',    cardDx:14,  cardDy:-14 },
        { source:'R_AGENDA', target:'E_CITA',    card:'0..N', cardDx:14,  cardDy:-20 },
        { source:'E_TRAT',   target:'R_INCLUYE', card:'1',    cardDx:14,  cardDy:-14 },
        { source:'R_INCLUYE',target:'E_PRES',    card:'0..N', cardDx:14,  cardDy:-20 },
        { source:'E_MED',    target:'R_DISPONE', card:'1',    cardDx:14,  cardDy:-14 },
        { source:'R_DISPONE',target:'E_LOTE',    card:'0..N', cardDx:14,  cardDy:-20 },
        // Row 2
        { source:'E_VET',     target:'R_ATIENDE',card:'0..N', cardDx:14,  cardDy:-18 },
        { source:'R_ATIENDE', target:'E_CITA',   card:'1',    cardDx:-28, cardDy:-18 },
        // Vertical Row 2↔3
        { source:'E_CITA',  target:'R_DERIVA',   card:'1',    cardDx:14,  cardDy:-14 },
        { source:'R_DERIVA',target:'E_CONS',     card:'0..1', cardDx:14,  cardDy:-20 },
        { source:'E_DIAG',  target:'R_ORIGINA',  card:'1',    cardDx:14,  cardDy:18 },
        { source:'R_ORIGINA',target:'E_TRAT',    card:'1',    cardDx:14,  cardDy:18 },
        // Row 3
        { source:'E_CONS',      target:'R_DETERMINA', card:'1',    cardDx:14,  cardDy:-18 },
        { source:'R_DETERMINA', target:'E_DIAG',      card:'1..N', cardDx:-38, cardDy:-18 }
    ]
};

// ── SVG Rendering ─────────────────────────────────────────────

function measureText(text, fontSize) {
    return text.length * fontSize * 0.58;
}

function renderDiagram(data) {
    const W = data.canvasW || 2600;
    const H = data.canvasH || 1100;
    const ns = 'http://www.w3.org/2000/svg';
    let s = '';

    // --- Lines (draw first, behind everything) ---
    data.links.forEach(link => {
        const src = data.entities.find(e => e.id === link.source)
                 || data.relationships.find(r => r.id === link.source);
        const tgt = data.entities.find(e => e.id === link.target)
                 || data.relationships.find(r => r.id === link.target);
        if (!src || !tgt) return;

        s += `<line x1="${src.x}" y1="${src.y}" x2="${tgt.x}" y2="${tgt.y}" class="svg-line"/>`;

        // Cardinality label near the correct end
        const isSourceEntity = !!data.entities.find(e => e.id === link.source);
        let cx, cy;
        if (isSourceEntity) {
            cx = src.x + (tgt.x - src.x) * 0.18 + (link.cardDx || 0);
            cy = src.y + (tgt.y - src.y) * 0.18 + (link.cardDy || 0);
        } else {
            cx = tgt.x - (tgt.x - src.x) * 0.18 + (link.cardDx || 0);
            cy = tgt.y - (tgt.y - src.y) * 0.18 + (link.cardDy || 0);
        }
        s += `<text x="${cx}" y="${cy}" class="svg-cardinality">${link.card}</text>`;
    });

    // --- Attribute ellipses + connecting lines ---
    data.entities.forEach(ent => {
        ent.attrs.forEach(attr => {
            const ax = ent.x + attr.dx;
            const ay = ent.y + attr.dy;
            const textW = measureText(attr.name, 11);
            const rx = Math.max(32, textW / 2 + 10);
            const ry = 17;

            s += `<line x1="${ent.x}" y1="${ent.y}" x2="${ax}" y2="${ay}" class="svg-line"/>`;
            s += `<ellipse cx="${ax}" cy="${ay}" rx="${rx}" ry="${ry}" class="svg-attr-ellipse"/>`;
            if (attr.isPk) {
                s += `<text x="${ax}" y="${ay}" class="svg-attr-text pk">${attr.name}</text>`;
            } else {
                s += `<text x="${ax}" y="${ay}" class="svg-attr-text">${attr.name}</text>`;
            }
        });
    });

    // --- Relationship diamonds ---
    data.relationships.forEach(rel => {
        const d = 42;
        const pts = `${rel.x},${rel.y - d * 0.65} ${rel.x + d},${rel.y} ${rel.x},${rel.y + d * 0.65} ${rel.x - d},${rel.y}`;
        s += `<polygon points="${pts}" class="svg-rel-diamond"/>`;
        s += `<text x="${rel.x}" y="${rel.y}" class="svg-rel-text">${rel.name}</text>`;
    });

    // --- Entity rectangles ---
    data.entities.forEach(ent => {
        const textW = measureText(ent.name, 13);
        const w = Math.max(120, textW + 28);
        const h = 42;
        const rx = ent.x - w / 2;
        const ry = ent.y - h / 2;

        if (ent.type === 'weak') {
            // Outer rect (double border effect)
            s += `<rect x="${rx - 5}" y="${ry - 5}" width="${w + 10}" height="${h + 10}" rx="5" class="svg-ent-outer-weak"/>`;
            s += `<rect x="${rx}" y="${ry}" width="${w}" height="${h}" rx="4" class="svg-ent-rect svg-ent-weak"/>`;
        } else if (ent.type === 'associative') {
            // Rectangle + inner diamond overlay
            s += `<rect x="${rx}" y="${ry}" width="${w}" height="${h}" rx="4" class="svg-ent-rect svg-ent-assoc"/>`;
            const ad = 18;
            s += `<polygon points="${ent.x},${ent.y - ad} ${ent.x + ad},${ent.y} ${ent.x},${ent.y + ad} ${ent.x - ad},${ent.y}" class="svg-assoc-diamond"/>`;
        } else {
            s += `<rect x="${rx}" y="${ry}" width="${w}" height="${h}" rx="4" class="svg-ent-rect"/>`;
        }
        s += `<text x="${ent.x}" y="${ent.y}" class="svg-ent-text">${ent.name}</text>`;
    });

    svgContainer.innerHTML = `<svg width="${W}" height="${H}" xmlns="http://www.w3.org/2000/svg">${s}</svg>`;
}

// ── Init ──────────────────────────────────────────────────────
renderDiagram(case1Data);
resetZoom();
