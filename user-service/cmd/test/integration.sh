#!/bin/bash

# Load environment variables from .env file
ENV_FILE="../build-tools/.env"
if [ -f "$ENV_FILE" ]; then
  export $(grep -v '^#' "$ENV_FILE" | xargs)
else
  echo "⚠️ .env file not found at $ENV_FILE"
  exit 1
fi

# Install jq (JSON parsing utility) if not already installed
if ! command -v jq &> /dev/null
then
  echo "jq could not be found, installing..."
  sudo apt-get update
  sudo apt-get install -y jq
fi

# Define user details
USERNAME="testuser"
MAILADDRESS="testuser@example.com"
PASSWORD="TestPassword123"
ROLE="Admin"

# Define new parameters
NEW_PASSWORD="NewTestPassword123"
NEW_EMAIL="newmail@example.com"
NEW_ROLE="MANAGER"


# Define API URLs
# Read port from .env file
BASE_URL="http://localhost:$USER_SERVICE_PORT"
HEALTH_CHECK_URL="$BASE_URL/health"
REGISTER_URL="$BASE_URL/register"
LOGIN_URL="$BASE_URL/login"
USER_URL="$BASE_URL/user"


# (Require JWT authentication)
UPDATE_PASSWORD_URL="$BASE_URL/update-password"
UPDATE_USER_URL="$BASE_URL/update-user"
DEACTIVATE_USER_URL="$BASE_URL/deactivate-user"
ACTIVATE_USER_URL="$BASE_URL/activate-user"
UPDATE_EMAIL_URL="$BASE_URL/update-email"
UPDATE_ROLE_URL="$BASE_URL/update-role"
DELETE_USER_URL="$BASE_URL/delete-user"


health_check() {
  echo "<------HEALTH CHECK------>"
  echo "Checking service health at: $HEALTH_CHECK_URL"

  RESPONSE=$(curl -s -X GET "$HEALTH_CHECK_URL")

  if [[ -z "$RESPONSE" ]]; then
    echo "❌ Error: No response from service!"
    exit 1
  fi

  echo "✅ Health Check Response: $RESPONSE"
  echo "--------------------------"
}



# Function to check if the user exists (using the registration endpoint)
register_user() {
  echo "Test: REGISTER NEW USER"
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
    echo "User already exists...Deleting this user to continue test..."
    delete_user
  else
    echo "Registration failed with status code $HTTP_STATUS. Response: $HTTP_BODY"
    exit 1
  fi

  echo "--------------------------"
}

