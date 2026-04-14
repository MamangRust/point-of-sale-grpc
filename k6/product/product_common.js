import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:5000';
const TOKEN = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzY5NjkzMjk0fQ.VSNNewCyWsh69LIqN3wxuIk11qyrxn76Kr7rIMIAkMY';

const PRODUCT_ID = 1;
const MERCHANT_ID = 1;
const CATEGORY_NAME = 'Electronics';

export default function () {
  const params = {
    headers: { Authorization: TOKEN, 'Content-Type': 'application/json' },
  };

  const basicEndpoints = [
    `/api/product?page=1&page_size=10`,
    `/api/product/${PRODUCT_ID}`,
    `/api/product/merchant/${MERCHANT_ID}?page=1&page_size=10`,
    `/api/product/category/${CATEGORY_NAME}?page=1&page_size=10`,
    `/api/product/active?page=1&page_size=10`,
    `/api/product/trashed?page=1&page_size=10`,
  ];

  for (let endpoint of basicEndpoints) {
    let res = http.get(`${BASE_URL}${endpoint}`, params);
    check(res, { [`GET ${endpoint} success`]: (r) => r.status === 200 });
  }

  sleep(0.1);
}
