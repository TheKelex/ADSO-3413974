/* Login screen. It is the only screen reachable without a session. */
import { useState } from 'react';
import { Navigate, useNavigate } from 'react-router-dom';

import { ApiError } from '../../services/api_client';
import { ErrorBanner } from '../../shared/DataState';
import { useSession } from '../../shared/SessionContext';

export function LoginPage() {
  const { session, signIn } = useSession();
  const navigate = useNavigate();
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [sending, setSending] = useState(false);

  if (session) {
    return <Navigate to="/dashboard" replace />;
  }

  const submit = async (event: React.FormEvent) => {
    event.preventDefault();
    setError('');
    setSending(true);
    try {
      await signIn(username, password);
      navigate('/dashboard', { replace: true });
    } catch (failure) {
      setPassword('');
      setError(
        failure instanceof ApiError ? failure.message : 'No se pudo iniciar sesion.',
      );
    } finally {
      setSending(false);
    }
  };

  return (
    <div className="login-layout">
      <section className="card login-card">
        <h1 className="card__title">Soporte Tecnico Automotriz</h1>
        <form onSubmit={submit} noValidate>
          <ErrorBanner message={error} />
          <div className="field">
            <label className="field__label" htmlFor="username">
              Usuario
            </label>
            <input
              className="field__input"
              id="username"
              name="username"
              autoComplete="username"
              value={username}
              onChange={(event) => setUsername(event.target.value)}
              required
            />
          </div>
          <div className="field">
            <label className="field__label" htmlFor="password">
              Contrasena
            </label>
            <input
              className="field__input"
              id="password"
              name="password"
              type="password"
              autoComplete="current-password"
              value={password}
              onChange={(event) => setPassword(event.target.value)}
              required
            />
          </div>
          <button type="submit" className="button button--primary" disabled={sending}>
            {sending ? 'Ingresando...' : 'Ingresar'}
          </button>
        </form>
      </section>
    </div>
  );
}