# Function to log in and get JWT token
login_user() {
  echo "Test: LOGIN USER"
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
  echo "<------USER DETAILS------>"
  echo "--------------------------"
  echo "$USERNAME"
  echo "URL:$USER_URL?username=$USERNAME"

  RESPONSE=$(curl -s -X GET "$USER_URL?username=$USERNAME" -H "Authorization: Bearer $JWT_TOKEN")

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
  echo "Test: DEACTIVATE USER"
  echo "--------------------------"
  echo "USERNAME: $USERNAME"
  echo "URL: $DEACTIVATE_USER_URL"

  # Construct JSON payload
  JSON_PAYLOAD=$(jq -n --arg username "$USERNAME" '{username: $username}')

  # Make the PUT request with JSON body
  DEACTIVATE_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$DEACTIVATE_USER_URL" \
    -H "Authorization: Bearer $JWT_TOKEN" \
    -H "Content-Type: application/json" \
    -d "$JSON_PAYLOAD")

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

# Function to activate user
activate_user() {
  echo "Test: ACTIVATE USER"
  echo "--------------------------"
  echo "USERNAME: $USERNAME"
  echo "URL: $ACTIVATE_USER_URL"

  # Construct JSON payload
  JSON_PAYLOAD=$(jq -n --arg username "$USERNAME" '{username: $username}')

  # Make the PUT request with JSON body
  ACTIVATE_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$ACTIVATE_USER_URL" \
    -H "Authorization: Bearer $JWT_TOKEN" \
    -H "Content-Type: application/json" \
    -d "$JSON_PAYLOAD")

  HTTP_BODY=$(echo "$ACTIVATE_RESPONSE" | sed '$ d')
  HTTP_STATUS=$(echo "$ACTIVATE_RESPONSE" | tail -n1)

  echo "Activate response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: User activation failed."
    exit 1
  fi

  echo "User activated successfully."
  echo "--------------------------"
}



# Function to update user details
update_user() {
  echo "Test: UPDATE USER"
  echo "--------------------------"
  echo "USERNAME: $USERNAME"
  echo "URL: $UPDATE_USER_URL"

  # Construct JSON payload dynamically
  JSON_PAYLOAD=$(jq -n \
    --arg username "$USERNAME" \
    --arg email "$EMAIL" \
    --arg role "$ROLE" \
    '{
      username: $username,
      email: ($email // empty),
      role: ($role // empty)
    }')

  # Make the PUT request with JSON body
  UPDATE_USER_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$UPDATE_USER_URL" \
    -H "Authorization: Bearer $JWT_TOKEN" \
    -H "Content-Type: application/json" \
    -d "$JSON_PAYLOAD")

  # Extract HTTP status and response body
  HTTP_BODY=$(echo "$UPDATE_USER_RESPONSE" | sed '$ d')
  HTTP_STATUS=$(echo "$UPDATE_USER_RESPONSE" | tail -n1)

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
  echo "Test: UPDATE NEW PASSWORD"
  echo "--------------------------"
  echo "URL:$UPDATE_PASSWORD_URL"

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
  echo "--------------------------"
}

# Function to update user email address
update_email() {
  echo "Test: UPDATE EMAIL ADDRESS"
  echo "--------------------------"
  echo "USERNAME=$USERNAME"
  echo "URL: $UPDATE_EMAIL_URL"

  # Construct JSON payload dynamically
  JSON_PAYLOAD=$(jq -n \
    --arg username "$USERNAME" \
    --arg new_email "$NEW_EMAIL" \
    '{
      username: $username,
      new_email: $new_email
    }')

  # Make the PUT request
  UPDATE_EMAIL_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$UPDATE_EMAIL_URL" \
    -H "Authorization: Bearer $JWT_TOKEN" \
    -H "Content-Type: application/json" \
    -d "$JSON_PAYLOAD")

  # Extract HTTP status and response body
  HTTP_BODY=$(echo "$UPDATE_EMAIL_RESPONSE" | sed '$ d')
  HTTP_STATUS=$(echo "$UPDATE_EMAIL_RESPONSE" | tail -n1)

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
  echo "Test: UPDATE USER ROLE"
  echo "--------------------------"
  echo "USER NAME=$USERNAME"
  
  echo "URL:$UPDATE_ROLE_URL"

  UPDATE_ROLE_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$UPDATE_ROLE_URL" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{
    "username": "'$USERNAME'",
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



# Function to delete user by username
delete_user() {
  echo "Test: DELETE USER"
  echo "--------------------------"
  echo "USERNAME=$USERNAME"
  echo "URL:$DELETE_USER_URL"

  # Perform the DELETE request and capture both status code and response body
  DELETE_RESPONSE=$(curl -s -w "%{http_code}" -X DELETE "$DELETE_USER_URL" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{
    "username": "'$USERNAME'"
  }')

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



show_databas_table(){
  
  # Get the container ID using the container name
  CONTAINER_ID=$(docker ps -qf "name=$USER_POSTGRES_DB_CONTAINER_NAME")

  # Check if the container exists
  if [ -z "$CONTAINER_ID" ]; then
      echo "Error: No running container found with name '$CONTAINER_NAME'."
      exit 1
  fi

  # Run the query to list all rows in the 'users' table
  docker exec -i "$CONTAINER_ID" psql -U "$USER_POSTGRES_DB_USER" -d "$USER_POSTGRES_DB_NAME" -c "SELECT * FROM users;"

}

### **🚀 TEST EXECUTION FLOW 🚀**


health_check


# First Register
register_user

# Start to test all end points
login_user
show_databas_table
#get_user_details

deactivate_user
show_databas_table
#get_user_details


activate_user
show_databas_table
#get_user_details

update_email
show_databas_table
#get_user_details

update_password
show_databas_table
#get_user_details

update_role
show_databas_table
#get_user_details

update_user
show_databas_table
#get_user_details


delete_user
show_databas_table

# Final message
echo "ALL TESTS ARE DONE!!!"
