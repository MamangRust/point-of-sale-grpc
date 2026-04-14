# Category API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find All Categories
GET {{baseUrl}}/api/category
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Category
POST {{baseUrl}}/api/category/create
Authorization: Bearer {{accessToken}}
[MultipartFormData]
name: "Electronics"
description: "Gadgets and more"
image_category: file,tests/hurl/dummy.jpg;
HTTP 201
[Captures]
categoryId: jsonpath "$.data.id"

# Find Category by ID
GET {{baseUrl}}/api/category/{{categoryId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Category
POST {{baseUrl}}/api/category/update/{{categoryId}}
Authorization: Bearer {{accessToken}}
[MultipartFormData]
name: "Electronics Updated"
description: "New description"
HTTP 200

# Find Monthly Total Pricing
GET {{baseUrl}}/api/category/monthly-total-pricing?year=2024&month=4
Authorization: Bearer {{accessToken}}
HTTP 200

# Trash Category
POST {{baseUrl}}/api/category/trashed/{{categoryId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Category
POST {{baseUrl}}/api/category/restore/{{categoryId}}
Authorization: Bearer {{accessToken}}
HTTP 200
