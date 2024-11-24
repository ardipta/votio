PROJECT_NAME = votio

# Install dependencies using asdf
install:
	asdf install

# Set up the project
setup:
	@echo "Setting up the project..."
	@make install
	@docker-compose build
	@echo "Project setup complete!"

# Docker Compose Commands
build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down

restart:
	@make down
	@make up

clean:
	docker system prune -f

serve:
	@make build
	@make up

