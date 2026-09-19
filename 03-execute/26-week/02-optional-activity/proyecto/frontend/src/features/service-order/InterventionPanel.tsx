/* Intervention panel: the work executed on the vehicle, with a repeatable part
   row, plus the action that issues a warranty over a recorded intervention. */
import { useState } from 'react';

import { ApiError } from '../../services/api_client';
import { listIntervention, registerIntervention } from '../../services/service_order_service';
import type { PartUsage } from '../../services/service_order_service';
import { issueWarranty } from '../../services/warranty_service';
import type { WarrantyKind } from '../../services/warranty_service';
import { DataState, ErrorBanner, SuccessBanner } from '../../shared/DataState';
import { useAsyncData } from '../../shared/useAsyncData';
import { useSession, useToken } from '../../shared/SessionContext';
import { formatDateTime } from '../../shared/format';

interface InterventionPanelProps {
  serviceOrderId: string;
  delivered: boolean;
  onChange: () => void;
}

const EMPTY_PART: PartUsage = { partName: '', quantity: 1 };

export function InterventionPanel({ serviceOrderId, delivered, onChange }: InterventionPanelProps) {
  const token = useToken();
  const { isAdministrator } = useSession();
  const intervention = useAsyncData(
    () => listIntervention(token, serviceOrderId),
    [token, serviceOrderId],
  );
  const [description, setDescription] = useState('');
  const [laborHourCount, setLaborHourCount] = useState('1');
  const [part, setPart] = useState<PartUsage[]>([EMPTY_PART]);
  const [error, setError] = useState('');
  const [confirmation, setConfirmation] = useState('');
  const [sending, setSending] = useState(false);

  const updatePart = (index: number, field: keyof PartUsage, value: string) =>
    setPart((previous) =>
      previous.map((item, position) =>
        position === index
          ? { ...item, [field]: field === 'quantity' ? Number(value) : value }
          : item,
      ),
    );

  const submit = async (event: React.FormEvent) => {
    event.preventDefault();
    setError('');
    setConfirmation('');
    setSending(true);
    try {
      await registerIntervention(
        token,
        serviceOrderId,
        description,
        Number(laborHourCount),
        part.filter((item) => item.partName.trim().length > 0),
      );
      setDescription('');
      setLaborHourCount('1');
      setPart([EMPTY_PART]);
      setConfirmation('Intervencion registrada.');
      intervention.reload();
      onChange();
    } catch (failure) {
      setError(
        failure instanceof ApiError ? failure.message : 'No se pudo registrar la intervencion.',
      );
    } finally {
      setSending(false);
    }
  };

  const issue = async (interventionId: string, kind: WarrantyKind) => {
    setError('');
    setConfirmation('');
    try {
      await issueWarranty(token, interventionId, kind, 12);
      setConfirmation('Garantia emitida por 12 meses.');
      intervention.reload();
    } catch (failure) {
      setError(failure instanceof ApiError ? failure.message : 'No se pudo emitir la garantia.');
    }
  };

  return (
    <section className="card">
      <h3 className="card__title">Intervenciones</h3>
      <ErrorBanner message={error} />
      <SuccessBanner message={confirmation} />

      {delivered ? (
        <p className="state-message">La orden ya fue entregada.</p>
      ) : isAdministrator ? null : (
        <form onSubmit={submit} noValidate>
          <div className="form-grid">
            <div className="field">
              <label className="field__label" htmlFor="description">
                Descripcion
              </label>
              <input
                className="field__input"
                id="description"
                value={description}
                onChange={(event) => setDescription(event.target.value)}
                required
              />
            </div>
            <div className="field">
              <label className="field__label" htmlFor="laborHourCount">
                Horas de trabajo
              </label>
              <input
                className="field__input"
                id="laborHourCount"
                type="number"
                min="0.5"
                step="0.5"
                value={laborHourCount}
                onChange={(event) => setLaborHourCount(event.target.value)}
                required
              />
            </div>
          </div>
          {part.map((item, index) => (
            <div className="form-grid" key={index}>
              <div className="field">
                <label className="field__label" htmlFor={'partName-' + index}>
                  Repuesto
                </label>
                <input
                  className="field__input"
                  id={'partName-' + index}
                  value={item.partName}
                  onChange={(event) => updatePart(index, 'partName', event.target.value)}
                />
              </div>
              <div className="field">
                <label className="field__label" htmlFor={'quantity-' + index}>
                  Cantidad
                </label>
                <input
                  className="field__input"
                  id={'quantity-' + index}
                  type="number"
                  min="1"
                  value={String(item.quantity)}
                  onChange={(event) => updatePart(index, 'quantity', event.target.value)}
                />
              </div>
            </div>
          ))}
          <button
            type="button"
            className="button button--secondary"
            onClick={() => setPart((previous) => [...previous, { ...EMPTY_PART }])}
          >
            Agregar repuesto
          </button>
          <button type="submit" className="button button--primary" disabled={sending}>
            {sending ? 'Registrando...' : 'Registrar intervencion'}
          </button>
        </form>
      )}

      <DataState
        loading={intervention.loading}
        error={intervention.error}
        empty={(intervention.data ?? []).length === 0}
        emptyMessage="Esta orden aun no tiene intervenciones."
      >
        <div className="table-scroll">
          <table className="data-table">
            <thead>
              <tr>
                <th scope="col">Descripcion</th>
                <th scope="col">Horas</th>
                <th scope="col">Repuestos</th>
                <th scope="col">Fecha</th>
                <th scope="col">Garantia</th>
              </tr>
            </thead>
            <tbody>
              {(intervention.data ?? []).map((item) => (
                <tr key={item.id}>
                  <td>{item.description}</td>
                  <td>{item.laborHourCount}</td>
                  <td>
                    {item.part.length === 0
                      ? 'Sin repuestos'
                      : item.part.map((used) => used.partName).join(', ')}
                  </td>
                  <td>{formatDateTime(item.performedAt)}</td>
                  <td>
                    {item.warrantyId ? (
                      <span>{item.valid ? 'Garantia vigente' : 'Garantia vencida'}</span>
                    ) : isAdministrator && !delivered ? (
                      <button
                        type="button"
                        className="button button--secondary"
                        onClick={() => issue(item.id, 'LABOR')}
                      >
                        Emitir garantia
                      </button>
                    ) : null}
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
