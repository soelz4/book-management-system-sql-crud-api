# Please Don't Change
SRC_DIR := src
.DEFAULT_GOAL := help
BINARY_NAME := main
BINARY_DIR := bin

help:  ## 💬 This Help Message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Linting and Formatting without Fix
lint: ## 🔎 Lint & Format, will not Fix but Sets Exit Code on Error
	gofmt -l $(SRC_DIR) \
	&& gofmt -d $(SRC_DIR) \
	&& golangci-lint run $(SRC_DIR)/...

# Linting and Formatting with Try to Fix and Modify Code
lint-fix: ## 📜 Lint & Format, will Try to Fix Errors and Modify Code
	gofmt -w $(SRC_DIR) \
	&& golangci-lint run $(SRC_DIR)/...

# Build Binary File
build: ## 🔨 Build Binary File
	go build -o $(BINARY_DIR)/$(BINARY_NAME) $(SRC_DIR)/cmd/main/main.go

# RUN
run: build ## 🏃 Run the Web Server Locally at PORT 9010
	$(BINARY_DIR)/$(BINARY_NAME)

# Resolve Dependencies
init: ## 📥 Download Dependencies From go.mod File
	go mod download

# Clean up Project
clean: ## 🧹 Clean up Project
	go clean
	rm $(BINARY_DIR)/$(BINARY_NAME)

# Docker
IMAGE_REPO := soelz/golang-book-management-system-sql-crud-api
IMAGE_TAG := 0.1
DATABASE_URL := postgres://gorm:1234@db:5432/bank
PostgreSQL_IMAGE := postgres:alpine3.18

# Pull PostgreSQL Docker Image from Docker Hub Registry
postgres: ## 📚 Pull PostgreSQL Docker Image from Docker Hub Registry
	docker pull $(PostgreSQL_IMAGE)

# Create Docker Network
docker-network: ## 🪡 Create Docker Network
	docker network create -d bridge backend

# Build Docker Image
image:  ## 📦 Build Docker Container Image from Dockerfile 
	docker build . --file docker/Dockerfile \
	--tag $(IMAGE_REPO):$(IMAGE_TAG)

# Push Docker Image to Docker Hub Registry
push:  ## 📤 Push Container Image to Registry 
	docker push $(IMAGE_REPO):$(IMAGE_TAG)

# RUN Containers with Docker Compose
compose-up: ## 🧷 Create and Start Containers
	docker compose up --build

# Stop and Remove Containers, Networks
compose-down: ## 🧼 Stop and Remove Containers, Networks
	docker compose down
