/* Service order screen: the check-in form, the status filter and the list of
   orders with the badge of their lifecycle status. */
import { useState } from 'react';
import { Link } from 'react-router-dom';

import { ApiError } from '../../services/api_client';
import { listVehicle } from '../../services/vehicle_service';
import { listServiceOrder, openServiceOrder } from '../../services/service_order_service';
import { DataState, ErrorBanner, SuccessBanner } from '../../shared/DataState';
import { StatusBadge } from '../../shared/StatusBadge';
import { useAsyncData } from '../../shared/useAsyncData';
import { useSession, useToken } from '../../shared/SessionContext';
import { STATUS_ORDER, formatDateTime, statusLabel } from '../../shared/format';

export function ServiceOrderPage() {
  const token = useToken();
  const { isAdministrator } = useSession();
  const [status, setStatus] = useState('');
  const order = useAsyncData(() => listServiceOrder(token, status), [token, status]);
  const vehicle = useAsyncData(
    () => (isAdministrator ? listVehicle(token) : Promise.resolve([])),
    [token, isAdministrator],
  );
  const [vehicleId, setVehicleId] = useState('');
  const [reportedFailure, setReportedFailure] = useState('');
  const [formError, setFormError] = useState('');
  const [confirmation, setConfirmation] = useState('');
  const [sending, setSending] = useState(false);

  const submit = async (event: React.FormEvent) => {
    event.preventDefault();
    setFormError('');
    setConfirmation('');
    setSending(true);
    try {
      await openServiceOrder(token, vehicleId, reportedFailure);
      setVehicleId('');
      setReportedFailure('');
      setConfirmation('Ingreso registrado.');
      order.reload();
    } catch (failure) {
      setFormError(
        failure instanceof ApiError ? failure.message : 'No se pudo registrar el ingreso.',
      );
    } finally {
      setSending(false);
    }
  };

  return (
    <section>
      <h2 className="screen-title">Ordenes de servicio</h2>

      {isAdministrator ? (
        <section className="card">
          <h3 className="card__title">Nueva orden</h3>
          <form onSubmit={submit} noValidate>
            <ErrorBanner message={formError} />
            <SuccessBanner message={confirmation} />
            <div className="form-grid">
              <div className="field">
                <label className="field__label" htmlFor="vehicleId">
                  Vehiculo
                </label>
                <select
                  className="field__input"
                  id="vehicleId"
                  value={vehicleId}
                  onChange={(event) => setVehicleId(event.target.value)}
                  required
                >
                  <option value="">Seleccione un vehiculo</option>
                  {(vehicle.data ?? []).map((item) => (
                    <option key={item.id} value={item.id}>
                      {item.plate + ' - ' + item.brand + ' ' + item.model}
                    </option>
                  ))}
                </select>
              </div>
              <div className="field">
                <label className="field__label" htmlFor="reportedFailure">
                  Falla reportada
                </label>
                <input
                  className="field__input"
                  id="reportedFailure"
                  value={reportedFailure}
                  onChange={(event) => setReportedFailure(event.target.value)}
                  required
                />
              </div>
            </div>
            <button type="submit" className="button button--primary" disabled={sending}>
              {sending ? 'Registrando...' : 'Registrar ingreso'}
            </button>
          </form>
        </section>
      ) : null}

      <section className="card">
        <div className="field">
          <label className="field__label" htmlFor="statusFilter">
            Filtrar por estado
          </label>
          <select
            className="field__input"
            id="statusFilter"
            value={status}
            onChange={(event) => setStatus(event.target.value)}
          >
            <option value="">Todos los estados</option>
            {STATUS_ORDER.map((item) => (
              <option key={item} value={item}>
                {statusLabel(item)}
              </option>
            ))}
          </select>
        </div>
        <DataState
          loading={order.loading}
          error={order.error}
          empty={(order.data ?? []).length === 0}
          emptyMessage="No hay ordenes con ese estado."
        >
          <div className="table-scroll">
            <table className="data-table">
              <thead>
                <tr>
                  <th scope="col">Orden</th>
                  <th scope="col">Placa</th>
                  <th scope="col">Estado</th>
                  <th scope="col">Tecnico</th>
                  <th scope="col">Ingreso</th>
                  <th scope="col">Detalle</th>
                </tr>
              </thead>
              <tbody>
                {(order.data ?? []).map((item) => (
                  <tr key={item.id}>
                    <td>{item.orderNumber}</td>
                    <td>{item.vehiclePlate}</td>
                    <td>
                      <StatusBadge status={item.status} />
                    </td>
                    <td>{item.technicianName || 'Sin asignar'}</td>
                    <td>{formatDateTime(item.receivedAt)}</td>
                    <td>
                      <Link className="button button--secondary" to={'/service-orders/' + item.id}>
                        Abrir
                      </Link>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </DataState>
      </section>
    </section>
  );
}
