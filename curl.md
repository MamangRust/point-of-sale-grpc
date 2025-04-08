## Auth

### Login
```sh
curl -X POST http://localhost:5000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "securepassword"
  }'
```

### Register

```sh
curl -X POST http://localhost:5000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "firstname": "John",
    "lastname": "Doe",
    "email": "john.doe@example.com",
    "password": "securepassword",
    "confirm_password": "securepassword"
  }'
```

### Refresh Token

```sh
curl -X POST \
  http://localhost:5000/api/auth/refresh-token \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1Nzg2MTYxfQ.yEx98MCuT0fg8b63VuLl9XcPxszYG2BTlQtRVvEsMbI' \
  -d '{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJyZWZyZXNoIl0sImV4cCI6MTczNTc4NjE2MX0.Ti5BTb8xMbMUYDNE-vFU8MVbr6o7zQLWJ-CIetByFd4"
}'
```

### GetMe

```sh
curl -X GET http://localhost:5000/api/auth/me \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'

```

## Role


### Find All
```sh
curl -X GET "http://localhost:5000/api/role \
 -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
```

### Find Id
```sh
curl -X GET "http://localhost:5000/api/role/1 \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```


### Find Active
```sh
curl -X GET "http://localhost:5000/api/role/active \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```


### Find Trashed
```sh
curl -X GET "http://localhost:5000/api/role/trashed \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```


### Find User Id
```sh
curl -X GET "http://localhost:5000/api/role/user/123 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Create Role
```sh
curl -X POST "http://localhost:5000/api/role" \
     -H "Content-Type: application/json" \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8' \
     -d '{"name": "Admin"}'
```

### Update Role
```sh
curl -X POST "http://localhost:5000/api/role/1" \
     -H "Content-Type: application/json" \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8' \
     -d '{"id": 1, "name": "Super Admin"}'
```

### Trashed Role
```sh
curl -X DELETE "http://localhost:5000/api/role/1 \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Restore Role
```sh
curl -X PUT "http://localhost:5000/api/role/restore/1 \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```


### Delete Permenant Role

```sh
curl -X DELETE "http://localhost:5000/api/role/permanent/1 \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Restore All Role
```sh
curl -X PUT "http://localhost:5000/api/role/restore-all \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Delete All Role
```sh
curl -X DELETE "http://localhost:5000/api/role/permanent-all \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

## User


### Find All
```sh
curl -X GET "http://localhost:5000/api/user \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Find ID
```sh
curl -X GET "http://localhost:5000/api/user/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Find Active
```sh
curl -X GET "http://localhost:5000/api/user/active \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Find Trashed
```sh
curl -X GET "http://localhost:5000/api/user/trashed \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Create User
```sh
curl -X POST "http://localhost:5000/api/user/create" \
     -H "Content-Type: application/json" \
      -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8' \
     -d '{"name": "John Doe", "email": "john@example.com", "password": "securepassword"}'
```

### Update User
```sh
curl -X POST "http://localhost:5000/api/user/update/1" \
     -H "Content-Type: application/json" \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'\
     -d '{"name": "John Updated", "email": "john.updated@example.com"}'
```

### Trashed User
```sh
curl -X POST "http://localhost:5000/api/user/trashed/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Restore user
```sh
curl -X POST "http://localhost:5000/api/user/restore/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Delete Permanent User
```sh
curl -X DELETE "http://localhost:5000/api/user/permanent/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Restore All User
```sh
curl -X POST "http://localhost:5000/api/user/restore/all \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Delete All User
```sh
curl -X POST "http://localhost:5000/api/user/permanent/all \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```



## Category

### Find All

```sh
curl -X GET "http://localhost:5000/api/category \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'

```

### Find Id

```sh
curl -X GET "http://localhost:5000/api/category/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Find Active

```sh
curl -X GET "http://localhost:5000/api/category/active \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
```

### Find Trashed

```sh
curl -X GET "http://localhost:5000/api/category/trashed \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
```


### Create Category 
```sh
curl -X POST "http://localhost:5000/api/category/create" \
-H "Content-Type: application/json" \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8' \
-d '{
    "name": "Electronics",
    "description": "This is a category for electronic devices"
}'
```


### Update Category

```sh
curl -X POST "http://localhost:5000/api/category/update/1" \
-H "Content-Type: application/json" \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8' \
-d '{
    "name": "Updated Electronics",
    "description": "This is an updated category for electronic devices"
}'
```

### Trashed Category

```sh
curl -X POST "http://localhost:5000/api/category/trashed/1 \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```


