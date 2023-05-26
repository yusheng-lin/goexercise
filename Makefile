all: postgres run

postgres:
	docker-compose -f docker-compose.01.yml up -d; sleep 5

build:
	cd src;\
	wire .; swag init

run: 
	cd src;\
	go run goexercise

down:
	docker-compose down