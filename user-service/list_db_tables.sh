#!/bin/bash

# Set the container name
CONTAINER_NAME="user-db"

# Get the container ID using the container name
CONTAINER_ID=$(docker ps -qf "name=$CONTAINER_NAME")

# Check if the container exists
if [ -z "$CONTAINER_ID" ]; then
    echo "Error: No running container found with name '$CONTAINER_NAME'."
    exit 1
fi

# Set PostgreSQL details from .env
DB_NAME="user_db"  # Match the correct database name from  .env file
DB_USER="user"     # Database user from your .env file

# Run the query to list all rows in the 'users' table
docker exec -i "$CONTAINER_ID" psql -U "$DB_USER" -d "$DB_NAME" -c "SELECT * FROM users;"