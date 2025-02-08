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

This is how Docker networks isolate them.


