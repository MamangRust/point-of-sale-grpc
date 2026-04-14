# Role Lifecycle Test

# 1. Login to get token
POST {{baseUrl}}/api/auth/login
{
    "email": "john.doe@example.com",
    "password": "password123"
}
HTTP 200
[Captures]
accessToken: jsonpath "$.data.access_token"

# 2. Create Role
POST {{baseUrl}}/api/role
Authorization: Bearer {{accessToken}}
{
    "name": "Hurl Role"
}
HTTP 200
[Captures]
roleId: jsonpath "$.data.id"

# 3. Get Role
GET {{baseUrl}}/api/role/{{roleId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# 4. Update Role
POST {{baseUrl}}/api/role/update/{{roleId}}
Authorization: Bearer {{accessToken}}
{
    "name": "Hurl Role Updated"
}
HTTP 200

# 5. List Roles
GET {{baseUrl}}/api/role?page=1&page_size=10
Authorization: Bearer {{accessToken}}
HTTP 200

# 6. Trash Role
POST {{baseUrl}}/api/role/trashed/{{roleId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# 7. List Trashed Roles
GET {{baseUrl}}/api/role/trashed?page=1&page_size=10
Authorization: Bearer {{accessToken}}
HTTP 200

# 8. Restore Role
POST {{baseUrl}}/api/role/restore/{{roleId}}
Authorization: Bearer {{accessToken}}
HTTP 200

# 9. Delete Permanent
POST {{baseUrl}}/api/role/trashed/{{roleId}}
Authorization: Bearer {{accessToken}}
HTTP 200

DELETE {{baseUrl}}/api/role/permanent/{{roleId}}
Authorization: Bearer {{accessToken}}
HTTP 200
