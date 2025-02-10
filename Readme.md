


# ✅ Running Containers Check:

In this study, a total of 6 separate services running on Docker were designed. Development was made on Golang language as the Backend Software language due to its speed and efficiency, and on Ubuntu Linux as the operating system. With Bash shell script language and Makefile, services were created automatically from scratch and all API functions were tested automatically.

![alt text](image.png)



Each micro service has its own postgres database.

    Microservices:
    ✅ user-service (Port: 8080)
    ✅ customer-service (Port: 8081)
    ✅ salestracking-service (Port: 8082)

    PostgreSQL Databases:
    ✅ user-db (Port: 5432)
    ✅ customer-db (Port: 5433)
    ✅ salestracking-db (Port: 5434)

#### Backend SW Structure for all services:
![alt text](image-1.png)


All variables are taken parametrically from the .env file for modularity, reusability, efficiency and easy addition of new services.
![alt text](image-2.png)

# Explanation of `.env` File

This `.env` file is used to store environment variables that configure various services and database connections for your Golang microservices. Below is a detailed breakdown of each section.

### Golang Config
GOFULLPATH: Specifies the full path to the Go binary. This is important for your Makefile or other scripts that need to know where Go is installed.
```bash
GOFULLPATH=/usr/local/go/bin/go
```

### User Service Config

USER_SERVICE_PORT=8080
USER_SERVICE_NAME=USER-SERVICE
USER_SERVICE_IMAGE_NAME=user-service-img
USER_SERVICE_CONTAINER_NAME=user-service
USER_SERVICE_BINARY=userServiceApp
USER_SERVICE_JWT_SECRET=6$8fjZ2@sjKl#F8tTr1&n!X2ZjzGp#nJ2k2ZoLs45!Vqa5m0F!ztr7@1f#Vjz1j


    USER_SERVICE_PORT: Port number for the user service to listen on.
    USER_SERVICE_NAME: The name of the user service.
    USER_SERVICE_IMAGE_NAME: The Docker image name for the user service.
    USER_SERVICE_CONTAINER_NAME: The name of the Docker container for the user service.
    USER_SERVICE_BINARY: The name of the compiled Go binary for the user service.
    USER_SERVICE_JWT_SECRET: A secret key used for signing JWT tokens in the user service (important for authentication).


### Customer Service Config

CUSTOMER_SERVICE_PORT=8081
CUSTOMER_SERVICE_NAME=CUSTOMER-SERVICE
CUSTOMER_SERVICE_IMAGE_NAME=customer-service-img
CUSTOMER_SERVICE_CONTAINER_NAME=customer-service
CUSTOMER_SERVICE_BINARY=customerServiceApp

    CUSTOMER_SERVICE_PORT: Port number for the customer service.
    CUSTOMER_SERVICE_NAME: The name of the customer service.
    CUSTOMER_SERVICE_IMAGE_NAME: Docker image name for the customer service.
    CUSTOMER_SERVICE_CONTAINER_NAME: Docker container name for the customer service.
    CUSTOMER_SERVICE_BINARY: Name of the compiled Go binary for the customer service.

### Sales Tracking Service Config

SALESTRACKING_SERVICE_PORT=8082
SALESTRACKING_SERVICE_NAME=SALESTRACKING-SERVICE
SALESTRACKING_SERVICE_IMAGE_NAME=salestracking-service-img
SALESTRACKING_SERVICE_CONTAINER_NAME=salestracking-service
SALESTRACKING_SERVICE_BINARY=salestrackingServiceApp

    SALESTRACKING_SERVICE_PORT: Port number for the sales tracking service.
    SALESTRACKING_SERVICE_NAME: Name of the sales tracking service.
    SALESTRACKING_SERVICE_IMAGE_NAME: Docker image name for the sales tracking service.
    SALESTRACKING_SERVICE_CONTAINER_NAME: Docker container name for the sales tracking service.
    SALESTRACKING_SERVICE_BINARY: Name of the compiled Go binary for the sales tracking service.

