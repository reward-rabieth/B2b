DB_URL=postgres://root:root@localhost:5432/B2B?sslmode=disable
createdb:
	 docker exec -it postgres createdb --username=root --owner root B2B

network:
	docker network create shatta-network

postgres:
	docker run --name postgres  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16.2-alpine

postgresNet:
	docker run --name postgres --network shatta-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16.2-alpine

b2b:
	docker run --name b2b-api --network shatta-network -p 7979:7979 -e db-source="postgres://root:root@localhost:5432/B2B?sslmode=disable" b2b:latest

dropdb:
	docker exec -it postgres dropdb B2B

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrationfix:
	migrate -path db/migration -database postgres://root:root@localhost/B2B?sslmode=disable force 1


migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)