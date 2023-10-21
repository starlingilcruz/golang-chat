
postgres:
	docker run --name postgresDB -p 5435:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -d postgres

clean:
	docker rm $(docker ps -aq -f name=postgresDB)

# access database container
epo: 
	docker exec -ti postgresDB bash