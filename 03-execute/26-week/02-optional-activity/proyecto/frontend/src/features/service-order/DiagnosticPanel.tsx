/* Diagnostic panel. Only the assigned technician can write here; the backend
   is the authority and its refusal is shown as it comes. */
import { useState } from 'react';

import { ApiError } from '../../services/api_client';
import { findDiagnostic, recordDiagnostic } from '../../services/service_order_service';
import { ErrorBanner, SuccessBanner } from '../../shared/DataState';
import { useAsyncData } from '../../shared/useAsyncData';
import { useSession, useToken } from '../../shared/SessionContext';
import { formatDateTime } from '../../shared/format';

interface DiagnosticPanelProps {
  serviceOrderId: string;
  delivered: boolean;
  onChange: () => void;
}

export function DiagnosticPanel({ serviceOrderId, delivered, onChange }: DiagnosticPanelProps) {
  const token = useToken();
  const { isAdministrator } = useSession();
  const diagnostic = useAsyncData(
    () =>
      findDiagnostic(token, serviceOrderId).catch((failure: unknown) => {
        if (failure instanceof ApiError && failure.status === 404) {
          return null;
        }
        throw failure;
      }),
    [token, serviceOrderId],
  );
  const [finding, setFinding] = useState('');
  const [componentToRepair, setComponentToRepair] = useState('');
  const [error, setError] = useState('');
  const [confirmation, setConfirmation] = useState('');
  const [sending, setSending] = useState(false);

  const submit = async (event: React.FormEvent) => {
    event.preventDefault();
    setError('');
    setConfirmation('');
    setSending(true);
    try {
      await recordDiagnostic(token, serviceOrderId, finding, componentToRepair);
      setFinding('');
      setComponentToRepair('');
      setConfirmation('Diagnostico guardado.');
      diagnostic.reload();
      onChange();
    } catch (failure) {
      setError(
        failure instanceof ApiError ? failure.message : 'No se pudo guardar el diagnostico.',
      );
    } finally {
      setSending(false);
    }
  };

  return (
    <section className="card">
      <h3 className="card__title">Diagnostico</h3>
      {diagnostic.loading ? (
        <p className="state-message" role="status" aria-live="polite">
          Cargando...
        </p>
      ) : null}
      {diagnostic.error ? <ErrorBanner message={diagnostic.error} /> : null}
      {!diagnostic.loading && !diagnostic.error && !diagnostic.data ? (
        <p className="state-message">Esta orden aun no tiene diagnostico.</p>
      ) : null}
      {diagnostic.data ? (
        <div>
          <p>
            <strong>Hallazgo:</strong> {diagnostic.data.finding}
          </p>
          <p>
            <strong>Componentes a reparar:</strong> {diagnostic.data.componentToRepair}
          </p>
          <p className="timeline__date">{formatDateTime(diagnostic.data.createdAt)}</p>
        </div>
      ) : null}

      {delivered ? (
        <p className="state-message">La orden ya fue entregada.</p>
      ) : isAdministrator ? (
        <p className="state-message">Solo el tecnico asignado puede escribir el diagnostico.</p>
      ) : (
        <form onSubmit={submit} noValidate>
          <ErrorBanner message={error} />
          <SuccessBanner message={confirmation} />
          <div className="field">
            <label className="field__label" htmlFor="finding">
              Hallazgo
            </label>
            <input
              className="field__input"
              id="finding"
              value={finding}
              onChange={(event) => setFinding(event.target.value)}
              required
            />
          </div>
          <div className="field">
            <label className="field__label" htmlFor="componentToRepair">
              Componentes a reparar
            </label>
            <input
              className="field__input"
              id="componentToRepair"
              value={componentToRepair}
              onChange={(event) => setComponentToRepair(event.target.value)}
              required
            />
          </div>
          <button type="submit" className="button button--primary" disabled={sending}>
            {sending ? 'Guardando...' : 'Guardar diagnostico'}
          </button>
        </form>
      )}
    </section>
  );
}
