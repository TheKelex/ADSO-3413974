// Test setup shared by every suite: it adds the DOM matchers and clears the
// stored session between tests so one test never inherits another session.
import '@testing-library/jest-dom/vitest';
import { afterEach, beforeEach, vi } from 'vitest';

beforeEach(() => {
  window.localStorage.clear();
});

afterEach(() => {
  vi.restoreAllMocks();
});
