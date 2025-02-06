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


