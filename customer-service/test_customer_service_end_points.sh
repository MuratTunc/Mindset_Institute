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

# Define customer details
CUSTOMERNAME="testcustomer"
MAILADDRESS="testcustomer@example.com"
PASSWORD="TestPassword123"


# Define new parameters
NEW_PASSWORD="NewTestPassword123"
NEW_EMAIL="newmail@example.com"


# Define API URLs
# Read port from .env file
BASE_URL="http://localhost:$CUSTOMER_SERVICE_PORT"
HEALTH_CHECK_URL="$BASE_URL/health"
REGISTER_URL="$BASE_URL/register"
LOGIN_URL="$BASE_URL/login"
CUSTOMER_URL="$BASE_URL/customer"


# (Require JWT authentication)
UPDATE_PASSWORD_URL="$BASE_URL/update-password"
UPDATE_CUSTOMER_URL="$BASE_URL/update-customer"
DEACTIVATE_CUSTOMER_URL="$BASE_URL/deactivate-customer"
ACTIVATE_CUSTOMER_URL="$BASE_URL/activate-customer"
UPDATE_EMAIL_URL="$BASE_URL/update-email"
DELETE_CUSTOMER_URL="$BASE_URL/delete-customer"


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



# Function to check if the customer exists (using the registration endpoint)
register_customer() {
  echo "Test: REGISTER NEW CUSTOMER"
  echo "--------------------------"
  echo "URL:$REGISTER_URL"

  REGISTER_RESPONSE=$(curl -s -w "\n%{http_code}" -X POST "$REGISTER_URL" -H "Content-Type: application/json" -d '{
    "customername": "'$CUSTOMERNAME'",
    "mailAddress": "'$MAILADDRESS'",
    "password": "'$PASSWORD'",
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
login_customer() {
  echo "Test: LOGIN CUSTOMER"
  echo "--------------------------"
  echo "URL:$LOGIN_URL"

  LOGIN_RESPONSE=$(curl -s -X POST "$LOGIN_URL" -H "Content-Type: application/json" -d '{
    "customername": "'$CUSTOMERNAME'",
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


# Function to get customer details
get_customer_details() {
  echo "<------CUSTOMER DETAILS------>"
  echo "--------------------------"
  echo "$CUSTOMERNAME"
  echo "URL:$CUSTOMER_URL?customername=$CUSTOMERNAME"

  RESPONSE=$(curl -s -X GET "$CUSTOMER_URL?customername=$CUSTOMERNAME" -H "Authorization: Bearer $JWT_TOKEN")

  echo "$RESPONSE" | jq .

  CUSTOMER_ID=$(echo "$RESPONSE" | jq -r '.ID')

  if [[ "$CUSTOMER_ID" == "null" || -z "$CUSTOMER_ID" ]]; then
    echo "Error: Could not retrieve customer ID."
    exit 1
  fi

  echo "User ID retrieved: $CUSTOMER_ID"
  echo "--------------------------"
}

# Function to deactivate customer
deactivate_customer() {
  echo "Test: DEACTIVATE CUSTOMER"
  echo "--------------------------"
  echo "CUSTOMER ID: $CUSTOMER_ID"
  echo "URL:$DEACTIVATE_CUSTOMER_URL/$CUSTOMER_ID"

  # Use the DEACTIVATE_CUSTOMER_URL variable and append the customer ID directly
  DEACTIVATE_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$DEACTIVATE_CUSTOMER_URL/$CUSTOMER_ID" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json")

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

# Function to activate customer
activate_customer() {
  echo "Test: ACTIVATE CUSTOMER"
  echo "--------------------------"
  echo "CUSTOMER ID: $CUSTOMER_ID"
  echo "URL:$ACTIVATE_CUSTOMER_URL/$CUSTOMER_ID"

  # Use the ACTIVATE_CUSTOMER_URL variable and append the customer ID directly
  ACTIVATE_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$ACTIVATE_CUSTOMER_URL/$CUSTOMER_ID" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json")

  HTTP_BODY=$(echo "$ACTIVATE_RESPONSE" | sed '$ d')
  HTTP_STATUS=$(echo "$ACTIVATE_RESPONSE" | tail -n1)

  echo "Activate response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: User activation failed."
    exit 1
  fi

  echo "User deactivated successfully."
  echo "--------------------------"
}


# Function to update customer details
update_customer() {
  echo "Test: UPDATE CUSTOMER"
  echo "--------------------------"
  echo "URL:$UPDATE_CUSTOMER_URL/$CUSTOMER_ID"

  UPDATE_CUSTOMER_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$UPDATE_CUSTOMER_URL/$CUSTOMER_ID" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{
  "customername": "updatedcustomer",
  "mailAddress": "updatedcustomer@example.com",
  }')

  # Get the HTTP status code (last 3 characters of the response)
  HTTP_STATUS="${UPDATE_CUSTOMER_RESPONSE: -3}"
  HTTP_BODY="${UPDATE_CUSTOMER_RESPONSE%???}"

  echo "Update response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: User update failed."
    exit 1
  fi

  echo "User updated successfully."
  echo "--------------------------"
}

# Function to update customer password
update_password() {
  echo "Test: UPDATE NEW PASSWORD"
  echo "--------------------------"
  echo "URL:$UPDATE_PASSWORD_URL"

  UPDATE_PASSWORD_RESPONSE=$(curl -s -w "%{http_code}" -X POST "$UPDATE_PASSWORD_URL" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{
    "customername": "'$CUSTOMERNAME'",
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

# Function to update customer email address
update_email() {
  echo "Test: UPDATE EMAIL ADDRESS"
  echo "--------------------------"
  echo "CUSTOMER ID=$CUSTOMER_ID"
  echo "URL:$UPDATE_EMAIL_URL/$CUSTOMER_ID"

  UPDATE_EMAIL_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$UPDATE_EMAIL_URL/$CUSTOMER_ID" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{
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



# Function to delete customer
delete_customer() {
  echo "Test: DELETE CUSTOMER"
  echo "--------------------------"
  echo "CUSTOMER ID=$CUSTOMER_ID"
  echo "URL:$DELETE_CUSTOMER_URL/$CUSTOMER_ID"

  # Perform the DELETE request and capture both status code and response body
  DELETE_RESPONSE=$(curl -s -w "%{http_code}" -X DELETE "$DELETE_CUSTOMER_URL/$CUSTOMER_ID" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json")

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
  CONTAINER_ID=$(docker ps -qf "name=$CUSTOMER_POSTGRES_DB_CONTAINER_NAME")

  # Check if the container exists
  if [ -z "$CONTAINER_ID" ]; then
      echo "Error: No running container found with name '$CONTAINER_NAME'."
      exit 1
  fi

  # Run the query to list all rows in the 'customers' table
  docker exec -i "$CONTAINER_ID" psql -U "$CUSTOMER_POSTGRES_DB_CUSTOMER" -d "$CUSTOMER_POSTGRES_DB_NAME" -c "SELECT * FROM customers;"

}

### **🚀 TEST EXECUTION FLOW 🚀**


health_check


# First Register
register_customer

# Start to test all end points
login_customer
get_customer_details

deactivate_customer
get_customer_details

activate_customer
get_customer_details

update_email
get_customer_details

update_password
get_customer_details

update_customer
get_customer_details

delete_customer
show_databas_table

# Final message
echo "ALL TESTS ARE DONE!!!"
