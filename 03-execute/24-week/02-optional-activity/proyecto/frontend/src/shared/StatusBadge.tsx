/* Status badge. The colour comes from the token palette and the meaning is
   always written as text, never carried by colour alone. */
import { statusLabel, statusModifier } from './format';
import type { ServiceOrderStatus } from '../services/service_order_service';

export function StatusBadge({ status }: { status: ServiceOrderStatus }) {
  return (
    <span className={'badge badge--' + statusModifier(status)}>{statusLabel(status)}</span>
  );
}

/** Badge for a warranty that is still valid or already expired. */
export function ValidityBadge({ valid }: { valid: boolean }) {
  return (
    <span className={'badge badge--' + (valid ? 'ready' : 'delivered')}>
      {valid ? 'Vigente' : 'Vencida'}
    </span>
  );
}

/** Badge that tells whether a technician can receive a new order. */
export function AvailabilityBadge({ busy }: { busy: boolean }) {
  return (
    <span className={'badge badge--' + (busy ? 'repair' : 'ready')}>
      {busy ? 'Ocupado' : 'Disponible'}
    </span>
  );
}
