import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:5000';
const TOKEN =
  'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzY5Njc2Mjk5fQ.WR1WaNTuYi-zBybFhxoEadw1-AAgmhqEvanDI50cjgo';

export const options = {
  scenarios: {
    stress_test: {
      executor: 'ramping-vus',
      startVUs: 100,
      stages: [
        { duration: '30s', target: 300 },
        { duration: '30s', target: 600 },
        { duration: '30s', target: 1000 },
        { duration: '30s', target: 1500 },
      ],
    },
  },
};

const year = 2026;
const month = 1;
const merchantId = 10;
const page = 1;
const pageSize = 10;

export default function () {
  const params = { headers: { Authorization: TOKEN } };

  const basicEndpoints = [
    `/api/transaction?page=${page}&page_size=${pageSize}`,
    `/api/transaction/active?page=${page}&page_size=${pageSize}`,
    `/api/transaction/trashed?page=${page}&page_size=${pageSize}`,

    // ===== BY MERCHANT =====
    `/api/transaction/merchant/${merchantId}?page=${page}&page_size=${pageSize}`,

    // ===== GLOBAL TRANSACTION =====
    `/api/transaction/monthly-success?year=${year}&month=${month}`,
    `/api/transaction/yearly-success?year=${year}`,
    `/api/transaction/monthly-failed?year=${year}&month=${month}`,
    `/api/transaction/yearly-failed?year=${year}`,

    // ===== BY MERCHANT =====
    `/api/transaction/merchant/monthly-success?year=${year}&month=${month}&merchant_id=${merchantId}`,
    `/api/transaction/merchant/yearly-success?year=${year}&merchant_id=${merchantId}`,
    `/api/transaction/merchant/monthly-failed?year=${year}&month=${month}&merchant_id=${merchantId}`,
    `/api/transaction/merchant/yearly-failed?year=${year}&merchant_id=${merchantId}`,

    // ===== BY METHOD (GLOBAL) =====
    `/api/transaction/monthly-method-success?year=${year}&month=${month}`,
    `/api/transaction/yearly-method-success?year=${year}`,
    `/api/transaction/monthly-method-failed?year=${year}&month=${month}`,
    `/api/transaction/yearly-method-failed?year=${year}`,

    // ===== BY METHOD + MERCHANT =====
    `/api/transaction/merchant/monthly-method-success?year=${year}&month=${month}&merchant_id=${merchantId}`,
    `/api/transaction/merchant/yearly-method-success?year=${year}&merchant_id=${merchantId}`,
    `/api/transaction/merchant/monthly-method-failed?year=${year}&month=${month}&merchant_id=${merchantId}`,
    `/api/transaction/merchant/yearly-method-failed?year=${year}&merchant_id=${merchantId}`,
  ];

  for (let endpoint of basicEndpoints) {
    let res = http.get(`${BASE_URL}${endpoint}`, params);
    check(res, { [`GET ${endpoint} success`]: (r) => r.status === 200 });
  }

  sleep(0.1);
}
