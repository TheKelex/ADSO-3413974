/* Route guard: without a session the user is sent back to the login screen. */
import { Navigate } from 'react-router-dom';
import type { ReactNode } from 'react';

import { useSession } from '../shared/SessionContext';
import type { Session } from '../services/session_service';

interface ProtectedRouteProps {
  children: ReactNode;
  requiredRole?: Session['role'];
}

export function ProtectedRoute({ children, requiredRole }: ProtectedRouteProps) {
  const { session } = useSession();
  if (!session) {
    return <Navigate to="/login" replace />;
  }
  if (requiredRole && session.role !== requiredRole) {
    return <Navigate to="/dashboard" replace />;
  }
  return <>{children}</>;
}
