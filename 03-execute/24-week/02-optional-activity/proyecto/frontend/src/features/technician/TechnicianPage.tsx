/* Technician screen: who is available and who is holding a vehicle. */
import { DataState } from '../../shared/DataState';
import { AvailabilityBadge } from '../../shared/StatusBadge';
import { useAsyncData } from '../../shared/useAsyncData';
import { useToken } from '../../shared/SessionContext';
import { listTechnician } from '../../services/technician_service';

export function TechnicianPage() {
  const token = useToken();
  const { data, loading, error } = useAsyncData(() => listTechnician(token), [token]);

  return (
    <section>
      <h2 className="screen-title">Tecnicos</h2>
      <section className="card">
        <DataState
          loading={loading}
          error={error}
          empty={(data ?? []).length === 0}
          emptyMessage="No hay tecnicos registrados."
        >
          <div className="table-scroll">
            <table className="data-table">
              <thead>
                <tr>
                  <th scope="col">Tecnico</th>
                  <th scope="col">Especialidad</th>
                  <th scope="col">Estado</th>
                  <th scope="col">Orden activa</th>
                </tr>
              </thead>
              <tbody>
                {(data ?? []).map((technician) => (
                  <tr key={technician.id}>
                    <td>{technician.fullName}</td>
                    <td>{technician.specialty}</td>
                    <td>
                      <AvailabilityBadge busy={technician.busy} />
                    </td>
                    <td>{technician.activeOrderNumber || 'Sin orden'}</td>
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
