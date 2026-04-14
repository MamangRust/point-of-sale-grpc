# Banner API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find All Banners
GET {{baseUrl}}/api/banner
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Banner
POST {{baseUrl}}/api/banner/create
Authorization: Bearer {{accessToken}}
{
    "name": "Summer Sale",
    "start_date": "2026-06-01",
    "end_date": "2026-08-31",
    "start_time": "00:00:00",
    "end_time": "23:59:59",
    "is_active": true
}
HTTP 200
[Captures]
bannerId: jsonpath "$.data.id"

# Find Banner by ID
GET {{baseUrl}}/api/banner/{{bannerId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Banner
POST {{baseUrl}}/api/banner/update/{{bannerId}}
Authorization: Bearer {{accessToken}}
{
    "name": "Summer Sale Updated",
    "start_date": "2026-06-15",
    "end_date": "2026-09-15",
    "start_time": "08:00:00",
    "end_time": "22:00:00",
    "is_active": true
}
HTTP 200

# Trash Banner
POST {{baseUrl}}/api/banner/trashed/{{bannerId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Banner
POST {{baseUrl}}/api/banner/restore/{{bannerId}}
Authorization: Bearer {{accessToken}}
HTTP 200
