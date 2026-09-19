import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';

import { AvailabilityBadge, StatusBadge, ValidityBadge } from './StatusBadge';
import { nextStatus, statusLabel } from './format';

describe('status badge', () => {
  it('writes every lifecycle status in Spanish', () => {
    render(<StatusBadge status="IN_DIAGNOSIS" />);
    expect(screen.getByText('En diagnostico')).toBeInTheDocument();
  });

  it('carries its meaning as text, never as colour alone', () => {
    render(<StatusBadge status="READY" />);
    const badge = screen.getByText('Listo');
    expect(badge.className).toContain('badge--ready');
    expect(badge.textContent).not.toBe('');
  });

  it('tells a valid warranty apart from an expired one', () => {
    const { rerender } = render(<ValidityBadge valid />);
    expect(screen.getByText('Vigente')).toBeInTheDocument();
    rerender(<ValidityBadge valid={false} />);
    expect(screen.getByText('Vencida')).toBeInTheDocument();
  });

  it('tells an available technician apart from a busy one', () => {
    const { rerender } = render(<AvailabilityBadge busy={false} />);
    expect(screen.getByText('Disponible')).toBeInTheDocument();
    rerender(<AvailabilityBadge busy />);
    expect(screen.getByText('Ocupado')).toBeInTheDocument();
  });
});

describe('lifecycle helpers', () => {
  it('follows the declared order of the lifecycle', () => {
    expect(nextStatus('RECEIVED')).toBe('IN_DIAGNOSIS');
    expect(nextStatus('IN_DIAGNOSIS')).toBe('IN_REPAIR');
    expect(nextStatus('IN_REPAIR')).toBe('READY');
    expect(nextStatus('READY')).toBe('DELIVERED');
  });

  it('offers no next step once the vehicle was delivered', () => {
    expect(nextStatus('DELIVERED')).toBeNull();
  });

  it('labels every status in Spanish', () => {
    expect(statusLabel('RECEIVED')).toBe('Recibido');
    expect(statusLabel('DELIVERED')).toBe('Entregado');
  });
});
