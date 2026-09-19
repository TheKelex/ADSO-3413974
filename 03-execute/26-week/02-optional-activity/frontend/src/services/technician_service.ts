/* Technician allocation service. */
import { request } from './api_client';

export interface Technician {
  id: string;
  userId: string;
  fullName: string;
  specialty: string;
  busy: boolean;
  activeOrderId: string;
  activeOrderNumber: string;
  activeVehiclePlate: string;
}

export function listTechnician(token: string): Promise<Technician[]> {
  return request<Technician[]>('/technician', { token });
}
