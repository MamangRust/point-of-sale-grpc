# Merchant Detail API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find All Merchant Details
GET {{baseUrl}}/api/merchant-detail
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Merchant Detail (Multipart)
POST {{baseUrl}}/api/merchant-detail/create
Authorization: Bearer {{accessToken}}
[MultipartFormData]
merchant_id: 1
display_name: "Fresh Foods Market"
short_description: "Quality fresh produce"
website_url: "http://freshfoods.com"
cover_image_url: file,dummy.jpg;
logo_url: file,dummy.jpg;
social_links: [{"platform": "instagram", "url": "https://instagram.com/freshfoods"}]
HTTP 200
[Captures]
detailId: jsonpath "$.data.id"

# Find Merchant Detail by ID
GET {{baseUrl}}/api/merchant-detail/{{detailId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Update Merchant Detail (Multipart)
POST {{baseUrl}}/api/merchant-detail/update/{{detailId}}
Authorization: Bearer {{accessToken}}
[MultipartFormData]
display_name: "Fresh Foods Market Updated"
short_description: "Premium quality fresh produce"
website_url: "http://freshfoods.org"
social_links: [{"id": 1, "platform": "facebook", "url": "https://facebook.com/freshfoods", "merchant_detail_id": 1}]
HTTP 200

# Trash Merchant Detail
POST {{baseUrl}}/api/merchant-detail/trashed/{{detailId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Merchant Detail
POST {{baseUrl}}/api/merchant-detail/restore/{{detailId}}
Authorization: Bearer {{accessToken}}
HTTP 200
