all: postgres run

postgres:
	docker-compose up -d

build: 
	wire .; swag init

run: 
	go run goexercise