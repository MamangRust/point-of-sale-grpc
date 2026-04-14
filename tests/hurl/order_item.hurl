# Order Item API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find All Order Items
GET {{baseUrl}}/api/order-item
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Order Items by Order ID
GET {{baseUrl}}/api/order-item/1
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Active Order Items
GET {{baseUrl}}/api/order-item/active
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Trashed Order Items
GET {{baseUrl}}/api/order-item/trashed
Authorization: Bearer {{accessToken}}
HTTP 200
