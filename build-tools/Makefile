# Environment file and variables
ENV_FILE := .env

## up_build: Stops and removes all running containers, builds the project and starts docker-compose
build: stop_down_all_containers check_postgres build_user_service build_customer_service build_salestracking_service
	echo "🚀 Building (when required) and starting docker images with environment variables..."
	docker-compose up --build -d
	echo "✅ Docker images built and started!"
	echo "📜 Fetching logs for all services..."
	docker-compose logs --tail=20
	echo "🚀 Runnig Containers:"
	docker ps

## stop_down_all_containers: Stops and removes all running Docker containers (if any exist)
stop_down_all_containers:
	echo "🔍 Checking for running containers..."
	@if [ -n "$$(docker ps -q)" ]; then \
		echo "🛑 Stopping all running Docker containers..."; \
		docker stop $$(docker ps -q); \
		echo "🗑️ Removing all stopped containers..."; \
		docker rm $$(docker ps -aq); \
		echo "✅ All containers stopped and removed."; \
		echo "🛑 Stopping docker-compose..." \
	    docker-compose down \
	    echo "✅ Done!" \
	else \
		echo "⚡ No running containers found. Skipping stop and remove."; \
	fi

check_postgres:
	@if sudo lsof -i :5432 | grep LISTEN; then \
		echo "PostgreSQL is running on port 5432. Stopping it..."; \
		sudo systemctl stop postgresql; \
		sleep 2; \
		if sudo lsof -i :5432 | grep LISTEN; then \
			echo "Failed to stop PostgreSQL!"; \
			exit 1; \
		else \
			echo "PostgreSQL stopped successfully!"; \
		fi \
	else \
		echo "PostgreSQL is not running on port 5432."; \
	fi	


## build_user_service: Builds the user-service binary as a Linux executable
build_user_service:
	echo "1- 🚀 Building user-service binary..."
	@set -a; . ./$(ENV_FILE); set +a; \
	cd ../user-service && $$GOFULLPATH build -o $$USER_SERVICE_BINARY ./cmd/api
	echo "✅ Done! user-service binary..."

## build_customer_service: Builds the customer-service binary as a Linux executable
build_customer_service:
	echo "2- 🚀 Building customer-service binary..."
	@set -a; . ./$(ENV_FILE); set +a; \
	cd ../customer-service && $$GOFULLPATH build -o $$CUSTOMER_SERVICE_BINARY ./cmd/api
	echo "✅ Done! customer-service binary..."

## build_salestracking_service: Builds the salestracking-service binary as a Linux executable
build_salestracking_service:
	echo "3- 🚀 Building salestracking-service binary..."
	@set -a; . ./$(ENV_FILE); set +a; \
	cd ../salestracking-service && $$GOFULLPATH build -o $$SALESTRACKING_SERVICE_BINARY ./cmd/api
	echo "✅ Done! salestracking-service binary..."	

## logs: Shows logs from all services
logs:
	echo "📜 Fetching last 20 logs for all services..."
	docker-compose logs --tail=20 -f

## help: Displays the list of available commands
help:
	@grep -E '^##' $(MAKEFILE_LIST) | sed -e 's/## //'