all: postgres

postgres:
	docker-compose up -d

#remove: 
#	docker-compose down