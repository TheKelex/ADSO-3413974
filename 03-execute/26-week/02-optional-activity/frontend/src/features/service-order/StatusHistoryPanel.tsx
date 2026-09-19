/* Status history panel: every accepted transition with its author and its
   timestamp, which is what makes the repair auditable. */
import { listTransition } from '../../services/service_order_service';
import { DataState } from '../../shared/DataState';
import { StatusBadge } from '../../shared/StatusBadge';
import { useAsyncData } from '../../shared/useAsyncData';
import { useToken } from '../../shared/SessionContext';
import { formatDateTime } from '../../shared/format';

interface StatusHistoryPanelProps {
  serviceOrderId: string;
  refreshToken: number;
}

export function StatusHistoryPanel({ serviceOrderId, refreshToken }: StatusHistoryPanelProps) {
  const token = useToken();
  const history = useAsyncData(
    () => listTransition(token, serviceOrderId),
    [token, serviceOrderId, refreshToken],
  );

  return (
    <section className="card">
      <h3 className="card__title">Historial de estados</h3>
      <DataState
        loading={history.loading}
        error={history.error}
        empty={(history.data ?? []).length === 0}
        emptyMessage="La orden aun no ha cambiado de estado."
      >
        <div className="table-scroll">
          <table className="data-table">
            <thead>
              <tr>
                <th scope="col">Desde</th>
                <th scope="col">Hasta</th>
                <th scope="col">Responsable</th>
                <th scope="col">Fecha</th>
              </tr>
            </thead>
            <tbody>
              {(history.data ?? []).map((item) => (
                <tr key={item.id}>
                  <td>
                    <StatusBadge status={item.fromStatus} />
                  </td>
                  <td>
                    <StatusBadge status={item.toStatus} />
                  </td>
                  <td>{item.changedByName}</td>
                  <td>{formatDateTime(item.changedAt)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </DataState>
    </section>
  );
}