### Restore Category

```sh
curl -X POST "http://localhost:5000/api/category/restore/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Delete Category Permanent

```sh
curl -X DELETE "http://localhost:5000/api/category/permanent/1 \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```


### Restore All Category

```sh
curl -X POST "http://localhost:5000/api/category/restore/all \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Delete All Category

```sh
curl -X POST "http://localhost:5000/api/category/permanent/all \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Month Total Pricing
```sh
curl -X GET "http://localhost:5000/api/category/monthly-total-pricing?year=2025&month=4" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Year Total Pricing
```sh
curl -X GET "http://localhost:5000/api/category/yearly-total-pricing?year=2025" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Month Total Pricing Category
```sh
curl -X GET "http://localhost:5000/api/category/mycategory/monthly-total-pricing?year=2025&&month=4&category_id=2" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Year Total Pricing Category
```sh
curl -X GET "http://localhost:5000/api/category/mycategory/yearly-total-pricing?year=2025&category_id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```


### Month Total Pricing Merchant
```sh
curl -X GET "http://localhost:5000/api/category/merchant/yearly-total-pricing?year=2025&merchant_id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Year Total Pricing Merchant
```sh
curl -X GET "http://localhost:5000/api/category/merchant/monthly-total-pricing?year=2025&&month=4&category_id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```


### Month Pricing
```sh
curl -X GET "http://localhost:5000/api/category/monthly-pricing?year=2025" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year Pricing
```sh
curl -X GET "http://localhost:5000/api/category/yearly-pricing?year=2025" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Month Pricing Merchant
```sh
curl -X GET "http://localhost:5000/api/category/merchant/monthly-pricing?year=2025&merchant_id=1" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year Pricing Merchant
```sh
curl -X GET "http://localhost:5000/api/category/merchant/yearly-pricing?year=2025&merchant_id=1" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Month Pricing Category
```sh
curl -X GET "http://localhost:5000/api/category/mycategory/monthly-pricing?year=2025&category_id=1" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year Pricing Category
```sh
curl -X GET "http://localhost:5000/api/category/mycategory/yearly-pricing?year=2025&category_id=1" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```



## Merchant


### Find All
```sh
curl -X GET "http://localhost:5000/api/merchant \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Find ID
```sh
curl -X GET "http://localhost:5000/api/merchant/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Find Active
```sh
curl -X GET "http://localhost:5000/api/merchant/active \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Find Trashed
```sh
curl -X GET "http://localhost:5000/api/merchant/trashed?page=1&page_size=10&search=store \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Create Merchant
```sh
curl -X POST "http://localhost:5000/api/merchant/create" \
     -H "Content-Type: application/json" \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8' \
     -d '{
           "user_id": 123,
           "name": "Toko Sukses",
           "description": "Menjual berbagai kebutuhan sehari-hari",
           "address": "Jl. Merdeka No. 45",
           "contact_email": "tokosukses@example.com",
           "contact_phone": "08123456789",
           "status": "active"
         }'
```


### Update Merchant Id
```sh
curl -X POST "http://localhost:5000/api/merchant/update/1" \
     -H "Content-Type: application/json" \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8' \
     -d '{
           "merchant_id": 1,
           "user_id": 123,
           "name": "Toko Makmur",
           "description": "Menjual sembako dan kebutuhan rumah tangga",
           "address": "Jl. Merdeka No. 99",
           "contact_email": "tokomakmur@example.com",
           "contact_phone": "08129876543",
           "status": "inactive"
         }'
```


### Trashed merchant
```sh
curl -X POST "http://localhost:5000/api/merchant/trashed/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Restore merchant
```sh
curl -X POST "http://localhost:5000/api/merchant/restore/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Delete Permanent merchant
```sh
curl -X DELETE "http://localhost:5000/api/merchant/permanent/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Restore all merchant
```sh
curl -X POST "http://localhost:5000/api/merchant/restore/all \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Delete all merchant
```sh
curl -X POST "http://localhost:5000/api/merchant/permanent/all \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```




