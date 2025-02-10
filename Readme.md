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




### API DOCUMENTATION

## USER-SERVICE API

### TEST END POINT--->HEALTH CHECK
```bash
REQUEST URL: http://localhost:8080/health
REQUEST TYPE: GET
COMMAND: curl -X GET "http://localhost:8080/health"
Health Check Response Body: OK
HTTP Status Code: 200
Service is healthy!
✅ Health Check successfully
```

### TEST END POINT-->REGISTER NEW USER
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

### TEST END POINT-->LOGIN USER
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

### TEST END POINT-->DEACTIVATE USER
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

### TEST END POINT-->ACTIVATE USER
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

### TEST END POINT-->UPDATE EMAIL ADDRESS
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
### TEST END POINT-->UPDATE NEW PASSWORD
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
### TEST END POINT-->UPDATE USER ROLE
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
### TEST END POINT-->UPDATE USER
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
### TEST END POINT-->DELETE USER
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

## CUSTOMER SERVICE API

### TEST END POINT--->HEALTH CHECK
```bash
REQUEST URL: http://localhost:8081/health
REQUEST TYPE: GET
COMMAND: curl -X GET "http://localhost:8081/health"
Health Check Response Body: OK
HTTP Status Code: 200
Service is healthy!
✅ Health Check successfully
```

### TEST END POINT--->REGISTER NEW CUSTOMER
```bash
REQUEST URL: http://localhost:8081/register
REQUEST TYPE: POST
JSON BODY: {
  "customername": "testcustomer",
  "mailAddress": "testcustomer@example.com",
  "password": "TestPassword123"
}
COMMAND: curl -X POST "http://localhost:8081/register" -H "Content-Type: application/json" -d '{
  "customername": "testcustomer",
  "mailAddress": "testcustomer@example.com",
  "password": "TestPassword123"
}'
Registration response: {"mailAddress":"testcustomer@example.com","message":"Customer created successfully"}
HTTP Status Code: 200
Customer registered successfully!
✅ Registration successful!
```

