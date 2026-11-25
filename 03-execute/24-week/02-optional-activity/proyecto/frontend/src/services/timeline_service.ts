/* Clinical timeline service: the chronological history of one vehicle. */
import { request } from './api_client';
import type { Vehicle } from './vehicle_service';

export type TimelineKind = 'ORDER' | 'DIAGNOSTIC' | 'INTERVENTION' | 'WARRANTY';

export interface TimelineEntry {
  kind: TimelineKind;
  occurredAt: string;
  title: string;
  description: string;
  reference: string;
}

export interface Timeline {
  vehicle: Vehicle;
  entry: TimelineEntry[];
}

export function readTimeline(token: string, vehicleId: string): Promise<Timeline> {
  return request<Timeline>('/vehicle/' + vehicleId + '/timeline', { token });
}
