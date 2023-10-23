
postgres:
	docker run --name postgresDB -p 5435:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -d postgres

rabbitmq:
	docker run -d --name rabbitmqGO -p 15672:15672 -p 5672:5672 -e RABBITMQ_DEFAULT_USER=rabbitmq -e RABBITMQ_DEFAULT_PASS=rabbitmq rabbitmq:3-management

clean:
	- docker stop $$(docker ps -aq)
	- docker rm $$(docker ps -aq -f name=postgresDB)
	- docker rm $$(docker ps -aq -f name=rabbitmqGO)

run:
	- make postgres
	- make rabbitmq

# access database container
epo: 
	docker exec -ti postgresDB bash