# Merchant Policy API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find All Merchant Policies
GET {{baseUrl}}/api/merchant-policy
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Merchant Policy
POST {{baseUrl}}/api/merchant-policy/create
Authorization: Bearer {{accessToken}}
{
    "merchant_id": 1,
    "policy_type": "Shipping",
    "title": "Standard Shipping Policy",
    "description": "Items are shipped within 2-3 business days."
}
HTTP 200
[Captures]
policyId: jsonpath "$.data.id"

# Find Policy by ID
GET {{baseUrl}}/api/merchant-policy/{{policyId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Policy
POST {{baseUrl}}/api/merchant-policy/update/{{policyId}}
Authorization: Bearer {{accessToken}}
{
    "policy_type": "Return",
    "title": "Return Policy",
    "description": "Returns are accepted within 30 days of purchase."
}
HTTP 200

# Trash Policy
POST {{baseUrl}}/api/merchant-policy/trashed/{{policyId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Policy
POST {{baseUrl}}/api/merchant-policy/restore/{{policyId}}
Authorization: Bearer {{accessToken}}
HTTP 200
