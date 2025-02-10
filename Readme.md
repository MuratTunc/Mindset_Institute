
# Mindset Institute Case Study

In this study, a total of 6 separate services running on Docker were designed. Development was made on Golang language as the Backend Software language due to its speed and efficiency, and on Ubuntu Linux as the operating system. With Bash shell script language and Makefile, services were created automatically from scratch and all API functions were tested automatically.


## Getting the project running on your PC

Download the project git repository to your local PC. Run the following commands in order.

_Curl, Docker, Go, Docker Compose features must be installed on your development PC. If they are not installed or you are not sure, you can run the prepare_development_pc.sh script to automatically install the necessary packages._

```bash
cd Mindset_Institute/build-tools/
sudo ./prepare_development_pc.sh
sudo make -s build
```

__*You should get terminal output like the one below.*__

```bash
mutu@mutu:~/projects/Mindset_Institute/build-tools$ sudo make -s build
[sudo] password for mutu: 
🧹 Cleaning up all database volumes...
Stopping customer-service      ... done
Stopping user-service          ... done
Stopping salestracking-service ... done
Stopping user-db               ... done
Stopping salestracking-db      ... done
Stopping customer-db           ... done
Removing customer-service      ... done
Removing user-service          ... done
Removing salestracking-service ... done
Removing user-db               ... done
Removing salestracking-db      ... done
Removing customer-db           ... done
Removing network build-tools_default
Removing volume build-tools_user_db_data
Removing volume build-tools_customer_db_data
Removing volume build-tools_salestracking_db_data
✅ All volumes removed!
🔍 Checking for running containers...
PostgreSQL is not running on port 5432.
1- 🚀 Building user-service binary...
/bin/sh: 11: ./.env: n!X2ZjzGp#nJ2k2ZoLs45!Vqa5m0F!ztr7@1f#Vjz1j: not found
✅ Done! user-service binary...
2- 🚀 Building customer-service binary...
/bin/sh: 11: ./.env: n!X2ZjzGp#nJ2k2ZoLs45!Vqa5m0F!ztr7@1f#Vjz1j: not found
✅ Done! customer-service binary...
3- 🚀 Building salestracking-service binary...
/bin/sh: 11: ./.env: n!X2ZjzGp#nJ2k2ZoLs45!Vqa5m0F!ztr7@1f#Vjz1j: not found
✅ Done! salestracking-service binary...
🚀 Building (when required) and starting docker images with environment variables...
Creating network "build-tools_default" with the default driver
Creating volume "build-tools_user_db_data" with default driver
Creating volume "build-tools_customer_db_data" with default driver
Creating volume "build-tools_salestracking_db_data" with default driver
Building user-service
[+] Building 2.3s (15/15) FINISHED                                                                                                                                       docker:default
 => [internal] load build definition from user-service.dockerfile                                                                                                                  0.0s
 => => transferring dockerfile: 746B                                                                                                                                               0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                                                                                   2.2s
 => [internal] load metadata for docker.io/library/golang:1.23-alpine                                                                                                              0.8s
 => [internal] load .dockerignore                                                                                                                                                  0.0s
 => => transferring context: 2B                                                                                                                                                    0.0s
 => [builder 1/6] FROM docker.io/library/golang:1.23-alpine@sha256:2c49857f2295e89b23b28386e57e018a86620a8fede5003900f2d138ba9c4037                                                0.0s
 => [internal] load build context                                                                                                                                                  0.0s
 => => transferring context: 370B                                                                                                                                                  0.0s
 => [stage-1 1/3] FROM docker.io/library/alpine:latest@sha256:56fa17d2a7e7f168a043a2712e63aed1f8543aeafdcee47c58dcffe38ed51099                                                     0.0s
 => CACHED [stage-1 2/3] WORKDIR /app                                                                                                                                              0.0s
 => CACHED [builder 2/6] WORKDIR /app                                                                                                                                              0.0s
 => CACHED [builder 3/6] COPY go.mod go.sum ./                                                                                                                                     0.0s
 => CACHED [builder 4/6] RUN go mod download                                                                                                                                       0.0s
 => CACHED [builder 5/6] COPY cmd/api /app/cmd/api                                                                                                                                 0.0s
 => CACHED [builder 6/6] RUN go build -o userServiceApp ./cmd/api                                                                                                                  0.0s
 => CACHED [stage-1 3/3] COPY --from=builder /app/userServiceApp .                                                                                                                 0.0s
 => exporting to image                                                                                                                                                             0.0s
 => => exporting layers                                                                                                                                                            0.0s
 => => writing image sha256:bd123c0c62631151699d0b5fa6c755be2860e588ac4f90a99f0c14b8e40f0dcb                                                                                       0.0s
 => => naming to docker.io/library/user-service-img                                                                                                                                0.0s
Building customer-service
[+] Building 0.6s (15/15) FINISHED                                                                                                                                       docker:default
 => [internal] load build definition from customer-service.dockerfile                                                                                                              0.0s
 => => transferring dockerfile: 762B                                                                                                                                               0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                                                                                   0.2s
 => [internal] load metadata for docker.io/library/golang:1.23-alpine                                                                                                              0.5s
 => [internal] load .dockerignore                                                                                                                                                  0.0s
 => => transferring context: 2B                                                                                                                                                    0.0s
 => [stage-1 1/3] FROM docker.io/library/alpine:latest@sha256:56fa17d2a7e7f168a043a2712e63aed1f8543aeafdcee47c58dcffe38ed51099                                                     0.0s
 => [builder 1/6] FROM docker.io/library/golang:1.23-alpine@sha256:2c49857f2295e89b23b28386e57e018a86620a8fede5003900f2d138ba9c4037                                                0.0s
 => [internal] load build context                                                                                                                                                  0.0s
 => => transferring context: 370B                                                                                                                                                  0.0s
 => CACHED [stage-1 2/3] WORKDIR /app                                                                                                                                              0.0s
 => CACHED [builder 2/6] WORKDIR /app                                                                                                                                              0.0s
 => CACHED [builder 3/6] COPY go.mod go.sum ./                                                                                                                                     0.0s
 => CACHED [builder 4/6] RUN go mod download                                                                                                                                       0.0s
 => CACHED [builder 5/6] COPY cmd/api /app/cmd/api                                                                                                                                 0.0s
 => CACHED [builder 6/6] RUN go build -o customerServiceApp ./cmd/api                                                                                                              0.0s
 => CACHED [stage-1 3/3] COPY --from=builder /app/customerServiceApp .                                                                                                             0.0s
 => exporting to image                                                                                                                                                             0.0s
 => => exporting layers                                                                                                                                                            0.0s
 => => writing image sha256:4a0d85831b11417021ae358d4af89702e9eb0c01dd6f88f97950781abf284a0c                                                                                       0.0s
 => => naming to docker.io/library/customer-service-img                                                                                                                            0.0s
Building salestracking-service
[+] Building 0.4s (15/15) FINISHED                                                                                                                                       docker:default
 => [internal] load build definition from salestracking-service.dockerfile                                                                                                         0.0s
 => => transferring dockerfile: 782B                                                                                                                                               0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                                                                                   0.3s
 => [internal] load metadata for docker.io/library/golang:1.23-alpine                                                                                                              0.2s
 => [internal] load .dockerignore                                                                                                                                                  0.0s
 => => transferring context: 2B                                                                                                                                                    0.0s
 => [stage-1 1/3] FROM docker.io/library/alpine:latest@sha256:56fa17d2a7e7f168a043a2712e63aed1f8543aeafdcee47c58dcffe38ed51099                                                     0.0s
 => [builder 1/6] FROM docker.io/library/golang:1.23-alpine@sha256:2c49857f2295e89b23b28386e57e018a86620a8fede5003900f2d138ba9c4037                                                0.0s
 => [internal] load build context                                                                                                                                                  0.0s
 => => transferring context: 369B                                                                                                                                                  0.0s
 => CACHED [stage-1 2/3] WORKDIR /app                                                                                                                                              0.0s
 => CACHED [builder 2/6] WORKDIR /app                                                                                                                                              0.0s
 => CACHED [builder 3/6] COPY go.mod go.sum ./                                                                                                                                     0.0s
 => CACHED [builder 4/6] RUN go mod download                                                                                                                                       0.0s
 => CACHED [builder 5/6] COPY cmd/api /app/cmd/api                                                                                                                                 0.0s
 => CACHED [builder 6/6] RUN go build -o salestrackingServiceApp ./cmd/api                                                                                                         0.0s
 => CACHED [stage-1 3/3] COPY --from=builder /app/salestrackingServiceApp .                                                                                                        0.0s
 => exporting to image                                                                                                                                                             0.0s
 => => exporting layers                                                                                                                                                            0.0s
 => => writing image sha256:07f08e19d2dd25c879749a90a9a4984f4663a410d3a467485a80e88de4200c17                                                                                       0.0s
 => => naming to docker.io/library/salestracking-service-img                                                                                                                       0.0s
Creating customer-db      ... done
Creating salestracking-db ... done
Creating user-db          ... done
Creating customer-service      ... done
Creating user-service          ... done
Creating salestracking-service ... done
✅ Docker images built and started!
⏳ Waiting for 5   seconds to allow services to initialize ..... ✅
✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅
📜 Fetching logs for all services...
Attaching to salestracking-service, user-service, customer-service, user-db, salestracking-db, customer-db
customer-service         | 🔧 Loaded Environment Variables - CUSTOMER_SERVICE
customer-service         | DBHost: customer-db
customer-service         | DBUser: customer
customer-service         | DBPassword: customer_password
customer-service         | DBName: customer_db
customer-service         | DBPort: 5432
customer-service         | ServicePort: 8081
customer-service         | ServiceName: CUSTOMER-SERVICE
customer-service         | ✅ DATABASE connection success!
customer-service         | 🚀 CUSTOMER-SERVICE is running on port: 8081
salestracking-db         | /usr/local/bin/docker-entrypoint.sh: ignoring /docker-entrypoint-initdb.d/*
salestracking-db         | 
salestracking-db         | waiting for server to shut down...2025-02-10 16:21:46.139 UTC [48] LOG:  received fast shutdown request
salestracking-db         | .2025-02-10 16:21:46.142 UTC [48] LOG:  aborting any active transactions
salestracking-db         | 2025-02-10 16:21:46.143 UTC [48] LOG:  background worker "logical replication launcher" (PID 54) exited with exit code 1
salestracking-db         | 2025-02-10 16:21:46.143 UTC [49] LOG:  shutting down
salestracking-db         | 2025-02-10 16:21:46.147 UTC [49] LOG:  checkpoint starting: shutdown immediate
salestracking-db         | 2025-02-10 16:21:46.201 UTC [49] LOG:  checkpoint complete: wrote 918 buffers (5.6%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.014 s, sync=0.031 s, total=0.058 s; sync files=301, longest=0.003 s, average=0.001 s; distance=4222 kB, estimate=4222 kB
salestracking-db         | 2025-02-10 16:21:46.206 UTC [48] LOG:  database system is shut down
salestracking-db         |  done
salestracking-db         | server stopped
salestracking-db         | 
salestracking-db         | PostgreSQL init process complete; ready for start up.
salestracking-db         | 
salestracking-db         | 2025-02-10 16:21:46.282 UTC [1] LOG:  starting PostgreSQL 15.10 (Debian 15.10-1.pgdg120+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 12.2.0-14) 12.2.0, 64-bit
salestracking-db         | 2025-02-10 16:21:46.282 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
salestracking-db         | 2025-02-10 16:21:46.282 UTC [1] LOG:  listening on IPv6 address "::", port 5432
salestracking-db         | 2025-02-10 16:21:46.293 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
salestracking-db         | 2025-02-10 16:21:46.306 UTC [64] LOG:  database system was shut down at 2025-02-10 16:21:46 UTC
salestracking-db         | 2025-02-10 16:21:46.319 UTC [1] LOG:  database system is ready to accept connections
customer-db              | /usr/local/bin/docker-entrypoint.sh: ignoring /docker-entrypoint-initdb.d/*
customer-db              | 
customer-db              | 2025-02-10 16:21:46.215 UTC [48] LOG:  received fast shutdown request
customer-db              | waiting for server to shut down....2025-02-10 16:21:46.218 UTC [48] LOG:  aborting any active transactions
customer-db              | 2025-02-10 16:21:46.223 UTC [48] LOG:  background worker "logical replication launcher" (PID 54) exited with exit code 1
customer-db              | 2025-02-10 16:21:46.223 UTC [49] LOG:  shutting down
customer-db              | 2025-02-10 16:21:46.228 UTC [49] LOG:  checkpoint starting: shutdown immediate
customer-db              | 2025-02-10 16:21:46.341 UTC [49] LOG:  checkpoint complete: wrote 918 buffers (5.6%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.029 s, sync=0.061 s, total=0.119 s; sync files=301, longest=0.006 s, average=0.001 s; distance=4222 kB, estimate=4222 kB
customer-db              | 2025-02-10 16:21:46.357 UTC [48] LOG:  database system is shut down
customer-db              |  done
customer-db              | server stopped
customer-db              | 
customer-db              | PostgreSQL init process complete; ready for start up.
customer-db              | 
customer-db              | 2025-02-10 16:21:46.469 UTC [1] LOG:  starting PostgreSQL 15.10 (Debian 15.10-1.pgdg120+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 12.2.0-14) 12.2.0, 64-bit
customer-db              | 2025-02-10 16:21:46.469 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
customer-db              | 2025-02-10 16:21:46.469 UTC [1] LOG:  listening on IPv6 address "::", port 5432
customer-db              | 2025-02-10 16:21:46.478 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
customer-db              | 2025-02-10 16:21:46.493 UTC [64] LOG:  database system was shut down at 2025-02-10 16:21:46 UTC
customer-db              | 2025-02-10 16:21:46.505 UTC [1] LOG:  database system is ready to accept connections
user-service             | 🔧 Loaded Environment Variables - USER_SERVICE
user-service             | DBHost: user-db
user-service             | DBUser: user
user-service             | DBPassword: user_password
user-service             | DBName: user_db
user-service             | DBPort: 5432
user-service             | ServicePort: 8080
user-service             | ServiceName: USER-SERVICE
user-service             | JWTSecret: 6$8fjZ2@sjKl#F8tTr1&n!X2ZjzGp#nJ2k2ZoLs45!Vqa5m0F!ztr7@1f#Vjz1j
user-service             | ✅ DATABASE connection success!
user-service             | 🚀 USER-SERVICE is running on port: 8080
salestracking-service    | 🔧 Loaded Environment Variables -SALESTRACKING_SERVICE
salestracking-service    | DBHost: salestracking-db
salestracking-service    | DBUser: salestracking
salestracking-service    | DBPassword: salestracking_password
salestracking-service    | DBName: salestracking_db
salestracking-service    | DBPort: 5432
salestracking-service    | ServicePort: 8082
salestracking-service    | ServiceName: SALESTRACKING-SERVICE
salestracking-service    | ✅ DATABASE connection success!
salestracking-service    | 🚀 SALESTRACKING-SERVICE is running on port: 8082
user-db                  | /usr/local/bin/docker-entrypoint.sh: ignoring /docker-entrypoint-initdb.d/*
user-db                  | 
user-db                  | waiting for server to shut down....2025-02-10 16:21:46.215 UTC [48] LOG:  received fast shutdown request
user-db                  | 2025-02-10 16:21:46.218 UTC [48] LOG:  aborting any active transactions
user-db                  | 2025-02-10 16:21:46.222 UTC [48] LOG:  background worker "logical replication launcher" (PID 54) exited with exit code 1
user-db                  | 2025-02-10 16:21:46.223 UTC [49] LOG:  shutting down
user-db                  | 2025-02-10 16:21:46.226 UTC [49] LOG:  checkpoint starting: shutdown immediate
user-db                  | 2025-02-10 16:21:46.338 UTC [49] LOG:  checkpoint complete: wrote 918 buffers (5.6%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.031 s, sync=0.061 s, total=0.116 s; sync files=301, longest=0.006 s, average=0.001 s; distance=4222 kB, estimate=4222 kB
user-db                  | 2025-02-10 16:21:46.354 UTC [48] LOG:  database system is shut down
user-db                  |  done
user-db                  | server stopped
user-db                  | 
user-db                  | PostgreSQL init process complete; ready for start up.
user-db                  | 
user-db                  | 2025-02-10 16:21:46.469 UTC [1] LOG:  starting PostgreSQL 15.10 (Debian 15.10-1.pgdg120+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 12.2.0-14) 12.2.0, 64-bit
user-db                  | 2025-02-10 16:21:46.469 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
user-db                  | 2025-02-10 16:21:46.469 UTC [1] LOG:  listening on IPv6 address "::", port 5432
user-db                  | 2025-02-10 16:21:46.481 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
user-db                  | 2025-02-10 16:21:46.495 UTC [64] LOG:  database system was shut down at 2025-02-10 16:21:46 UTC
user-db                  | 2025-02-10 16:21:46.508 UTC [1] LOG:  database system is ready to accept connections
🚀 Running Containers:
CONTAINER ID   IMAGE                       COMMAND                  CREATED          STATUS                    PORTS                                         NAMES
fb91183bd8bc   salestracking-service-img   "/app/salestrackingS…"   6 seconds ago    Up 5 seconds              0.0.0.0:8082->8082/tcp, :::8082->8082/tcp     salestracking-service
8d5f288dca7b   user-service-img            "/app/userServiceApp"    6 seconds ago    Up 5 seconds              0.0.0.0:8080->8080/tcp, :::8080->8080/tcp     user-service
e55ccbd1c2f3   customer-service-img        "/app/customerServic…"   6 seconds ago    Up 5 seconds              0.0.0.0:8081->8081/tcp, :::8081->8081/tcp     customer-service
1b4c74172afb   postgres:15                 "docker-entrypoint.s…"   17 seconds ago   Up 16 seconds (healthy)   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp     user-db
756e1794e1a9   postgres:15                 "docker-entrypoint.s…"   17 seconds ago   Up 16 seconds (healthy)   0.0.0.0:5434->5432/tcp, [::]:5434->5432/tcp   salestracking-db
78778f43a0a6   postgres:15                 "docker-entrypoint.s…"   17 seconds ago   Up 16 seconds (healthy)   0.0.0.0:5433->5432/tcp, [::]:5433->5432/tcp   customer-db
✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅

```

