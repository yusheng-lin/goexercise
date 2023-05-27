all: postgres run

build:
	cd src;\
	wire .; swag init

run: 
	docker-compose up -d

down:
	docker-compose down