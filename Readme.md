## ✅ Running Containers Check:

    Microservices:
    ✅ user-service (Port: 8080)
    ✅ customer-service (Port: 8081)
    ✅ salestracking-service (Port: 8082)

    PostgreSQL Databases:
    ✅ user-db (Port: 5432)
    ✅ customer-db (Port: 5433)
    ✅ salestracking-db (Port: 5434)


## Health Check


curl -X GET http://localhost:8080/health
curl -X GET http://localhost:8081/health
curl -X GET http://localhost:8082/health


## Docker services logs

docker logs user-service
docker logs customer-service
docker logs salestracking-service

## Generating sample jwt-secret key
openssl rand -base64 32


## user-db
sudo docker exec -it <container id> psql -U user -d user_db



## see ttable
sudo docker exec -it e25ff201e765 bash
psql -U user -d user_db
SELECT * FROM users;


## All Databases Can Use 5432 Internally

Each database container must use 5432 internally because that's the default port for PostgreSQL inside Docker. However, they are still separate because each has a unique service name (host).


## The Important Part: Containers Use Service Names, Not Ports

Docker uses container names (service names in docker-compose.yml) instead of IPs/ports to allow services to talk to each other.

So your services connect like this:

    user-service connects to user-db at postgres://user:user_password@user-db:5432/user_db
    customer-service connects to customer-db at postgres://customer:customer_password@customer-db:5432/customer_db
    salestracking-service connects to salestracking-db at postgres://salestracking:salestracking_password@salestracking-db:5432/salestracking_db

# This is how Docker networks isolate them.

USER_POSTGRES_DB_PORT=5432
CUSTOMER_POSTGRES_DB_PORT=5433
SALESTRACKING_POSTGRES_DB_PORT=5434

This allows you to run:

psql -U user -h localhost -p 5432 -d user_db
psql -U customer -h localhost -p 5433 -d customer_db
psql -U salestracking -h localhost -p 5434 -d salestracking_db


## user-service unit test:

Explanation:

    TestCreateUserHandler:
        This test simulates a POST request to create a new user and checks that the response contains the success message and status code 200.

    TestLoginUserHandler:
        This test checks that the login endpoint returns a successful response (status code 200) and includes a JWT token in the response body.

    TestHealthCheckHandler:
        A basic test to check if the health check endpoint correctly responds with "OK" when the database is available.


    Mocking the Database:
        You'll need to use a mocking library or mock the GORM database interactions to ensure the tests are isolated and do not require an actual database connection.

    Other Tests:
        You can similarly add unit tests for other handler functions, such as UpdatePasswordHandler, UpdateEmailHandler, DeactivateUserHandler, and so on.

    Testing JWT Authentication:
        For functions that require JWT authentication, mock the JWT token validation and test different scenarios (e.g., valid/invalid tokens).



## Steps to Integrate Swagger:

Install Swagger in your project (if you haven’t already):

go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

Generate Swagger Docs:

swag init

Run Your Server and access Swagger UI at:

http://localhost:8080/swagger/index.html


### API DOCUMENTATION

## USER-SERVICE

### ===>TEST END POINT--->HEALTH CHECK
```bash
REQUEST URL: http://localhost:8080/health
REQUEST TYPE: GET
COMMAND: curl -X GET "http://localhost:8080/health"
Health Check Response Body: OK
HTTP Status Code: 200
Service is healthy!
✅ Health Check successfully
```

### ===>TEST END POINT-->REGISTER NEW USER
```bash
REQUEST URL: http://localhost:8080/register
REQUEST TYPE: POST
COMMAND: curl -X POST "http://localhost:8080/register" -H "Content-Type: application/json" -d '{
    "username": "testuser",
    "mailAddress": "testuser@example.com",
    "password": "TestPassword123",
    "role": "Admin"
  }'
Registration response: {"mailAddress":"testuser@example.com","message":"User created successfully","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkyNTgxNzYsInJvbGUiOiJBZG1pbiIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.qw4d0eoMUQmf1TaH-rViI5iJKPHKlwYrVr6kCKs2GtI"}
HTTP Status Code: 200
User registered successfully!
✅ Register New User User successfully

 id | username |     mail_address     |                           password                           | role  | activated | login_status |          created_at           |          updated_at           
----+----------+----------------------+--------------------------------------------------------------+-------+-----------+--------------+-------------------------------+-------------------------------
 12 | testuser | testuser@example.com | $2a$10$xW2C69tOT7SfWvYQe88BS..S57SXG91vHqHX19569esIqZvn5t4XC | Admin | t         | f            | 2025-02-10 07:16:16.791345+00 | 2025-02-10 07:16:16.791345+00
(1 row)
```