### TEST END POINT--->LOGIN CUSTOMER
```bash
REQUEST URL: http://localhost:8081/login
REQUEST TYPE: POST
JSON BODY: {
  "customername": "testcustomer",
  "password": "TestPassword123"
}
COMMAND: curl -X POST "http://localhost:8081/login" -H "Content-Type: application/json" -d '{
  "customername": "testcustomer",
  "password": "TestPassword123"
}'
Login response body: {"loginStatus":"true","message":"Login successful"}
HTTP Status Code: 200
Login successful.
✅ Login successful!

 id | customername |       mail_address       |                           password                           | activated | login_status | note |          created_at           |          updated_at           
----+--------------+--------------------------+--------------------------------------------------------------+-----------+--------------+------+-------------------------------+-------------------------------
  1 | testcustomer | testcustomer@example.com | $2a$10$VJGBI4cH4QMZqOJl.DY33ebzSxQHQQVRYSKAQJFofdIfGoACHpr1y | t         | t            |      | 2025-02-10 08:46:26.492897+00 | 2025-02-10 08:46:26.558239+00
(1 row)
```
### TEST END POINT--->DEACTIVATE CUSTOMER
```bash
REQUEST URL: http://localhost:8081/deactivate-customer
REQUEST TYPE: PUT
JSON BODY: {
  "customername": "testcustomer"
}
COMMAND: curl -X PUT "http://localhost:8081/deactivate-customer" -H "Authorization: Bearer " -H "Content-Type: application/json" -d "{
  "customername": "testcustomer"
}"
Deactivate response body: {"message":"Customer deactivated successfully"}
HTTP Status Code: 200
✅ Customer deactivated successfully.

 id | customername |       mail_address       |                           password                           | activated | login_status | note |          created_at           |          updated_at           
----+--------------+--------------------------+--------------------------------------------------------------+-----------+--------------+------+-------------------------------+-------------------------------
  1 | testcustomer | testcustomer@example.com | $2a$10$VJGBI4cH4QMZqOJl.DY33ebzSxQHQQVRYSKAQJFofdIfGoACHpr1y | f         | t            |      | 2025-02-10 08:46:26.492897+00 | 2025-02-10 08:46:26.661808+00
(1 row)
```
### TEST END POINT--->ACTIVATE CUSTOMER
```basH
REQUEST URL: http://localhost:8081/activate-customer
REQUEST TYPE: PUT
JSON BODY: {
  "customername": "testcustomer"
}
COMMAND: curl -X PUT "http://localhost:8081/activate-customer" -H "Authorization: Bearer " -H "Content-Type: application/json" -d "{
  "customername": "testcustomer"
}"
Activate response body: {"message":"Customer activated successfully"}
HTTP Status Code: 200
✅ Customer activated successfully.

 id | customername |       mail_address       |                           password                           | activated | login_status | note |          created_at           |          updated_at           
----+--------------+--------------------------+--------------------------------------------------------------+-----------+--------------+------+-------------------------------+-------------------------------
  1 | testcustomer | testcustomer@example.com | $2a$10$VJGBI4cH4QMZqOJl.DY33ebzSxQHQQVRYSKAQJFofdIfGoACHpr1y | t         | t            |      | 2025-02-10 08:46:26.492897+00 | 2025-02-10 08:46:26.774977+00
(1 row)
```
### TEST END POINT--->UPDATE EMAIL ADDRESS
```bash
REQUEST URL: http://localhost:8081/update-email
REQUEST TYPE: PUT
JSON BODY: {
  "customername": "testcustomer",
  "new_email": "newmail@example.com"
}
COMMAND: curl -X PUT "http://localhost:8081/update-email" -H "Content-Type: application/json" -d "{
  "customername": "testcustomer",
  "new_email": "newmail@example.com"
}"
Update email response body: {"message":"Email updated successfully"}
HTTP Status Code: 200
✅ Email updated successfully.

 id | customername |    mail_address     |                           password                           | activated | login_status | note |          created_at           |          updated_at           
----+--------------+---------------------+--------------------------------------------------------------+-----------+--------------+------+-------------------------------+-------------------------------
  1 | testcustomer | newmail@example.com | $2a$10$VJGBI4cH4QMZqOJl.DY33ebzSxQHQQVRYSKAQJFofdIfGoACHpr1y | t         | t            |      | 2025-02-10 08:46:26.492897+00 | 2025-02-10 08:46:26.883792+00
(1 row)
```
### TEST END POINT--->UPDATE PASSWORD
```bash
REQUEST URL: http://localhost:8081/update-password
REQUEST TYPE: POST
JSON BODY: {
  "customername": "testcustomer",
  "new_password": "NewTestPassword123"
}
COMMAND: curl -X POST "http://localhost:8081/update-password" -H "Authorization: Bearer " -H "Content-Type: application/json" -d "{
  "customername": "testcustomer",
  "new_password": "NewTestPassword123"
}"
Update password response body: Password updated successfully
HTTP Status Code: 200
✅ Password updated successfully.

 id | customername |    mail_address     |                           password                           | activated | login_status | note |          created_at           |          updated_at           
----+--------------+---------------------+--------------------------------------------------------------+-----------+--------------+------+-------------------------------+-------------------------------
  1 | testcustomer | newmail@example.com | $2a$10$KW2qLeYzR/5zORziOza/IuRwuC.3wGkbA4hVXIVALjjVZmajr58dq | t         | t            |      | 2025-02-10 08:46:26.492897+00 | 2025-02-10 08:46:27.146958+00
(1 row)
```
### TEST END POINT--->INSERT NOTE
```bash
REQUEST URL: http://localhost:8081/insert-note
REQUEST TYPE: PUT
JSON BODY: {
  "customername": "testcustomer",
  "new_note": "This is a new note to append."
}
COMMAND: curl -X PUT "http://localhost:8081/insert-note" -H "Content-Type: application/json" -d "{
  "customername": "testcustomer",
  "new_note": "This is a new note to append."
}"
Insert Note response body: {"message":"Note appended successfully"}
HTTP Status Code: 200
✅ Note inserted successfully.

 id | customername |    mail_address     |                           password                           | activated | login_status |             note              |          created_at           |          updated_at           
----+--------------+---------------------+--------------------------------------------------------------+-----------+--------------+-------------------------------+-------------------------------+-------------------------------
  1 | testcustomer | newmail@example.com | $2a$10$KW2qLeYzR/5zORziOza/IuRwuC.3wGkbA4hVXIVALjjVZmajr58dq | t         | t            | This is a new note to append. | 2025-02-10 08:46:26.492897+00 | 2025-02-10 08:46:27.249643+00
(1 row)
```
### TEST END POINT--->UPDATE NOTE
```bash
REQUEST URL: http://localhost:8081/update-note
REQUEST TYPE: PUT
JSON BODY: {
  "customername": "testcustomer",
  "note": "This is the completely new note."
}
COMMAND: curl -X PUT "http://localhost:8081/update-note" -H "Content-Type: application/json" -d "{
  "customername": "testcustomer",
  "note": "This is the completely new note."
}"
Update Note response body: {"message":"Note updated successfully"}
HTTP Status Code: 200
✅ Note updated successfully.

 id | customername |    mail_address     |                           password                           | activated | login_status |               note               |          created_at           |          updated_at          
----+--------------+---------------------+--------------------------------------------------------------+-----------+--------------+----------------------------------+-------------------------------+------------------------------
  1 | testcustomer | newmail@example.com | $2a$10$KW2qLeYzR/5zORziOza/IuRwuC.3wGkbA4hVXIVALjjVZmajr58dq | t         | t            | This is the completely new note. | 2025-02-10 08:46:26.492897+00 | 2025-02-10 08:46:27.34629+00
(1 row)
```
### TEST END POINT--->UPDATE CUSTOMER
```bash
REQUEST URL: http://localhost:8081/update-customer
REQUEST TYPE: PUT
JSON BODY: {
  "customername": "testcustomer",
  "mailAddress": "updatedcustomer@example.com"
}
COMMAND: curl -X PUT "http://localhost:8081/update-customer" -H "Authorization: Bearer " -H "Content-Type: application/json" -d "{
  "customername": "testcustomer",
  "mailAddress": "updatedcustomer@example.com"
}"
Update response body: Customer updated successfully
HTTP Status Code: 200
✅ Customer updated successfully.

 id | customername |    mail_address     |                           password                           | activated | login_status |               note               |          created_at           |          updated_at          
----+--------------+---------------------+--------------------------------------------------------------+-----------+--------------+----------------------------------+-------------------------------+------------------------------
  1 | testcustomer | newmail@example.com | $2a$10$KW2qLeYzR/5zORziOza/IuRwuC.3wGkbA4hVXIVALjjVZmajr58dq | t         | t            | This is the completely new note. | 2025-02-10 08:46:26.492897+00 | 2025-02-10 08:46:27.45609+00
(1 row)
```
### TEST END POINT--->ORDER CUSTOMERS
```bash
REQUEST URL: http://localhost:8081/order-customers
Testing ordering by created_at (default):
COMMAND: curl -X GET "http://localhost:8081/order-customers"
Response Body: [{"ID":1,"Customername":"testcustomer","MailAddress":"newmail@example.com","Password":"$2a$10$KW2qLeYzR/5zORziOza/IuRwuC.3wGkbA4hVXIVALjjVZmajr58dq","Activated":true,"LoginStatus":true,"Note":"This is the completely new note.","CreatedAt":"2025-02-10T08:46:26.492897Z","UpdatedAt":"2025-02-10T08:46:27.45609Z"}]
HTTP Status Code: 200
✅ Customers ordered by created_at.
Testing ordering by customername:
COMMAND: curl -X GET "http://localhost:8081/order-customers?order_by=customername"
Response Body: [{"ID":1,"Customername":"testcustomer","MailAddress":"newmail@example.com","Password":"$2a$10$KW2qLeYzR/5zORziOza/IuRwuC.3wGkbA4hVXIVALjjVZmajr58dq","Activated":true,"LoginStatus":true,"Note":"This is the completely new note.","CreatedAt":"2025-02-10T08:46:26.492897Z","UpdatedAt":"2025-02-10T08:46:27.45609Z"}]
HTTP Status Code: 200
✅ Customers ordered by customername.
Testing ordering by updated_at:
COMMAND: curl -X GET "http://localhost:8081/order-customers?order_by=updated_at"
Response Body: [{"ID":1,"Customername":"testcustomer","MailAddress":"newmail@example.com","Password":"$2a$10$KW2qLeYzR/5zORziOza/IuRwuC.3wGkbA4hVXIVALjjVZmajr58dq","Activated":true,"LoginStatus":true,"Note":"This is the completely new note.","CreatedAt":"2025-02-10T08:46:26.492897Z","UpdatedAt":"2025-02-10T08:46:27.45609Z"}]
HTTP Status Code: 200
✅ Customers ordered by updated_at.
Testing ordering with invalid 'order_by' field (should default):
COMMAND: curl -X GET "http://localhost:8081/order-customers?order_by=invalid_field"
Response Body: [{"ID":1,"Customername":"testcustomer","MailAddress":"newmail@example.com","Password":"$2a$10$KW2qLeYzR/5zORziOza/IuRwuC.3wGkbA4hVXIVALjjVZmajr58dq","Activated":true,"LoginStatus":true,"Note":"This is the completely new note.","CreatedAt":"2025-02-10T08:46:26.492897Z","UpdatedAt":"2025-02-10T08:46:27.45609Z"}]
HTTP Status Code: 200
✅ Customers ordered with invalid 'order_by' field (default applied).
```
### TEST END POINT--->GET ACTIVATED CUSTOMERS
```bash
REQUEST URL: http://localhost:8081/activated-customers
REQUEST TYPE: GET
COMMAND: curl -X GET "http://localhost:8081/activated-customers"
Response Body: ["testcustomer"]
HTTP Status Code: 200
✅ Successfully retrieved activated customer names.
```
### TEST END POINT--->GET LOGGED-IN CUSTOMERS
```bash
REQUEST URL: http://localhost:8081/logged-in-customers
REQUEST TYPE: GET
COMMAND: curl -X GET "http://localhost:8081/logged-in-customers" -H "Content-Type: application/json"
Response Body: [{"ID":1,"Customername":"testcustomer","MailAddress":"newmail@example.com","Password":"$2a$10$KW2qLeYzR/5zORziOza/IuRwuC.3wGkbA4hVXIVALjjVZmajr58dq","Activated":true,"LoginStatus":true,"Note":"This is the completely new note.","CreatedAt":"2025-02-10T08:46:26.492897Z","UpdatedAt":"2025-02-10T08:46:27.45609Z"}]
HTTP Status Code: 200
✅ Successfully retrieved logged-in customers.
```
### TEST END POINT--->DELETE CUSTOMER
```bash
REQUEST URL: http://localhost:8081/delete-customer
REQUEST TYPE: DELETE
JSON BODY: {
  "customername": "testcustomer"
}
COMMAND: curl -X DELETE "http://localhost:8081/delete-customer" -H "Content-Type: application/json" -d "{
  "customername": "testcustomer"
}"
Delete response body: Customer deleted successfully
HTTP Status Code: 200
✅ Customer deleted successfully.

 id | customername | mail_address | password | activated | login_status | note | created_at | updated_at 
----+--------------+--------------+----------+-----------+--------------+------+------------+------------
(0 rows)
```




