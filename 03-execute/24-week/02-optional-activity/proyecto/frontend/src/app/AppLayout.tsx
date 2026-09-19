/* Shell of every signed in screen: the brand header, the navigation the role
   is allowed to see, and the outlet where the active screen renders. */
import { NavLink, Outlet, useNavigate } from 'react-router-dom';

import { useSession } from '../shared/SessionContext';

interface NavigationItem {
  to: string;
  label: string;
  administratorOnly: boolean;
}

const NAVIGATION: NavigationItem[] = [
  { to: '/dashboard', label: 'Panel', administratorOnly: false },
  { to: '/service-orders', label: 'Ordenes', administratorOnly: false },
  { to: '/customers', label: 'Clientes', administratorOnly: true },
  { to: '/vehicles', label: 'Vehiculos', administratorOnly: true },
  { to: '/technicians', label: 'Tecnicos', administratorOnly: true },
  { to: '/warranties', label: 'Garantias', administratorOnly: true },
];

export function AppLayout() {
  const { session, signOut, isAdministrator } = useSession();
  const navigate = useNavigate();

  const leave = () => {
    signOut();
    navigate('/login', { replace: true });
  };

  return (
    <>
      <header className="app-header">
        <h1 className="app-header__brand">Soporte Tecnico Automotriz</h1>
        <nav className="app-header__nav" aria-label="Navegacion principal">
          {NAVIGATION.filter((item) => isAdministrator || !item.administratorOnly).map((item) => (
            <NavLink
              key={item.to}
              to={item.to}
              className={({ isActive }) =>
                'app-header__link' + (isActive ? ' app-header__link--active' : '')
              }
            >
              {item.label}
            </NavLink>
          ))}
        </nav>
        <div className="app-header__user">
          <span>{session?.fullName}</span>
          <button type="button" className="button button--secondary" onClick={leave}>
            Salir
          </button>
        </div>
      </header>
      <main className="app-main">
        <Outlet />
      </main>
    </>
  );
}
