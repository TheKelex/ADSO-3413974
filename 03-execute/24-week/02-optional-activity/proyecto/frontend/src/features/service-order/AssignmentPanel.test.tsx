import { render, screen } from '@testing-library/react';
import { describe, expect, it, vi } from 'vitest';

import { AssignmentPanel } from './AssignmentPanel';
import { SessionProvider } from '../../shared/SessionContext';

function storeTechnicianSession(): void {
  window.localStorage.setItem(
    'workshop.session',
    JSON.stringify({
      token: 'technician-token',
      expiresAt: new Date(Date.now() + 3600000).toISOString(),
      userId: 'user-2',
      username: 'jperez',
      fullName: 'Juan Perez',
      role: 'TECHNICIAN',
    }),
  );
}

describe('assignment panel', () => {
  it('does not request the technician catalog for a technician', async () => {
    storeTechnicianSession();
    vi.stubGlobal('fetch', vi.fn());

    render(
      <SessionProvider>
        <AssignmentPanel serviceOrderId="order-1" technicianName="Juan Perez" onChange={vi.fn()} />
      </SessionProvider>,
    );

    expect(await screen.findByText('Técnico asignado: Juan Perez.')).toBeInTheDocument();
    expect(fetch).not.toHaveBeenCalled();
  });
});