## Running Containers

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




All variables are taken parametrically from the .env file for modularity, reusability, efficiency and easy addition of new services.
![alt text](image-2.png)

# Explanation of `.env` File

This `.env` file is used to store environment variables that configure various services and database connections for your Golang microservices. Below is a detailed breakdown of each section.
Provides a centralized location for storing environment-specific configuration values for our services.
Helps our configure service ports, Docker container names, binary names, database connections, and other sensitive data like JWT secrets.


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

#### Backend SW Structure for all services:
![alt text](image-1.png)

### USER-SERVICE

User service is the service that communicates with the database where user admin information is kept and provides management. It starts the web service with port __8080__.


#### API Endpoints

###### Public Routes (No Authentication Required)
| Method | Endpoint      | Description |
|--------|-------------|-------------|
| `POST` | `/register`  | Registers a new user |
| `POST` | `/login`     | Logs in an existing user |
| `GET`  | `/health`    | Checks if the service is healthy |
| `GET`  | `/swagger/*` | Serves Swagger API documentation |

###### Protected Routes (Require JWT Authentication)
| Method  | Endpoint              | Description |
|---------|----------------------|-------------|
| `GET`   | `/user`               | Retrieves a user by their ID |
| `POST`  | `/update-password`    | Updates the user's password |
| `PUT`   | `/update-user`        | Updates user information |
| `PUT`   | `/deactivate-user`    | Deactivates a user by username |
| `PUT`   | `/activate-user`      | Activates a deactivated user |
| `PUT`   | `/update-email`       | Updates the user's email address |
| `PUT`   | `/update-role`        | Updates the user's role |
| `DELETE`| `/delete-user`        | Deletes a user |




