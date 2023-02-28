all: postgres run

postgres:
	docker-compose up -d; sleep 5

build: 
	wire .; swag init

run: 
	go run goexercise

down:
	docker-compose down