## Cashier


### Find All
```sh
curl -X GET "http://localhost:5000/api/cashier \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Find ID
```sh
curl -X GET "http://localhost:5000/api/cashier/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Find Active
```sh
curl -X GET "http://localhost:5000/api/cashier/active \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Find Trashed
```sh
curl -X GET "http://localhost:5000/api/cashier/trashed \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Create Cashier
```sh
curl -X POST "http://localhost:5000/api/cashier/create" \
     -H "Content-Type: application/json" \
      -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8' \
     -d '{
           "merchant_id": 1,
           "user_id": 123,
           "name": "John Doe"
         }'
```

### Update cashier by ID
```sh
curl -X POST "http://localhost:5000/api/cashier/update/1" \
     -H "Content-Type: application/json" \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8' \
     -d '{
           "cashier_id": 1,
           "name": "John Updated"
         }'
```


### Trashed Cashier
```sh
curl -X POST "http://localhost:5000/api/cashier/trashed/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Restore Cashier
```sh
curl -X POST "http://localhost:5000/api/cashier/restore/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Delete Permanent Cashier
```sh
curl -X DELETE "http://localhost:5000/api/cashier/permanent/1 \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```

### Restore All Cashier
```sh
curl -X POST "http://localhost:5000/api/cashier/restore/all \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```


### Delete All Cashier
```sh
curl -X POST "http://localhost:5000/api/cashier/permanent/all \
-H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzM1ODI5MDk0fQ.0MAChuYO1G458hK_HqVmFYAOdOnmeYkqFbTjbY0QDi8'
"
```


### Month Total Sales
```sh
curl -X GET "http://localhost:5000/api/cashier/monthly-total-sales?year=2025&&month=4" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Year Total Sales
```sh
curl -X GET "http://localhost:5000/api/cashier/yearly-total-sales?year=2025" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Month Total Sales Merchant
```sh
curl -X GET "http://localhost:5000/api/cashier/merchant/monthly-total-sales?year=2025&&month=4&merchant_id=2" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Year Total Sales Merchant
```sh
curl -X GET "http://localhost:5000/api/cashier/merchant/yearly-total-sales?year=2025&merchant_id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Month Total Sales Cashier
```sh
curl -X GET "http://localhost:5000/api/cashier/mycashier/monthly-total-sales?year=2025&&month=4&cashier_id=2" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Year Total Sales Cashier
```sh
curl -X GET "http://localhost:5000/api/cashier/mycashier/yearly-total-sales?year=2025&cashier_id=2" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Month Sales
```sh
curl -X GET "http://localhost:5000/api/cashier/monthly-sales?year=2025" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Year Sales
```sh
curl -X GET "http://localhost:5000/api/cashier/yearly-sales?year=2025" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Monht Sales Cashier
```sh
curl -X GET "http://localhost:5000/api/cashier/mycashier/monthly-sales?year=2025&cashier_id=3" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year Sales Cashier
```sh
curl -X GET "http://localhost:5000/api/cashier/mycashier/yearly-sales?year=2025&cashier_id=3" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Month Sales Merchant
```sh
curl -X GET "http://localhost:5000/api/cashier/merchant/monthly-sales?year=2025&merchant_id=4" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year Sales Merchant
```sh
curl -X GET "http://localhost:5000/api/cashier/merchant/yearly-sales?year=2025&merchant_id=4" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


## Product

### Find All
```sh
curl -X GET "http://localhost:5000/api/product \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
"
```

### Find ID
```sh
curl -X GET "http://localhost:5000/api/product/1 \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
"
```

### Find Product By Merchant
```sh
curl -X GET "http://localhost:5000/api/product/merchant/5 \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
" 

```

### Find Product By Category 
```sh
curl -X GET "http://localhost:5000/api/product/category/electronics \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
"
```

### Find Active
```sh
curl -X GET "http://localhost:5000/api/product/active \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
"
```

### Find Trashed
```sh
curl -X GET "http://localhost:5000/api/product/trashed \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
"
```

