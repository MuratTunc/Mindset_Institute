#!/bin/bash

# Install jq (JSON parsing utility) if not already installed
if ! command -v jq &> /dev/null
then
  echo "jq could not be found, installing..."
  sudo apt-get update
  sudo apt-get install -y jq
fi

# Define user details
USERNAME="testuser3"
MAILADDRESS="testuser3@example.com"
PASSWORD="TestPassword123"
NEW_PASSWORD="NewTestPassword123"
ROLE="Admin"

# Define API URLs
BASE_URL="http://localhost:8080"
REGISTER_URL="$BASE_URL/register"
LOGIN_URL="$BASE_URL/login"
USER_URL="$BASE_URL/user"
UPDATE_USER_URL="$BASE_URL/user"
UPDATE_PASSWORD_URL="$BASE_URL/update-password"
DELETE_USER_URL="$BASE_URL/user"

# Function to check if the user exists
check_user_exists() {
  echo "Checking if user exists..."

  RESPONSE=$(curl -s -X GET "$USER_URL?username=$USERNAME")

  # Validate if the response contains a user ID
  if echo "$RESPONSE" | jq empty > /dev/null 2>&1; then
    USER_ID=$(echo "$RESPONSE" | jq -r '.ID')

    if [[ "$USER_ID" != "null" && -n "$USER_ID" ]]; then
      echo "User already exists with ID: $USER_ID"
      return 0
    fi
  fi

  echo "User does not exist."
  return 1
}

# Function to register the user
register_user() {
  echo "Registering new user..."

  REGISTER_RESPONSE=$(curl -s -w "\n%{http_code}" -X POST "$REGISTER_URL" -H "Content-Type: application/json" -d '{
    "username": "'$USERNAME'",
    "mailAddress": "'$MAILADDRESS'",
    "password": "'$PASSWORD'",
    "role": "'$ROLE'"
  }')

  HTTP_BODY=$(echo "$REGISTER_RESPONSE" | sed '$ d')
  HTTP_STATUS=$(echo "$REGISTER_RESPONSE" | tail -n1)

  echo "Registration response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  if [ "$HTTP_STATUS" -eq 201 ]; then
    echo "User registered successfully!"
  elif [ "$HTTP_STATUS" -eq 409 ]; then
    echo "User already exists. Proceeding with login..."
  else
    echo "Registration failed with status code $HTTP_STATUS. Response: $HTTP_BODY"
    exit 1
  fi
}

# Function to log in and get JWT token
login_user() {
  echo "Logging in with user credentials..."

  LOGIN_RESPONSE=$(curl -s -X POST "$LOGIN_URL" -H "Content-Type: application/json" -d '{
    "username": "'$USERNAME'",
    "password": "'$PASSWORD'"
  }')

  echo "Login response: $LOGIN_RESPONSE"

  JWT_TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token')

  if [[ "$JWT_TOKEN" == "null" || -z "$JWT_TOKEN" ]]; then
    echo "Error: JWT token not received from login."
    exit 1
  fi

  echo "Login successful. JWT token received."
}

# Function to get user details
get_user_details() {
  echo "Fetching user details..."

  RESPONSE=$(curl -s -X GET "$USER_URL?username=$USERNAME" -H "Authorization: Bearer $JWT_TOKEN")

  echo "$RESPONSE" | jq .

  USER_ID=$(echo "$RESPONSE" | jq -r '.ID')

  if [[ "$USER_ID" == "null" || -z "$USER_ID" ]]; then
    echo "Error: Could not retrieve user ID."
    exit 1
  fi

  echo "User ID retrieved: $USER_ID"
}

# Function to update user details
update_user() {
  echo "Updating user details..."

  UPDATE_USER_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$UPDATE_USER_URL/$USER_ID" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{
  "username": "updateduser",
  "mailAddress": "updateduser@example.com",
  "role": "Sales Representative"
  }')

  # Get the HTTP status code (last 3 characters of the response)
  HTTP_STATUS="${UPDATE_USER_RESPONSE: -3}"
  HTTP_BODY="${UPDATE_USER_RESPONSE%???}"

  echo "Update response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: User update failed."
    exit 1
  fi

  echo "User updated successfully."
}


# Function to update user password
update_password() {
  echo "Updating user password..."

  UPDATE_PASSWORD_RESPONSE=$(curl -s -w "%{http_code}" -X POST "$UPDATE_PASSWORD_URL" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{
    "username": "'$USERNAME'",
    "new_password": "'$NEW_PASSWORD'"
  }')

  # Get the HTTP status code (last 3 characters of the response)
  HTTP_STATUS="${UPDATE_PASSWORD_RESPONSE: -3}"
  HTTP_BODY="${UPDATE_PASSWORD_RESPONSE%???}"

  echo "Update password response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: Password update failed."
    exit 1
  fi

  echo "Password updated successfully."
}


# Function to delete user
delete_user() {
  echo "Deleting the user..."

  DELETE_RESPONSE=$(curl -s -w "\n%{http_code}" -X DELETE "$DELETE_USER_URL/$USER_ID" -H "Authorization: Bearer $JWT_TOKEN")

  HTTP_BODY=$(echo "$DELETE_RESPONSE" | sed '$ d')
  HTTP_STATUS=$(echo "$DELETE_RESPONSE" | tail -n1)

  echo "Delete response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: User deletion failed."
    exit 1
  fi

  echo "User deleted successfully."
}

### **🚀 EXECUTION FLOW 🚀**

# Check if user exists
if check_user_exists; then
  echo "Skipping registration since user already exists."
else
  register_user
fi

# Log in
login_user

# Fetch user details
get_user_details

# Update user details
update_user

# Update password
update_password

# Get user details again to confirm updates
get_user_details

# Delete user
delete_user

# Final message
echo "Test script finished successfully!"
