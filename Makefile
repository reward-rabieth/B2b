DB_URL=postgres://root:root@?dbname=B2B&sslmode=disable
createdb:
	 docker exec -it postgres createdb --username=root --owner root B2B

dropdb:
	docker exec -it postgres dropdb B2B

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)