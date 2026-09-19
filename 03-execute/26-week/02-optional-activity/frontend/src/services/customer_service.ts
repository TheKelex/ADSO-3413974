/* Customer registry service. */
import { request } from './api_client';

export interface Customer {
  id: string;
  fullName: string;
  documentNumber: string;
  phone: string;
  email: string;
  createdAt: string;
}

export interface NewCustomer {
  fullName: string;
  documentNumber: string;
  phone: string;
  email: string;
}

export function listCustomer(token: string): Promise<Customer[]> {
  return request<Customer[]>('/customer', { token });
}

export function createCustomer(token: string, customer: NewCustomer): Promise<Customer> {
  return request<Customer>('/customer', { method: 'POST', body: customer, token });
}
