export

IMAGE_NAME = blevels/openweatherapi

# HELP =================================================================================================================
# This will output the help for each task

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

test:
	${DOCKER_RUN} go test -cover ./...

up: ### Run docker-compose command to spin up the docker container
	docker-compose up --build -d && docker-compose logs -f

down: ### Run docker-compose command to shut down the docker container and clean up
	docker-compose down --remove-orphans

logs: ### Run docker-compose command to view logs
	docker-compose logs -f app

build: ### Run docker command to build the image
	docker build -t ${IMAGE_NAME} -f Dockerfile .