### Create Product
```sh
curl -X POST "http://localhost:5000/api/product/create" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ" \
     -H "Content-Type: multipart/form-data" \
     -F "merchant_id=5" \
     -F "category_id=2" \
     -F "name=Shampoo Herbal" \
     -F "description=Shampoo herbal alami tanpa bahan kimia" \
     -F "price=50000" \
     -F "count_in_stock=100" \
     -F "brand=HerbalCare" \
     -F "weight=200" \
     -F "rating=5" \
     -F "slug_product=shampoo-herbal" \
     -F "barcode=123456789" \
     -F "image_product=@/path/to/image.jpg"
```

### Update Product
```sh
curl -X POST "http://localhost:5000/api/product/update/1" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ" \
     -H "Content-Type: multipart/form-data" \
     -F "product_id=1" \
     -F "merchant_id=5" \
     -F "category_id=2" \
     -F "name=Shampoo Herbal Premium" \
     -F "description=Shampoo herbal dengan tambahan ekstrak lidah buaya" \
     -F "price=55000" \
     -F "count_in_stock=90" \
     -F "brand=HerbalCare" \
     -F "weight=250" \
     -F "rating=5" \
     -F "slug_product=shampoo-herbal-premium" \
     -F "barcode=987654321" \
     -F "image_product=@/path/to/image.jpg"  
```


### Trashed Product
```sh
curl -X POST "http://localhost:5000/api/product/trashed/1 \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Restore Product
```sh
curl -X POST "http://localhost:5000/api/product/restore/1 \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
"
```

### Delete Permanent Product
```sh
curl -X DELETE "http://localhost:5000/api/product/permanent/1 \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
"
```

### Restore All Product
```sh
curl -X POST "http://localhost:5000/api/product/restore/all \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
"
```

### Delete All Product
```sh
curl -X POST "http://localhost:5000/api/product/permanent/all \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
"
```


## Order

### Find All
```sh
curl -X GET "http://localhost:5000/api/order" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Find ID

```sh
curl -X GET "http://localhost:5000/api/order/1" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Find Active
```sh
curl -X GET "http://localhost:5000/api/order/active" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Find Trashed

```sh
curl -X GET "http://localhost:5000/api/order/trashed" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


### Create Order
```sh
curl -X POST "http://localhost:5000/api/order/create" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ" \
     -H "Content-Type: application/json" \
     -d '{
           "merchant_id": 1,
           "cashier_id": 2,
           "total_price": 150000,
           "items": [
             {
               "product_id": 101,
               "quantity": 2,
               "price": 50000
             },
             {
               "product_id": 102,
               "quantity": 1,
               "price": 50000
             }
           ]
         }'
```


### Update Order
```sh
curl -X POST "http://localhost:5000/api/order/update/1" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ" \
     -H "Content-Type: application/json" \
     -d '{
           "order_id": 1,
           "total_price": 140000,
           "items": [
             {
               "order_item_id": 201,
               "product_id": 101,
               "quantity": 1,
               "price": 50000
             },
             {
               "order_item_id": 202,
               "product_id": 103,
               "quantity": 2,
               "price": 45000
             }
           ]
         }'
```


### Trashed Order

```sh
curl -X POST "http://localhost:5000/api/order/trashed/1" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


### Restore Order

```sh
curl -X POST "http://localhost:5000/api/order/restore/1" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


### Delete Permanent Order

```sh
curl -X DELETE "http://localhost:5000/api/order/permanent/1" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


### Restore All Order

```sh
curl -X POST "http://localhost:5000/api/order/restore/all" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Delete All Order

```sh
curl -X POST "http://localhost:5000/api/order/permanent/all" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Month Total Revenue
```sh
curl -X GET "http://localhost:5000/api/order/monthly-total-revenue?year=2025&month=4" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```
### Year Total Revenue
```sh
curl -X GET "http://localhost:5000/api/order/yearly-total-revenue?year=2025" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Month Total Revenue Merchant
```sh
curl -X GET "http://localhost:5000/api/order/merchant/monthly-total-revenue?year=2025&month=4&merchant_id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Year Total Revenue Merchant
```sh

curl -X GET "http://localhost:5000/api/order/merchant/yearly-total-revenue?year=2025&merchant_id=1" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzOTU1NzU1fQ.XLVUuVp6cOL_ugSBOL7zbrR_jilsLsM4lRUWHwVApuM"
```

