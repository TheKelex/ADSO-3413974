/* Dashboard service: the operational summary of the workshop. */
import { request } from './api_client';
import type { Technician } from './technician_service';
import type { ServiceOrderStatus } from './service_order_service';

export interface StatusCount {
  status: ServiceOrderStatus;
  count: number;
}

export interface Dashboard {
  openOrderCount: number;
  statusCount: StatusCount[];
  busyTechnician: Technician[];
}

export function readDashboard(token: string): Promise<Dashboard> {
  return request<Dashboard>('/dashboard', { token });
}
