# Slider API Tests

# Login to get access token
POST {{baseUrl}}/api/auth/login
{
    "email": "test@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# Find All Sliders
GET {{baseUrl}}/api/slider
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Slider (Multipart)
POST {{baseUrl}}/api/slider/create
Authorization: Bearer {{accessToken}}
[MultipartFormData]
name: "Summer Sale"
image_slider: file,dummy.jpg;
HTTP 200
[Captures]
sliderId: jsonpath "$.data.id"

# Update Slider (Multipart)
POST {{baseUrl}}/api/slider/update/{{sliderId}}
Authorization: Bearer {{accessToken}}
[MultipartFormData]
name: "Summer Sale Updated"
image_slider: file,dummy.jpg;
HTTP 200

# Trash Slider
POST {{baseUrl}}/api/slider/trashed/{{sliderId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Slider
POST {{baseUrl}}/api/slider/restore/{{sliderId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Delete Slider Permanent
DELETE {{baseUrl}}/api/slider/permanent/{{sliderId}}
Authorization: Bearer {{accessToken}}
HTTP 200