### User Service Database Config

USER_POSTGRES_DB_HOST=user-db
USER_POSTGRES_DB_PORT=5432
USER_POSTGRES_DB_USER=user
USER_POSTGRES_DB_PASSWORD=user_password
USER_POSTGRES_DB_NAME=user_db
USER_POSTGRES_DB_CONTAINER_NAME=user-db

    USER_POSTGRES_DB_HOST: The hostname or IP address of the PostgreSQL database for the user service.
    USER_POSTGRES_DB_PORT: The port number on which the PostgreSQL service is running for the user service.
    USER_POSTGRES_DB_USER: Username for the user service's PostgreSQL database.
    USER_POSTGRES_DB_PASSWORD: Password for the user service's PostgreSQL database.
    USER_POSTGRES_DB_NAME: The name of the database for the user service.
    USER_POSTGRES_DB_CONTAINER_NAME: Docker container name for the user service's PostgreSQL database.

### Customer Service Database Config

CUSTOMER_POSTGRES_DB_HOST=customer-db
CUSTOMER_POSTGRES_DB_PORT=5433
CUSTOMER_POSTGRES_DB_USER=customer
CUSTOMER_POSTGRES_DB_PASSWORD=customer_password
CUSTOMER_POSTGRES_DB_NAME=customer_db
CUSTOMER_POSTGRES_DB_CONTAINER_NAME=customer-db

    This section follows the same structure as the User Service Database Config, but it's for the customer service database.

### Sales Tracking Service Database Config

SALESTRACKING_POSTGRES_DB_HOST=salestracking-db
SALESTRACKING_POSTGRES_DB_PORT=5434
SALESTRACKING_POSTGRES_DB_USER=salestracking
SALESTRACKING_POSTGRES_DB_PASSWORD=salestracking_password
SALESTRACKING_POSTGRES_DB_NAME=salestracking_db
SALESTRACKING_POSTGRES_DB_CONTAINER_NAME=salestracking-db

    This section follows the same structure as the previous two, but it's for the sales tracking service database.

_Purpose of This .env File_

This .env file:

    Provides a centralized location for storing environment-specific configuration values for your services.
    Helps you configure service ports, Docker container names, binary names, database connections, and other sensitive data like JWT secrets.


### Makefile Purpose:
Summary:

This Makefile is a comprehensive tool for managing the build, testing, and cleanup of your services. It includes:

    Build targets for user, customer, and sales tracking services.
    Integration tests that can be customized with a wait time parameter.
    Docker container management (starting, stopping, removing).
    Integration test execution after the Docker containers have been successfully built and started.

It uses make to automate service building, waiting, and testing, with flexibility to adjust wait times between steps.

### Docker-Compose Yaml file Purpose:
Purpose of This YAML File:

    Defines services and dependencies: It sets up three services (user-service, customer-service, salestracking-service), each of which depends on its corresponding PostgreSQL database (user-db, customer-db, salestracking-db).

    Docker Image Build and Container Setup:
        Each service (user-service, customer-service, salestracking-service) is built from a specified directory and Dockerfile.
        The containers for each service are set to restart automatically (restart: always).
        Each service is bound to its respective database container, ensuring that the service waits for the database to be healthy before starting (depends_on and condition: service_healthy).

    Environment Variables:
        Environment variables are loaded from a .env file for each service. This allows sensitive information like database credentials and service ports to be easily configured.

    Health Checks:
        Health checks are configured for each database to ensure that PostgreSQL is ready and accessible before starting the related services.

    Persistent Storage:
        Volumes (user_db_data, customer_db_data, and salestracking_db_data) are defined for each database to store their data persistently.

How It Works:

    The services are built and run in Docker containers.
    The containers for the user, customer, and sales tracking services depend on their respective PostgreSQL database containers.
    Health checks ensure that the databases are ready before the services can start, which helps in coordinating the startup sequence.
    Volumes are used for persisting database data to ensure data durability across container restarts.


# BACK-END-SERVICES

