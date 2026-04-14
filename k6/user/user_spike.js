import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:5000';
const TOKEN = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzY5NjkzMjk0fQ.VSNNewCyWsh69LIqN3wxuIk11qyrxn76Kr7rIMIAkMY';

const USER_ID = 11;

export const options = {
  scenarios: {
    spike_test: {
      executor: 'ramping-vus',
      startVUs: 100,
      stages: [
        { duration: '10s', target: 100 },
        { duration: '30s', target: 1500 },
        { duration: '10s', target: 100 },
        { duration: '1m', target: 0 },
      ],
    },
  },
};

export default function () {
  const params = {
    headers: { Authorization: TOKEN, 'Content-Type': 'application/json' },
  };

  const basicEndpoints = [
    `/api/user?page=1&page_size=10`,
    `/api/user/${USER_ID}`,
    `/api/user/active?page=1&page_size=10`,
    `/api/user/trashed?page=1&page_size=10`,
  ];

  for (let endpoint of basicEndpoints) {
    let res = http.get(`${BASE_URL}${endpoint}`, params);
    check(res, { [`GET ${endpoint} success`]: (r) => r.status === 200 });
  }

  sleep(0.1);
}
