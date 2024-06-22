## Setup 
```
cp env.example .env
```
config environment variables in .env 

run migrations
```
make db-migrate-up
```



## REST API Doc (I wanted to implement swagger too but I changed my mind because I didn't had enough time on friday🥴 )
### Authentication & JWT
Login 
```
curl -X POST http://localhost:8080/api/auth/login \
     -H "Content-Type: application/json" \
     -d '{
           "email": "user@example.com",
           "password": "password123",
           "remember": true
         }'
```

Renew Access Token 
```

```

Verify access token 
```
curl -X POST http://localhost:8080/api/auth/access-token-verify \
     -H "Content-Type: application/json" \
     -d '{
           "accessToken": "your-access-token"
         }'
```


### User
Create User 
```
curl -X POST http://localhost:8080/api/users \
    -H "Content-Type: application/json" \
    -d '{
        "name": "John Doe",
        "email": "john@example.com",
        "password": "password123"
    }'
```

Get user with ID 

```
curl -X GET http://localhost:8080/api/users/1
```


Update user with id 
```
curl -X PUT http://localhost:8080/api/users/1 \
    -H "Content-Type: application/json" \
    -d '{
        "name": "John Doe Updated",
        "email": "john.updated@example.com",
        "password": "newpassword123"
    }'
```


Delete User with id
```
curl -X DELETE http://localhost:8080/api/users/1
```


Get users with pagination 
```
curl -X GET "http://localhost:8080/users?page=1&pageSize=10"
```



### Extra Features 

- Complete unit tests for user service in /domain/service/user_test.go
- makefile 
- environment variables file
- password encryption
- general response generator for cleaner responses