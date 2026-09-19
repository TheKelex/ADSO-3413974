import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter } from 'react-router-dom';
import { describe, expect, it, vi } from 'vitest';

import { LoginPage } from './LoginPage';
import { SessionProvider } from '../../shared/SessionContext';

function renderLogin() {
  return render(
    <SessionProvider>
      <MemoryRouter initialEntries={['/login']}>
        <LoginPage />
      </MemoryRouter>
    </SessionProvider>,
  );
}

function jsonResponse(status: number, payload: unknown): Response {
  return new Response(JSON.stringify(payload), {
    status,
    headers: { 'Content-Type': 'application/json' },
  });
}

describe('login screen', () => {
  it('shows the form in Spanish with labelled fields', () => {
    renderLogin();
    expect(screen.getByLabelText('Usuario')).toBeInTheDocument();
    expect(screen.getByLabelText('Contrasena')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: 'Ingresar' })).toBeInTheDocument();
  });

  it('stores the session when the credentials are valid', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue(
        jsonResponse(200, {
          token: 'a-token',
          expiresAt: new Date(Date.now() + 3600000).toISOString(),
          userId: 'user-1',
          username: 'admin',
          fullName: 'Administrador del taller',
          role: 'ADMINISTRATOR',
        }),
      ),
    );
    renderLogin();

    await userEvent.type(screen.getByLabelText('Usuario'), 'admin');
    await userEvent.type(screen.getByLabelText('Contrasena'), 'Admin2026');
    await userEvent.click(screen.getByRole('button', { name: 'Ingresar' }));

    await waitFor(() => {
      expect(window.localStorage.getItem('workshop.session')).toContain('a-token');
    });
  });

  it('shows the Spanish error and keeps no session when the password is wrong', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue(
        jsonResponse(401, { code: 'unauthorized', message: 'Usuario o contrasena incorrectos.' }),
      ),
    );
    renderLogin();

    await userEvent.type(screen.getByLabelText('Usuario'), 'admin');
    await userEvent.type(screen.getByLabelText('Contrasena'), 'wrong');
    await userEvent.click(screen.getByRole('button', { name: 'Ingresar' }));

    expect(await screen.findByRole('alert')).toHaveTextContent('Usuario o contrasena incorrectos.');
    expect(window.localStorage.getItem('workshop.session')).toBeNull();
    expect(screen.getByLabelText('Contrasena')).toHaveValue('');
  });
});
