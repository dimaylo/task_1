DB_DSN := "postgres://postgres:21814046098@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/app/main.go

gen:
	oapi-codegen -config openapi/.openapi.tasks -include-tags tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	oapi-codegen -config openapi/.openapi.users -include-tags users openapi/openapi.yaml > ./internal/web/users/api.gen.go


lint:
	golangci-lint run --out-format=colored-line-number

git:
	git add .
	git commit -m "$(commit)"
	git push
