# Review API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find All Reviews
GET {{baseUrl}}/api/review
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Review
POST {{baseUrl}}/api/review/create
Authorization: Bearer {{accessToken}}
{
    "user_id": 1,
    "product_id": 1,
    "comment": "Great product!",
    "rating": 5
}
HTTP 200
[Captures]
reviewId: jsonpath "$.data.id"

# Find Reviews by Product
GET {{baseUrl}}/api/review/product/1
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Review
POST {{baseUrl}}/api/review/update/{{reviewId}}
Authorization: Bearer {{accessToken}}
{
    "comment": "Great product, updated!",
    "rating": 4
}
HTTP 200

# Trash Review
POST {{baseUrl}}/api/review/trashed/{{reviewId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Review
POST {{baseUrl}}/api/review/restore/{{reviewId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Delete Review Permanent
DELETE {{baseUrl}}/api/review/permanent/{{reviewId}}
Authorization: Bearer {{accessToken}}
HTTP 200
