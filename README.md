Postgress connection example:

postgresql://username:password@host:port/database_name?sslmode=disable

migration:
migrate create -ext sql -dir ./migrations -seq <migration_name>
migrate -path ./migrations -database "postgresql://postgres:admin@localhost:5432/cookie_bomb?sslmode=disable" -verbose up
