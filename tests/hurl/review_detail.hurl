# Review Detail API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find All Review Details
GET {{baseUrl}}/api/review-detail
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Review Detail (Multipart)
POST {{baseUrl}}/api/review-detail/create
Authorization: Bearer {{accessToken}}
[MultipartFormData]
review_id: 1
type: "image"
caption: "Product image"
url: file,dummy.jpg;
HTTP 200
[Captures]
detailId: jsonpath "$.data.id"

# Find By ID
GET {{baseUrl}}/api/review-detail/{{detailId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Review Detail (Multipart)
POST {{baseUrl}}/api/review-detail/update/{{detailId}}
Authorization: Bearer {{accessToken}}
[MultipartFormData]
type: "image"
caption: "Updated product image"
url: file,dummy.jpg;
HTTP 200

# Trash Review Detail
POST {{baseUrl}}/api/review-detail/trashed/{{detailId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Review Detail
POST {{baseUrl}}/api/review-detail/restore/{{detailId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Delete Review Detail Permanent
DELETE {{baseUrl}}/api/review-detail/permanent/{{detailId}}
Authorization: Bearer {{accessToken}}
HTTP 200
