POST http://localhost:5000/api/auth/login
Content-Type: application/json
{
  "email": "john.doe@example.com",
  "password": "securepassword"
}
HTTP 200
[Captures]
access_token: jsonpath "$.data.access_token"
refresh_token: jsonpath "$.data.refresh_token"
