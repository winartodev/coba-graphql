.PHOHY: run mod setup_db remove_db start_db

DB_CONTAINER_NAME = db-graphql-pg
DB_VOLUME_NAME = test-graphql-volume
DB_NAME = db_graphql_pg

run:
	go run ./...

mod: 
	go mod tidy
	go mod download

setup_db:
	docker volume create ${DB_VOLUME_NAME}
	docker run -d \
		--name ${DB_CONTAINER_NAME} \
		-e POSTGRES_PASSWORD=postgres \
		-e PGDATA=/var/lib/postgresql/data/pgdata \
		-e POSTGRES_DB=${DB_NAME}\
		-v ${DB_VOLUME_NAME}:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres:14

remove_db:
	docker container rm -f ${DB_CONTAINER_NAME}
	docker volume rm ${DB_VOLUME_NAME}

start_db: 
	docker start ${DB_CONTAINER_NAME}

test: 
	go test -cover -coverprofile=coverage.out $$(go list ./...)

coverage: 
	make test 
	go tool cover -html=coverage.out
