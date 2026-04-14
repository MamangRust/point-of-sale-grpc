# Merchant Business API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find All Merchant Business Info
GET {{baseUrl}}/api/merchant-business
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Merchant Business Info
POST {{baseUrl}}/api/merchant-business/create
Authorization: Bearer {{accessToken}}
{
    "merchant_id": 1,
    "business_type": "LLC",
    "tax_id": "TAX-123456",
    "established_year": 2010,
    "number_of_employees": 50,
    "website_url": "http://example.com"
}
HTTP 200
[Captures]
businessId: jsonpath "$.data.id"

# Find Business Info by ID
GET {{baseUrl}}/api/merchant-business/{{businessId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Business Info
POST {{baseUrl}}/api/merchant-business/update/{{businessId}}
Authorization: Bearer {{accessToken}}
{
    "business_type": "Corporation",
    "tax_id": "TAX-654321",
    "established_year": 2010,
    "number_of_employees": 100,
    "website_url": "http://example.org"
}
HTTP 200

# Trash Business Info
POST {{baseUrl}}/api/merchant-business/trashed/{{businessId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Business Info
POST {{baseUrl}}/api/merchant-business/restore/{{businessId}}
Authorization: Bearer {{accessToken}}
HTTP 200
