/*
 * Session provider. It holds the signed in identity for the whole tree and is
 * the only consumer of the session service, so no screen reads storage.
 */
import { createContext, useCallback, useContext, useMemo, useState } from 'react';
import type { ReactNode } from 'react';

import {
  clearStoredSession,
  readStoredSession,
  signIn as requestSignIn,
  storeSession,
} from '../services/session_service';
import type { Session } from '../services/session_service';

interface SessionContextValue {
  session: Session | null;
  signIn: (username: string, password: string) => Promise<void>;
  signOut: () => void;
  isAdministrator: boolean;
}

const SessionContext = createContext<SessionContextValue | null>(null);

export function SessionProvider({ children }: { children: ReactNode }) {
  const [session, setSession] = useState<Session | null>(() => readStoredSession());

  const signIn = useCallback(async (username: string, password: string) => {
    const created = await requestSignIn(username, password);
    storeSession(created);
    setSession(created);
  }, []);

  const signOut = useCallback(() => {
    clearStoredSession();
    setSession(null);
  }, []);

  const value = useMemo<SessionContextValue>(
    () => ({
      session,
      signIn,
      signOut,
      isAdministrator: session?.role === 'ADMINISTRATOR',
    }),
    [session, signIn, signOut],
  );

  return <SessionContext.Provider value={value}>{children}</SessionContext.Provider>;
}

export function useSession(): SessionContextValue {
  const value = useContext(SessionContext);
  if (!value) {
    throw new Error('useSession must be used inside SessionProvider');
  }
  return value;
}

/** The token of the current session, or an empty string when signed out. */
export function useToken(): string {
  return useSession().session?.token ?? '';
}