## SALESTRACKING SERVICE API

### TEST END POINT--->HEALTH CHECK
```bash
REQUEST URL: http://localhost:8082/health
REQUEST TYPE: GET
COMMAND: curl -X GET "http://localhost:8082/health"
Health Check Response Body: OK
HTTP Status Code: 200
Service is healthy!
✅ Health Check successfully
```

### TEST END POINT--->INSERT NEW SALE RECORD
```bash
REQUEST URL: http://localhost:8082/insert-sale
REQUEST_TYPE: POST
JSON_BODY: {
    "salename": "TestSale123",
    "note": "This is a test note for the sale record."
  }
Curl Command: curl -X POST "http://localhost:8082/insert-sale" -H "Content-Type: application/json" -d '{
    "salename": "TestSale123",
    "note": "This is a test note for the sale record."
  }'
Insert sale response body: {"message":"Sale record created successfully"}
HTTP Status Code: 200
Sale record inserted successfully!
✅ INSERT NEW SALE RECORD successfully

 id |  salename   | new | in_communication | deal | closed |                   note                   |          created_at           |          updated_at           
----+-------------+-----+------------------+------+--------+------------------------------------------+-------------------------------+-------------------------------
  5 | TestSale123 | t   | f                | f    | f      | This is a test note for the sale record. | 2025-02-10 09:49:48.453175+00 | 2025-02-10 09:49:48.453175+00
(1 row)
```