### USER-SERVIS

User service is the service that communicates with the database where user admin information is kept and provides management. It starts the web service with port __8080__.

### CUSTOMER-SERVIS

It is the service that listening to port __8081__ and transfers customer information that will make the purchase to the database and manages it.

### SALESTRACKING-SERVIS

It is a service that listens to __8082__ and enables the transfer, updating and management of purchased product information and product status information during sales processes to the database.


# DATABASE TABLE FIELDS

### USER-DB
## User Model Breakdown

| Field        | Type        | GORM Tag                     | Description |
|-------------|------------|-----------------------------|-------------|
| `ID`        | `uint`      | `gorm:"primaryKey"`         | Auto-incremented primary key |
| `Username`  | `string`    | `gorm:"unique;not null"`    | Must be unique and cannot be null |
| `MailAddress` | `string`  | `gorm:"unique;not null"`    | Must be unique and cannot be null |
| `Password`  | `string`    | `gorm:"not null"`           | Cannot be null (hashed in DB) |
| `Role`      | `string`    | `gorm:"not null"`           | Can be `"Admin"` or `"Sales Representative"` |
| `Activated` | `bool`      | `gorm:"default:false"`      | Defaults to `false` (user is inactive by default) |
| `LoginStatus` | `bool`    | `gorm:"default:false"`      | Tracks whether the user is logged in |
| `CreatedAt` | `time.Time` | `gorm:"autoCreateTime"`     | Automatically set when the user is created |
| `UpdatedAt` | `time.Time` | `gorm:"autoUpdateTime"`     | Automatically updates when the user data is modified |



## CUSTOMER-DB

## Customer Model Breakdown

| Field          | Type        | GORM Tag                     | Description |
|--------------|------------|-----------------------------|-------------|
| `ID`         | `uint`      | `gorm:"primaryKey"`         | Auto-incremented primary key |
| `Customername` | `string`  | `gorm:"unique;not null"`    | Must be unique and cannot be null |
| `MailAddress`  | `string`  | `gorm:"unique;not null"`    | Must be unique and cannot be null |
| `Password`    | `string`   | `gorm:"not null"`           | Cannot be null (hashed in DB) |
| `Activated`   | `bool`     | `gorm:"default:false"`      | Defaults to `false` (customer is inactive by default) |
| `LoginStatus` | `bool`     | `gorm:"default:false"`      | Tracks whether the customer is logged in |
| `Note`        | `string`   | `gorm:"type:text"`          | Stores additional text information |
| `CreatedAt`   | `time.Time` | `gorm:"autoCreateTime"`     | Automatically set when the customer is created |
| `UpdatedAt`   | `time.Time` | `gorm:"autoUpdateTime"`     | Automatically updates when the customer data is modified |

## SALESTRACKING-DB

## Sale Model Breakdown

| Field            | Type        | GORM Tag                                | Description |
|----------------|------------|----------------------------------------|-------------|
| `ID`          | `uint`      | `gorm:"primaryKey"`                    | Auto-incremented primary key |
| `Salename`    | `string`    | `gorm:"type:varchar(255);uniqueIndex;not null"` | Unique sale name (max 255 chars), cannot be null |
| `New`         | `bool`      | `gorm:"default:false"`                  | Indicates if the sale is new |
| `InCommunication` | `bool`  | `gorm:"default:false"`                  | Tracks if the sale is in communication phase |
| `Deal`        | `bool`      | `gorm:"default:false"`                  | Indicates if a deal has been made |
| `Closed`      | `bool`      | `gorm:"default:false"`                  | Indicates if the sale is closed |
| `Note`        | `string`    | `gorm:"type:text"`                      | Stores additional text information about the sale |
| `CreatedAt`   | `time.Time` | `gorm:"autoCreateTime"`                 | Automatically set when the sale is created |
| `UpdatedAt`   | `time.Time` | `gorm:"autoUpdateTime"`                 | Automatically updates when the sale data is modified |




# API DOCUMENTATION

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


