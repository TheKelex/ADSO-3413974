import { render, screen } from '@testing-library/react';
import { MemoryRouter, Route, Routes } from 'react-router-dom';
import { describe, expect, it } from 'vitest';

import { ProtectedRoute } from './ProtectedRoute';
import { SessionProvider } from '../shared/SessionContext';
import type { Session } from '../services/session_service';

function storeSession(role: Session['role']): void {
  window.localStorage.setItem(
    'workshop.session',
    JSON.stringify({
      token: 'token',
      expiresAt: new Date(Date.now() + 3600000).toISOString(),
      userId: 'user-1',
      username: role === 'ADMINISTRATOR' ? 'admin' : 'jperez',
      fullName: 'Usuario de prueba',
      role,
    }),
  );
}

function renderAdminRoute(role: Session['role']) {
  storeSession(role);
  return render(
    <SessionProvider>
      <MemoryRouter initialEntries={['/customers']}>
        <Routes>
          <Route path="/dashboard" element={<p>Panel principal</p>} />
          <Route
            path="/customers"
            element={
              <ProtectedRoute requiredRole="ADMINISTRATOR">
                <p>Clientes</p>
              </ProtectedRoute>
            }
          />
        </Routes>
      </MemoryRouter>
    </SessionProvider>,
  );
}

describe('ProtectedRoute', () => {
  it('redirige a un técnico desde una ruta exclusiva de administrador', () => {
    renderAdminRoute('TECHNICIAN');

    expect(screen.getByText('Panel principal')).toBeInTheDocument();
    expect(screen.queryByText('Clientes')).not.toBeInTheDocument();
  });

  it('permite el acceso de un administrador a su ruta privada', () => {
    renderAdminRoute('ADMINISTRATOR');

    expect(screen.getByText('Clientes')).toBeInTheDocument();
  });
});
