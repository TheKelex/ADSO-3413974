/*
 * Hook that owns the four states of a server read: loading, error, empty and
 * success. Every screen consumes it instead of reimplementing three pieces of
 * state per view.
 */
import { useCallback, useEffect, useState } from 'react';

import { ApiError } from '../services/api_client';

export interface AsyncData<T> {
  data: T | null;
  loading: boolean;
  error: string;
  reload: () => void;
}

export function useAsyncData<T>(loader: () => Promise<T>, dependency: unknown[]): AsyncData<T> {
  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [reloadToken, setReloadToken] = useState(0);

  const reload = useCallback(() => setReloadToken((previous) => previous + 1), []);

  useEffect(() => {
    let active = true;
    setLoading(true);
    setError('');
    loader()
      .then((result) => {
        if (active) {
          setData(result);
        }
      })
      .catch((failure: unknown) => {
        if (!active) {
          return;
        }
        setData(null);
        setError(
          failure instanceof ApiError ? failure.message : 'No se pudo cargar la informacion.',
        );
      })
      .finally(() => {
        if (active) {
          setLoading(false);
        }
      });
    return () => {
      active = false;
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [...dependency, reloadToken]);

  return { data, loading, error, reload };
}
