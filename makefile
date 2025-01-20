mysql:
	docker run --name mysql -e MYSQL_ROOT_PASSWORD=password -p 3306:3306 -d mysql

migratefile:
	migrate create -ext sql -dir ./migration -seq migration_name

createdb:
	docker exec -it mysql mysql -u root -p -e "CREATE DATABASE absensi;"

dropdb:
	docker exec -it mysql mysql -u root -p -e "DROP DATABASE IF EXISTS absensi;"

migrateup:
	migrate -path ./migration -database "mysql://root:password@tcp(localhost:3306)/absensi?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

migratedown:
	migrate -path ./migration -database "mysql://root:password@tcp(localhost:3306)/absensi?charset=utf8mb4&parseTime=True&loc=Local" -verbose down

server:
	go run ./cmd/main.go

.PHONY: createdb dropdb migratefile migrateup migratedown server mysql