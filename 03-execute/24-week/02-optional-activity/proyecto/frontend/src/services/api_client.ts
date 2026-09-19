/*
 * Single HTTP client of the product. Every service goes through it, so the
 * session token, the JSON headers, the timeout and the translation of a
 * backend error into a message the user reads live in one place.
 */
export const API_BASE = '/api';

const REQUEST_TIMEOUT_MILLISECOND = 15000;

export class ApiError extends Error {
  readonly status: number;
  readonly code: string;

  constructor(status: number, code: string, message: string) {
    super(message);
    this.name = 'ApiError';
    this.status = status;
    this.code = code;
  }
}

export interface RequestOption {
  method?: 'GET' | 'POST';
  body?: unknown;
  token?: string | null;
  signal?: AbortSignal;
}

function buildHeader(option: RequestOption): HeadersInit {
  const header: Record<string, string> = { Accept: 'application/json' };
  if (option.body !== undefined) {
    header['Content-Type'] = 'application/json';
  }
  if (option.token) {
    header.Authorization = 'Bearer ' + option.token;
  }
  return header;
}

async function readError(response: Response): Promise<ApiError> {
  let code = 'internal_error';
  let message = 'Ocurrio un error inesperado. Intente de nuevo.';
  try {
    const payload = (await response.json()) as { code?: string; message?: string };
    if (payload && typeof payload.message === 'string' && payload.message.length > 0) {
      message = payload.message;
    }
    if (payload && typeof payload.code === 'string' && payload.code.length > 0) {
      code = payload.code;
    }
  } catch {
    if (response.status === 401) {
      message = 'La sesion expiro. Vuelva a ingresar.';
      code = 'unauthorized';
    }
  }
  return new ApiError(response.status, code, message);
}

/** Perform a request against the API and return the parsed payload. */
export async function request<T>(path: string, option: RequestOption = {}): Promise<T> {
  const controller = new AbortController();
  const timer = setTimeout(() => controller.abort(), REQUEST_TIMEOUT_MILLISECOND);
  if (option.signal) {
    option.signal.addEventListener('abort', () => controller.abort());
  }
  try {
    const response = await fetch(API_BASE + path, {
      method: option.method ?? 'GET',
      headers: buildHeader(option),
      body: option.body === undefined ? undefined : JSON.stringify(option.body),
      signal: controller.signal,
    });
    if (!response.ok) {
      throw await readError(response);
    }
    if (response.status === 204) {
      return undefined as T;
    }
    return (await response.json()) as T;
  } catch (error) {
    if (error instanceof ApiError) {
      throw error;
    }
    if (error instanceof DOMException && error.name === 'AbortError') {
      throw new ApiError(0, 'timeout', 'La peticion tardo demasiado. Intente de nuevo.');
    }
    throw new ApiError(0, 'network_error', 'No se pudo conectar con el servidor.');
  } finally {
    clearTimeout(timer);
  }
}
