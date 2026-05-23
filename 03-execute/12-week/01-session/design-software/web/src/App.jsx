import { useEffect, useMemo, useState } from 'react';
import {
  AlertTriangle,
  CalendarDays,
  Clock,
  ClipboardList,
  Database,
  Plus,
  RefreshCw,
  Save,
  School,
  Search,
  UserRound,
} from 'lucide-react';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';
const DEMO_STORE_KEY = 'sena-horarios-demo-store-v2';

const emptySchedule = {
  date: '',
  startTime: '',
  endTime: '',
  instructorId: '',
  fichaId: '',
  environmentId: '',
  programId: '',
  competence: '',
  rap: '',
  learningActivity: '',
  observation: '',
};

function App() {
  const [instructors, setInstructors] = useState([]);
  const [programs, setPrograms] = useState([]);
  const [fichas, setFichas] = useState([]);
  const [environments, setEnvironments] = useState([]);
  const [schedules, setSchedules] = useState([]);
  const [noProgrammingDays, setNoProgrammingDays] = useState([]);
  const [schedule, setSchedule] = useState(emptySchedule);
  const [notice, setNotice] = useState(null);
  const [loading, setLoading] = useState(false);
  const [filter, setFilter] = useState('');
  const [selectedInstructorId, setSelectedInstructorId] = useState('');
  const [demoMode, setDemoMode] = useState(false);

  const [catalogForms, setCatalogForms] = useState({
    instructorName: '',
    instructorType: 'planta',
    programName: '',
    programLevel: '',
    fichaCode: '',
    fichaProgramId: '',
    environmentName: '',
  });

  const [noProgrammingForm, setNoProgrammingForm] = useState({
    date: '',
    fichaId: '',
    instructorId: '',
    reason: '',
  });

  useEffect(() => {
    loadAll();
  }, []);

  useEffect(() => {
    if (instructors.length > 0 && !instructors.some((item) => String(item.id) === String(selectedInstructorId))) {
      setSelectedInstructorId(String(instructors[0].id));
    }
  }, [instructors, selectedInstructorId]);

  const filteredSchedules = useMemo(() => {
    const term = filter.trim().toLowerCase();
    if (!term) return schedules;
    return schedules.filter((item) =>
      [
        item.date,
        item.fichaCode,
        item.instructorName,
        item.environmentName,
        item.programName,
        item.competence,
        item.rap,
        item.observation,
      ]
        .filter(Boolean)
        .some((value) => String(value).toLowerCase().includes(term))
    );
  }, [filter, schedules]);

  const instructorSchedules = useMemo(() => {
    const selected = selectedInstructorId || instructors[0]?.id;
    return schedules
      .filter((item) => String(item.instructorId) === String(selected))
      .sort((left, right) => `${left.date} ${left.startTime}`.localeCompare(`${right.date} ${right.startTime}`));
  }, [instructors, schedules, selectedInstructorId]);

  const selectedInstructor = useMemo(() => {
    const selected = selectedInstructorId || instructors[0]?.id;
    return instructors.find((item) => String(item.id) === String(selected));
  }, [instructors, selectedInstructorId]);

  async function loadAll() {
    setLoading(true);
    setNotice(null);
    try {
      const [instructorsData, programsData, fichasData, environmentsData, schedulesData, noProgrammingData] =
        await Promise.all([
          apiGet('/api/instructors'),
          apiGet('/api/programs'),
          apiGet('/api/fichas'),
          apiGet('/api/environments'),
          apiGet('/api/schedules'),
          apiGet('/api/no-programming-days'),
        ]);
      setInstructors(instructorsData);
      setPrograms(programsData);
      setFichas(fichasData);
      setEnvironments(environmentsData);
      setSchedules(schedulesData);
      setNoProgrammingDays(noProgrammingData);
      setDemoMode(false);
    } catch (error) {
      const demoStore = loadDemoStore();
      applyStore(demoStore);
      setDemoMode(true);
      setNotice({
        type: 'warning',
        message: 'API no disponible. Estas viendo un MVP funcional en modo demo local.',
      });
    } finally {
      setLoading(false);
    }
  }

  function applyStore(store) {
    setInstructors(store.instructors);
    setPrograms(store.programs);
    setFichas(store.fichas);
    setEnvironments(store.environments);
    setSchedules(store.schedules);
    setNoProgrammingDays(store.noProgrammingDays);
  }

  async function createCatalog(kind) {
    setNotice(null);
    try {
      if (demoMode) {
        const store = loadDemoStore();
        if (kind === 'instructor') {
          requireText(catalogForms.instructorName, 'Nombre del instructor');
          store.instructors.push({
            id: nextID(store.instructors),
            name: catalogForms.instructorName.trim(),
            instructorType: catalogForms.instructorType,
          });
          setCatalogForms((current) => ({ ...current, instructorName: '' }));
        }
        if (kind === 'program') {
          requireText(catalogForms.programName, 'Programa de formacion');
          store.programs.push({
            id: nextID(store.programs),
            name: catalogForms.programName.trim(),
            level: catalogForms.programLevel.trim(),
          });
          setCatalogForms((current) => ({ ...current, programName: '', programLevel: '' }));
        }
        if (kind === 'ficha') {
          requireText(catalogForms.fichaCode, 'Ficha');
          const program = findByID(store.programs, catalogForms.fichaProgramId);
          store.fichas.push({
            id: nextID(store.fichas),
            code: catalogForms.fichaCode.trim(),
            programId: optionalNumber(catalogForms.fichaProgramId),
            programName: program?.name || '',
          });
          setCatalogForms((current) => ({ ...current, fichaCode: '', fichaProgramId: '' }));
        }
        if (kind === 'environment') {
          requireText(catalogForms.environmentName, 'Ambiente de formacion');
          store.environments.push({
            id: nextID(store.environments),
            name: catalogForms.environmentName.trim(),
          });
          setCatalogForms((current) => ({ ...current, environmentName: '' }));
        }
        saveDemoStore(store);
        applyStore(store);
        setNotice({ type: 'success', message: 'Catalogo actualizado en modo demo.' });
        return;
      }

      if (kind === 'instructor') {
        await apiPost('/api/instructors', {
          name: catalogForms.instructorName,
          instructorType: catalogForms.instructorType,
        });
        setCatalogForms((current) => ({ ...current, instructorName: '' }));
      }
      if (kind === 'program') {
        await apiPost('/api/programs', {
          name: catalogForms.programName,
          level: catalogForms.programLevel,
        });
        setCatalogForms((current) => ({ ...current, programName: '', programLevel: '' }));
      }
      if (kind === 'ficha') {
        await apiPost('/api/fichas', {
          code: catalogForms.fichaCode,
          programId: optionalNumber(catalogForms.fichaProgramId),
        });
        setCatalogForms((current) => ({ ...current, fichaCode: '', fichaProgramId: '' }));
      }
      if (kind === 'environment') {
        await apiPost('/api/environments', { name: catalogForms.environmentName });
        setCatalogForms((current) => ({ ...current, environmentName: '' }));
      }
      await loadAll();
      setNotice({ type: 'success', message: 'Catalogo actualizado.' });
    } catch (error) {
      setNotice({ type: 'error', message: error.message });
    }
  }

  async function submitSchedule(event) {
    event.preventDefault();
    setNotice(null);
    try {
      if (demoMode) {
        const store = loadDemoStore();
        validateDemoSchedule(schedule);
        const conflicts = findDemoConflicts(store.schedules, schedule);
        if (conflicts.length) {
          throw new Error(`La programacion cruza con ${conflicts.join(', ')} ya asignado.`);
        }
        const instructor = findByID(store.instructors, schedule.instructorId);
        const ficha = findByID(store.fichas, schedule.fichaId);
        const environment = findByID(store.environments, schedule.environmentId);
        const program = findByID(store.programs, schedule.programId);
        store.schedules.push({
          id: nextID(store.schedules),
          ...schedule,
          instructorId: Number(schedule.instructorId),
          instructorName: instructor.name,
          instructorType: instructor.instructorType,
          fichaId: Number(schedule.fichaId),
          fichaCode: ficha.code,
          environmentId: Number(schedule.environmentId),
          environmentName: environment.name,
          programId: optionalNumber(schedule.programId),
          programName: program?.name || '',
          programLevel: program?.level || '',
        });
        saveDemoStore(store);
        applyStore(store);
        setSchedule(emptySchedule);
        setNotice({ type: 'success', message: 'Horario registrado en modo demo.' });
        return;
      }

      await apiPost('/api/schedules', {
        ...schedule,
        instructorId: Number(schedule.instructorId),
        fichaId: Number(schedule.fichaId),
        environmentId: Number(schedule.environmentId),
        programId: optionalNumber(schedule.programId),
      });
      setSchedule(emptySchedule);
      await loadAll();
      setNotice({ type: 'success', message: 'Horario registrado.' });
    } catch (error) {
      setNotice({ type: 'error', message: error.message });
    }
  }

  async function submitNoProgrammingDay(event) {
    event.preventDefault();
    setNotice(null);
    try {
      if (demoMode) {
        const store = loadDemoStore();
        requireText(noProgrammingForm.date, 'Fecha');
        requireText(noProgrammingForm.reason, 'Motivo u observacion');
        const ficha = findByID(store.fichas, noProgrammingForm.fichaId);
        const instructor = findByID(store.instructors, noProgrammingForm.instructorId);
        store.noProgrammingDays.push({
          id: nextID(store.noProgrammingDays),
          date: noProgrammingForm.date,
          fichaId: optionalNumber(noProgrammingForm.fichaId),
          fichaCode: ficha?.code || '',
          instructorId: optionalNumber(noProgrammingForm.instructorId),
          instructorName: instructor?.name || '',
          reason: noProgrammingForm.reason.trim(),
        });
        saveDemoStore(store);
        applyStore(store);
        setNoProgrammingForm({ date: '', fichaId: '', instructorId: '', reason: '' });
        setNotice({ type: 'success', message: 'Dia sin programacion registrado en modo demo.' });
        return;
      }

      await apiPost('/api/no-programming-days', {
        date: noProgrammingForm.date,
        fichaId: optionalNumber(noProgrammingForm.fichaId),
        instructorId: optionalNumber(noProgrammingForm.instructorId),
        reason: noProgrammingForm.reason,
      });
      setNoProgrammingForm({ date: '', fichaId: '', instructorId: '', reason: '' });
      await loadAll();
      setNotice({ type: 'success', message: 'Dia sin programacion registrado.' });
    } catch (error) {
      setNotice({ type: 'error', message: error.message });
    }
  }

  return (
    <main className="app-shell">
      <section className="top-bar">
        <div>
          <span className="eyebrow">MVP operativo</span>
          <h1>Sistema de Horarios SENA</h1>
          <p>Programacion basica de instructores, fichas y ambientes de formacion.</p>
        </div>
        <button className="icon-button" onClick={loadAll} title="Actualizar datos" disabled={loading}>
          <RefreshCw size={18} />
          <span>{loading ? 'Actualizando' : 'Actualizar'}</span>
        </button>
      </section>

      {notice && (
        <div className={`notice ${notice.type}`}>
          <AlertTriangle size={18} />
          <span>{notice.message}</span>
        </div>
      )}

      {demoMode && (
        <div className="demo-ribbon">
          Datos locales de demostracion. Para persistencia real, levanta `api` y `db` con contenedores.
        </div>
      )}

      <section className="metrics-band">
        <Metric icon={UserRound} label="Instructores" value={instructors.length} />
        <Metric icon={School} label="Fichas" value={fichas.length} />
        <Metric icon={Database} label="Ambientes" value={environments.length} />
        <Metric icon={CalendarDays} label="Horarios" value={schedules.length} />
      </section>

      <section className="panel instructor-schedule-panel">
        <div className="section-toolbar">
          <PanelTitle icon={Clock} title="Horario por instructor" />
          <label className="instructor-picker">
            Instructor
            <select value={selectedInstructorId} onChange={(event) => setSelectedInstructorId(event.target.value)}>
              {instructors.map((item) => (
                <option key={item.id} value={item.id}>
                  {item.name} ({item.instructorType})
                </option>
              ))}
            </select>
          </label>
        </div>

        <div className="instructor-summary">
          <div>
            <span className="eyebrow">Vista semanal demo</span>
            <h2>{selectedInstructor?.name || 'Instructor Pendiente'}</h2>
            <p>{selectedInstructor?.instructorType || 'Tipo Pendiente'} · {instructorSchedules.length} bloques programados</p>
          </div>
        </div>

        <div className="schedule-board">
          {instructorSchedules.map((item) => (
            <article className="schedule-card" key={item.id}>
              <div className="schedule-card-time">
                <strong>{item.startTime} - {item.endTime}</strong>
                <span>{formatDisplayDate(item.date)}</span>
              </div>
              <div className="schedule-card-body">
                <h3>{item.fichaCode}</h3>
                <p>{item.environmentName}</p>
                <span>{item.programName || 'Programa Pendiente'}</span>
              </div>
              <div className="schedule-card-note">
                <strong>{item.rap || 'RAP Pendiente'}</strong>
                <span>{item.observation || 'Sin novedad'}</span>
              </div>
            </article>
          ))}
          {instructorSchedules.length === 0 && (
            <div className="empty-state schedule-empty">No hay bloques programados para este instructor.</div>
          )}
        </div>
      </section>

      <section className="workspace-grid">
        <div className="panel catalog-panel">
          <PanelTitle icon={Plus} title="Catalogos basicos" />
          <CatalogBlock
            label="Instructor"
            fields={
              <>
                <input
                  placeholder="Nombre del instructor"
                  value={catalogForms.instructorName}
                  onChange={(event) => setCatalogForms({ ...catalogForms, instructorName: event.target.value })}
                />
                <select
                  value={catalogForms.instructorType}
                  onChange={(event) => setCatalogForms({ ...catalogForms, instructorType: event.target.value })}
                >
                  <option value="planta">Planta</option>
                  <option value="contratista">Contratista</option>
                </select>
              </>
            }
            onSubmit={() => createCatalog('instructor')}
          />
          <CatalogBlock
            label="Programa"
            fields={
              <>
                <input
                  placeholder="Programa de formacion"
                  value={catalogForms.programName}
                  onChange={(event) => setCatalogForms({ ...catalogForms, programName: event.target.value })}
                />
                <input
                  placeholder="Nivel"
                  value={catalogForms.programLevel}
                  onChange={(event) => setCatalogForms({ ...catalogForms, programLevel: event.target.value })}
                />
              </>
            }
            onSubmit={() => createCatalog('program')}
          />
          <CatalogBlock
            label="Ficha"
            fields={
              <>
                <input
                  placeholder="Ficha"
                  value={catalogForms.fichaCode}
                  onChange={(event) => setCatalogForms({ ...catalogForms, fichaCode: event.target.value })}
                />
                <select
                  value={catalogForms.fichaProgramId}
                  onChange={(event) => setCatalogForms({ ...catalogForms, fichaProgramId: event.target.value })}
                >
                  <option value="">Programa Pendiente</option>
                  {programs.map((programItem) => (
                    <option key={programItem.id} value={programItem.id}>
                      {programItem.name}
                    </option>
                  ))}
                </select>
              </>
            }
            onSubmit={() => createCatalog('ficha')}
          />
          <CatalogBlock
            label="Ambiente"
            fields={
              <input
                placeholder="Ambiente de formacion"
                value={catalogForms.environmentName}
                onChange={(event) => setCatalogForms({ ...catalogForms, environmentName: event.target.value })}
              />
            }
            onSubmit={() => createCatalog('environment')}
          />
        </div>

        <form className="panel schedule-form" onSubmit={submitSchedule}>
          <PanelTitle icon={ClipboardList} title="Registrar horario" />
          <div className="form-grid">
            <label>
              Fecha
              <input type="date" value={schedule.date} onChange={(event) => setSchedule({ ...schedule, date: event.target.value })} required />
            </label>
            <label>
              Inicio
              <input type="time" value={schedule.startTime} onChange={(event) => setSchedule({ ...schedule, startTime: event.target.value })} required />
            </label>
            <label>
              Fin
              <input type="time" value={schedule.endTime} onChange={(event) => setSchedule({ ...schedule, endTime: event.target.value })} required />
            </label>
            <label>
              Instructor
              <select value={schedule.instructorId} onChange={(event) => setSchedule({ ...schedule, instructorId: event.target.value })} required>
                <option value="">Seleccionar</option>
                {instructors.map((item) => (
                  <option key={item.id} value={item.id}>
                    {item.name} ({item.instructorType})
                  </option>
                ))}
              </select>
            </label>
            <label>
              Ficha
              <select value={schedule.fichaId} onChange={(event) => setSchedule({ ...schedule, fichaId: event.target.value })} required>
                <option value="">Seleccionar</option>
                {fichas.map((item) => (
                  <option key={item.id} value={item.id}>
                    {item.code}
                  </option>
                ))}
              </select>
            </label>
            <label>
              Ambiente
              <select value={schedule.environmentId} onChange={(event) => setSchedule({ ...schedule, environmentId: event.target.value })} required>
                <option value="">Seleccionar</option>
                {environments.map((item) => (
                  <option key={item.id} value={item.id}>
                    {item.name}
                  </option>
                ))}
              </select>
            </label>
            <label>
              Programa
              <select value={schedule.programId} onChange={(event) => setSchedule({ ...schedule, programId: event.target.value })}>
                <option value="">Pendiente</option>
                {programs.map((item) => (
                  <option key={item.id} value={item.id}>
                    {item.name}
                  </option>
                ))}
              </select>
            </label>
            <label>
              Competencia
              <input value={schedule.competence} onChange={(event) => setSchedule({ ...schedule, competence: event.target.value })} placeholder="Pendiente si no aplica" />
            </label>
            <label>
              RAP
              <input value={schedule.rap} onChange={(event) => setSchedule({ ...schedule, rap: event.target.value })} placeholder="Pendiente si no aplica" />
            </label>
            <label className="wide">
              Actividad de aprendizaje
              <input value={schedule.learningActivity} onChange={(event) => setSchedule({ ...schedule, learningActivity: event.target.value })} placeholder="Cuando aplique" />
            </label>
            <label className="wide">
              Observaciones o novedades
              <textarea value={schedule.observation} onChange={(event) => setSchedule({ ...schedule, observation: event.target.value })} rows={3} />
            </label>
          </div>
          <button className="primary-action" type="submit">
            <Save size={18} />
            <span>Guardar horario</span>
          </button>
        </form>
      </section>

      <section className="panel">
        <div className="section-toolbar">
          <PanelTitle icon={CalendarDays} title="Programacion registrada" />
          <label className="search-box">
            <Search size={16} />
            <input placeholder="Filtrar por ficha, instructor, ambiente..." value={filter} onChange={(event) => setFilter(event.target.value)} />
          </label>
        </div>
        <div className="table-wrap">
          <table>
            <thead>
              <tr>
                <th>Fecha</th>
                <th>Horario</th>
                <th>Ficha</th>
                <th>Instructor</th>
                <th>Ambiente</th>
                <th>Programa</th>
                <th>RAP</th>
                <th>Observacion</th>
              </tr>
            </thead>
            <tbody>
              {filteredSchedules.map((item) => (
                <tr key={item.id}>
                  <td>{item.date}</td>
                  <td>{item.startTime} - {item.endTime}</td>
                  <td>{item.fichaCode}</td>
                  <td>{item.instructorName}</td>
                  <td>{item.environmentName}</td>
                  <td>{item.programName || 'Pendiente'}</td>
                  <td>{item.rap || 'Pendiente'}</td>
                  <td>{item.observation || 'Sin novedad'}</td>
                </tr>
              ))}
              {filteredSchedules.length === 0 && (
                <tr>
                  <td colSpan="8" className="empty-state">Sin horarios registrados.</td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </section>

      <section className="panel no-programming">
        <PanelTitle icon={AlertTriangle} title="Dias sin programacion" />
        <form className="inline-form" onSubmit={submitNoProgrammingDay}>
          <input type="date" value={noProgrammingForm.date} onChange={(event) => setNoProgrammingForm({ ...noProgrammingForm, date: event.target.value })} required />
          <select value={noProgrammingForm.fichaId} onChange={(event) => setNoProgrammingForm({ ...noProgrammingForm, fichaId: event.target.value })}>
            <option value="">Ficha Pendiente</option>
            {fichas.map((item) => (
              <option key={item.id} value={item.id}>{item.code}</option>
            ))}
          </select>
          <select value={noProgrammingForm.instructorId} onChange={(event) => setNoProgrammingForm({ ...noProgrammingForm, instructorId: event.target.value })}>
            <option value="">Instructor Pendiente</option>
            {instructors.map((item) => (
              <option key={item.id} value={item.id}>{item.name}</option>
            ))}
          </select>
          <input placeholder="Motivo u observacion" value={noProgrammingForm.reason} onChange={(event) => setNoProgrammingForm({ ...noProgrammingForm, reason: event.target.value })} required />
          <button className="icon-button" type="submit" title="Registrar dia sin programacion">
            <Plus size={18} />
            <span>Registrar</span>
          </button>
        </form>
        <div className="chips">
          {noProgrammingDays.map((item) => (
            <span className="chip" key={item.id}>
              {item.date} · {item.fichaCode || item.instructorName || 'Pendiente'} · {item.reason}
            </span>
          ))}
          {noProgrammingDays.length === 0 && <span className="muted">No hay dias sin programacion registrados.</span>}
        </div>
      </section>
    </main>
  );
}

function Metric({ icon: Icon, label, value }) {
  return (
    <div className="metric">
      <Icon size={20} />
      <div>
        <strong>{value}</strong>
        <span>{label}</span>
      </div>
    </div>
  );
}

function PanelTitle({ icon: Icon, title }) {
  return (
    <div className="panel-title">
      <Icon size={18} />
      <h2>{title}</h2>
    </div>
  );
}

function CatalogBlock({ label, fields, onSubmit }) {
  return (
    <div className="catalog-block">
      <span>{label}</span>
      <div className="catalog-fields">{fields}</div>
      <button type="button" className="icon-only" onClick={onSubmit} title={`Crear ${label}`}>
        <Plus size={18} />
      </button>
    </div>
  );
}

async function apiGet(path) {
  const response = await fetchWithTimeout(`${API_URL}${path}`);
  return parseResponse(response);
}

async function apiPost(path, body) {
  const response = await fetchWithTimeout(`${API_URL}${path}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  });
  return parseResponse(response);
}

async function parseResponse(response) {
  const data = await response.json().catch(() => ({}));
  if (!response.ok) {
    if (data.conflicts?.length) {
      const details = data.conflicts.map((item) => `${item.kind} ${item.date} ${item.startTime}-${item.endTime}`).join(', ');
      throw new Error(`${data.message} Cruces: ${details}`);
    }
    throw new Error(data.message || 'No fue posible completar la operacion.');
  }
  return data;
}

function optionalNumber(value) {
  return value ? Number(value) : null;
}

async function fetchWithTimeout(url, options = {}) {
  const controller = new AbortController();
  const timeout = setTimeout(() => controller.abort(), 2500);
  try {
    return await fetch(url, { ...options, signal: controller.signal });
  } finally {
    clearTimeout(timeout);
  }
}

function loadDemoStore() {
  const raw = localStorage.getItem(DEMO_STORE_KEY);
  if (raw) {
    return JSON.parse(raw);
  }
  const store = {
    instructors: [
      { id: 1, name: 'Instructor Demo Planta 1', instructorType: 'planta' },
      { id: 2, name: 'Instructor Demo Planta 2', instructorType: 'planta' },
      { id: 3, name: 'Instructor Demo Contratista 1', instructorType: 'contratista' },
      { id: 4, name: 'Instructor Demo Contratista 2', instructorType: 'contratista' },
      { id: 5, name: 'Instructor Demo Contratista 3', instructorType: 'contratista' },
    ],
    programs: [
      { id: 1, name: 'Programa de formacion Pendiente Demo - Software', level: 'Tecnologo Pendiente' },
      { id: 2, name: 'Programa de formacion Pendiente Demo - Gestion', level: 'Tecnico Pendiente' },
      { id: 3, name: 'Programa de formacion Pendiente Demo - Redes', level: 'Tecnologo Pendiente' },
    ],
    fichas: [
      { id: 1, code: 'Ficha Demo 001 - Pendiente', programId: 1, programName: 'Programa de formacion Pendiente Demo - Software' },
      { id: 2, code: 'Ficha Demo 002 - Pendiente', programId: 2, programName: 'Programa de formacion Pendiente Demo - Gestion' },
      { id: 3, code: 'Ficha Demo 003 - Pendiente', programId: 3, programName: 'Programa de formacion Pendiente Demo - Redes' },
      { id: 4, code: 'Ficha Demo 004 - Pendiente', programId: 1, programName: 'Programa de formacion Pendiente Demo - Software' },
    ],
    environments: [
      { id: 1, name: 'Ambiente Demo 101 - Pendiente' },
      { id: 2, name: 'Ambiente Demo 202 - Pendiente' },
      { id: 3, name: 'Ambiente Demo Laboratorio - Pendiente' },
      { id: 4, name: 'Ambiente Demo Virtual - Pendiente' },
    ],
    schedules: [
      demoSchedule(1, '2026-05-18', '07:00', '09:00', 1, 1, 1, 1, 'RAP Pendiente Demo 1', 'Inicio de bloque formativo demo'),
      demoSchedule(2, '2026-05-18', '09:00', '11:00', 1, 2, 2, 2, 'RAP Pendiente Demo 2', 'Cambio de ambiente pendiente de confirmar'),
      demoSchedule(3, '2026-05-19', '07:00', '10:00', 1, 1, 3, 1, 'RAP Pendiente Demo 3', 'Practica demo'),
      demoSchedule(4, '2026-05-20', '13:00', '16:00', 1, 4, 4, 1, 'RAP Pendiente Demo 4', 'Sesion virtual demo'),
      demoSchedule(5, '2026-05-18', '07:00', '09:00', 2, 3, 3, 3, 'RAP Pendiente Demo 5', 'Sin novedad'),
      demoSchedule(6, '2026-05-19', '10:00', '12:00', 2, 2, 1, 2, 'RAP Pendiente Demo 6', 'Observacion demo'),
      demoSchedule(7, '2026-05-21', '08:00', '11:00', 3, 1, 2, 1, 'RAP Pendiente Demo 7', 'Pendiente validar RAP'),
      demoSchedule(8, '2026-05-22', '14:00', '17:00', 3, 3, 4, 3, 'RAP Pendiente Demo 8', 'Trabajo remoto demo'),
      demoSchedule(9, '2026-05-18', '11:00', '13:00', 4, 4, 3, 1, 'RAP Pendiente Demo 9', 'Sin novedad'),
      demoSchedule(10, '2026-05-20', '07:00', '09:00', 4, 2, 1, 2, 'RAP Pendiente Demo 10', 'Novedad pendiente'),
      demoSchedule(11, '2026-05-21', '13:00', '15:00', 5, 3, 2, 3, 'RAP Pendiente Demo 11', 'Sin novedad'),
      demoSchedule(12, '2026-05-22', '09:00', '12:00', 5, 4, 1, 1, 'RAP Pendiente Demo 12', 'Cierre semanal demo'),
    ],
    noProgrammingDays: [
      { id: 1, date: '2026-05-20', fichaId: 3, fichaCode: 'Ficha Demo 003 - Pendiente', reason: 'Dia sin programacion demo por ajuste de planeacion' },
      { id: 2, date: '2026-05-22', instructorId: 2, instructorName: 'Instructor Demo Planta 2', reason: 'Disponibilidad de instructor pendiente de confirmar' },
    ],
  };
  saveDemoStore(store);
  return store;
}

function saveDemoStore(store) {
  localStorage.setItem(DEMO_STORE_KEY, JSON.stringify(store));
}

function nextID(items) {
  return items.reduce((max, item) => Math.max(max, Number(item.id) || 0), 0) + 1;
}

function findByID(items, id) {
  const numericID = Number(id);
  return items.find((item) => Number(item.id) === numericID);
}

function requireText(value, label) {
  if (!String(value || '').trim()) {
    throw new Error(`${label} es requerido.`);
  }
}

function validateDemoSchedule(input) {
  requireText(input.date, 'Fecha');
  requireText(input.startTime, 'Inicio');
  requireText(input.endTime, 'Fin');
  if (!input.instructorId || !input.fichaId || !input.environmentId) {
    throw new Error('Instructor, ficha y ambiente son requeridos.');
  }
  if (input.endTime <= input.startTime) {
    throw new Error('La hora fin debe ser posterior a la hora inicio.');
  }
}

function findDemoConflicts(existingSchedules, input) {
  return existingSchedules
    .filter((item) => item.date === input.date && item.startTime < input.endTime && item.endTime > input.startTime)
    .flatMap((item) => {
      const conflicts = [];
      if (Number(item.instructorId) === Number(input.instructorId)) conflicts.push('instructor');
      if (Number(item.fichaId) === Number(input.fichaId)) conflicts.push('ficha');
      if (Number(item.environmentId) === Number(input.environmentId)) conflicts.push('ambiente');
      return conflicts;
    });
}

function demoSchedule(id, date, startTime, endTime, instructorId, fichaId, environmentId, programId, rap, observation) {
  const instructors = {
    1: ['Instructor Demo Planta 1', 'planta'],
    2: ['Instructor Demo Planta 2', 'planta'],
    3: ['Instructor Demo Contratista 1', 'contratista'],
    4: ['Instructor Demo Contratista 2', 'contratista'],
    5: ['Instructor Demo Contratista 3', 'contratista'],
  };
  const fichas = {
    1: 'Ficha Demo 001 - Pendiente',
    2: 'Ficha Demo 002 - Pendiente',
    3: 'Ficha Demo 003 - Pendiente',
    4: 'Ficha Demo 004 - Pendiente',
  };
  const environments = {
    1: 'Ambiente Demo 101 - Pendiente',
    2: 'Ambiente Demo 202 - Pendiente',
    3: 'Ambiente Demo Laboratorio - Pendiente',
    4: 'Ambiente Demo Virtual - Pendiente',
  };
  const programs = {
    1: ['Programa de formacion Pendiente Demo - Software', 'Tecnologo Pendiente'],
    2: ['Programa de formacion Pendiente Demo - Gestion', 'Tecnico Pendiente'],
    3: ['Programa de formacion Pendiente Demo - Redes', 'Tecnologo Pendiente'],
  };

  return {
    id,
    date,
    startTime,
    endTime,
    instructorId,
    instructorName: instructors[instructorId][0],
    instructorType: instructors[instructorId][1],
    fichaId,
    fichaCode: fichas[fichaId],
    environmentId,
    environmentName: environments[environmentId],
    programId,
    programName: programs[programId][0],
    programLevel: programs[programId][1],
    competence: `Competencia Pendiente Demo ${id}`,
    rap,
    learningActivity: `Actividad de aprendizaje Pendiente Demo ${id}`,
    observation,
  };
}

function formatDisplayDate(value) {
  if (!value) return 'Fecha Pendiente';
  const date = new Date(`${value}T00:00:00`);
  return new Intl.DateTimeFormat('es-CO', {
    weekday: 'short',
    day: '2-digit',
    month: 'short',
  }).format(date);
}

export default App;
