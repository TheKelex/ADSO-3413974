/*
 * Presentation helpers. Every label the user reads is written in Spanish here,
 * so a screen never spells a status or a date format on its own.
 */
import type { ServiceOrderStatus } from '../services/service_order_service';
import type { TimelineKind } from '../services/timeline_service';
import type { WarrantyKind } from '../services/warranty_service';

const STATUS_LABEL: Record<ServiceOrderStatus, string> = {
  RECEIVED: 'Recibido',
  IN_DIAGNOSIS: 'En diagnostico',
  IN_REPAIR: 'En reparacion',
  READY: 'Listo',
  DELIVERED: 'Entregado',
};

const STATUS_MODIFIER: Record<ServiceOrderStatus, string> = {
  RECEIVED: 'received',
  IN_DIAGNOSIS: 'diagnosis',
  IN_REPAIR: 'repair',
  READY: 'ready',
  DELIVERED: 'delivered',
};

const TIMELINE_LABEL: Record<TimelineKind, string> = {
  ORDER: 'Orden',
  DIAGNOSTIC: 'Diagnostico',
  INTERVENTION: 'Intervencion',
  WARRANTY: 'Garantia',
};

const WARRANTY_LABEL: Record<WarrantyKind, string> = {
  LABOR: 'Mano de obra',
  PART: 'Repuesto',
};

export const STATUS_ORDER: ServiceOrderStatus[] = [
  'RECEIVED',
  'IN_DIAGNOSIS',
  'IN_REPAIR',
  'READY',
  'DELIVERED',
];

/** The status that follows the current one, or null when the order is closed. */
export function nextStatus(status: ServiceOrderStatus): ServiceOrderStatus | null {
  const index = STATUS_ORDER.indexOf(status);
  if (index < 0 || index >= STATUS_ORDER.length - 1) {
    return null;
  }
  return STATUS_ORDER[index + 1];
}

export function statusLabel(status: ServiceOrderStatus): string {
  return STATUS_LABEL[status] ?? status;
}

export function statusModifier(status: ServiceOrderStatus): string {
  return STATUS_MODIFIER[status] ?? 'received';
}

export function timelineLabel(kind: TimelineKind): string {
  return TIMELINE_LABEL[kind] ?? kind;
}

export function warrantyLabel(kind: WarrantyKind): string {
  return WARRANTY_LABEL[kind] ?? kind;
}

/** Render an instant as a readable Spanish date and time. */
export function formatDateTime(value: string): string {
  if (!value) {
    return '';
  }
  const moment = new Date(value);
  if (Number.isNaN(moment.getTime())) {
    return value;
  }
  return moment.toLocaleString('es', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  });
}

/** Render an instant as a readable Spanish date. */
export function formatDate(value: string): string {
  if (!value) {
    return '';
  }
  const moment = new Date(value);
  if (Number.isNaN(moment.getTime())) {
    return value;
  }
  return moment.toLocaleDateString('es', { day: '2-digit', month: '2-digit', year: 'numeric' });
}
