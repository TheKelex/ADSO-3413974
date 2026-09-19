/* Allocation panel: pick an available technician for this order. A busy
   technician is shown with its badge and the backend rejection is surfaced as
   the Spanish message it returns. */
import { useState } from 'react';

import { ApiError } from '../../services/api_client';
import { assignTechnician } from '../../services/service_order_service';
import { listTechnician } from '../../services/technician_service';
import { DataState, ErrorBanner, SuccessBanner } from '../../shared/DataState';
import { AvailabilityBadge } from '../../shared/StatusBadge';
import { useAsyncData } from '../../shared/useAsyncData';
import { useSession, useToken } from '../../shared/SessionContext';

interface AssignmentPanelProps {
  serviceOrderId: string;
  technicianName: string;
  onChange: () => void;
}

export function AssignmentPanel({ serviceOrderId, technicianName, onChange }: AssignmentPanelProps) {
  const token = useToken();
  const { isAdministrator } = useSession();
  const technician = useAsyncData(
    () => (isAdministrator ? listTechnician(token) : Promise.resolve([])),
    [token, isAdministrator],
  );
  const [technicianId, setTechnicianId] = useState('');
  const [error, setError] = useState('');
  const [confirmation, setConfirmation] = useState('');
  const [sending, setSending] = useState(false);

  const submit = async (event: React.FormEvent) => {
    event.preventDefault();
    setError('');
    setConfirmation('');
    setSending(true);
    try {
      await assignTechnician(token, serviceOrderId, technicianId);
      setConfirmation('Tecnico asignado.');
      technician.reload();
      onChange();
    } catch (failure) {
      setError(failure instanceof ApiError ? failure.message : 'No se pudo asignar el tecnico.');
    } finally {
      setSending(false);
    }
  };

  if (!isAdministrator) {
    return (
      <section className="card">
        <h3 className="card__title">Asignacion</h3>
        <p className="state-message">Técnico asignado: {technicianName || 'Sin asignar'}.</p>
      </section>
    );
  }

  return (
    <section className="card">
      <h3 className="card__title">Asignacion</h3>
      <DataState
        loading={technician.loading}
        error={technician.error}
        empty={(technician.data ?? []).length === 0}
        emptyMessage="No hay tecnicos registrados."
      >
        <form onSubmit={submit} noValidate>
          <ErrorBanner message={error} />
          <SuccessBanner message={confirmation} />
          <div className="field">
            <label className="field__label" htmlFor="technicianId">
              Tecnico
            </label>
            <select
              className="field__input"
              id="technicianId"
              value={technicianId}
              onChange={(event) => setTechnicianId(event.target.value)}
              required
            >
              <option value="">Seleccione un tecnico</option>
              {(technician.data ?? []).map((item) => (
                <option key={item.id} value={item.id}>
                  {item.fullName + (item.busy ? ' (ocupado)' : '')}
                </option>
              ))}
            </select>
          </div>
          <button type="submit" className="button button--primary" disabled={sending}>
            {sending ? 'Asignando...' : 'Asignar tecnico'}
          </button>
        </form>
        <div className="table-scroll">
          <table className="data-table">
            <thead>
              <tr>
                <th scope="col">Tecnico</th>
                <th scope="col">Especialidad</th>
                <th scope="col">Estado</th>
              </tr>
            </thead>
            <tbody>
              {(technician.data ?? []).map((item) => (
                <tr key={item.id}>
                  <td>{item.fullName}</td>
                  <td>{item.specialty}</td>
                  <td>
                    <AvailabilityBadge busy={item.busy} />
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </DataState>
    </section>
  );
}
