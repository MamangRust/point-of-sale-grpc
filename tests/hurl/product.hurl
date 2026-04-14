# Product API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"
userId: jsonpath "$.data.user.id"

# Find All Products
GET {{baseUrl}}/api/product
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Product by ID
# Assuming product ID 1 exists for this test, or we'd capture it from a create call.
# However, create requires an image upload which is tricky in Hurl without a real file.
GET {{baseUrl}}/api/product/1
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Products by Merchant
GET {{baseUrl}}/api/product/merchant/1
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Products by Category
GET {{baseUrl}}/api/product/category/electronics
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Active Products
GET {{baseUrl}}/api/product/active
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Trashed Products
GET {{baseUrl}}/api/product/trashed
Authorization: Bearer {{accessToken}}
HTTP 200
