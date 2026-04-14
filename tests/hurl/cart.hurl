# Cart API Tests

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

# Find All Carts for User
GET {{baseUrl}}/api/cart?user_id={{userId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Cart Item
POST {{baseUrl}}/api/cart/create
Authorization: Bearer {{accessToken}}
{
    "user_id": {{userId}},
    "product_id": 1,
    "quantity": 2
}
HTTP 201
[Captures]
cartId: jsonpath "$.data.id"

# Delete Cart Item
DELETE {{baseUrl}}/api/cart/{{cartId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Delete All Carts
POST {{baseUrl}}/api/cart/delete-all
Authorization: Bearer {{accessToken}}
{
    "cart_ids": [1, 2, 3]
}
HTTP 200
