
# rebuild everything, args=""
build:
	-docker-compose up --build $(args)

# run the compose.yml
up:
	-docker-compose up