### Month Revenue 
```sh
curl -X GET "http://localhost:5000/api/order/monthly-revenue?year=2025" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year  Revenue 
```sh
curl -X GET "http://localhost:5000/api/order/yearly-revenue?year=2025" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"

```

### Month Revenue Merchant
```sh
curl -X GET "http://localhost:5000/api/order/merchant/monthly-revenue?year=2025&merchant_id=4" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year  Revenue Merchant
```sh
curl -X GET "http://localhost:5000/api/order/merchant/yearly-revenue?year=2025&merchant_id=4" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


## Order Item


###  Find All
```sh
curl -X GET "http://localhost:5000/api/order-item" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Find Order ID

```sh
curl -X GET "http://localhost:5000/api/order-item/1" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Find Active

```sh
curl -X GET "http://localhost:5000/api/order-item/active" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Find Trashed

```sh
curl -X GET "http://localhost:5000/api/order-item/trashed" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


## Transaction

### Find All

```sh
curl -X GET "http://localhost:5000/api/transaction" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


### Find ID

```sh
curl -X GET "http://localhost:5000/api/transaction/merchant/5" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Find Transaction By Merchant

```sh
curl -X GET "http://localhost:5000/api/transaction/merchant/5" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Find Active

```sh
curl -X GET "http://localhost:5000/api/transaction/active" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Find Trashed

```sh
curl -X GET "http://localhost:5000/api/transaction/trashed" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Create Transaction

```sh
curl -X POST "http://localhost:5000/api/transaction/create" \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ" \
     -d '{
           "order_id": 10,
           "merchant_id": 5,
           "payment_method": "credit_card",
           "amount": 250000,
           "change_amount": 0,
           "payment_status": "pending"
         }'
```

### Update Transaction
```sh
curl -X POST "http://localhost:5000/api/transaction/update/1" \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ" \
     -d '{
           "transaction_id": 1,
           "order_id": 10,
           "merchant_id": 5,
           "payment_method": "debit_card",
           "amount": 255000,
           "change_amount": 5000,
           "payment_status": "completed"
         }'
```

### Trashed Transaction

```sh
curl -X POST "http://localhost:5000/api/transaction/trashed/1" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


### Restore Transaction
```sh
curl -X POST "http://localhost:5000/api/transaction/restore/1" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Restore All Transaction
```sh
curl -X POST "http://localhost:5000/api/transaction/restore/all" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Delete All Transaction
```sh
curl -X POST "http://localhost:5000/api/transaction/permanent/all" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


### Month Success

```sh
curl -X GET "http://localhost:5000/api/transaction/monthly-success?year=2025&month=4" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year Success

```sh
curl -X GET "http://localhost:5000/api/transaction/yearly-success?year=2025" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Month Failed

```sh
curl -X GET "http://localhost:5000/api/transaction/monthly-failed?year=2025&month=4" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year Failed
```sh
curl -X GET "http://localhost:5000/api/transaction/yearly-failed?year=2025" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Month Success Merchant

```sh
curl -X GET "http://localhost:5000/api/transaction/monthly-success?year=2025&month=4&merchant_id=4" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year Success Merchant
```sh
curl -X GET "http://localhost:5000/api/transaction/yearly-success?year=2025&merchant_id=4" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Month Failed Merchant
```sh
curl -X GET "http://localhost:5000/api/transaction/monthly-failed?year=2025&month=4&merchant_id=4" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year Failed Merchant
```sh
curl -X GET "http://localhost:5000/api/transaction/yearly-failed?year=2025&merchant_id=4" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Month Method
```sh
curl -X GET "http://localhost:5000/api/transaction/monthly-methods?year=2025" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


### Year Method
```sh
curl -X GET "http://localhost:5000/api/transaction/yearly-methods?year=2025" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```


### Month Method Merchant
```sh
curl -X GET "http://localhost:5000/api/transaction/merchant/monthly-methods?year=2025&merchant_id=4" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```

### Year Method Merchant
```sh
curl -X GET "http://localhost:5000/api/transaction/merchant/yearly-methods?year=2025&merchant_id=4" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsImF1ZCI6WyJhY2Nlc3MiXSwiZXhwIjoxNzQzNjk2OTgwfQ.Zd0HzzrBOHn06nnTB4ClRYrIrwDYKc1Q5Imp3lqjxyQ"
```
