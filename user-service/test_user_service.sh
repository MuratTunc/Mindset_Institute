#!/bin/bash

# Install jq (JSON parsing utility) if not already installed
if ! command -v jq &> /dev/null
then
  echo "jq could not be found, installing..."
  sudo apt-get update
  sudo apt-get install -y jq
fi

# Define user details
USERNAME="testuser1"
MAILADDRESS="testuser1@example.com"
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
DEACTIVATE_USER_URL="$BASE_URL/deactivate-user"

NEW_EMAIL="newmail@example.com"
NEW_ROLE="MANAGER"

# Function to check if the user exists (using the registration endpoint)
register_user() {
  echo "Test-1: REGISTER NEW USER"
  echo "--------------------------"
  echo "URL:$REGISTER_URL"

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

  if [ "$HTTP_STATUS" -eq 200 ]; then
    echo "User registered successfully!"
  elif [ "$HTTP_STATUS" -eq 409 ]; then
    echo "User already exists."
  else
    echo "Registration failed with status code $HTTP_STATUS. Response: $HTTP_BODY"
    exit 1
  fi

  echo "--------------------------"
}

# Function to log in and get JWT token
login_user() {
  echo "Test-2: LOGIN USER"
  echo "--------------------------"
  echo "URL:$LOGIN_URL"

  LOGIN_RESPONSE=$(curl -s -X POST "$LOGIN_URL" -H "Content-Type: application/json" -d '{
    "username": "'$USERNAME'",
    "password": "'$PASSWORD'"
  }')

  echo "Login response: $LOGIN_RESPONSE"  # Add this line to print the raw response

  JWT_TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token')

  if [[ "$JWT_TOKEN" == "null" || -z "$JWT_TOKEN" ]]; then
    echo "Error: JWT token not received from login."
    exit 1
  fi

  echo "Login successful. JWT token received."
  echo "--------------------------"
  
}


# Function to get user details
get_user_details() {
  echo "Test-3: FETCH USER DETAILS"
  echo "--------------------------"
  echo "URL:$USER_URL?username=updateduser"

  RESPONSE=$(curl -s -X GET "$USER_URL?username=updateduser" -H "Authorization: Bearer $JWT_TOKEN")

  echo "$RESPONSE" | jq .

  USER_ID=$(echo "$RESPONSE" | jq -r '.ID')

  if [[ "$USER_ID" == "null" || -z "$USER_ID" ]]; then
    echo "Error: Could not retrieve user ID."
    exit 1
  fi

  echo "User ID retrieved: $USER_ID"
  echo "--------------------------"
}

# Function to deactivate user
deactivate_user() {
  echo "Test-4: DEACTIVATE USER"
  echo "--------------------------"
  echo "USER ID: $USER_ID"
  echo "URL:$DEACTIVATE_USER_URL/$USER_ID"

  # Use the DEACTIVATE_USER_URL variable and append the user ID directly
  DEACTIVATE_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$DEACTIVATE_USER_URL/$USER_ID" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json")

  HTTP_BODY=$(echo "$DEACTIVATE_RESPONSE" | sed '$ d')
  HTTP_STATUS=$(echo "$DEACTIVATE_RESPONSE" | tail -n1)

  echo "Deactivate response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: User deactivation failed."
    exit 1
  fi

  echo "User deactivated successfully."
  echo "--------------------------"
}


# Function to update user details
update_user() {
  echo "Test-5: UPDATE USER"
  echo "--------------------------"
  echo "URL:$UPDATE_USER_URL/$USER_ID"

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
  echo "--------------------------"
}

# Function to update user password
update_password() {
  echo "Test-6: UPDATE NEW PASSWORD"
  echo "--------------------------"
  echo "URL:$UPDATE_PASSWORD_URL"

  UPDATE_PASSWORD_RESPONSE=$(curl -s -w "%{http_code}" -X POST "$UPDATE_PASSWORD_URL" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{
    "username": "updateduser",
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
  echo "--------------------------"
}

# Function to update user email address
update_email() {
  echo "Test-7: UPDATE EMAIL ADDRESS"
  echo "--------------------------"
  echo "URL:$UPDATE_EMAIL_URL"

  UPDATE_EMAIL_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$UPDATE_EMAIL_URL" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{
    "new_email": "'$NEW_EMAIL'"
  }')

  # Get the HTTP status code (last 3 characters of the response)
  HTTP_STATUS="${UPDATE_EMAIL_RESPONSE: -3}"
  HTTP_BODY="${UPDATE_EMAIL_RESPONSE%???}"

  echo "Update email response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: Email update failed."
    exit 1
  fi

  echo "Email updated successfully."
  echo "--------------------------"
}

# Function to update user role
update_role() {
  echo "Test-8: UPDATE USER ROLE"
  echo "--------------------------"
  echo "URL:$UPDATE_ROLE_URL"

  UPDATE_ROLE_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$UPDATE_ROLE_URL" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{
    "role": "'$NEW_ROLE'"
  }')

  # Get the HTTP status code (last 3 characters of the response)
  HTTP_STATUS="${UPDATE_ROLE_RESPONSE: -3}"
  HTTP_BODY="${UPDATE_ROLE_RESPONSE%???}"

  echo "Update role response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: Role update failed."
    exit 1
  fi

  echo "Role updated successfully."
  echo "--------------------------"
}


# Function to delete user
delete_user() {
  echo "Test-9: DELETE USER"
  echo "--------------------------"
  echo "USER ID=$USER_ID"
  echo "URL:$DELETE_USER_URL/$USER_ID"

  # Perform the DELETE request and capture both status code and response body
  DELETE_RESPONSE=$(curl -s -w "%{http_code}" -X DELETE "$DELETE_USER_URL/$USER_ID" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json")

  # Extract the response body and HTTP status code
  HTTP_STATUS=$(echo "$DELETE_RESPONSE" | tail -n1)  # Extract the last line as the HTTP status code
  HTTP_BODY=$(echo "$DELETE_RESPONSE" | sed '$ d')   # Remove the last line (HTTP status code) to get the body

  echo "Delete response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  # Check if the HTTP status code is 200
  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: User deletion failed."
    exit 1
  fi

  echo "User deleted successfully."
  echo "--------------------------"
}


### **🚀 EXECUTION FLOW 🚀**

# register_user
register_user

# Log in
login_user

# Fetch user details
get_user_details

# Deactivate user
deactivate_user

# Update user details
update_user

# Update password
update_password

# Update update_emai
update_email

# Update ROLE
update_role


# Get user details again to confirm updates
get_user_details

# Delete user
delete_user

# Final message
echo "ALL TESTS ARE DONE!!!"
