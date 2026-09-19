import { render, screen } from '@testing-library/react';
import { describe, expect, it, vi } from 'vitest';

import { WarrantyPage } from './WarrantyPage';
import { SessionProvider } from '../../shared/SessionContext';

const VALID_WARRANTY = {
  id: 'warranty-1',
  interventionId: 'intervention-1',
  orderNumber: 'OS-0001',
  vehiclePlate: 'ABC123',
  kind: 'LABOR',
  coverageMonthCount: 12,
  issuedAt: '2026-01-15T10:00:00Z',
  expirationDate: '2027-01-15T10:00:00Z',
  valid: true,
};

function storeSession() {
  window.localStorage.setItem(
    'workshop.session',
    JSON.stringify({
      token: 'a-token',
      expiresAt: new Date(Date.now() + 3600000).toISOString(),
      userId: 'user-1',
      username: 'admin',
      fullName: 'Administrador del taller',
      role: 'ADMINISTRATOR',
    }),
  );
}

function stubWarranty(payload: unknown) {
  vi.stubGlobal(
    'fetch',
    vi.fn().mockResolvedValue(
      new Response(JSON.stringify(payload), {
        status: 200,
        headers: { 'Content-Type': 'application/json' },
      }),
    ),
  );
}

function renderWarranty() {
  return render(
    <SessionProvider>
      <WarrantyPage />
    </SessionProvider>,
  );
}

describe('warranty screen', () => {
  it('reads a warranty that still covers the intervention as valid', async () => {
    storeSession();
    stubWarranty([VALID_WARRANTY]);
    renderWarranty();

    expect(await screen.findByText('Vigente')).toBeInTheDocument();
    expect(screen.getByText('Mano de obra')).toBeInTheDocument();
    expect(screen.getByText('12 meses')).toBeInTheDocument();
  });

  it('reads a warranty past its expiration date as expired', async () => {
    storeSession();
    stubWarranty([{ ...VALID_WARRANTY, valid: false, kind: 'PART' }]);
    renderWarranty();

    expect(await screen.findByText('Vencida')).toBeInTheDocument();
    expect(screen.getByText('Repuesto')).toBeInTheDocument();
  });

  it('offers the date field that drives the validity report', async () => {
    storeSession();
    stubWarranty([VALID_WARRANTY]);
    renderWarranty();

    expect(await screen.findByLabelText('Consultar a la fecha')).toBeInTheDocument();
  });

  it('shows the empty state when no warranty was issued yet', async () => {
    storeSession();
    stubWarranty([]);
    renderWarranty();

    expect(await screen.findByText('No hay garantias emitidas.')).toBeInTheDocument();
  });
});