### TEST END POINT--->UPDATE INCOMMUNICATION FIELD
```bash
Salename: TestSale123
InCommunication: true
Note: This is the completely new note.
URL: http://localhost:8082/update-incommunication
REQUEST_TYPE: PUT
JSON Payload: {
  "salename": "TestSale123",
  "in_communication": true,
  "note": "This is the completely new note."
}
Curl Command: curl -X PUT "http://localhost:8082/update-incommunication" -H "Content-Type: application/json" -d "{
  "salename": "TestSale123",
  "in_communication": true,
  "note": "This is the completely new note."
}"
Update response: {"message":"Sale record updated successfully"}
HTTP Status Code: 200
Sale record updated successfully!
✅ UPDATE INCOMMUNICATION successfully

 id |  salename   | new | in_communication | deal | closed |                   note                   |          created_at           |          updated_at           
----+-------------+-----+------------------+------+--------+------------------------------------------+-------------------------------+-------------------------------
  5 | TestSale123 | f   | t                | f    | f      | This is a test note for the sale record.+| 2025-02-10 09:49:48.453175+00 | 2025-02-10 09:49:48.641957+00
    |             |     |                  |      |        | This is the completely new note.         |                               | 
(1 row)
```

### TEST END POINT--->UPDATE INCOMMUNICATION FIELD
```bash
Salename: TestSale123
InCommunication: false
Note: This is the completely new note.
URL: http://localhost:8082/update-incommunication
REQUEST_TYPE: PUT
JSON Payload: {
  "salename": "TestSale123",
  "in_communication": false,
  "note": "This is the completely new note."
}
Curl Command: curl -X PUT "http://localhost:8082/update-incommunication" -H "Content-Type: application/json" -d "{
  "salename": "TestSale123",
  "in_communication": false,
  "note": "This is the completely new note."
}"
Update response: {"message":"Sale record updated successfully"}
HTTP Status Code: 200
Sale record updated successfully!
✅ UPDATE INCOMMUNICATION successfully

 id |  salename   | new | in_communication | deal | closed |                   note                   |          created_at           |          updated_at           
----+-------------+-----+------------------+------+--------+------------------------------------------+-------------------------------+-------------------------------
  5 | TestSale123 | t   | f                | f    | f      | This is a test note for the sale record.+| 2025-02-10 09:49:48.453175+00 | 2025-02-10 09:49:48.751244+00
    |             |     |                  |      |        | This is the completely new note.        +|                               | 
    |             |     |                  |      |        | This is the completely new note.         |                               | 
(1 row)
```

