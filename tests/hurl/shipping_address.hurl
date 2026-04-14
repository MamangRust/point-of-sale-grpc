# Shipping Address API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find All Shipping Addresses
GET {{baseUrl}}/api/shipping-address
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Shipping Address by ID
GET {{baseUrl}}/api/shipping-address/1
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Shipping Address by Order ID
GET {{baseUrl}}/api/shipping-address/order/1
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Active Shipping Addresses
GET {{baseUrl}}/api/shipping-address/active
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Trashed Shipping Addresses
GET {{baseUrl}}/api/shipping-address/trashed
Authorization: Bearer {{accessToken}}
HTTP 200

# Trash Shipping Address
POST {{baseUrl}}/api/shipping-address/trashed/1
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Shipping Address
POST {{baseUrl}}/api/shipping-address/restore/1
Authorization: Bearer {{accessToken}}
HTTP 200
