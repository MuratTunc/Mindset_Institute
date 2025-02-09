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

# Test: REGISTER NEW USER


- **REQUEST URL**: `http://localhost:8080/register`
- **REQUEST TYPE**: `POST`
- **COMMAND**:
  ```bash
  curl -X POST "http://localhost:8080/register" -H "Content-Type: application/json" -d '{
    "username": "testuser",
    "mailAddress": "testuser@example.com",
    "password": "TestPassword123",
    "role": "Admin"
  }'

