# Order API Tests

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

# Find All Orders
GET {{baseUrl}}/api/order
Authorization: Bearer {{accessToken}}
HTTP 200

# Create Order
POST {{baseUrl}}/api/order/create
Authorization: Bearer {{accessToken}}
{
    "merchant_id": 1,
    "user_id": {{userId}},
    "total_price": 150000,
    "items": [
        {
            "product_id": 1,
            "quantity": 2,
            "price": 50000
        },
        {
            "product_id": 2,
            "quantity": 1,
            "price": 50000
        }
    ],
    "shipping_address": {
        "alamat": "Jl. Merdeka No. 1",
        "provinsi": "Jawa Barat",
        "kota": "Bandung",
        "courier": "JNE",
        "shipping_method": "REG",
        "shipping_cost": 10000,
        "negara": "Indonesia"
    }
}
HTTP 200
[Captures]
orderId: jsonpath "$.data.id"

# Find Order by ID
GET {{baseUrl}}/api/order/{{orderId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Find Monthly Total Revenue
GET {{baseUrl}}/api/order/monthly-total-revenue?year=2024&month=4
Authorization: Bearer {{accessToken}}
HTTP 200

# Trash Order
POST {{baseUrl}}/api/order/trashed/{{orderId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# Restore Order
POST {{baseUrl}}/api/order/restore/{{orderId}}
Authorization: Bearer {{accessToken}}
HTTP 200
