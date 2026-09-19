import { render, screen } from '@testing-library/react';
import { describe, expect, it, vi } from 'vitest';

import { DiagnosticPanel } from './DiagnosticPanel';
import { InterventionPanel } from './InterventionPanel';
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

function renderDeliveredPanels() {
  storeTechnicianSession();
  vi.stubGlobal(
    'fetch',
    vi.fn().mockImplementation((url: string) =>
      Promise.resolve(
        new Response(url.endsWith('/diagnostic') ? '' : '[]', {
          status: url.endsWith('/diagnostic') ? 404 : 200,
          headers: { 'Content-Type': 'application/json' },
        }),
      ),
    ),
  );
  return render(
    <SessionProvider>
      <DiagnosticPanel serviceOrderId="order-1" delivered onChange={vi.fn()} />
      <InterventionPanel serviceOrderId="order-1" delivered onChange={vi.fn()} />
    </SessionProvider>,
  );
}

describe('delivered order panels', () => {
  it('hide diagnostic and intervention write controls', async () => {
    renderDeliveredPanels();

    expect(await screen.findAllByText('La orden ya fue entregada.')).toHaveLength(2);
    expect(screen.queryByLabelText('Hallazgo')).not.toBeInTheDocument();
    expect(screen.queryByLabelText('Descripcion')).not.toBeInTheDocument();
    expect(screen.queryByRole('button', { name: 'Registrar intervencion' })).not.toBeInTheDocument();
  });
});
