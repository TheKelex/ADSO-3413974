/* App composes the providers and the routing. It holds no business rule, no
   API client and no demo state. */
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';

import { AppLayout } from './AppLayout';
import { ProtectedRoute } from './ProtectedRoute';
import { SessionProvider } from '../shared/SessionContext';
import { LoginPage } from '../features/login/LoginPage';
import { DashboardPage } from '../features/dashboard/DashboardPage';
import { CustomerPage } from '../features/customer/CustomerPage';
import { VehiclePage } from '../features/vehicle/VehiclePage';
import { TechnicianPage } from '../features/technician/TechnicianPage';
import { ServiceOrderPage } from '../features/service-order/ServiceOrderPage';
import { ServiceOrderDetailPage } from '../features/service-order/ServiceOrderDetailPage';
import { WarrantyPage } from '../features/warranty/WarrantyPage';
import { VehicleTimelinePage } from '../features/timeline/VehicleTimelinePage';

export function App() {
  return (
    <SessionProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route
            element={
              <ProtectedRoute>
                <AppLayout />
              </ProtectedRoute>
            }
          >
            <Route path="/dashboard" element={<DashboardPage />} />
            <Route
              path="/customers"
              element={
                <ProtectedRoute requiredRole="ADMINISTRATOR">
                  <CustomerPage />
                </ProtectedRoute>
              }
            />
            <Route
              path="/vehicles"
              element={
                <ProtectedRoute requiredRole="ADMINISTRATOR">
                  <VehiclePage />
                </ProtectedRoute>
              }
            />
            <Route
              path="/vehicles/:vehicleId/timeline"
              element={
                <ProtectedRoute requiredRole="ADMINISTRATOR">
                  <VehicleTimelinePage />
                </ProtectedRoute>
              }
            />
            <Route
              path="/technicians"
              element={
                <ProtectedRoute requiredRole="ADMINISTRATOR">
                  <TechnicianPage />
                </ProtectedRoute>
              }
            />
            <Route path="/service-orders" element={<ServiceOrderPage />} />
            <Route path="/service-orders/:serviceOrderId" element={<ServiceOrderDetailPage />} />
            <Route
              path="/warranties"
              element={
                <ProtectedRoute requiredRole="ADMINISTRATOR">
                  <WarrantyPage />
                </ProtectedRoute>
              }
            />
          </Route>
          <Route path="*" element={<Navigate to="/dashboard" replace />} />
        </Routes>
      </BrowserRouter>
    </SessionProvider>
  );
}