### TEST END POINT--->UPDATE DEAL FIELD
```bash
Salename: TestSale123
Deal: true
Note: This is the completely new note.
URL: http://localhost:8082/update-deal
REQUEST_TYPE: PUT
curl -s -w "%{http_code}" -X PUT "http://localhost:8082/update-deal" -H "Content-Type: application/json" -d '{
    "salename": "TestSale123",
    "deal": true,
    "note": "This is the completely new note."
  }'
Update response: {"message":"Sale record updated successfully"}

HTTP Status Code: 200
Sale record updated successfully!
✅ UPDATE DEAL FIELD successfully

 id |  salename   | new | in_communication | deal | closed |                   note                   |          created_at           |          updated_at           
----+-------------+-----+------------------+------+--------+------------------------------------------+-------------------------------+-------------------------------
  5 | TestSale123 | f   | f                | t    | f      | This is a test note for the sale record.+| 2025-02-10 09:49:48.453175+00 | 2025-02-10 09:49:48.865938+00
    |             |     |                  |      |        | This is the completely new note.        +|                               | 
    |             |     |                  |      |        | This is the completely new note.        +|                               | 
    |             |     |                  |      |        | This is the completely new note.         |                               | 
(1 row)
```

### TEST END POINT--->UPDATE CLOSED FIELD
```bash
Salename: TestSale123
Note: This is the completely new note.
URL: http://localhost:8082/update-closed
REQUEST_TYPE: PUT
curl -s -w "%{http_code}" -X PUT "http://localhost:8082/update-closed" -H "Content-Type: application/json" -d '{
    "salename": "TestSale123",
    "note": "This is the completely new note."
  }'
Update response: {"message":"Sale record closed successfully"}

HTTP Status Code: 200
Sale record closed successfully!
✅ UPDATE CLOSED FIELD successfully

 id |  salename   | new | in_communication | deal | closed |                   note                   |          created_at           |          updated_at           
----+-------------+-----+------------------+------+--------+------------------------------------------+-------------------------------+-------------------------------
  5 | TestSale123 | f   | f                | f    | t      | This is a test note for the sale record.+| 2025-02-10 09:49:48.453175+00 | 2025-02-10 09:49:48.969303+00
    |             |     |                  |      |        | This is the completely new note.        +|                               | 
    |             |     |                  |      |        | This is the completely new note.        +|                               | 
    |             |     |                  |      |        | This is the completely new note.        +|                               | 
    |             |     |                  |      |        | This is the completely new note.         |                               | 
(1 row)
```

### TEST END POINT---> DELETE SALE
```bash
SALE NAME: TestSale123
URL: http://localhost:8082/delete-sale
REQUEST_TYPE: DELETE
curl -s -w "%{http_code}" -X DELETE "http://localhost:8082/delete-sale" -H "Content-Type: application/json" -d '{
    "salename": "TestSale123"
  }'
Delete response: {"message":"Sale deleted successfully","salename":"TestSale123"}

HTTP Status Code: 200
Sale deleted successfully.
✅ DELETE SALE successfully

 id | salename | new | in_communication | deal | closed | note | created_at | updated_at 
----+----------+-----+------------------+------+--------+------+------------+------------
(0 rows)
```
