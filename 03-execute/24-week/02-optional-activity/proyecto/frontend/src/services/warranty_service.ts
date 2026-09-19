/* Warranty service: issuance and the validity report at a consulted date. */
import { request } from './api_client';

export type WarrantyKind = 'LABOR' | 'PART';

export interface Warranty {
  id: string;
  interventionId: string;
  orderNumber: string;
  vehiclePlate: string;
  kind: WarrantyKind;
  coverageMonthCount: number;
  issuedAt: string;
  expirationDate: string;
  valid: boolean;
}

export function listWarranty(token: string, consultedAt: string): Promise<Warranty[]> {
  const query = consultedAt ? '?consultedAt=' + encodeURIComponent(consultedAt) : '';
  return request<Warranty[]>('/warranty' + query, { token });
}

export function issueWarranty(
  token: string,
  interventionId: string,
  kind: WarrantyKind,
  coverageMonthCount: number,
): Promise<Warranty> {
  return request<Warranty>('/warranty', {
    method: 'POST',
    body: { interventionId, kind, coverageMonthCount },
    token,
  });
}
