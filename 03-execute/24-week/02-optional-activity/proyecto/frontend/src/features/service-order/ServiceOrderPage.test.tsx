import { render, screen, within } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import { describe, expect, it, vi } from 'vitest';

import { ServiceOrderPage } from './ServiceOrderPage';
import { SessionProvider } from '../../shared/SessionContext';

const ORDER = {
  id: 'order-1',
  orderNumber: 'OS-0001',
  vehicleId: 'vehicle-1',
  vehiclePlate: 'ABC123',
  technicianName: '',
  reportedFailure: 'Ruido en el motor',
  status: 'RECEIVED',
  receivedAt: '2026-03-01T09:00:00Z',
  updatedAt: '2026-03-01T09:00:00Z',
};

function storeAdministratorSession() {
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

function storeTechnicianSession() {
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

function stubApi(order: unknown[]) {
  vi.stubGlobal(
    'fetch',
    vi.fn().mockImplementation((url: string) => {
      const payload = url.startsWith('/api/service-order') ? order : [];
      return Promise.resolve(
        new Response(JSON.stringify(payload), {
          status: 200,
          headers: { 'Content-Type': 'application/json' },
        }),
      );
    }),
  );
}

function renderOrders() {
  return render(
    <SessionProvider>
      <MemoryRouter>
        <ServiceOrderPage />
      </MemoryRouter>
    </SessionProvider>,
  );
}

describe('service order screen', () => {
  it('announces the loading state before the data arrives', () => {
    storeAdministratorSession();
    stubApi([ORDER]);
    renderOrders();
    expect(screen.getAllByRole('status').length).toBeGreaterThan(0);
  });

  it('lists the order with its plate and its Spanish status badge', async () => {
    storeAdministratorSession();
    stubApi([ORDER]);
    renderOrders();

    expect(await screen.findByText('OS-0001')).toBeInTheDocument();
    const table = within(screen.getByRole('table'));
    expect(table.getByText('ABC123')).toBeInTheDocument();
    expect(table.getByText('Recibido')).toBeInTheDocument();
    expect(table.getByText('Sin asignar')).toBeInTheDocument();
  });

  it('shows the empty state when the workshop has no order in that status', async () => {
    storeAdministratorSession();
    stubApi([]);
    renderOrders();

    expect(await screen.findByText('No hay ordenes con ese estado.')).toBeInTheDocument();
  });

  it('offers the check-in form to the workshop manager', async () => {
    storeAdministratorSession();
    stubApi([ORDER]);
    renderOrders();

    expect(await screen.findByLabelText('Falla reportada')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: 'Registrar ingreso' })).toBeInTheDocument();
  });

  it('does not request the vehicle catalog for a technician', async () => {
    storeTechnicianSession();
    stubApi([ORDER]);
    renderOrders();

    expect(await screen.findByText('OS-0001')).toBeInTheDocument();
    expect(vi.mocked(fetch)).toHaveBeenCalledWith(
      '/api/service-order',
      expect.anything(),
    );
    expect(vi.mocked(fetch)).not.toHaveBeenCalledWith('/api/vehicle', expect.anything());
  });
});
