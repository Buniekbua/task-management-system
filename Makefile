postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root task_managment_db

dropdb:
	docker exec -it postgres16 dropdb task_managment_db

migrateup:
	migrate -path internal/db/migration -database "postgresql://root:password@localhost:5432/task_managment_db?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://root:password@localhost:5432/task_managment_db?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown