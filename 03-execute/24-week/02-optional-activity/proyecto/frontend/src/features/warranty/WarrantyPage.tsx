/* Warranty screen: the whole list with the badge that says whether each one is
   valid at the date the user consults. */
import { useState } from 'react';

import { listWarranty } from '../../services/warranty_service';
import { DataState } from '../../shared/DataState';
import { ValidityBadge } from '../../shared/StatusBadge';
import { useAsyncData } from '../../shared/useAsyncData';
import { useToken } from '../../shared/SessionContext';
import { formatDate, warrantyLabel } from '../../shared/format';

export function WarrantyPage() {
  const token = useToken();
  const [consultedAt, setConsultedAt] = useState('');
  const warranty = useAsyncData(() => listWarranty(token, consultedAt), [token, consultedAt]);

  return (
    <section>
      <h2 className="screen-title">Garantias</h2>
      <section className="card">
        <div className="field">
          <label className="field__label" htmlFor="consultedAt">
            Consultar a la fecha
          </label>
          <input
            className="field__input"
            id="consultedAt"
            type="date"
            value={consultedAt}
            onChange={(event) => setConsultedAt(event.target.value)}
          />
        </div>
        <DataState
          loading={warranty.loading}
          error={warranty.error}
          empty={(warranty.data ?? []).length === 0}
          emptyMessage="No hay garantias emitidas."
        >
          <div className="table-scroll">
            <table className="data-table">
              <thead>
                <tr>
                  <th scope="col">Orden</th>
                  <th scope="col">Placa</th>
                  <th scope="col">Tipo</th>
                  <th scope="col">Cobertura</th>
                  <th scope="col">Emitida</th>
                  <th scope="col">Vence</th>
                  <th scope="col">Estado</th>
                </tr>
              </thead>
              <tbody>
                {(warranty.data ?? []).map((item) => (
                  <tr key={item.id}>
                    <td>{item.orderNumber}</td>
                    <td>{item.vehiclePlate}</td>
                    <td>{warrantyLabel(item.kind)}</td>
                    <td>{item.coverageMonthCount + ' meses'}</td>
                    <td>{formatDate(item.issuedAt)}</td>
                    <td>{formatDate(item.expirationDate)}</td>
                    <td>
                      <ValidityBadge valid={item.valid} />
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
