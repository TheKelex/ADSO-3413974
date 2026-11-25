/* Clinical timeline: everything that happened to one vehicle, oldest first. */
import { useParams } from 'react-router-dom';

import { readTimeline } from '../../services/timeline_service';
import { DataState } from '../../shared/DataState';
import { useAsyncData } from '../../shared/useAsyncData';
import { useToken } from '../../shared/SessionContext';
import { formatDateTime, timelineLabel } from '../../shared/format';

export function VehicleTimelinePage() {
  const { vehicleId = '' } = useParams();
  const token = useToken();
  const timeline = useAsyncData(() => readTimeline(token, vehicleId), [token, vehicleId]);
  const vehicle = timeline.data?.vehicle;

  return (
    <section>
      <h2 className="screen-title">Historial del vehiculo</h2>
      <DataState loading={timeline.loading} error={timeline.error} empty={!timeline.data}>
        <section className="card">
          <h3 className="card__title">{vehicle?.plate}</h3>
          <p>
            {vehicle?.brand} {vehicle?.model} {vehicle?.modelYear}
          </p>
          <p className="timeline__date">Propietario: {vehicle?.ownerName}</p>
        </section>
        <section className="card">
          <h3 className="card__title">Cronologia</h3>
          {(timeline.data?.entry ?? []).length === 0 ? (
            <p className="state-message">Este vehiculo aun no tiene historial.</p>
          ) : (
            <ol className="timeline">
              {(timeline.data?.entry ?? []).map((item, index) => (
                <li className="timeline__item" key={item.reference + index}>
                  <span className="badge badge--diagnosis">{timelineLabel(item.kind)}</span>
                  <p className="timeline__date">{formatDateTime(item.occurredAt)}</p>
                  <p>
                    <strong>{item.title}</strong>
                  </p>
                  {item.description ? <p>{item.description}</p> : null}
                </li>
              ))}
            </ol>
          )}
        </section>
      </DataState>
    </section>
  );
}
