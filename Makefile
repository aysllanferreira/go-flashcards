# Variables
APP_NAME=flashcards
DOCKER_COMPOSE_FILE=docker-compose.yml

# Default target: Run the Go project
run:
	@echo "Running $(APP_NAME) application..."
	go run cmd/server/main.go

# Generate Swagger documentation
swagger:
	@echo "Generating Swagger docs..."
	swag init -g cmd/server/main.go

# Start Docker services in the background
docker-up:
	@echo "Starting Docker containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop Docker services
docker-down:
	@echo "Stopping Docker containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Clean Docker resources
docker-clean:
	@echo "Cleaning Docker resources..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --volumes --remove-orphans

# Build Go application (optional, if you want to compile the app)
build:
	@echo "Building the Go application..."
	go build -o $(APP_NAME) cmd/server/main.go

# Full setup (Swagger, Docker, and running the app)
setup: swagger docker-up run
	@echo "Setup completed."

# Help command to list available make targets
help:
	@echo "Available make targets:"
	@echo "  run           - Run the Go application"
	@echo "  swagger       - Generate Swagger documentation"
	@echo "  docker-up     - Start Docker services in the background"
	@echo "  docker-down   - Stop Docker services"
	@echo "  docker-clean  - Clean Docker resources"
	@echo "  build         - Build the Go application"
	@echo "  setup         - Full setup (Swagger, Docker, and running the app)"
	@echo "  help          - Show this help message"
