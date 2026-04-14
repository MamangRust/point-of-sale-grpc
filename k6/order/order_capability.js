import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:5000';
const TOKEN = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzY5NjkzMjk0fQ.VSNNewCyWsh69LIqN3wxuIk11qyrxn76Kr7rIMIAkMY';

const ORDER_ID = 1;

export const options = {
  scenarios: {
    scalability_test: {
      executor: 'ramping-arrival-rate',
      startRate: 50,
      timeUnit: '1s',
      stages: [
        { duration: '1m', target: 100 },
        { duration: '1m', target: 300 },
        { duration: '1m', target: 600 },
      ],
      preAllocatedVUs: 100,
      maxVUs: 900,
    },
  },
};

export default function () {
  const params = {
    headers: { Authorization: TOKEN, 'Content-Type': 'application/json' },
  };

  const basicEndpoints = [
    `/api/order?page=1&page_size=10`,
    `/api/order/${ORDER_ID}`,
    `/api/order/active?page=1&page_size=10`,
    `/api/order/trashed?page=1&page_size=10`,
    `/api/order/monthly-total-revenue?year=2025&month=1`,
    `/api/order/yearly-total-revenue?year=2025`,
  ];

  for (let endpoint of basicEndpoints) {
    let res = http.get(`${BASE_URL}${endpoint}`, params);
    check(res, { [`GET ${endpoint} success`]: (r) => r.status === 200 });
  }

  sleep(0.1);
}
