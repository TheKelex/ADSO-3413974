/* Wrapper that renders the four states of a server read with one rule for the
   whole product: loading, error, empty and success. */
import type { ReactNode } from 'react';

interface DataStateProps {
  loading: boolean;
  error: string;
  empty: boolean;
  emptyMessage?: string;
  children: ReactNode;
}

export function DataState({ loading, error, empty, emptyMessage, children }: DataStateProps) {
  if (loading) {
    return (
      <p className="state-message" role="status" aria-live="polite">
        Cargando...
      </p>
    );
  }
  if (error) {
    return (
      <p className="banner banner--error" role="alert">
        {error}
      </p>
    );
  }
  if (empty) {
    return <p className="state-message">{emptyMessage ?? 'No hay registros.'}</p>;
  }
  return <>{children}</>;
}

/** Inline error banner for a rejected write. */
export function ErrorBanner({ message }: { message: string }) {
  if (!message) {
    return null;
  }
  return (
    <p className="banner banner--error" role="alert">
      {message}
    </p>
  );
}

/** Inline confirmation for an accepted write. */
export function SuccessBanner({ message }: { message: string }) {
  if (!message) {
    return null;
  }
  return (
    <p className="banner banner--success" role="status">
      {message}
    </p>
  );
}
