/* Customer registry screen: the creation form above the customer table. */
import { useState } from 'react';

import { ApiError } from '../../services/api_client';
import { createCustomer, listCustomer } from '../../services/customer_service';
import { DataState, ErrorBanner, SuccessBanner } from '../../shared/DataState';
import { useAsyncData } from '../../shared/useAsyncData';
import { useToken } from '../../shared/SessionContext';
import { formatDate } from '../../shared/format';

const EMPTY_FORM = { fullName: '', documentNumber: '', phone: '', email: '' };

export function CustomerPage() {
  const token = useToken();
  const { data, loading, error, reload } = useAsyncData(() => listCustomer(token), [token]);
  const [form, setForm] = useState(EMPTY_FORM);
  const [formError, setFormError] = useState('');
  const [confirmation, setConfirmation] = useState('');
  const [sending, setSending] = useState(false);

  const update = (field: keyof typeof EMPTY_FORM) => (event: React.ChangeEvent<HTMLInputElement>) =>
    setForm((previous) => ({ ...previous, [field]: event.target.value }));

  const submit = async (event: React.FormEvent) => {
    event.preventDefault();
    setFormError('');
    setConfirmation('');
    setSending(true);
    try {
      await createCustomer(token, form);
      setForm(EMPTY_FORM);
      setConfirmation('Cliente guardado.');
      reload();
    } catch (failure) {
      setFormError(
        failure instanceof ApiError ? failure.message : 'No se pudo guardar el cliente.',
      );
    } finally {
      setSending(false);
    }
  };

  return (
    <section>
      <h2 className="screen-title">Clientes</h2>
      <section className="card">
        <h3 className="card__title">Nuevo cliente</h3>
        <form onSubmit={submit} noValidate>
          <ErrorBanner message={formError} />
          <SuccessBanner message={confirmation} />
          <div className="form-grid">
            <div className="field">
              <label className="field__label" htmlFor="fullName">
                Nombre
              </label>
              <input
                className="field__input"
                id="fullName"
                value={form.fullName}
                onChange={update('fullName')}
                required
              />
            </div>
            <div className="field">
              <label className="field__label" htmlFor="documentNumber">
                Documento
              </label>
              <input
                className="field__input"
                id="documentNumber"
                value={form.documentNumber}
                onChange={update('documentNumber')}
                required
              />
            </div>
            <div className="field">
              <label className="field__label" htmlFor="phone">
                Telefono
              </label>
              <input
                className="field__input"
                id="phone"
                value={form.phone}
                onChange={update('phone')}
                required
              />
            </div>
            <div className="field">
              <label className="field__label" htmlFor="email">
                Correo
              </label>
              <input
                className="field__input"
                id="email"
                type="email"
                value={form.email}
                onChange={update('email')}
                required
              />
            </div>
          </div>
          <button type="submit" className="button button--primary" disabled={sending}>
            {sending ? 'Guardando...' : 'Guardar cliente'}
          </button>
        </form>
      </section>

      <section className="card">
        <h3 className="card__title">Clientes registrados</h3>
        <DataState
          loading={loading}
          error={error}
          empty={(data ?? []).length === 0}
          emptyMessage="No hay clientes registrados."
        >
          <div className="table-scroll">
            <table className="data-table">
              <thead>
                <tr>
                  <th scope="col">Nombre</th>
                  <th scope="col">Documento</th>
                  <th scope="col">Telefono</th>
                  <th scope="col">Correo</th>
                  <th scope="col">Registrado</th>
                </tr>
              </thead>
              <tbody>
                {(data ?? []).map((customer) => (
                  <tr key={customer.id}>
                    <td>{customer.fullName}</td>
                    <td>{customer.documentNumber}</td>
                    <td>{customer.phone}</td>
                    <td>{customer.email}</td>
                    <td>{formatDate(customer.createdAt)}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </DataState>
      </section>
    </section>
  );
}