### CUSTOMER-SERVICE

It is the service that listening to port __8081__ and transfers customer information that will make the purchase to the database and manages it.

#### Customer Service API Endpoints

#### Public Routes (No Authentication Required)
| Method | Endpoint       | Description |
|--------|--------------|-------------|
| `POST` | `/register`   | Registers a new customer |
| `POST` | `/login`      | Logs in an existing customer |
| `GET`  | `/health`     | Checks if the service is healthy |

##### Protected Routes (Require Authentication)
###### **Customer Management**
| Method  | Endpoint                | Description |
|---------|-------------------------|-------------|
| `GET`   | `/get_all_customer`      | Retrieves all customers |
| `GET`   | `/order-customers`       | Orders customers based on criteria |
| `GET`   | `/activated-customers`   | Retrieves activated customers |
| `GET`   | `/logged-in-customers`   | Retrieves logged-in customers |
| `GET`   | `/customer`              | Retrieves a specific customer by criteria |
| `PUT`   | `/update-customer`       | Updates customer information |
| `PUT`   | `/deactivate-customer`   | Deactivates a customer account |
| `PUT`   | `/activate-customer`     | Activates a customer account |
| `DELETE`| `/delete-customer`       | Deletes a customer |

###### **Customer Notes**
| Method  | Endpoint         | Description |
|---------|----------------|-------------|
| `PUT`   | `/update-note`  | Updates an existing customer note |
| `PUT`   | `/insert-note`  | Inserts a new note for a customer |

