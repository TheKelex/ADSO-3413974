/* Vehicle registry screen: the owner is chosen from the real customer list,
   never typed as a raw identifier. */
import { useState } from 'react';
import { Link } from 'react-router-dom';

import { ApiError } from '../../services/api_client';
import { listCustomer } from '../../services/customer_service';
import { createVehicle, listVehicle } from '../../services/vehicle_service';
import { DataState, ErrorBanner, SuccessBanner } from '../../shared/DataState';
import { useAsyncData } from '../../shared/useAsyncData';
import { useToken } from '../../shared/SessionContext';

const CURRENT_YEAR = new Date().getFullYear();
const EMPTY_FORM = {
  customerId: '',
  plate: '',
  vin: '',
  brand: '',
  model: '',
  modelYear: String(CURRENT_YEAR),
};

export function VehiclePage() {
  const token = useToken();
  const vehicle = useAsyncData(() => listVehicle(token), [token]);
  const customer = useAsyncData(() => listCustomer(token), [token]);
  const [form, setForm] = useState(EMPTY_FORM);
  const [formError, setFormError] = useState('');
  const [confirmation, setConfirmation] = useState('');
  const [sending, setSending] = useState(false);

  const update =
    (field: keyof typeof EMPTY_FORM) =>
    (event: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) =>
      setForm((previous) => ({ ...previous, [field]: event.target.value }));

  const submit = async (event: React.FormEvent) => {
    event.preventDefault();
    setFormError('');
    setConfirmation('');
    setSending(true);
    try {
      await createVehicle(token, {
        customerId: form.customerId,
        plate: form.plate,
        vin: form.vin,
        brand: form.brand,
        model: form.model,
        modelYear: Number(form.modelYear),
      });
      setForm(EMPTY_FORM);
      setConfirmation('Vehiculo guardado.');
      vehicle.reload();
    } catch (failure) {
      setFormError(
        failure instanceof ApiError ? failure.message : 'No se pudo guardar el vehiculo.',
      );
    } finally {
      setSending(false);
    }
  };

  return (
    <section>
      <h2 className="screen-title">Vehiculos</h2>
      <section className="card">
        <h3 className="card__title">Nuevo vehiculo</h3>
        <form onSubmit={submit} noValidate>
          <ErrorBanner message={formError} />
          <SuccessBanner message={confirmation} />
          <div className="form-grid">
            <div className="field">
              <label className="field__label" htmlFor="customerId">
                Propietario
              </label>
              <select
                className="field__input"
                id="customerId"
                value={form.customerId}
                onChange={update('customerId')}
                required
              >
                <option value="">Seleccione un cliente</option>
                {(customer.data ?? []).map((item) => (
                  <option key={item.id} value={item.id}>
                    {item.fullName}
                  </option>
                ))}
              </select>
            </div>
            <div className="field">
              <label className="field__label" htmlFor="plate">
                Placa
              </label>
              <input
                className="field__input"
                id="plate"
                value={form.plate}
                onChange={update('plate')}
                required
              />
            </div>
            <div className="field">
              <label className="field__label" htmlFor="vin">
                VIN
              </label>
              <input
                className="field__input"
                id="vin"
                value={form.vin}
                onChange={update('vin')}
                required
              />
            </div>
            <div className="field">
              <label className="field__label" htmlFor="brand">
                Marca
              </label>
              <input
                className="field__input"
                id="brand"
                value={form.brand}
                onChange={update('brand')}
                required
              />
            </div>
            <div className="field">
              <label className="field__label" htmlFor="model">
                Modelo
              </label>
              <input
                className="field__input"
                id="model"
                value={form.model}
                onChange={update('model')}
                required
              />
            </div>
            <div className="field">
              <label className="field__label" htmlFor="modelYear">
                Ano
              </label>
              <select
                className="field__input"
                id="modelYear"
                value={form.modelYear}
                onChange={update('modelYear')}
                required
              >
                {Array.from({ length: 40 }, (_, index) => CURRENT_YEAR - index).map((year) => (
                  <option key={year} value={String(year)}>
                    {year}
                  </option>
                ))}
              </select>
            </div>
          </div>
          <button type="submit" className="button button--primary" disabled={sending}>
            {sending ? 'Guardando...' : 'Guardar vehiculo'}
          </button>
        </form>
      </section>

      <section className="card">
        <h3 className="card__title">Vehiculos registrados</h3>
        <DataState
          loading={vehicle.loading}
          error={vehicle.error}
          empty={(vehicle.data ?? []).length === 0}
          emptyMessage="No hay vehiculos registrados."
        >
          <div className="table-scroll">
            <table className="data-table">
              <thead>
                <tr>
                  <th scope="col">Placa</th>
                  <th scope="col">VIN</th>
                  <th scope="col">Marca</th>
                  <th scope="col">Modelo</th>
                  <th scope="col">Ano</th>
                  <th scope="col">Propietario</th>
                  <th scope="col">Historial</th>
                </tr>
              </thead>
              <tbody>
                {(vehicle.data ?? []).map((item) => (
                  <tr key={item.id}>
                    <td>{item.plate}</td>
                    <td>{item.vin}</td>
                    <td>{item.brand}</td>
                    <td>{item.model}</td>
                    <td>{item.modelYear}</td>
                    <td>{item.ownerName}</td>
                    <td>
                      <Link className="button button--secondary" to={'/vehicles/' + item.id + '/timeline'}>
                        Ver historial
                      </Link>
                    </td>
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
