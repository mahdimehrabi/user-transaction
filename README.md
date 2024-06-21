## Setup 
```
cp env.example .env
```
config environment variables in .env 

run migrations
```
make db-migrate-up
```



### REST API 
Create User 
```
curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d '{
        "name": "John Doe",
        "email": "john@example.com",
        "password": "password123"
    }'
```