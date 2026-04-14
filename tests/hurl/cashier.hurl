# Cashier Lifecycle Test

# 1. Login to get token
POST {{baseUrl}}/api/auth/login
{
    "email": "john.doe@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# 2. Setup Merchant (Assuming merchant exists or use auth me to get user id)
GET {{baseUrl}}/api/auth/me
Authorization: Bearer {{accessToken}}
HTTP 200
[Captures]
userId: jsonpath "$.data.user_id"

# Create Merchant for testing if needed, or assume one exists from migrations/other tests
# For Hurl tests we usually expect a clean state but some seeded data
# Let's try to find an existing merchant or create one
GET {{baseUrl}}/api/merchant?page=1&page_size=1
Authorization: Bearer {{accessToken}}
HTTP 200
[Captures]
merchantId: jsonpath "$.data[0].id"

# 3. Create Cashier
POST {{baseUrl}}/api/cashier/create
Authorization: Bearer {{accessToken}}
{
    "name": "Hurl Cashier",
    "merchant_id": {{merchantId}},
    "user_id": {{userId}}
}
HTTP 201
[Captures]
cashierId: jsonpath "$.data.id"

# 4. Get Cashier
GET {{baseUrl}}/api/cashier/{{cashierId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# 5. Update Cashier
POST {{baseUrl}}/api/cashier/update/{{cashierId}}
Authorization: Bearer {{accessToken}}
{
    "name": "Hurl Cashier Updated"
}
HTTP 200

# 6. List Cashiers
GET {{baseUrl}}/api/cashier?page=1&page_size=10
Authorization: Bearer {{accessToken}}
HTTP 200

# 7. Trash Cashier
POST {{baseUrl}}/api/cashier/trashed/{{cashierId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# 8. List Trashed Cashiers
GET {{baseUrl}}/api/cashier/trashed?page=1&page_size=10
Authorization: Bearer {{accessToken}}
HTTP 200

# 9. Restore Cashier
POST {{baseUrl}}/api/cashier/restore/{{cashierId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# 10. Delete Permanent
POST {{baseUrl}}/api/cashier/trashed/{{cashierId}}
Authorization: Bearer {{accessToken}}
HTTP 200

DELETE {{baseUrl}}/api/cashier/permanent/{{cashierId}}
Authorization: Bearer {{accessToken}}
HTTP 200
