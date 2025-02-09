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
NEW_NOTE="This is a new note to append."
UPDATED_NOTE="This is the completely new note."


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
UPDATE_NOTE_URL="$BASE_URL/update-note"
INSERT_NOTE_URL="$BASE_URL/insert-note"
GET_ALL_CUSTOMERS_URL="$BASE_URL/get_all_customer"
ORDER_CUSTOMERS_URL="$BASE_URL/order-customers"
GET_ACTIVATED_CUSTOMERS_URL="$BASE_URL/activated-customers"
GET_LOGGED_IN_CUSTOMERS_URL="$BASE_URL/logged-in-customers"
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


# Function to register new customer to customer-db
register_customer() {
  echo "Test: REGISTER NEW CUSTOMER"
  echo "--------------------------"
  echo "URL:$REGISTER_URL"

  REGISTER_RESPONSE=$(curl -s -w "\n%{http_code}" -X POST "$REGISTER_URL" -H "Content-Type: application/json" -d '{
    "customername": "'$CUSTOMERNAME'",
    "mailAddress": "'$MAILADDRESS'",
    "password": "'$PASSWORD'"
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


# Function to log in
login_customer() {
  echo "Test: LOGIN CUSTOMER"
  echo "--------------------------"
  echo "URL:$LOGIN_URL"

  LOGIN_RESPONSE=$(curl -s -w "%{http_code}" -X POST "$LOGIN_URL" -H "Content-Type: application/json" -d '{
    "customername": "'$CUSTOMERNAME'",
    "password": "'$PASSWORD'"
  }')

  # Extract HTTP status code (last 3 characters of response)
  HTTP_STATUS="${LOGIN_RESPONSE: -3}"
  HTTP_BODY="${LOGIN_RESPONSE%???}"  # Remove the last 3 characters (status code) from the body

  echo "Login response body: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  if [ "$HTTP_STATUS" -eq 200 ]; then
    echo "Login successful."
  elif [ "$HTTP_STATUS" -eq 401 ]; then
    echo "Error: Invalid credentials."
    exit 1
  else
    echo "Login failed with status code $HTTP_STATUS. Response: $HTTP_BODY"
    exit 1
  fi

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

  echo "User activated successfully."
  echo "--------------------------"
}


# Function to update customer details
update_customer() {
  echo "Test: UPDATE CUSTOMER"
  echo "--------------------------"
  echo "URL:$UPDATE_CUSTOMER_URL/$CUSTOMER_ID"

  UPDATE_CUSTOMER_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$UPDATE_CUSTOMER_URL/$CUSTOMER_ID" -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{
    "customername": "updatedcustomer",
    "mailAddress": "updatedcustomer@example.com"
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

  # Update email using PUT method
  UPDATE_EMAIL_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$UPDATE_EMAIL_URL/$CUSTOMER_ID" -H "Content-Type: application/json" -d '{
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

# Function to update the note for a customer
update_note() {
  echo "Test: UPDATE NOTE"
  echo "--------------------------"
  echo "Customer Name: $CUSTOMERNAME"
  echo "URL: $UPDATE_NOTE_URL"

  # Send PUT request to the /update-note endpoint with the customer name and updated note
  UPDATE_NOTE_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$UPDATE_NOTE_URL" -H "Content-Type: application/json" -d '{
    "customername": "'$CUSTOMERNAME'",
    "note": "'"$UPDATED_NOTE"'"
  }')

  # Extract the HTTP status code and response body
  HTTP_STATUS="${UPDATE_NOTE_RESPONSE: -3}"
  HTTP_BODY="${UPDATE_NOTE_RESPONSE%???}"

  echo "Update Note response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  # Check if the HTTP status code is 200
  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: Note update failed."
    exit 1
  fi

  echo "Note updated successfully."
  echo "--------------------------"
}



# Function to insert or append a note for a customer
insert_note() {
  echo "Test: INSERT NOTE"
  echo "--------------------------"
  echo "Customer Name: $CUSTOMERNAME"
  echo "URL: $INSERT_NOTE_URL"

  # Send PUT request to the /insert-note endpoint with the customer name and new note
  INSERT_NOTE_RESPONSE=$(curl -s -w "%{http_code}" -X PUT "$INSERT_NOTE_URL" -H "Content-Type: application/json" -d '{
    "customername": "'$CUSTOMERNAME'",
    "new_note": "'"$NEW_NOTE"'"
  }')

  # Extract the HTTP status code and response body
  HTTP_STATUS="${INSERT_NOTE_RESPONSE: -3}"
  HTTP_BODY="${INSERT_NOTE_RESPONSE%???}"

  echo "Insert Note response: $HTTP_BODY"
  echo "HTTP Status Code: $HTTP_STATUS"

  # Check if the HTTP status code is 200
  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error: Note insertion failed."
    exit 1
  fi

  echo "Note inserted successfully."
  echo "--------------------------"
}




# Function to get all customers
get_all_customers() {
  echo "--------------------------"
  echo "Test: GET ALL CUSTOMERS"
  echo "--------------------------"
  echo "URL: $GET_ALL_CUSTOMERS_URL"

  # Send GET request to the endpoint
  RESPONSE=$(curl -s -w "\nHTTP Status Code: %{http_code}\n" -X GET "$GET_ALL_CUSTOMERS_URL" -H "Content-Type: application/json")

  # Display the response
  echo -e "$RESPONSE"

  # Check if the request was successful (HTTP 200)
  if [[ "$RESPONSE" == *"HTTP Status Code: 200"* ]]; then
    echo "✅ Successfully retrieved all customers."
  else
    echo "❌ Failed to fetch customers."
  fi
}


# Function to test ordering customers
order_customers() {
  echo "--------------------------"
  echo "Test: ORDER CUSTOMERS"
  echo "--------------------------"
  
  # Test ordering by created_at (default)
  echo "URL: $ORDER_CUSTOMERS_URL"
  RESPONSE=$(curl -s -w "\nHTTP Status Code: %{http_code}\n" -X GET "$ORDER_CUSTOMERS_URL")
  echo -e "$RESPONSE"
  if [[ "$RESPONSE" == *"HTTP Status Code: 200"* ]]; then
    echo "✅ Customers ordered by created_at."
  else
    echo "❌ Failed to order customers by created_at."
  fi
  
  # Test ordering by customername
  echo "URL: $ORDER_CUSTOMERS_URL?order_by=customername"
  RESPONSE=$(curl -s -w "\nHTTP Status Code: %{http_code}\n" -X GET "$ORDER_CUSTOMERS_URL?order_by=customername")
  echo -e "$RESPONSE"
  if [[ "$RESPONSE" == *"HTTP Status Code: 200"* ]]; then
    echo "✅ Customers ordered by customername."
  else
    echo "❌ Failed to order customers by customername."
  fi
  
  # Test ordering by updated_at
  echo "URL: $ORDER_CUSTOMERS_URL?order_by=updated_at"
  RESPONSE=$(curl -s -w "\nHTTP Status Code: %{http_code}\n" -X GET "$ORDER_CUSTOMERS_URL?order_by=updated_at")
  echo -e "$RESPONSE"
  if [[ "$RESPONSE" == *"HTTP Status Code: 200"* ]]; then
    echo "✅ Customers ordered by updated_at."
  else
    echo "❌ Failed to order customers by updated_at."
  fi
  
  # Test with an invalid 'order_by' parameter (should default)
  echo "URL: $ORDER_CUSTOMERS_URL?order_by=invalid_field"
  RESPONSE=$(curl -s -w "\nHTTP Status Code: %{http_code}\n" -X GET "$ORDER_CUSTOMERS_URL?order_by=invalid_field")
  echo -e "$RESPONSE"
  if [[ "$RESPONSE" == *"HTTP Status Code: 200"* ]]; then
    echo "✅ Customers ordered with invalid 'order_by' field (default applied)."
  else
    echo "❌ Failed to apply default ordering."
  fi
}

# Function to get activated customer names
get_activated_customers() {
  echo "--------------------------"
  echo "Test: GET ACTIVATED CUSTOMERS"
  echo "--------------------------"
  echo "URL: $GET_ACTIVATED_CUSTOMERS_URL"

  # Send GET request to the endpoint
  RESPONSE=$(curl -s -w "\nHTTP Status Code: %{http_code}\n" -X GET "$GET_ACTIVATED_CUSTOMERS_URL")

  # Display the response
  echo -e "$RESPONSE"

  # Check if the request was successful (HTTP 200)
  if [[ "$RESPONSE" == *"HTTP Status Code: 200"* ]]; then
    echo "✅ Successfully retrieved activated customer names."
  else
    echo "❌ Failed to fetch activated customer names."
  fi
}


# Function to get all logged-in customers
get_logged_in_customers() {
  echo "--------------------------"
  echo "Test: GET LOGGED-IN CUSTOMERS"
  echo "--------------------------"
  echo "URL: $GET_LOGGED_IN_CUSTOMERS_URL"

  # Send GET request to the endpoint
  RESPONSE=$(curl -s -w "\nHTTP Status Code: %{http_code}\n" -X GET "$GET_LOGGED_IN_CUSTOMERS_URL" -H "Content-Type: application/json")

  # Display the response
  echo -e "$RESPONSE"

  # Check if the request was successful (HTTP 200)
  if [[ "$RESPONSE" == *"HTTP Status Code: 200"* ]]; then
    echo "✅ Successfully retrieved logged-in customers."
  else
    echo "❌ Failed to fetch logged-in customers."
  fi
}



# Function to delete customer
delete_customer() {
  echo "Test: DELETE CUSTOMER"
  echo "--------------------------"
  echo "CUSTOMER ID=$CUSTOMER_ID"
  echo "URL:$DELETE_CUSTOMER_URL/$CUSTOMER_ID"

  # Perform the DELETE request and capture both status code and response body
  DELETE_RESPONSE=$(curl -s -w "%{http_code}" -X DELETE "$DELETE_CUSTOMER_URL/$CUSTOMER_ID" -H "Content-Type: application/json")

  # Extract the response body and HTTP status code
  HTTP_STATUS="${DELETE_RESPONSE: -3}"
  HTTP_BODY="${DELETE_RESPONSE%???}"

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






show_database_table(){
  
  # Get the container ID using the container name
  CONTAINER_ID=$(docker ps -qf "name=$CUSTOMER_POSTGRES_DB_CONTAINER_NAME")

  # Check if the container exists
  if [ -z "$CONTAINER_ID" ]; then
      echo "Error: No running container found with name '$CONTAINER_NAME'."
      exit 1
  fi

  # Run the query to list all rows in the 'customers' table
  docker exec -i "$CONTAINER_ID" psql -U "$CUSTOMER_POSTGRES_DB_USER" -d "$CUSTOMER_POSTGRES_DB_NAME" -c "SELECT * FROM customers;"

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

insert_note
get_customer_details

update_note
get_customer_details

update_customer
get_customer_details

order_customers
get_activated_customers
get_logged_in_customers

delete_customer
show_database_table

# Final message
echo "ALL TESTS ARE DONE!!!"
