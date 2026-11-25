/* Vehicle registry service. */
import { request } from './api_client';

export interface Vehicle {
  id: string;
  customerId: string;
  ownerName: string;
  plate: string;
  vin: string;
  brand: string;
  model: string;
  modelYear: number;
  createdAt: string;
}

export interface NewVehicle {
  customerId: string;
  plate: string;
  vin: string;
  brand: string;
  model: string;
  modelYear: number;
}

export function listVehicle(token: string): Promise<Vehicle[]> {
  return request<Vehicle[]>('/vehicle', { token });
}

export function createVehicle(token: string, vehicle: NewVehicle): Promise<Vehicle> {
  return request<Vehicle>('/vehicle', { method: 'POST', body: vehicle, token });
}
