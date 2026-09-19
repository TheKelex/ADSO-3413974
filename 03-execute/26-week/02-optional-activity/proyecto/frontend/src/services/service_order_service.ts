/* Service order service: check-in, listing, detail, lifecycle and the panels
   that hang from an order (assignment, diagnostic, intervention, history). */
import { request } from './api_client';

export type ServiceOrderStatus =
  | 'RECEIVED'
  | 'IN_DIAGNOSIS'
  | 'IN_REPAIR'
  | 'READY'
  | 'DELIVERED';

export interface ServiceOrder {
  id: string;
  orderNumber: string;
  vehicleId: string;
  vehiclePlate: string;
  technicianName: string;
  reportedFailure: string;
  status: ServiceOrderStatus;
  receivedAt: string;
  updatedAt: string;
}

export interface StatusTransition {
  id: string;
  fromStatus: ServiceOrderStatus;
  toStatus: ServiceOrderStatus;
  changedByName: string;
  changedAt: string;
}

export interface Assignment {
  id: string;
  serviceOrderId: string;
  technicianId: string;
  isActive: boolean;
  assignedAt: string;
}

export interface Diagnostic {
  id: string;
  serviceOrderId: string;
  technicianId: string;
  finding: string;
  componentToRepair: string;
  createdAt: string;
}

export interface PartUsage {
  partName: string;
  quantity: number;
}

export interface Intervention {
  id: string;
  serviceOrderId: string;
  technicianId: string;
  description: string;
  laborHourCount: number;
  performedAt: string;
  part: PartUsage[];
  warrantyId: string | null;
  kind: 'LABOR' | 'PART' | null;
  valid: boolean | null;
}

export function listServiceOrder(token: string, status: string): Promise<ServiceOrder[]> {
  const query = status ? '?status=' + encodeURIComponent(status) : '';
  return request<ServiceOrder[]>('/service-order' + query, { token });
}

export function findServiceOrder(token: string, serviceOrderId: string): Promise<ServiceOrder> {
  return request<ServiceOrder>('/service-order/' + serviceOrderId, { token });
}

export function openServiceOrder(
  token: string,
  vehicleId: string,
  reportedFailure: string,
): Promise<ServiceOrder> {
  return request<ServiceOrder>('/service-order', {
    method: 'POST',
    body: { vehicleId, reportedFailure },
    token,
  });
}

export function advanceServiceOrder(
  token: string,
  serviceOrderId: string,
  status: ServiceOrderStatus,
): Promise<ServiceOrder> {
  return request<ServiceOrder>('/service-order/' + serviceOrderId + '/status', {
    method: 'POST',
    body: { status },
    token,
  });
}

export function listTransition(token: string, serviceOrderId: string): Promise<StatusTransition[]> {
  return request<StatusTransition[]>('/service-order/' + serviceOrderId + '/transition', { token });
}

export function assignTechnician(
  token: string,
  serviceOrderId: string,
  technicianId: string,
): Promise<Assignment> {
  return request<Assignment>('/service-order/' + serviceOrderId + '/assignment', {
    method: 'POST',
    body: { technicianId },
    token,
  });
}

export function findAssignment(token: string, serviceOrderId: string): Promise<Assignment> {
  return request<Assignment>('/service-order/' + serviceOrderId + '/assignment', { token });
}

export function findDiagnostic(token: string, serviceOrderId: string): Promise<Diagnostic> {
  return request<Diagnostic>('/service-order/' + serviceOrderId + '/diagnostic', { token });
}

export function recordDiagnostic(
  token: string,
  serviceOrderId: string,
  finding: string,
  componentToRepair: string,
): Promise<Diagnostic> {
  return request<Diagnostic>('/service-order/' + serviceOrderId + '/diagnostic', {
    method: 'POST',
    body: { finding, componentToRepair },
    token,
  });
}

export function listIntervention(token: string, serviceOrderId: string): Promise<Intervention[]> {
  return request<Intervention[]>('/service-order/' + serviceOrderId + '/intervention', { token });
}

export function registerIntervention(
  token: string,
  serviceOrderId: string,
  description: string,
  laborHourCount: number,
  part: PartUsage[],
): Promise<Intervention> {
  return request<Intervention>('/service-order/' + serviceOrderId + '/intervention', {
    method: 'POST',
    body: { description, laborHourCount, part },
    token,
  });
}
