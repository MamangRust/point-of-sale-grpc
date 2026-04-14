# Merchant Award (Certification) API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find All Merchant Awards
GET {{baseUrl}}/api/merchant-certification
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Merchant Award
POST {{baseUrl}}/api/merchant-certification/create
Authorization: Bearer {{accessToken}}
{
    "merchant_id": 1,
    "title": "Best Organic Farm 2024",
    "description": "Awarded for sustainable practices",
    "issued_by": "EcoCert",
    "issue_date": "2024-01-01",
    "expiry_date": "2025-01-01",
    "certificate_url": "http://ecocert.com/cert/123"
}
HTTP 200
[Captures]
awardId: jsonpath "$.data.id"

# Find Award by ID
GET {{baseUrl}}/api/merchant-certification/{{awardId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Award
POST {{baseUrl}}/api/merchant-certification/update/{{awardId}}
Authorization: Bearer {{accessToken}}
{
    "title": "Best Organic Farm 2024 Updated",
    "description": "Updated description",
    "issued_by": "EcoCert Global",
    "issue_date": "2024-01-01",
    "expiry_date": "2026-01-01",
    "certificate_url": "http://ecocert.com/cert/123-new"
}
HTTP 200

# Trash Award
POST {{baseUrl}}/api/merchant-certification/trashed/{{awardId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Award
POST {{baseUrl}}/api/merchant-certification/restore/{{awardId}}
Authorization: Bearer {{accessToken}}
HTTP 200