### ===>TEST END POINT-->LOGIN USER
```bash
REQUEST URL: http://localhost:8080/login
REQUEST TYPE: POST
COMMAND: curl -X POST "http://localhost:8080/login" -H "Content-Type: application/json" -d '{
    "username": "testuser",
    "password": "TestPassword123"
  }'
JSON BODY: {
    "username": "testuser",
    "password": "TestPassword123"
  }
Login response: {"loginStatus":"true","message":"Login successful","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkyNTgxNzYsInJvbGUiOiJBZG1pbiIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.qw4d0eoMUQmf1TaH-rViI5iJKPHKlwYrVr6kCKs2GtI"}
HTTP Status Code: 200
✅ Login successful. JWT token received.

 id | username |     mail_address     |                           password                           | role  | activated | login_status |          created_at           |          updated_at           
----+----------+----------------------+--------------------------------------------------------------+-------+-----------+--------------+-------------------------------+-------------------------------
 12 | testuser | testuser@example.com | $2a$10$xW2C69tOT7SfWvYQe88BS..S57SXG91vHqHX19569esIqZvn5t4XC | Admin | t         | t            | 2025-02-10 07:16:16.791345+00 | 2025-02-10 07:16:16.951427+00
(1 row)
```

### ===>TEST END POINT-->DEACTIVATE USER
```bash
REQUEST URL: http://localhost:8080/deactivate-user
JSON BODY: {
  "username": "testuser"
}
REQUEST TYPE: PUT
COMMAND: curl -X PUT "http://localhost:8080/deactivate-user" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkyNTgxNzYsInJvbGUiOiJBZG1pbiIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.qw4d0eoMUQmf1TaH-rViI5iJKPHKlwYrVr6kCKs2GtI" -H "Content-Type: application/json" -d '{
  "username": "testuser"
}'
Deactivate response: {"message":"User deactivated successfully","username":"testuser"}
HTTP Status Code: 200
✅ User deactivated successfully.

 id | username |     mail_address     |                           password                           | role  | activated | login_status |          created_at           |          updated_at           
----+----------+----------------------+--------------------------------------------------------------+-------+-----------+--------------+-------------------------------+-------------------------------
 12 | testuser | testuser@example.com | $2a$10$xW2C69tOT7SfWvYQe88BS..S57SXG91vHqHX19569esIqZvn5t4XC | Admin | f         | t            | 2025-02-10 07:16:16.791345+00 | 2025-02-10 07:16:17.068763+00
(1 row)
```

### ===>TEST END POINT-->ACTIVATE USER
```bash
REQUEST URL: http://localhost:8080/activate-user
JSON BODY: {
  "username": "testuser"
}
REQUEST TYPE: PUT
COMMAND: curl -X PUT "http://localhost:8080/activate-user" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkyNTgxNzYsInJvbGUiOiJBZG1pbiIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.qw4d0eoMUQmf1TaH-rViI5iJKPHKlwYrVr6kCKs2GtI" -H "Content-Type: application/json" -d '{
  "username": "testuser"
}'
Activate response: {"message":"User activated successfully","username":"testuser"}
HTTP Status Code: 200
✅ User activated successfully.

 id | username |     mail_address     |                           password                           | role  | activated | login_status |          created_at           |          updated_at           
----+----------+----------------------+--------------------------------------------------------------+-------+-----------+--------------+-------------------------------+-------------------------------
 12 | testuser | testuser@example.com | $2a$10$xW2C69tOT7SfWvYQe88BS..S57SXG91vHqHX19569esIqZvn5t4XC | Admin | t         | t            | 2025-02-10 07:16:16.791345+00 | 2025-02-10 07:16:17.245721+00
(1 row)
```