###### **Customer Account Updates**
| Method  | Endpoint            | Description |
|---------|---------------------|-------------|
| `POST`  | `/update-password`  | Updates the customer's password |
| `PUT`   | `/update-email`     | Updates the customer's email address |



### SALESTRACKING-SERVICE

It is a service that listens to __8082__ and enables the transfer, updating and management of purchased product information and product status information during sales processes to the database.

#### Sales Tracking Service API Endpoints

###### Public Routes (No Authentication Required)
| Method | Endpoint  | Description |
|--------|----------|-------------|
| `GET`  | `/health` | Checks if the service is healthy |

###### **Sales Management Endpoints**
| Method   | Endpoint                  | Description |
|----------|---------------------------|-------------|
| `POST`   | `/insert-sale`             | Inserts a new sale record |
| `DELETE` | `/delete-sale`             | Deletes an existing sale record |
| `PUT`    | `/update-incommunication`  | Updates the "in communication" status of a sale |
| `PUT`    | `/update-deal`             | Updates the "deal" status of a sale |
| `PUT`    | `/update-closed`           | Updates the "closed" status of a sale |


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

## Getting integration tests.
```bash
cd Mindset_Institute/integration_tests/
./integration_user_service.sh
./integration_customer_service.sh
./integration_salestracking_service.sh 
```



## Docker services logs

docker logs user-service
docker logs customer-service
docker logs salestracking-service

## Generating sample jwt-secret key
openssl rand -base64 32

## db analze from container
sudo docker exec -it <container id> psql -U user -d user_db


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