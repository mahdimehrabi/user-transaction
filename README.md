### Checkout Makefile for database migrations

### Create user and addresses using file
```go run ./cmd/main.go -file=users_data.json```

# Create a user with its addresses using http
```
curl -X POST \
http://localhost:8080/api/users \
-H 'Content-Type: application/json' \
-d '{
"name": "John",
"lastname": "Doe",
"addresses": {
"street": "123 Main St",
"city": "Anytown",
"state": "Anystate",
"zipCode": "12345",
"country": "USA"
}
}' 
```

# GET user data
``` curl -X GET http://localhost:8080/api/users/58870 ```