### ===>TEST END POINT-->UPDATE EMAIL ADDRESS
```bash
REQUEST URL: http://localhost:8080/update-email
JSON BODY: {
  "username": "testuser",
  "new_email": "newmail@example.com"
}
REQUEST TYPE: PUT
COMMAND: curl -X PUT "http://localhost:8080/update-email" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkyNTgxNzYsInJvbGUiOiJBZG1pbiIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.qw4d0eoMUQmf1TaH-rViI5iJKPHKlwYrVr6kCKs2GtI" -H "Content-Type: application/json" -d '{
  "username": "testuser",
  "new_email": "newmail@example.com"
}'
Update email response: {"message":"Email updated successfully","new_email":"newmail@example.com","username":"testuser"}

HTTP Status Code: 200
✅ Email updated successfully.

 id | username |    mail_address     |                           password                           | role  | activated | login_status |          created_at           |          updated_at           
----+----------+---------------------+--------------------------------------------------------------+-------+-----------+--------------+-------------------------------+-------------------------------
 12 | testuser | newmail@example.com | $2a$10$xW2C69tOT7SfWvYQe88BS..S57SXG91vHqHX19569esIqZvn5t4XC | Admin | t         | t            | 2025-02-10 07:16:16.791345+00 | 2025-02-10 07:16:17.343587+00
(1 row)
```
### ===>TEST END POINT-->UPDATE NEW PASSWORD
```bash
REQUEST URL: http://localhost:8080/update-password
JSON BODY: {
  "username": "testuser",
  "new_password": "NewTestPassword123"
}
REQUEST TYPE: POST
COMMAND: curl -X POST "http://localhost:8080/update-password" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkyNTgxNzYsInJvbGUiOiJBZG1pbiIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.qw4d0eoMUQmf1TaH-rViI5iJKPHKlwYrVr6kCKs2GtI" -H "Content-Type: application/json" -d '{
  "username": "testuser",
  "new_password": "NewTestPassword123"
}'
Update password response: Password updated successfully

HTTP Status Code: 200
✅ Password updated successfully.

 id | username |    mail_address     |                           password                           | role  | activated | login_status |          created_at           |          updated_at           
----+----------+---------------------+--------------------------------------------------------------+-------+-----------+--------------+-------------------------------+-------------------------------
 12 | testuser | newmail@example.com | $2a$10$1cG4hQxKE2eRXmPet05NWuzocy29U0NCEJh.zAf42u3lVKMXTB8iW | Admin | t         | t            | 2025-02-10 07:16:16.791345+00 | 2025-02-10 07:16:17.492229+00
(1 row)
```
### ===>TEST END POINT-->UPDATE USER ROLE
```bash
REQUEST URL: http://localhost:8080/update-role
JSON BODY: {
  "username": "testuser",
  "role": "MANAGER"
}
REQUEST TYPE: PUT
COMMAND: curl -X PUT "http://localhost:8080/update-role" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkyNTgxNzYsInJvbGUiOiJBZG1pbiIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.qw4d0eoMUQmf1TaH-rViI5iJKPHKlwYrVr6kCKs2GtI" -H "Content-Type: application/json" -d '{
  "username": "testuser",
  "role": "MANAGER"
}'
Update role response: User role updated to: MANAGER
HTTP Status Code: 200
✅ Role updated successfully.

 id | username |    mail_address     |                           password                           |  role   | activated | login_status |          created_at           |          updated_at           
----+----------+---------------------+--------------------------------------------------------------+---------+-----------+--------------+-------------------------------+-------------------------------
 12 | testuser | newmail@example.com | $2a$10$1cG4hQxKE2eRXmPet05NWuzocy29U0NCEJh.zAf42u3lVKMXTB8iW | MANAGER | t         | t            | 2025-02-10 07:16:16.791345+00 | 2025-02-10 07:16:17.603123+00
(1 row)
```
### ===>TEST END POINT-->UPDATE USER
```bash
REQUEST URL: http://localhost:8080/update-user
JSON BODY: {
  "username": "testuser",
  "email": "",
  "role": "Admin"
}
REQUEST TYPE: PUT
COMMAND: curl -X PUT "http://localhost:8080/update-user" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkyNTgxNzYsInJvbGUiOiJBZG1pbiIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.qw4d0eoMUQmf1TaH-rViI5iJKPHKlwYrVr6kCKs2GtI" -H "Content-Type: application/json" -d '{
  "username": "testuser",
  "email": "",
  "role": "Admin"
}'
Update response: {"message":"User updated successfully","username":"testuser"}
HTTP Status Code: 200
✅ User updated successfully.

 id | username |    mail_address     |                           password                           | role  | activated | login_status |          created_at           |          updated_at           
----+----------+---------------------+--------------------------------------------------------------+-------+-----------+--------------+-------------------------------+-------------------------------
 12 | testuser | newmail@example.com | $2a$10$1cG4hQxKE2eRXmPet05NWuzocy29U0NCEJh.zAf42u3lVKMXTB8iW | Admin | t         | t            | 2025-02-10 07:16:16.791345+00 | 2025-02-10 07:16:17.702148+00
(1 row)
```
### ===>TEST END POINT-->DELETE USER
```bash
REQUEST URL: http://localhost:8080/delete-user
JSON BODY: {
  "username": "testuser"
}
REQUEST TYPE: DELETE
COMMAND: curl -X DELETE "http://localhost:8080/delete-user" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkyNTgxNzYsInJvbGUiOiJBZG1pbiIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.qw4d0eoMUQmf1TaH-rViI5iJKPHKlwYrVr6kCKs2GtI" -H "Content-Type: application/json" -d '{
  "username": "testuser"
}'
Delete response: User deleted successfully
HTTP Status Code: 200
✅ User deleted successfully.

 id | username | mail_address | password | role | activated | login_status | created_at | updated_at 
----+----------+--------------+----------+------+-----------+--------------+------------+------------
(0 rows)
```

ALL TESTS ARE DONE!!!

