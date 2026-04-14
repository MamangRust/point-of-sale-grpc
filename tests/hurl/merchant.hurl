# Login
POST {{baseUrl}}/api/auth/login
{
    "email": "john.doe@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"
refreshToken: jsonpath "$.data.refresh_token"

# Find All Merchant
GET {{baseUrl}}/api/merchant
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Merchant
POST {{baseUrl}}/api/merchant/create
Authorization: Bearer {{accessToken}}
{
    "name": "Global Store",
    "user_id": {{userId}}
}
HTTP 200
[Captures]
merchantId: jsonpath "$.data.id"

# Find Merchant by ID
GET {{baseUrl}}/api/merchant/{{merchantId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Merchant
POST {{baseUrl}}/api/merchant/update/{{merchantId}}
Authorization: Bearer {{accessToken}}
{
    "name": "Global Store Updated",
    "user_id": {{userId}},
    "status": "active"
}
HTTP 200

# Find Active Merchant
GET {{baseUrl}}/api/merchant/active
Authorization: Bearer {{accessToken}}
HTTP 200

# Trash Merchant
POST {{baseUrl}}/api/merchant/trashed/{{merchantId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Merchant
POST {{baseUrl}}/api/merchant/restore/{{merchantId}}
Authorization: Bearer {{accessToken}}
HTTP 200
