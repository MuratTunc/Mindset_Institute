# Environment file and variables
ENV_FILE := .env
WAIT_TIME ?= 5  # Default wait time is 5 seconds, can be overridden

## up_build: Stops and removes all running containers, builds the project and starts docker-compose
build: clean_volumes stop_down_all_containers check_postgres build_user_service build_customer_service build_salestracking_service
	echo "🚀 Building (when required) and starting docker images with environment variables..."
	docker-compose up --build -d
	echo "✅ Docker images built and started!"
	$(MAKE) wait_for_services  # Call the wait block
	echo "✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅"
	echo "📜 Fetching logs for all services..."
	docker-compose logs --tail=20
	echo "🚀 Running Containers:"
	docker ps
	echo "✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅✅"
	$(MAKE) wait_for_integration_test
	$(MAKE) integration_test

## wait_for_services: Waits for the services to initialize by showing animated dots
wait_for_services:
	echo -n "⏳ Waiting for $(WAIT_TIME) seconds to allow services to initialize "
	for i in $$(seq 1 $(WAIT_TIME)); do \
		echo -n "."; \
		sleep 1; \
	done
	echo " ✅"

## wait_for_services: Waits for the services to initialize by showing animated dots
wait_for_integration_test:
	echo -n "⏳ Waiting for $(WAIT_TIME) seconds to allow services to ready "
	for i in $$(seq 1 $(WAIT_TIME)); do \
		echo -n "."; \
		sleep 1; \
	done
	echo " ✅"

## wait_general: Waits before running the next integration test
wait_general:
	echo -n "⏳ Waiting for $(WAIT_TIME) next integration test... "
	for i in $$(seq 1 $(WAIT_TIME)); do \
		echo -n "."; \
		sleep 1; \
	done
	echo " ✅"

## clean_volumes: Removes all volumes (database data) to start fresh
clean_volumes:
	echo "🧹 Cleaning up all database volumes..."
	docker-compose down -v
	echo "✅ All volumes removed!"

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

## integration_test: Runs integration tests for all services
integration_test:
	echo "🔍 Running integration tests..."
	cd ../integration_tests && \
	./integration_user_service.sh && \
	./integration_customer_service.sh && \
	./integration_salestracking_service.sh
	echo "✅ Integration tests completed successfully!"

## migration_up: Runs the migration tool to apply all migrations
migration_up:
	@echo "🚀 Running database migrations..."
	@set -a; . ./$(ENV_FILE); set +a; \
	cd cmd/migration-tool && go run main.go
	echo "✅ Migrations applied successfully!"

## migration_down: Rolls back the most recent migration
migration_down:
	@echo "⏪ Rolling back the most recent migration..."
	@set -a; . ./$(ENV_FILE); set +a; \
	cd cmd/migration-tool && go run main.go down
	echo "✅ Most recent migration rolled back!"

## migration_status: Checks the current migration version
migration_status:
	@echo "🔍 Checking current migration version..."
	@set -a; . ./$(ENV_FILE); set +a; \
	cd cmd/migration-tool && go run main.go status
	echo "✅ Migration version check completed!"
